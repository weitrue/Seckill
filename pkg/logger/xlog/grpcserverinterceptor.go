/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/12/21 10:34 PM
 * Description:
 **/

package xlog

import (
	"context"
	"path"
	"time"

	"github.com/weitrue/Seckill/pkg/logger/meta"
	"github.com/weitrue/Seckill/pkg/logger/xzap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var (

	// GrpcSystemField 日志系统域
	GrpcSystemField = meta.NewField("system", "grpc")

	// ServerField 服务端日志
	ServerField = meta.NewField("span.kind", "server")
)

// PayloadUnaryServerInterceptor 一元服务器拦截器，用于记录服务端请求和响应
func PayloadUnaryServerInterceptor(zapLogger *xzap.ZapLogger, opts ...Option) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		o := evaluateServerOpt(opts...)
		startTime := time.Now()
		newCtx := newServerLoggerCaller(ctx, zapLogger, info.FullMethod, startTime)
		defer func() {
			if r := recover(); r != nil {
				err = recoverFrom(newCtx, r, o.recoveryHandlerFunc)
			}
		}()

		resp, err := handler(newCtx, req)
		if !o.shouldLog(info.FullMethod, err) {
			return resp, err
		}

		meta := protoMessageToFields(req, "grpc.request")
		if err == nil {
			meta = append(meta, protoMessageToFields(resp, "grpc.response")...)
		}
		code := o.codeFunc(err)
		level := o.levelFunc(code)
		meta = append(meta, o.durationFunc(time.Since(startTime)))
		o.messageFunc(newCtx, "info", level, err, meta...)

		return resp, err
	}
}

// PayloadStreamServerInterceptor 流拦截器，用于记录服务端请求和响应
func PayloadStreamServerInterceptor(zapLogger *xzap.ZapLogger, opts ...Option) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		o := evaluateServerOpt(opts...)
		startTime := time.Now()
		ctx := newServerLoggerCaller(stream.Context(), zapLogger, info.FullMethod, startTime)
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
		o.messageFunc(ctx, "info", level, err, o.durationFunc(time.Since(startTime)))

		return err
	}
}

func evaluateServerOpt(opts ...Option) *option {
	optCopy := &option{}
	*optCopy = *defaultOptions
	for _, o := range opts {
		o(optCopy)
	}

	return optCopy
}

func newServerLoggerCaller(ctx context.Context, zapLogger *xzap.ZapLogger, methodString string, start time.Time) context.Context {
	var fields []meta.Field
	fields = append(fields, meta.NewField("grpc.start_time", start.Format(xzap.DefaultTimeLayout)))
	if d, ok := ctx.Deadline(); ok {
		fields = append(fields, meta.NewField("grpc.request.deadline", d.Format(xzap.DefaultTimeLayout)))
	}

	if p, ok := peer.FromContext(ctx); ok {
		fields = append(fields, meta.NewField("grpc.address", p.Addr.String()))
	}

	// tr, ok := ctx.Value(tracespec.TracingKey).(tracespec.Trace)
	// if ok {
	// 	fields = append(fields, meta.NewField("trace", tr.TraceId()))
	// 	fields = append(fields, meta.NewField("span", tr.SpanId()))
	// }

	// token, ok := jwt.FromContext(ctx)
	// if ok {
	// 	fields = append(fields, meta.NewField("user_id", token.UserId))
	// }

	return ToContext(ctx, zapLogger.With(append(fields, serverCallFields(methodString)...)...))
}

func serverCallFields(methodString string) []meta.Field {
	service := path.Dir(methodString)[1:]
	method := path.Base(methodString)
	return []meta.Field{
		GrpcSystemField,
		ServerField,
		meta.NewField("grpc.service", service),
		meta.NewField("grpc.method", method),
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
		AddFields(l.Context(), protoMessageToFields(m, "grpc.response")...)
	}

	return err
}

// RecvMsg 接收消息
func (l *wrappedServerStream) RecvMsg(m interface{}) error {
	err := l.ServerStream.RecvMsg(m)
	if err == nil {
		AddFields(l.Context(), protoMessageToFields(m, "grpc.request")...)
	}

	return err
}

// Context 返回封装的上下文
func (l *wrappedServerStream) Context() context.Context {
	return l.wrappedContext
}
