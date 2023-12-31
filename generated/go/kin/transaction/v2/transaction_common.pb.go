// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: kin/transaction/v2/transaction_common.proto

package transaction

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/sokkit-io/xiphias-model-common/generated/go/kikoptions"
	v2 "github.com/sokkit-io/xiphias-model-common/generated/go/kin/account/v2"
	v21 "github.com/sokkit-io/xiphias-model-common/generated/go/kin/payment/v2"
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

type TransactionPaymentInfo_TransactionStatus int32

const (
	// Indicates that the transaction is submitted to the block-chain
	TransactionPaymentInfo_IN_PROGRESS TransactionPaymentInfo_TransactionStatus = 0
	// Indicates that the transaction has executed successfully on the block-chain
	TransactionPaymentInfo_SUCCESS TransactionPaymentInfo_TransactionStatus = 1
	// Indicates that the transaction has failed
	TransactionPaymentInfo_FAIL TransactionPaymentInfo_TransactionStatus = 2
)

// Enum value maps for TransactionPaymentInfo_TransactionStatus.
var (
	TransactionPaymentInfo_TransactionStatus_name = map[int32]string{
		0: "IN_PROGRESS",
		1: "SUCCESS",
		2: "FAIL",
	}
	TransactionPaymentInfo_TransactionStatus_value = map[string]int32{
		"IN_PROGRESS": 0,
		"SUCCESS":     1,
		"FAIL":        2,
	}
)

func (x TransactionPaymentInfo_TransactionStatus) Enum() *TransactionPaymentInfo_TransactionStatus {
	p := new(TransactionPaymentInfo_TransactionStatus)
	*p = x
	return p
}

func (x TransactionPaymentInfo_TransactionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionPaymentInfo_TransactionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_kin_transaction_v2_transaction_common_proto_enumTypes[0].Descriptor()
}

func (TransactionPaymentInfo_TransactionStatus) Type() protoreflect.EnumType {
	return &file_kin_transaction_v2_transaction_common_proto_enumTypes[0]
}

func (x TransactionPaymentInfo_TransactionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionPaymentInfo_TransactionStatus.Descriptor instead.
func (TransactionPaymentInfo_TransactionStatus) EnumDescriptor() ([]byte, []int) {
	return file_kin_transaction_v2_transaction_common_proto_rawDescGZIP(), []int{0, 0}
}

type TransactionPaymentInfo_TransactionFailureReason int32

const (
	TransactionPaymentInfo_UNKNOWN                 TransactionPaymentInfo_TransactionFailureReason = 0
	TransactionPaymentInfo_BLOCKCHAIN_TIMEOUT      TransactionPaymentInfo_TransactionFailureReason = 1
	TransactionPaymentInfo_INSUFFICIENT_FUNDS      TransactionPaymentInfo_TransactionFailureReason = 2
	TransactionPaymentInfo_INVALID_SEQUENCE_NUMBER TransactionPaymentInfo_TransactionFailureReason = 3
)

// Enum value maps for TransactionPaymentInfo_TransactionFailureReason.
var (
	TransactionPaymentInfo_TransactionFailureReason_name = map[int32]string{
		0: "UNKNOWN",
		1: "BLOCKCHAIN_TIMEOUT",
		2: "INSUFFICIENT_FUNDS",
		3: "INVALID_SEQUENCE_NUMBER",
	}
	TransactionPaymentInfo_TransactionFailureReason_value = map[string]int32{
		"UNKNOWN":                 0,
		"BLOCKCHAIN_TIMEOUT":      1,
		"INSUFFICIENT_FUNDS":      2,
		"INVALID_SEQUENCE_NUMBER": 3,
	}
)

func (x TransactionPaymentInfo_TransactionFailureReason) Enum() *TransactionPaymentInfo_TransactionFailureReason {
	p := new(TransactionPaymentInfo_TransactionFailureReason)
	*p = x
	return p
}

func (x TransactionPaymentInfo_TransactionFailureReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionPaymentInfo_TransactionFailureReason) Descriptor() protoreflect.EnumDescriptor {
	return file_kin_transaction_v2_transaction_common_proto_enumTypes[1].Descriptor()
}

func (TransactionPaymentInfo_TransactionFailureReason) Type() protoreflect.EnumType {
	return &file_kin_transaction_v2_transaction_common_proto_enumTypes[1]
}

func (x TransactionPaymentInfo_TransactionFailureReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionPaymentInfo_TransactionFailureReason.Descriptor instead.
func (TransactionPaymentInfo_TransactionFailureReason) EnumDescriptor() ([]byte, []int) {
	return file_kin_transaction_v2_transaction_common_proto_rawDescGZIP(), []int{0, 1}
}

type TransactionPaymentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Transaction id uniquely identifies a transaction.  This won't necessarily be the transaction hash
	TransactionId *TransactionId `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// Kin account id (wallet) of the sender
	SenderKinAccountId *v2.KinAccountId `protobuf:"bytes,2,opt,name=sender_kin_account_id,json=senderKinAccountId,proto3" json:"sender_kin_account_id,omitempty"`
	// Kin account id (wallet) of the recipient
	RecipientKinAccountId *v2.KinAccountId `protobuf:"bytes,3,opt,name=recipient_kin_account_id,json=recipientKinAccountId,proto3" json:"recipient_kin_account_id,omitempty"`
	// Transaction total amount
	Amount *v21.KinAmount `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	// Transaction type (earn | spend)
	TransactionType v21.TransactionType `protobuf:"varint,5,opt,name=transaction_type,json=transactionType,proto3,enum=common.kin.payment.v2.TransactionType" json:"transaction_type,omitempty"`
	// The UTC time according to the transaction service, when the transaction was submitted to the blockchain
	SubmissionTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=submission_time,json=submissionTime,proto3" json:"submission_time,omitempty"`
	// The UTC time according to the blockchain when the transaction was confirmed to the blockchain
	ConfirmationTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=confirmation_time,json=confirmationTime,proto3" json:"confirmation_time,omitempty"`
	// Status of the transaction
	Status TransactionPaymentInfo_TransactionStatus        `protobuf:"varint,8,opt,name=status,proto3,enum=common.kin.transaction.v2.TransactionPaymentInfo_TransactionStatus" json:"status,omitempty"`
	Reason TransactionPaymentInfo_TransactionFailureReason `protobuf:"varint,9,opt,name=reason,proto3,enum=common.kin.transaction.v2.TransactionPaymentInfo_TransactionFailureReason" json:"reason,omitempty"`
}

func (x *TransactionPaymentInfo) Reset() {
	*x = TransactionPaymentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kin_transaction_v2_transaction_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionPaymentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionPaymentInfo) ProtoMessage() {}

func (x *TransactionPaymentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_kin_transaction_v2_transaction_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionPaymentInfo.ProtoReflect.Descriptor instead.
func (*TransactionPaymentInfo) Descriptor() ([]byte, []int) {
	return file_kin_transaction_v2_transaction_common_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionPaymentInfo) GetTransactionId() *TransactionId {
	if x != nil {
		return x.TransactionId
	}
	return nil
}

func (x *TransactionPaymentInfo) GetSenderKinAccountId() *v2.KinAccountId {
	if x != nil {
		return x.SenderKinAccountId
	}
	return nil
}

func (x *TransactionPaymentInfo) GetRecipientKinAccountId() *v2.KinAccountId {
	if x != nil {
		return x.RecipientKinAccountId
	}
	return nil
}

func (x *TransactionPaymentInfo) GetAmount() *v21.KinAmount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *TransactionPaymentInfo) GetTransactionType() v21.TransactionType {
	if x != nil {
		return x.TransactionType
	}
	return v21.TransactionType(0)
}

func (x *TransactionPaymentInfo) GetSubmissionTime() *timestamp.Timestamp {
	if x != nil {
		return x.SubmissionTime
	}
	return nil
}

func (x *TransactionPaymentInfo) GetConfirmationTime() *timestamp.Timestamp {
	if x != nil {
		return x.ConfirmationTime
	}
	return nil
}

func (x *TransactionPaymentInfo) GetStatus() TransactionPaymentInfo_TransactionStatus {
	if x != nil {
		return x.Status
	}
	return TransactionPaymentInfo_IN_PROGRESS
}

func (x *TransactionPaymentInfo) GetReason() TransactionPaymentInfo_TransactionFailureReason {
	if x != nil {
		return x.Reason
	}
	return TransactionPaymentInfo_UNKNOWN
}

// TransactionDisplayInfo contains any information that kik has added on top of what is in the
// blockchain, used for display purposes on clients.
// This could be absent if this is a transaction that Kik didn't manage.
type TransactionDisplayInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Title of transaction that will be displayed in client transaction history (For ex. Tipping)
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Description of transaction that will be displayed in client transaction history (For ex. Tipped user X in group Y)
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *TransactionDisplayInfo) Reset() {
	*x = TransactionDisplayInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kin_transaction_v2_transaction_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionDisplayInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionDisplayInfo) ProtoMessage() {}

func (x *TransactionDisplayInfo) ProtoReflect() protoreflect.Message {
	mi := &file_kin_transaction_v2_transaction_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionDisplayInfo.ProtoReflect.Descriptor instead.
func (*TransactionDisplayInfo) Descriptor() ([]byte, []int) {
	return file_kin_transaction_v2_transaction_common_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionDisplayInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TransactionDisplayInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Stellar transaction envelope contains raw signed transaction data
// For all offers involving user sending Kin, the client is required to created a transaction using recipient account
// public address and offer amount data. The transaction is then signed by the user's private wallet key (which is
// stored on the users device, and is never transmitted to Kik's server backend)
type TransactionEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawValue []byte `protobuf:"bytes,1,opt,name=raw_value,json=rawValue,proto3" json:"raw_value,omitempty"`
}

func (x *TransactionEnvelope) Reset() {
	*x = TransactionEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kin_transaction_v2_transaction_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionEnvelope) ProtoMessage() {}

func (x *TransactionEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_kin_transaction_v2_transaction_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionEnvelope.ProtoReflect.Descriptor instead.
func (*TransactionEnvelope) Descriptor() ([]byte, []int) {
	return file_kin_transaction_v2_transaction_common_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionEnvelope) GetRawValue() []byte {
	if x != nil {
		return x.RawValue
	}
	return nil
}

// Transaction id is a generic id used to uniquely identify a transaction history record
// Transaction id can be the transaction hash (for spend transaction) or a combination of different block chain
// attributes to uniquely identify a transaction (including transaction hash, operation id...etc.)
type TransactionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Blockchain transaction hash for the transaction the record corresponds to
	TransactionHash string `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	// The operation from the transaction associated with the transaction_hash that the record corresponds to
	OperationIndex uint32 `protobuf:"varint,2,opt,name=operation_index,json=operationIndex,proto3" json:"operation_index,omitempty"`
}

func (x *TransactionId) Reset() {
	*x = TransactionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kin_transaction_v2_transaction_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionId) ProtoMessage() {}

func (x *TransactionId) ProtoReflect() protoreflect.Message {
	mi := &file_kin_transaction_v2_transaction_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionId.ProtoReflect.Descriptor instead.
func (*TransactionId) Descriptor() ([]byte, []int) {
	return file_kin_transaction_v2_transaction_common_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionId) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *TransactionId) GetOperationIndex() uint32 {
	if x != nil {
		return x.OperationIndex
	}
	return 0
}

var File_kin_transaction_v2_transaction_common_proto protoreflect.FileDescriptor

var file_kin_transaction_v2_transaction_common_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6b, 0x69, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6b, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6b, 0x69, 0x6e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6b, 0x69, 0x6e, 0x2f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6,
	0x07, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x57, 0x0a, 0x0e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6b, 0x69, 0x6e, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x06, 0xca, 0x9d, 0x25,
	0x02, 0x08, 0x01, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x5e, 0x0a, 0x15, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6b, 0x69, 0x6e,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6b, 0x69, 0x6e, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4b, 0x69, 0x6e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x01, 0x52, 0x12,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x64, 0x0a, 0x18, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x6b, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6b, 0x69,
	0x6e, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x4b, 0x69, 0x6e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08,
	0x01, 0x52, 0x15, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x69, 0x6e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x6b, 0x69, 0x6e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x4b, 0x69, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02,
	0x08, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x51, 0x0a, 0x10, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6b, 0x69,
	0x6e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a,
	0x0f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x01, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x11, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x43, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6b, 0x69, 0x6e,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x62, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x4a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6b, 0x69, 0x6e, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f,
	0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10,
	0x02, 0x22, 0x74, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x4c,
	0x4f, 0x43, 0x4b, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x53, 0x55, 0x46, 0x46, 0x49, 0x43, 0x49, 0x45,
	0x4e, 0x54, 0x5f, 0x46, 0x55, 0x4e, 0x44, 0x53, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x45, 0x51, 0x55, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x4e,
	0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x03, 0x22, 0x64, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1e, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xca, 0x9d, 0x25, 0x04, 0x08, 0x01, 0x20, 0x20, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xca, 0x9d, 0x25, 0x04, 0x08, 0x01, 0x20, 0x40,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e, 0x0a,
	0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0xca, 0x9d, 0x25, 0x06, 0x08, 0x01, 0x30,
	0x80, 0x80, 0x02, 0x52, 0x08, 0x72, 0x61, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x80, 0x01,
	0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x35, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xca, 0x9d, 0x25, 0x06, 0x08,
	0x01, 0x28, 0x40, 0x30, 0x40, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x38, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x0f, 0xca, 0x9d, 0x25, 0x0b, 0x08, 0x01, 0x41, 0x63, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x52, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x42, 0x7d, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x6b, 0x69, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x78,
	0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f,
	0x2f, 0x6b, 0x69, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x32, 0x3b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kin_transaction_v2_transaction_common_proto_rawDescOnce sync.Once
	file_kin_transaction_v2_transaction_common_proto_rawDescData = file_kin_transaction_v2_transaction_common_proto_rawDesc
)

func file_kin_transaction_v2_transaction_common_proto_rawDescGZIP() []byte {
	file_kin_transaction_v2_transaction_common_proto_rawDescOnce.Do(func() {
		file_kin_transaction_v2_transaction_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_kin_transaction_v2_transaction_common_proto_rawDescData)
	})
	return file_kin_transaction_v2_transaction_common_proto_rawDescData
}

var file_kin_transaction_v2_transaction_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_kin_transaction_v2_transaction_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_kin_transaction_v2_transaction_common_proto_goTypes = []interface{}{
	(TransactionPaymentInfo_TransactionStatus)(0),        // 0: common.kin.transaction.v2.TransactionPaymentInfo.TransactionStatus
	(TransactionPaymentInfo_TransactionFailureReason)(0), // 1: common.kin.transaction.v2.TransactionPaymentInfo.TransactionFailureReason
	(*TransactionPaymentInfo)(nil),                       // 2: common.kin.transaction.v2.TransactionPaymentInfo
	(*TransactionDisplayInfo)(nil),                       // 3: common.kin.transaction.v2.TransactionDisplayInfo
	(*TransactionEnvelope)(nil),                          // 4: common.kin.transaction.v2.TransactionEnvelope
	(*TransactionId)(nil),                                // 5: common.kin.transaction.v2.TransactionId
	(*v2.KinAccountId)(nil),                              // 6: common.kin.account.v2.KinAccountId
	(*v21.KinAmount)(nil),                                // 7: common.kin.payment.v2.KinAmount
	(v21.TransactionType)(0),                             // 8: common.kin.payment.v2.TransactionType
	(*timestamp.Timestamp)(nil),                          // 9: google.protobuf.Timestamp
}
var file_kin_transaction_v2_transaction_common_proto_depIdxs = []int32{
	5, // 0: common.kin.transaction.v2.TransactionPaymentInfo.transaction_id:type_name -> common.kin.transaction.v2.TransactionId
	6, // 1: common.kin.transaction.v2.TransactionPaymentInfo.sender_kin_account_id:type_name -> common.kin.account.v2.KinAccountId
	6, // 2: common.kin.transaction.v2.TransactionPaymentInfo.recipient_kin_account_id:type_name -> common.kin.account.v2.KinAccountId
	7, // 3: common.kin.transaction.v2.TransactionPaymentInfo.amount:type_name -> common.kin.payment.v2.KinAmount
	8, // 4: common.kin.transaction.v2.TransactionPaymentInfo.transaction_type:type_name -> common.kin.payment.v2.TransactionType
	9, // 5: common.kin.transaction.v2.TransactionPaymentInfo.submission_time:type_name -> google.protobuf.Timestamp
	9, // 6: common.kin.transaction.v2.TransactionPaymentInfo.confirmation_time:type_name -> google.protobuf.Timestamp
	0, // 7: common.kin.transaction.v2.TransactionPaymentInfo.status:type_name -> common.kin.transaction.v2.TransactionPaymentInfo.TransactionStatus
	1, // 8: common.kin.transaction.v2.TransactionPaymentInfo.reason:type_name -> common.kin.transaction.v2.TransactionPaymentInfo.TransactionFailureReason
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_kin_transaction_v2_transaction_common_proto_init() }
func file_kin_transaction_v2_transaction_common_proto_init() {
	if File_kin_transaction_v2_transaction_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kin_transaction_v2_transaction_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionPaymentInfo); i {
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
		file_kin_transaction_v2_transaction_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionDisplayInfo); i {
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
		file_kin_transaction_v2_transaction_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionEnvelope); i {
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
		file_kin_transaction_v2_transaction_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionId); i {
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
			RawDescriptor: file_kin_transaction_v2_transaction_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kin_transaction_v2_transaction_common_proto_goTypes,
		DependencyIndexes: file_kin_transaction_v2_transaction_common_proto_depIdxs,
		EnumInfos:         file_kin_transaction_v2_transaction_common_proto_enumTypes,
		MessageInfos:      file_kin_transaction_v2_transaction_common_proto_msgTypes,
	}.Build()
	File_kin_transaction_v2_transaction_common_proto = out.File
	file_kin_transaction_v2_transaction_common_proto_rawDesc = nil
	file_kin_transaction_v2_transaction_common_proto_goTypes = nil
	file_kin_transaction_v2_transaction_common_proto_depIdxs = nil
}
