// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal.proto

package Internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InternalRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalRequest) Reset()         { *m = InternalRequest{} }
func (m *InternalRequest) String() string { return proto.CompactTextString(m) }
func (*InternalRequest) ProtoMessage()    {}
func (*InternalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_a53dc05fbff8e250, []int{0}
}
func (m *InternalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalRequest.Unmarshal(m, b)
}
func (m *InternalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalRequest.Marshal(b, m, deterministic)
}
func (dst *InternalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalRequest.Merge(dst, src)
}
func (m *InternalRequest) XXX_Size() int {
	return xxx_messageInfo_InternalRequest.Size(m)
}
func (m *InternalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InternalRequest proto.InternalMessageInfo

func (m *InternalRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type InternalResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalResponse) Reset()         { *m = InternalResponse{} }
func (m *InternalResponse) String() string { return proto.CompactTextString(m) }
func (*InternalResponse) ProtoMessage()    {}
func (*InternalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_a53dc05fbff8e250, []int{1}
}
func (m *InternalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalResponse.Unmarshal(m, b)
}
func (m *InternalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalResponse.Marshal(b, m, deterministic)
}
func (dst *InternalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalResponse.Merge(dst, src)
}
func (m *InternalResponse) XXX_Size() int {
	return xxx_messageInfo_InternalResponse.Size(m)
}
func (m *InternalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InternalResponse proto.InternalMessageInfo

func (m *InternalResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*InternalRequest)(nil), "Internal.InternalRequest")
	proto.RegisterType((*InternalResponse)(nil), "Internal.InternalResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InternalServiceClient is the client API for InternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InternalServiceClient interface {
	Message(ctx context.Context, in *InternalRequest, opts ...grpc.CallOption) (*InternalResponse, error)
}

type internalServiceClient struct {
	cc *grpc.ClientConn
}

func NewInternalServiceClient(cc *grpc.ClientConn) InternalServiceClient {
	return &internalServiceClient{cc}
}

func (c *internalServiceClient) Message(ctx context.Context, in *InternalRequest, opts ...grpc.CallOption) (*InternalResponse, error) {
	out := new(InternalResponse)
	err := c.cc.Invoke(ctx, "/Internal.InternalService/Message", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalServiceServer is the server API for InternalService service.
type InternalServiceServer interface {
	Message(context.Context, *InternalRequest) (*InternalResponse, error)
}

func RegisterInternalServiceServer(s *grpc.Server, srv InternalServiceServer) {
	s.RegisterService(&_InternalService_serviceDesc, srv)
}

func _InternalService_Message_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Message(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Internal.InternalService/Message",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Message(ctx, req.(*InternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Internal.InternalService",
	HandlerType: (*InternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Message",
			Handler:    _InternalService_Message_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal.proto",
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor_internal_a53dc05fbff8e250) }

var fileDescriptor_internal_a53dc05fbff8e250 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xf0, 0x84, 0xf2, 0x95,
	0xb4, 0xb9, 0xf8, 0x61, 0xec, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x09, 0x2e, 0xf6,
	0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57,
	0x49, 0x87, 0x4b, 0x00, 0xa1, 0xb8, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0xb7, 0x6a, 0xa3, 0x60,
	0x84, 0xd1, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x0e, 0x5c, 0xec, 0xbe, 0x10, 0x59,
	0x21, 0x49, 0x3d, 0x98, 0xa4, 0x1e, 0x9a, 0x03, 0xa4, 0xa4, 0xb0, 0x49, 0x41, 0xac, 0x4b, 0x62,
	0x03, 0x7b, 0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x89, 0x9d, 0x1b, 0xc2, 0xd2, 0x00, 0x00,
	0x00,
}
