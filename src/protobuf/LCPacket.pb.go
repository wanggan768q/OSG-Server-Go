// Code generated by protoc-gen-go.
// source: LCPacket.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	LCPacket.proto

It has these top-level messages:
	LC_CheckAccountResult
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

type LC_CheckAccountResult_Result int32

const (
	LC_CheckAccountResult_OK           LC_CheckAccountResult_Result = 1
	LC_CheckAccountResult_SERVERERROR  LC_CheckAccountResult_Result = 2
	LC_CheckAccountResult_USERNOTFOUND LC_CheckAccountResult_Result = 3
	LC_CheckAccountResult_AUTH_FAILED  LC_CheckAccountResult_Result = 4
	LC_CheckAccountResult_ISONFIRE     LC_CheckAccountResult_Result = 5
)

var LC_CheckAccountResult_Result_name = map[int32]string{
	1: "OK",
	2: "SERVERERROR",
	3: "USERNOTFOUND",
	4: "AUTH_FAILED",
	5: "ISONFIRE",
}
var LC_CheckAccountResult_Result_value = map[string]int32{
	"OK":           1,
	"SERVERERROR":  2,
	"USERNOTFOUND": 3,
	"AUTH_FAILED":  4,
	"ISONFIRE":     5,
}

func (x LC_CheckAccountResult_Result) Enum() *LC_CheckAccountResult_Result {
	p := new(LC_CheckAccountResult_Result)
	*p = x
	return p
}
func (x LC_CheckAccountResult_Result) String() string {
	return proto.EnumName(LC_CheckAccountResult_Result_name, int32(x))
}
func (x LC_CheckAccountResult_Result) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *LC_CheckAccountResult_Result) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LC_CheckAccountResult_Result_value, data, "LC_CheckAccountResult_Result")
	if err != nil {
		return err
	}
	*x = LC_CheckAccountResult_Result(value)
	return nil
}

type LC_CheckAccountResult struct {
	Result           *LC_CheckAccountResult_Result `protobuf:"varint,1,req,name=result,enum=protobuf.LC_CheckAccountResult_Result,def=1" json:"result,omitempty"`
	ServerTime       *uint32                       `protobuf:"varint,2,req,name=server_time" json:"server_time,omitempty"`
	SessionKey       *string                       `protobuf:"bytes,3,req,name=sessionKey" json:"sessionKey,omitempty"`
	Uid              *string                       `protobuf:"bytes,4,req,name=uid" json:"uid,omitempty"`
	Errmsg           *string                       `protobuf:"bytes,5,opt,name=errmsg" json:"errmsg,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *LC_CheckAccountResult) Reset()         { *m = LC_CheckAccountResult{} }
func (m *LC_CheckAccountResult) String() string { return proto.CompactTextString(m) }
func (*LC_CheckAccountResult) ProtoMessage()    {}

const Default_LC_CheckAccountResult_Result LC_CheckAccountResult_Result = LC_CheckAccountResult_OK

func (m *LC_CheckAccountResult) GetResult() LC_CheckAccountResult_Result {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return Default_LC_CheckAccountResult_Result
}

func (m *LC_CheckAccountResult) GetServerTime() uint32 {
	if m != nil && m.ServerTime != nil {
		return *m.ServerTime
	}
	return 0
}

func (m *LC_CheckAccountResult) GetSessionKey() string {
	if m != nil && m.SessionKey != nil {
		return *m.SessionKey
	}
	return ""
}

func (m *LC_CheckAccountResult) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

func (m *LC_CheckAccountResult) GetErrmsg() string {
	if m != nil && m.Errmsg != nil {
		return *m.Errmsg
	}
	return ""
}

func (m *LC_CheckAccountResult) SetResult(value LC_CheckAccountResult_Result) {
	if m != nil {
		m.Result = &value
	}
}

func (m *LC_CheckAccountResult) SetServerTime(value uint32) {
	if m != nil {
		m.ServerTime = &value
	}
}

func (m *LC_CheckAccountResult) SetSessionKey(value string) {
	if m != nil {
		m.SessionKey = &value
	}
}

func (m *LC_CheckAccountResult) SetUid(value string) {
	if m != nil {
		m.Uid = &value
	}
}

func (m *LC_CheckAccountResult) SetErrmsg(value string) {
	if m != nil {
		m.Errmsg = &value
	}
}

func init() {
	proto.RegisterEnum("protobuf.LC_CheckAccountResult_Result", LC_CheckAccountResult_Result_name, LC_CheckAccountResult_Result_value)
}
