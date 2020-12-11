// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: box/account.proto

package eboxpb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AcountType int32

const (
	// 占位 为了和数据库的常量一致
	AcountType_AccountTypeNone   AcountType = 0
	AcountType_AccountTypeAdmin  AcountType = 1
	AcountType_AccountTypeGuest  AcountType = 2
	AcountType_AccountTypeMember AcountType = 3
)

// Enum value maps for AcountType.
var (
	AcountType_name = map[int32]string{
		0: "AccountTypeNone",
		1: "AccountTypeAdmin",
		2: "AccountTypeGuest",
		3: "AccountTypeMember",
	}
	AcountType_value = map[string]int32{
		"AccountTypeNone":   0,
		"AccountTypeAdmin":  1,
		"AccountTypeGuest":  2,
		"AccountTypeMember": 3,
	}
)

func (x AcountType) Enum() *AcountType {
	p := new(AcountType)
	*p = x
	return p
}

func (x AcountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AcountType) Descriptor() protoreflect.EnumDescriptor {
	return file_box_account_proto_enumTypes[0].Descriptor()
}

func (AcountType) Type() protoreflect.EnumType {
	return &file_box_account_proto_enumTypes[0]
}

func (x AcountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AcountType.Descriptor instead.
func (AcountType) EnumDescriptor() ([]byte, []int) {
	return file_box_account_proto_rawDescGZIP(), []int{0}
}

type SetBasicInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname  string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	AvatarKey string `protobuf:"bytes,2,opt,name=avatar_key,json=avatarKey,proto3" json:"avatar_key,omitempty"`
}

func (x *SetBasicInfoRequest) Reset() {
	*x = SetBasicInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBasicInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBasicInfoRequest) ProtoMessage() {}

func (x *SetBasicInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBasicInfoRequest.ProtoReflect.Descriptor instead.
func (*SetBasicInfoRequest) Descriptor() ([]byte, []int) {
	return file_box_account_proto_rawDescGZIP(), []int{0}
}

func (x *SetBasicInfoRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SetBasicInfoRequest) GetAvatarKey() string {
	if x != nil {
		return x.AvatarKey
	}
	return ""
}

type AuthenticationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//用户ID
	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	//向用户控制器申请的token
	BoxToken string `protobuf:"bytes,2,opt,name=box_token,json=boxToken,proto3" json:"box_token,omitempty"`
	//如果不是连接的自己盒子，需要传递此参数，
	ConversationId string `protobuf:"bytes,3,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
}

func (x *AuthenticationRequest) Reset() {
	*x = AuthenticationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationRequest) ProtoMessage() {}

func (x *AuthenticationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return file_box_account_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticationRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AuthenticationRequest) GetBoxToken() string {
	if x != nil {
		return x.BoxToken
	}
	return ""
}

func (x *AuthenticationRequest) GetConversationId() string {
	if x != nil {
		return x.ConversationId
	}
	return ""
}

type AuthenticationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAdmin bool `protobuf:"varint,1,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
}

func (x *AuthenticationResponse) Reset() {
	*x = AuthenticationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationResponse) ProtoMessage() {}

func (x *AuthenticationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return file_box_account_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticationResponse) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type GetProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_box_account_proto_rawDescGZIP(), []int{3}
}

func (x *GetProfileResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type DelShareAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// string box_token = 2; // 盒子调动控制器接口的token
	VerifyCode string `protobuf:"bytes,3,opt,name=verify_code,json=verifyCode,proto3" json:"verify_code,omitempty"`
}

func (x *DelShareAccountRequest) Reset() {
	*x = DelShareAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelShareAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelShareAccountRequest) ProtoMessage() {}

func (x *DelShareAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelShareAccountRequest.ProtoReflect.Descriptor instead.
func (*DelShareAccountRequest) Descriptor() ([]byte, []int) {
	return file_box_account_proto_rawDescGZIP(), []int{4}
}

func (x *DelShareAccountRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DelShareAccountRequest) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

type DelShareAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelShareAccountResponse) Reset() {
	*x = DelShareAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelShareAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelShareAccountResponse) ProtoMessage() {}

func (x *DelShareAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelShareAccountResponse.ProtoReflect.Descriptor instead.
func (*DelShareAccountResponse) Descriptor() ([]byte, []int) {
	return file_box_account_proto_rawDescGZIP(), []int{5}
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         int64  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname    string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	AvatarKey   string `protobuf:"bytes,3,opt,name=avatar_key,json=avatarKey,proto3" json:"avatar_key,omitempty"`
	TotalInByte int64  `protobuf:"varint,4,opt,name=total_in_byte,json=totalInByte,proto3" json:"total_in_byte,omitempty"`
	UseIntByte  int64  `protobuf:"varint,5,opt,name=use_int_byte,json=useIntByte,proto3" json:"use_int_byte,omitempty"`
	// 用户类型，1 管理员；2 共享盒子的客人；3 只有群聊权限的用户
	AcountType int32 `protobuf:"varint,6,opt,name=AcountType,proto3" json:"AcountType,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_box_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_box_account_proto_rawDescGZIP(), []int{6}
}

func (x *Account) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Account) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Account) GetAvatarKey() string {
	if x != nil {
		return x.AvatarKey
	}
	return ""
}

func (x *Account) GetTotalInByte() int64 {
	if x != nil {
		return x.TotalInByte
	}
	return 0
}

func (x *Account) GetUseIntByte() int64 {
	if x != nil {
		return x.UseIntByte
	}
	return 0
}

func (x *Account) GetAcountType() int32 {
	if x != nil {
		return x.AcountType
	}
	return 0
}

type ListShareAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShareAccounts []*Account `protobuf:"bytes,1,rep,name=share_accounts,json=shareAccounts,proto3" json:"share_accounts,omitempty"`
}

func (x *ListShareAccountResponse) Reset() {
	*x = ListShareAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShareAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShareAccountResponse) ProtoMessage() {}

func (x *ListShareAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShareAccountResponse.ProtoReflect.Descriptor instead.
func (*ListShareAccountResponse) Descriptor() ([]byte, []int) {
	return file_box_account_proto_rawDescGZIP(), []int{7}
}

func (x *ListShareAccountResponse) GetShareAccounts() []*Account {
	if x != nil {
		return x.ShareAccounts
	}
	return nil
}

type QuitShareAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string box_token = 1; // 盒子调动控制器接口的token
	VerifyCode string `protobuf:"bytes,2,opt,name=verify_code,json=verifyCode,proto3" json:"verify_code,omitempty"` // 解绑验证码，需要调用验证码
}

func (x *QuitShareAccountRequest) Reset() {
	*x = QuitShareAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuitShareAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuitShareAccountRequest) ProtoMessage() {}

func (x *QuitShareAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuitShareAccountRequest.ProtoReflect.Descriptor instead.
func (*QuitShareAccountRequest) Descriptor() ([]byte, []int) {
	return file_box_account_proto_rawDescGZIP(), []int{8}
}

func (x *QuitShareAccountRequest) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

type UnbindAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string box_token = 1; // 盒子调动控制器接口的token
	VerifyCode string `protobuf:"bytes,2,opt,name=verify_code,json=verifyCode,proto3" json:"verify_code,omitempty"`
}

func (x *UnbindAccountRequest) Reset() {
	*x = UnbindAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnbindAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbindAccountRequest) ProtoMessage() {}

func (x *UnbindAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnbindAccountRequest.ProtoReflect.Descriptor instead.
func (*UnbindAccountRequest) Descriptor() ([]byte, []int) {
	return file_box_account_proto_rawDescGZIP(), []int{9}
}

func (x *UnbindAccountRequest) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

var File_box_account_proto protoreflect.FileDescriptor

var file_box_account_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x6f, 0x78, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x74, 0x73, 0x6d, 0x65, 0x2e, 0x65, 0x62, 0x6f, 0x78, 0x22,
	0x50, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4b, 0x65,
	0x79, 0x22, 0x6f, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x6f, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x6f, 0x78, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x32, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x43, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x65, 0x74, 0x73, 0x6d, 0x65, 0x2e, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4b, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0d,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x6e, 0x42, 0x79, 0x74, 0x65,
	0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x42, 0x79,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x41, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x56, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x0e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x74, 0x73, 0x6d, 0x65, 0x2e, 0x65,
	0x62, 0x6f, 0x78, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0d, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x3a, 0x0a, 0x17, 0x51, 0x75,
	0x69, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x37, 0x0a, 0x14, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x2a,
	0x64, 0x0a, 0x0a, 0x41, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a,
	0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x6f, 0x6e, 0x65,
	0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x47, 0x75, 0x65, 0x73, 0x74, 0x10, 0x02, 0x12, 0x15,
	0x0a, 0x11, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x10, 0x03, 0x42, 0x61, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x74, 0x73,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x6f, 0x78, 0x42,
	0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x65, 0x74, 0x73, 0x75, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x2f, 0x65, 0x62, 0x6f, 0x78, 0x2f, 0x65, 0x62, 0x6f, 0x78, 0x2d, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x62, 0x6f, 0x78, 0x70, 0x62, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x05, 0x45, 0x50, 0x42, 0x6f, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_box_account_proto_rawDescOnce sync.Once
	file_box_account_proto_rawDescData = file_box_account_proto_rawDesc
)

func file_box_account_proto_rawDescGZIP() []byte {
	file_box_account_proto_rawDescOnce.Do(func() {
		file_box_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_box_account_proto_rawDescData)
	})
	return file_box_account_proto_rawDescData
}

var file_box_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_box_account_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_box_account_proto_goTypes = []interface{}{
	(AcountType)(0),                  // 0: etsme.ebox.AcountType
	(*SetBasicInfoRequest)(nil),      // 1: etsme.ebox.SetBasicInfoRequest
	(*AuthenticationRequest)(nil),    // 2: etsme.ebox.AuthenticationRequest
	(*AuthenticationResponse)(nil),   // 3: etsme.ebox.AuthenticationResponse
	(*GetProfileResponse)(nil),       // 4: etsme.ebox.GetProfileResponse
	(*DelShareAccountRequest)(nil),   // 5: etsme.ebox.DelShareAccountRequest
	(*DelShareAccountResponse)(nil),  // 6: etsme.ebox.DelShareAccountResponse
	(*Account)(nil),                  // 7: etsme.ebox.Account
	(*ListShareAccountResponse)(nil), // 8: etsme.ebox.ListShareAccountResponse
	(*QuitShareAccountRequest)(nil),  // 9: etsme.ebox.QuitShareAccountRequest
	(*UnbindAccountRequest)(nil),     // 10: etsme.ebox.UnbindAccountRequest
}
var file_box_account_proto_depIdxs = []int32{
	7, // 0: etsme.ebox.GetProfileResponse.account:type_name -> etsme.ebox.Account
	7, // 1: etsme.ebox.ListShareAccountResponse.share_accounts:type_name -> etsme.ebox.Account
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_box_account_proto_init() }
func file_box_account_proto_init() {
	if File_box_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_box_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetBasicInfoRequest); i {
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
		file_box_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationRequest); i {
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
		file_box_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationResponse); i {
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
		file_box_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileResponse); i {
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
		file_box_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelShareAccountRequest); i {
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
		file_box_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelShareAccountResponse); i {
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
		file_box_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_box_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShareAccountResponse); i {
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
		file_box_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuitShareAccountRequest); i {
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
		file_box_account_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnbindAccountRequest); i {
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
			RawDescriptor: file_box_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_box_account_proto_goTypes,
		DependencyIndexes: file_box_account_proto_depIdxs,
		EnumInfos:         file_box_account_proto_enumTypes,
		MessageInfos:      file_box_account_proto_msgTypes,
	}.Build()
	File_box_account_proto = out.File
	file_box_account_proto_rawDesc = nil
	file_box_account_proto_goTypes = nil
	file_box_account_proto_depIdxs = nil
}
