/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/12/21 10:55 PM
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
	// ClientField 客户端日志
	ClientField = meta.NewField("span.kind", "client")
)

// PayloadUnaryClientInterceptor 一元拦截器，用于记录客户端端请求和响应
func PayloadUnaryClientInterceptor(zapLogger *xzap.ZapLogger, options ...Option) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, resp interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		o := evaluateClientOpt(options...)
		startTime := time.Now()
		newCtx := newClientLoggerCaller(ctx, zapLogger, method, startTime)
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
		o.messageFunc(newCtx, "info", level, err, o.durationFunc(time.Since(startTime)))

		return err
	}
}

// PayloadStreamClientInterceptor 流拦截器，用于记录客户端请求和响应
func PayloadStreamClientInterceptor(zapLogger *xzap.ZapLogger, options ...Option) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (_ grpc.ClientStream, err error) {
		o := evaluateClientOpt(options...)
		startTime := time.Now()
		newCtx := newClientLoggerCaller(ctx, zapLogger, method, startTime)
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
		o.messageFunc(newCtx, "info", level, err, o.durationFunc(time.Since(startTime)))

		return &wrappedClientStream{
			ClientStream:   clientStream,
			wrappedContext: newCtx,
		}, nil
	}
}

func evaluateClientOpt(opts ...Option) *option {
	optCopy := &option{}
	*optCopy = *defaultOptions
	optCopy.levelFunc = DefaultClientCodeToLevel
	for _, o := range opts {
		o(optCopy)
	}
	return optCopy
}

func newClientLoggerCaller(ctx context.Context, logger *xzap.ZapLogger, methodString string, start time.Time) context.Context {
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

	return ToContext(ctx, logger.With(append(fields, clientLoggerFields(methodString)...)...))
}

func clientLoggerFields(methodString string) []meta.Field {
	service := path.Dir(methodString)[1:]
	method := path.Base(methodString)
	return []meta.Field{
		GrpcSystemField,
		ClientField,
		meta.NewField("grpc.service", service),
		meta.NewField("grpc.method", method),
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
		AddFields(l.Context(), protoMessageToFields(m, "grpc.request")...)
	}

	return err
}

// RecvMsg 接收消息
func (l *wrappedClientStream) RecvMsg(m interface{}) error {
	err := l.ClientStream.RecvMsg(m)
	if err == nil {
		AddFields(l.Context(), protoMessageToFields(m, "grpc.response")...)
	}

	return err
}

// Context 返回封装的上下文, 用于覆盖 grpc.ServerStream.Context()
func (l *wrappedClientStream) Context() context.Context {
	return l.wrappedContext
}
