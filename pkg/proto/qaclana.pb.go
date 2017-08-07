// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/proto/qaclana.proto

/*
Package qaclana is a generated protocol buffer package.

It is generated from these files:
	pkg/proto/qaclana.proto

It has these top-level messages:
	Request
	MultiValue
	SystemState
	Empty
*/
package qaclana

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

type State int32

const (
	State_DISABLED   State = 0
	State_PERMISSIVE State = 1
	State_ENFORCING  State = 2
)

var State_name = map[int32]string{
	0: "DISABLED",
	1: "PERMISSIVE",
	2: "ENFORCING",
}
var State_value = map[string]int32{
	"DISABLED":   0,
	"PERMISSIVE": 1,
	"ENFORCING":  2,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}
func (State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Request struct {
	ClientAddress   string                 `protobuf:"bytes,1,opt,name=clientAddress" json:"clientAddress,omitempty"`
	Id              string                 `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Headers         map[string]*MultiValue `protobuf:"bytes,3,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	QueryParameters map[string]*MultiValue `protobuf:"bytes,4,rep,name=queryParameters" json:"queryParameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	PostParameters  map[string]*MultiValue `protobuf:"bytes,5,rep,name=postParameters" json:"postParameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetClientAddress() string {
	if m != nil {
		return m.ClientAddress
	}
	return ""
}

func (m *Request) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Request) GetHeaders() map[string]*MultiValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Request) GetQueryParameters() map[string]*MultiValue {
	if m != nil {
		return m.QueryParameters
	}
	return nil
}

func (m *Request) GetPostParameters() map[string]*MultiValue {
	if m != nil {
		return m.PostParameters
	}
	return nil
}

type MultiValue struct {
	Value []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (m *MultiValue) Reset()                    { *m = MultiValue{} }
func (m *MultiValue) String() string            { return proto.CompactTextString(m) }
func (*MultiValue) ProtoMessage()               {}
func (*MultiValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MultiValue) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type SystemState struct {
	State State `protobuf:"varint,1,opt,name=state,enum=State" json:"state,omitempty"`
}

func (m *SystemState) Reset()                    { *m = SystemState{} }
func (m *SystemState) String() string            { return proto.CompactTextString(m) }
func (*SystemState) ProtoMessage()               {}
func (*SystemState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SystemState) GetState() State {
	if m != nil {
		return m.State
	}
	return State_DISABLED
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*MultiValue)(nil), "MultiValue")
	proto.RegisterType((*SystemState)(nil), "SystemState")
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterEnum("State", State_name, State_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RequestService service

type RequestServiceClient interface {
	Process(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error)
	SystemStateChange(ctx context.Context, in *Empty, opts ...grpc.CallOption) (RequestService_SystemStateChangeClient, error)
}

type requestServiceClient struct {
	cc *grpc.ClientConn
}

func NewRequestServiceClient(cc *grpc.ClientConn) RequestServiceClient {
	return &requestServiceClient{cc}
}

func (c *requestServiceClient) Process(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/RequestService/Process", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestServiceClient) SystemStateChange(ctx context.Context, in *Empty, opts ...grpc.CallOption) (RequestService_SystemStateChangeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RequestService_serviceDesc.Streams[0], c.cc, "/RequestService/SystemStateChange", opts...)
	if err != nil {
		return nil, err
	}
	x := &requestServiceSystemStateChangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RequestService_SystemStateChangeClient interface {
	Recv() (*SystemState, error)
	grpc.ClientStream
}

type requestServiceSystemStateChangeClient struct {
	grpc.ClientStream
}

func (x *requestServiceSystemStateChangeClient) Recv() (*SystemState, error) {
	m := new(SystemState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RequestService service

type RequestServiceServer interface {
	Process(context.Context, *Request) (*Empty, error)
	SystemStateChange(*Empty, RequestService_SystemStateChangeServer) error
}

func RegisterRequestServiceServer(s *grpc.Server, srv RequestServiceServer) {
	s.RegisterService(&_RequestService_serviceDesc, srv)
}

func _RequestService_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestServiceServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RequestService/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestServiceServer).Process(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestService_SystemStateChange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RequestServiceServer).SystemStateChange(m, &requestServiceSystemStateChangeServer{stream})
}

type RequestService_SystemStateChangeServer interface {
	Send(*SystemState) error
	grpc.ServerStream
}

type requestServiceSystemStateChangeServer struct {
	grpc.ServerStream
}

func (x *requestServiceSystemStateChangeServer) Send(m *SystemState) error {
	return x.ServerStream.SendMsg(m)
}

var _RequestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RequestService",
	HandlerType: (*RequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _RequestService_Process_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SystemStateChange",
			Handler:       _RequestService_SystemStateChange_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/proto/qaclana.proto",
}

// Client API for SystemStateService service

type SystemStateServiceClient interface {
	Receive(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SystemStateService_ReceiveClient, error)
}

type systemStateServiceClient struct {
	cc *grpc.ClientConn
}

func NewSystemStateServiceClient(cc *grpc.ClientConn) SystemStateServiceClient {
	return &systemStateServiceClient{cc}
}

func (c *systemStateServiceClient) Receive(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SystemStateService_ReceiveClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SystemStateService_serviceDesc.Streams[0], c.cc, "/SystemStateService/Receive", opts...)
	if err != nil {
		return nil, err
	}
	x := &systemStateServiceReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SystemStateService_ReceiveClient interface {
	Recv() (*SystemState, error)
	grpc.ClientStream
}

type systemStateServiceReceiveClient struct {
	grpc.ClientStream
}

func (x *systemStateServiceReceiveClient) Recv() (*SystemState, error) {
	m := new(SystemState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SystemStateService service

type SystemStateServiceServer interface {
	Receive(*Empty, SystemStateService_ReceiveServer) error
}

func RegisterSystemStateServiceServer(s *grpc.Server, srv SystemStateServiceServer) {
	s.RegisterService(&_SystemStateService_serviceDesc, srv)
}

func _SystemStateService_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SystemStateServiceServer).Receive(m, &systemStateServiceReceiveServer{stream})
}

type SystemStateService_ReceiveServer interface {
	Send(*SystemState) error
	grpc.ServerStream
}

type systemStateServiceReceiveServer struct {
	grpc.ServerStream
}

func (x *systemStateServiceReceiveServer) Send(m *SystemState) error {
	return x.ServerStream.SendMsg(m)
}

var _SystemStateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SystemStateService",
	HandlerType: (*SystemStateServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receive",
			Handler:       _SystemStateService_Receive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/proto/qaclana.proto",
}

func init() { proto.RegisterFile("pkg/proto/qaclana.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x6d, 0x9c, 0x34, 0x93, 0xd4, 0x98, 0xa5, 0x88, 0x28, 0x6a, 0xa5, 0xb0, 0x70, 0x88,
	0x40, 0x6c, 0x90, 0xe1, 0x00, 0xdc, 0x9a, 0xc6, 0x84, 0x48, 0x34, 0x75, 0x6d, 0xa9, 0x47, 0xa4,
	0xc5, 0x1e, 0xb5, 0x56, 0x93, 0x38, 0x59, 0x6f, 0x22, 0xf9, 0x93, 0xf9, 0x0b, 0xe4, 0xb5, 0x53,
	0xec, 0xaa, 0xea, 0xa1, 0xb7, 0x9d, 0xf7, 0xe6, 0xbd, 0x99, 0xd1, 0xcc, 0xc2, 0xeb, 0xf5, 0xed,
	0xf5, 0x68, 0x2d, 0x12, 0x99, 0x8c, 0x36, 0x3c, 0x5c, 0xf0, 0x15, 0x67, 0x2a, 0xa2, 0x7f, 0x0d,
	0x68, 0xf9, 0xb8, 0xd9, 0x62, 0x2a, 0xc9, 0x3b, 0x38, 0x0c, 0x17, 0x31, 0xae, 0xe4, 0x69, 0x14,
	0x09, 0x4c, 0xd3, 0x9e, 0x36, 0xd0, 0x86, 0x6d, 0xbf, 0x0e, 0x12, 0x0b, 0xf4, 0x38, 0xea, 0xe9,
	0x8a, 0xd2, 0xe3, 0x88, 0x8c, 0xa0, 0x75, 0x83, 0x3c, 0x42, 0x91, 0xf6, 0x8c, 0x81, 0x31, 0xec,
	0x38, 0xaf, 0x58, 0x69, 0xc8, 0x7e, 0x16, 0xb8, 0xbb, 0x92, 0x22, 0xf3, 0xf7, 0x59, 0x64, 0x0a,
	0xcf, 0x37, 0x5b, 0x14, 0x99, 0xc7, 0x05, 0x5f, 0xa2, 0xcc, 0x85, 0xcf, 0x94, 0xf0, 0xe4, 0x4e,
	0x78, 0x59, 0xe7, 0x0b, 0x83, 0xfb, 0x2a, 0x32, 0x01, 0x6b, 0x9d, 0xa4, 0xb2, 0xe2, 0x63, 0x2a,
	0x9f, 0xe3, 0x3b, 0x1f, 0xaf, 0x46, 0x17, 0x36, 0xf7, 0x34, 0xfd, 0x29, 0x74, 0xab, 0x7d, 0x12,
	0x1b, 0x8c, 0x5b, 0xcc, 0xca, 0xd9, 0xf3, 0x27, 0x79, 0x03, 0xe6, 0x8e, 0x2f, 0xb6, 0xa8, 0x86,
	0xee, 0x38, 0x1d, 0x76, 0xbe, 0x5d, 0xc8, 0xf8, 0x2a, 0x87, 0xfc, 0x82, 0xf9, 0xae, 0x7f, 0xd5,
	0xfa, 0x17, 0x70, 0xf4, 0x50, 0xdf, 0x4f, 0x37, 0x9c, 0xc3, 0xcb, 0x07, 0x06, 0x78, 0xb2, 0x1f,
	0xa5, 0x00, 0xff, 0x09, 0x72, 0xb4, 0x17, 0x69, 0x03, 0x63, 0xd8, 0x2e, 0xf3, 0xe8, 0x07, 0xe8,
	0x04, 0x59, 0x2a, 0x71, 0x19, 0x48, 0x2e, 0x91, 0x1c, 0x83, 0x99, 0xe6, 0x0f, 0x55, 0xcd, 0x72,
	0x9a, 0x4c, 0xc1, 0x7e, 0x01, 0xd2, 0x16, 0x98, 0xee, 0x72, 0x2d, 0xb3, 0xf7, 0x5f, 0xc0, 0x2c,
	0xf2, 0xbb, 0x70, 0x30, 0x99, 0x05, 0xa7, 0xe3, 0x5f, 0xee, 0xc4, 0x6e, 0x10, 0x0b, 0xc0, 0x73,
	0xfd, 0xf3, 0x59, 0x10, 0xcc, 0xae, 0x5c, 0x5b, 0x23, 0x87, 0xd0, 0x76, 0xe7, 0x3f, 0x2e, 0xfc,
	0xb3, 0xd9, 0x7c, 0x6a, 0xeb, 0xce, 0x6f, 0xb0, 0xca, 0x45, 0x05, 0x28, 0x76, 0x71, 0x88, 0xe4,
	0x04, 0x5a, 0x9e, 0x48, 0xc2, 0xfc, 0xcc, 0x0e, 0xf6, 0x4b, 0xec, 0x37, 0x99, 0x2a, 0x42, 0x1b,
	0xe4, 0x23, 0xbc, 0xa8, 0x34, 0x77, 0x76, 0xc3, 0x57, 0xd7, 0x48, 0x4a, 0xba, 0xdf, 0x65, 0x15,
	0x8e, 0x36, 0x3e, 0x69, 0xce, 0x37, 0x20, 0x15, 0x68, 0x5f, 0xe3, 0x6d, 0x7e, 0xf0, 0x21, 0xc6,
	0xbb, 0x47, 0xa4, 0x63, 0x0a, 0x76, 0x9c, 0xb0, 0xda, 0x57, 0x19, 0x77, 0x2f, 0x8b, 0xd0, 0xcb,
	0x23, 0x4f, 0xfb, 0xd3, 0x54, 0xf0, 0xe7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x42, 0xb2, 0x6a,
	0x23, 0x5c, 0x03, 0x00, 0x00,
}
