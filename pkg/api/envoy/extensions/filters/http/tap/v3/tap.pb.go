// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/tap/v3/tap.proto

package envoy_extensions_filters_http_tap_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/extensions/common/tap/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Tap struct {
	CommonConfig         *v3.CommonExtensionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Tap) Reset()         { *m = Tap{} }
func (m *Tap) String() string { return proto.CompactTextString(m) }
func (*Tap) ProtoMessage()    {}
func (*Tap) Descriptor() ([]byte, []int) {
	return fileDescriptor_184267ac15318ad7, []int{0}
}

func (m *Tap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tap.Unmarshal(m, b)
}
func (m *Tap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tap.Marshal(b, m, deterministic)
}
func (m *Tap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tap.Merge(m, src)
}
func (m *Tap) XXX_Size() int {
	return xxx_messageInfo_Tap.Size(m)
}
func (m *Tap) XXX_DiscardUnknown() {
	xxx_messageInfo_Tap.DiscardUnknown(m)
}

var xxx_messageInfo_Tap proto.InternalMessageInfo

func (m *Tap) GetCommonConfig() *v3.CommonExtensionConfig {
	if m != nil {
		return m.CommonConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*Tap)(nil), "envoy.extensions.filters.http.tap.v3.Tap")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/tap/v3/tap.proto", fileDescriptor_184267ac15318ad7)
}

var fileDescriptor_184267ac15318ad7 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0x28, 0x29, 0x29, 0xd0, 0x2f, 0x49, 0x2c, 0xd0, 0x2f, 0x33,
	0x06, 0x51, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2a, 0x60, 0xf5, 0x7a, 0x08, 0xf5, 0x7a,
	0x50, 0xf5, 0x7a, 0x20, 0xf5, 0x7a, 0x20, 0x85, 0x65, 0xc6, 0x52, 0xda, 0x18, 0xa6, 0x26, 0xe7,
	0xe7, 0xe6, 0xe6, 0xe7, 0xc1, 0xcc, 0x83, 0xf0, 0x20, 0x46, 0x4a, 0xc9, 0x96, 0xa6, 0x14, 0x24,
	0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x80, 0x15, 0x17, 0x97, 0x24, 0x96, 0x94, 0x16,
	0x43, 0xa5, 0x15, 0x31, 0xa4, 0xcb, 0x52, 0x8b, 0x40, 0x86, 0x66, 0xe6, 0xa5, 0x43, 0x95, 0x88,
	0x97, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xa4, 0xea, 0xc3, 0x18, 0x10, 0x09, 0xa5, 0x39, 0x8c,
	0x5c, 0xcc, 0x21, 0x89, 0x05, 0x42, 0x29, 0x5c, 0xbc, 0x10, 0x2b, 0xe3, 0x93, 0xf3, 0xf3, 0xd2,
	0x32, 0xd3, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x4c, 0xf5, 0x30, 0x7c, 0x03, 0x75, 0x19,
	0xc4, 0x1f, 0x7a, 0xce, 0x60, 0x9e, 0x2b, 0x4c, 0xda, 0x19, 0xac, 0xd9, 0x89, 0xe3, 0x97, 0x13,
	0x6b, 0x17, 0x23, 0x93, 0x00, 0x63, 0x10, 0x0f, 0x44, 0x39, 0x44, 0xdc, 0x4a, 0x7f, 0xd6, 0xd1,
	0x0e, 0x39, 0x2d, 0x2e, 0x0d, 0x88, 0xa1, 0x10, 0x9b, 0xa0, 0xc1, 0x83, 0x14, 0x3a, 0x46, 0x89,
	0x39, 0x05, 0x19, 0x89, 0x7a, 0x21, 0x89, 0x05, 0x4e, 0x1e, 0xbb, 0x1a, 0x4e, 0x5c, 0x64, 0x63,
	0x12, 0x60, 0xe2, 0x32, 0xca, 0xcc, 0x87, 0xb8, 0xa5, 0xa0, 0x28, 0xbf, 0xa2, 0x52, 0x8f, 0x98,
	0x40, 0x76, 0xe2, 0x08, 0x49, 0x2c, 0x08, 0x00, 0x79, 0x33, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x5f,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xc2, 0xa9, 0x20, 0xcf, 0x01, 0x00, 0x00,
}
