// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: matching/v1/matching_common.proto

package matching

import (
	_go "github.com/sokkit-io/xiphias-model-common/generated/go"
	v1 "github.com/sokkit-io/xiphias-model-common/generated/go/common/v1"
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

type ChatSessionState int32

const (
	// Participants can chat, end chat, make friends
	ChatSessionState_CHAT_ACTIVE ChatSessionState = 0
	// Chatting and voting is closed, session is considered marked for cleanup at this state
	// If users end the chat early, chat sessions are transitioned to EXPIRY with no voting period
	ChatSessionState_EXPIRED ChatSessionState = 2
	// Session is considered ended, but users have opted to upgrade the conversation to a real jid conversation
	// No more messages should be allowed through to the chat participants, though the end session status message
	// May need to be allowed through an upgraded session, and messages to an upgraded chat session may need forwarding
	// to the real participants to resolve some race conditions surrounding upgrades.
	ChatSessionState_UPGRADED ChatSessionState = 3
)

// Enum value maps for ChatSessionState.
var (
	ChatSessionState_name = map[int32]string{
		0: "CHAT_ACTIVE",
		2: "EXPIRED",
		3: "UPGRADED",
	}
	ChatSessionState_value = map[string]int32{
		"CHAT_ACTIVE": 0,
		"EXPIRED":     2,
		"UPGRADED":    3,
	}
)

func (x ChatSessionState) Enum() *ChatSessionState {
	p := new(ChatSessionState)
	*p = x
	return p
}

func (x ChatSessionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatSessionState) Descriptor() protoreflect.EnumDescriptor {
	return file_matching_v1_matching_common_proto_enumTypes[0].Descriptor()
}

func (ChatSessionState) Type() protoreflect.EnumType {
	return &file_matching_v1_matching_common_proto_enumTypes[0]
}

func (x ChatSessionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatSessionState.Descriptor instead.
func (ChatSessionState) EnumDescriptor() ([]byte, []int) {
	return file_matching_v1_matching_common_proto_rawDescGZIP(), []int{0}
}

type AnonMatchingSessionKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A session can be identified by the session id, or either of the ajids in the conversation
	// session_id is preferred wherever possible because it is less costly than retrieving by alias
	//
	// Types that are assignable to ChatSessionId:
	//	*AnonMatchingSessionKey_SessionId
	//	*AnonMatchingSessionKey_ChatPartnerAlias
	ChatSessionId isAnonMatchingSessionKey_ChatSessionId `protobuf_oneof:"chat_session_id"`
}

func (x *AnonMatchingSessionKey) Reset() {
	*x = AnonMatchingSessionKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matching_v1_matching_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnonMatchingSessionKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnonMatchingSessionKey) ProtoMessage() {}

func (x *AnonMatchingSessionKey) ProtoReflect() protoreflect.Message {
	mi := &file_matching_v1_matching_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnonMatchingSessionKey.ProtoReflect.Descriptor instead.
func (*AnonMatchingSessionKey) Descriptor() ([]byte, []int) {
	return file_matching_v1_matching_common_proto_rawDescGZIP(), []int{0}
}

func (m *AnonMatchingSessionKey) GetChatSessionId() isAnonMatchingSessionKey_ChatSessionId {
	if m != nil {
		return m.ChatSessionId
	}
	return nil
}

func (x *AnonMatchingSessionKey) GetSessionId() *_go.XiUuid {
	if x, ok := x.GetChatSessionId().(*AnonMatchingSessionKey_SessionId); ok {
		return x.SessionId
	}
	return nil
}

func (x *AnonMatchingSessionKey) GetChatPartnerAlias() *v1.XiAliasJid {
	if x, ok := x.GetChatSessionId().(*AnonMatchingSessionKey_ChatPartnerAlias); ok {
		return x.ChatPartnerAlias
	}
	return nil
}

type isAnonMatchingSessionKey_ChatSessionId interface {
	isAnonMatchingSessionKey_ChatSessionId()
}

type AnonMatchingSessionKey_SessionId struct {
	SessionId *_go.XiUuid `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3,oneof"`
}

type AnonMatchingSessionKey_ChatPartnerAlias struct {
	ChatPartnerAlias *v1.XiAliasJid `protobuf:"bytes,2,opt,name=chat_partner_alias,json=chatPartnerAlias,proto3,oneof"`
}

func (*AnonMatchingSessionKey_SessionId) isAnonMatchingSessionKey_ChatSessionId() {}

func (*AnonMatchingSessionKey_ChatPartnerAlias) isAnonMatchingSessionKey_ChatSessionId() {}

var File_matching_v1_matching_common_proto protoreflect.FileDescriptor

var file_matching_v1_matching_common_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01,
	0x0a, 0x16, 0x41, 0x6e, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x58, 0x69, 0x55, 0x75, 0x69, 0x64, 0x48, 0x00, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x12, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x58, 0x69, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x4a, 0x69, 0x64, 0x48, 0x00, 0x52, 0x10,
	0x63, 0x68, 0x61, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x41, 0x6c, 0x69, 0x61, 0x73,
	0x42, 0x11, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x2a, 0x3e, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x41, 0x54, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49,
	0x52, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45,
	0x44, 0x10, 0x03, 0x42, 0x6e, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5a, 0x4b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74,
	0x2d, 0x69, 0x6f, 0x2f, 0x78, 0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x3b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0xa0, 0x01, 0x01, 0xa2, 0x02, 0x03,
	0x4d, 0x41, 0x54, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_matching_v1_matching_common_proto_rawDescOnce sync.Once
	file_matching_v1_matching_common_proto_rawDescData = file_matching_v1_matching_common_proto_rawDesc
)

func file_matching_v1_matching_common_proto_rawDescGZIP() []byte {
	file_matching_v1_matching_common_proto_rawDescOnce.Do(func() {
		file_matching_v1_matching_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_matching_v1_matching_common_proto_rawDescData)
	})
	return file_matching_v1_matching_common_proto_rawDescData
}

var file_matching_v1_matching_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_matching_v1_matching_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_matching_v1_matching_common_proto_goTypes = []interface{}{
	(ChatSessionState)(0),          // 0: common.matching.v1.ChatSessionState
	(*AnonMatchingSessionKey)(nil), // 1: common.matching.v1.AnonMatchingSessionKey
	(*_go.XiUuid)(nil),             // 2: common.XiUuid
	(*v1.XiAliasJid)(nil),          // 3: common.v1.XiAliasJid
}
var file_matching_v1_matching_common_proto_depIdxs = []int32{
	2, // 0: common.matching.v1.AnonMatchingSessionKey.session_id:type_name -> common.XiUuid
	3, // 1: common.matching.v1.AnonMatchingSessionKey.chat_partner_alias:type_name -> common.v1.XiAliasJid
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_matching_v1_matching_common_proto_init() }
func file_matching_v1_matching_common_proto_init() {
	if File_matching_v1_matching_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_matching_v1_matching_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnonMatchingSessionKey); i {
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
	file_matching_v1_matching_common_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AnonMatchingSessionKey_SessionId)(nil),
		(*AnonMatchingSessionKey_ChatPartnerAlias)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_matching_v1_matching_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_matching_v1_matching_common_proto_goTypes,
		DependencyIndexes: file_matching_v1_matching_common_proto_depIdxs,
		EnumInfos:         file_matching_v1_matching_common_proto_enumTypes,
		MessageInfos:      file_matching_v1_matching_common_proto_msgTypes,
	}.Build()
	File_matching_v1_matching_common_proto = out.File
	file_matching_v1_matching_common_proto_rawDesc = nil
	file_matching_v1_matching_common_proto_goTypes = nil
	file_matching_v1_matching_common_proto_depIdxs = nil
}
