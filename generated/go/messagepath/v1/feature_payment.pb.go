// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: messagepath/v1/feature_payment.proto

package messagepath

import (
	v11 "github.com/sokkit-io/xiphias-model-common/generated/go/common/v1"
	_ "github.com/sokkit-io/xiphias-model-common/generated/go/kikoptions"
	v1 "github.com/sokkit-io/xiphias-model-common/generated/go/kin/payment/v1"
	v12 "github.com/sokkit-io/xiphias-model-common/generated/go/offer/v1"
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

type TransactionDetailsAttachment_Target int32

const (
	TransactionDetailsAttachment_UNKNOWN   TransactionDetailsAttachment_Target = 0
	TransactionDetailsAttachment_SENDER    TransactionDetailsAttachment_Target = 1
	TransactionDetailsAttachment_RECIPIENT TransactionDetailsAttachment_Target = 2
)

// Enum value maps for TransactionDetailsAttachment_Target.
var (
	TransactionDetailsAttachment_Target_name = map[int32]string{
		0: "UNKNOWN",
		1: "SENDER",
		2: "RECIPIENT",
	}
	TransactionDetailsAttachment_Target_value = map[string]int32{
		"UNKNOWN":   0,
		"SENDER":    1,
		"RECIPIENT": 2,
	}
)

func (x TransactionDetailsAttachment_Target) Enum() *TransactionDetailsAttachment_Target {
	p := new(TransactionDetailsAttachment_Target)
	*p = x
	return p
}

func (x TransactionDetailsAttachment_Target) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionDetailsAttachment_Target) Descriptor() protoreflect.EnumDescriptor {
	return file_messagepath_v1_feature_payment_proto_enumTypes[0].Descriptor()
}

func (TransactionDetailsAttachment_Target) Type() protoreflect.EnumType {
	return &file_messagepath_v1_feature_payment_proto_enumTypes[0]
}

func (x TransactionDetailsAttachment_Target) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionDetailsAttachment_Target.Descriptor instead.
func (TransactionDetailsAttachment_Target) EnumDescriptor() ([]byte, []int) {
	return file_messagepath_v1_feature_payment_proto_rawDescGZIP(), []int{0, 0}
}

// TransactionDetailsAttachment is meant to be attached to status messages related to feature payments
// It contains the intended target of the status message as well as the tip amount
type TransactionDetailsAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the intended target of the attachment (This helps mobile client carry out custom actions for sender
	// and recipient)
	Target        TransactionDetailsAttachment_Target `protobuf:"varint,1,opt,name=target,proto3,enum=common.messagepath.v1.TransactionDetailsAttachment_Target" json:"target,omitempty"`
	Amount        *v1.KinAmount                       `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	SenderJid     *v11.XiBareUserJidOrAliasJid        `protobuf:"bytes,3,opt,name=sender_jid,json=senderJid,proto3" json:"sender_jid,omitempty"`
	RecipientJid  *v11.XiBareUserJidOrAliasJid        `protobuf:"bytes,4,opt,name=recipient_jid,json=recipientJid,proto3" json:"recipient_jid,omitempty"`
	UserOfferData *v12.UserOfferData                  `protobuf:"bytes,5,opt,name=user_offer_data,json=userOfferData,proto3" json:"user_offer_data,omitempty"`
}

func (x *TransactionDetailsAttachment) Reset() {
	*x = TransactionDetailsAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagepath_v1_feature_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionDetailsAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionDetailsAttachment) ProtoMessage() {}

func (x *TransactionDetailsAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_messagepath_v1_feature_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionDetailsAttachment.ProtoReflect.Descriptor instead.
func (*TransactionDetailsAttachment) Descriptor() ([]byte, []int) {
	return file_messagepath_v1_feature_payment_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionDetailsAttachment) GetTarget() TransactionDetailsAttachment_Target {
	if x != nil {
		return x.Target
	}
	return TransactionDetailsAttachment_UNKNOWN
}

func (x *TransactionDetailsAttachment) GetAmount() *v1.KinAmount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *TransactionDetailsAttachment) GetSenderJid() *v11.XiBareUserJidOrAliasJid {
	if x != nil {
		return x.SenderJid
	}
	return nil
}

func (x *TransactionDetailsAttachment) GetRecipientJid() *v11.XiBareUserJidOrAliasJid {
	if x != nil {
		return x.RecipientJid
	}
	return nil
}

func (x *TransactionDetailsAttachment) GetUserOfferData() *v12.UserOfferData {
	if x != nil {
		return x.UserOfferData
	}
	return nil
}

var File_messagepath_v1_feature_payment_proto protoreflect.FileDescriptor

var file_messagepath_v1_feature_payment_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x6b, 0x69, 0x6e, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6b,
	0x69, 0x6b, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x03, 0x0a, 0x1c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x52, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6b, 0x69, 0x6e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4b, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6a,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x69, 0x42, 0x61, 0x72, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4a,
	0x69, 0x64, 0x4f, 0x72, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x4a, 0x69, 0x64, 0x52, 0x09, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x4a, 0x69, 0x64, 0x12, 0x47, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6a, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x69, 0x42, 0x61, 0x72,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x69, 0x64, 0x4f, 0x72, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x4a,
	0x69, 0x64, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4a, 0x69, 0x64,
	0x12, 0x46, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4f, 0x66, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x52,
	0x45, 0x43, 0x49, 0x50, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x42, 0x74, 0x0a, 0x19, 0x63, 0x6f,
	0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74,
	0x68, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x78, 0x69,
	0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2d, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0xa2, 0x02, 0x03, 0x4b, 0x50, 0x42,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messagepath_v1_feature_payment_proto_rawDescOnce sync.Once
	file_messagepath_v1_feature_payment_proto_rawDescData = file_messagepath_v1_feature_payment_proto_rawDesc
)

func file_messagepath_v1_feature_payment_proto_rawDescGZIP() []byte {
	file_messagepath_v1_feature_payment_proto_rawDescOnce.Do(func() {
		file_messagepath_v1_feature_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_messagepath_v1_feature_payment_proto_rawDescData)
	})
	return file_messagepath_v1_feature_payment_proto_rawDescData
}

var file_messagepath_v1_feature_payment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messagepath_v1_feature_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_messagepath_v1_feature_payment_proto_goTypes = []interface{}{
	(TransactionDetailsAttachment_Target)(0), // 0: common.messagepath.v1.TransactionDetailsAttachment.Target
	(*TransactionDetailsAttachment)(nil),     // 1: common.messagepath.v1.TransactionDetailsAttachment
	(*v1.KinAmount)(nil),                     // 2: common.kin.payment.v1.KinAmount
	(*v11.XiBareUserJidOrAliasJid)(nil),      // 3: common.v1.XiBareUserJidOrAliasJid
	(*v12.UserOfferData)(nil),                // 4: common.offer.v1.UserOfferData
}
var file_messagepath_v1_feature_payment_proto_depIdxs = []int32{
	0, // 0: common.messagepath.v1.TransactionDetailsAttachment.target:type_name -> common.messagepath.v1.TransactionDetailsAttachment.Target
	2, // 1: common.messagepath.v1.TransactionDetailsAttachment.amount:type_name -> common.kin.payment.v1.KinAmount
	3, // 2: common.messagepath.v1.TransactionDetailsAttachment.sender_jid:type_name -> common.v1.XiBareUserJidOrAliasJid
	3, // 3: common.messagepath.v1.TransactionDetailsAttachment.recipient_jid:type_name -> common.v1.XiBareUserJidOrAliasJid
	4, // 4: common.messagepath.v1.TransactionDetailsAttachment.user_offer_data:type_name -> common.offer.v1.UserOfferData
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_messagepath_v1_feature_payment_proto_init() }
func file_messagepath_v1_feature_payment_proto_init() {
	if File_messagepath_v1_feature_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messagepath_v1_feature_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionDetailsAttachment); i {
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
			RawDescriptor: file_messagepath_v1_feature_payment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messagepath_v1_feature_payment_proto_goTypes,
		DependencyIndexes: file_messagepath_v1_feature_payment_proto_depIdxs,
		EnumInfos:         file_messagepath_v1_feature_payment_proto_enumTypes,
		MessageInfos:      file_messagepath_v1_feature_payment_proto_msgTypes,
	}.Build()
	File_messagepath_v1_feature_payment_proto = out.File
	file_messagepath_v1_feature_payment_proto_rawDesc = nil
	file_messagepath_v1_feature_payment_proto_goTypes = nil
	file_messagepath_v1_feature_payment_proto_depIdxs = nil
}
