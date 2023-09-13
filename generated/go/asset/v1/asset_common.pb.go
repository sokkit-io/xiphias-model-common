// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: asset/v1/asset_common.proto

package asset

import (
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

// Used to request assets of a specific pixel density from a service
type PixelDensity int32

const (
	PixelDensity_NODPI           PixelDensity = 0
	PixelDensity_ANDROID_LDPI    PixelDensity = 1
	PixelDensity_ANDROID_MDPI    PixelDensity = 2
	PixelDensity_ANDROID_HDPI    PixelDensity = 3
	PixelDensity_ANDROID_XHDPI   PixelDensity = 4
	PixelDensity_ANDROID_XXHDPI  PixelDensity = 5
	PixelDensity_ANDROID_XXXHDPI PixelDensity = 6
	PixelDensity_IOS_X1          PixelDensity = 7
	PixelDensity_IOS_X2          PixelDensity = 8
	PixelDensity_IOS_X3          PixelDensity = 9
)

// Enum value maps for PixelDensity.
var (
	PixelDensity_name = map[int32]string{
		0: "NODPI",
		1: "ANDROID_LDPI",
		2: "ANDROID_MDPI",
		3: "ANDROID_HDPI",
		4: "ANDROID_XHDPI",
		5: "ANDROID_XXHDPI",
		6: "ANDROID_XXXHDPI",
		7: "IOS_X1",
		8: "IOS_X2",
		9: "IOS_X3",
	}
	PixelDensity_value = map[string]int32{
		"NODPI":           0,
		"ANDROID_LDPI":    1,
		"ANDROID_MDPI":    2,
		"ANDROID_HDPI":    3,
		"ANDROID_XHDPI":   4,
		"ANDROID_XXHDPI":  5,
		"ANDROID_XXXHDPI": 6,
		"IOS_X1":          7,
		"IOS_X2":          8,
		"IOS_X3":          9,
	}
)

func (x PixelDensity) Enum() *PixelDensity {
	p := new(PixelDensity)
	*p = x
	return p
}

func (x PixelDensity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PixelDensity) Descriptor() protoreflect.EnumDescriptor {
	return file_asset_v1_asset_common_proto_enumTypes[0].Descriptor()
}

func (PixelDensity) Type() protoreflect.EnumType {
	return &file_asset_v1_asset_common_proto_enumTypes[0]
}

func (x PixelDensity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PixelDensity.Descriptor instead.
func (PixelDensity) EnumDescriptor() ([]byte, []int) {
	return file_asset_v1_asset_common_proto_rawDescGZIP(), []int{0}
}

type ProductContent_Type int32

const (
	ProductContent_UNKNOWN    ProductContent_Type = 0
	ProductContent_CHAT_THEME ProductContent_Type = 1
	ProductContent_AVATAR     ProductContent_Type = 2
)

// Enum value maps for ProductContent_Type.
var (
	ProductContent_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "CHAT_THEME",
		2: "AVATAR",
	}
	ProductContent_Type_value = map[string]int32{
		"UNKNOWN":    0,
		"CHAT_THEME": 1,
		"AVATAR":     2,
	}
)

func (x ProductContent_Type) Enum() *ProductContent_Type {
	p := new(ProductContent_Type)
	*p = x
	return p
}

func (x ProductContent_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductContent_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_asset_v1_asset_common_proto_enumTypes[1].Descriptor()
}

func (ProductContent_Type) Type() protoreflect.EnumType {
	return &file_asset_v1_asset_common_proto_enumTypes[1]
}

func (x ProductContent_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductContent_Type.Descriptor instead.
func (ProductContent_Type) EnumDescriptor() ([]byte, []int) {
	return file_asset_v1_asset_common_proto_rawDescGZIP(), []int{0, 0}
}

type MediaContent_Mimetype int32

const (
	MediaContent_UNKNOWN    MediaContent_Mimetype = 0
	MediaContent_IMAGE_JPEG MediaContent_Mimetype = 1
	MediaContent_IMAGE_PNG  MediaContent_Mimetype = 2
)

// Enum value maps for MediaContent_Mimetype.
var (
	MediaContent_Mimetype_name = map[int32]string{
		0: "UNKNOWN",
		1: "IMAGE_JPEG",
		2: "IMAGE_PNG",
	}
	MediaContent_Mimetype_value = map[string]int32{
		"UNKNOWN":    0,
		"IMAGE_JPEG": 1,
		"IMAGE_PNG":  2,
	}
)

func (x MediaContent_Mimetype) Enum() *MediaContent_Mimetype {
	p := new(MediaContent_Mimetype)
	*p = x
	return p
}

func (x MediaContent_Mimetype) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaContent_Mimetype) Descriptor() protoreflect.EnumDescriptor {
	return file_asset_v1_asset_common_proto_enumTypes[2].Descriptor()
}

func (MediaContent_Mimetype) Type() protoreflect.EnumType {
	return &file_asset_v1_asset_common_proto_enumTypes[2]
}

func (x MediaContent_Mimetype) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaContent_Mimetype.Descriptor instead.
func (MediaContent_Mimetype) EnumDescriptor() ([]byte, []int) {
	return file_asset_v1_asset_common_proto_rawDescGZIP(), []int{2, 0}
}

// Contains the universal details of a product
type ProductContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A map between the name of the asset and its content
	Assets map[string]*Asset   `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type   ProductContent_Type `protobuf:"varint,2,opt,name=type,proto3,enum=common.asset.v1.ProductContent_Type" json:"type,omitempty"`
}

func (x *ProductContent) Reset() {
	*x = ProductContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asset_v1_asset_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductContent) ProtoMessage() {}

func (x *ProductContent) ProtoReflect() protoreflect.Message {
	mi := &file_asset_v1_asset_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductContent.ProtoReflect.Descriptor instead.
func (*ProductContent) Descriptor() ([]byte, []int) {
	return file_asset_v1_asset_common_proto_rawDescGZIP(), []int{0}
}

func (x *ProductContent) GetAssets() map[string]*Asset {
	if x != nil {
		return x.Assets
	}
	return nil
}

func (x *ProductContent) GetType() ProductContent_Type {
	if x != nil {
		return x.Type
	}
	return ProductContent_UNKNOWN
}

// Contains the content of an asset (a component of a product). This model should be used for retrieving existing
// assets, not creating or updating them. media_content or simple_content (or both) must be set.
type Asset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaContent *MediaContent `protobuf:"bytes,10,opt,name=media_content,json=mediaContent,proto3" json:"media_content,omitempty"`
	// The preview for the media content being presented (optional)
	MediaContentPreview *MediaContent `protobuf:"bytes,11,opt,name=media_content_preview,json=mediaContentPreview,proto3" json:"media_content_preview,omitempty"`
	// If the asset is not a media object (e.g. a color hex code), its data will be stored in this map. The expected
	// max size for a value in simple_content is roughly 1024 bytes and there should be no more than 10 key-value pairs.
	SimpleContent map[string]string `protobuf:"bytes,12,rep,name=simple_content,json=simpleContent,proto3" json:"simple_content,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Asset) Reset() {
	*x = Asset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asset_v1_asset_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_asset_v1_asset_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_asset_v1_asset_common_proto_rawDescGZIP(), []int{1}
}

func (x *Asset) GetMediaContent() *MediaContent {
	if x != nil {
		return x.MediaContent
	}
	return nil
}

func (x *Asset) GetMediaContentPreview() *MediaContent {
	if x != nil {
		return x.MediaContentPreview
	}
	return nil
}

func (x *Asset) GetSimpleContent() map[string]string {
	if x != nil {
		return x.SimpleContent
	}
	return nil
}

// Contains the details of a media component (e.g. image, video) of an asset
type MediaContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentUrl string                `protobuf:"bytes,1,opt,name=content_url,json=contentUrl,proto3" json:"content_url,omitempty"`
	Mimetype   MediaContent_Mimetype `protobuf:"varint,2,opt,name=mimetype,proto3,enum=common.asset.v1.MediaContent_Mimetype" json:"mimetype,omitempty"`
}

func (x *MediaContent) Reset() {
	*x = MediaContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asset_v1_asset_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MediaContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MediaContent) ProtoMessage() {}

func (x *MediaContent) ProtoReflect() protoreflect.Message {
	mi := &file_asset_v1_asset_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MediaContent.ProtoReflect.Descriptor instead.
func (*MediaContent) Descriptor() ([]byte, []int) {
	return file_asset_v1_asset_common_proto_rawDescGZIP(), []int{2}
}

func (x *MediaContent) GetContentUrl() string {
	if x != nil {
		return x.ContentUrl
	}
	return ""
}

func (x *MediaContent) GetMimetype() MediaContent_Mimetype {
	if x != nil {
		return x.Mimetype
	}
	return MediaContent_UNKNOWN
}

var File_asset_v1_asset_common_proto protoreflect.FileDescriptor

var file_asset_v1_asset_common_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x19,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x02, 0x0a, 0x0e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x06,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x09, 0xd2, 0x9d, 0x25, 0x05, 0x28,
	0x01, 0x30, 0x80, 0x01, 0x52, 0x06, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x51, 0x0a, 0x0b, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2f, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x54, 0x48, 0x45, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x10, 0x02, 0x22, 0xbc, 0x02, 0x0a, 0x05, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x12, 0x42, 0x0a, 0x0d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x51, 0x0a, 0x15, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x13, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x5a, 0x0a, 0x0e, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08,
	0xd2, 0x9d, 0x25, 0x04, 0x28, 0x01, 0x30, 0x40, 0x52, 0x0d, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x40, 0x0a, 0x12, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb6, 0x01, 0x0a, 0x0c, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xca, 0x9d, 0x25, 0x05, 0x28, 0x01, 0x30, 0xff, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x42, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x22, 0x36, 0x0a, 0x08, 0x4d, 0x69,
	0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x4a, 0x50, 0x45,
	0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x50, 0x4e, 0x47,
	0x10, 0x02, 0x2a, 0xaf, 0x01, 0x0a, 0x0c, 0x50, 0x69, 0x78, 0x65, 0x6c, 0x44, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x79, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x4f, 0x44, 0x50, 0x49, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x4c, 0x44, 0x50, 0x49, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x4d, 0x44, 0x50, 0x49,
	0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x48, 0x44,
	0x50, 0x49, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f,
	0x58, 0x48, 0x44, 0x50, 0x49, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4e, 0x44, 0x52, 0x4f,
	0x49, 0x44, 0x5f, 0x58, 0x58, 0x48, 0x44, 0x50, 0x49, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x41,
	0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x58, 0x58, 0x58, 0x48, 0x44, 0x50, 0x49, 0x10, 0x06,
	0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4f, 0x53, 0x5f, 0x58, 0x31, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06,
	0x49, 0x4f, 0x53, 0x5f, 0x58, 0x32, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x4f, 0x53, 0x5f,
	0x58, 0x33, 0x10, 0x09, 0x42, 0x5c, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5a, 0x45, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d, 0x69,
	0x6f, 0x2f, 0x78, 0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2d,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_asset_v1_asset_common_proto_rawDescOnce sync.Once
	file_asset_v1_asset_common_proto_rawDescData = file_asset_v1_asset_common_proto_rawDesc
)

func file_asset_v1_asset_common_proto_rawDescGZIP() []byte {
	file_asset_v1_asset_common_proto_rawDescOnce.Do(func() {
		file_asset_v1_asset_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_asset_v1_asset_common_proto_rawDescData)
	})
	return file_asset_v1_asset_common_proto_rawDescData
}

var file_asset_v1_asset_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_asset_v1_asset_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_asset_v1_asset_common_proto_goTypes = []interface{}{
	(PixelDensity)(0),          // 0: common.asset.v1.PixelDensity
	(ProductContent_Type)(0),   // 1: common.asset.v1.ProductContent.Type
	(MediaContent_Mimetype)(0), // 2: common.asset.v1.MediaContent.Mimetype
	(*ProductContent)(nil),     // 3: common.asset.v1.ProductContent
	(*Asset)(nil),              // 4: common.asset.v1.Asset
	(*MediaContent)(nil),       // 5: common.asset.v1.MediaContent
	nil,                        // 6: common.asset.v1.ProductContent.AssetsEntry
	nil,                        // 7: common.asset.v1.Asset.SimpleContentEntry
}
var file_asset_v1_asset_common_proto_depIdxs = []int32{
	6, // 0: common.asset.v1.ProductContent.assets:type_name -> common.asset.v1.ProductContent.AssetsEntry
	1, // 1: common.asset.v1.ProductContent.type:type_name -> common.asset.v1.ProductContent.Type
	5, // 2: common.asset.v1.Asset.media_content:type_name -> common.asset.v1.MediaContent
	5, // 3: common.asset.v1.Asset.media_content_preview:type_name -> common.asset.v1.MediaContent
	7, // 4: common.asset.v1.Asset.simple_content:type_name -> common.asset.v1.Asset.SimpleContentEntry
	2, // 5: common.asset.v1.MediaContent.mimetype:type_name -> common.asset.v1.MediaContent.Mimetype
	4, // 6: common.asset.v1.ProductContent.AssetsEntry.value:type_name -> common.asset.v1.Asset
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_asset_v1_asset_common_proto_init() }
func file_asset_v1_asset_common_proto_init() {
	if File_asset_v1_asset_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_asset_v1_asset_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductContent); i {
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
		file_asset_v1_asset_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Asset); i {
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
		file_asset_v1_asset_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MediaContent); i {
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
			RawDescriptor: file_asset_v1_asset_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_asset_v1_asset_common_proto_goTypes,
		DependencyIndexes: file_asset_v1_asset_common_proto_depIdxs,
		EnumInfos:         file_asset_v1_asset_common_proto_enumTypes,
		MessageInfos:      file_asset_v1_asset_common_proto_msgTypes,
	}.Build()
	File_asset_v1_asset_common_proto = out.File
	file_asset_v1_asset_common_proto_rawDesc = nil
	file_asset_v1_asset_common_proto_goTypes = nil
	file_asset_v1_asset_common_proto_depIdxs = nil
}
