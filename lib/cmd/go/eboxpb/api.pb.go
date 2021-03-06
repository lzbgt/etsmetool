// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: box/api.proto

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

type ApiOperation int32

const (
	ApiOperation_AuthenticationOp ApiOperation = 0
	// 创建群
	ApiOperation_CreateConversationOp ApiOperation = 4
	// 加入群
	// JoinConversationOp = 5;
	// 成员主动退出会话
	ApiOperation_LeaveConversationOp ApiOperation = 6
	// 管理员解散会话
	ApiOperation_DisbandConversationOp ApiOperation = 8
	// 群列表
	ApiOperation_ListConversationOp ApiOperation = 7
	// 设置头像
	// SetBasicInfoOp=15;
	// 获取会话成员
	ApiOperation_ListConversationMembersOp ApiOperation = 16
	// 获取会话历史消息
	ApiOperation_ListConversationMessageOp ApiOperation = 17
	// 综合group和message history搜索 ConversationRelatedSearchRequest ConversationRelatedSearchResponse
	ApiOperation_ConversationRelatedSearchOp ApiOperation = 21
	// 获取绑定码
	ApiOperation_GetBindCodeOp ApiOperation = 22
	// 获取入群邀请码
	ApiOperation_GetConversationInviteCodeOp ApiOperation = 23
	// 根据关键词搜索历史消息，返回的是会话列表，每一个会话命中多少消息 MessageSearchRequest MessageSearchResponse
	ApiOperation_MessageSearchOp ApiOperation = 26
	//个人中心-容量列表
	//Request:  google.protobuf.Empty
	//Response: ListCapacityResponse
	ApiOperation_ListCapacityOp ApiOperation = 30
	// 置顶会话
	ApiOperation_TopConversationOp ApiOperation = 32
	// 更新成员备注，针对当前用户有效
	ApiOperation_UpdateMemberNoteOp ApiOperation = 33
	// 管理员删除共享目录，取消与TA共享设备
	ApiOperation_DelShareAccountOp ApiOperation = 34
	// 当前盒子的管理员，获取共享的用户列表
	ApiOperation_ListShareAccountOp ApiOperation = 35
	// 用户主动退出共享
	ApiOperation_QuitShareAccountOp ApiOperation = 36
	// 获取设备的信息
	ApiOperation_GetDeviceInfoOp ApiOperation = 37
	// 设备重启
	ApiOperation_RebootOp ApiOperation = 38
	// 恢复出厂设置
	ApiOperation_FactoryResetOp ApiOperation = 39
	// 解绑设备
	ApiOperation_UnbindAccountOp ApiOperation = 40
	// 会话搜索 ConversationSearchRequest & ConversationSearchResponse
	ApiOperation_ConversationSearchOp ApiOperation = 50
	// 通用文件搜索，包括图片，视频，音频，文件，其他文件 CommonFileSearchRequest & CommonFileSearchResponse
	ApiOperation_CommonFileSearchOp ApiOperation = 57
	// GetMemberProfileRequest GetMemberProfileResponse 在group内，或者member的profile，注意，只能在群内使用
	ApiOperation_GetMemberProfileOp ApiOperation = 58
	// ConversationCalendarSearchOp 查询日历
	ApiOperation_ConversationCalendarSearchOp ApiOperation = 59
	// GlobalSearchResponseOp 查询日历
	ApiOperation_GlobalSearchOp ApiOperation = 60
	// SendMessageOp 发送消息，注意双向接口，服务端向客户端发送消息也是这个操作。
	ApiOperation_SendMessageOp ApiOperation = 61
	// 已都消息
	ApiOperation_ReadMessageOp ApiOperation = 62
	//更新自己在群里的资料
	ApiOperation_UpdateProfileInConversationOp ApiOperation = 63
	//更新会话信息，只有群主有权限
	ApiOperation_UpdateConversationOp ApiOperation = 64
	//时间线接口
	ApiOperation_TimelineOp ApiOperation = 65
	//设备运行状态
	ApiOperation_MonitorOp ApiOperation = 66
	//消息撤回  CancelSendMessageRequest
	ApiOperation_CancelSendMessageOp ApiOperation = 67
	// 更新群名称，只有群主有权限
	ApiOperation_UpdateConversationNameOp ApiOperation = 68
	// 查询任务ID列表
	ApiOperation_GetInterBoxFileTransferTaskIDsOp ApiOperation = 70
	// 文件转发功能
	ApiOperation_InterBoxFileTransferOp ApiOperation = 71
	// App请求Ebox任务状态
	ApiOperation_DurableTaskStatusRequestOp ApiOperation = 72
	// App取消文件转发任务
	ApiOperation_CancelDurableTaskRequestOp ApiOperation = 73
	// 升级系统
	ApiOperation_UpgradeOp ApiOperation = 74
	// 获取升级状态
	ApiOperation_GetUpgradeStatusOp ApiOperation = 75
	// 文件上传完成（不通过群上传文件，例如自动备份中的文件上传完成以后需要调用该接口，才能再盒子上搜索到文件）
	ApiOperation_UploadCompleteOp ApiOperation = 76
	// 群主踢出群成员
	ApiOperation_KickOutConversationMemberOp ApiOperation = 77
	// 查询附件消息
	ApiOperation_ListMessageAttachmentOp ApiOperation = 78
)

// Enum value maps for ApiOperation.
var (
	ApiOperation_name = map[int32]string{
		0:  "AuthenticationOp",
		4:  "CreateConversationOp",
		6:  "LeaveConversationOp",
		8:  "DisbandConversationOp",
		7:  "ListConversationOp",
		16: "ListConversationMembersOp",
		17: "ListConversationMessageOp",
		21: "ConversationRelatedSearchOp",
		22: "GetBindCodeOp",
		23: "GetConversationInviteCodeOp",
		26: "MessageSearchOp",
		30: "ListCapacityOp",
		32: "TopConversationOp",
		33: "UpdateMemberNoteOp",
		34: "DelShareAccountOp",
		35: "ListShareAccountOp",
		36: "QuitShareAccountOp",
		37: "GetDeviceInfoOp",
		38: "RebootOp",
		39: "FactoryResetOp",
		40: "UnbindAccountOp",
		50: "ConversationSearchOp",
		57: "CommonFileSearchOp",
		58: "GetMemberProfileOp",
		59: "ConversationCalendarSearchOp",
		60: "GlobalSearchOp",
		61: "SendMessageOp",
		62: "ReadMessageOp",
		63: "UpdateProfileInConversationOp",
		64: "UpdateConversationOp",
		65: "TimelineOp",
		66: "MonitorOp",
		67: "CancelSendMessageOp",
		68: "UpdateConversationNameOp",
		70: "GetInterBoxFileTransferTaskIDsOp",
		71: "InterBoxFileTransferOp",
		72: "DurableTaskStatusRequestOp",
		73: "CancelDurableTaskRequestOp",
		74: "UpgradeOp",
		75: "GetUpgradeStatusOp",
		76: "UploadCompleteOp",
		77: "KickOutConversationMemberOp",
		78: "ListMessageAttachmentOp",
	}
	ApiOperation_value = map[string]int32{
		"AuthenticationOp":                 0,
		"CreateConversationOp":             4,
		"LeaveConversationOp":              6,
		"DisbandConversationOp":            8,
		"ListConversationOp":               7,
		"ListConversationMembersOp":        16,
		"ListConversationMessageOp":        17,
		"ConversationRelatedSearchOp":      21,
		"GetBindCodeOp":                    22,
		"GetConversationInviteCodeOp":      23,
		"MessageSearchOp":                  26,
		"ListCapacityOp":                   30,
		"TopConversationOp":                32,
		"UpdateMemberNoteOp":               33,
		"DelShareAccountOp":                34,
		"ListShareAccountOp":               35,
		"QuitShareAccountOp":               36,
		"GetDeviceInfoOp":                  37,
		"RebootOp":                         38,
		"FactoryResetOp":                   39,
		"UnbindAccountOp":                  40,
		"ConversationSearchOp":             50,
		"CommonFileSearchOp":               57,
		"GetMemberProfileOp":               58,
		"ConversationCalendarSearchOp":     59,
		"GlobalSearchOp":                   60,
		"SendMessageOp":                    61,
		"ReadMessageOp":                    62,
		"UpdateProfileInConversationOp":    63,
		"UpdateConversationOp":             64,
		"TimelineOp":                       65,
		"MonitorOp":                        66,
		"CancelSendMessageOp":              67,
		"UpdateConversationNameOp":         68,
		"GetInterBoxFileTransferTaskIDsOp": 70,
		"InterBoxFileTransferOp":           71,
		"DurableTaskStatusRequestOp":       72,
		"CancelDurableTaskRequestOp":       73,
		"UpgradeOp":                        74,
		"GetUpgradeStatusOp":               75,
		"UploadCompleteOp":                 76,
		"KickOutConversationMemberOp":      77,
		"ListMessageAttachmentOp":          78,
	}
)

func (x ApiOperation) Enum() *ApiOperation {
	p := new(ApiOperation)
	*p = x
	return p
}

func (x ApiOperation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiOperation) Descriptor() protoreflect.EnumDescriptor {
	return file_box_api_proto_enumTypes[0].Descriptor()
}

func (ApiOperation) Type() protoreflect.EnumType {
	return &file_box_api_proto_enumTypes[0]
}

func (x ApiOperation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiOperation.Descriptor instead.
func (ApiOperation) EnumDescriptor() ([]byte, []int) {
	return file_box_api_proto_rawDescGZIP(), []int{0}
}

type ApiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation     ApiOperation `protobuf:"varint,1,opt,name=operation,proto3,enum=etsme.ebox.ApiOperation" json:"operation,omitempty"`
	Content       []byte       `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	ServerVersion int32        `protobuf:"varint,3,opt,name=serverVersion,proto3" json:"serverVersion,omitempty"`
	From          int64        `protobuf:"varint,4,opt,name=from,proto3" json:"from,omitempty"` // message source
	To            int64        `protobuf:"varint,5,opt,name=to,proto3" json:"to,omitempty"`     // next hop: 0 - usage defined by application; otherwise to one specific
}

func (x *ApiRequest) Reset() {
	*x = ApiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiRequest) ProtoMessage() {}

func (x *ApiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_box_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiRequest.ProtoReflect.Descriptor instead.
func (*ApiRequest) Descriptor() ([]byte, []int) {
	return file_box_api_proto_rawDescGZIP(), []int{0}
}

func (x *ApiRequest) GetOperation() ApiOperation {
	if x != nil {
		return x.Operation
	}
	return ApiOperation_AuthenticationOp
}

func (x *ApiRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ApiRequest) GetServerVersion() int32 {
	if x != nil {
		return x.ServerVersion
	}
	return 0
}

func (x *ApiRequest) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *ApiRequest) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

type ApiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation ApiOperation `protobuf:"varint,1,opt,name=operation,proto3,enum=etsme.ebox.ApiOperation" json:"operation,omitempty"`
	Content   []byte       `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Message   string       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Code      int32        `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	ErrDetail string       `protobuf:"bytes,5,opt,name=errDetail,proto3" json:"errDetail,omitempty"`
}

func (x *ApiResponse) Reset() {
	*x = ApiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_box_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponse) ProtoMessage() {}

func (x *ApiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_box_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponse.ProtoReflect.Descriptor instead.
func (*ApiResponse) Descriptor() ([]byte, []int) {
	return file_box_api_proto_rawDescGZIP(), []int{1}
}

func (x *ApiResponse) GetOperation() ApiOperation {
	if x != nil {
		return x.Operation
	}
	return ApiOperation_AuthenticationOp
}

func (x *ApiResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ApiResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApiResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ApiResponse) GetErrDetail() string {
	if x != nil {
		return x.ErrDetail
	}
	return ""
}

var File_box_api_proto protoreflect.FileDescriptor

var file_box_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x65, 0x74, 0x73, 0x6d, 0x65, 0x2e, 0x65, 0x62, 0x6f, 0x78, 0x22, 0xa8, 0x01, 0x0a, 0x0a,
	0x41, 0x70, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x65, 0x74, 0x73, 0x6d, 0x65, 0x2e, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x41, 0x70, 0x69, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x74, 0x73, 0x6d,
	0x65, 0x2e, 0x65, 0x62, 0x6f, 0x78, 0x2e, 0x41, 0x70, 0x69, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2a, 0xbe, 0x08, 0x0a, 0x0c, 0x41, 0x70, 0x69, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x10, 0x06, 0x12, 0x19,
	0x0a, 0x15, 0x44, 0x69, 0x73, 0x62, 0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x10,
	0x07, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x4f, 0x70, 0x10, 0x10,
	0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x10, 0x11, 0x12,
	0x1f, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x10, 0x15,
	0x12, 0x11, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x4f,
	0x70, 0x10, 0x16, 0x12, 0x1f, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x4f, 0x70, 0x10, 0x17, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x10, 0x1a, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x4f, 0x70, 0x10, 0x1e, 0x12, 0x15, 0x0a,
	0x11, 0x54, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x10, 0x20, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x65, 0x4f, 0x70, 0x10, 0x21, 0x12, 0x15, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x53, 0x68, 0x61, 0x72, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f,
	0x70, 0x10, 0x22, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x70, 0x10, 0x23, 0x12, 0x16, 0x0a, 0x12, 0x51,
	0x75, 0x69, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f,
	0x70, 0x10, 0x24, 0x12, 0x13, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x4f, 0x70, 0x10, 0x25, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x62, 0x6f,
	0x6f, 0x74, 0x4f, 0x70, 0x10, 0x26, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4f, 0x70, 0x10, 0x27, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x6e,
	0x62, 0x69, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x70, 0x10, 0x28, 0x12,
	0x18, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x10, 0x32, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x10,
	0x39, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x10, 0x3a, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x10, 0x3b, 0x12, 0x12, 0x0a, 0x0e, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x10, 0x3c, 0x12,
	0x11, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x10, 0x3d, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x10, 0x3e, 0x12, 0x21, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x10, 0x3f, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x10, 0x40, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4f, 0x70,
	0x10, 0x41, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x10,
	0x42, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x10, 0x43, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x70, 0x10, 0x44, 0x12, 0x24, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x73, 0x4f, 0x70, 0x10, 0x46, 0x12, 0x1a,
	0x0a, 0x16, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x6f, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4f, 0x70, 0x10, 0x47, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x75,
	0x72, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x70, 0x10, 0x48, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x70, 0x10, 0x49, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x4f, 0x70, 0x10, 0x4a, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x70, 0x10,
	0x4b, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x70, 0x10, 0x4c, 0x12, 0x1f, 0x0a, 0x1b, 0x4b, 0x69, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x70, 0x10, 0x4d, 0x12, 0x1b, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x4f, 0x70, 0x10, 0x4e, 0x42, 0x5d, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x74, 0x73,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x6f, 0x78, 0x42,
	0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x2e, 0x65, 0x74, 0x73, 0x75, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x65, 0x62,
	0x6f, 0x78, 0x2f, 0x65, 0x62, 0x6f, 0x78, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x62, 0x6f, 0x78, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x05, 0x45,
	0x50, 0x42, 0x6f, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_box_api_proto_rawDescOnce sync.Once
	file_box_api_proto_rawDescData = file_box_api_proto_rawDesc
)

func file_box_api_proto_rawDescGZIP() []byte {
	file_box_api_proto_rawDescOnce.Do(func() {
		file_box_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_box_api_proto_rawDescData)
	})
	return file_box_api_proto_rawDescData
}

var file_box_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_box_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_box_api_proto_goTypes = []interface{}{
	(ApiOperation)(0),   // 0: etsme.ebox.ApiOperation
	(*ApiRequest)(nil),  // 1: etsme.ebox.ApiRequest
	(*ApiResponse)(nil), // 2: etsme.ebox.ApiResponse
}
var file_box_api_proto_depIdxs = []int32{
	0, // 0: etsme.ebox.ApiRequest.operation:type_name -> etsme.ebox.ApiOperation
	0, // 1: etsme.ebox.ApiResponse.operation:type_name -> etsme.ebox.ApiOperation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_box_api_proto_init() }
func file_box_api_proto_init() {
	if File_box_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_box_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiRequest); i {
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
		file_box_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResponse); i {
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
			RawDescriptor: file_box_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_box_api_proto_goTypes,
		DependencyIndexes: file_box_api_proto_depIdxs,
		EnumInfos:         file_box_api_proto_enumTypes,
		MessageInfos:      file_box_api_proto_msgTypes,
	}.Build()
	File_box_api_proto = out.File
	file_box_api_proto_rawDesc = nil
	file_box_api_proto_goTypes = nil
	file_box_api_proto_depIdxs = nil
}
