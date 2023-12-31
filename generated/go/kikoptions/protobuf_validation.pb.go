// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: protobuf_validation.proto

package kikoptions

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type FieldValidation_Order int32

const (
	// Ascending (smaller to bigger)
	FieldValidation_ASC FieldValidation_Order = 0
	// Descending (bigger to smaller)
	FieldValidation_DESC FieldValidation_Order = 1
)

// Enum value maps for FieldValidation_Order.
var (
	FieldValidation_Order_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	FieldValidation_Order_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x FieldValidation_Order) Enum() *FieldValidation_Order {
	p := new(FieldValidation_Order)
	*p = x
	return p
}

func (x FieldValidation_Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FieldValidation_Order) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_validation_proto_enumTypes[0].Descriptor()
}

func (FieldValidation_Order) Type() protoreflect.EnumType {
	return &file_protobuf_validation_proto_enumTypes[0]
}

func (x FieldValidation_Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *FieldValidation_Order) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = FieldValidation_Order(num)
	return nil
}

// Deprecated: Use FieldValidation_Order.Descriptor instead.
func (FieldValidation_Order) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_validation_proto_rawDescGZIP(), []int{0, 0}
}

type FieldValidation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO: Document semantics of this
	Mandatory *bool   `protobuf:"varint,1,opt,name=mandatory,def=0" json:"mandatory,omitempty"`
	Regex     *string `protobuf:"bytes,2,opt,name=regex" json:"regex,omitempty"`
	// For strings:
	MinCodepoints *int32 `protobuf:"varint,3,opt,name=min_codepoints,json=minCodepoints" json:"min_codepoints,omitempty"`
	MaxCodepoints *int32 `protobuf:"varint,4,opt,name=max_codepoints,json=maxCodepoints" json:"max_codepoints,omitempty"`
	// Valid for strings and byte arrays:
	MinByteLength  *int32   `protobuf:"varint,5,opt,name=min_byte_length,json=minByteLength" json:"min_byte_length,omitempty"`
	MaxByteLength  *int32   `protobuf:"varint,6,opt,name=max_byte_length,json=maxByteLength" json:"max_byte_length,omitempty"`
	MinVal         *int64   `protobuf:"fixed64,7,opt,name=min_val,json=minVal" json:"min_val,omitempty"`
	MaxVal         *int64   `protobuf:"fixed64,8,opt,name=max_val,json=maxVal" json:"max_val,omitempty"`
	MinDoubleVal   *float64 `protobuf:"fixed64,11,opt,name=min_double_val,json=minDoubleVal" json:"min_double_val,omitempty"`
	MaxDoubleVal   *float64 `protobuf:"fixed64,12,opt,name=max_double_val,json=maxDoubleVal" json:"max_double_val,omitempty"`
	MinRepetitions *uint32  `protobuf:"varint,15,opt,name=min_repetitions,json=minRepetitions" json:"min_repetitions,omitempty"`
	MaxRepetitions *uint32  `protobuf:"varint,16,opt,name=max_repetitions,json=maxRepetitions" json:"max_repetitions,omitempty"`
	// Valid on repeated fields ONLY (with some additional restrictions on what those fields can contain)
	Ordered *FieldValidation_Order `protobuf:"varint,17,opt,name=ordered,enum=kik.validation.FieldValidation_Order" json:"ordered,omitempty"`
}

// Default values for FieldValidation fields.
const (
	Default_FieldValidation_Mandatory = bool(false)
)

func (x *FieldValidation) Reset() {
	*x = FieldValidation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_validation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldValidation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldValidation) ProtoMessage() {}

func (x *FieldValidation) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_validation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldValidation.ProtoReflect.Descriptor instead.
func (*FieldValidation) Descriptor() ([]byte, []int) {
	return file_protobuf_validation_proto_rawDescGZIP(), []int{0}
}

func (x *FieldValidation) GetMandatory() bool {
	if x != nil && x.Mandatory != nil {
		return *x.Mandatory
	}
	return Default_FieldValidation_Mandatory
}

func (x *FieldValidation) GetRegex() string {
	if x != nil && x.Regex != nil {
		return *x.Regex
	}
	return ""
}

func (x *FieldValidation) GetMinCodepoints() int32 {
	if x != nil && x.MinCodepoints != nil {
		return *x.MinCodepoints
	}
	return 0
}

func (x *FieldValidation) GetMaxCodepoints() int32 {
	if x != nil && x.MaxCodepoints != nil {
		return *x.MaxCodepoints
	}
	return 0
}

func (x *FieldValidation) GetMinByteLength() int32 {
	if x != nil && x.MinByteLength != nil {
		return *x.MinByteLength
	}
	return 0
}

func (x *FieldValidation) GetMaxByteLength() int32 {
	if x != nil && x.MaxByteLength != nil {
		return *x.MaxByteLength
	}
	return 0
}

func (x *FieldValidation) GetMinVal() int64 {
	if x != nil && x.MinVal != nil {
		return *x.MinVal
	}
	return 0
}

func (x *FieldValidation) GetMaxVal() int64 {
	if x != nil && x.MaxVal != nil {
		return *x.MaxVal
	}
	return 0
}

func (x *FieldValidation) GetMinDoubleVal() float64 {
	if x != nil && x.MinDoubleVal != nil {
		return *x.MinDoubleVal
	}
	return 0
}

func (x *FieldValidation) GetMaxDoubleVal() float64 {
	if x != nil && x.MaxDoubleVal != nil {
		return *x.MaxDoubleVal
	}
	return 0
}

func (x *FieldValidation) GetMinRepetitions() uint32 {
	if x != nil && x.MinRepetitions != nil {
		return *x.MinRepetitions
	}
	return 0
}

func (x *FieldValidation) GetMaxRepetitions() uint32 {
	if x != nil && x.MaxRepetitions != nil {
		return *x.MaxRepetitions
	}
	return 0
}

func (x *FieldValidation) GetOrdered() FieldValidation_Order {
	if x != nil && x.Ordered != nil {
		return *x.Ordered
	}
	return FieldValidation_ASC
}

var file_protobuf_validation_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*FieldValidation)(nil),
		Field:         76249,
		Name:          "kik.validation.field_validation",
		Tag:           "bytes,76249,opt,name=field_validation",
		Filename:      "protobuf_validation.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*FieldValidation)(nil),
		Field:         76250,
		Name:          "kik.validation.map_key",
		Tag:           "bytes,76250,opt,name=map_key",
		Filename:      "protobuf_validation.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// Note tag is just a random number between 50000-99999
	// (see https://developers.google.com/protocol-buffers/docs/proto#customoptions)
	//
	// optional kik.validation.FieldValidation field_validation = 76249;
	E_FieldValidation = &file_protobuf_validation_proto_extTypes[0]
	// optional kik.validation.FieldValidation map_key = 76250;
	E_MapKey = &file_protobuf_validation_proto_extTypes[1] // Tag number 76251 is used in xiphias-model-common/proto/messagepath/v1/core_message_options.proto
)

var File_protobuf_validation_proto protoreflect.FileDescriptor

var file_protobuf_validation_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6b, 0x69, 0x6b,
	0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x04,
	0x0a, 0x0f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x23, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x6e,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x12, 0x25, 0x0a, 0x0e,
	0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x61, 0x78,
	0x43, 0x6f, 0x64, 0x65, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x69,
	0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x69, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x4c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x61, 0x78,
	0x42, 0x79, 0x74, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x69,
	0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x10, 0x52, 0x06, 0x6d, 0x69, 0x6e,
	0x56, 0x61, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x10, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0e,
	0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56,
	0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x69, 0x6e, 0x5f,
	0x72, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x52,
	0x65, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6b, 0x69,
	0x6b, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x22, 0x1a, 0x0a, 0x05, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x3a, 0x6b, 0x0a, 0x10, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd9, 0xd3, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x59, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xda,
	0xd3, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x61, 0x70, 0x4b, 0x65, 0x79, 0x42,
	0x67, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d, 0x69, 0x6f,
	0x2f, 0x78, 0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x6b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3b, 0x6b, 0x69,
	0x6b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
}

var (
	file_protobuf_validation_proto_rawDescOnce sync.Once
	file_protobuf_validation_proto_rawDescData = file_protobuf_validation_proto_rawDesc
)

func file_protobuf_validation_proto_rawDescGZIP() []byte {
	file_protobuf_validation_proto_rawDescOnce.Do(func() {
		file_protobuf_validation_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_validation_proto_rawDescData)
	})
	return file_protobuf_validation_proto_rawDescData
}

var file_protobuf_validation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_validation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_validation_proto_goTypes = []interface{}{
	(FieldValidation_Order)(0),      // 0: kik.validation.FieldValidation.Order
	(*FieldValidation)(nil),         // 1: kik.validation.FieldValidation
	(*descriptor.FieldOptions)(nil), // 2: google.protobuf.FieldOptions
}
var file_protobuf_validation_proto_depIdxs = []int32{
	0, // 0: kik.validation.FieldValidation.ordered:type_name -> kik.validation.FieldValidation.Order
	2, // 1: kik.validation.field_validation:extendee -> google.protobuf.FieldOptions
	2, // 2: kik.validation.map_key:extendee -> google.protobuf.FieldOptions
	1, // 3: kik.validation.field_validation:type_name -> kik.validation.FieldValidation
	1, // 4: kik.validation.map_key:type_name -> kik.validation.FieldValidation
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	3, // [3:5] is the sub-list for extension type_name
	1, // [1:3] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_validation_proto_init() }
func file_protobuf_validation_proto_init() {
	if File_protobuf_validation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_validation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldValidation); i {
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
			RawDescriptor: file_protobuf_validation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_validation_proto_goTypes,
		DependencyIndexes: file_protobuf_validation_proto_depIdxs,
		EnumInfos:         file_protobuf_validation_proto_enumTypes,
		MessageInfos:      file_protobuf_validation_proto_msgTypes,
		ExtensionInfos:    file_protobuf_validation_proto_extTypes,
	}.Build()
	File_protobuf_validation_proto = out.File
	file_protobuf_validation_proto_rawDesc = nil
	file_protobuf_validation_proto_goTypes = nil
	file_protobuf_validation_proto_depIdxs = nil
}
