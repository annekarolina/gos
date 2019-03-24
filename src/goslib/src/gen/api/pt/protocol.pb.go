// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

/*
Package pt is a generated protocol buffer package.

It is generated from these files:
	protocol.proto

It has these top-level messages:
	SessionAuthParams
	SessionAuthResponse
	Ok
	Fail
	EquipLoadParams
	EquipLoadResponse
	EquipUnLoadParams
	EquipUnLoadResponse
	LoginResponse
	RoomJoinParams
	RoomJoinResponse
	RoomJoinNotice
*/
package pt

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

type SessionAuthParams struct {
	AccountId string `protobuf:"bytes,1,opt,name=AccountId" json:"AccountId,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=Token" json:"Token,omitempty"`
}

func (m *SessionAuthParams) Reset()                    { *m = SessionAuthParams{} }
func (m *SessionAuthParams) String() string            { return proto.CompactTextString(m) }
func (*SessionAuthParams) ProtoMessage()               {}
func (*SessionAuthParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SessionAuthParams) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *SessionAuthParams) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SessionAuthResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *SessionAuthResponse) Reset()                    { *m = SessionAuthResponse{} }
func (m *SessionAuthResponse) String() string            { return proto.CompactTextString(m) }
func (*SessionAuthResponse) ProtoMessage()               {}
func (*SessionAuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SessionAuthResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type Ok struct {
	OK bool `protobuf:"varint,1,opt,name=OK" json:"OK,omitempty"`
}

func (m *Ok) Reset()                    { *m = Ok{} }
func (m *Ok) String() string            { return proto.CompactTextString(m) }
func (*Ok) ProtoMessage()               {}
func (*Ok) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ok) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

type Fail struct {
	Fail string `protobuf:"bytes,1,opt,name=Fail" json:"Fail,omitempty"`
}

func (m *Fail) Reset()                    { *m = Fail{} }
func (m *Fail) String() string            { return proto.CompactTextString(m) }
func (*Fail) ProtoMessage()               {}
func (*Fail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Fail) GetFail() string {
	if m != nil {
		return m.Fail
	}
	return ""
}

type EquipLoadParams struct {
	PlayerID string `protobuf:"bytes,1,opt,name=PlayerID" json:"PlayerID,omitempty"`
	EquipId  string `protobuf:"bytes,2,opt,name=EquipId" json:"EquipId,omitempty"`
	HeroId   string `protobuf:"bytes,3,opt,name=HeroId" json:"HeroId,omitempty"`
}

func (m *EquipLoadParams) Reset()                    { *m = EquipLoadParams{} }
func (m *EquipLoadParams) String() string            { return proto.CompactTextString(m) }
func (*EquipLoadParams) ProtoMessage()               {}
func (*EquipLoadParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EquipLoadParams) GetPlayerID() string {
	if m != nil {
		return m.PlayerID
	}
	return ""
}

func (m *EquipLoadParams) GetEquipId() string {
	if m != nil {
		return m.EquipId
	}
	return ""
}

func (m *EquipLoadParams) GetHeroId() string {
	if m != nil {
		return m.HeroId
	}
	return ""
}

type EquipLoadResponse struct {
	PlayerID string `protobuf:"bytes,1,opt,name=PlayerID" json:"PlayerID,omitempty"`
	EquipId  string `protobuf:"bytes,2,opt,name=EquipId" json:"EquipId,omitempty"`
	Level    uint32 `protobuf:"varint,3,opt,name=Level" json:"Level,omitempty"`
}

func (m *EquipLoadResponse) Reset()                    { *m = EquipLoadResponse{} }
func (m *EquipLoadResponse) String() string            { return proto.CompactTextString(m) }
func (*EquipLoadResponse) ProtoMessage()               {}
func (*EquipLoadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EquipLoadResponse) GetPlayerID() string {
	if m != nil {
		return m.PlayerID
	}
	return ""
}

func (m *EquipLoadResponse) GetEquipId() string {
	if m != nil {
		return m.EquipId
	}
	return ""
}

func (m *EquipLoadResponse) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type EquipUnLoadParams struct {
	PlayerID string `protobuf:"bytes,1,opt,name=PlayerID" json:"PlayerID,omitempty"`
	EquipId  string `protobuf:"bytes,2,opt,name=EquipId" json:"EquipId,omitempty"`
	HeroId   string `protobuf:"bytes,3,opt,name=HeroId" json:"HeroId,omitempty"`
}

func (m *EquipUnLoadParams) Reset()                    { *m = EquipUnLoadParams{} }
func (m *EquipUnLoadParams) String() string            { return proto.CompactTextString(m) }
func (*EquipUnLoadParams) ProtoMessage()               {}
func (*EquipUnLoadParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *EquipUnLoadParams) GetPlayerID() string {
	if m != nil {
		return m.PlayerID
	}
	return ""
}

func (m *EquipUnLoadParams) GetEquipId() string {
	if m != nil {
		return m.EquipId
	}
	return ""
}

func (m *EquipUnLoadParams) GetHeroId() string {
	if m != nil {
		return m.HeroId
	}
	return ""
}

type EquipUnLoadResponse struct {
	PlayerID string `protobuf:"bytes,1,opt,name=PlayerID" json:"PlayerID,omitempty"`
	EquipId  string `protobuf:"bytes,2,opt,name=EquipId" json:"EquipId,omitempty"`
	Level    uint32 `protobuf:"varint,3,opt,name=Level" json:"Level,omitempty"`
}

func (m *EquipUnLoadResponse) Reset()                    { *m = EquipUnLoadResponse{} }
func (m *EquipUnLoadResponse) String() string            { return proto.CompactTextString(m) }
func (*EquipUnLoadResponse) ProtoMessage()               {}
func (*EquipUnLoadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *EquipUnLoadResponse) GetPlayerID() string {
	if m != nil {
		return m.PlayerID
	}
	return ""
}

func (m *EquipUnLoadResponse) GetEquipId() string {
	if m != nil {
		return m.EquipId
	}
	return ""
}

func (m *EquipUnLoadResponse) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type LoginResponse struct {
	Uuid      string               `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
	Level     uint32               `protobuf:"varint,2,opt,name=Level" json:"Level,omitempty"`
	Exp       float32              `protobuf:"fixed32,3,opt,name=Exp" json:"Exp,omitempty"`
	Equips    []*EquipLoadResponse `protobuf:"bytes,4,rep,name=Equips" json:"Equips,omitempty"`
	HeadEquip *EquipLoadParams     `protobuf:"bytes,5,opt,name=HeadEquip" json:"HeadEquip,omitempty"`
	Friends   []string             `protobuf:"bytes,6,rep,name=Friends" json:"Friends,omitempty"`
	Ages      []int32              `protobuf:"varint,7,rep,packed,name=Ages" json:"Ages,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *LoginResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *LoginResponse) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *LoginResponse) GetExp() float32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *LoginResponse) GetEquips() []*EquipLoadResponse {
	if m != nil {
		return m.Equips
	}
	return nil
}

func (m *LoginResponse) GetHeadEquip() *EquipLoadParams {
	if m != nil {
		return m.HeadEquip
	}
	return nil
}

func (m *LoginResponse) GetFriends() []string {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *LoginResponse) GetAges() []int32 {
	if m != nil {
		return m.Ages
	}
	return nil
}

type RoomJoinParams struct {
	RoomId   string `protobuf:"bytes,1,opt,name=RoomId" json:"RoomId,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=PlayerId" json:"PlayerId,omitempty"`
}

func (m *RoomJoinParams) Reset()                    { *m = RoomJoinParams{} }
func (m *RoomJoinParams) String() string            { return proto.CompactTextString(m) }
func (*RoomJoinParams) ProtoMessage()               {}
func (*RoomJoinParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RoomJoinParams) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *RoomJoinParams) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

type RoomJoinResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *RoomJoinResponse) Reset()                    { *m = RoomJoinResponse{} }
func (m *RoomJoinResponse) String() string            { return proto.CompactTextString(m) }
func (*RoomJoinResponse) ProtoMessage()               {}
func (*RoomJoinResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RoomJoinResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RoomJoinNotice struct {
	RoomId      string `protobuf:"bytes,1,opt,name=RoomId" json:"RoomId,omitempty"`
	NewPlayerId string `protobuf:"bytes,2,opt,name=NewPlayerId" json:"NewPlayerId,omitempty"`
}

func (m *RoomJoinNotice) Reset()                    { *m = RoomJoinNotice{} }
func (m *RoomJoinNotice) String() string            { return proto.CompactTextString(m) }
func (*RoomJoinNotice) ProtoMessage()               {}
func (*RoomJoinNotice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *RoomJoinNotice) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *RoomJoinNotice) GetNewPlayerId() string {
	if m != nil {
		return m.NewPlayerId
	}
	return ""
}

func init() {
	proto.RegisterType((*SessionAuthParams)(nil), "pt.SessionAuthParams")
	proto.RegisterType((*SessionAuthResponse)(nil), "pt.SessionAuthResponse")
	proto.RegisterType((*Ok)(nil), "pt.Ok")
	proto.RegisterType((*Fail)(nil), "pt.Fail")
	proto.RegisterType((*EquipLoadParams)(nil), "pt.EquipLoadParams")
	proto.RegisterType((*EquipLoadResponse)(nil), "pt.EquipLoadResponse")
	proto.RegisterType((*EquipUnLoadParams)(nil), "pt.EquipUnLoadParams")
	proto.RegisterType((*EquipUnLoadResponse)(nil), "pt.EquipUnLoadResponse")
	proto.RegisterType((*LoginResponse)(nil), "pt.LoginResponse")
	proto.RegisterType((*RoomJoinParams)(nil), "pt.RoomJoinParams")
	proto.RegisterType((*RoomJoinResponse)(nil), "pt.RoomJoinResponse")
	proto.RegisterType((*RoomJoinNotice)(nil), "pt.RoomJoinNotice")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcd, 0x6e, 0xdb, 0x30,
	0x0c, 0xc7, 0x61, 0x3b, 0x76, 0x1b, 0x06, 0xcd, 0x5a, 0x25, 0x2b, 0x84, 0x62, 0x07, 0x43, 0x27,
	0x1f, 0xb6, 0x0c, 0xeb, 0x9e, 0x20, 0x40, 0xdb, 0xd5, 0x6d, 0xd0, 0x14, 0xea, 0x7a, 0x2e, 0x34,
	0x5b, 0xc8, 0x84, 0x38, 0x92, 0x67, 0xd9, 0xfb, 0x78, 0xd6, 0xbd, 0xcc, 0x60, 0x59, 0xfe, 0xd8,
	0x86, 0x01, 0x43, 0x81, 0x9c, 0xcc, 0xbf, 0x44, 0xf2, 0x47, 0x52, 0x26, 0x4c, 0xf3, 0x42, 0x95,
	0x2a, 0x51, 0xd9, 0xc2, 0x18, 0xc8, 0xcd, 0x4b, 0xf2, 0x01, 0x4e, 0x1e, 0xb8, 0xd6, 0x42, 0xc9,
	0x65, 0x55, 0x7e, 0xbe, 0x67, 0x05, 0xdb, 0x69, 0xf4, 0x0a, 0xc6, 0xcb, 0x24, 0x51, 0x95, 0x2c,
	0xe3, 0x14, 0x3b, 0xa1, 0x13, 0x8d, 0x69, 0x7f, 0x80, 0xe6, 0xe0, 0x7f, 0x54, 0x5b, 0x2e, 0xb1,
	0x6b, 0x6e, 0x1a, 0x41, 0xde, 0xc2, 0x6c, 0x90, 0x88, 0x72, 0x9d, 0x2b, 0xa9, 0x39, 0xc2, 0x70,
	0xf0, 0x50, 0x25, 0x09, 0xd7, 0xda, 0x24, 0x3a, 0xa4, 0xad, 0x24, 0x73, 0x70, 0xd7, 0x5b, 0x34,
	0x05, 0x77, 0x7d, 0x6b, 0xaf, 0xdc, 0xf5, 0x2d, 0x39, 0x83, 0xd1, 0x15, 0x13, 0x19, 0x42, 0xcd,
	0xd7, 0xd2, 0x8d, 0x4d, 0x9e, 0xe0, 0xc5, 0xe5, 0x97, 0x4a, 0xe4, 0x2b, 0xc5, 0x52, 0x5b, 0xe9,
	0x19, 0x1c, 0xde, 0x67, 0xec, 0x07, 0x2f, 0xe2, 0x0b, 0xeb, 0xda, 0xe9, 0x1a, 0x6d, 0xdc, 0xe3,
	0xd4, 0x56, 0xda, 0x4a, 0x74, 0x0a, 0xc1, 0x35, 0x2f, 0x54, 0x9c, 0x62, 0xcf, 0x5c, 0x58, 0x45,
	0x9e, 0xe0, 0xa4, 0x03, 0x74, 0x1d, 0x3c, 0x0f, 0x31, 0x07, 0x7f, 0xc5, 0xbf, 0xf2, 0xcc, 0x10,
	0x8e, 0x68, 0x23, 0x08, 0xb3, 0x80, 0x47, 0xb9, 0xb7, 0x1e, 0x18, 0xcc, 0x06, 0x88, 0xbd, 0x74,
	0xf1, 0xd3, 0x81, 0xa3, 0x95, 0xda, 0x08, 0xd9, 0x65, 0x47, 0x30, 0x7a, 0xac, 0x44, 0xfb, 0xaf,
	0x18, 0xbb, 0x8f, 0x75, 0x07, 0xb1, 0xe8, 0x18, 0xbc, 0xcb, 0xef, 0xb9, 0xc9, 0xe7, 0xd2, 0xda,
	0x44, 0x6f, 0x20, 0x30, 0x38, 0x8d, 0x47, 0xa1, 0x17, 0x4d, 0xce, 0x5f, 0x2e, 0xf2, 0x72, 0xf1,
	0xd7, 0x33, 0x50, 0xeb, 0x84, 0xde, 0xc1, 0xf8, 0x9a, 0xb3, 0xd4, 0x28, 0xec, 0x87, 0x4e, 0x34,
	0x39, 0x9f, 0xfd, 0x16, 0xd1, 0x4c, 0x95, 0xf6, 0x5e, 0x75, 0x7f, 0x57, 0x85, 0xe0, 0x32, 0xd5,
	0x38, 0x08, 0xbd, 0xba, 0x3f, 0x2b, 0xeb, 0xba, 0x97, 0x1b, 0xae, 0xf1, 0x41, 0xe8, 0x45, 0x3e,
	0x35, 0x36, 0xb9, 0x80, 0x29, 0x55, 0x6a, 0x77, 0xa3, 0x84, 0xb4, 0x0f, 0x74, 0x0a, 0x41, 0x7d,
	0xd2, 0xed, 0x82, 0x55, 0x83, 0x99, 0xb6, 0x83, 0xeb, 0x34, 0x79, 0x0d, 0xc7, 0x6d, 0x96, 0xff,
	0xd8, 0x85, 0x9b, 0x9e, 0x79, 0xa7, 0x4a, 0x91, 0xf0, 0x7f, 0x32, 0x43, 0x98, 0xdc, 0xf1, 0x6f,
	0x7f, 0x60, 0x87, 0x47, 0x9f, 0x02, 0xb3, 0xdc, 0xef, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x19,
	0x4c, 0x9d, 0x07, 0xee, 0x03, 0x00, 0x00,
}
