// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: idl/enum.proto

package pb

import (
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

type RequestType int32

const (
	RequestType_HEART_BEAT      RequestType = 0
	RequestType_ONLINE          RequestType = 1
	RequestType_LIST_USERS      RequestType = 2
	RequestType_SEND_MESSAGE    RequestType = 3
	RequestType_OFFLINE         RequestType = 4
	RequestType_RECEIVE_MESSAGE RequestType = 5
)

// Enum value maps for RequestType.
var (
	RequestType_name = map[int32]string{
		0: "HEART_BEAT",
		1: "ONLINE",
		2: "LIST_USERS",
		3: "SEND_MESSAGE",
		4: "OFFLINE",
		5: "RECEIVE_MESSAGE",
	}
	RequestType_value = map[string]int32{
		"HEART_BEAT":      0,
		"ONLINE":          1,
		"LIST_USERS":      2,
		"SEND_MESSAGE":    3,
		"OFFLINE":         4,
		"RECEIVE_MESSAGE": 5,
	}
)

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}

func (x RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_enum_proto_enumTypes[0].Descriptor()
}

func (RequestType) Type() protoreflect.EnumType {
	return &file_idl_enum_proto_enumTypes[0]
}

func (x RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestType.Descriptor instead.
func (RequestType) EnumDescriptor() ([]byte, []int) {
	return file_idl_enum_proto_rawDescGZIP(), []int{0}
}

type NoticeType int32

const (
	NoticeType_Temp NoticeType = 0
)

// Enum value maps for NoticeType.
var (
	NoticeType_name = map[int32]string{
		0: "Temp",
	}
	NoticeType_value = map[string]int32{
		"Temp": 0,
	}
)

func (x NoticeType) Enum() *NoticeType {
	p := new(NoticeType)
	*p = x
	return p
}

func (x NoticeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NoticeType) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_enum_proto_enumTypes[1].Descriptor()
}

func (NoticeType) Type() protoreflect.EnumType {
	return &file_idl_enum_proto_enumTypes[1]
}

func (x NoticeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NoticeType.Descriptor instead.
func (NoticeType) EnumDescriptor() ([]byte, []int) {
	return file_idl_enum_proto_rawDescGZIP(), []int{1}
}

type MessageState int32

const (
	MessageState_Sent      MessageState = 0
	MessageState_DELIVERED MessageState = 1
	MessageState_Seen      MessageState = 2
)

// Enum value maps for MessageState.
var (
	MessageState_name = map[int32]string{
		0: "Sent",
		1: "DELIVERED",
		2: "Seen",
	}
	MessageState_value = map[string]int32{
		"Sent":      0,
		"DELIVERED": 1,
		"Seen":      2,
	}
)

func (x MessageState) Enum() *MessageState {
	p := new(MessageState)
	*p = x
	return p
}

func (x MessageState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageState) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_enum_proto_enumTypes[2].Descriptor()
}

func (MessageState) Type() protoreflect.EnumType {
	return &file_idl_enum_proto_enumTypes[2]
}

func (x MessageState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageState.Descriptor instead.
func (MessageState) EnumDescriptor() ([]byte, []int) {
	return file_idl_enum_proto_rawDescGZIP(), []int{2}
}

type AckType int32

const (
	AckType_REQUEST AckType = 0
	AckType_NOTICE  AckType = 1
)

// Enum value maps for AckType.
var (
	AckType_name = map[int32]string{
		0: "REQUEST",
		1: "NOTICE",
	}
	AckType_value = map[string]int32{
		"REQUEST": 0,
		"NOTICE":  1,
	}
)

func (x AckType) Enum() *AckType {
	p := new(AckType)
	*p = x
	return p
}

func (x AckType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AckType) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_enum_proto_enumTypes[3].Descriptor()
}

func (AckType) Type() protoreflect.EnumType {
	return &file_idl_enum_proto_enumTypes[3]
}

func (x AckType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AckType.Descriptor instead.
func (AckType) EnumDescriptor() ([]byte, []int) {
	return file_idl_enum_proto_rawDescGZIP(), []int{3}
}

type SessionType int32

const (
	SessionType_USER  SessionType = 0
	SessionType_GROUP SessionType = 1
)

// Enum value maps for SessionType.
var (
	SessionType_name = map[int32]string{
		0: "USER",
		1: "GROUP",
	}
	SessionType_value = map[string]int32{
		"USER":  0,
		"GROUP": 1,
	}
)

func (x SessionType) Enum() *SessionType {
	p := new(SessionType)
	*p = x
	return p
}

func (x SessionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SessionType) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_enum_proto_enumTypes[4].Descriptor()
}

func (SessionType) Type() protoreflect.EnumType {
	return &file_idl_enum_proto_enumTypes[4]
}

func (x SessionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SessionType.Descriptor instead.
func (SessionType) EnumDescriptor() ([]byte, []int) {
	return file_idl_enum_proto_rawDescGZIP(), []int{4}
}

type MessageType int32

const (
	MessageType_TXT   MessageType = 0
	MessageType_IMG   MessageType = 1
	MessageType_VIDEO MessageType = 2
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "TXT",
		1: "IMG",
		2: "VIDEO",
	}
	MessageType_value = map[string]int32{
		"TXT":   0,
		"IMG":   1,
		"VIDEO": 2,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_enum_proto_enumTypes[5].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_idl_enum_proto_enumTypes[5]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_idl_enum_proto_rawDescGZIP(), []int{5}
}

type UserState int32

const (
	UserState_ON     UserState = 0
	UserState_OFF    UserState = 1
	UserState_LOGOUT UserState = 2
)

// Enum value maps for UserState.
var (
	UserState_name = map[int32]string{
		0: "ON",
		1: "OFF",
		2: "LOGOUT",
	}
	UserState_value = map[string]int32{
		"ON":     0,
		"OFF":    1,
		"LOGOUT": 2,
	}
)

func (x UserState) Enum() *UserState {
	p := new(UserState)
	*p = x
	return p
}

func (x UserState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserState) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_enum_proto_enumTypes[6].Descriptor()
}

func (UserState) Type() protoreflect.EnumType {
	return &file_idl_enum_proto_enumTypes[6]
}

func (x UserState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserState.Descriptor instead.
func (UserState) EnumDescriptor() ([]byte, []int) {
	return file_idl_enum_proto_rawDescGZIP(), []int{6}
}

var File_idl_enum_proto protoreflect.FileDescriptor

var file_idl_enum_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x2a, 0x6d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x45, 0x41, 0x52, 0x54, 0x5f, 0x42, 0x45, 0x41,
	0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x53, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10,
	0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x04, 0x12, 0x13,
	0x0a, 0x0f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x10, 0x05, 0x2a, 0x16, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x65, 0x6d, 0x70, 0x10, 0x00, 0x2a, 0x31, 0x0a, 0x0c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x53,
	0x65, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x65, 0x65, 0x6e, 0x10, 0x02, 0x2a, 0x22,
	0x0a, 0x07, 0x41, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x49, 0x43, 0x45,
	0x10, 0x01, 0x2a, 0x22, 0x0a, 0x0b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x10, 0x01, 0x2a, 0x2a, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x58, 0x54, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x49, 0x4d, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f,
	0x10, 0x02, 0x2a, 0x28, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x46, 0x46, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x47, 0x4f, 0x55, 0x54, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_enum_proto_rawDescOnce sync.Once
	file_idl_enum_proto_rawDescData = file_idl_enum_proto_rawDesc
)

func file_idl_enum_proto_rawDescGZIP() []byte {
	file_idl_enum_proto_rawDescOnce.Do(func() {
		file_idl_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_enum_proto_rawDescData)
	})
	return file_idl_enum_proto_rawDescData
}

var file_idl_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 7)
var file_idl_enum_proto_goTypes = []interface{}{
	(RequestType)(0),  // 0: pb.RequestType
	(NoticeType)(0),   // 1: pb.NoticeType
	(MessageState)(0), // 2: pb.MessageState
	(AckType)(0),      // 3: pb.AckType
	(SessionType)(0),  // 4: pb.SessionType
	(MessageType)(0),  // 5: pb.MessageType
	(UserState)(0),    // 6: pb.UserState
}
var file_idl_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_enum_proto_init() }
func file_idl_enum_proto_init() {
	if File_idl_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_enum_proto_rawDesc,
			NumEnums:      7,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_idl_enum_proto_goTypes,
		DependencyIndexes: file_idl_enum_proto_depIdxs,
		EnumInfos:         file_idl_enum_proto_enumTypes,
	}.Build()
	File_idl_enum_proto = out.File
	file_idl_enum_proto_rawDesc = nil
	file_idl_enum_proto_goTypes = nil
	file_idl_enum_proto_depIdxs = nil
}
