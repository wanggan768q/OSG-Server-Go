// Code generated by protoc-gen-go.
// source: PB_PacketDefine.proto
// DO NOT EDIT!

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CL_Protocol int32

const (
	CL_Protocol_eCL_ZERO        CL_Protocol = 0
	CL_Protocol_eCL_PacketBegin CL_Protocol = 33554432
	// ----------------------------
	CL_Protocol_eCL_Connected    CL_Protocol = 33554432
	CL_Protocol_eCL_Disconnected CL_Protocol = 33554433
	CL_Protocol_eCL_CheckAccount CL_Protocol = 33554434
	// ----------------------------
	CL_Protocol_eCL_PacketEnd CL_Protocol = 34603008
)

var CL_Protocol_name = map[int32]string{
	0:        "eCL_ZERO",
	33554432: "eCL_PacketBegin",
	// Duplicate value: 33554432: "eCL_Connected",
	33554433: "eCL_Disconnected",
	33554434: "eCL_CheckAccount",
	34603008: "eCL_PacketEnd",
}
var CL_Protocol_value = map[string]int32{
	"eCL_ZERO":         0,
	"eCL_PacketBegin":  33554432,
	"eCL_Connected":    33554432,
	"eCL_Disconnected": 33554433,
	"eCL_CheckAccount": 33554434,
	"eCL_PacketEnd":    34603008,
}

func (x CL_Protocol) String() string {
	return proto.EnumName(CL_Protocol_name, int32(x))
}
func (CL_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

type LC_Protocol int32

const (
	LC_Protocol_eLC_ZERO        LC_Protocol = 0
	LC_Protocol_eLC_PacketBegin LC_Protocol = 34603008
	// ----------------------------
	LC_Protocol_eLC_Connected          LC_Protocol = 34603008
	LC_Protocol_eLC_Disconnected       LC_Protocol = 34603009
	LC_Protocol_eLC_CheckAccountResult LC_Protocol = 34603010
	// ----------------------------
	LC_Protocol_eLC_PacketEnd LC_Protocol = 35651584
)

var LC_Protocol_name = map[int32]string{
	0:        "eLC_ZERO",
	34603008: "eLC_PacketBegin",
	// Duplicate value: 34603008: "eLC_Connected",
	34603009: "eLC_Disconnected",
	34603010: "eLC_CheckAccountResult",
	35651584: "eLC_PacketEnd",
}
var LC_Protocol_value = map[string]int32{
	"eLC_ZERO":               0,
	"eLC_PacketBegin":        34603008,
	"eLC_Connected":          34603008,
	"eLC_Disconnected":       34603009,
	"eLC_CheckAccountResult": 34603010,
	"eLC_PacketEnd":          35651584,
}

func (x LC_Protocol) String() string {
	return proto.EnumName(LC_Protocol_name, int32(x))
}
func (LC_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

type CS_Protocol int32

const (
	CS_Protocol_eCS_ZERO        CS_Protocol = 0
	CS_Protocol_eCS_PacketBegin CS_Protocol = 50331648
	// ----------------------------
	CS_Protocol_eCS_Connected    CS_Protocol = 50331648
	CS_Protocol_eCS_Disconnected CS_Protocol = 50331649
	CS_Protocol_eCS_CheckSession CS_Protocol = 50331650
	CS_Protocol_eCS_Ping         CS_Protocol = 50331651
	// ----------------------------
	CS_Protocol_eCS_PacketEnd CS_Protocol = 83886080
)

var CS_Protocol_name = map[int32]string{
	0:        "eCS_ZERO",
	50331648: "eCS_PacketBegin",
	// Duplicate value: 50331648: "eCS_Connected",
	50331649: "eCS_Disconnected",
	50331650: "eCS_CheckSession",
	50331651: "eCS_Ping",
	83886080: "eCS_PacketEnd",
}
var CS_Protocol_value = map[string]int32{
	"eCS_ZERO":         0,
	"eCS_PacketBegin":  50331648,
	"eCS_Connected":    50331648,
	"eCS_Disconnected": 50331649,
	"eCS_CheckSession": 50331650,
	"eCS_Ping":         50331651,
	"eCS_PacketEnd":    83886080,
}

func (x CS_Protocol) String() string {
	return proto.EnumName(CS_Protocol_name, int32(x))
}
func (CS_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

type SC_Protocol int32

const (
	SC_Protocol_eSC_ZERO        SC_Protocol = 0
	SC_Protocol_eSC_PacketBegin SC_Protocol = 83886080
	// ----------------------------
	SC_Protocol_eSC_Connected          SC_Protocol = 83886080
	SC_Protocol_eSC_Disconnected       SC_Protocol = 83886081
	SC_Protocol_eSC_CheckSessionResult SC_Protocol = 83886082
	SC_Protocol_eCS_PingResult         SC_Protocol = 83886083
	// ----------------------------
	SC_Protocol_eSC_PacketEnd SC_Protocol = 117440512
)

var SC_Protocol_name = map[int32]string{
	0:        "eSC_ZERO",
	83886080: "eSC_PacketBegin",
	// Duplicate value: 83886080: "eSC_Connected",
	83886081:  "eSC_Disconnected",
	83886082:  "eSC_CheckSessionResult",
	83886083:  "eCS_PingResult",
	117440512: "eSC_PacketEnd",
}
var SC_Protocol_value = map[string]int32{
	"eSC_ZERO":               0,
	"eSC_PacketBegin":        83886080,
	"eSC_Connected":          83886080,
	"eSC_Disconnected":       83886081,
	"eSC_CheckSessionResult": 83886082,
	"eCS_PingResult":         83886083,
	"eSC_PacketEnd":          117440512,
}

func (x SC_Protocol) String() string {
	return proto.EnumName(SC_Protocol_name, int32(x))
}
func (SC_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func init() {
	proto.RegisterEnum("protobuf.CL_Protocol", CL_Protocol_name, CL_Protocol_value)
	proto.RegisterEnum("protobuf.LC_Protocol", LC_Protocol_name, LC_Protocol_value)
	proto.RegisterEnum("protobuf.CS_Protocol", CS_Protocol_name, CS_Protocol_value)
	proto.RegisterEnum("protobuf.SC_Protocol", SC_Protocol_name, SC_Protocol_value)
}

var fileDescriptor7 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0x2d, 0x0b, 0x43, 0xc6, 0x1f, 0xc6, 0x46, 0xb0, 0x31, 0x71, 0x4f, 0x58, 0xb8, 0x71,
	0xe3, 0xd6, 0x16, 0x12, 0x17, 0x24, 0x36, 0xdc, 0x9d, 0x1b, 0x22, 0xe3, 0x80, 0x0d, 0x64, 0xc6,
	0xd8, 0x76, 0x5f, 0xe5, 0x19, 0x5c, 0xb9, 0xf7, 0x01, 0x7c, 0x82, 0x3e, 0x9a, 0xf7, 0xce, 0x68,
	0x99, 0x29, 0x2b, 0x72, 0xcf, 0x21, 0x67, 0x3e, 0xbe, 0xc0, 0xfa, 0x69, 0x3c, 0x4f, 0x9f, 0xc4,
	0x5a, 0x16, 0x63, 0xb9, 0xcc, 0x94, 0xbc, 0x7e, 0x7d, 0xd3, 0x85, 0x0e, 0xbb, 0xe6, 0x63, 0x51,
	0x2e, 0x47, 0x9f, 0x01, 0x3b, 0x4a, 0xa6, 0xf3, 0x94, 0x6e, 0xa1, 0x37, 0xe1, 0x31, 0xeb, 0x4a,
	0xbc, 0x1f, 0x27, 0xb3, 0x07, 0x7e, 0x10, 0x0e, 0x58, 0x8f, 0x2e, 0xbb, 0x10, 0xcb, 0x55, 0xa6,
	0x78, 0x55, 0x55, 0x3c, 0x3c, 0x67, 0x27, 0x94, 0x27, 0x5a, 0x29, 0x29, 0x0a, 0xf9, 0x6c, 0xd3,
	0x0b, 0xc6, 0x29, 0x1d, 0x67, 0xb9, 0x68, 0x8a, 0x77, 0xa7, 0x48, 0x5e, 0xa4, 0x58, 0xdf, 0x09,
	0xa1, 0x4b, 0x55, 0xf0, 0x0f, 0x67, 0xc7, 0xee, 0x4f, 0x14, 0xed, 0xd4, 0xfc, 0xb2, 0xc3, 0x83,
	0xd1, 0x17, 0x72, 0x4d, 0x13, 0x9f, 0x0b, 0x6f, 0x87, 0x8b, 0x5a, 0x8f, 0xab, 0xb6, 0x7b, 0x98,
	0xbb, 0x5c, 0xb5, 0x7d, 0x1e, 0xd3, 0x16, 0x17, 0x16, 0x57, 0x6c, 0x60, 0xbe, 0xee, 0x70, 0xcd,
	0x64, 0x5e, 0x6e, 0x88, 0x6e, 0xb7, 0xe6, 0xd2, 0x55, 0x67, 0x86, 0xee, 0x9b, 0xac, 0x41, 0xcb,
	0x1a, 0x78, 0xd6, 0xa0, 0x6d, 0x2d, 0xb2, 0xbf, 0x16, 0x7c, 0x6b, 0x91, 0x95, 0x03, 0x7b, 0xd6,
	0x9a, 0xc2, 0xd0, 0x81, 0xcc, 0xf3, 0x4c, 0x2b, 0xb2, 0x16, 0x85, 0x3d, 0xfb, 0x5a, 0x9a, 0xa9,
	0x15, 0xdf, 0x3a, 0xc3, 0x1e, 0xe8, 0xd0, 0x80, 0xfe, 0x20, 0x28, 0xb4, 0x34, 0x82, 0xa7, 0x11,
	0xda, 0x1a, 0xab, 0xa1, 0xd9, 0x03, 0x5f, 0x23, 0xa6, 0xc4, 0x03, 0x7b, 0x1a, 0xb1, 0x20, 0x8d,
	0x90, 0x78, 0xa0, 0x8d, 0x46, 0xac, 0xfb, 0xec, 0xf4, 0x1f, 0xf7, 0x2f, 0xde, 0x3a, 0x8f, 0x78,
	0xd0, 0xb7, 0x04, 0x1d, 0x77, 0xee, 0x83, 0xc5, 0xa1, 0xf9, 0x87, 0xde, 0xfc, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x60, 0x7e, 0x66, 0x1e, 0xc1, 0x02, 0x00, 0x00,
}
