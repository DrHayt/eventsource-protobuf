// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	events.proto

It has these top-level messages:
	A
	B
	EventContainer
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

type A struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	At      int64  `protobuf:"varint,3,opt,name=at" json:"at,omitempty"`
}

func (m *A) Reset()                    { *m = A{} }
func (m *A) String() string            { return proto.CompactTextString(m) }
func (*A) ProtoMessage()               {}
func (*A) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *A) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *A) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *A) GetAt() int64 {
	if m != nil {
		return m.At
	}
	return 0
}

type B struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	At      int64  `protobuf:"varint,3,opt,name=at" json:"at,omitempty"`
}

func (m *B) Reset()                    { *m = B{} }
func (m *B) String() string            { return proto.CompactTextString(m) }
func (*B) ProtoMessage()               {}
func (*B) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *B) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *B) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *B) GetAt() int64 {
	if m != nil {
		return m.At
	}
	return 0
}

type EventContainer struct {
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Ma   *A    `protobuf:"bytes,2,opt,name=ma" json:"ma,omitempty"`
	Mb   *B    `protobuf:"bytes,3,opt,name=mb" json:"mb,omitempty"`
}

func (m *EventContainer) Reset()                    { *m = EventContainer{} }
func (m *EventContainer) String() string            { return proto.CompactTextString(m) }
func (*EventContainer) ProtoMessage()               {}
func (*EventContainer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EventContainer) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *EventContainer) GetMa() *A {
	if m != nil {
		return m.Ma
	}
	return nil
}

func (m *EventContainer) GetMb() *B {
	if m != nil {
		return m.Mb
	}
	return nil
}

func init() {
	proto.RegisterType((*A)(nil), "main.a")
	proto.RegisterType((*B)(nil), "main.b")
	proto.RegisterType((*EventContainer)(nil), "main.event_container")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x8f, 0xb1, 0xae, 0xc2, 0x30,
	0x0c, 0x45, 0x15, 0xb7, 0x7d, 0xd5, 0x33, 0x08, 0xa4, 0x2c, 0x64, 0xac, 0x98, 0x3a, 0x75, 0x80,
	0x99, 0x9f, 0xf0, 0xc2, 0x88, 0x1c, 0x9a, 0x21, 0x43, 0x92, 0xaa, 0xb5, 0x2a, 0xf1, 0xf7, 0xa8,
	0x96, 0xe0, 0x07, 0xd8, 0x6c, 0x5d, 0xdf, 0xa3, 0x63, 0xdc, 0x87, 0x35, 0x64, 0x59, 0x86, 0x69,
	0x2e, 0x52, 0x6c, 0x9d, 0x38, 0xe6, 0xf3, 0x0d, 0x0d, 0xdb, 0x03, 0x42, 0x1c, 0x9d, 0xe9, 0x4c,
	0xff, 0x4f, 0x10, 0x47, 0xeb, 0xb0, 0x5d, 0xc3, 0xbc, 0xc4, 0x92, 0x1d, 0x74, 0xa6, 0x6f, 0xe8,
	0xb3, 0x6e, 0x97, 0x2c, 0xae, 0xea, 0x4c, 0x5f, 0x11, 0xb0, 0x6c, 0x75, 0xff, 0x43, 0xfd, 0x8e,
	0x47, 0x75, 0x7a, 0x3c, 0x4b, 0x16, 0x8e, 0x39, 0xcc, 0xd6, 0x62, 0x2d, 0xaf, 0x29, 0x28, 0xae,
	0x21, 0x9d, 0xed, 0x09, 0x21, 0xb1, 0xb2, 0x76, 0x97, 0x76, 0xd8, 0xbc, 0x07, 0x26, 0x48, 0xac,
	0x81, 0x57, 0xde, 0x37, 0xf0, 0x04, 0xc9, 0xfb, 0x3f, 0xfd, 0xf1, 0xfa, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0xc1, 0xa2, 0x31, 0xf3, 0x00, 0x00, 0x00,
}
