-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf"
local PB_PacketServerDefine_pb = require("PB_PacketServerDefine_pb")
module('LSPacket_pb')


local LS_UPDATEPLAYERCOUNTRESULT = protobuf.Descriptor();
local LS_UPDATEPLAYERCOUNTRESULT_RESULT = protobuf.EnumDescriptor();
local LS_UPDATEPLAYERCOUNTRESULT_RESULT_OK_ENUM = protobuf.EnumValueDescriptor();
local LS_UPDATEPLAYERCOUNTRESULT_RESULT_SERVERERROR_ENUM = protobuf.EnumValueDescriptor();
local LS_UPDATEPLAYERCOUNTRESULT_RESULT_FIELD = protobuf.FieldDescriptor();
local LS_UPDATEPLAYERCOUNTRESULT_SERVER_TIME_FIELD = protobuf.FieldDescriptor();

LS_UPDATEPLAYERCOUNTRESULT_RESULT_OK_ENUM.name = "OK"
LS_UPDATEPLAYERCOUNTRESULT_RESULT_OK_ENUM.index = 0
LS_UPDATEPLAYERCOUNTRESULT_RESULT_OK_ENUM.number = 0
LS_UPDATEPLAYERCOUNTRESULT_RESULT_SERVERERROR_ENUM.name = "SERVERERROR"
LS_UPDATEPLAYERCOUNTRESULT_RESULT_SERVERERROR_ENUM.index = 1
LS_UPDATEPLAYERCOUNTRESULT_RESULT_SERVERERROR_ENUM.number = 1
LS_UPDATEPLAYERCOUNTRESULT_RESULT.name = "Result"
LS_UPDATEPLAYERCOUNTRESULT_RESULT.full_name = ".protobuf.LS_UpdatePlayerCountResult.Result"
LS_UPDATEPLAYERCOUNTRESULT_RESULT.values = {LS_UPDATEPLAYERCOUNTRESULT_RESULT_OK_ENUM,LS_UPDATEPLAYERCOUNTRESULT_RESULT_SERVERERROR_ENUM}
LS_UPDATEPLAYERCOUNTRESULT_RESULT_FIELD.name = "result"
LS_UPDATEPLAYERCOUNTRESULT_RESULT_FIELD.full_name = ".protobuf.LS_UpdatePlayerCountResult.result"
LS_UPDATEPLAYERCOUNTRESULT_RESULT_FIELD.number = 1
LS_UPDATEPLAYERCOUNTRESULT_RESULT_FIELD.index = 0
LS_UPDATEPLAYERCOUNTRESULT_RESULT_FIELD.label = 2
LS_UPDATEPLAYERCOUNTRESULT_RESULT_FIELD.has_default_value = true
LS_UPDATEPLAYERCOUNTRESULT_RESULT_FIELD.default_value = OK
LS_UPDATEPLAYERCOUNTRESULT_RESULT_FIELD.enum_type = LS_UPDATEPLAYERCOUNTRESULT_RESULT
LS_UPDATEPLAYERCOUNTRESULT_RESULT_FIELD.type = 14
LS_UPDATEPLAYERCOUNTRESULT_RESULT_FIELD.cpp_type = 8

LS_UPDATEPLAYERCOUNTRESULT_SERVER_TIME_FIELD.name = "server_time"
LS_UPDATEPLAYERCOUNTRESULT_SERVER_TIME_FIELD.full_name = ".protobuf.LS_UpdatePlayerCountResult.server_time"
LS_UPDATEPLAYERCOUNTRESULT_SERVER_TIME_FIELD.number = 2
LS_UPDATEPLAYERCOUNTRESULT_SERVER_TIME_FIELD.index = 1
LS_UPDATEPLAYERCOUNTRESULT_SERVER_TIME_FIELD.label = 2
LS_UPDATEPLAYERCOUNTRESULT_SERVER_TIME_FIELD.has_default_value = false
LS_UPDATEPLAYERCOUNTRESULT_SERVER_TIME_FIELD.default_value = 0
LS_UPDATEPLAYERCOUNTRESULT_SERVER_TIME_FIELD.type = 13
LS_UPDATEPLAYERCOUNTRESULT_SERVER_TIME_FIELD.cpp_type = 3

LS_UPDATEPLAYERCOUNTRESULT.name = "LS_UpdatePlayerCountResult"
LS_UPDATEPLAYERCOUNTRESULT.full_name = ".protobuf.LS_UpdatePlayerCountResult"
LS_UPDATEPLAYERCOUNTRESULT.nested_types = {}
LS_UPDATEPLAYERCOUNTRESULT.enum_types = {LS_UPDATEPLAYERCOUNTRESULT_RESULT}
LS_UPDATEPLAYERCOUNTRESULT.fields = {LS_UPDATEPLAYERCOUNTRESULT_RESULT_FIELD, LS_UPDATEPLAYERCOUNTRESULT_SERVER_TIME_FIELD}
LS_UPDATEPLAYERCOUNTRESULT.is_extendable = false
LS_UPDATEPLAYERCOUNTRESULT.extensions = {}

LS_UpdatePlayerCountResult = protobuf.Message(LS_UPDATEPLAYERCOUNTRESULT)

