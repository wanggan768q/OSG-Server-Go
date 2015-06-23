// Code generated by protoc-gen-go.
// source: CLPacket.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	CLPacket.proto

It has these top-level messages:
	CL_CheckAccount
*/
package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import protobuf1 "PB_PacketDefine.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CL_CheckAccount struct {
	Uid              *string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Account          *string `protobuf:"bytes,2,req,name=account" json:"account,omitempty"`
	Password         *string `protobuf:"bytes,3,req,name=password" json:"password,omitempty"`
	Option           *string `protobuf:"bytes,4,opt,name=option" json:"option,omitempty"`
	Language         *uint32 `protobuf:"varint,5,opt,name=language" json:"language,omitempty"`
	Udid             *string `protobuf:"bytes,6,opt,name=udid" json:"udid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CL_CheckAccount) Reset()         { *m = CL_CheckAccount{} }
func (m *CL_CheckAccount) String() string { return proto.CompactTextString(m) }
func (*CL_CheckAccount) ProtoMessage()    {}

func (m *CL_CheckAccount) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

func (m *CL_CheckAccount) GetAccount() string {
	if m != nil && m.Account != nil {
		return *m.Account
	}
	return ""
}

func (m *CL_CheckAccount) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *CL_CheckAccount) GetOption() string {
	if m != nil && m.Option != nil {
		return *m.Option
	}
	return ""
}

func (m *CL_CheckAccount) GetLanguage() uint32 {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return 0
}

func (m *CL_CheckAccount) GetUdid() string {
	if m != nil && m.Udid != nil {
		return *m.Udid
	}
	return ""
}

func (m *CL_CheckAccount) SetUid(value string) {
	if m != nil {
		m.Uid = &value
	}
}

func (m *CL_CheckAccount) SetAccount(value string) {
	if m != nil {
		m.Account = &value
	}
}

func (m *CL_CheckAccount) SetPassword(value string) {
	if m != nil {
		m.Password = &value
	}
}

func (m *CL_CheckAccount) SetOption(value string) {
	if m != nil {
		m.Option = &value
	}
}

func (m *CL_CheckAccount) SetLanguage(value uint32) {
	if m != nil {
		m.Language = &value
	}
}

func (m *CL_CheckAccount) SetUdid(value string) {
	if m != nil {
		m.Udid = &value
	}
}

func init() {
}