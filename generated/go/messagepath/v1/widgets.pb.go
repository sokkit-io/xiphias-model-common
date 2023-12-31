// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: messagepath/v1/widgets.proto

package messagepath

import (
	_go "github.com/sokkit-io/xiphias-model-common/generated/go"
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

// WidgetAttachment defines the mechanism that allow bots to alter the behavior of dynamic
// content that supplements the chatbot experience.
//
// Specifying a Widget allows a bot to give the user additional information that provides
// more context to a conversation. TextWidget is an example of a widget that stays at the
// top of the screen, providing persistent information all users in a conversation.
//
// Up to 51 of these may be present in a single message to be delivered. The case where this can
// occur is in the case of a bot responding to a mention in a group containing 50 individuals, and
// also having a default.
//
// If a client receives multiple widgets in this list directed at a single user, the first one
// should be taken as the one to display. Support for displaying multiple widgets may be added in
// the future.
//
// See: https://docs.google.com/a/kik.com/document/d/1Y2tnA5KfCma0wmGoqnoqFKWTJZa-beU0mvU3v8A3UN0/edit
type WidgetAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Widgets []*Widget `protobuf:"bytes,1,rep,name=widgets,proto3" json:"widgets,omitempty"`
}

func (x *WidgetAttachment) Reset() {
	*x = WidgetAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagepath_v1_widgets_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WidgetAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WidgetAttachment) ProtoMessage() {}

func (x *WidgetAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_messagepath_v1_widgets_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WidgetAttachment.ProtoReflect.Descriptor instead.
func (*WidgetAttachment) Descriptor() ([]byte, []int) {
	return file_messagepath_v1_widgets_proto_rawDescGZIP(), []int{0}
}

func (x *WidgetAttachment) GetWidgets() []*Widget {
	if x != nil {
		return x.Widgets
	}
	return nil
}

type Widget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The user that will receive this particular Widget instance.
	//
	// In the event that this field is omitted, this widget should be sent to all users that will
	// receive the message containing this widget.
	To *_go.XiBareUserJid `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	// Types that are assignable to Type:
	//	*Widget_TextWidget
	Type isWidget_Type `protobuf_oneof:"type"`
}

func (x *Widget) Reset() {
	*x = Widget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagepath_v1_widgets_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Widget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Widget) ProtoMessage() {}

func (x *Widget) ProtoReflect() protoreflect.Message {
	mi := &file_messagepath_v1_widgets_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Widget.ProtoReflect.Descriptor instead.
func (*Widget) Descriptor() ([]byte, []int) {
	return file_messagepath_v1_widgets_proto_rawDescGZIP(), []int{1}
}

func (x *Widget) GetTo() *_go.XiBareUserJid {
	if x != nil {
		return x.To
	}
	return nil
}

func (m *Widget) GetType() isWidget_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Widget) GetTextWidget() *TextWidget {
	if x, ok := x.GetType().(*Widget_TextWidget); ok {
		return x.TextWidget
	}
	return nil
}

type isWidget_Type interface {
	isWidget_Type()
}

type Widget_TextWidget struct {
	TextWidget *TextWidget `protobuf:"bytes,32,opt,name=text_widget,json=textWidget,proto3,oneof"`
}

func (*Widget_TextWidget) isWidget_Type() {}

// TextWidget allow bots to provide text to be shown at the top of the
// conversation window.
type TextWidget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body  string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// You can attach a Keyboard to a TextWidget, which specifies
	// the keyboard to be shown when a reply button is pressed on the widget
	//
	// The `to` field of this Keyboard is ignored and overrided by the
	// `to` field of the widget
	Keyboard *Keyboard `protobuf:"bytes,3,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
}

func (x *TextWidget) Reset() {
	*x = TextWidget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messagepath_v1_widgets_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextWidget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextWidget) ProtoMessage() {}

func (x *TextWidget) ProtoReflect() protoreflect.Message {
	mi := &file_messagepath_v1_widgets_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextWidget.ProtoReflect.Descriptor instead.
func (*TextWidget) Descriptor() ([]byte, []int) {
	return file_messagepath_v1_widgets_proto_rawDescGZIP(), []int{2}
}

func (x *TextWidget) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *TextWidget) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TextWidget) GetKeyboard() *Keyboard {
	if x != nil {
		return x.Keyboard
	}
	return nil
}

var File_messagepath_v1_widgets_proto protoreflect.FileDescriptor

var file_messagepath_v1_widgets_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x2f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74,
	0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x10, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x07, 0x77, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x42, 0x0c, 0xca, 0x9d, 0x25, 0x08, 0x08, 0x01,
	0x78, 0x01, 0x80, 0x01, 0xe8, 0x07, 0x52, 0x07, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x22,
	0x85, 0x01, 0x0a, 0x06, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x58, 0x69, 0x42, 0x61, 0x72, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4a, 0x69, 0x64, 0x42, 0x06, 0xca,
	0x9d, 0x25, 0x02, 0x08, 0x00, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x44, 0x0a, 0x0b, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70,
	0x61, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65,
	0x74, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x42,
	0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x0a, 0x54, 0x65, 0x78, 0x74,
	0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xca, 0x9d, 0x25, 0x07, 0x08, 0x01, 0x28, 0x01, 0x30, 0xe8,
	0x07, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xca, 0x9d, 0x25, 0x05, 0x08, 0x00, 0x30, 0xe8,
	0x07, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x06, 0xca, 0x9d, 0x25,
	0x02, 0x08, 0x00, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x42, 0x75, 0x0a,
	0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x69, 0x6b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x70, 0x61, 0x74, 0x68, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6b, 0x6b, 0x69, 0x74, 0x2d, 0x69, 0x6f,
	0x2f, 0x78, 0x69, 0x70, 0x68, 0x69, 0x61, 0x73, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x76,
	0x31, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0xa2, 0x02, 0x04,
	0x4d, 0x50, 0x54, 0x48, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messagepath_v1_widgets_proto_rawDescOnce sync.Once
	file_messagepath_v1_widgets_proto_rawDescData = file_messagepath_v1_widgets_proto_rawDesc
)

func file_messagepath_v1_widgets_proto_rawDescGZIP() []byte {
	file_messagepath_v1_widgets_proto_rawDescOnce.Do(func() {
		file_messagepath_v1_widgets_proto_rawDescData = protoimpl.X.CompressGZIP(file_messagepath_v1_widgets_proto_rawDescData)
	})
	return file_messagepath_v1_widgets_proto_rawDescData
}

var file_messagepath_v1_widgets_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_messagepath_v1_widgets_proto_goTypes = []interface{}{
	(*WidgetAttachment)(nil),  // 0: common.messagepath.v1.WidgetAttachment
	(*Widget)(nil),            // 1: common.messagepath.v1.Widget
	(*TextWidget)(nil),        // 2: common.messagepath.v1.TextWidget
	(*_go.XiBareUserJid)(nil), // 3: common.XiBareUserJid
	(*Keyboard)(nil),          // 4: common.messagepath.v1.Keyboard
}
var file_messagepath_v1_widgets_proto_depIdxs = []int32{
	1, // 0: common.messagepath.v1.WidgetAttachment.widgets:type_name -> common.messagepath.v1.Widget
	3, // 1: common.messagepath.v1.Widget.to:type_name -> common.XiBareUserJid
	2, // 2: common.messagepath.v1.Widget.text_widget:type_name -> common.messagepath.v1.TextWidget
	4, // 3: common.messagepath.v1.TextWidget.keyboard:type_name -> common.messagepath.v1.Keyboard
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_messagepath_v1_widgets_proto_init() }
func file_messagepath_v1_widgets_proto_init() {
	if File_messagepath_v1_widgets_proto != nil {
		return
	}
	file_messagepath_v1_keyboards_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_messagepath_v1_widgets_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WidgetAttachment); i {
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
		file_messagepath_v1_widgets_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Widget); i {
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
		file_messagepath_v1_widgets_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextWidget); i {
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
	file_messagepath_v1_widgets_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Widget_TextWidget)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messagepath_v1_widgets_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messagepath_v1_widgets_proto_goTypes,
		DependencyIndexes: file_messagepath_v1_widgets_proto_depIdxs,
		MessageInfos:      file_messagepath_v1_widgets_proto_msgTypes,
	}.Build()
	File_messagepath_v1_widgets_proto = out.File
	file_messagepath_v1_widgets_proto_rawDesc = nil
	file_messagepath_v1_widgets_proto_goTypes = nil
	file_messagepath_v1_widgets_proto_depIdxs = nil
}
