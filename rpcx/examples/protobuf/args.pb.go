// Code generated by protoc-gen-go. DO NOT EDIT.
// source: args.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	args.proto
	reply.proto

It has these top-level messages:
	Args
	Reply
*/
package main

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

type Args struct {
	A int32 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *Args) Reset()                    { *m = Args{} }
func (m *Args) String() string            { return proto.CompactTextString(m) }
func (*Args) ProtoMessage()               {}
func (*Args) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Args) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Args) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

func init() {
	proto.RegisterType((*Args)(nil), "main.Args")
}

func init() { proto.RegisterFile("args.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 77 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x4a, 0x2f,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x52, 0xe2, 0x62,
	0x71, 0x2c, 0x4a, 0x2f, 0x16, 0xe2, 0xe1, 0x62, 0x4c, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d,
	0x62, 0x4c, 0x04, 0xf1, 0x92, 0x24, 0x98, 0x20, 0xbc, 0xa4, 0x24, 0x36, 0xb0, 0x06, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0xe8, 0xb5, 0x2b, 0x3e, 0x00, 0x00, 0x00,
}
