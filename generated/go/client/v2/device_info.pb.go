// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: client/v2/device_info.proto

package client

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type DeviceInfoEnvelope_AlgorithmIdentifier int32

const (
	// In java land, this is "RSA/ECB/OAEPWithSHA1AndMGF1Padding" for asymmetric encryption of the symmetric key
	// (thats RSA_PKCS1_OAEP_PADDING in OpenSSL) and "AES/CBC/PKCS5Padding" for the encrypted payload.
	DeviceInfoEnvelope_RSA_OAEP_WITH_AES_128_CBC DeviceInfoEnvelope_AlgorithmIdentifier = 0
)

// Enum value maps for DeviceInfoEnvelope_AlgorithmIdentifier.
var (
	DeviceInfoEnvelope_AlgorithmIdentifier_name = map[int32]string{
		0: "RSA_OAEP_WITH_AES_128_CBC",
	}
	DeviceInfoEnvelope_AlgorithmIdentifier_value = map[string]int32{
		"RSA_OAEP_WITH_AES_128_CBC": 0,
	}
)

func (x DeviceInfoEnvelope_AlgorithmIdentifier) Enum() *DeviceInfoEnvelope_AlgorithmIdentifier {
	p := new(DeviceInfoEnvelope_AlgorithmIdentifier)
	*p = x
	return p
}

func (x DeviceInfoEnvelope_AlgorithmIdentifier) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceInfoEnvelope_AlgorithmIdentifier) Descriptor() protoreflect.EnumDescriptor {
	return file_client_v2_device_info_proto_enumTypes[0].Descriptor()
}

func (DeviceInfoEnvelope_AlgorithmIdentifier) Type() protoreflect.EnumType {
	return &file_client_v2_device_info_proto_enumTypes[0]
}

func (x DeviceInfoEnvelope_AlgorithmIdentifier) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceInfoEnvelope_AlgorithmIdentifier.Descriptor instead.
func (DeviceInfoEnvelope_AlgorithmIdentifier) EnumDescriptor() ([]byte, []int) {
	return file_client_v2_device_info_proto_rawDescGZIP(), []int{0, 0}
}

type DeviceInfoEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier for the asymmetric public key being used
	EncryptionKeyVersion uint32                                 `protobuf:"varint,1,opt,name=encryption_key_version,json=encryptionKeyVersion,proto3" json:"encryption_key_version,omitempty"`
	AlgorithmIdentifier  DeviceInfoEnvelope_AlgorithmIdentifier `protobuf:"varint,2,opt,name=algorithm_identifier,json=algorithmIdentifier,proto3,enum=common.client.v2.DeviceInfoEnvelope_AlgorithmIdentifier" json:"algorithm_identifier,omitempty"`
	// The symmetric key encrypted by the servers asymmetric public key. This key should be generated in a securely random
	// way and be used only once. The IV is expected to be all zeros (which is
	// safe provided the key is used only once).
	// Currently only using AES 128, so key size is 128 bits = 16 bytes
	// However, the key is encrypted using RSA and using a 2048 bit = 256 byte modulus
	// So the resulting encrypted key size will be 256 bytes
	EncryptedSymmetricKey []byte `protobuf:"bytes,3,opt,name=encrypted_symmetric_key,json=encryptedSymmetricKey,proto3" json:"encrypted_symmetric_key,omitempty"`
	// The protobuf serialized DeviceInfo, encrypted with the symmetric key.
	// 27 field tags - 15 bytes + (2 * 12) bytes = 39 bytes
	// 11 strings - upperbound 1599 bytes
	// 16 ints - upperbound 160 bytes
	// 1 bool - 1 byte
	// 1 Timestamp - 20 bytes + 2 bytes for field tags
	// Round up to 2048 bytes
	EncryptedDeviceInfo []byte `protobuf:"bytes,4,opt,name=encrypted_device_info,json=encryptedDeviceInfo,proto3" json:"encrypted_device_info,omitempty"`
}

func (x *DeviceInfoEnvelope) Reset() {
	*x = DeviceInfoEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_v2_device_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfoEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfoEnvelope) ProtoMessage() {}

func (x *DeviceInfoEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_client_v2_device_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfoEnvelope.ProtoReflect.Descriptor instead.
func (*DeviceInfoEnvelope) Descriptor() ([]byte, []int) {
	return file_client_v2_device_info_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceInfoEnvelope) GetEncryptionKeyVersion() uint32 {
	if x != nil {
		return x.EncryptionKeyVersion
	}
	return 0
}

func (x *DeviceInfoEnvelope) GetAlgorithmIdentifier() DeviceInfoEnvelope_AlgorithmIdentifier {
	if x != nil {
		return x.AlgorithmIdentifier
	}
	return DeviceInfoEnvelope_RSA_OAEP_WITH_AES_128_CBC
}

func (x *DeviceInfoEnvelope) GetEncryptedSymmetricKey() []byte {
	if x != nil {
		return x.EncryptedSymmetricKey
	}
	return nil
}

func (x *DeviceInfoEnvelope) GetEncryptedDeviceInfo() []byte {
	if x != nil {
		return x.EncryptedDeviceInfo
	}
	return nil
}

type DeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A mandatory time of creation which will be validated against the server
	// to prevent (but not completely) replay attacks.
	CreatedAt                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	InstallReferrer           string               `protobuf:"bytes,2,opt,name=install_referrer,json=installReferrer,proto3" json:"install_referrer,omitempty"`
	InstallDate               *timestamp.Timestamp `protobuf:"bytes,3,opt,name=install_date,json=installDate,proto3" json:"install_date,omitempty"`
	LoginsSinceInstall        uint64               `protobuf:"varint,4,opt,name=logins_since_install,json=loginsSinceInstall,proto3" json:"logins_since_install,omitempty"`
	RegistrationsSinceInstall uint64               `protobuf:"varint,5,opt,name=registrations_since_install,json=registrationsSinceInstall,proto3" json:"registrations_since_install,omitempty"`
	// Carrier code e.g. "41201", "60302", "54411", etc.
	Operator   string `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	AndroidSdk string `protobuf:"bytes,7,opt,name=android_sdk,json=androidSdk,proto3" json:"android_sdk,omitempty"`
	// Android only. e.g. "2ac70f97", "HKL1N5CZ"
	SerialNumber string `protobuf:"bytes,8,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	// Android only i.e. "android-build",
	BuildUser string `protobuf:"bytes,9,opt,name=build_user,json=buildUser,proto3" json:"build_user,omitempty"`
	// e.g. "wpef12.hot.corp.google.com"
	BuildHost string `protobuf:"bytes,10,opt,name=build_host,json=buildHost,proto3" json:"build_host,omitempty"`
	// Android only. e.g. "ac101cpl", "msm8998", "A2010-a"
	BoardName string `protobuf:"bytes,11,opt,name=board_name,json=boardName,proto3" json:"board_name,omitempty"`
	// Current battery level (0-255).
	// If monitoring isn't enabled on iPhone, -255.
	// Garbage values on android will be set to -1.
	BatteryLevel int32 `protobuf:"varint,12,opt,name=battery_level,json=batteryLevel,proto3" json:"battery_level,omitempty"`
	// Is the battery charging?
	BatteryIsCharging bool `protobuf:"varint,13,opt,name=battery_is_charging,json=batteryIsCharging,proto3" json:"battery_is_charging,omitempty"`
	// Current battery temp (10ths of a degree C). Android only.
	// Not guaranteed to be postive.
	BatteryTemp int32 `protobuf:"varint,14,opt,name=battery_temp,json=batteryTemp,proto3" json:"battery_temp,omitempty"`
	// Battery voltage in mA. Android only.
	BatteryVoltage int32 `protobuf:"varint,15,opt,name=battery_voltage,json=batteryVoltage,proto3" json:"battery_voltage,omitempty"`
	// System device version (i.e. "8.0.0", "9")
	DeviceVersion string `protobuf:"bytes,16,opt,name=device_version,json=deviceVersion,proto3" json:"device_version,omitempty"`
	// System device model (i.e. "HTC Desire 626", "SAMSUNG-SM-G900A")
	DeviceModel string `protobuf:"bytes,17,opt,name=device_model,json=deviceModel,proto3" json:"device_model,omitempty"`
	// System device type (i.e. "star2qltesq", "x86_64", "iPhone6,1")
	DeviceType string `protobuf:"bytes,18,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	// System device brand (Android only)
	DeviceBrand string `protobuf:"bytes,19,opt,name=device_brand,json=deviceBrand,proto3" json:"device_brand,omitempty"`
	// The duration since the last device reset
	SystemUptime *duration.Duration `protobuf:"bytes,20,opt,name=system_uptime,json=systemUptime,proto3" json:"system_uptime,omitempty"`
	// The difference between device uptime and actual, steady uptime
	ClockDrift *duration.Duration `protobuf:"bytes,21,opt,name=clock_drift,json=clockDrift,proto3" json:"clock_drift,omitempty"`
	// Screen brightness (0-255, where 255 is 100% brightness)
	ScreenBrightness int32 `protobuf:"varint,22,opt,name=screen_brightness,json=screenBrightness,proto3" json:"screen_brightness,omitempty"`
	// The number of processors onboard the device.
	NumberOfProcessors uint32 `protobuf:"varint,23,opt,name=number_of_processors,json=numberOfProcessors,proto3" json:"number_of_processors,omitempty"`
	// Total disk space in megabytes (MB). IPHONE ONLY.
	TotalDiskSpace uint64 `protobuf:"varint,24,opt,name=total_disk_space,json=totalDiskSpace,proto3" json:"total_disk_space,omitempty"`
	// Total used disk space in megabytes (MB). IPHONE ONLY.
	UsedDiskSpace uint64 `protobuf:"varint,25,opt,name=used_disk_space,json=usedDiskSpace,proto3" json:"used_disk_space,omitempty"`
	// Total memory (RAM) available in megabytes (MB).
	TotalMemory uint32 `protobuf:"varint,26,opt,name=total_memory,json=totalMemory,proto3" json:"total_memory,omitempty"`
	// Total memory (RAM) currently in use in megabytes (MB).
	// Not guaranteed to be positive.
	UsedMemory int32 `protobuf:"varint,27,opt,name=used_memory,json=usedMemory,proto3" json:"used_memory,omitempty"`
	// Benchmarks
	HashCashTime uint32 `protobuf:"varint,28,opt,name=hash_cash_time,json=hashCashTime,proto3" json:"hash_cash_time,omitempty"`
}

func (x *DeviceInfo) Reset() {
	*x = DeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_v2_device_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfo) ProtoMessage() {}

func (x *DeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_client_v2_device_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfo.ProtoReflect.Descriptor instead.
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return file_client_v2_device_info_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceInfo) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DeviceInfo) GetInstallReferrer() string {
	if x != nil {
		return x.InstallReferrer
	}
	return ""
}

func (x *DeviceInfo) GetInstallDate() *timestamp.Timestamp {
	if x != nil {
		return x.InstallDate
	}
	return nil
}

func (x *DeviceInfo) GetLoginsSinceInstall() uint64 {
	if x != nil {
		return x.LoginsSinceInstall
	}
	return 0
}

func (x *DeviceInfo) GetRegistrationsSinceInstall() uint64 {
	if x != nil {
		return x.RegistrationsSinceInstall
	}
	return 0
}

func (x *DeviceInfo) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *DeviceInfo) GetAndroidSdk() string {
	if x != nil {
		return x.AndroidSdk
	}
	return ""
}

func (x *DeviceInfo) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *DeviceInfo) GetBuildUser() string {
	if x != nil {
		return x.BuildUser
	}
	return ""
}

func (x *DeviceInfo) GetBuildHost() string {
	if x != nil {
		return x.BuildHost
	}
	return ""
}

func (x *DeviceInfo) GetBoardName() string {
	if x != nil {
		return x.BoardName
	}
	return ""
}

func (x *DeviceInfo) GetBatteryLevel() int32 {
	if x != nil {
		return x.BatteryLevel
	}
	return 0
}

func (x *DeviceInfo) GetBatteryIsCharging() bool {
	if x != nil {
		return x.BatteryIsCharging
	}
	return false
}

func (x *DeviceInfo) GetBatteryTemp() int32 {
	if x != nil {
		return x.BatteryTemp
	}
	return 0
}

func (x *DeviceInfo) GetBatteryVoltage() int32 {
	if x != nil {
		return x.BatteryVoltage
	}
	return 0
}

func (x *DeviceInfo) GetDeviceVersion() string {
	if x != nil {
		return x.DeviceVersion
	}
	return ""
}

func (x *DeviceInfo) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

func (x *DeviceInfo) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *DeviceInfo) GetDeviceBrand() string {
	if x != nil {
		return x.DeviceBrand
	}
	return ""
}

func (x *DeviceInfo) GetSystemUptime() *duration.Duration {
	if x != nil {
		return x.SystemUptime
	}
	return nil
}

func (x *DeviceInfo) GetClockDrift() *duration.Duration {
	if x != nil {
		return x.ClockDrift
	}
	return nil
}

func (x *DeviceInfo) GetScreenBrightness() int32 {
	if x != nil {
		return x.ScreenBrightness
	}
	return 0
}

func (x *DeviceInfo) GetNumberOfProcessors() uint32 {
	if x != nil {
		return x.NumberOfProcessors
	}
	return 0
}

func (x *DeviceInfo) GetTotalDiskSpace() uint64 {
	if x != nil {
		return x.TotalDiskSpace
	}
	return 0
}

func (x *DeviceInfo) GetUsedDiskSpace() uint64 {
	if x != nil {
		return x.UsedDiskSpace
	}
	return 0
}

func (x *DeviceInfo) GetTotalMemory() uint32 {
	if x != nil {
		return x.TotalMemory
	}
	return 0
}

func (x *DeviceInfo) GetUsedMemory() int32 {
	if x != nil {
		return x.UsedMemory
	}
	return 0
}

func (x *DeviceInfo) GetHashCashTime() uint32 {
	if x != nil {
		return x.HashCashTime
	}
	return 0
}

var File_client_v2_device_info_proto protoreflect.FileDescriptor

var file_client_v2_device_info_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x1a,
	0x11, 0x6b, 0x69, 0x6b, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2,
	0x02, 0x0a, 0x12, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x6b, 0x0a, 0x14, 0x61,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x2e,
	0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x52, 0x13, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x17, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0c, 0xca, 0x9d, 0x25, 0x08, 0x08,
	0x01, 0x28, 0x80, 0x02, 0x30, 0x80, 0x02, 0x52, 0x15, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x53, 0x79, 0x6d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x3d,
	0x0a, 0x15, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x09, 0xca,
	0x9d, 0x25, 0x05, 0x08, 0x01, 0x30, 0x80, 0x10, 0x52, 0x13, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x34, 0x0a,
	0x13, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x53, 0x41, 0x5f, 0x4f, 0x41, 0x45, 0x50,
	0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x41, 0x45, 0x53, 0x5f, 0x31, 0x32, 0x38, 0x5f, 0x43, 0x42,
	0x43, 0x10, 0x00, 0x22, 0xc4, 0x0a, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x32, 0x0a,
	0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xca, 0x9d, 0x25, 0x03, 0x30, 0xff, 0x01,
	0x52, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65,
	0x72, 0x12, 0x3d, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x30, 0x0a, 0x14, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x12, 0x3e, 0x0a, 0x1b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x19, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x12, 0x22, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xca, 0x9d, 0x25, 0x02, 0x30, 0x40, 0x52, 0x08, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0b, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69,
	0x64, 0x5f, 0x73, 0x64, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xca, 0x9d, 0x25,
	0x02, 0x30, 0x40, 0x52, 0x0a, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x53, 0x64, 0x6b, 0x12,
	0x2c, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xca, 0x9d, 0x25, 0x03, 0x30, 0x80, 0x02, 0x52,
	0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a,
	0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xca, 0x9d, 0x25, 0x03, 0x30, 0x80, 0x01, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0a, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xca, 0x9d, 0x25, 0x03, 0x30,
	0x80, 0x02, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x0a, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xca, 0x9d, 0x25, 0x03, 0x30, 0x80, 0x01, 0x52, 0x09, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0d, 0xca, 0x9d,
	0x25, 0x09, 0x41, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x0c, 0x62, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x62, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x49,
	0x73, 0x43, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x30, 0x0a, 0x0c, 0x62, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0d, 0xca, 0x9d, 0x25, 0x09, 0x41, 0x20, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x0b,
	0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x12, 0x36, 0x0a, 0x0f, 0x62,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x5f, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x0d, 0xca, 0x9d, 0x25, 0x09, 0x41, 0x10, 0x27, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x52, 0x0e, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x56, 0x6f, 0x6c, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xca, 0x9d, 0x25,
	0x02, 0x30, 0x40, 0x52, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xca, 0x9d, 0x25, 0x03, 0x30, 0x80,
	0x01, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x28,
	0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xca, 0x9d, 0x25, 0x03, 0x30, 0x80, 0x01, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x0c, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xca, 0x9d, 0x25, 0x03, 0x30, 0x80, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x3e, 0x0a, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x72,
	0x69, 0x66, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x72, 0x69, 0x66, 0x74,
	0x12, 0x3a, 0x0a, 0x11, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x5f, 0x62, 0x72, 0x69, 0x67, 0x68,
	0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0d, 0xca, 0x9d, 0x25,
	0x09, 0x41, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x10, 0x73, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x42, 0x72, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x14,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x4f, 0x66, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x12, 0x28,
	0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44,
	0x69, 0x73, 0x6b, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x64,
	0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x64, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x63, 0x61, 0x73,
	0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x68, 0x61,
	0x73, 0x68, 0x43, 0x61, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x96, 0x01, 0xaa, 0xa3, 0x2a,
	0x02, 0x08, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x67, 0x65, 0x6e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x32, 0x42, 0x0f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x78, 0x69, 0x70, 0x68, 0x69, 0x61, 0x73,
	0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x32, 0x3b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xa2, 0x02, 0x15, 0x4b, 0x50,
	0x42, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_v2_device_info_proto_rawDescOnce sync.Once
	file_client_v2_device_info_proto_rawDescData = file_client_v2_device_info_proto_rawDesc
)

func file_client_v2_device_info_proto_rawDescGZIP() []byte {
	file_client_v2_device_info_proto_rawDescOnce.Do(func() {
		file_client_v2_device_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_v2_device_info_proto_rawDescData)
	})
	return file_client_v2_device_info_proto_rawDescData
}

var file_client_v2_device_info_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_client_v2_device_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_v2_device_info_proto_goTypes = []interface{}{
	(DeviceInfoEnvelope_AlgorithmIdentifier)(0), // 0: common.client.v2.DeviceInfoEnvelope.AlgorithmIdentifier
	(*DeviceInfoEnvelope)(nil),                  // 1: common.client.v2.DeviceInfoEnvelope
	(*DeviceInfo)(nil),                          // 2: common.client.v2.DeviceInfo
	(*timestamp.Timestamp)(nil),                 // 3: google.protobuf.Timestamp
	(*duration.Duration)(nil),                   // 4: google.protobuf.Duration
}
var file_client_v2_device_info_proto_depIdxs = []int32{
	0, // 0: common.client.v2.DeviceInfoEnvelope.algorithm_identifier:type_name -> common.client.v2.DeviceInfoEnvelope.AlgorithmIdentifier
	3, // 1: common.client.v2.DeviceInfo.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: common.client.v2.DeviceInfo.install_date:type_name -> google.protobuf.Timestamp
	4, // 3: common.client.v2.DeviceInfo.system_uptime:type_name -> google.protobuf.Duration
	4, // 4: common.client.v2.DeviceInfo.clock_drift:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_client_v2_device_info_proto_init() }
func file_client_v2_device_info_proto_init() {
	if File_client_v2_device_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_v2_device_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfoEnvelope); i {
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
		file_client_v2_device_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfo); i {
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
			RawDescriptor: file_client_v2_device_info_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_v2_device_info_proto_goTypes,
		DependencyIndexes: file_client_v2_device_info_proto_depIdxs,
		EnumInfos:         file_client_v2_device_info_proto_enumTypes,
		MessageInfos:      file_client_v2_device_info_proto_msgTypes,
	}.Build()
	File_client_v2_device_info_proto = out.File
	file_client_v2_device_info_proto_rawDesc = nil
	file_client_v2_device_info_proto_goTypes = nil
	file_client_v2_device_info_proto_depIdxs = nil
}
