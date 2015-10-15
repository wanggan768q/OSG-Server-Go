-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf"
module('PB_PacketCommon_pb')


local PACKET = protobuf.Descriptor();
local PACKET_ID_FIELD = protobuf.FieldDescriptor();
local PACKET_METHOD_FIELD = protobuf.FieldDescriptor();
local PACKET_SERIALIZED_PACKET_FIELD = protobuf.FieldDescriptor();
local PACKET_TIMER_FIELD = protobuf.FieldDescriptor();
local PACKET_DESC_FIELD = protobuf.FieldDescriptor();
local RPCERRORRESPONSE = protobuf.Descriptor();
local RPCERRORRESPONSE_METHOD_FIELD = protobuf.FieldDescriptor();
local RPCERRORRESPONSE_TEXT_FIELD = protobuf.FieldDescriptor();
local VECTOR3 = protobuf.Descriptor();
local VECTOR3_X_FIELD = protobuf.FieldDescriptor();
local VECTOR3_Y_FIELD = protobuf.FieldDescriptor();
local VECTOR3_Z_FIELD = protobuf.FieldDescriptor();
local VECTOR2 = protobuf.Descriptor();
local VECTOR2_X_FIELD = protobuf.FieldDescriptor();
local VECTOR2_Y_FIELD = protobuf.FieldDescriptor();

PACKET_ID_FIELD.name = "id"
PACKET_ID_FIELD.full_name = ".protobuf.Packet.id"
PACKET_ID_FIELD.number = 1
PACKET_ID_FIELD.index = 0
PACKET_ID_FIELD.label = 2
PACKET_ID_FIELD.has_default_value = false
PACKET_ID_FIELD.default_value = 0
PACKET_ID_FIELD.type = 4
PACKET_ID_FIELD.cpp_type = 4

PACKET_METHOD_FIELD.name = "method"
PACKET_METHOD_FIELD.full_name = ".protobuf.Packet.method"
PACKET_METHOD_FIELD.number = 2
PACKET_METHOD_FIELD.index = 1
PACKET_METHOD_FIELD.label = 2
PACKET_METHOD_FIELD.has_default_value = false
PACKET_METHOD_FIELD.default_value = 0
PACKET_METHOD_FIELD.type = 13
PACKET_METHOD_FIELD.cpp_type = 3

PACKET_SERIALIZED_PACKET_FIELD.name = "serialized_packet"
PACKET_SERIALIZED_PACKET_FIELD.full_name = ".protobuf.Packet.serialized_packet"
PACKET_SERIALIZED_PACKET_FIELD.number = 3
PACKET_SERIALIZED_PACKET_FIELD.index = 2
PACKET_SERIALIZED_PACKET_FIELD.label = 1
PACKET_SERIALIZED_PACKET_FIELD.has_default_value = false
PACKET_SERIALIZED_PACKET_FIELD.default_value = ""
PACKET_SERIALIZED_PACKET_FIELD.type = 12
PACKET_SERIALIZED_PACKET_FIELD.cpp_type = 9

PACKET_TIMER_FIELD.name = "timer"
PACKET_TIMER_FIELD.full_name = ".protobuf.Packet.timer"
PACKET_TIMER_FIELD.number = 4
PACKET_TIMER_FIELD.index = 3
PACKET_TIMER_FIELD.label = 1
PACKET_TIMER_FIELD.has_default_value = false
PACKET_TIMER_FIELD.default_value = 0
PACKET_TIMER_FIELD.type = 4
PACKET_TIMER_FIELD.cpp_type = 4

PACKET_DESC_FIELD.name = "desc"
PACKET_DESC_FIELD.full_name = ".protobuf.Packet.desc"
PACKET_DESC_FIELD.number = 5
PACKET_DESC_FIELD.index = 4
PACKET_DESC_FIELD.label = 1
PACKET_DESC_FIELD.has_default_value = false
PACKET_DESC_FIELD.default_value = ""
PACKET_DESC_FIELD.type = 9
PACKET_DESC_FIELD.cpp_type = 9

PACKET.name = "Packet"
PACKET.full_name = ".protobuf.Packet"
PACKET.nested_types = {}
PACKET.enum_types = {}
PACKET.fields = {PACKET_ID_FIELD, PACKET_METHOD_FIELD, PACKET_SERIALIZED_PACKET_FIELD, PACKET_TIMER_FIELD, PACKET_DESC_FIELD}
PACKET.is_extendable = true
PACKET.extensions = {}
RPCERRORRESPONSE_METHOD_FIELD.name = "method"
RPCERRORRESPONSE_METHOD_FIELD.full_name = ".protobuf.RpcErrorResponse.method"
RPCERRORRESPONSE_METHOD_FIELD.number = 1
RPCERRORRESPONSE_METHOD_FIELD.index = 0
RPCERRORRESPONSE_METHOD_FIELD.label = 2
RPCERRORRESPONSE_METHOD_FIELD.has_default_value = false
RPCERRORRESPONSE_METHOD_FIELD.default_value = 0
RPCERRORRESPONSE_METHOD_FIELD.type = 13
RPCERRORRESPONSE_METHOD_FIELD.cpp_type = 3

RPCERRORRESPONSE_TEXT_FIELD.name = "text"
RPCERRORRESPONSE_TEXT_FIELD.full_name = ".protobuf.RpcErrorResponse.text"
RPCERRORRESPONSE_TEXT_FIELD.number = 2
RPCERRORRESPONSE_TEXT_FIELD.index = 1
RPCERRORRESPONSE_TEXT_FIELD.label = 2
RPCERRORRESPONSE_TEXT_FIELD.has_default_value = false
RPCERRORRESPONSE_TEXT_FIELD.default_value = ""
RPCERRORRESPONSE_TEXT_FIELD.type = 9
RPCERRORRESPONSE_TEXT_FIELD.cpp_type = 9

RPCERRORRESPONSE.name = "RpcErrorResponse"
RPCERRORRESPONSE.full_name = ".protobuf.RpcErrorResponse"
RPCERRORRESPONSE.nested_types = {}
RPCERRORRESPONSE.enum_types = {}
RPCERRORRESPONSE.fields = {RPCERRORRESPONSE_METHOD_FIELD, RPCERRORRESPONSE_TEXT_FIELD}
RPCERRORRESPONSE.is_extendable = false
RPCERRORRESPONSE.extensions = {}
VECTOR3_X_FIELD.name = "x"
VECTOR3_X_FIELD.full_name = ".protobuf.Vector3.x"
VECTOR3_X_FIELD.number = 1
VECTOR3_X_FIELD.index = 0
VECTOR3_X_FIELD.label = 2
VECTOR3_X_FIELD.has_default_value = false
VECTOR3_X_FIELD.default_value = 0.0
VECTOR3_X_FIELD.type = 2
VECTOR3_X_FIELD.cpp_type = 6

VECTOR3_Y_FIELD.name = "y"
VECTOR3_Y_FIELD.full_name = ".protobuf.Vector3.y"
VECTOR3_Y_FIELD.number = 2
VECTOR3_Y_FIELD.index = 1
VECTOR3_Y_FIELD.label = 2
VECTOR3_Y_FIELD.has_default_value = false
VECTOR3_Y_FIELD.default_value = 0.0
VECTOR3_Y_FIELD.type = 2
VECTOR3_Y_FIELD.cpp_type = 6

VECTOR3_Z_FIELD.name = "z"
VECTOR3_Z_FIELD.full_name = ".protobuf.Vector3.z"
VECTOR3_Z_FIELD.number = 3
VECTOR3_Z_FIELD.index = 2
VECTOR3_Z_FIELD.label = 2
VECTOR3_Z_FIELD.has_default_value = false
VECTOR3_Z_FIELD.default_value = 0.0
VECTOR3_Z_FIELD.type = 2
VECTOR3_Z_FIELD.cpp_type = 6

VECTOR3.name = "Vector3"
VECTOR3.full_name = ".protobuf.Vector3"
VECTOR3.nested_types = {}
VECTOR3.enum_types = {}
VECTOR3.fields = {VECTOR3_X_FIELD, VECTOR3_Y_FIELD, VECTOR3_Z_FIELD}
VECTOR3.is_extendable = false
VECTOR3.extensions = {}
VECTOR2_X_FIELD.name = "x"
VECTOR2_X_FIELD.full_name = ".protobuf.Vector2.x"
VECTOR2_X_FIELD.number = 1
VECTOR2_X_FIELD.index = 0
VECTOR2_X_FIELD.label = 2
VECTOR2_X_FIELD.has_default_value = false
VECTOR2_X_FIELD.default_value = 0.0
VECTOR2_X_FIELD.type = 2
VECTOR2_X_FIELD.cpp_type = 6

VECTOR2_Y_FIELD.name = "y"
VECTOR2_Y_FIELD.full_name = ".protobuf.Vector2.y"
VECTOR2_Y_FIELD.number = 2
VECTOR2_Y_FIELD.index = 1
VECTOR2_Y_FIELD.label = 2
VECTOR2_Y_FIELD.has_default_value = false
VECTOR2_Y_FIELD.default_value = 0.0
VECTOR2_Y_FIELD.type = 2
VECTOR2_Y_FIELD.cpp_type = 6

VECTOR2.name = "Vector2"
VECTOR2.full_name = ".protobuf.Vector2"
VECTOR2.nested_types = {}
VECTOR2.enum_types = {}
VECTOR2.fields = {VECTOR2_X_FIELD, VECTOR2_Y_FIELD}
VECTOR2.is_extendable = false
VECTOR2.extensions = {}

Packet = protobuf.Message(PACKET)
RpcErrorResponse = protobuf.Message(RPCERRORRESPONSE)
Vector2 = protobuf.Message(VECTOR2)
Vector3 = protobuf.Message(VECTOR3)

