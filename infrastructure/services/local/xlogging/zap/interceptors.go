package xzap

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/weitrue/Seckill/infrastructure/services/local/xlogging"
)

var (
	// GatewaySystemField 日志系统域
	GatewaySystemField = zap.String("system", "gateway")

	// GrpcSystemField 日志系统域
	GrpcSystemField = zap.String("system", "grpc")

	// ServerField 服务端日志
	ServerField = zap.String("span.kind", "server")

	// ClientField 客户端日志
	ClientField = zap.String("span.kind", "client")

	// JsonPbMarshaller 序列化protobuf消息
	JsonPbMarshaller xlogging.JsonPbMarshaler = &jsonpb.Marshaler{}
)

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w BodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func GinZapLogger(zapLogger *ZapLogger, msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		requestBody, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)
		bodyLogWriter := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		start := time.Now()

		c.Next()

		responseBody := bodyLogWriter.body.Bytes()
		log := WithContext(c.Request.Context())
		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				log.Error(e)
			}
		} else {
			latency := float64(time.Now().Sub(start).Nanoseconds() / 1000000.0)
			fields := []zapcore.Field{
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("function", c.HandlerName()),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("token", c.Request.Header.Get("session_id")),
				zap.String("content-type", c.Request.Header.Get("Content-Type")),
				zap.Float64("latency", latency),
				zap.String("request", string(requestBody)),
				zap.String("response", string(responseBody)),
			}
			logger.Info(msg, fields...)
		}
	}
}

// PayloadUnaryServerInterceptor 一元服务器拦截器，用于记录服务端请求和响应
func PayloadUnaryServerInterceptor(zapLogger *ZapLogger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		o := evaluateServerOpt(zapLogger.Opts)
		startTime := time.Now()
		newCtx := newServerLoggerCaller(ctx, zapLogger.Logger, info.FullMethod, startTime, o.timestampFormat)
		defer func() {
			if r := recover(); r != nil {
				err = recoverFrom(newCtx, r, o.recoveryHandlerFunc)
			}
		}()

		resp, err := handler(newCtx, req)
		if !o.shouldLog(info.FullMethod, err) {
			return resp, err
		}

		fields := protoMessageToFields(req, "grpc.request")
		if err == nil {
			fields = append(fields, protoMessageToFields(resp, "grpc.response")...)
		}
		code := o.codeFunc(err)
		level := o.levelFunc(code)
		o.messageFunc(newCtx, "info", level, err, append(fields, o.durationFunc(time.Since(startTime))))

		return resp, err
	}
}

// PayloadStreamServerInterceptor 流拦截器，用于记录服务端请求和响应
func PayloadStreamServerInterceptor(zapLogger *ZapLogger) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		o := evaluateServerOpt(zapLogger.Opts)
		startTime := time.Now()
		ctx := newServerLoggerCaller(stream.Context(), zapLogger.Logger, info.FullMethod, startTime, o.timestampFormat)
		wrapped := &wrappedServerStream{ServerStream: stream, wrappedContext: ctx}
		defer func() {
			if r := recover(); r != nil {
				err = recoverFrom(stream.Context(), r, o.recoveryHandlerFunc)
			}
		}()

		err = handler(srv, wrapped)
		if !o.shouldLog(info.FullMethod, err) {
			return err
		}

		code := o.codeFunc(err)
		level := o.levelFunc(code)
		o.messageFunc(ctx, "info", level, err, []zap.Field{o.durationFunc(time.Since(startTime))})

		return err
	}
}

// wrappedServerStream 包装后的服务端流对象
type wrappedServerStream struct {
	grpc.ServerStream
	wrappedContext context.Context
}

// SendMsg 发送消息
func (l *wrappedServerStream) SendMsg(m interface{}) error {
	err := l.ServerStream.SendMsg(m)
	if err == nil {
		addFields(l.Context(), protoMessageToFields(m, "grpc.response")...)
	}

	return err
}

// RecvMsg 接收消息
func (l *wrappedServerStream) RecvMsg(m interface{}) error {
	err := l.ServerStream.RecvMsg(m)
	if err == nil {
		addFields(l.Context(), protoMessageToFields(m, "grpc.request")...)
	}

	return err
}

// Context 返回封装的上下文
func (l *wrappedServerStream) Context() context.Context {
	return l.wrappedContext
}

// PayloadUnaryClientInterceptor 一元拦截器，用于记录客户端端请求和响应
func PayloadUnaryClientInterceptor(zapLogger *ZapLogger) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, resp interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		o := evaluateClientOpt(zapLogger.Opts)
		startTime := time.Now()
		newCtx := newClientLoggerCaller(ctx, zapLogger.Logger, method, startTime, o.timestampFormat)
		defer func() {
			if r := recover(); r != nil {
				err = recoverFrom(newCtx, r, o.recoveryHandlerFunc)
			}
		}()

		err = invoker(newCtx, method, req, resp, cc, opts...)
		if !o.shouldLog(method, err) {
			return err
		}

		fields := protoMessageToFields(req, "grpc.request")
		if err == nil {
			fields = append(fields, protoMessageToFields(resp, "grpc.response")...)
		}

		code := o.codeFunc(err)
		level := o.levelFunc(code)
		duration := o.durationFunc(time.Since(startTime))
		fields = append(fields, duration)
		o.messageFunc(newCtx, "info", level, err, fields)

		return err
	}
}

// PayloadStreamClientInterceptor 流拦截器，用于记录客户端请求和响应
func PayloadStreamClientInterceptor(zapLogger *ZapLogger) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (_ grpc.ClientStream, err error) {
		o := evaluateClientOpt(zapLogger.Opts)
		startTime := time.Now()
		newCtx := newClientLoggerCaller(ctx, logger, method, startTime, o.timestampFormat)
		defer func() {
			if r := recover(); r != nil {
				err = recoverFrom(newCtx, r, o.recoveryHandlerFunc)
			}
		}()

		clientStream, err := streamer(newCtx, desc, cc, method, opts...)
		if !o.shouldLog(method, err) {
			if err != nil {
				return nil, err
			}

			return &wrappedClientStream{
				ClientStream:   clientStream,
				wrappedContext: newCtx,
			}, err
		}

		code := o.codeFunc(err)
		level := o.levelFunc(code)
		var fields []zap.Field
		duration := o.durationFunc(time.Since(startTime))
		fields = append(fields, duration)
		o.messageFunc(newCtx, "info", level, err, fields)

		return &wrappedClientStream{
			ClientStream:   clientStream,
			wrappedContext: newCtx,
		}, nil
	}
}

// wrappedClientStream 包装后的客户端流对象
type wrappedClientStream struct {
	grpc.ClientStream
	wrappedContext context.Context
}

// SendMsg 发送消息
func (l *wrappedClientStream) SendMsg(m interface{}) error {
	err := l.ClientStream.SendMsg(m)
	if err == nil {
		addFields(l.Context(), protoMessageToFields(m, "grpc.request")...)
	}

	return err
}

// RecvMsg 接收消息
func (l *wrappedClientStream) RecvMsg(m interface{}) error {
	err := l.ClientStream.RecvMsg(m)
	if err == nil {
		addFields(l.Context(), protoMessageToFields(m, "grpc.response")...)
	}

	return err
}

// Context 返回封装的上下文, 用于覆盖 grpc.ServerStream.Context()
func (l *wrappedClientStream) Context() context.Context {
	return l.wrappedContext
}

type protoMessageObject struct {
	pb proto.Message
}

// MarshalLogObject 序列化成日志对象
func (j *protoMessageObject) MarshalLogObject(oe zapcore.ObjectEncoder) error {
	return oe.AddReflected("content", j)
}

// MarshalJSON 序列化成json
func (j *protoMessageObject) MarshalJSON() ([]byte, error) {
	b := &bytes.Buffer{}
	if err := JsonPbMarshaller.Marshal(b, j.pb); err != nil {
		return nil, fmt.Errorf("jsonpb serializer failed: %v", err)
	}

	return b.Bytes(), nil
}

// protoMessageToFields 将message序列化成json，并写入存储
func protoMessageToFields(pbMsg interface{}, key string) []zap.Field {
	var fields []zap.Field
	if p, ok := pbMsg.(proto.Message); ok {
		fields = append(fields, zap.Object(key, &protoMessageObject{pb: p}))
	}

	return fields
}

func recoverFrom(ctx context.Context, p interface{}, r RecoveryHandlerFuncContext) error {
	if r == nil {
		return status.Errorf(codes.Internal, "%v", p)
	}
	return r(ctx, p)
}
