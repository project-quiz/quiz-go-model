// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Answer.proto

package question

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Answer struct {
	Guid                 string   `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Answer) Reset()         { *m = Answer{} }
func (m *Answer) String() string { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()    {}
func (*Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac4f4feadc7a0433, []int{0}
}

func (m *Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Answer.Unmarshal(m, b)
}
func (m *Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Answer.Marshal(b, m, deterministic)
}
func (m *Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Answer.Merge(m, src)
}
func (m *Answer) XXX_Size() int {
	return xxx_messageInfo_Answer.Size(m)
}
func (m *Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_Answer proto.InternalMessageInfo

func (m *Answer) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *Answer) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Answer) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func init() {
	proto.RegisterType((*Answer)(nil), "Data.Question.Answer")
}

func init() { proto.RegisterFile("Answer.proto", fileDescriptor_ac4f4feadc7a0433) }

var fileDescriptor_ac4f4feadc7a0433 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x71, 0xcc, 0x2b, 0x2e,
	0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x75, 0x49, 0x2c, 0x49, 0xd4, 0x0b,
	0x2c, 0x4d, 0x2d, 0x2e, 0xc9, 0xcc, 0xcf, 0x53, 0x72, 0xe3, 0x62, 0x83, 0x48, 0x0b, 0x09, 0x71,
	0xb1, 0xa4, 0x97, 0x66, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x20, 0xb1,
	0x92, 0xd4, 0x8a, 0x12, 0x09, 0x26, 0x88, 0x18, 0x88, 0x2d, 0x24, 0xc2, 0xc5, 0x9a, 0x99, 0x9b,
	0x98, 0x9e, 0x2a, 0xc1, 0x0c, 0x16, 0x84, 0x70, 0x92, 0xd8, 0xc0, 0xa6, 0x1b, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x99, 0x8b, 0x44, 0x06, 0x6d, 0x00, 0x00, 0x00,
}
