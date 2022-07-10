// Code generated by protoc-gen-go. DO NOT EDIT.
// source: canarysvc.proto

package canary

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

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1092e078f25a3f2, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type GetReply struct {
	Err                  string   `protobuf:"bytes,1,opt,name=Err,proto3" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}
func (*GetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1092e078f25a3f2, []int{1}
}

func (m *GetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReply.Unmarshal(m, b)
}
func (m *GetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReply.Marshal(b, m, deterministic)
}
func (m *GetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReply.Merge(m, src)
}
func (m *GetReply) XXX_Size() int {
	return xxx_messageInfo_GetReply.Size(m)
}
func (m *GetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReply proto.InternalMessageInfo

func (m *GetReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ServiceStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceStatusRequest) Reset()         { *m = ServiceStatusRequest{} }
func (m *ServiceStatusRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceStatusRequest) ProtoMessage()    {}
func (*ServiceStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1092e078f25a3f2, []int{2}
}

func (m *ServiceStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceStatusRequest.Unmarshal(m, b)
}
func (m *ServiceStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceStatusRequest.Marshal(b, m, deterministic)
}
func (m *ServiceStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceStatusRequest.Merge(m, src)
}
func (m *ServiceStatusRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceStatusRequest.Size(m)
}
func (m *ServiceStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceStatusRequest proto.InternalMessageInfo

type ServiceStatusReply struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceStatusReply) Reset()         { *m = ServiceStatusReply{} }
func (m *ServiceStatusReply) String() string { return proto.CompactTextString(m) }
func (*ServiceStatusReply) ProtoMessage()    {}
func (*ServiceStatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1092e078f25a3f2, []int{3}
}

func (m *ServiceStatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceStatusReply.Unmarshal(m, b)
}
func (m *ServiceStatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceStatusReply.Marshal(b, m, deterministic)
}
func (m *ServiceStatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceStatusReply.Merge(m, src)
}
func (m *ServiceStatusReply) XXX_Size() int {
	return xxx_messageInfo_ServiceStatusReply.Size(m)
}
func (m *ServiceStatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceStatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceStatusReply proto.InternalMessageInfo

func (m *ServiceStatusReply) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ServiceStatusReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "canary.GetRequest")
	proto.RegisterType((*GetReply)(nil), "canary.GetReply")
	proto.RegisterType((*ServiceStatusRequest)(nil), "canary.ServiceStatusRequest")
	proto.RegisterType((*ServiceStatusReply)(nil), "canary.ServiceStatusReply")
}

func init() { proto.RegisterFile("canarysvc.proto", fileDescriptor_d1092e078f25a3f2) }

var fileDescriptor_d1092e078f25a3f2 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4e, 0xcc, 0x4b,
	0x2c, 0xaa, 0x2c, 0x2e, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x08, 0x28,
	0xf1, 0x70, 0x71, 0xb9, 0xa7, 0x96, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0xc9, 0x70,
	0x71, 0x80, 0x79, 0x05, 0x39, 0x95, 0x42, 0x02, 0x5c, 0xcc, 0xae, 0x45, 0x45, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x92, 0x18, 0x97, 0x48, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72,
	0x6a, 0x70, 0x49, 0x62, 0x49, 0x69, 0x31, 0x4c, 0x97, 0x15, 0x97, 0x10, 0x9a, 0x38, 0x48, 0xbf,
	0x10, 0x17, 0x4b, 0x72, 0x7e, 0x4a, 0x2a, 0xd8, 0x00, 0xe6, 0x20, 0x30, 0x1b, 0x64, 0x66, 0x6a,
	0x51, 0x91, 0x04, 0x13, 0xc4, 0xcc, 0xd4, 0xa2, 0x22, 0xa3, 0x16, 0x46, 0x2e, 0x36, 0x67, 0xb0,
	0x53, 0x84, 0x74, 0xb9, 0x98, 0xdd, 0x53, 0x4b, 0x84, 0x84, 0xf4, 0x20, 0x4e, 0xd3, 0x43, 0xb8,
	0x4b, 0x4a, 0x00, 0x45, 0xac, 0x20, 0xa7, 0x52, 0x89, 0x41, 0xc8, 0x9b, 0x8b, 0x17, 0xc5, 0x56,
	0x21, 0x19, 0x98, 0x22, 0x6c, 0x8e, 0x94, 0x92, 0xc2, 0x21, 0x0b, 0x36, 0x2c, 0x89, 0x0d, 0x1c,
	0x2a, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x49, 0x6f, 0xfb, 0x30, 0x28, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CanaryClient is the client API for Canary service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CanaryClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error)
}

type canaryClient struct {
	cc *grpc.ClientConn
}

func NewCanaryClient(cc *grpc.ClientConn) CanaryClient {
	return &canaryClient{cc}
}

func (c *canaryClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/canary.Canary/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *canaryClient) ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error) {
	out := new(ServiceStatusReply)
	err := c.cc.Invoke(ctx, "/canary.Canary/ServiceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CanaryServer is the server API for Canary service.
type CanaryServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
	ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error)
}

// UnimplementedCanaryServer can be embedded to have forward compatible implementations.
type UnimplementedCanaryServer struct {
}

func (*UnimplementedCanaryServer) Get(ctx context.Context, req *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCanaryServer) ServiceStatus(ctx context.Context, req *ServiceStatusRequest) (*ServiceStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceStatus not implemented")
}

func RegisterCanaryServer(s *grpc.Server, srv CanaryServer) {
	s.RegisterService(&_Canary_serviceDesc, srv)
}

func _Canary_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanaryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/canary.Canary/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanaryServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Canary_ServiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanaryServer).ServiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/canary.Canary/ServiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanaryServer).ServiceStatus(ctx, req.(*ServiceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Canary_serviceDesc = grpc.ServiceDesc{
	ServiceName: "canary.Canary",
	HandlerType: (*CanaryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Canary_Get_Handler,
		},
		{
			MethodName: "ServiceStatus",
			Handler:    _Canary_ServiceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "canarysvc.proto",
}