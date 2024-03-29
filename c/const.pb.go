// Code generated by protoc-gen-go. DO NOT EDIT.
// source: const.proto

/*
Package c is a generated protocol buffer package.

It is generated from these files:
	const.proto

It has these top-level messages:
*/
package c

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

type Gender int32

const (
	Gender_Null    Gender = 0
	Gender_Unknown Gender = 1
	Gender_Male    Gender = 2
	Gender_Famale  Gender = 3
)

var Gender_name = map[int32]string{
	0: "Null",
	1: "Unknown",
	2: "Male",
	3: "Famale",
}
var Gender_value = map[string]int32{
	"Null":    0,
	"Unknown": 1,
	"Male":    2,
	"Famale":  3,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}
func (Gender) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("c.Gender", Gender_name, Gender_value)
}

func init() { proto.RegisterFile("const.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xce, 0xcf, 0x2b,
	0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4c, 0xd6, 0x32, 0xe5, 0x62, 0x73, 0x4f,
	0xcd, 0x4b, 0x49, 0x2d, 0x12, 0xe2, 0xe0, 0x62, 0xf1, 0x2b, 0xcd, 0xc9, 0x11, 0x60, 0x10, 0xe2,
	0xe6, 0x62, 0x0f, 0xcd, 0xcb, 0xce, 0xcb, 0x2f, 0xcf, 0x13, 0x60, 0x04, 0x09, 0xfb, 0x26, 0xe6,
	0xa4, 0x0a, 0x30, 0x09, 0x71, 0x71, 0xb1, 0xb9, 0x25, 0xe6, 0x82, 0xd8, 0xcc, 0x49, 0x6c, 0x60,
	0x03, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x5a, 0xda, 0x84, 0x4f, 0x00, 0x00, 0x00,
}
