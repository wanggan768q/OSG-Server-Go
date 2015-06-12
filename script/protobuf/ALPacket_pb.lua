-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf"
local PB_PacketServerDefine_pb = require("PB_PacketServerDefine_pb")
module('ALPacket_pb')


local AL_CHECKACCOUNTRESULT = protobuf.Descriptor();
local AL_CHECKACCOUNTRESULT_RESULT = protobuf.EnumDescriptor();
local AL_CHECKACCOUNTRESULT_RESULT_OK_ENUM = protobuf.EnumValueDescriptor();
local AL_CHECKACCOUNTRESULT_RESULT_SERVERERROR_ENUM = protobuf.EnumValueDescriptor();
local AL_CHECKACCOUNTRESULT_RESULT_USERNOTFOUND_ENUM = protobuf.EnumValueDescriptor();
local AL_CHECKACCOUNTRESULT_RESULT_AUTH_FAILED_ENUM = protobuf.EnumValueDescriptor();
local AL_CHECKACCOUNTRESULT_RESULT_ISONFIRE_ENUM = protobuf.EnumValueDescriptor();
local AL_CHECKACCOUNTRESULT_RESULT_FIELD = protobuf.FieldDescriptor();
local AL_CHECKACCOUNTRESULT_SERVER_TIME_FIELD = protobuf.FieldDescriptor();
local AL_CHECKACCOUNTRESULT_SESSIONKEY_FIELD = protobuf.FieldDescriptor();
local AL_CHECKACCOUNTRESULT_UID_FIELD = protobuf.FieldDescriptor();
local AL_CHECKACCOUNTRESULT_ERRMSG_FIELD = protobuf.FieldDescriptor();

AL_CHECKACCOUNTRESULT_RESULT_OK_ENUM.name = "OK"
AL_CHECKACCOUNTRESULT_RESULT_OK_ENUM.index = 0
AL_CHECKACCOUNTRESULT_RESULT_OK_ENUM.number = 0
AL_CHECKACCOUNTRESULT_RESULT_SERVERERROR_ENUM.name = "SERVERERROR"
AL_CHECKACCOUNTRESULT_RESULT_SERVERERROR_ENUM.index = 1
AL_CHECKACCOUNTRESULT_RESULT_SERVERERROR_ENUM.number = 1
AL_CHECKACCOUNTRESULT_RESULT_USERNOTFOUND_ENUM.name = "USERNOTFOUND"
AL_CHECKACCOUNTRESULT_RESULT_USERNOTFOUND_ENUM.index = 2
AL_CHECKACCOUNTRESULT_RESULT_USERNOTFOUND_ENUM.number = 2
AL_CHECKACCOUNTRESULT_RESULT_AUTH_FAILED_ENUM.name = "AUTH_FAILED"
AL_CHECKACCOUNTRESULT_RESULT_AUTH_FAILED_ENUM.index = 3
AL_CHECKACCOUNTRESULT_RESULT_AUTH_FAILED_ENUM.number = 3
AL_CHECKACCOUNTRESULT_RESULT_ISONFIRE_ENUM.name = "ISONFIRE"
AL_CHECKACCOUNTRESULT_RESULT_ISONFIRE_ENUM.index = 4
AL_CHECKACCOUNTRESULT_RESULT_ISONFIRE_ENUM.number = 4
AL_CHECKACCOUNTRESULT_RESULT.name = "Result"
AL_CHECKACCOUNTRESULT_RESULT.full_name = ".protobuf.AL_CheckAccountResult.Result"
AL_CHECKACCOUNTRESULT_RESULT.values = {AL_CHECKACCOUNTRESULT_RESULT_OK_ENUM,AL_CHECKACCOUNTRESULT_RESULT_SERVERERROR_ENUM,AL_CHECKACCOUNTRESULT_RESULT_USERNOTFOUND_ENUM,AL_CHECKACCOUNTRESULT_RESULT_AUTH_FAILED_ENUM,AL_CHECKACCOUNTRESULT_RESULT_ISONFIRE_ENUM}
AL_CHECKACCOUNTRESULT_RESULT_FIELD.name = "result"
AL_CHECKACCOUNTRESULT_RESULT_FIELD.full_name = ".protobuf.AL_CheckAccountResult.result"
AL_CHECKACCOUNTRESULT_RESULT_FIELD.number = 1
AL_CHECKACCOUNTRESULT_RESULT_FIELD.index = 0
AL_CHECKACCOUNTRESULT_RESULT_FIELD.label = 2
AL_CHECKACCOUNTRESULT_RESULT_FIELD.has_default_value = true
AL_CHECKACCOUNTRESULT_RESULT_FIELD.default_value = OK
AL_CHECKACCOUNTRESULT_RESULT_FIELD.enum_type = AL_CHECKACCOUNTRESULT_RESULT
AL_CHECKACCOUNTRESULT_RESULT_FIELD.type = 14
AL_CHECKACCOUNTRESULT_RESULT_FIELD.cpp_type = 8

AL_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.name = "server_time"
AL_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.full_name = ".protobuf.AL_CheckAccountResult.server_time"
AL_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.number = 2
AL_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.index = 1
AL_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.label = 2
AL_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.has_default_value = false
AL_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.default_value = 0
AL_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.type = 13
AL_CHECKACCOUNTRESULT_SERVER_TIME_FIELD.cpp_type = 3

AL_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.name = "sessionKey"
AL_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.full_name = ".protobuf.AL_CheckAccountResult.sessionKey"
AL_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.number = 3
AL_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.index = 2
AL_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.label = 2
AL_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.has_default_value = false
AL_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.default_value = ""
AL_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.type = 9
AL_CHECKACCOUNTRESULT_SESSIONKEY_FIELD.cpp_type = 9

AL_CHECKACCOUNTRESULT_UID_FIELD.name = "uid"
AL_CHECKACCOUNTRESULT_UID_FIELD.full_name = ".protobuf.AL_CheckAccountResult.uid"
AL_CHECKACCOUNTRESULT_UID_FIELD.number = 4
AL_CHECKACCOUNTRESULT_UID_FIELD.index = 3
AL_CHECKACCOUNTRESULT_UID_FIELD.label = 2
AL_CHECKACCOUNTRESULT_UID_FIELD.has_default_value = false
AL_CHECKACCOUNTRESULT_UID_FIELD.default_value = ""
AL_CHECKACCOUNTRESULT_UID_FIELD.type = 9
AL_CHECKACCOUNTRESULT_UID_FIELD.cpp_type = 9

AL_CHECKACCOUNTRESULT_ERRMSG_FIELD.name = "errmsg"
AL_CHECKACCOUNTRESULT_ERRMSG_FIELD.full_name = ".protobuf.AL_CheckAccountResult.errmsg"
AL_CHECKACCOUNTRESULT_ERRMSG_FIELD.number = 5
AL_CHECKACCOUNTRESULT_ERRMSG_FIELD.index = 4
AL_CHECKACCOUNTRESULT_ERRMSG_FIELD.label = 1
AL_CHECKACCOUNTRESULT_ERRMSG_FIELD.has_default_value = false
AL_CHECKACCOUNTRESULT_ERRMSG_FIELD.default_value = ""
AL_CHECKACCOUNTRESULT_ERRMSG_FIELD.type = 9
AL_CHECKACCOUNTRESULT_ERRMSG_FIELD.cpp_type = 9

AL_CHECKACCOUNTRESULT.name = "AL_CheckAccountResult"
AL_CHECKACCOUNTRESULT.full_name = ".protobuf.AL_CheckAccountResult"
AL_CHECKACCOUNTRESULT.nested_types = {}
AL_CHECKACCOUNTRESULT.enum_types = {AL_CHECKACCOUNTRESULT_RESULT}
AL_CHECKACCOUNTRESULT.fields = {AL_CHECKACCOUNTRESULT_RESULT_FIELD, AL_CHECKACCOUNTRESULT_SERVER_TIME_FIELD, AL_CHECKACCOUNTRESULT_SESSIONKEY_FIELD, AL_CHECKACCOUNTRESULT_UID_FIELD, AL_CHECKACCOUNTRESULT_ERRMSG_FIELD}
AL_CHECKACCOUNTRESULT.is_extendable = false
AL_CHECKACCOUNTRESULT.extensions = {}

AL_CheckAccountResult = protobuf.Message(AL_CHECKACCOUNTRESULT)

