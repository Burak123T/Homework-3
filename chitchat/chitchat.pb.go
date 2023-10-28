// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: chitchat/chitchat.proto

package chitchat

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

type ClientMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Text    string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Lamport int32  `protobuf:"varint,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *ClientMessage) Reset() {
	*x = ClientMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chitchat_chitchat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMessage) ProtoMessage() {}

func (x *ClientMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chitchat_chitchat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMessage.ProtoReflect.Descriptor instead.
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return file_chitchat_chitchat_proto_rawDescGZIP(), []int{0}
}

func (x *ClientMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ClientMessage) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type ServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Text    string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Lamport int32  `protobuf:"varint,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *ServerMessage) Reset() {
	*x = ServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chitchat_chitchat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMessage) ProtoMessage() {}

func (x *ServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chitchat_chitchat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMessage.ProtoReflect.Descriptor instead.
func (*ServerMessage) Descriptor() ([]byte, []int) {
	return file_chitchat_chitchat_proto_rawDescGZIP(), []int{1}
}

func (x *ServerMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServerMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ServerMessage) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type SentChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confirmation bool `protobuf:"varint,1,opt,name=confirmation,proto3" json:"confirmation,omitempty"`
}

func (x *SentChatResponse) Reset() {
	*x = SentChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chitchat_chitchat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SentChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SentChatResponse) ProtoMessage() {}

func (x *SentChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chitchat_chitchat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SentChatResponse.ProtoReflect.Descriptor instead.
func (*SentChatResponse) Descriptor() ([]byte, []int) {
	return file_chitchat_chitchat_proto_rawDescGZIP(), []int{2}
}

func (x *SentChatResponse) GetConfirmation() bool {
	if x != nil {
		return x.Confirmation
	}
	return false
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lamport int32  `protobuf:"varint,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chitchat_chitchat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_chitchat_chitchat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_chitchat_chitchat_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

var File_chitchat_chitchat_proto protoreflect.FileDescriptor

var file_chitchat_chitchat_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x68, 0x69, 0x74, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x69, 0x74, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x68, 0x69, 0x74, 0x63,
	0x68, 0x61, 0x74, 0x22, 0x51, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c,
	0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x51, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x36, 0x0a, 0x10, 0x53, 0x65, 0x6e,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x44, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x32, 0xf7, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12,
	0x0e, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x17, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x30, 0x01, 0x12, 0x33, 0x0a, 0x05, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x12, 0x0e, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53,
	0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x42, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17,
	0x2e, 0x63, 0x68, 0x69, 0x74, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x17, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x63, 0x68, 0x69, 0x74, 0x63, 0x68, 0x61, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chitchat_chitchat_proto_rawDescOnce sync.Once
	file_chitchat_chitchat_proto_rawDescData = file_chitchat_chitchat_proto_rawDesc
)

func file_chitchat_chitchat_proto_rawDescGZIP() []byte {
	file_chitchat_chitchat_proto_rawDescOnce.Do(func() {
		file_chitchat_chitchat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chitchat_chitchat_proto_rawDescData)
	})
	return file_chitchat_chitchat_proto_rawDescData
}

var file_chitchat_chitchat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chitchat_chitchat_proto_goTypes = []interface{}{
	(*ClientMessage)(nil),    // 0: chitchat.ClientMessage
	(*ServerMessage)(nil),    // 1: chitchat.ServerMessage
	(*SentChatResponse)(nil), // 2: chitchat.SentChatResponse
	(*User)(nil),             // 3: chitchat.User
}
var file_chitchat_chitchat_proto_depIdxs = []int32{
	3, // 0: chitchat.ChatService.Join:input_type -> chitchat.User
	3, // 1: chitchat.ChatService.Leave:input_type -> chitchat.User
	0, // 2: chitchat.ChatService.SendMessage:input_type -> chitchat.ClientMessage
	3, // 3: chitchat.ChatService.BroadcastListener:input_type -> chitchat.User
	1, // 4: chitchat.ChatService.Join:output_type -> chitchat.ServerMessage
	2, // 5: chitchat.ChatService.Leave:output_type -> chitchat.SentChatResponse
	2, // 6: chitchat.ChatService.SendMessage:output_type -> chitchat.SentChatResponse
	0, // 7: chitchat.ChatService.BroadcastListener:output_type -> chitchat.ClientMessage
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chitchat_chitchat_proto_init() }
func file_chitchat_chitchat_proto_init() {
	if File_chitchat_chitchat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chitchat_chitchat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMessage); i {
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
		file_chitchat_chitchat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMessage); i {
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
		file_chitchat_chitchat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SentChatResponse); i {
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
		file_chitchat_chitchat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_chitchat_chitchat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chitchat_chitchat_proto_goTypes,
		DependencyIndexes: file_chitchat_chitchat_proto_depIdxs,
		MessageInfos:      file_chitchat_chitchat_proto_msgTypes,
	}.Build()
	File_chitchat_chitchat_proto = out.File
	file_chitchat_chitchat_proto_rawDesc = nil
	file_chitchat_chitchat_proto_goTypes = nil
	file_chitchat_chitchat_proto_depIdxs = nil
}
