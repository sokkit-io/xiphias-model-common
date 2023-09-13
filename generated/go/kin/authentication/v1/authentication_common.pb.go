// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: kin/authentication/v1/authentication_common.proto

package authentication

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

// OfferId contains the id of any offer that's meant to be paid for on the Kin marketplace
type OfferId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*OfferId_ProductId
	//	*OfferId_FeatureOfferId
	Value isOfferId_Value `protobuf_oneof:"value"`
}

func (x *OfferId) Reset() {
	*x = OfferId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kin_authentication_v1_authentication_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfferId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferId) ProtoMessage() {}

func (x *OfferId) ProtoReflect() protoreflect.Message {
	mi := &file_kin_authentication_v1_authentication_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfferId.ProtoReflect.Descriptor instead.
func (*OfferId) Descriptor() ([]byte, []int) {
	return file_kin_authentication_v1_authentication_common_proto_rawDescGZIP(), []int{0}
}

func (m *OfferId) GetValue() isOfferId_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *OfferId) GetProductId() *_go.XiUuid {
	if x, ok := x.GetValue().(*OfferId_ProductId); ok {
		return x.ProductId
	}
	return nil
}

func (x *OfferId) GetFeatureOfferId() string {
	if x, ok := x.GetValue().(*OfferId_FeatureOfferId); ok {
		return x.FeatureOfferId
	}
	return ""
}

type isOfferId_Value interface {
	isOfferId_Value()
}

type OfferId_ProductId struct {
	// Uuid of a AMS product
	ProductId *_go.XiUuid `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3,oneof"`
}

type OfferId_FeatureOfferId struct {
	// The offer id of a feature - used for JWT generation and tracking in the Kin SDK.
	// The offer id should ideally be in the format 'FEATURE_ID-TIMESTAMP'. Timestamp is epoch time in milliseconds
	// This will help verify payment confirmation JWTs
	FeatureOfferId string `protobuf:"bytes,2,opt,name=feature_offer_id,json=featureOfferId,proto3,oneof"`
}

func (*OfferId_ProductId) isOfferId_Value() {}

func (*OfferId_FeatureOfferId) isOfferId_Value() {}

// An offer JWT is generated by the authentication service in response to a service requesting a JWT for a Kin offer.
// Currently, the product data service is the main user of this JWT.
type OfferJwt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of an offer. An offer as an entity offered by Kik that has a Kin value and can be purchased. Currently,
	// product Ids are supported.
	Id *OfferId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// JWT generated by the authentication service for an offer.
	Jwt *v1.XiJWT `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *OfferJwt) Reset() {
	*x = OfferJwt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kin_authentication_v1_authentication_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfferJwt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfferJwt) ProtoMessage() {}

func (x *OfferJwt) ProtoReflect() protoreflect.Message {
	mi := &file_kin_authentication_v1_authentication_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfferJwt.ProtoReflect.Descriptor instead.
func (*OfferJwt) Descriptor() ([]byte, []int) {
	return file_kin_authentication_v1_authentication_common_proto_rawDescGZIP(), []int{1}
}

func (x *OfferJwt) GetId() *OfferId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *OfferJwt) GetJwt() *v1.XiJWT {
	if x != nil {
		return x.Jwt
	}
	return nil
}

var File_kin_authentication_v1_authentication_common_proto protoreflect.FileDescriptor

var file_kin_authentication_v1_authentication_common_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6b, 0x69, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6b, 0x69, 0x6e, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x07, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x58, 0x69, 0x55, 0x75, 0x69, 0x64, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x10, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xca,
	0x9d, 0x25, 0x05, 0x28, 0x01, 0x30, 0x80, 0x01, 0x48, 0x00, 0x52, 0x0e, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x75, 0x0a, 0x08, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x4a, 0x77, 0x74, 0x12,
	0x3d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6b, 0x69, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x49, 0x64, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x08, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a,
	0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x58, 0x69, 0x4a, 0x57, 0x54, 0x42, 0x06, 0xca,
	0x9d, 0x25, 0x02, 0x08, 0x01, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x42, 0x7f, 0x0a, 0x20, 0x63, 0x6f,
	0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x6b, 0x69, 0x6e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5a, 0x5b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69,
	0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x78, 0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_kin_authentication_v1_authentication_common_proto_rawDescOnce sync.Once
	file_kin_authentication_v1_authentication_common_proto_rawDescData = file_kin_authentication_v1_authentication_common_proto_rawDesc
)

func file_kin_authentication_v1_authentication_common_proto_rawDescGZIP() []byte {
	file_kin_authentication_v1_authentication_common_proto_rawDescOnce.Do(func() {
		file_kin_authentication_v1_authentication_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_kin_authentication_v1_authentication_common_proto_rawDescData)
	})
	return file_kin_authentication_v1_authentication_common_proto_rawDescData
}

var file_kin_authentication_v1_authentication_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kin_authentication_v1_authentication_common_proto_goTypes = []interface{}{
	(*OfferId)(nil),    // 0: common.kin.authentication.v1.OfferId
	(*OfferJwt)(nil),   // 1: common.kin.authentication.v1.OfferJwt
	(*_go.XiUuid)(nil), // 2: common.XiUuid
	(*v1.XiJWT)(nil),   // 3: common.v1.XiJWT
}
var file_kin_authentication_v1_authentication_common_proto_depIdxs = []int32{
	2, // 0: common.kin.authentication.v1.OfferId.product_id:type_name -> common.XiUuid
	0, // 1: common.kin.authentication.v1.OfferJwt.id:type_name -> common.kin.authentication.v1.OfferId
	3, // 2: common.kin.authentication.v1.OfferJwt.jwt:type_name -> common.v1.XiJWT
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kin_authentication_v1_authentication_common_proto_init() }
func file_kin_authentication_v1_authentication_common_proto_init() {
	if File_kin_authentication_v1_authentication_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kin_authentication_v1_authentication_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfferId); i {
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
		file_kin_authentication_v1_authentication_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfferJwt); i {
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
	file_kin_authentication_v1_authentication_common_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OfferId_ProductId)(nil),
		(*OfferId_FeatureOfferId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kin_authentication_v1_authentication_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kin_authentication_v1_authentication_common_proto_goTypes,
		DependencyIndexes: file_kin_authentication_v1_authentication_common_proto_depIdxs,
		MessageInfos:      file_kin_authentication_v1_authentication_common_proto_msgTypes,
	}.Build()
	File_kin_authentication_v1_authentication_common_proto = out.File
	file_kin_authentication_v1_authentication_common_proto_rawDesc = nil
	file_kin_authentication_v1_authentication_common_proto_goTypes = nil
	file_kin_authentication_v1_authentication_common_proto_depIdxs = nil
}
