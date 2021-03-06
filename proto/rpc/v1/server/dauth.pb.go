// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/rpc/v1/server/dauth.proto

package server

import (
	fmt "fmt"
	math "math"

	v1 "github.com/airbloc/airbloc-go/proto/rpc/v1"
	proto "github.com/golang/protobuf/proto"
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

type DAuthRequest struct {
	CollectionId         string   `protobuf:"bytes,1,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
	IdentityHash         []byte   `protobuf:"bytes,2,opt,name=identityHash,proto3" json:"identityHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DAuthRequest) Reset()         { *m = DAuthRequest{} }
func (m *DAuthRequest) String() string { return proto.CompactTextString(m) }
func (*DAuthRequest) ProtoMessage()    {}
func (*DAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66cc8b72c8d7fd76, []int{0}
}

func (m *DAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DAuthRequest.Unmarshal(m, b)
}
func (m *DAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DAuthRequest.Marshal(b, m, deterministic)
}
func (m *DAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DAuthRequest.Merge(m, src)
}
func (m *DAuthRequest) XXX_Size() int {
	return xxx_messageInfo_DAuthRequest.Size(m)
}
func (m *DAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DAuthRequest proto.InternalMessageInfo

func (m *DAuthRequest) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *DAuthRequest) GetIdentityHash() []byte {
	if m != nil {
		return m.IdentityHash
	}
	return nil
}

func init() {
	proto.RegisterType((*DAuthRequest)(nil), "airbloc.rpc.v1.DAuthRequest")
}

func init() { proto.RegisterFile("proto/rpc/v1/server/dauth.proto", fileDescriptor_66cc8b72c8d7fd76) }

var fileDescriptor_66cc8b72c8d7fd76 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2,
	0x4f, 0x49, 0x2c, 0x2d, 0xc9, 0xd0, 0x03, 0xcb, 0x08, 0xf1, 0x25, 0x66, 0x16, 0x25, 0xe5, 0xe4,
	0x27, 0xeb, 0x15, 0x15, 0x24, 0xeb, 0x95, 0x19, 0x4a, 0x49, 0xa0, 0x68, 0x28, 0xa9, 0x2c, 0x48,
	0x2d, 0x86, 0xa8, 0x54, 0x0a, 0xe3, 0xe2, 0x71, 0x71, 0x2c, 0x2d, 0xc9, 0x08, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xc9, 0x49, 0x4d, 0x2e, 0xc9, 0xcc,
	0xcf, 0xf3, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x11, 0x03, 0xa9, 0xc9, 0x4c,
	0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0xf4, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x09, 0x42, 0x11, 0x33, 0x6a, 0x61, 0xe4, 0x62, 0x05, 0x1b, 0x2c, 0x64, 0xcb, 0xc5, 0xea, 0x98,
	0x93, 0x93, 0x5f, 0x2e, 0x24, 0xa3, 0x87, 0xea, 0x2a, 0x3d, 0x64, 0x8b, 0xa5, 0xc4, 0xd0, 0x65,
	0x83, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x84, 0x6c, 0xb8, 0x58, 0x5c, 0x52, 0xf3, 0x2a, 0xc9, 0xd3,
	0xed, 0x64, 0x1c, 0x65, 0x98, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f,
	0x55, 0x03, 0xa3, 0x75, 0xd3, 0xf3, 0xf5, 0xb1, 0x84, 0x64, 0x12, 0x1b, 0x58, 0xd0, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xf3, 0x50, 0xee, 0xea, 0x67, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DAuthClient is the client API for DAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DAuthClient interface {
	Allow(ctx context.Context, in *DAuthRequest, opts ...grpc.CallOption) (*v1.Result, error)
	Deny(ctx context.Context, in *DAuthRequest, opts ...grpc.CallOption) (*v1.Result, error)
}

type dAuthClient struct {
	cc *grpc.ClientConn
}

func NewDAuthClient(cc *grpc.ClientConn) DAuthClient {
	return &dAuthClient{cc}
}

func (c *dAuthClient) Allow(ctx context.Context, in *DAuthRequest, opts ...grpc.CallOption) (*v1.Result, error) {
	out := new(v1.Result)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.DAuth/Allow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dAuthClient) Deny(ctx context.Context, in *DAuthRequest, opts ...grpc.CallOption) (*v1.Result, error) {
	out := new(v1.Result)
	err := c.cc.Invoke(ctx, "/airbloc.rpc.v1.DAuth/Deny", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DAuthServer is the server API for DAuth service.
type DAuthServer interface {
	Allow(context.Context, *DAuthRequest) (*v1.Result, error)
	Deny(context.Context, *DAuthRequest) (*v1.Result, error)
}

func RegisterDAuthServer(s *grpc.Server, srv DAuthServer) {
	s.RegisterService(&_DAuth_serviceDesc, srv)
}

func _DAuth_Allow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DAuthServer).Allow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.DAuth/Allow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DAuthServer).Allow(ctx, req.(*DAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DAuth_Deny_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DAuthServer).Deny(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/airbloc.rpc.v1.DAuth/Deny",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DAuthServer).Deny(ctx, req.(*DAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DAuth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "airbloc.rpc.v1.DAuth",
	HandlerType: (*DAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allow",
			Handler:    _DAuth_Allow_Handler,
		},
		{
			MethodName: "Deny",
			Handler:    _DAuth_Deny_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/v1/server/dauth.proto",
}
