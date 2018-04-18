// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

/*
Package event is a generated protocol buffer package.

It is generated from these files:
	event.proto

It has these top-level messages:
	ListenRequest
	Event
*/
package event

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

type ListenRequest struct {
}

func (m *ListenRequest) Reset()                    { *m = ListenRequest{} }
func (m *ListenRequest) String() string            { return proto.CompactTextString(m) }
func (*ListenRequest) ProtoMessage()               {}
func (*ListenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Event struct {
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ListenRequest)(nil), "event.ListenRequest")
	proto.RegisterType((*Event)(nil), "event.Event")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Service service

type ServiceClient interface {
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (Service_ListenClient, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (Service_ListenClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Service_serviceDesc.Streams[0], c.cc, "/event.Service/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_ListenClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type serviceListenClient struct {
	grpc.ClientStream
}

func (x *serviceListenClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Service service

type ServiceServer interface {
	Listen(*ListenRequest, Service_ListenServer) error
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).Listen(m, &serviceListenServer{stream})
}

type Service_ListenServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type serviceListenServer struct {
	grpc.ServerStream
}

func (x *serviceListenServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _Service_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "event.proto",
}

func init() { proto.RegisterFile("event.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x4b, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xf8, 0xb9, 0x78, 0x7d,
	0x32, 0x8b, 0x4b, 0x52, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0xd8, 0xb9, 0x58,
	0x5d, 0x41, 0x32, 0x46, 0xd6, 0x5c, 0xec, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x06,
	0x5c, 0x6c, 0x10, 0x45, 0x42, 0x22, 0x7a, 0x10, 0x33, 0x50, 0xf4, 0x48, 0xf1, 0x40, 0x45, 0xc1,
	0x1a, 0x95, 0x18, 0x0c, 0x18, 0x93, 0xd8, 0xc0, 0x96, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x7c, 0x5a, 0x80, 0xad, 0x73, 0x00, 0x00, 0x00,
}
