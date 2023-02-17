// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/DownloadBot/v1/greeter.proto

package v1

import (
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

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_DownloadBot_v1_greeter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_DownloadBot_v1_greeter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_api_DownloadBot_v1_greeter_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_DownloadBot_v1_greeter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_DownloadBot_v1_greeter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_api_DownloadBot_v1_greeter_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SuddenMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SuddenMessage) Reset() {
	*x = SuddenMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_DownloadBot_v1_greeter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuddenMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuddenMessage) ProtoMessage() {}

func (x *SuddenMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_DownloadBot_v1_greeter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuddenMessage.ProtoReflect.Descriptor instead.
func (*SuddenMessage) Descriptor() ([]byte, []int) {
	return file_api_DownloadBot_v1_greeter_proto_rawDescGZIP(), []int{2}
}

func (x *SuddenMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TMStopMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid string `protobuf:"bytes,1,opt,name=gid,proto3" json:"gid,omitempty"`
}

func (x *TMStopMsg) Reset() {
	*x = TMStopMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_DownloadBot_v1_greeter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TMStopMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TMStopMsg) ProtoMessage() {}

func (x *TMStopMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_DownloadBot_v1_greeter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TMStopMsg.ProtoReflect.Descriptor instead.
func (*TMStopMsg) Descriptor() ([]byte, []int) {
	return file_api_DownloadBot_v1_greeter_proto_rawDescGZIP(), []int{3}
}

func (x *TMStopMsg) GetGid() string {
	if x != nil {
		return x.Gid
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_DownloadBot_v1_greeter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_api_DownloadBot_v1_greeter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_api_DownloadBot_v1_greeter_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName string `protobuf:"bytes,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_DownloadBot_v1_greeter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_api_DownloadBot_v1_greeter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_api_DownloadBot_v1_greeter_proto_rawDescGZIP(), []int{5}
}

func (x *Ping) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

var File_api_DownloadBot_v1_greeter_proto protoreflect.FileDescriptor

var file_api_DownloadBot_v1_greeter_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x0c, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x53, 0x75, 0x64, 0x64, 0x65,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x1d, 0x0a, 0x09, 0x54, 0x4d, 0x53, 0x74, 0x6f, 0x70, 0x4d, 0x73, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x69,
	0x64, 0x22, 0x36, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x04, 0x70, 0x69, 0x6e,
	0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x32, 0xde, 0x01, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x38, 0x0a,
	0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x75, 0x64, 0x64, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x64, 0x64, 0x65, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x54, 0x4d, 0x53, 0x74, 0x6f,
	0x70, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x4d, 0x53, 0x74,
	0x6f, 0x70, 0x4d, 0x73, 0x67, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x0d, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x69, 0x6e, 0x67, 0x1a,
	0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_DownloadBot_v1_greeter_proto_rawDescOnce sync.Once
	file_api_DownloadBot_v1_greeter_proto_rawDescData = file_api_DownloadBot_v1_greeter_proto_rawDesc
)

func file_api_DownloadBot_v1_greeter_proto_rawDescGZIP() []byte {
	file_api_DownloadBot_v1_greeter_proto_rawDescOnce.Do(func() {
		file_api_DownloadBot_v1_greeter_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_DownloadBot_v1_greeter_proto_rawDescData)
	})
	return file_api_DownloadBot_v1_greeter_proto_rawDescData
}

var file_api_DownloadBot_v1_greeter_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_DownloadBot_v1_greeter_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),  // 0: greeter.HelloRequest
	(*HelloReply)(nil),    // 1: greeter.HelloReply
	(*SuddenMessage)(nil), // 2: greeter.SuddenMessage
	(*TMStopMsg)(nil),     // 3: greeter.TMStopMsg
	(*Status)(nil),        // 4: greeter.Status
	(*Ping)(nil),          // 5: greeter.ping
}
var file_api_DownloadBot_v1_greeter_proto_depIdxs = []int32{
	0, // 0: greeter.Greeter.SayHello:input_type -> greeter.HelloRequest
	2, // 1: greeter.Greeter.SendSuddenMessage:input_type -> greeter.SuddenMessage
	3, // 2: greeter.Greeter.TMStop:input_type -> greeter.TMStopMsg
	5, // 3: greeter.Greeter.Ping:input_type -> greeter.ping
	1, // 4: greeter.Greeter.SayHello:output_type -> greeter.HelloReply
	4, // 5: greeter.Greeter.SendSuddenMessage:output_type -> greeter.Status
	4, // 6: greeter.Greeter.TMStop:output_type -> greeter.Status
	4, // 7: greeter.Greeter.Ping:output_type -> greeter.Status
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_DownloadBot_v1_greeter_proto_init() }
func file_api_DownloadBot_v1_greeter_proto_init() {
	if File_api_DownloadBot_v1_greeter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_DownloadBot_v1_greeter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_api_DownloadBot_v1_greeter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_api_DownloadBot_v1_greeter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuddenMessage); i {
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
		file_api_DownloadBot_v1_greeter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TMStopMsg); i {
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
		file_api_DownloadBot_v1_greeter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_api_DownloadBot_v1_greeter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
			RawDescriptor: file_api_DownloadBot_v1_greeter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_DownloadBot_v1_greeter_proto_goTypes,
		DependencyIndexes: file_api_DownloadBot_v1_greeter_proto_depIdxs,
		MessageInfos:      file_api_DownloadBot_v1_greeter_proto_msgTypes,
	}.Build()
	File_api_DownloadBot_v1_greeter_proto = out.File
	file_api_DownloadBot_v1_greeter_proto_rawDesc = nil
	file_api_DownloadBot_v1_greeter_proto_goTypes = nil
	file_api_DownloadBot_v1_greeter_proto_depIdxs = nil
}
