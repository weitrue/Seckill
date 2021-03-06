// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application/api/rpc/activity.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Goods struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Img                  string   `protobuf:"bytes,3,opt,name=img,proto3" json:"img,omitempty"`
	Price                string   `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	ActivityPrice        string   `protobuf:"bytes,5,opt,name=activity_price,json=activityPrice,proto3" json:"activity_price,omitempty"`
	ActivityStock        int32    `protobuf:"varint,6,opt,name=activity_stock,json=activityStock,proto3" json:"activity_stock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Goods) Reset()         { *m = Goods{} }
func (m *Goods) String() string { return proto.CompactTextString(m) }
func (*Goods) ProtoMessage()    {}
func (*Goods) Descriptor() ([]byte, []int) {
	return fileDescriptor_8133469beee19a97, []int{0}
}

func (m *Goods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Goods.Unmarshal(m, b)
}
func (m *Goods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Goods.Marshal(b, m, deterministic)
}
func (m *Goods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Goods.Merge(m, src)
}
func (m *Goods) XXX_Size() int {
	return xxx_messageInfo_Goods.Size(m)
}
func (m *Goods) XXX_DiscardUnknown() {
	xxx_messageInfo_Goods.DiscardUnknown(m)
}

var xxx_messageInfo_Goods proto.InternalMessageInfo

func (m *Goods) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Goods) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Goods) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *Goods) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Goods) GetActivityPrice() string {
	if m != nil {
		return m.ActivityPrice
	}
	return ""
}

func (m *Goods) GetActivityStock() int32 {
	if m != nil {
		return m.ActivityStock
	}
	return 0
}

type Activity struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TopicId              int32    `protobuf:"varint,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	StartTime            int64    `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              int64    `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Limit                int32    `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	GoodsList            []*Goods `protobuf:"bytes,6,rep,name=goods_list,json=goodsList,proto3" json:"goods_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Activity) Reset()         { *m = Activity{} }
func (m *Activity) String() string { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()    {}
func (*Activity) Descriptor() ([]byte, []int) {
	return fileDescriptor_8133469beee19a97, []int{1}
}

func (m *Activity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Activity.Unmarshal(m, b)
}
func (m *Activity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Activity.Marshal(b, m, deterministic)
}
func (m *Activity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Activity.Merge(m, src)
}
func (m *Activity) XXX_Size() int {
	return xxx_messageInfo_Activity.Size(m)
}
func (m *Activity) XXX_DiscardUnknown() {
	xxx_messageInfo_Activity.DiscardUnknown(m)
}

var xxx_messageInfo_Activity proto.InternalMessageInfo

func (m *Activity) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Activity) GetTopicId() int32 {
	if m != nil {
		return m.TopicId
	}
	return 0
}

func (m *Activity) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Activity) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Activity) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Activity) GetGoodsList() []*Goods {
	if m != nil {
		return m.GoodsList
	}
	return nil
}

type Topic struct {
	//
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Banner               string   `protobuf:"bytes,4,opt,name=banner,proto3" json:"banner,omitempty"`
	StartTime            int64    `protobuf:"varint,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              int64    `protobuf:"varint,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_8133469beee19a97, []int{2}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Topic) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Topic) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Topic) GetBanner() string {
	if m != nil {
		return m.Banner
	}
	return ""
}

func (m *Topic) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Topic) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type Response struct {
	// ??????
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8133469beee19a97, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Goods)(nil), "rpc.Goods")
	proto.RegisterType((*Activity)(nil), "rpc.Activity")
	proto.RegisterType((*Topic)(nil), "rpc.Topic")
	proto.RegisterType((*Response)(nil), "rpc.Response")
}

func init() {
	proto.RegisterFile("application/api/rpc/activity.proto", fileDescriptor_8133469beee19a97)
}

var fileDescriptor_8133469beee19a97 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x8a, 0xd4, 0x30,
	0x14, 0xa6, 0xd3, 0x49, 0x77, 0xe6, 0x8c, 0x3b, 0x4a, 0x18, 0xa4, 0x0a, 0xc2, 0x50, 0x10, 0x66,
	0x6f, 0xa6, 0xb2, 0x3e, 0x81, 0x78, 0x21, 0x82, 0xe0, 0x12, 0xf7, 0xbe, 0x74, 0x93, 0xec, 0x70,
	0xb0, 0x4d, 0x42, 0x13, 0x04, 0x9f, 0xc3, 0x27, 0xf0, 0x0d, 0x7c, 0x07, 0x5f, 0x4c, 0x7a, 0x92,
	0x16, 0xdd, 0x41, 0xf1, 0xee, 0x7c, 0x3f, 0xc9, 0xf9, 0xf2, 0xb5, 0x50, 0xb5, 0xce, 0x75, 0x28,
	0xdb, 0x80, 0xd6, 0xd4, 0xad, 0xc3, 0x7a, 0x70, 0xb2, 0x6e, 0x65, 0xc0, 0x2f, 0x18, 0xbe, 0x1e,
	0xdd, 0x60, 0x83, 0xe5, 0xf9, 0xe0, 0x64, 0xf5, 0x3d, 0x03, 0xf6, 0xce, 0x5a, 0xe5, 0xf9, 0x16,
	0x16, 0xa8, 0xca, 0x6c, 0x9f, 0x1d, 0x98, 0x58, 0xa0, 0xe2, 0x1c, 0x96, 0x4a, 0x7b, 0x59, 0x2e,
	0xf6, 0xd9, 0x61, 0x2d, 0x68, 0xe6, 0x4f, 0x20, 0xc7, 0xfe, 0x54, 0xe6, 0x44, 0x8d, 0x23, 0xdf,
	0x01, 0x73, 0x03, 0x4a, 0x5d, 0x2e, 0x89, 0x8b, 0x80, 0xbf, 0x84, 0xed, 0xb4, 0xac, 0x89, 0x32,
	0x23, 0xf9, 0x72, 0x62, 0x6f, 0xce, 0x6c, 0x3e, 0x58, 0xf9, 0xb9, 0x2c, 0x68, 0xfd, 0x6c, 0xfb,
	0x34, 0x92, 0xd5, 0x8f, 0x0c, 0x56, 0x6f, 0x12, 0x73, 0x16, 0xf3, 0x19, 0xac, 0x82, 0x75, 0x28,
	0x1b, 0x54, 0x14, 0x95, 0x89, 0x0b, 0xc2, 0xef, 0x15, 0x7f, 0x01, 0xe0, 0x43, 0x3b, 0x84, 0x26,
	0x60, 0xaf, 0x29, 0x74, 0x2e, 0xd6, 0xc4, 0xdc, 0x62, 0xaf, 0xc7, 0x93, 0xda, 0xa8, 0x28, 0x2e,
	0x49, 0xbc, 0xd0, 0x46, 0x91, 0xb4, 0x03, 0xd6, 0x61, 0x8f, 0x81, 0x62, 0x33, 0x11, 0x01, 0xbf,
	0x02, 0x38, 0x8d, 0x55, 0x35, 0x1d, 0xfa, 0x50, 0x16, 0xfb, 0xfc, 0xb0, 0xb9, 0x86, 0xe3, 0xe0,
	0xe4, 0x91, 0x1a, 0x14, 0x6b, 0x52, 0x3f, 0xa0, 0x0f, 0xd5, 0xb7, 0x0c, 0xd8, 0xed, 0x18, 0xe3,
	0x2c, 0xef, 0x0e, 0x58, 0xc0, 0xd0, 0xe9, 0xd4, 0x6b, 0x04, 0x73, 0xd9, 0xf9, 0x6f, 0x65, 0x3f,
	0x85, 0xe2, 0xae, 0x35, 0x46, 0x0f, 0xa9, 0xdb, 0x84, 0x1e, 0x3c, 0x8b, 0xfd, 0xeb, 0x59, 0xc5,
	0x1f, 0xcf, 0xaa, 0x5e, 0xc1, 0x4a, 0x68, 0xef, 0xac, 0xf1, 0xb4, 0x51, 0x5a, 0xa5, 0x53, 0x32,
	0x9a, 0xc7, 0xcf, 0xdb, 0xfb, 0x53, 0x4a, 0x36, 0x8e, 0xd7, 0x3f, 0x33, 0xd8, 0x4c, 0xd5, 0x8b,
	0x9b, 0xb7, 0xfc, 0x08, 0xdb, 0x09, 0x7e, 0x34, 0x1d, 0x1a, 0xcd, 0x2f, 0xa9, 0x80, 0x89, 0x7c,
	0x1e, 0xe1, 0xbc, 0xa5, 0x86, 0xc7, 0xb3, 0xff, 0xfe, 0xfe, 0x3f, 0x0e, 0x1c, 0x60, 0x43, 0xbd,
	0xa5, 0xdb, 0x63, 0xbd, 0xc4, 0x3c, 0x74, 0x5e, 0xc1, 0xa3, 0xe8, 0x4c, 0xf7, 0xfe, 0xdd, 0x7a,
	0x57, 0xd0, 0x0f, 0xff, 0xfa, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0xbe, 0x3b, 0x05, 0x16,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ActivityRPCClient is the client API for ActivityRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActivityRPCClient interface {
	ActivityOnline(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error)
	ActivityOffline(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error)
	TopicOnline(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Response, error)
	TopicOffline(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Response, error)
}

type activityRPCClient struct {
	cc *grpc.ClientConn
}

func NewActivityRPCClient(cc *grpc.ClientConn) ActivityRPCClient {
	return &activityRPCClient{cc}
}

func (c *activityRPCClient) ActivityOnline(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.ActivityRPC/ActivityOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityRPCClient) ActivityOffline(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.ActivityRPC/ActivityOffline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityRPCClient) TopicOnline(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.ActivityRPC/TopicOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityRPCClient) TopicOffline(ctx context.Context, in *Topic, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.ActivityRPC/TopicOffline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityRPCServer is the server API for ActivityRPC service.
type ActivityRPCServer interface {
	ActivityOnline(context.Context, *Activity) (*Response, error)
	ActivityOffline(context.Context, *Activity) (*Response, error)
	TopicOnline(context.Context, *Topic) (*Response, error)
	TopicOffline(context.Context, *Topic) (*Response, error)
}

// UnimplementedActivityRPCServer can be embedded to have forward compatible implementations.
type UnimplementedActivityRPCServer struct {
}

func (*UnimplementedActivityRPCServer) ActivityOnline(ctx context.Context, req *Activity) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityOnline not implemented")
}
func (*UnimplementedActivityRPCServer) ActivityOffline(ctx context.Context, req *Activity) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityOffline not implemented")
}
func (*UnimplementedActivityRPCServer) TopicOnline(ctx context.Context, req *Topic) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopicOnline not implemented")
}
func (*UnimplementedActivityRPCServer) TopicOffline(ctx context.Context, req *Topic) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopicOffline not implemented")
}

func RegisterActivityRPCServer(s *grpc.Server, srv ActivityRPCServer) {
	s.RegisterService(&_ActivityRPC_serviceDesc, srv)
}

func _ActivityRPC_ActivityOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Activity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityRPCServer).ActivityOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ActivityRPC/ActivityOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityRPCServer).ActivityOnline(ctx, req.(*Activity))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityRPC_ActivityOffline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Activity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityRPCServer).ActivityOffline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ActivityRPC/ActivityOffline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityRPCServer).ActivityOffline(ctx, req.(*Activity))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityRPC_TopicOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityRPCServer).TopicOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ActivityRPC/TopicOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityRPCServer).TopicOnline(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityRPC_TopicOffline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Topic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityRPCServer).TopicOffline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.ActivityRPC/TopicOffline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityRPCServer).TopicOffline(ctx, req.(*Topic))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActivityRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.ActivityRPC",
	HandlerType: (*ActivityRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ActivityOnline",
			Handler:    _ActivityRPC_ActivityOnline_Handler,
		},
		{
			MethodName: "ActivityOffline",
			Handler:    _ActivityRPC_ActivityOffline_Handler,
		},
		{
			MethodName: "TopicOnline",
			Handler:    _ActivityRPC_TopicOnline_Handler,
		},
		{
			MethodName: "TopicOffline",
			Handler:    _ActivityRPC_TopicOffline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application/api/rpc/activity.proto",
}
