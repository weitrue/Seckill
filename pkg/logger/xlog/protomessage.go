/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2022/12/21 11:01 PM
 * Description:
 **/

package xlog

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/weitrue/Seckill/pkg/logger/meta"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// JsonPbMarshaller 序列化protobuf消息
	JsonPbMarshaller JsonPbMarshaler = &jsonpb.Marshaler{}
)

// JsonPbMarshaler 序列化protobuf消息
type JsonPbMarshaler interface {
	Marshal(out io.Writer, pb proto.Message) error
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
func protoMessageToFields(pbMsg any, key string) []meta.Field {
	var fields []meta.Field
	if p, ok := pbMsg.(proto.Message); ok {
		fields = append(fields, meta.NewField(key, &protoMessageObject{pb: p}))
	}

	return fields
}

func recoverFrom(ctx context.Context, p any, r RecoveryHandlerFuncContext) error {
	if r == nil {
		return status.Errorf(codes.Internal, "%v", p)
	}
	return r(ctx, p)
}
