// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: conversation.proto

package pb

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

type GroupCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
}

func (x *GroupCreateReq) Reset() {
	*x = GroupCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCreateReq) ProtoMessage() {}

func (x *GroupCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCreateReq.ProtoReflect.Descriptor instead.
func (*GroupCreateReq) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{0}
}

func (x *GroupCreateReq) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type GroupCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
}

func (x *GroupCreateResp) Reset() {
	*x = GroupCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conversation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCreateResp) ProtoMessage() {}

func (x *GroupCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_conversation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCreateResp.ProtoReflect.Descriptor instead.
func (*GroupCreateResp) Descriptor() ([]byte, []int) {
	return file_conversation_proto_rawDescGZIP(), []int{1}
}

func (x *GroupCreateResp) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_conversation_proto protoreflect.FileDescriptor

var file_conversation_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x32, 0x4f, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_conversation_proto_rawDescOnce sync.Once
	file_conversation_proto_rawDescData = file_conversation_proto_rawDesc
)

func file_conversation_proto_rawDescGZIP() []byte {
	file_conversation_proto_rawDescOnce.Do(func() {
		file_conversation_proto_rawDescData = protoimpl.X.CompressGZIP(file_conversation_proto_rawDescData)
	})
	return file_conversation_proto_rawDescData
}

var file_conversation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_conversation_proto_goTypes = []interface{}{
	(*GroupCreateReq)(nil),  // 0: pb.GroupCreateReq
	(*GroupCreateResp)(nil), // 1: pb.GroupCreateResp
	(*RequestHeader)(nil),   // 2: pb.RequestHeader
	(*ResponseHeader)(nil),  // 3: pb.ResponseHeader
}
var file_conversation_proto_depIdxs = []int32{
	2, // 0: pb.GroupCreateReq.header:type_name -> pb.RequestHeader
	3, // 1: pb.GroupCreateResp.header:type_name -> pb.ResponseHeader
	0, // 2: pb.conversationService.GroupCreate:input_type -> pb.GroupCreateReq
	1, // 3: pb.conversationService.GroupCreate:output_type -> pb.GroupCreateResp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_conversation_proto_init() }
func file_conversation_proto_init() {
	if File_conversation_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_conversation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCreateReq); i {
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
		file_conversation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCreateResp); i {
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
			RawDescriptor: file_conversation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_conversation_proto_goTypes,
		DependencyIndexes: file_conversation_proto_depIdxs,
		MessageInfos:      file_conversation_proto_msgTypes,
	}.Build()
	File_conversation_proto = out.File
	file_conversation_proto_rawDesc = nil
	file_conversation_proto_goTypes = nil
	file_conversation_proto_depIdxs = nil
}