// Code generated by protoc-gen-go.
// source: SFPacket.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	SFPacket.proto

It has these top-level messages:
	SF_Battle
	SF_EnterBattle
	SF_ExitBattle
	SF_BattleQuery
*/
package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import protobuf1 "PB_PacketServerDefine.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type SF_Battle struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SF_Battle) Reset()         { *m = SF_Battle{} }
func (m *SF_Battle) String() string { return proto.CompactTextString(m) }
func (*SF_Battle) ProtoMessage()    {}

type SF_EnterBattle struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SF_EnterBattle) Reset()         { *m = SF_EnterBattle{} }
func (m *SF_EnterBattle) String() string { return proto.CompactTextString(m) }
func (*SF_EnterBattle) ProtoMessage()    {}

type SF_ExitBattle struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SF_ExitBattle) Reset()         { *m = SF_ExitBattle{} }
func (m *SF_ExitBattle) String() string { return proto.CompactTextString(m) }
func (*SF_ExitBattle) ProtoMessage()    {}

type SF_BattleQuery struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SF_BattleQuery) Reset()         { *m = SF_BattleQuery{} }
func (m *SF_BattleQuery) String() string { return proto.CompactTextString(m) }
func (*SF_BattleQuery) ProtoMessage()    {}

func init() {
}
