// Code generated by protoc-gen-go.
// source: ALPacket.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	ALPacket.proto

It has these top-level messages:
	AL_CheckAccountResult
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

type AL_CheckAccountResult_Result int32

const (
	AL_CheckAccountResult_OK           AL_CheckAccountResult_Result = 1
	AL_CheckAccountResult_SERVERERROR  AL_CheckAccountResult_Result = 2
	AL_CheckAccountResult_USERNOTFOUND AL_CheckAccountResult_Result = 3
	AL_CheckAccountResult_AUTH_FAILED  AL_CheckAccountResult_Result = 4
	AL_CheckAccountResult_ISONFIRE     AL_CheckAccountResult_Result = 5
)

var AL_CheckAccountResult_Result_name = map[int32]string{
	1: "OK",
	2: "SERVERERROR",
	3: "USERNOTFOUND",
	4: "AUTH_FAILED",
	5: "ISONFIRE",
}
var AL_CheckAccountResult_Result_value = map[string]int32{
	"OK":           1,
	"SERVERERROR":  2,
	"USERNOTFOUND": 3,
	"AUTH_FAILED":  4,
	"ISONFIRE":     5,
}

func (x AL_CheckAccountResult_Result) Enum() *AL_CheckAccountResult_Result {
	p := new(AL_CheckAccountResult_Result)
	*p = x
	return p
}
func (x AL_CheckAccountResult_Result) String() string {
	return proto.EnumName(AL_CheckAccountResult_Result_name, int32(x))
}
func (x AL_CheckAccountResult_Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *AL_CheckAccountResult_Result) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AL_CheckAccountResult_Result_value, data, "AL_CheckAccountResult_Result")
	if err != nil {
		return err
	}
	*x = AL_CheckAccountResult_Result(value)
	return nil
}

type AL_CheckAccountResult struct {
	Result           *AL_CheckAccountResult_Result `protobuf:"varint,1,req,name=result,enum=protobuf.AL_CheckAccountResult_Result" json:"result,omitempty"`
	SessionKey       *string                       `protobuf:"bytes,2,req,name=sessionKey" json:"sessionKey,omitempty"`
	Uid              *string                       `protobuf:"bytes,3,req,name=uid" json:"uid,omitempty"`
	Errmsg           *string                       `protobuf:"bytes,4,opt,name=errmsg" json:"errmsg,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *AL_CheckAccountResult) Reset()         { *m = AL_CheckAccountResult{} }
func (m *AL_CheckAccountResult) String() string { return proto.CompactTextString(m) }
func (*AL_CheckAccountResult) ProtoMessage()    {}

func (m *AL_CheckAccountResult) GetResult() AL_CheckAccountResult_Result {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return AL_CheckAccountResult_OK
}

func (m *AL_CheckAccountResult) GetSessionKey() string {
	if m != nil && m.SessionKey != nil {
		return *m.SessionKey
	}
	return ""
}

func (m *AL_CheckAccountResult) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

func (m *AL_CheckAccountResult) GetErrmsg() string {
	if m != nil && m.Errmsg != nil {
		return *m.Errmsg
	}
	return ""
}

func (m *AL_CheckAccountResult) SetResult(value AL_CheckAccountResult_Result) {
	if m != nil {
		m.Result = &value
	}
}

func (m *AL_CheckAccountResult) SetSessionKey(value string) {
	if m != nil {
		m.SessionKey = &value
	}
}

func (m *AL_CheckAccountResult) SetUid(value string) {
	if m != nil {
		m.Uid = &value
	}
}

func (m *AL_CheckAccountResult) SetErrmsg(value string) {
	if m != nil {
		m.Errmsg = &value
	}
}

func init() {
	proto.RegisterEnum("protobuf.AL_CheckAccountResult_Result", AL_CheckAccountResult_Result_name, AL_CheckAccountResult_Result_value)
}
