// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

/*
Package lnrpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
	SendManyRequest
	SendManyResponse
	NewAddressRequest
	NewAddressResponse
	ConnectPeerRequest
	ConnectPeerResponse
*/
package lnrpc

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
const _ = proto.ProtoPackageIsVersion1

type SendManyRequest struct {
	AddrToAmount map[string]int64 `protobuf:"bytes,1,rep,name=AddrToAmount" json:"AddrToAmount,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *SendManyRequest) Reset()                    { *m = SendManyRequest{} }
func (m *SendManyRequest) String() string            { return proto.CompactTextString(m) }
func (*SendManyRequest) ProtoMessage()               {}
func (*SendManyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SendManyRequest) GetAddrToAmount() map[string]int64 {
	if m != nil {
		return m.AddrToAmount
	}
	return nil
}

type SendManyResponse struct {
	Txid string `protobuf:"bytes,1,opt,name=txid" json:"txid,omitempty"`
}

func (m *SendManyResponse) Reset()                    { *m = SendManyResponse{} }
func (m *SendManyResponse) String() string            { return proto.CompactTextString(m) }
func (*SendManyResponse) ProtoMessage()               {}
func (*SendManyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NewAddressRequest struct {
}

func (m *NewAddressRequest) Reset()                    { *m = NewAddressRequest{} }
func (m *NewAddressRequest) String() string            { return proto.CompactTextString(m) }
func (*NewAddressRequest) ProtoMessage()               {}
func (*NewAddressRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type NewAddressResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *NewAddressResponse) Reset()                    { *m = NewAddressResponse{} }
func (m *NewAddressResponse) String() string            { return proto.CompactTextString(m) }
func (*NewAddressResponse) ProtoMessage()               {}
func (*NewAddressResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ConnectPeerRequest struct {
	IdAtHost string `protobuf:"bytes,1,opt,name=idAtHost" json:"idAtHost,omitempty"`
}

func (m *ConnectPeerRequest) Reset()                    { *m = ConnectPeerRequest{} }
func (m *ConnectPeerRequest) String() string            { return proto.CompactTextString(m) }
func (*ConnectPeerRequest) ProtoMessage()               {}
func (*ConnectPeerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ConnectPeerResponse struct {
	LnID []byte `protobuf:"bytes,1,opt,name=lnID,proto3" json:"lnID,omitempty"`
}

func (m *ConnectPeerResponse) Reset()                    { *m = ConnectPeerResponse{} }
func (m *ConnectPeerResponse) String() string            { return proto.CompactTextString(m) }
func (*ConnectPeerResponse) ProtoMessage()               {}
func (*ConnectPeerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*SendManyRequest)(nil), "lnrpc.SendManyRequest")
	proto.RegisterType((*SendManyResponse)(nil), "lnrpc.SendManyResponse")
	proto.RegisterType((*NewAddressRequest)(nil), "lnrpc.NewAddressRequest")
	proto.RegisterType((*NewAddressResponse)(nil), "lnrpc.NewAddressResponse")
	proto.RegisterType((*ConnectPeerRequest)(nil), "lnrpc.ConnectPeerRequest")
	proto.RegisterType((*ConnectPeerResponse)(nil), "lnrpc.ConnectPeerResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Lightning service

type LightningClient interface {
	SendMany(ctx context.Context, in *SendManyRequest, opts ...grpc.CallOption) (*SendManyResponse, error)
	NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error)
	ConnectPeer(ctx context.Context, in *ConnectPeerRequest, opts ...grpc.CallOption) (*ConnectPeerResponse, error)
}

type lightningClient struct {
	cc *grpc.ClientConn
}

func NewLightningClient(cc *grpc.ClientConn) LightningClient {
	return &lightningClient{cc}
}

func (c *lightningClient) SendMany(ctx context.Context, in *SendManyRequest, opts ...grpc.CallOption) (*SendManyResponse, error) {
	out := new(SendManyResponse)
	err := grpc.Invoke(ctx, "/lnrpc.Lightning/SendMany", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightningClient) NewAddress(ctx context.Context, in *NewAddressRequest, opts ...grpc.CallOption) (*NewAddressResponse, error) {
	out := new(NewAddressResponse)
	err := grpc.Invoke(ctx, "/lnrpc.Lightning/NewAddress", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lightningClient) ConnectPeer(ctx context.Context, in *ConnectPeerRequest, opts ...grpc.CallOption) (*ConnectPeerResponse, error) {
	out := new(ConnectPeerResponse)
	err := grpc.Invoke(ctx, "/lnrpc.Lightning/ConnectPeer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Lightning service

type LightningServer interface {
	SendMany(context.Context, *SendManyRequest) (*SendManyResponse, error)
	NewAddress(context.Context, *NewAddressRequest) (*NewAddressResponse, error)
	ConnectPeer(context.Context, *ConnectPeerRequest) (*ConnectPeerResponse, error)
}

func RegisterLightningServer(s *grpc.Server, srv LightningServer) {
	s.RegisterService(&_Lightning_serviceDesc, srv)
}

func _Lightning_SendMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServer).SendMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lnrpc.Lightning/SendMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServer).SendMany(ctx, req.(*SendManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lightning_NewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServer).NewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lnrpc.Lightning/NewAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServer).NewAddress(ctx, req.(*NewAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lightning_ConnectPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LightningServer).ConnectPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lnrpc.Lightning/ConnectPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LightningServer).ConnectPeer(ctx, req.(*ConnectPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lightning_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lnrpc.Lightning",
	HandlerType: (*LightningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMany",
			Handler:    _Lightning_SendMany_Handler,
		},
		{
			MethodName: "NewAddress",
			Handler:    _Lightning_NewAddress_Handler,
		},
		{
			MethodName: "ConnectPeer",
			Handler:    _Lightning_ConnectPeer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x4e, 0xb3, 0x40,
	0x10, 0x85, 0x43, 0xfb, 0xf7, 0xb7, 0x1d, 0x30, 0x6d, 0xb7, 0x89, 0x22, 0x57, 0x0d, 0x46, 0xc3,
	0x15, 0x17, 0xed, 0x8d, 0xd1, 0xc4, 0x84, 0x58, 0x13, 0x4d, 0xd4, 0x18, 0xf5, 0x05, 0xb0, 0x4c,
	0x2a, 0x11, 0x67, 0x91, 0x5d, 0x54, 0x5e, 0xc0, 0xf7, 0xf3, 0x8d, 0x44, 0x58, 0x02, 0x0a, 0x97,
	0x73, 0xe6, 0x3b, 0x87, 0x33, 0x64, 0x61, 0x94, 0xc4, 0x6b, 0x37, 0x4e, 0xb8, 0xe4, 0x6c, 0x10,
	0x51, 0x3e, 0xd8, 0x9f, 0x1a, 0x8c, 0xef, 0x91, 0x82, 0x6b, 0x9f, 0xb2, 0x3b, 0x7c, 0x4d, 0x51,
	0x48, 0x76, 0x0a, 0x86, 0x17, 0x04, 0xc9, 0x03, 0xf7, 0x5e, 0x78, 0x4a, 0xd2, 0xd4, 0xe6, 0x7d,
	0x47, 0x5f, 0x38, 0x6e, 0xe1, 0x70, 0xff, 0xd0, 0x6e, 0x13, 0x3d, 0x27, 0x99, 0x64, 0xd6, 0x12,
	0xa6, 0x2d, 0x91, 0xe9, 0xd0, 0x7f, 0xc6, 0x2c, 0xcf, 0xd2, 0x9c, 0x11, 0xdb, 0x86, 0xc1, 0x9b,
	0x1f, 0xa5, 0x68, 0xf6, 0xf2, 0xb1, 0x7f, 0xdc, 0x3b, 0xd2, 0xec, 0x39, 0x4c, 0xea, 0x64, 0x11,
	0x73, 0x12, 0xc8, 0x0c, 0xf8, 0x27, 0x3f, 0xc2, 0xa0, 0x34, 0xd9, 0x33, 0x98, 0xde, 0xe0, 0xfb,
	0x4f, 0x32, 0x0a, 0xa1, 0xbe, 0x6e, 0x1f, 0x00, 0x6b, 0x8a, 0xca, 0x38, 0x86, 0x2d, 0xbf, 0x94,
	0x94, 0xf7, 0x10, 0xd8, 0x19, 0x27, 0xc2, 0xb5, 0xbc, 0x45, 0x4c, 0xaa, 0x43, 0x27, 0x30, 0x0c,
	0x03, 0x4f, 0x5e, 0x70, 0x21, 0x15, 0xb7, 0x0f, 0xb3, 0x5f, 0x5c, 0x5d, 0x24, 0xa2, 0xcb, 0x55,
	0x01, 0x19, 0x8b, 0x2f, 0x0d, 0x46, 0x57, 0xe1, 0xe6, 0x49, 0x52, 0x48, 0x1b, 0x76, 0x02, 0xc3,
	0xaa, 0x38, 0xdb, 0xe9, 0xfe, 0x47, 0xd6, 0x6e, 0x4b, 0x57, 0xc1, 0x1e, 0x40, 0x5d, 0x9f, 0x99,
	0x0a, 0x6b, 0x9d, 0x69, 0xed, 0x75, 0x6c, 0x54, 0xc4, 0x0a, 0xf4, 0x46, 0x65, 0x56, 0x91, 0xed,
	0x73, 0x2d, 0xab, 0x6b, 0x55, 0xa6, 0x3c, 0xfe, 0x2f, 0x5e, 0xc5, 0xf2, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x15, 0x2e, 0xcf, 0x06, 0x22, 0x02, 0x00, 0x00,
}
