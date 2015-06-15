-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf"
local PB_PacketDefine_pb = require("PB_PacketDefine_pb")
module('LCPacket_pb')


local LC_CHECKACCOUNTRESULT = protobuf.Descriptor();
local LC_CHECKACCOUNTRESULT_RESULT = protobuf.EnumDescriptor();
local LC_CHECKACCOUNTRESULT_RESULT_OK_ENUM = protobuf.EnumValueDescriptor();
local LC_CHECKACCOUNTRESULT_RESULT_SERVERERROR_ENUM = protobuf.EnumValueDescriptor();
local LC_CHECKACCOUNTRESULT_RESULT_USERNOTFOUND_ENUM = protobuf.EnumValueDescriptor();
local LC_CHECKACCOUNTRESULT_RESULT_AUTH_FAILED_ENUM = protobuf.EnumValueDescriptor();
local LC_CHECKACCOUNTRESULT_RESULT_ISONFIRE_ENUM = protobuf.EnumValueDescriptor();
local LC_CHECKACCOUNTRESULT_RESULT_FIELD = protobuf.FieldDescriptor();
local LC_CHECKACCOUNTRESULT_SERVER_TIME_FIELD = protobuf.FieldDescriptor();
local LC_CHECKACCOUNTRESULT_SESSIONKEY_FIELD = protobuf.FieldDescriptor();
local LC_CHECKACCOUNTRESULT_UID_FIELD = protobuf.FieldDescriptor();
local LC_CHECKACCOUNTRESULT_GAMESERVERIP_FIELD = protobuf.FieldDescriptor();

LC_CHECKACCOUNTRESULT_RESULT_OK_ENUM.name = "OK"
LC_CHECKACCOUNTRESULT_RESULT_OK_ENUM.index = 0
LC_CHECKACCOUNTRESULT_RESULT_OK_ENUM.number = 1
LC_CHECKACCOUNTRESULT_RESULT_SERVERERROR_ENUM.name = "SERVERERROR"
LC_CHECKACCOUNTRESULT_RESULT_SERVERERROR_ENUM.index = 1
LC_CHECKACCOUNTRESULT_RESULT_SERVERERROR_ENUM.number = 2
LC_CHECKACCOUNTRESULT_RESULT_USERNOTFOUND_ENUM.name = "USERNOTFOUND"
LC_CHECKACCOUNTRESULT_RESULT_USERNOTFOUND_ENUM.index = 2
LC_CHECKACCOUNTRESULT_RESULT_USERNOTFOUND_ENUM.number = 3
LC_CHECKACCOUNTRESULT_RESULT_AUTH_FAILED_ENUM.name = "AUTH_FAILED"
LC_CHECKACCOUNTRESULT_RESULT_AUTH_FAILED_ENUM.index = 3
LC_CHECKACCOUNTRESULT_RESULT_AUTH_FAILED_ENUM.number = 4
LC_CHECKACCOUNTRESULT_RESULT_ISONFIRE_ENUM.name = "ISONFIRE"
LC_CHECKACCOUNTRESULT_RESULT_ISONFIRE_ENUM.index = 4
LC_CHECKACCOUNTRESULT_RESULT_ISONFIRE_ENUM.number = 5
LC_CHECKACCOUNTRESULT_RESULT.name = "Result"
LC_CHECKACCOUNTRESULT_RESULT.full_name = ".protobuf.LC_CheckAccountResult.Result"
LC_CHECKACCOUNTRESULT_RESULT.values = {LC_CHECKACCOUNTRESULT_RESULT_OK_ENUM,LC_CHECKACCOUNTRESULT_RESULT_SERVERERROR_ENUM,LC_CHECKACCOUNTRESULT_RESULT_USERNOTFOUND_ENUM,LC_CHECKACCOUNTRESULT_RESULT_AUTH_FAILED_ENUM,LC_CHECKACCOUNTRESULT_RESULT_ISONFIRE_ENUM}
LC_CHECKACCOUNTRESULT_RESULT_FIELD.name = "result"
LC_CHECKACCOUNTRESULT_RESULT_FIELD.full_name = ".protobuf.LC_CheckAccountResult.result"
LC_CHECKACCOUNTRESULT_RESULT_FIELD.number = 1
LC_CHECKACCOUNTRESULT_RESULT_FIELD.index = 0
LC_CHECKACCOUNTRESULT_RESULT_FIELD.label = 2
LC_CHECKACCOUNTRESULT_RESULT_FIELD.has_default_value = true
LC_CHECKACCOUNTRESULT_RESULT_FIELD.default_value = OK
LC_CHECKACCOUNTRESULT_RESULT_FIELD.enum_type = LC_CHECKACCOUNTRESULT_RESULT
LC_CHECKACCOUNTRESULT_RESULT_FIELD.type = 14
LC_CHECKACCOUNTRESULT_RESULT_FIELD.cpp_type = 8

LC_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.name = "server_time"
LC_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.full_name = ".protobuf.LC_CheckAccountResult.server_time"
LC_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.number = 2
LC_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.index = 1
LC_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.label = 2
LC_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.has_default_value = false
LC_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.default_value = 0
LC_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.type = 13
LC_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.cpp_type = 3

LC_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.name = "sessionKey"
LC_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.full_name = ".protobuf.LC_CheckAccountResult.sessionKey"
LC_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.number = 3
LC_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.index = 2
LC_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.label = 2
LC_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.has_default_value = false
LC_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.default_value = ""
LC_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.type = 9
LC_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.cpp_type = 9

LC_CHECKACCOUNTRESULT_UID_FIELD.name = "uid"
LC_CHECKACCOUNTRESULT_UID_FIELD.full_name = ".protobuf.LC_CheckAccountResult.uid"
LC_CHECKACCOUNTRESULT_UID_FIELD.number = 4
LC_CHECKACCOUNTRESULT_UID_FIELD.index = 3
LC_CHECKACCOUNTRESULT_UID_FIELD.label = 2
LC_CHECKACCOUNTRESULT_UID_FIELD.has_default_value = false
LC_CHECKACCOUNTRESULT_UID_FIELD.default_value = ""
LC_CHECKACCOUNTRESULT_UID_FIELD.type = 9
LC_CHECKACCOUNTRESULT_UID_FIELD.cpp_type = 9

LC_CHECKACCOUNTRESULT_GAMESERVERIP_FIELD.name = "gameServerIp"
LC_CHECKACCOUNTRESULT_GAMESERVERIP_FIELD.full_name = ".protobuf.LC_CheckAccountResult.gameServerIp"
LC_CHECKACCOUNTRESULT_GAMESERVERIP_FIELD.number = 5
LC_CHECKACCOUNTRESULT_GAMESERVERIP_FIELD.index = 4
LC_CHECKACCOUNTRESULT_GAMESERVERIP_FIELD.label = 1
LC_CHECKACCOUNTRESULT_GAMESERVERIP_FIELD.has_default_value = false
LC_CHECKACCOUNTRESULT_GAMESERVERIP_FIELD.default_value = ""
LC_CHECKACCOUNTRESULT_GAMESERVERIP_FIELD.type = 9
LC_CHECKACCOUNTRESULT_GAMESERVERIP_FIELD.cpp_type = 9

LC_CHECKACCOUNTRESULT.name = "LC_CheckAccountResult"
LC_CHECKACCOUNTRESULT.full_name = ".protobuf.LC_CheckAccountResult"
LC_CHECKACCOUNTRESULT.nested_types = {}
LC_CHECKACCOUNTRESULT.enum_types = {LC_CHECKACCOUNTRESULT_RESULT}
LC_CHECKACCOUNTRESULT.fields = {LC_CHECKACCOUNTRESULT_RESULT_FIELD, LC_CHECKACCOUNTRESULT_SERVER_TIME_FIELD, LC_CHECKACCOUNTRESULT_SESSIONKEY_FIELD, LC_CHECKACCOUNTRESULT_UID_FIELD, LC_CHECKACCOUNTRESULT_GAMESERVERIP_FIELD}
LC_CHECKACCOUNTRESULT.is_extendable = false
LC_CHECKACCOUNTRESULT.extensions = {}

LC_CheckAccountResult = protobuf.Message(LC_CHECKACCOUNTRESULT)

