// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: entity/v1/subscription_common.proto

package entity

import (
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

type RosterSyncToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Clients MUST NOT interpret this message.
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *RosterSyncToken) Reset() {
	*x = RosterSyncToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_v1_subscription_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RosterSyncToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RosterSyncToken) ProtoMessage() {}

func (x *RosterSyncToken) ProtoReflect() protoreflect.Message {
	mi := &file_entity_v1_subscription_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RosterSyncToken.ProtoReflect.Descriptor instead.
func (*RosterSyncToken) Descriptor() ([]byte, []int) {
	return file_entity_v1_subscription_common_proto_rawDescGZIP(), []int{0}
}

func (x *RosterSyncToken) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Pseudo-entity that indicates the given alias JID is blocked by the requesting user
type AliasBlockEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AliasJid *v1.XiAliasJid `protobuf:"bytes,1,opt,name=alias_jid,json=aliasJid,proto3" json:"alias_jid,omitempty"`
}

func (x *AliasBlockEntry) Reset() {
	*x = AliasBlockEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_v1_subscription_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AliasBlockEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AliasBlockEntry) ProtoMessage() {}

func (x *AliasBlockEntry) ProtoReflect() protoreflect.Message {
	mi := &file_entity_v1_subscription_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AliasBlockEntry.ProtoReflect.Descriptor instead.
func (*AliasBlockEntry) Descriptor() ([]byte, []int) {
	return file_entity_v1_subscription_common_proto_rawDescGZIP(), []int{1}
}

func (x *AliasBlockEntry) GetAliasJid() *v1.XiAliasJid {
	if x != nil {
		return x.AliasJid
	}
	return nil
}

type RosterEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OneOf user or group
	//
	// Types that are assignable to RosterEntryKind:
	//	*RosterEntry_UserData
	//	*RosterEntry_GroupData
	//	*RosterEntry_AliasBlockData
	RosterEntryKind isRosterEntry_RosterEntryKind `protobuf_oneof:"roster_entry_kind"`
	// Only applies to user rosters
	// Has this person being blocked by roster owner?
	IsBlocked bool `protobuf:"varint,3,opt,name=is_blocked,json=isBlocked,proto3" json:"is_blocked,omitempty"`
}

func (x *RosterEntry) Reset() {
	*x = RosterEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_v1_subscription_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RosterEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RosterEntry) ProtoMessage() {}

func (x *RosterEntry) ProtoReflect() protoreflect.Message {
	mi := &file_entity_v1_subscription_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RosterEntry.ProtoReflect.Descriptor instead.
func (*RosterEntry) Descriptor() ([]byte, []int) {
	return file_entity_v1_subscription_common_proto_rawDescGZIP(), []int{2}
}

func (m *RosterEntry) GetRosterEntryKind() isRosterEntry_RosterEntryKind {
	if m != nil {
		return m.RosterEntryKind
	}
	return nil
}

func (x *RosterEntry) GetUserData() *EntityUserRosterEntry {
	if x, ok := x.GetRosterEntryKind().(*RosterEntry_UserData); ok {
		return x.UserData
	}
	return nil
}

func (x *RosterEntry) GetGroupData() *EntityGroupRosterEntry {
	if x, ok := x.GetRosterEntryKind().(*RosterEntry_GroupData); ok {
		return x.GroupData
	}
	return nil
}

func (x *RosterEntry) GetAliasBlockData() *AliasBlockEntry {
	if x, ok := x.GetRosterEntryKind().(*RosterEntry_AliasBlockData); ok {
		return x.AliasBlockData
	}
	return nil
}

func (x *RosterEntry) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

type isRosterEntry_RosterEntryKind interface {
	isRosterEntry_RosterEntryKind()
}

type RosterEntry_UserData struct {
	UserData *EntityUserRosterEntry `protobuf:"bytes,1,opt,name=user_data,json=userData,proto3,oneof"`
}

type RosterEntry_GroupData struct {
	GroupData *EntityGroupRosterEntry `protobuf:"bytes,2,opt,name=group_data,json=groupData,proto3,oneof"`
}

type RosterEntry_AliasBlockData struct {
	AliasBlockData *AliasBlockEntry `protobuf:"bytes,4,opt,name=alias_block_data,json=aliasBlockData,proto3,oneof"`
}

func (*RosterEntry_UserData) isRosterEntry_RosterEntryKind() {}

func (*RosterEntry_GroupData) isRosterEntry_RosterEntryKind() {}

func (*RosterEntry_AliasBlockData) isRosterEntry_RosterEntryKind() {}

var File_entity_v1_subscription_common_proto protoreflect.FileDescriptor

var file_entity_v1_subscription_common_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0f, 0x52, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0b, 0xca, 0x9d,
	0x25, 0x07, 0x08, 0x01, 0x28, 0x01, 0x30, 0x80, 0x28, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x4d, 0x0a, 0x0f, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x09, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x5f, 0x6a,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x69, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x4a, 0x69, 0x64, 0x42,
	0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x01, 0x52, 0x08, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x4a, 0x69,
	0x64, 0x22, 0xa3, 0x02, 0x0a, 0x0b, 0x52, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x46, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x0a, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x4d, 0x0a, 0x10, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x48, 0x00, 0x52, 0x0e, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x72, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x69, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x6b,
	0x69, 0x6b, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5a,
	0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b,
	0x69, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x78, 0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x3b, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0xa0, 0x01, 0x01, 0xa2, 0x02, 0x04, 0x53, 0x55,
	0x42, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_v1_subscription_common_proto_rawDescOnce sync.Once
	file_entity_v1_subscription_common_proto_rawDescData = file_entity_v1_subscription_common_proto_rawDesc
)

func file_entity_v1_subscription_common_proto_rawDescGZIP() []byte {
	file_entity_v1_subscription_common_proto_rawDescOnce.Do(func() {
		file_entity_v1_subscription_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_v1_subscription_common_proto_rawDescData)
	})
	return file_entity_v1_subscription_common_proto_rawDescData
}

var file_entity_v1_subscription_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_entity_v1_subscription_common_proto_goTypes = []interface{}{
	(*RosterSyncToken)(nil),        // 0: common.entity.v1.RosterSyncToken
	(*AliasBlockEntry)(nil),        // 1: common.entity.v1.AliasBlockEntry
	(*RosterEntry)(nil),            // 2: common.entity.v1.RosterEntry
	(*v1.XiAliasJid)(nil),          // 3: common.v1.XiAliasJid
	(*EntityUserRosterEntry)(nil),  // 4: common.entity.v1.EntityUserRosterEntry
	(*EntityGroupRosterEntry)(nil), // 5: common.entity.v1.EntityGroupRosterEntry
}
var file_entity_v1_subscription_common_proto_depIdxs = []int32{
	3, // 0: common.entity.v1.AliasBlockEntry.alias_jid:type_name -> common.v1.XiAliasJid
	4, // 1: common.entity.v1.RosterEntry.user_data:type_name -> common.entity.v1.EntityUserRosterEntry
	5, // 2: common.entity.v1.RosterEntry.group_data:type_name -> common.entity.v1.EntityGroupRosterEntry
	1, // 3: common.entity.v1.RosterEntry.alias_block_data:type_name -> common.entity.v1.AliasBlockEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_entity_v1_subscription_common_proto_init() }
func file_entity_v1_subscription_common_proto_init() {
	if File_entity_v1_subscription_common_proto != nil {
		return
	}
	file_entity_v1_entity_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entity_v1_subscription_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RosterSyncToken); i {
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
		file_entity_v1_subscription_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AliasBlockEntry); i {
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
		file_entity_v1_subscription_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RosterEntry); i {
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
	file_entity_v1_subscription_common_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*RosterEntry_UserData)(nil),
		(*RosterEntry_GroupData)(nil),
		(*RosterEntry_AliasBlockData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entity_v1_subscription_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_v1_subscription_common_proto_goTypes,
		DependencyIndexes: file_entity_v1_subscription_common_proto_depIdxs,
		MessageInfos:      file_entity_v1_subscription_common_proto_msgTypes,
	}.Build()
	File_entity_v1_subscription_common_proto = out.File
	file_entity_v1_subscription_common_proto_rawDesc = nil
	file_entity_v1_subscription_common_proto_goTypes = nil
	file_entity_v1_subscription_common_proto_depIdxs = nil
}
