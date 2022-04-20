// 客户端消息

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: client.proto

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

type ClientActionType int32

const (
	ClientActionType_LOGIN    ClientActionType = 0
	ClientActionType_REGISTER ClientActionType = 1
	ClientActionType_LOGOUT   ClientActionType = 2
)

// Enum value maps for ClientActionType.
var (
	ClientActionType_name = map[int32]string{
		0: "LOGIN",
		1: "REGISTER",
		2: "LOGOUT",
	}
	ClientActionType_value = map[string]int32{
		"LOGIN":    0,
		"REGISTER": 1,
		"LOGOUT":   2,
	}
)

func (x ClientActionType) Enum() *ClientActionType {
	p := new(ClientActionType)
	*p = x
	return p
}

func (x ClientActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_client_proto_enumTypes[0].Descriptor()
}

func (ClientActionType) Type() protoreflect.EnumType {
	return &file_client_proto_enumTypes[0]
}

func (x ClientActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientActionType.Descriptor instead.
func (ClientActionType) EnumDescriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

type ClientMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action ClientActionType `protobuf:"varint,1,opt,name=action,proto3,enum=ClientActionType" json:"action,omitempty"` // 动作类型
	// Types that are assignable to Data:
	//	*ClientMsg_LoginMsgData
	//	*ClientMsg_RegisterMsgData
	//	*ClientMsg_LogoutUserId
	Data isClientMsg_Data `protobuf_oneof:"Data"`
}

func (x *ClientMsg) Reset() {
	*x = ClientMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMsg) ProtoMessage() {}

func (x *ClientMsg) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMsg.ProtoReflect.Descriptor instead.
func (*ClientMsg) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *ClientMsg) GetAction() ClientActionType {
	if x != nil {
		return x.Action
	}
	return ClientActionType_LOGIN
}

func (m *ClientMsg) GetData() isClientMsg_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ClientMsg) GetLoginMsgData() *LoginMsgData {
	if x, ok := x.GetData().(*ClientMsg_LoginMsgData); ok {
		return x.LoginMsgData
	}
	return nil
}

func (x *ClientMsg) GetRegisterMsgData() *RegisterMsgData {
	if x, ok := x.GetData().(*ClientMsg_RegisterMsgData); ok {
		return x.RegisterMsgData
	}
	return nil
}

func (x *ClientMsg) GetLogoutUserId() string {
	if x, ok := x.GetData().(*ClientMsg_LogoutUserId); ok {
		return x.LogoutUserId
	}
	return ""
}

type isClientMsg_Data interface {
	isClientMsg_Data()
}

type ClientMsg_LoginMsgData struct {
	LoginMsgData *LoginMsgData `protobuf:"bytes,2,opt,name=login_msg_data,json=loginMsgData,proto3,oneof"` // 登录消息
}

type ClientMsg_RegisterMsgData struct {
	RegisterMsgData *RegisterMsgData `protobuf:"bytes,3,opt,name=register_msg_data,json=registerMsgData,proto3,oneof"` // 注册消息
}

type ClientMsg_LogoutUserId struct {
	LogoutUserId string `protobuf:"bytes,4,opt,name=logout_user_id,json=logoutUserId,proto3,oneof"` // 登出用户id
}

func (*ClientMsg_LoginMsgData) isClientMsg_Data() {}

func (*ClientMsg_RegisterMsgData) isClientMsg_Data() {}

func (*ClientMsg_LogoutUserId) isClientMsg_Data() {}

type LoginMsgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAccount string `protobuf:"bytes,1,opt,name=UserAccount,proto3" json:"UserAccount,omitempty"`
	UserPwd     string `protobuf:"bytes,2,opt,name=UserPwd,proto3" json:"UserPwd,omitempty"`
}

func (x *LoginMsgData) Reset() {
	*x = LoginMsgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginMsgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginMsgData) ProtoMessage() {}

func (x *LoginMsgData) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginMsgData.ProtoReflect.Descriptor instead.
func (*LoginMsgData) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

func (x *LoginMsgData) GetUserAccount() string {
	if x != nil {
		return x.UserAccount
	}
	return ""
}

func (x *LoginMsgData) GetUserPwd() string {
	if x != nil {
		return x.UserPwd
	}
	return ""
}

type RegisterMsgData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAccount string `protobuf:"bytes,1,opt,name=UserAccount,proto3" json:"UserAccount,omitempty"`
	UserPwd     string `protobuf:"bytes,2,opt,name=UserPwd,proto3" json:"UserPwd,omitempty"`
	NickName    string `protobuf:"bytes,3,opt,name=NickName,proto3" json:"NickName,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *RegisterMsgData) Reset() {
	*x = RegisterMsgData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterMsgData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterMsgData) ProtoMessage() {}

func (x *RegisterMsgData) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterMsgData.ProtoReflect.Descriptor instead.
func (*RegisterMsgData) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterMsgData) GetUserAccount() string {
	if x != nil {
		return x.UserAccount
	}
	return ""
}

func (x *RegisterMsgData) GetUserPwd() string {
	if x != nil {
		return x.UserPwd
	}
	return ""
}

func (x *RegisterMsgData) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *RegisterMsgData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd,
	0x01, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x29, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00,
	0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3e,
	0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x26,
	0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x4a,
	0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x50, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x50, 0x77, 0x64, 0x22, 0x7f, 0x0a, 0x0f, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x50, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x50, 0x77, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x2a, 0x37, 0x0a, 0x10, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x09, 0x0a, 0x05, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x47, 0x4f,
	0x55, 0x54, 0x10, 0x02, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_client_proto_goTypes = []interface{}{
	(ClientActionType)(0),   // 0: ClientActionType
	(*ClientMsg)(nil),       // 1: ClientMsg
	(*LoginMsgData)(nil),    // 2: LoginMsgData
	(*RegisterMsgData)(nil), // 3: RegisterMsgData
}
var file_client_proto_depIdxs = []int32{
	0, // 0: ClientMsg.action:type_name -> ClientActionType
	2, // 1: ClientMsg.login_msg_data:type_name -> LoginMsgData
	3, // 2: ClientMsg.register_msg_data:type_name -> RegisterMsgData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMsg); i {
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
		file_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginMsgData); i {
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
		file_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterMsgData); i {
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
	file_client_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ClientMsg_LoginMsgData)(nil),
		(*ClientMsg_RegisterMsgData)(nil),
		(*ClientMsg_LogoutUserId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		EnumInfos:         file_client_proto_enumTypes,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}
