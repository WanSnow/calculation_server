// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: proto/api/AcceptGame.proto

package accept_game

import (
	player_info "github.com/wansnow/calculation_server/model/player_info"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GameMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GameMsg) Reset() {
	*x = GameMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_AcceptGame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameMsg) ProtoMessage() {}

func (x *GameMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_AcceptGame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameMsg.ProtoReflect.Descriptor instead.
func (*GameMsg) Descriptor() ([]byte, []int) {
	return file_proto_api_AcceptGame_proto_rawDescGZIP(), []int{0}
}

type GameMsg_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerInfo *player_info.PlayerInfo `protobuf:"bytes,1,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
	MissionId  string                  `protobuf:"bytes,2,opt,name=mission_id,json=missionId,proto3" json:"mission_id,omitempty"`
}

func (x *GameMsg_Request) Reset() {
	*x = GameMsg_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_AcceptGame_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameMsg_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameMsg_Request) ProtoMessage() {}

func (x *GameMsg_Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_AcceptGame_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameMsg_Request.ProtoReflect.Descriptor instead.
func (*GameMsg_Request) Descriptor() ([]byte, []int) {
	return file_proto_api_AcceptGame_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GameMsg_Request) GetPlayerInfo() *player_info.PlayerInfo {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

func (x *GameMsg_Request) GetMissionId() string {
	if x != nil {
		return x.MissionId
	}
	return ""
}

type GameMsg_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port   int64  `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	GameId string `protobuf:"bytes,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *GameMsg_Response) Reset() {
	*x = GameMsg_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_AcceptGame_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameMsg_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameMsg_Response) ProtoMessage() {}

func (x *GameMsg_Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_AcceptGame_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameMsg_Response.ProtoReflect.Descriptor instead.
func (*GameMsg_Response) Descriptor() ([]byte, []int) {
	return file_proto_api_AcceptGame_proto_rawDescGZIP(), []int{0, 1}
}

func (x *GameMsg_Response) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *GameMsg_Response) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

var File_proto_api_AcceptGame_proto protoreflect.FileDescriptor

var file_proto_api_AcceptGame_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x43, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x07, 0x47, 0x61,
	0x6d, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x76, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x4c, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x37, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x32, 0x7b, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x12, 0x6d, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x47, 0x61,
	0x6d, 0x65, 0x12, 0x2e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x61, 0x6e, 0x73, 0x6e, 0x6f, 0x77, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x67, 0x61, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_AcceptGame_proto_rawDescOnce sync.Once
	file_proto_api_AcceptGame_proto_rawDescData = file_proto_api_AcceptGame_proto_rawDesc
)

func file_proto_api_AcceptGame_proto_rawDescGZIP() []byte {
	file_proto_api_AcceptGame_proto_rawDescOnce.Do(func() {
		file_proto_api_AcceptGame_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_AcceptGame_proto_rawDescData)
	})
	return file_proto_api_AcceptGame_proto_rawDescData
}

var file_proto_api_AcceptGame_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_api_AcceptGame_proto_goTypes = []interface{}{
	(*GameMsg)(nil),                // 0: CompetitionPlatform.proto.api.GameMsg
	(*GameMsg_Request)(nil),        // 1: CompetitionPlatform.proto.api.GameMsg.Request
	(*GameMsg_Response)(nil),       // 2: CompetitionPlatform.proto.api.GameMsg.Response
	(*player_info.PlayerInfo)(nil), // 3: CompetitionPlatform.proto.model.PlayerInfo
}
var file_proto_api_AcceptGame_proto_depIdxs = []int32{
	3, // 0: CompetitionPlatform.proto.api.GameMsg.Request.player_info:type_name -> CompetitionPlatform.proto.model.PlayerInfo
	1, // 1: CompetitionPlatform.proto.api.AcceptGame.AcceptGame:input_type -> CompetitionPlatform.proto.api.GameMsg.Request
	2, // 2: CompetitionPlatform.proto.api.AcceptGame.AcceptGame:output_type -> CompetitionPlatform.proto.api.GameMsg.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_api_AcceptGame_proto_init() }
func file_proto_api_AcceptGame_proto_init() {
	if File_proto_api_AcceptGame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_AcceptGame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_AcceptGame_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameMsg_Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_AcceptGame_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameMsg_Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_api_AcceptGame_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_AcceptGame_proto_goTypes,
		DependencyIndexes: file_proto_api_AcceptGame_proto_depIdxs,
		MessageInfos:      file_proto_api_AcceptGame_proto_msgTypes,
	}.Build()
	File_proto_api_AcceptGame_proto = out.File
	file_proto_api_AcceptGame_proto_rawDesc = nil
	file_proto_api_AcceptGame_proto_goTypes = nil
	file_proto_api_AcceptGame_proto_depIdxs = nil
}
