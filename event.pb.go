// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: event.proto

package schema

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Status  Status   `protobuf:"varint,3,opt,name=status,proto3,enum=schema.Status" json:"status,omitempty"`
	Payload []*Field `protobuf:"bytes,10,rep,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Event) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNACTIVATED
}

func (x *Event) GetPayload() []*Field {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x86, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x27, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0xb4, 0x01, 0x0a, 0x0c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x41, 0x50, 0x49, 0x12, 0x4a, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x0d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x30, 0x2f, 0x69, 0x6e,
	0x73, 0x70, 0x65, 0x63, 0x74, 0x30, 0x01, 0x12, 0x58, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0f,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x0d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x31,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x30, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x7b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x7d, 0x32, 0xfc, 0x02, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x41, 0x50, 0x49, 0x12, 0x4a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x0d, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x30, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x5b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x0d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x1a,
	0x28, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x30, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f,
	0x7b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x08,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x0d, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x32, 0x2a, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x30, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x2f, 0x7b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d, 0x12, 0x62, 0x0a, 0x0a,
	0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x0d, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2e, 0x32, 0x2c, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x30, 0x2f, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x7b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x7d,
	0x42, 0x3f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x68, 0x69, 0x73, 0x75, 0x69, 0x74, 0x65,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x68, 0x69, 0x73, 0x75, 0x69, 0x74, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_event_proto_goTypes = []interface{}{
	(*Event)(nil),   // 0: schema.Event
	(Status)(0),     // 1: schema.Status
	(*Field)(nil),   // 2: schema.Field
	(*Options)(nil), // 3: schema.Options
}
var file_event_proto_depIdxs = []int32{
	1, // 0: schema.Event.status:type_name -> schema.Status
	2, // 1: schema.Event.payload:type_name -> schema.Field
	3, // 2: schema.EventReadAPI.List:input_type -> schema.Options
	3, // 3: schema.EventReadAPI.Get:input_type -> schema.Options
	0, // 4: schema.EventWriteAPI.Create:input_type -> schema.Event
	0, // 5: schema.EventWriteAPI.Update:input_type -> schema.Event
	3, // 6: schema.EventWriteAPI.Activate:input_type -> schema.Options
	3, // 7: schema.EventWriteAPI.Deactivate:input_type -> schema.Options
	0, // 8: schema.EventReadAPI.List:output_type -> schema.Event
	0, // 9: schema.EventReadAPI.Get:output_type -> schema.Event
	0, // 10: schema.EventWriteAPI.Create:output_type -> schema.Event
	0, // 11: schema.EventWriteAPI.Update:output_type -> schema.Event
	0, // 12: schema.EventWriteAPI.Activate:output_type -> schema.Event
	0, // 13: schema.EventWriteAPI.Deactivate:output_type -> schema.Event
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventReadAPIClient is the client API for EventReadAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventReadAPIClient interface {
	List(ctx context.Context, in *Options, opts ...grpc.CallOption) (EventReadAPI_ListClient, error)
	Get(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Event, error)
}

type eventReadAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewEventReadAPIClient(cc grpc.ClientConnInterface) EventReadAPIClient {
	return &eventReadAPIClient{cc}
}

func (c *eventReadAPIClient) List(ctx context.Context, in *Options, opts ...grpc.CallOption) (EventReadAPI_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventReadAPI_serviceDesc.Streams[0], "/schema.EventReadAPI/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventReadAPIListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventReadAPI_ListClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventReadAPIListClient struct {
	grpc.ClientStream
}

func (x *eventReadAPIListClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventReadAPIClient) Get(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/schema.EventReadAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventReadAPIServer is the server API for EventReadAPI service.
type EventReadAPIServer interface {
	List(*Options, EventReadAPI_ListServer) error
	Get(context.Context, *Options) (*Event, error)
}

// UnimplementedEventReadAPIServer can be embedded to have forward compatible implementations.
type UnimplementedEventReadAPIServer struct {
}

func (*UnimplementedEventReadAPIServer) List(*Options, EventReadAPI_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedEventReadAPIServer) Get(context.Context, *Options) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterEventReadAPIServer(s *grpc.Server, srv EventReadAPIServer) {
	s.RegisterService(&_EventReadAPI_serviceDesc, srv)
}

func _EventReadAPI_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Options)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventReadAPIServer).List(m, &eventReadAPIListServer{stream})
}

type EventReadAPI_ListServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventReadAPIListServer struct {
	grpc.ServerStream
}

func (x *eventReadAPIListServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _EventReadAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Options)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventReadAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.EventReadAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventReadAPIServer).Get(ctx, req.(*Options))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventReadAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schema.EventReadAPI",
	HandlerType: (*EventReadAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _EventReadAPI_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _EventReadAPI_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "event.proto",
}

// EventWriteAPIClient is the client API for EventWriteAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventWriteAPIClient interface {
	Create(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	Update(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	Activate(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Event, error)
	Deactivate(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Event, error)
}

type eventWriteAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewEventWriteAPIClient(cc grpc.ClientConnInterface) EventWriteAPIClient {
	return &eventWriteAPIClient{cc}
}

func (c *eventWriteAPIClient) Create(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/schema.EventWriteAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventWriteAPIClient) Update(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/schema.EventWriteAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventWriteAPIClient) Activate(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/schema.EventWriteAPI/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventWriteAPIClient) Deactivate(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/schema.EventWriteAPI/Deactivate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventWriteAPIServer is the server API for EventWriteAPI service.
type EventWriteAPIServer interface {
	Create(context.Context, *Event) (*Event, error)
	Update(context.Context, *Event) (*Event, error)
	Activate(context.Context, *Options) (*Event, error)
	Deactivate(context.Context, *Options) (*Event, error)
}

// UnimplementedEventWriteAPIServer can be embedded to have forward compatible implementations.
type UnimplementedEventWriteAPIServer struct {
}

func (*UnimplementedEventWriteAPIServer) Create(context.Context, *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEventWriteAPIServer) Update(context.Context, *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedEventWriteAPIServer) Activate(context.Context, *Options) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (*UnimplementedEventWriteAPIServer) Deactivate(context.Context, *Options) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deactivate not implemented")
}

func RegisterEventWriteAPIServer(s *grpc.Server, srv EventWriteAPIServer) {
	s.RegisterService(&_EventWriteAPI_serviceDesc, srv)
}

func _EventWriteAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventWriteAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.EventWriteAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventWriteAPIServer).Create(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventWriteAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventWriteAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.EventWriteAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventWriteAPIServer).Update(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventWriteAPI_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Options)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventWriteAPIServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.EventWriteAPI/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventWriteAPIServer).Activate(ctx, req.(*Options))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventWriteAPI_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Options)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventWriteAPIServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schema.EventWriteAPI/Deactivate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventWriteAPIServer).Deactivate(ctx, req.(*Options))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventWriteAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "schema.EventWriteAPI",
	HandlerType: (*EventWriteAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EventWriteAPI_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EventWriteAPI_Update_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _EventWriteAPI_Activate_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _EventWriteAPI_Deactivate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event.proto",
}
