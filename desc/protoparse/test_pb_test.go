// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package protoparse is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Simple
	Test
	Another
*/
package protoparse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Test_Nested__NestedNested_EEE int32

const (
	Test_Nested__NestedNested_OK Test_Nested__NestedNested_EEE = 0
	Test_Nested__NestedNested_V1 Test_Nested__NestedNested_EEE = 1
	Test_Nested__NestedNested_V2 Test_Nested__NestedNested_EEE = 2
	Test_Nested__NestedNested_V3 Test_Nested__NestedNested_EEE = 3
	Test_Nested__NestedNested_V4 Test_Nested__NestedNested_EEE = 4
	Test_Nested__NestedNested_V5 Test_Nested__NestedNested_EEE = 5
	Test_Nested__NestedNested_V6 Test_Nested__NestedNested_EEE = 6
)

var Test_Nested__NestedNested_EEE_name = map[int32]string{
	0: "OK",
	1: "V1",
	2: "V2",
	3: "V3",
	4: "V4",
	5: "V5",
	6: "V6",
}
var Test_Nested__NestedNested_EEE_value = map[string]int32{
	"OK": 0,
	"V1": 1,
	"V2": 2,
	"V3": 3,
	"V4": 4,
	"V5": 5,
	"V6": 6,
}

func (x Test_Nested__NestedNested_EEE) Enum() *Test_Nested__NestedNested_EEE {
	p := new(Test_Nested__NestedNested_EEE)
	*p = x
	return p
}
func (x Test_Nested__NestedNested_EEE) String() string {
	return proto.EnumName(Test_Nested__NestedNested_EEE_name, int32(x))
}
func (x *Test_Nested__NestedNested_EEE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Test_Nested__NestedNested_EEE_value, data, "Test_Nested__NestedNested_EEE")
	if err != nil {
		return err
	}
	*x = Test_Nested__NestedNested_EEE(value)
	return nil
}
func (Test_Nested__NestedNested_EEE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 1, 0, 0}
}

type Simple struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id               *uint64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Simple) Reset()                    { *m = Simple{} }
func (m *Simple) String() string            { return proto.CompactTextString(m) }
func (*Simple) ProtoMessage()               {}
func (*Simple) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Simple) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Simple) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type Test struct {
	Foo                          *string          `protobuf:"bytes,1,opt,name=foo,json=|foo|" json:"foo,omitempty"`
	Array                        []int32          `protobuf:"varint,2,rep,name=array" json:"array,omitempty"`
	S                            *Simple          `protobuf:"bytes,3,opt,name=s" json:"s,omitempty"`
	R                            []*Simple        `protobuf:"bytes,4,rep,name=r" json:"r,omitempty"`
	M                            map[string]int32 `protobuf:"bytes,5,rep,name=m" json:"m,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	B                            []byte           `protobuf:"bytes,6,opt,name=b,def=\\000\\001\\002\\003\\004\\005\\006\\007fubar!" json:"b,omitempty"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

var extRange_Test = []proto.ExtensionRange{
	{100, 200},
	{300, 350},
	{500, 550},
}

func (*Test) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Test
}

var Default_Test_B []byte = []byte("\\000\\001\\002\\003\\004\\005\\006\\007fubar!")

func (m *Test) GetFoo() string {
	if m != nil && m.Foo != nil {
		return *m.Foo
	}
	return ""
}

func (m *Test) GetArray() []int32 {
	if m != nil {
		return m.Array
	}
	return nil
}

func (m *Test) GetS() *Simple {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *Test) GetR() []*Simple {
	if m != nil {
		return m.R
	}
	return nil
}

func (m *Test) GetM() map[string]int32 {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *Test) GetB() []byte {
	if m != nil && m.B != nil {
		return m.B
	}
	return append([]byte(nil), Default_Test_B...)
}

type Test_Nested struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Test_Nested) Reset()                    { *m = Test_Nested{} }
func (m *Test_Nested) String() string            { return proto.CompactTextString(m) }
func (*Test_Nested) ProtoMessage()               {}
func (*Test_Nested) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

var E_Test_Nested_Fooblez = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         20003,
	Name:          "foo.bar.Test.Nested.fooblez",
	Tag:           "varint,20003,opt,name=fooblez",
	Filename:      "test.proto",
}

type Test_Nested__NestedNested struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Test_Nested__NestedNested) Reset()                    { *m = Test_Nested__NestedNested{} }
func (m *Test_Nested__NestedNested) String() string            { return proto.CompactTextString(m) }
func (*Test_Nested__NestedNested) ProtoMessage()               {}
func (*Test_Nested__NestedNested) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1, 0} }

var E_Test_Nested_XNestedNested_XGarblez = &proto.ExtensionDesc{
	ExtendedType:  (*Test)(nil),
	ExtensionType: (*string)(nil),
	Field:         100,
	Name:          "foo.bar.Test.Nested._NestedNested._garblez",
	Tag:           "bytes,100,opt,name=_garblez,json=Garblez",
	Filename:      "test.proto",
}

type Test_Nested__NestedNested_NestedNestedNested struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Test_Nested__NestedNested_NestedNestedNested) Reset() {
	*m = Test_Nested__NestedNested_NestedNestedNested{}
}
func (m *Test_Nested__NestedNested_NestedNestedNested) String() string {
	return proto.CompactTextString(m)
}
func (*Test_Nested__NestedNested_NestedNestedNested) ProtoMessage() {}
func (*Test_Nested__NestedNested_NestedNestedNested) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 1, 0, 0}
}

type Another struct {
	Test             *Test                          `protobuf:"bytes,1,opt,name=test" json:"test,omitempty"`
	Fff              *Test_Nested__NestedNested_EEE `protobuf:"varint,2,opt,name=fff,enum=foo.bar.Test_Nested__NestedNested_EEE,def=1" json:"fff,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *Another) Reset()                    { *m = Another{} }
func (m *Another) String() string            { return proto.CompactTextString(m) }
func (*Another) ProtoMessage()               {}
func (*Another) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

const Default_Another_Fff Test_Nested__NestedNested_EEE = Test_Nested__NestedNested_V1

func (m *Another) GetTest() *Test {
	if m != nil {
		return m.Test
	}
	return nil
}

func (m *Another) GetFff() Test_Nested__NestedNested_EEE {
	if m != nil && m.Fff != nil {
		return *m.Fff
	}
	return Default_Another_Fff
}

var E_Label = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.ExtensionRangeOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         20000,
	Name:          "foo.bar.label",
	Tag:           "bytes,20000,opt,name=label",
	Filename:      "test.proto",
}

var E_Rept = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: ([]*Test)(nil),
	Field:         20002,
	Name:          "foo.bar.rept",
	Tag:           "bytes,20002,rep,name=rept",
	Filename:      "test.proto",
}

var E_Eee = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*Test_Nested__NestedNested_EEE)(nil),
	Field:         20010,
	Name:          "foo.bar.eee",
	Tag:           "varint,20010,opt,name=eee,enum=foo.bar.Test_Nested__NestedNested_EEE",
	Filename:      "test.proto",
}

var E_A = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*Another)(nil),
	Field:         20020,
	Name:          "foo.bar.a",
	Tag:           "bytes,20020,opt,name=a",
	Filename:      "test.proto",
}

func init() {
	proto.RegisterType((*Simple)(nil), "foo.bar.Simple")
	proto.RegisterType((*Test)(nil), "foo.bar.Test")
	proto.RegisterType((*Test_Nested)(nil), "foo.bar.Test.Nested")
	proto.RegisterType((*Test_Nested__NestedNested)(nil), "foo.bar.Test.Nested._NestedNested")
	proto.RegisterType((*Test_Nested__NestedNested_NestedNestedNested)(nil), "foo.bar.Test.Nested._NestedNested.NestedNestedNested")
	proto.RegisterType((*Another)(nil), "foo.bar.Another")
	proto.RegisterEnum("foo.bar.Test_Nested__NestedNested_EEE", Test_Nested__NestedNested_EEE_name, Test_Nested__NestedNested_EEE_value)
	proto.RegisterExtension(E_Test_Nested_Fooblez)
	proto.RegisterExtension(E_Test_Nested_XNestedNested_XGarblez)
	proto.RegisterExtension(E_Label)
	proto.RegisterExtension(E_Rept)
	proto.RegisterExtension(E_Eee)
	proto.RegisterExtension(E_A)
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 784 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4f, 0x8f, 0xdb, 0x44,
	0x14, 0xef, 0xcc, 0xd8, 0x4e, 0xfc, 0xb6, 0x2d, 0xa3, 0x61, 0x91, 0x2c, 0x4b, 0xa8, 0x69, 0x90,
	0xa2, 0x28, 0x02, 0xef, 0x6c, 0xba, 0x2d, 0xc8, 0xfc, 0x91, 0x8a, 0x88, 0x38, 0xa0, 0xb6, 0x92,
	0x41, 0x3d, 0x40, 0xa4, 0x6a, 0xbc, 0x19, 0x67, 0x53, 0xec, 0x8c, 0x35, 0x76, 0x10, 0xbb, 0xf4,
	0x04, 0xdf, 0x20, 0xa7, 0x1e, 0xf6, 0x80, 0xc2, 0x01, 0xa9, 0x70, 0x44, 0xdc, 0xb9, 0x95, 0x03,
	0x12, 0xdf, 0x00, 0xc9, 0xf9, 0x0a, 0x70, 0x45, 0xc8, 0x63, 0xb7, 0xcb, 0x02, 0xd2, 0x72, 0xf1,
	0x9b, 0x79, 0xbf, 0xf7, 0xde, 0xcc, 0xfb, 0xbd, 0x9f, 0x07, 0xa0, 0x94, 0x45, 0x19, 0xe4, 0x5a,
	0x95, 0x8a, 0x75, 0x12, 0xa5, 0x82, 0x58, 0x68, 0xbf, 0x37, 0x57, 0x6a, 0x9e, 0xca, 0x3d, 0xe3,
	0x8e, 0x57, 0xc9, 0xde, 0x4c, 0x16, 0x87, 0x7a, 0x91, 0x97, 0x4a, 0x37, 0xa1, 0xfd, 0x57, 0xc1,
	0xf9, 0x70, 0x91, 0xe5, 0xa9, 0x64, 0x0c, 0xac, 0xa5, 0xc8, 0xa4, 0x87, 0x7a, 0x68, 0xe8, 0x46,
	0x66, 0xcd, 0xae, 0x02, 0x5e, 0xcc, 0x3c, 0xdc, 0x43, 0x43, 0x2b, 0xc2, 0x8b, 0x59, 0xff, 0x89,
	0x05, 0xd6, 0x47, 0xb2, 0x28, 0x19, 0x03, 0x92, 0x28, 0xd5, 0xc6, 0xda, 0x8f, 0x12, 0xa5, 0x1e,
	0xb1, 0x5d, 0xb0, 0x85, 0xd6, 0xe2, 0xd8, 0xc3, 0x3d, 0x32, 0xb4, 0xa3, 0x66, 0xc3, 0x5e, 0x06,
	0x54, 0x78, 0xa4, 0x87, 0x86, 0x3b, 0xe3, 0x17, 0x82, 0xf6, 0x5e, 0x41, 0x73, 0x64, 0x84, 0x8a,
	0x1a, 0xd6, 0x9e, 0xd5, 0x23, 0xff, 0x09, 0x6b, 0xd6, 0x07, 0x94, 0x79, 0xb6, 0x81, 0x77, 0x9f,
	0xc3, 0xf5, 0x0d, 0x82, 0x3b, 0x93, 0x65, 0xa9, 0x8f, 0x23, 0x94, 0xb1, 0x03, 0x40, 0xb1, 0xe7,
	0xf4, 0xd0, 0xf0, 0x72, 0x38, 0x98, 0x72, 0xce, 0xa7, 0x9c, 0xef, 0x4f, 0x39, 0x1f, 0x4f, 0x39,
	0xbf, 0x31, 0xe5, 0xfc, 0x60, 0xca, 0xf9, 0xcd, 0x29, 0xe7, 0xb7, 0xa6, 0x9c, 0xbf, 0x9e, 0xac,
	0x62, 0xa1, 0xaf, 0x47, 0x28, 0xf6, 0x0f, 0xc0, 0x69, 0x4a, 0x30, 0x0a, 0xe4, 0x53, 0x79, 0xdc,
	0xf6, 0x52, 0x2f, 0xeb, 0x4e, 0x3e, 0x13, 0xe9, 0x4a, 0x9a, 0xce, 0xed, 0xa8, 0xd9, 0x84, 0xf8,
	0x0d, 0xe4, 0xff, 0x89, 0xc0, 0xb9, 0x2b, 0x8b, 0x52, 0xce, 0xfc, 0x1f, 0x11, 0x5c, 0x79, 0xd0,
	0xac, 0x5b, 0xcf, 0x00, 0xd8, 0xdf, 0xf7, 0xcd, 0x37, 0xa4, 0xeb, 0xca, 0xbd, 0x0c, 0xe4, 0x48,
	0xa9, 0x8d, 0x63, 0x15, 0xb9, 0x52, 0xfd, 0xb7, 0x81, 0x4c, 0x26, 0x13, 0xe6, 0x00, 0xbe, 0xf7,
	0x01, 0xbd, 0x54, 0xdb, 0xfb, 0xfb, 0x14, 0x19, 0x3b, 0xa6, 0xd8, 0xd8, 0x1b, 0x94, 0x18, 0x7b,
	0x40, 0x2d, 0x63, 0x6f, 0x52, 0xdb, 0xd8, 0x5b, 0xd4, 0x19, 0x0f, 0xa1, 0xfb, 0x60, 0x2e, 0x74,
	0x9c, 0xca, 0x13, 0x76, 0xe5, 0x1c, 0x29, 0xde, 0xcc, 0xf4, 0xd0, 0x79, 0xbf, 0x41, 0xc3, 0xdd,
	0x75, 0xe5, 0xee, 0x00, 0x99, 0xd7, 0x47, 0x93, 0x58, 0xa9, 0xc7, 0x95, 0xfb, 0xc7, 0xdd, 0xf1,
	0x9b, 0x50, 0xeb, 0xc3, 0xa4, 0x5f, 0x0b, 0x1a, 0x81, 0x04, 0xcf, 0x04, 0x12, 0xdc, 0x91, 0x45,
	0x21, 0xe6, 0xf2, 0x5e, 0x5e, 0x2e, 0xd4, 0xb2, 0xf0, 0xbe, 0x39, 0x45, 0x86, 0x82, 0x67, 0x19,
	0x23, 0xbb, 0x3b, 0xa3, 0x3f, 0xa3, 0x11, 0xed, 0x7e, 0x8f, 0xe9, 0x6f, 0xd8, 0xef, 0x7e, 0x59,
	0xb9, 0xd6, 0x43, 0x71, 0x72, 0x32, 0xa2, 0xdd, 0xdf, 0x09, 0xfd, 0xd6, 0x3a, 0xf3, 0xf4, 0x1f,
	0x63, 0xe8, 0xdc, 0x5e, 0xaa, 0xf2, 0x48, 0x6a, 0x76, 0x1d, 0xac, 0x5a, 0x9f, 0x86, 0xe4, 0x9d,
	0xf1, 0xf9, 0x5b, 0x47, 0x06, 0x62, 0x6f, 0x01, 0x49, 0x92, 0xc4, 0x50, 0x7e, 0x75, 0x3c, 0x38,
	0x3f, 0xec, 0x86, 0xca, 0xe0, 0x1c, 0xdd, 0xc1, 0x64, 0x32, 0x09, 0xf1, 0xfd, 0xfd, 0xa8, 0x4e,
	0x0b, 0xbf, 0x43, 0xeb, 0xca, 0x7d, 0x05, 0x88, 0x88, 0x0f, 0x29, 0xa2, 0x98, 0x12, 0xbf, 0x63,
	0x24, 0x4a, 0xbf, 0xe8, 0x13, 0x40, 0x49, 0xfd, 0x29, 0xfa, 0x84, 0x3e, 0x25, 0xeb, 0xca, 0xbd,
	0x06, 0x64, 0x26, 0x13, 0x4a, 0x28, 0xa6, 0xc8, 0xef, 0x02, 0x89, 0x85, 0xa6, 0x3f, 0xe1, 0x3a,
	0x64, 0x6e, 0xe2, 0xd6, 0x95, 0x6b, 0x9b, 0x90, 0x5f, 0x2b, 0x17, 0x6d, 0xb6, 0x2e, 0xa6, 0x97,
	0x36, 0x5b, 0xf7, 0x25, 0x78, 0x71, 0xd4, 0xd6, 0x9c, 0x8d, 0xda, 0xac, 0xa7, 0x35, 0xdc, 0x01,
	0x1b, 0x48, 0x36, 0xc8, 0x36, 0x5b, 0x17, 0xa0, 0xeb, 0x3b, 0x60, 0x1d, 0xab, 0x54, 0x6d, 0xb6,
	0x6e, 0x17, 0x1c, 0xdf, 0xa2, 0xbf, 0x7c, 0xe5, 0x6c, 0xb6, 0xae, 0x05, 0x98, 0xa2, 0xd6, 0xe2,
	0xf0, 0x1d, 0xb0, 0x53, 0x11, 0xcb, 0x94, 0x0d, 0xfe, 0x35, 0x80, 0xc9, 0xe7, 0xa5, 0x5c, 0x16,
	0x0b, 0xb5, 0x8c, 0xc4, 0xf2, 0x6c, 0x0e, 0x5f, 0x9f, 0xb6, 0xbf, 0x9a, 0x49, 0x0b, 0xdf, 0x03,
	0x4b, 0xcb, 0xbc, 0xbc, 0x78, 0x7e, 0x9b, 0x53, 0x64, 0x7e, 0x9e, 0x7f, 0x32, 0x5e, 0x67, 0x87,
	0x9f, 0x00, 0x91, 0x52, 0x5e, 0x5c, 0xe4, 0x89, 0x39, 0xfc, 0x7f, 0x0f, 0x25, 0xaa, 0xab, 0x86,
	0xb7, 0x01, 0x89, 0x8b, 0x4b, 0xff, 0x70, 0xda, 0x28, 0x82, 0x3e, 0x2f, 0xdd, 0x2a, 0x26, 0x42,
	0xe2, 0xdd, 0xbd, 0x8f, 0x5f, 0x9b, 0x2f, 0xca, 0xa3, 0x55, 0x1c, 0x1c, 0xaa, 0x6c, 0xef, 0xe1,
	0xd1, 0x2a, 0xcb, 0x9b, 0x97, 0x4c, 0xcb, 0x24, 0x95, 0x87, 0xa5, 0x79, 0xcd, 0x1a, 0x4f, 0x2e,
	0x74, 0x21, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x9b, 0x29, 0x65, 0x04, 0x05, 0x00, 0x00,
}
