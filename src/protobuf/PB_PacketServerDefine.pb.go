// Code generated by protoc-gen-go.
// source: PB_PacketServerDefine.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	PB_PacketServerDefine.proto

It has these top-level messages:
*/
package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type LA_Protocol int32

const (
	LA_Protocol_eLA_PacketBegin LA_Protocol = 11000
	// ----------------------------
	LA_Protocol_eLA_Connected    LA_Protocol = 11000
	LA_Protocol_eLA_Disconnected LA_Protocol = 11001
	LA_Protocol_eLA_CheckAccount LA_Protocol = 11002
	// ----------------------------
	LA_Protocol_eLA_PacketEnd LA_Protocol = 12000
)

var LA_Protocol_name = map[int32]string{
	11000: "eLA_PacketBegin",
	// Duplicate value: 11000: "eLA_Connected",
	11001: "eLA_Disconnected",
	11002: "eLA_CheckAccount",
	12000: "eLA_PacketEnd",
}
var LA_Protocol_value = map[string]int32{
	"eLA_PacketBegin":  11000,
	"eLA_Connected":    11000,
	"eLA_Disconnected": 11001,
	"eLA_CheckAccount": 11002,
	"eLA_PacketEnd":    12000,
}

func (x LA_Protocol) Enum() *LA_Protocol {
	p := new(LA_Protocol)
	*p = x
	return p
}
func (x LA_Protocol) String() string {
	return proto.EnumName(LA_Protocol_name, int32(x))
}
func (x LA_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *LA_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LA_Protocol_value, data, "LA_Protocol")
	if err != nil {
		return err
	}
	*x = LA_Protocol(value)
	return nil
}

type AL_Protocol int32

const (
	AL_Protocol_eAL_PacketBegin AL_Protocol = 12000
	// ----------------------------
	AL_Protocol_eAL_Connected          AL_Protocol = 12000
	AL_Protocol_eAL_Disconnected       AL_Protocol = 12001
	AL_Protocol_eAL_CheckAccountResult AL_Protocol = 12002
	// ----------------------------
	AL_Protocol_eAL_PacketEnd AL_Protocol = 13000
)

var AL_Protocol_name = map[int32]string{
	12000: "eAL_PacketBegin",
	// Duplicate value: 12000: "eAL_Connected",
	12001: "eAL_Disconnected",
	12002: "eAL_CheckAccountResult",
	13000: "eAL_PacketEnd",
}
var AL_Protocol_value = map[string]int32{
	"eAL_PacketBegin":        12000,
	"eAL_Connected":          12000,
	"eAL_Disconnected":       12001,
	"eAL_CheckAccountResult": 12002,
	"eAL_PacketEnd":          13000,
}

func (x AL_Protocol) Enum() *AL_Protocol {
	p := new(AL_Protocol)
	*p = x
	return p
}
func (x AL_Protocol) String() string {
	return proto.EnumName(AL_Protocol_name, int32(x))
}
func (x AL_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *AL_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AL_Protocol_value, data, "AL_Protocol")
	if err != nil {
		return err
	}
	*x = AL_Protocol(value)
	return nil
}

type LS_Protocol int32

const (
	LS_Protocol_eLS_PacketBegin LS_Protocol = 13000
	// ----------------------------
	LS_Protocol_eLS_Connected               LS_Protocol = 13000
	LS_Protocol_eLS_Disconnected            LS_Protocol = 13001
	LS_Protocol_eLS_UpdatePlayerCountResult LS_Protocol = 13002
	// ----------------------------
	LS_Protocol_eSLS_PacketEnd LS_Protocol = 14000
)

var LS_Protocol_name = map[int32]string{
	13000: "eLS_PacketBegin",
	// Duplicate value: 13000: "eLS_Connected",
	13001: "eLS_Disconnected",
	13002: "eLS_UpdatePlayerCountResult",
	14000: "eSLS_PacketEnd",
}
var LS_Protocol_value = map[string]int32{
	"eLS_PacketBegin":             13000,
	"eLS_Connected":               13000,
	"eLS_Disconnected":            13001,
	"eLS_UpdatePlayerCountResult": 13002,
	"eSLS_PacketEnd":              14000,
}

func (x LS_Protocol) Enum() *LS_Protocol {
	p := new(LS_Protocol)
	*p = x
	return p
}
func (x LS_Protocol) String() string {
	return proto.EnumName(LS_Protocol_name, int32(x))
}
func (x LS_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *LS_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LS_Protocol_value, data, "LS_Protocol")
	if err != nil {
		return err
	}
	*x = LS_Protocol(value)
	return nil
}

type SL_Protocol int32

const (
	SL_Protocol_eSL_PacketBegin SL_Protocol = 14000
	// ----------------------------
	SL_Protocol_eSL_Connected         SL_Protocol = 14000
	SL_Protocol_eSL_Disconnected      SL_Protocol = 14001
	SL_Protocol_eSL_UpdatePlayerCount SL_Protocol = 14002
	// ----------------------------
	SL_Protocol_eSL_PacketEnd SL_Protocol = 20000
)

var SL_Protocol_name = map[int32]string{
	14000: "eSL_PacketBegin",
	// Duplicate value: 14000: "eSL_Connected",
	14001: "eSL_Disconnected",
	14002: "eSL_UpdatePlayerCount",
	20000: "eSL_PacketEnd",
}
var SL_Protocol_value = map[string]int32{
	"eSL_PacketBegin":       14000,
	"eSL_Connected":         14000,
	"eSL_Disconnected":      14001,
	"eSL_UpdatePlayerCount": 14002,
	"eSL_PacketEnd":         20000,
}

func (x SL_Protocol) Enum() *SL_Protocol {
	p := new(SL_Protocol)
	*p = x
	return p
}
func (x SL_Protocol) String() string {
	return proto.EnumName(SL_Protocol_name, int32(x))
}
func (x SL_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *SL_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SL_Protocol_value, data, "SL_Protocol")
	if err != nil {
		return err
	}
	*x = SL_Protocol(value)
	return nil
}

type SF_Protocol int32

const (
	SF_Protocol_eSF_PacketBegin SF_Protocol = 40000
	// ----------------------------
	SF_Protocol_eSF_Connected    SF_Protocol = 40000
	SF_Protocol_eSF_Disconnected SF_Protocol = 40001
	// Begin ClientScene & BattleScene
	SF_Protocol_eSF_Battle      SF_Protocol = 41000
	SF_Protocol_eSF_EnterBattle SF_Protocol = 41001
	SF_Protocol_eSF_ExitBattle  SF_Protocol = 41002
	SF_Protocol_eSF_BattleQuery SF_Protocol = 41003
	// ----------------------------
	SF_Protocol_eSF_PacketEnd SF_Protocol = 50000
)

var SF_Protocol_name = map[int32]string{
	40000: "eSF_PacketBegin",
	// Duplicate value: 40000: "eSF_Connected",
	40001: "eSF_Disconnected",
	41000: "eSF_Battle",
	41001: "eSF_EnterBattle",
	41002: "eSF_ExitBattle",
	41003: "eSF_BattleQuery",
	50000: "eSF_PacketEnd",
}
var SF_Protocol_value = map[string]int32{
	"eSF_PacketBegin":  40000,
	"eSF_Connected":    40000,
	"eSF_Disconnected": 40001,
	"eSF_Battle":       41000,
	"eSF_EnterBattle":  41001,
	"eSF_ExitBattle":   41002,
	"eSF_BattleQuery":  41003,
	"eSF_PacketEnd":    50000,
}

func (x SF_Protocol) Enum() *SF_Protocol {
	p := new(SF_Protocol)
	*p = x
	return p
}
func (x SF_Protocol) String() string {
	return proto.EnumName(SF_Protocol_name, int32(x))
}
func (x SF_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *SF_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SF_Protocol_value, data, "SF_Protocol")
	if err != nil {
		return err
	}
	*x = SF_Protocol(value)
	return nil
}

type FS_Protocol int32

const (
	FS_Protocol_eFS_PacketBegin FS_Protocol = 50000
	// ----------------------------
	FS_Protocol_eFS_Connected    FS_Protocol = 50000
	FS_Protocol_eFS_Disconnected FS_Protocol = 50001
	// Begin ClientScene & BattleScene
	FS_Protocol_eFS_Battle      FS_Protocol = 51000
	FS_Protocol_eFS_BattleBegin FS_Protocol = 51001
	FS_Protocol_eFS_BattleEnd   FS_Protocol = 51002
	FS_Protocol_eFS_BattleInfo  FS_Protocol = 51003
	// ----------------------------
	FS_Protocol_eFS_PacketEnd FS_Protocol = 60000
)

var FS_Protocol_name = map[int32]string{
	50000: "eFS_PacketBegin",
	// Duplicate value: 50000: "eFS_Connected",
	50001: "eFS_Disconnected",
	51000: "eFS_Battle",
	51001: "eFS_BattleBegin",
	51002: "eFS_BattleEnd",
	51003: "eFS_BattleInfo",
	60000: "eFS_PacketEnd",
}
var FS_Protocol_value = map[string]int32{
	"eFS_PacketBegin":  50000,
	"eFS_Connected":    50000,
	"eFS_Disconnected": 50001,
	"eFS_Battle":       51000,
	"eFS_BattleBegin":  51001,
	"eFS_BattleEnd":    51002,
	"eFS_BattleInfo":   51003,
	"eFS_PacketEnd":    60000,
}

func (x FS_Protocol) Enum() *FS_Protocol {
	p := new(FS_Protocol)
	*p = x
	return p
}
func (x FS_Protocol) String() string {
	return proto.EnumName(FS_Protocol_name, int32(x))
}
func (x FS_Protocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *FS_Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FS_Protocol_value, data, "FS_Protocol")
	if err != nil {
		return err
	}
	*x = FS_Protocol(value)
	return nil
}

func init() {
	proto.RegisterEnum("protobuf.LA_Protocol", LA_Protocol_name, LA_Protocol_value)
	proto.RegisterEnum("protobuf.AL_Protocol", AL_Protocol_name, AL_Protocol_value)
	proto.RegisterEnum("protobuf.LS_Protocol", LS_Protocol_name, LS_Protocol_value)
	proto.RegisterEnum("protobuf.SL_Protocol", SL_Protocol_name, SL_Protocol_value)
	proto.RegisterEnum("protobuf.SF_Protocol", SF_Protocol_name, SF_Protocol_value)
	proto.RegisterEnum("protobuf.FS_Protocol", FS_Protocol_name, FS_Protocol_value)
}
