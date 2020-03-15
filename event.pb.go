// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package schema

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

type Event struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Status               Status   `protobuf:"varint,3,opt,name=status,proto3,enum=schema.Status" json:"status,omitempty"`
	Payload              []*Field `protobuf:"bytes,4,rep,name=payload,proto3" json:"payload,omitempty"`
	Options              []*Field `protobuf:"bytes,5,rep,name=options,proto3" json:"options,omitempty"`
	Files                []*File  `protobuf:"bytes,6,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Event) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UNACTIVATED
}

func (m *Event) GetPayload() []*Field {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Event) GetOptions() []*Field {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Event) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "schema.Event")
}

func init() {
	proto.RegisterFile("event.proto", fileDescriptor_2d17a9d3f0ddf27e)
}

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd2, 0xdf, 0x4a, 0xbc, 0x40,
	0x14, 0xc0, 0x71, 0x66, 0x77, 0x75, 0xf7, 0x77, 0x76, 0x7f, 0x1b, 0x9c, 0x6e, 0x86, 0xbd, 0x92,
	0x8d, 0x4a, 0xba, 0x90, 0xb0, 0x27, 0xd8, 0xed, 0x1f, 0x41, 0x90, 0x18, 0x3d, 0xc0, 0xa4, 0x27,
	0x76, 0x40, 0x1d, 0x71, 0x66, 0x85, 0x2e, 0x7b, 0xc6, 0x5e, 0x28, 0x1c, 0xb5, 0x08, 0x16, 0xec,
	0x4e, 0xcf, 0xf7, 0xc3, 0xc8, 0xc1, 0x81, 0x39, 0xd5, 0x54, 0x98, 0xa0, 0xac, 0x94, 0x51, 0xe8,
	0xea, 0x64, 0x47, 0xb9, 0x58, 0x2d, 0x12, 0x95, 0xe7, 0xaa, 0x68, 0xa7, 0xeb, 0x4f, 0x06, 0xce,
	0x6d, 0xa3, 0x10, 0x61, 0x52, 0x88, 0x9c, 0x38, 0xf3, 0x98, 0xff, 0x2f, 0xb6, 0xcf, 0xc8, 0x61,
	0x5a, 0x53, 0xa5, 0xa5, 0x2a, 0xf8, 0xc8, 0x8e, 0xfb, 0x57, 0x3c, 0x03, 0x57, 0x1b, 0x61, 0xf6,
	0x9a, 0x8f, 0x3d, 0xe6, 0x2f, 0xc3, 0x65, 0xd0, 0x1e, 0x1f, 0x3c, 0xdb, 0x69, 0xdc, 0x55, 0x3c,
	0x87, 0x69, 0x29, 0xde, 0x33, 0x25, 0x52, 0x3e, 0xf1, 0xc6, 0xfe, 0x3c, 0xfc, 0xdf, 0xc3, 0x3b,
	0x49, 0x59, 0x1a, 0xf7, 0xb5, 0x81, 0xaa, 0x34, 0x52, 0x15, 0x9a, 0x3b, 0x07, 0x61, 0x57, 0x71,
	0x0d, 0xce, 0x9b, 0xcc, 0x48, 0x73, 0xd7, 0xb2, 0xc5, 0x0f, 0xcb, 0x28, 0x6e, 0x53, 0xf8, 0x31,
	0x82, 0x99, 0xdd, 0x6a, 0x13, 0x3d, 0xa0, 0x0f, 0x93, 0x47, 0xa9, 0x0d, 0x1e, 0xf5, 0xf2, 0xa9,
	0x3d, 0x69, 0xf5, 0xfd, 0x05, 0x4b, 0x2f, 0x19, 0x9e, 0xc2, 0xf8, 0x9e, 0x06, 0x61, 0xb3, 0xfb,
	0x75, 0x45, 0xc2, 0x10, 0xfe, 0x0e, 0x07, 0xdc, 0x4b, 0x99, 0x0e, 0x3b, 0x1f, 0x66, 0x9b, 0xc4,
	0xc8, 0x7a, 0x58, 0x5e, 0x00, 0xdc, 0x90, 0xf8, 0x93, 0xdd, 0x9e, 0xc0, 0x71, 0xa2, 0xf2, 0xa0,
	0xdc, 0x49, 0xbd, 0x97, 0x86, 0xba, 0xb8, 0x05, 0x5b, 0xa3, 0xe6, 0xe7, 0x47, 0xec, 0xd5, 0xb5,
	0xb7, 0xe0, 0xea, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xce, 0x27, 0xf5, 0x0f, 0x2a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventAPIClient is the client API for EventAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventAPIClient interface {
	List(ctx context.Context, in *Options, opts ...grpc.CallOption) (EventAPI_ListClient, error)
	Get(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Event, error)
	Create(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	Update(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	Activate(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	Deactivate(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
}

type eventAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewEventAPIClient(cc grpc.ClientConnInterface) EventAPIClient {
	return &eventAPIClient{cc}
}

func (c *eventAPIClient) List(ctx context.Context, in *Options, opts ...grpc.CallOption) (EventAPI_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventAPI_serviceDesc.Streams[0], "/schema.EventAPI/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventAPIListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventAPI_ListClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventAPIListClient struct {
	grpc.ClientStream
}

func (x *eventAPIListClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventAPIClient) Get(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/schema.EventAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventAPIClient) Create(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/schema.EventAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventAPIClient) Update(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/schema.EventAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventAPIClient) Activate(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/schema.EventAPI/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventAPIClient) Deactivate(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/schema.EventAPI/Deactivate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventAPIServer is the server API for EventAPI service.
type EventAPIServer interface {
	List(*Options, EventAPI_ListServer) error
	Get(context.Context, *Options) (*Event, error)
	Create(context.Context, *Event) (*Event, error)
	Update(context.Context, *Event) (*Event, error)
	Activate(context.Context, *Event) (*Event, error)
	Deactivate(context.Context, *Event) (*Event, error)
}

// UnimplementedEventAPIServer can be embedded to have forward compatible implementations.
type UnimplementedEventAPIServer struct {
}

func (*UnimplementedEventAPIServer) List(req *Options, srv EventAPI_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedEventAPIServer) Get(ctx context.Context, req *Options) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedEventAPIServer) Create(ctx context.Context, req *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEventAPIServer) Update(ctx context.Context, req *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedEventAPIServer) Activate(ctx context.Context, req *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (*UnimplementedEventAPIServer) Deactivate(ctx context.Context, req *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deactivate not implemented")
}

func RegisterEventAPIServer(s *grpc.Server, srv EventAPIServer) {
	s.RegisterService(&_EventAPI_serviceDesc, srv)
}

func _EventAPI_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Options)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventAPIServer).List(m, &eventAPIListServer{stream})
}

type EventAPI_ListServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventAPIListServer struct {
	grpc.ServerStream
}

func (x *eventAPIListServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _EventAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Options)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.EventAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventAPIServer).Get(ctx, req.(*Options))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.EventAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventAPIServer).Create(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.EventAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventAPIServer).Update(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventAPI_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventAPIServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.EventAPI/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventAPIServer).Activate(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventAPI_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventAPIServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.EventAPI/Deactivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventAPIServer).Deactivate(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schema.EventAPI",
	HandlerType: (*EventAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _EventAPI_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _EventAPI_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EventAPI_Update_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _EventAPI_Activate_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _EventAPI_Deactivate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _EventAPI_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "event.proto",
}