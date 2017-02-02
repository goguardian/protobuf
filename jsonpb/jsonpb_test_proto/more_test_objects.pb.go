// Code generated by protoc-gen-go.
// source: more_test_objects.proto
// DO NOT EDIT!

/*
Package jsonpb is a generated protocol buffer package.

It is generated from these files:
	more_test_objects.proto
	test_objects.proto

It has these top-level messages:
	Simple3
	Mappy
	Slicy
	Simple
	Repeats
	Widget
	Maps
	MsgWithOneof
	Real
	Complex
	KnownTypes
*/
package jsonpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Numeral int32

const (
	Numeral_UNKNOWN Numeral = 0
	Numeral_ARABIC  Numeral = 1
	Numeral_ROMAN   Numeral = 2
)

var Numeral_name = map[int32]string{
	0: "UNKNOWN",
	1: "ARABIC",
	2: "ROMAN",
}
var Numeral_value = map[string]int32{
	"UNKNOWN": 0,
	"ARABIC":  1,
	"ROMAN":   2,
}

func (x Numeral) String() string {
	return proto.EnumName(Numeral_name, int32(x))
}
func (Numeral) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Simple3 struct {
	Dub float64 `protobuf:"fixed64,1,opt,name=dub" json:"dub,omitempty"`
}

func (m *Simple3) Reset()                    { *m = Simple3{} }
func (m *Simple3) String() string            { return proto.CompactTextString(m) }
func (*Simple3) ProtoMessage()               {}
func (*Simple3) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Simple3) GetDub() float64 {
	if m != nil {
		return m.Dub
	}
	return 0
}

type Mappy struct {
	Nummy    map[int64]int32    `protobuf:"bytes,1,rep,name=nummy" json:"nummy,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Strry    map[string]string  `protobuf:"bytes,2,rep,name=strry" json:"strry,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Objjy    map[int32]*Simple3 `protobuf:"bytes,3,rep,name=objjy" json:"objjy,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Buggy    map[int64]string   `protobuf:"bytes,4,rep,name=buggy" json:"buggy,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Booly    map[bool]bool      `protobuf:"bytes,5,rep,name=booly" json:"booly,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Enumy    map[string]Numeral `protobuf:"bytes,6,rep,name=enumy" json:"enumy,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=jsonpb.Numeral"`
	S32Booly map[int32]bool     `protobuf:"bytes,7,rep,name=s32booly" json:"s32booly,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	S64Booly map[int64]bool     `protobuf:"bytes,8,rep,name=s64booly" json:"s64booly,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	U32Booly map[uint32]bool    `protobuf:"bytes,9,rep,name=u32booly" json:"u32booly,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	U64Booly map[uint64]bool    `protobuf:"bytes,10,rep,name=u64booly" json:"u64booly,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *Mappy) Reset()                    { *m = Mappy{} }
func (m *Mappy) String() string            { return proto.CompactTextString(m) }
func (*Mappy) ProtoMessage()               {}
func (*Mappy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Mappy) GetNummy() map[int64]int32 {
	if m != nil {
		return m.Nummy
	}
	return nil
}

func (m *Mappy) GetStrry() map[string]string {
	if m != nil {
		return m.Strry
	}
	return nil
}

func (m *Mappy) GetObjjy() map[int32]*Simple3 {
	if m != nil {
		return m.Objjy
	}
	return nil
}

func (m *Mappy) GetBuggy() map[int64]string {
	if m != nil {
		return m.Buggy
	}
	return nil
}

func (m *Mappy) GetBooly() map[bool]bool {
	if m != nil {
		return m.Booly
	}
	return nil
}

func (m *Mappy) GetEnumy() map[string]Numeral {
	if m != nil {
		return m.Enumy
	}
	return nil
}

func (m *Mappy) GetS32Booly() map[int32]bool {
	if m != nil {
		return m.S32Booly
	}
	return nil
}

func (m *Mappy) GetS64Booly() map[int64]bool {
	if m != nil {
		return m.S64Booly
	}
	return nil
}

func (m *Mappy) GetU32Booly() map[uint32]bool {
	if m != nil {
		return m.U32Booly
	}
	return nil
}

func (m *Mappy) GetU64Booly() map[uint64]bool {
	if m != nil {
		return m.U64Booly
	}
	return nil
}

type Slicy struct {
	Slicy []string `protobuf:"bytes,1,rep,name=slicy" json:"slicy,omitempty"`
}

func (m *Slicy) Reset()                    { *m = Slicy{} }
func (m *Slicy) String() string            { return proto.CompactTextString(m) }
func (*Slicy) ProtoMessage()               {}
func (*Slicy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Slicy) GetSlicy() []string {
	if m != nil {
		return m.Slicy
	}
	return nil
}

func init() {
	proto.RegisterType((*Simple3)(nil), "jsonpb.Simple3")
	proto.RegisterType((*Mappy)(nil), "jsonpb.Mappy")
	proto.RegisterType((*Slicy)(nil), "jsonpb.Slicy")
	proto.RegisterEnum("jsonpb.Numeral", Numeral_name, Numeral_value)
}

func init() { proto.RegisterFile("more_test_objects.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0xc1, 0x6b, 0xdb, 0x30,
	0x14, 0x87, 0xe7, 0xa4, 0x4e, 0xec, 0x17, 0xba, 0x05, 0x31, 0x98, 0x58, 0x19, 0x84, 0xc0, 0x20,
	0x0c, 0xe6, 0x43, 0x32, 0xb6, 0xb2, 0x9d, 0xd2, 0xd1, 0x43, 0x19, 0x75, 0xc0, 0x21, 0xec, 0x58,
	0xe2, 0x4e, 0x94, 0x66, 0x76, 0x64, 0x6c, 0x6b, 0xa0, 0x3f, 0x7e, 0x30, 0x9e, 0x24, 0xc7, 0x8a,
	0x51, 0xc8, 0x6e, 0x12, 0xbf, 0xef, 0xf3, 0x7b, 0x92, 0x1e, 0x86, 0x37, 0x39, 0x2f, 0xd9, 0x43,
	0xcd, 0xaa, 0xfa, 0x81, 0xa7, 0x3b, 0xf6, 0x58, 0x57, 0x51, 0x51, 0xf2, 0x9a, 0x93, 0xc1, 0xae,
	0xe2, 0xfb, 0x22, 0x9d, 0x5e, 0xc1, 0x70, 0xfd, 0x9c, 0x17, 0x19, 0x5b, 0x90, 0x31, 0xf4, 0x7f,
	0x89, 0x94, 0x7a, 0x13, 0x6f, 0xe6, 0x25, 0xb8, 0x9c, 0xfe, 0x0d, 0xc0, 0xbf, 0xdf, 0x16, 0x85,
	0x24, 0x11, 0xf8, 0x7b, 0x91, 0xe7, 0x92, 0x7a, 0x93, 0xfe, 0x6c, 0x34, 0xa7, 0x91, 0xd6, 0x23,
	0x95, 0x46, 0x31, 0x46, 0xb7, 0xfb, 0xba, 0x94, 0x89, 0xc6, 0x90, 0xaf, 0xea, 0xb2, 0x94, 0xb4,
	0xe7, 0xe2, 0xd7, 0x18, 0x19, 0x5e, 0x61, 0xc8, 0xf3, 0x74, 0xb7, 0x93, 0xb4, 0xef, 0xe2, 0x57,
	0x18, 0x19, 0x5e, 0x61, 0xc8, 0xa7, 0xe2, 0xe9, 0x49, 0xd2, 0x0b, 0x17, 0x7f, 0x83, 0x91, 0xe1,
	0x15, 0xa6, 0x78, 0xce, 0x33, 0x49, 0x7d, 0x27, 0x8f, 0x51, 0xc3, 0xe3, 0x1a, 0x79, 0xb6, 0x17,
	0xb9, 0xa4, 0x03, 0x17, 0x7f, 0x8b, 0x91, 0xe1, 0x15, 0x46, 0xbe, 0x40, 0x50, 0x2d, 0xe6, 0xba,
	0xc4, 0x50, 0x29, 0x57, 0x9d, 0x23, 0x9b, 0x54, 0x5b, 0x07, 0x58, 0x89, 0x9f, 0x3f, 0x69, 0x31,
	0x70, 0x8a, 0x26, 0x6d, 0x44, 0xb3, 0x45, 0x51, 0x34, 0x15, 0x43, 0x97, 0xb8, 0x39, 0xae, 0x28,
	0xac, 0x8a, 0xa2, 0xa9, 0x08, 0x4e, 0xf1, 0xb8, 0x62, 0x03, 0xbf, 0xbd, 0x06, 0x68, 0x1f, 0x1a,
	0xa7, 0xe5, 0x37, 0x93, 0x6a, 0x5a, 0xfa, 0x09, 0x2e, 0xc9, 0x6b, 0xf0, 0xff, 0x6c, 0x33, 0xc1,
	0x68, 0x6f, 0xe2, 0xcd, 0xfc, 0x44, 0x6f, 0xbe, 0xf6, 0xae, 0x3d, 0x34, 0xdb, 0x27, 0xb7, 0xcd,
	0xd0, 0x61, 0x86, 0xb6, 0x79, 0x07, 0xd0, 0x3e, 0xbe, 0x6d, 0xfa, 0xda, 0x7c, 0x6f, 0x9b, 0xa3,
	0xf9, 0xab, 0xe6, 0x24, 0x66, 0xa6, 0x3b, 0x4d, 0xb4, 0x73, 0x71, 0xae, 0xfd, 0xb0, 0x6b, 0x1e,
	0x2e, 0xc4, 0x36, 0x03, 0x87, 0x19, 0x74, 0xda, 0x6f, 0x67, 0xc5, 0x71, 0xf0, 0xa3, 0xf6, 0x5f,
	0xb6, 0xed, 0xc7, 0x22, 0x67, 0xe5, 0x36, 0xb3, 0x3f, 0xf5, 0x0d, 0x2e, 0x8f, 0x66, 0xc8, 0x71,
	0x19, 0xa7, 0xfb, 0x40, 0xd9, 0x7e, 0xd5, 0x73, 0xc7, 0xef, 0xca, 0x9b, 0x53, 0x95, 0x2f, 0xff,
	0x47, 0x3e, 0x55, 0xf9, 0xe2, 0x8c, 0x3c, 0x7d, 0x07, 0xfe, 0x3a, 0x7b, 0x7e, 0x54, 0x48, 0x85,
	0x0b, 0xf5, 0xfb, 0x09, 0x13, 0xbd, 0xf9, 0xf0, 0x11, 0x86, 0xe6, 0xa2, 0xc8, 0x08, 0x86, 0x9b,
	0xf8, 0x47, 0xbc, 0xfa, 0x19, 0x8f, 0x5f, 0x10, 0x80, 0xc1, 0x32, 0x59, 0xde, 0xdc, 0x7d, 0x1f,
	0x7b, 0x24, 0x04, 0x3f, 0x59, 0xdd, 0x2f, 0xe3, 0x71, 0x2f, 0x1d, 0xa8, 0x3f, 0xdf, 0xe2, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x07, 0x03, 0x86, 0x14, 0x05, 0x00, 0x00,
}
