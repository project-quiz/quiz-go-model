// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package model

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

// Messages client -> server
type JoinGame struct {
	Player               *Player  `protobuf:"bytes,1,opt,name=Player,proto3" json:"Player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinGame) Reset()         { *m = JoinGame{} }
func (m *JoinGame) String() string { return proto.CompactTextString(m) }
func (*JoinGame) ProtoMessage()    {}
func (*JoinGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{0}
}

func (m *JoinGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinGame.Unmarshal(m, b)
}
func (m *JoinGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinGame.Marshal(b, m, deterministic)
}
func (m *JoinGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinGame.Merge(m, src)
}
func (m *JoinGame) XXX_Size() int {
	return xxx_messageInfo_JoinGame.Size(m)
}
func (m *JoinGame) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinGame.DiscardUnknown(m)
}

var xxx_messageInfo_JoinGame proto.InternalMessageInfo

func (m *JoinGame) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

// Messages server -> client
type GameJoined struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameJoined) Reset()         { *m = GameJoined{} }
func (m *GameJoined) String() string { return proto.CompactTextString(m) }
func (*GameJoined) ProtoMessage()    {}
func (*GameJoined) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{1}
}

func (m *GameJoined) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameJoined.Unmarshal(m, b)
}
func (m *GameJoined) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameJoined.Marshal(b, m, deterministic)
}
func (m *GameJoined) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameJoined.Merge(m, src)
}
func (m *GameJoined) XXX_Size() int {
	return xxx_messageInfo_GameJoined.Size(m)
}
func (m *GameJoined) XXX_DiscardUnknown() {
	xxx_messageInfo_GameJoined.DiscardUnknown(m)
}

var xxx_messageInfo_GameJoined proto.InternalMessageInfo

type GameStart struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameStart) Reset()         { *m = GameStart{} }
func (m *GameStart) String() string { return proto.CompactTextString(m) }
func (*GameStart) ProtoMessage()    {}
func (*GameStart) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{2}
}

func (m *GameStart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameStart.Unmarshal(m, b)
}
func (m *GameStart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameStart.Marshal(b, m, deterministic)
}
func (m *GameStart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameStart.Merge(m, src)
}
func (m *GameStart) XXX_Size() int {
	return xxx_messageInfo_GameStart.Size(m)
}
func (m *GameStart) XXX_DiscardUnknown() {
	xxx_messageInfo_GameStart.DiscardUnknown(m)
}

var xxx_messageInfo_GameStart proto.InternalMessageInfo

type GameStop struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameStop) Reset()         { *m = GameStop{} }
func (m *GameStop) String() string { return proto.CompactTextString(m) }
func (*GameStop) ProtoMessage()    {}
func (*GameStop) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{3}
}

func (m *GameStop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameStop.Unmarshal(m, b)
}
func (m *GameStop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameStop.Marshal(b, m, deterministic)
}
func (m *GameStop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameStop.Merge(m, src)
}
func (m *GameStop) XXX_Size() int {
	return xxx_messageInfo_GameStop.Size(m)
}
func (m *GameStop) XXX_DiscardUnknown() {
	xxx_messageInfo_GameStop.DiscardUnknown(m)
}

var xxx_messageInfo_GameStop proto.InternalMessageInfo

type GameEnd struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameEnd) Reset()         { *m = GameEnd{} }
func (m *GameEnd) String() string { return proto.CompactTextString(m) }
func (*GameEnd) ProtoMessage()    {}
func (*GameEnd) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{4}
}

func (m *GameEnd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameEnd.Unmarshal(m, b)
}
func (m *GameEnd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameEnd.Marshal(b, m, deterministic)
}
func (m *GameEnd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameEnd.Merge(m, src)
}
func (m *GameEnd) XXX_Size() int {
	return xxx_messageInfo_GameEnd.Size(m)
}
func (m *GameEnd) XXX_DiscardUnknown() {
	xxx_messageInfo_GameEnd.DiscardUnknown(m)
}

var xxx_messageInfo_GameEnd proto.InternalMessageInfo

type GameResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameResult) Reset()         { *m = GameResult{} }
func (m *GameResult) String() string { return proto.CompactTextString(m) }
func (*GameResult) ProtoMessage()    {}
func (*GameResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_38fc58335341d769, []int{5}
}

func (m *GameResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameResult.Unmarshal(m, b)
}
func (m *GameResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameResult.Marshal(b, m, deterministic)
}
func (m *GameResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameResult.Merge(m, src)
}
func (m *GameResult) XXX_Size() int {
	return xxx_messageInfo_GameResult.Size(m)
}
func (m *GameResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GameResult.DiscardUnknown(m)
}

var xxx_messageInfo_GameResult proto.InternalMessageInfo

func init() {
	proto.RegisterType((*JoinGame)(nil), "model.JoinGame")
	proto.RegisterType((*GameJoined)(nil), "model.GameJoined")
	proto.RegisterType((*GameStart)(nil), "model.GameStart")
	proto.RegisterType((*GameStop)(nil), "model.GameStop")
	proto.RegisterType((*GameEnd)(nil), "model.GameEnd")
	proto.RegisterType((*GameResult)(nil), "model.GameResult")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor_38fc58335341d769) }

var fileDescriptor_38fc58335341d769 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4f, 0xcc, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x91, 0xe2, 0x29,
	0xc8, 0x49, 0xac, 0x4c, 0x2d, 0x82, 0x08, 0x2a, 0x19, 0x72, 0x71, 0x78, 0xe5, 0x67, 0xe6, 0xb9,
	0x27, 0xe6, 0xa6, 0x0a, 0xa9, 0x72, 0xb1, 0x05, 0x80, 0xe5, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8,
	0x8d, 0x78, 0xf5, 0xc0, 0x3a, 0xf4, 0x20, 0x82, 0x41, 0x50, 0x49, 0x25, 0x1e, 0x2e, 0x2e, 0x90,
	0x72, 0x90, 0xb6, 0xd4, 0x14, 0x25, 0x6e, 0x2e, 0x4e, 0x10, 0x2f, 0xb8, 0x24, 0xb1, 0xa8, 0x44,
	0x89, 0x8b, 0x8b, 0x03, 0xc2, 0xc9, 0x2f, 0x50, 0xe2, 0xe4, 0x62, 0x07, 0xb1, 0x5d, 0xf3, 0x52,
	0x60, 0x3a, 0x82, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x92, 0xd8, 0xc0, 0x36, 0x1b, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x99, 0x8d, 0x56, 0x38, 0x9c, 0x00, 0x00, 0x00,
}
