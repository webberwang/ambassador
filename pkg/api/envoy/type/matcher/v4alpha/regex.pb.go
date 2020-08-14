// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/type/matcher/v4alpha/regex.proto

package envoy_type_matcher_v4alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RegexMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to EngineType:
	//	*RegexMatcher_GoogleRe2
	EngineType isRegexMatcher_EngineType `protobuf_oneof:"engine_type"`
	Regex      string                    `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
}

func (x *RegexMatcher) Reset() {
	*x = RegexMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_v4alpha_regex_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegexMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegexMatcher) ProtoMessage() {}

func (x *RegexMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_v4alpha_regex_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegexMatcher.ProtoReflect.Descriptor instead.
func (*RegexMatcher) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_v4alpha_regex_proto_rawDescGZIP(), []int{0}
}

func (m *RegexMatcher) GetEngineType() isRegexMatcher_EngineType {
	if m != nil {
		return m.EngineType
	}
	return nil
}

func (x *RegexMatcher) GetGoogleRe2() *RegexMatcher_GoogleRE2 {
	if x, ok := x.GetEngineType().(*RegexMatcher_GoogleRe2); ok {
		return x.GoogleRe2
	}
	return nil
}

func (x *RegexMatcher) GetRegex() string {
	if x != nil {
		return x.Regex
	}
	return ""
}

type isRegexMatcher_EngineType interface {
	isRegexMatcher_EngineType()
}

type RegexMatcher_GoogleRe2 struct {
	GoogleRe2 *RegexMatcher_GoogleRE2 `protobuf:"bytes,1,opt,name=google_re2,json=googleRe2,proto3,oneof"`
}

func (*RegexMatcher_GoogleRe2) isRegexMatcher_EngineType() {}

type RegexMatchAndSubstitute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pattern      *RegexMatcher `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Substitution string        `protobuf:"bytes,2,opt,name=substitution,proto3" json:"substitution,omitempty"`
}

func (x *RegexMatchAndSubstitute) Reset() {
	*x = RegexMatchAndSubstitute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_v4alpha_regex_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegexMatchAndSubstitute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegexMatchAndSubstitute) ProtoMessage() {}

func (x *RegexMatchAndSubstitute) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_v4alpha_regex_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegexMatchAndSubstitute.ProtoReflect.Descriptor instead.
func (*RegexMatchAndSubstitute) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_v4alpha_regex_proto_rawDescGZIP(), []int{1}
}

func (x *RegexMatchAndSubstitute) GetPattern() *RegexMatcher {
	if x != nil {
		return x.Pattern
	}
	return nil
}

func (x *RegexMatchAndSubstitute) GetSubstitution() string {
	if x != nil {
		return x.Substitution
	}
	return ""
}

type RegexMatcher_GoogleRE2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	HiddenEnvoyDeprecatedMaxProgramSize *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=hidden_envoy_deprecated_max_program_size,json=hiddenEnvoyDeprecatedMaxProgramSize,proto3" json:"hidden_envoy_deprecated_max_program_size,omitempty"`
}

func (x *RegexMatcher_GoogleRE2) Reset() {
	*x = RegexMatcher_GoogleRE2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_v4alpha_regex_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegexMatcher_GoogleRE2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegexMatcher_GoogleRE2) ProtoMessage() {}

func (x *RegexMatcher_GoogleRE2) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_v4alpha_regex_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegexMatcher_GoogleRE2.ProtoReflect.Descriptor instead.
func (*RegexMatcher_GoogleRE2) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_v4alpha_regex_proto_rawDescGZIP(), []int{0, 0}
}

// Deprecated: Do not use.
func (x *RegexMatcher_GoogleRE2) GetHiddenEnvoyDeprecatedMaxProgramSize() *wrappers.UInt32Value {
	if x != nil {
		return x.HiddenEnvoyDeprecatedMaxProgramSize
	}
	return nil
}

var File_envoy_type_matcher_v4alpha_regex_proto protoreflect.FileDescriptor

var file_envoy_type_matcher_v4alpha_regex_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x67,
	0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x87, 0x03, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x5d, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x32, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x45, 0x32, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x48, 0x00, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x32, 0x12,
	0x1d, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x1a, 0xb9,
	0x01, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x45, 0x32, 0x12, 0x77, 0x0a, 0x28,
	0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x5f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x23, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x44, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x78, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x53, 0x69, 0x7a, 0x65, 0x3a, 0x33, 0x9a, 0xc5, 0x88, 0x1e, 0x2e, 0x0a, 0x2c, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x45, 0x32, 0x3a, 0x29, 0x9a, 0xc5, 0x88, 0x1e,
	0x24, 0x0a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x12, 0x0a, 0x0b, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0xc1, 0x01, 0x0a, 0x17, 0x52, 0x65,
	0x67, 0x65, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x6e, 0x64, 0x53, 0x75, 0x62, 0x73, 0x74,
	0x69, 0x74, 0x75, 0x74, 0x65, 0x12, 0x4c, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x34, 0x9a, 0xc5, 0x88, 0x1e, 0x2f, 0x0a, 0x2d,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x41, 0x6e, 0x64, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x42, 0x40, 0x0a,
	0x28, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0a, 0x52, 0x65, 0x67, 0x65, 0x78,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x03, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_type_matcher_v4alpha_regex_proto_rawDescOnce sync.Once
	file_envoy_type_matcher_v4alpha_regex_proto_rawDescData = file_envoy_type_matcher_v4alpha_regex_proto_rawDesc
)

func file_envoy_type_matcher_v4alpha_regex_proto_rawDescGZIP() []byte {
	file_envoy_type_matcher_v4alpha_regex_proto_rawDescOnce.Do(func() {
		file_envoy_type_matcher_v4alpha_regex_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_matcher_v4alpha_regex_proto_rawDescData)
	})
	return file_envoy_type_matcher_v4alpha_regex_proto_rawDescData
}

var file_envoy_type_matcher_v4alpha_regex_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_type_matcher_v4alpha_regex_proto_goTypes = []interface{}{
	(*RegexMatcher)(nil),            // 0: envoy.type.matcher.v4alpha.RegexMatcher
	(*RegexMatchAndSubstitute)(nil), // 1: envoy.type.matcher.v4alpha.RegexMatchAndSubstitute
	(*RegexMatcher_GoogleRE2)(nil),  // 2: envoy.type.matcher.v4alpha.RegexMatcher.GoogleRE2
	(*wrappers.UInt32Value)(nil),    // 3: google.protobuf.UInt32Value
}
var file_envoy_type_matcher_v4alpha_regex_proto_depIdxs = []int32{
	2, // 0: envoy.type.matcher.v4alpha.RegexMatcher.google_re2:type_name -> envoy.type.matcher.v4alpha.RegexMatcher.GoogleRE2
	0, // 1: envoy.type.matcher.v4alpha.RegexMatchAndSubstitute.pattern:type_name -> envoy.type.matcher.v4alpha.RegexMatcher
	3, // 2: envoy.type.matcher.v4alpha.RegexMatcher.GoogleRE2.hidden_envoy_deprecated_max_program_size:type_name -> google.protobuf.UInt32Value
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_type_matcher_v4alpha_regex_proto_init() }
func file_envoy_type_matcher_v4alpha_regex_proto_init() {
	if File_envoy_type_matcher_v4alpha_regex_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_type_matcher_v4alpha_regex_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegexMatcher); i {
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
		file_envoy_type_matcher_v4alpha_regex_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegexMatchAndSubstitute); i {
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
		file_envoy_type_matcher_v4alpha_regex_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegexMatcher_GoogleRE2); i {
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
	file_envoy_type_matcher_v4alpha_regex_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RegexMatcher_GoogleRe2)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_type_matcher_v4alpha_regex_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_matcher_v4alpha_regex_proto_goTypes,
		DependencyIndexes: file_envoy_type_matcher_v4alpha_regex_proto_depIdxs,
		MessageInfos:      file_envoy_type_matcher_v4alpha_regex_proto_msgTypes,
	}.Build()
	File_envoy_type_matcher_v4alpha_regex_proto = out.File
	file_envoy_type_matcher_v4alpha_regex_proto_rawDesc = nil
	file_envoy_type_matcher_v4alpha_regex_proto_goTypes = nil
	file_envoy_type_matcher_v4alpha_regex_proto_depIdxs = nil
}
