// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: common_rpc.proto

package common

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	_ "github.com/sokkit-io/xiphias-model-common/generated/go/kikoptions"
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

type VoidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VoidRequest) Reset() {
	*x = VoidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoidRequest) ProtoMessage() {}

func (x *VoidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoidRequest.ProtoReflect.Descriptor instead.
func (*VoidRequest) Descriptor() ([]byte, []int) {
	return file_common_rpc_proto_rawDescGZIP(), []int{0}
}

type VoidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VoidResponse) Reset() {
	*x = VoidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoidResponse) ProtoMessage() {}

func (x *VoidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoidResponse.ProtoReflect.Descriptor instead.
func (*VoidResponse) Descriptor() ([]byte, []int) {
	return file_common_rpc_proto_rawDescGZIP(), []int{1}
}

// XiRequestId identifies a "logical request", which may consist of many
// xiphias requests. In other words, if an XiRequestId is provided for a
// call, than any xiphias request spawning from that call should re-use
// the provided XiRequestId.
//
// The string representation of an XiRequestId is the canonical UUID
// representation. This representation should only be used where using
// an XiRequestId directly is not possible.
//
// It is primarily used for cases such as metrics and request tracing.
type XiRequestId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *XiUuid `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *XiRequestId) Reset() {
	*x = XiRequestId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XiRequestId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XiRequestId) ProtoMessage() {}

func (x *XiRequestId) ProtoReflect() protoreflect.Message {
	mi := &file_common_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XiRequestId.ProtoReflect.Descriptor instead.
func (*XiRequestId) Descriptor() ([]byte, []int) {
	return file_common_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *XiRequestId) GetId() *XiUuid {
	if x != nil {
		return x.Id
	}
	return nil
}

// A string token that is used to route calls to services in a
// deterministic (sticky) way.
//
// The contents of the token string is left up to the client, but it should
// be something consistent such as JID or device ID. The users of this
// token should not care about what the client has used as the token string.
type XiRoutingToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *XiRoutingToken) Reset() {
	*x = XiRoutingToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XiRoutingToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XiRoutingToken) ProtoMessage() {}

func (x *XiRoutingToken) ProtoReflect() protoreflect.Message {
	mi := &file_common_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XiRoutingToken.ProtoReflect.Descriptor instead.
func (*XiRoutingToken) Descriptor() ([]byte, []int) {
	return file_common_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *XiRoutingToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SelfDescribingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains a single file descriptor that describes the message
	// that is serialized in message_data.
	FieldDescriptorSet *descriptor.FileDescriptorSet `protobuf:"bytes,1,opt,name=field_descriptor_set,json=fieldDescriptorSet,proto3" json:"field_descriptor_set,omitempty"`
	// The message name that is encoded in message data
	// For example, if `message MyCustomMessage` is encoded,
	// then the message name should be MyCustomMessage.
	MessageName string `protobuf:"bytes,2,opt,name=message_name,json=messageName,proto3" json:"message_name,omitempty"`
	// The serialized message data.
	MessageData []byte `protobuf:"bytes,3,opt,name=message_data,json=messageData,proto3" json:"message_data,omitempty"`
}

func (x *SelfDescribingMessage) Reset() {
	*x = SelfDescribingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelfDescribingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelfDescribingMessage) ProtoMessage() {}

func (x *SelfDescribingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_common_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelfDescribingMessage.ProtoReflect.Descriptor instead.
func (*SelfDescribingMessage) Descriptor() ([]byte, []int) {
	return file_common_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *SelfDescribingMessage) GetFieldDescriptorSet() *descriptor.FileDescriptorSet {
	if x != nil {
		return x.FieldDescriptorSet
	}
	return nil
}

func (x *SelfDescribingMessage) GetMessageName() string {
	if x != nil {
		return x.MessageName
	}
	return ""
}

func (x *SelfDescribingMessage) GetMessageData() []byte {
	if x != nil {
		return x.MessageData
	}
	return nil
}

var File_common_rpc_proto protoreflect.FileDescriptor

var file_common_rpc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6b, 0x69, 0x6b,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d,
	0x0a, 0x0b, 0x56, 0x6f, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a,
	0x0c, 0x56, 0x6f, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x0a,
	0x0b, 0x58, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x58, 0x69, 0x55, 0x75, 0x69, 0x64, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x0e, 0x58, 0x69, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xca, 0x9d, 0x25, 0x05, 0x28, 0x01, 0x30, 0x80, 0x01,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb3, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x6c, 0x66,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x54, 0x0a, 0x14, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x53, 0x65, 0x74, 0x52, 0x12, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6c, 0xaa,
	0xa3, 0x2a, 0x02, 0x10, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x78,
	0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x0e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d,
	0x69, 0x6f, 0x2f, 0x78, 0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x67, 0x6f, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_rpc_proto_rawDescOnce sync.Once
	file_common_rpc_proto_rawDescData = file_common_rpc_proto_rawDesc
)

func file_common_rpc_proto_rawDescGZIP() []byte {
	file_common_rpc_proto_rawDescOnce.Do(func() {
		file_common_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_rpc_proto_rawDescData)
	})
	return file_common_rpc_proto_rawDescData
}

var file_common_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_rpc_proto_goTypes = []interface{}{
	(*VoidRequest)(nil),                  // 0: common.VoidRequest
	(*VoidResponse)(nil),                 // 1: common.VoidResponse
	(*XiRequestId)(nil),                  // 2: common.XiRequestId
	(*XiRoutingToken)(nil),               // 3: common.XiRoutingToken
	(*SelfDescribingMessage)(nil),        // 4: common.SelfDescribingMessage
	(*XiUuid)(nil),                       // 5: common.XiUuid
	(*descriptor.FileDescriptorSet)(nil), // 6: google.protobuf.FileDescriptorSet
}
var file_common_rpc_proto_depIdxs = []int32{
	5, // 0: common.XiRequestId.id:type_name -> common.XiUuid
	6, // 1: common.SelfDescribingMessage.field_descriptor_set:type_name -> google.protobuf.FileDescriptorSet
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_rpc_proto_init() }
func file_common_rpc_proto_init() {
	if File_common_rpc_proto != nil {
		return
	}
	file_common_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoidRequest); i {
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
		file_common_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoidResponse); i {
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
		file_common_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XiRequestId); i {
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
		file_common_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XiRoutingToken); i {
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
		file_common_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelfDescribingMessage); i {
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
			RawDescriptor: file_common_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_rpc_proto_goTypes,
		DependencyIndexes: file_common_rpc_proto_depIdxs,
		MessageInfos:      file_common_rpc_proto_msgTypes,
	}.Build()
	File_common_rpc_proto = out.File
	file_common_rpc_proto_rawDesc = nil
	file_common_rpc_proto_goTypes = nil
	file_common_rpc_proto_depIdxs = nil
}
