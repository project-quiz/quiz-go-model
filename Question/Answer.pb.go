// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Answer.proto

package Data_Question

import (
	fmt "fmt"
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

type Answer struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
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

func (m *Answer) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
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
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x71, 0xcc, 0x2b, 0x2e,
	0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x75, 0x49, 0x2c, 0x49, 0xd4, 0x0b,
	0x2c, 0x4d, 0x2d, 0x2e, 0xc9, 0xcc, 0xcf, 0x53, 0x72, 0xe2, 0x62, 0x83, 0x48, 0x0b, 0xf1, 0x71,
	0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71,
	0xb1, 0x94, 0xa4, 0x56, 0x94, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x22,
	0x5c, 0xac, 0x99, 0xb9, 0x89, 0xe9, 0xa9, 0x12, 0xcc, 0x60, 0x41, 0x08, 0x27, 0x89, 0x0d, 0x6c,
	0xb2, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xef, 0xc0, 0x0d, 0x69, 0x00, 0x00, 0x00,
}
