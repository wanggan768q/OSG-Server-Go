-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf"
local PB_PacketServerDefine_pb = require("PB_PacketServerDefine_pb")
module('FSPacket_pb')


local FS_BATTLE = protobuf.Descriptor();
local FS_BATTLEBEGIN = protobuf.Descriptor();
local FS_BATTLEEND = protobuf.Descriptor();
local FS_BATTLEINFO = protobuf.Descriptor();

FS_BATTLE.name = "FS_Battle"
FS_BATTLE.full_name = ".protobuf.FS_Battle"
FS_BATTLE.nested_types = {}
FS_BATTLE.enum_types = {}
FS_BATTLE.fields = {}
FS_BATTLE.is_extendable = false
FS_BATTLE.extensions = {}
FS_BATTLEBEGIN.name = "FS_BattleBegin"
FS_BATTLEBEGIN.full_name = ".protobuf.FS_BattleBegin"
FS_BATTLEBEGIN.nested_types = {}
FS_BATTLEBEGIN.enum_types = {}
FS_BATTLEBEGIN.fields = {}
FS_BATTLEBEGIN.is_extendable = false
FS_BATTLEBEGIN.extensions = {}
FS_BATTLEEND.name = "FS_BattleEnd"
FS_BATTLEEND.full_name = ".protobuf.FS_BattleEnd"
FS_BATTLEEND.nested_types = {}
FS_BATTLEEND.enum_types = {}
FS_BATTLEEND.fields = {}
FS_BATTLEEND.is_extendable = false
FS_BATTLEEND.extensions = {}
FS_BATTLEINFO.name = "FS_BattleInfo"
FS_BATTLEINFO.full_name = ".protobuf.FS_BattleInfo"
FS_BATTLEINFO.nested_types = {}
FS_BATTLEINFO.enum_types = {}
FS_BATTLEINFO.fields = {}
FS_BATTLEINFO.is_extendable = false
FS_BATTLEINFO.extensions = {}

FS_Battle = protobuf.Message(FS_BATTLE)
FS_BattleBegin = protobuf.Message(FS_BATTLEBEGIN)
FS_BattleEnd = protobuf.Message(FS_BATTLEEND)
FS_BattleInfo = protobuf.Message(FS_BATTLEINFO)
