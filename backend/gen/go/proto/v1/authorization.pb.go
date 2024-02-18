// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/v1/authorization.proto

package pegismv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signature of message
	Sig string `protobuf:"bytes,1,opt,name=sig,proto3" json:"sig,omitempty"`
	// message thats been signed. usually it is UnsignedAuthMessage.content
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AuthArgs) Reset() {
	*x = AuthArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_authorization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthArgs) ProtoMessage() {}

func (x *AuthArgs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_authorization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthArgs.ProtoReflect.Descriptor instead.
func (*AuthArgs) Descriptor() ([]byte, []int) {
	return file_proto_v1_authorization_proto_rawDescGZIP(), []int{0}
}

func (x *AuthArgs) GetSig() string {
	if x != nil {
		return x.Sig
	}
	return ""
}

func (x *AuthArgs) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// UnsignedAuthMessage is a message that includes unsgined content and deadline.
type UnsignedAuthMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// deadline of content
	Deadline *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (x *UnsignedAuthMessage) Reset() {
	*x = UnsignedAuthMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_authorization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsignedAuthMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsignedAuthMessage) ProtoMessage() {}

func (x *UnsignedAuthMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_authorization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsignedAuthMessage.ProtoReflect.Descriptor instead.
func (*UnsignedAuthMessage) Descriptor() ([]byte, []int) {
	return file_proto_v1_authorization_proto_rawDescGZIP(), []int{1}
}

func (x *UnsignedAuthMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UnsignedAuthMessage) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

// Request for authenticating a user, includes authentication arguments.
type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Args *AuthArgs `protobuf:"bytes,1,opt,name=args,proto3" json:"args,omitempty"` // Authentication arguments for the request.
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_authorization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_authorization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_authorization_proto_rawDescGZIP(), []int{2}
}

func (x *AuthRequest) GetArgs() *AuthArgs {
	if x != nil {
		return x.Args
	}
	return nil
}

// Response for authenticating a user, returns AuthorizationDataContainer.
type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *AuthResponse_AuthorizationDataContainer `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` // Authorization data container.
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_authorization_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_authorization_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_authorization_proto_rawDescGZIP(), []int{3}
}

func (x *AuthResponse) GetData() *AuthResponse_AuthorizationDataContainer {
	if x != nil {
		return x.Data
	}
	return nil
}

// Request for obtaining an unsigned authentication message.
type AuthMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ethereum address for which the unsigned auth message is requested.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AuthMessageRequest) Reset() {
	*x = AuthMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_authorization_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthMessageRequest) ProtoMessage() {}

func (x *AuthMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_authorization_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthMessageRequest.ProtoReflect.Descriptor instead.
func (*AuthMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_authorization_proto_rawDescGZIP(), []int{4}
}

func (x *AuthMessageRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Response for obtaining an unsigned authentication message, returns UnsignedAuthMessage.
type AuthMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unsigned authentication message.
	Message *UnsignedAuthMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AuthMessageResponse) Reset() {
	*x = AuthMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_authorization_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthMessageResponse) ProtoMessage() {}

func (x *AuthMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_authorization_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthMessageResponse.ProtoReflect.Descriptor instead.
func (*AuthMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_authorization_proto_rawDescGZIP(), []int{5}
}

func (x *AuthMessageResponse) GetMessage() *UnsignedAuthMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

// Request for obtaining an unsigned authentication message.
type RefreshMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshMessageRequest) Reset() {
	*x = RefreshMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_authorization_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshMessageRequest) ProtoMessage() {}

func (x *RefreshMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_authorization_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshMessageRequest.ProtoReflect.Descriptor instead.
func (*RefreshMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_authorization_proto_rawDescGZIP(), []int{6}
}

func (x *RefreshMessageRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

// Request for obtaining an unsigned authentication message.
type RefreshMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The access token itself.
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Timestamp for access token expiration.
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *RefreshMessageResponse) Reset() {
	*x = RefreshMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_authorization_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshMessageResponse) ProtoMessage() {}

func (x *RefreshMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_authorization_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshMessageResponse.ProtoReflect.Descriptor instead.
func (*RefreshMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_authorization_proto_rawDescGZIP(), []int{7}
}

func (x *RefreshMessageResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RefreshMessageResponse) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

// Container for authorization data.
type AuthResponse_AuthorizationDataContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Access token for authentication.
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// Refresh token for refreshing authentication.
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	// Timestamp for access token expiration.
	AccessTokenExpiresAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=access_token_expires_at,json=accessTokenExpiresAt,proto3" json:"access_token_expires_at,omitempty"`
	// Timestamp for refresh token expiration.
	RefreshTokenExpiresAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=refresh_token_expires_at,json=refreshTokenExpiresAt,proto3" json:"refresh_token_expires_at,omitempty"`
}

func (x *AuthResponse_AuthorizationDataContainer) Reset() {
	*x = AuthResponse_AuthorizationDataContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_authorization_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse_AuthorizationDataContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse_AuthorizationDataContainer) ProtoMessage() {}

func (x *AuthResponse_AuthorizationDataContainer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_authorization_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse_AuthorizationDataContainer.ProtoReflect.Descriptor instead.
func (*AuthResponse_AuthorizationDataContainer) Descriptor() ([]byte, []int) {
	return file_proto_v1_authorization_proto_rawDescGZIP(), []int{3, 0}
}

func (x *AuthResponse_AuthorizationDataContainer) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthResponse_AuthorizationDataContainer) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AuthResponse_AuthorizationDataContainer) GetAccessTokenExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AccessTokenExpiresAt
	}
	return nil
}

func (x *AuthResponse_AuthorizationDataContainer) GetRefreshTokenExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RefreshTokenExpiresAt
	}
	return nil
}

var File_proto_v1_authorization_proto protoreflect.FileDescriptor

var file_proto_v1_authorization_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x70, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x41, 0x75,
	0x74, 0x68, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x67, 0x0a, 0x13, 0x55, 0x6e,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x64,
	0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x22, 0x36, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x41, 0x72, 0x67, 0x73, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0xe5, 0x02, 0x0a, 0x0c,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x65, 0x67,
	0x69, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x8c, 0x02, 0x0a, 0x1a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x51, 0x0a, 0x17, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x53,
	0x0a, 0x18, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x15, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x41, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x65,
	0x67, 0x69, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a, 0x15, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x76, 0x0a, 0x16, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x42, 0x9c, 0x01, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45,
	0x78, 0x63, 0x61, 0x2d, 0x44, 0x4b, 0x2f, 0x70, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x65,
	0x67, 0x69, 0x73, 0x6d, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x50,
	0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x50, 0x65, 0x67, 0x69, 0x73,
	0x6d, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x50, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x50,
	0x65, 0x67, 0x69, 0x73, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_v1_authorization_proto_rawDescOnce sync.Once
	file_proto_v1_authorization_proto_rawDescData = file_proto_v1_authorization_proto_rawDesc
)

func file_proto_v1_authorization_proto_rawDescGZIP() []byte {
	file_proto_v1_authorization_proto_rawDescOnce.Do(func() {
		file_proto_v1_authorization_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_authorization_proto_rawDescData)
	})
	return file_proto_v1_authorization_proto_rawDescData
}

var file_proto_v1_authorization_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_v1_authorization_proto_goTypes = []interface{}{
	(*AuthArgs)(nil),                                // 0: pegism.v1.AuthArgs
	(*UnsignedAuthMessage)(nil),                     // 1: pegism.v1.UnsignedAuthMessage
	(*AuthRequest)(nil),                             // 2: pegism.v1.AuthRequest
	(*AuthResponse)(nil),                            // 3: pegism.v1.AuthResponse
	(*AuthMessageRequest)(nil),                      // 4: pegism.v1.AuthMessageRequest
	(*AuthMessageResponse)(nil),                     // 5: pegism.v1.AuthMessageResponse
	(*RefreshMessageRequest)(nil),                   // 6: pegism.v1.RefreshMessageRequest
	(*RefreshMessageResponse)(nil),                  // 7: pegism.v1.RefreshMessageResponse
	(*AuthResponse_AuthorizationDataContainer)(nil), // 8: pegism.v1.AuthResponse.AuthorizationDataContainer
	(*timestamppb.Timestamp)(nil),                   // 9: google.protobuf.Timestamp
}
var file_proto_v1_authorization_proto_depIdxs = []int32{
	9, // 0: pegism.v1.UnsignedAuthMessage.deadline:type_name -> google.protobuf.Timestamp
	0, // 1: pegism.v1.AuthRequest.args:type_name -> pegism.v1.AuthArgs
	8, // 2: pegism.v1.AuthResponse.data:type_name -> pegism.v1.AuthResponse.AuthorizationDataContainer
	1, // 3: pegism.v1.AuthMessageResponse.message:type_name -> pegism.v1.UnsignedAuthMessage
	9, // 4: pegism.v1.RefreshMessageResponse.expires_at:type_name -> google.protobuf.Timestamp
	9, // 5: pegism.v1.AuthResponse.AuthorizationDataContainer.access_token_expires_at:type_name -> google.protobuf.Timestamp
	9, // 6: pegism.v1.AuthResponse.AuthorizationDataContainer.refresh_token_expires_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_v1_authorization_proto_init() }
func file_proto_v1_authorization_proto_init() {
	if File_proto_v1_authorization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_authorization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthArgs); i {
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
		file_proto_v1_authorization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsignedAuthMessage); i {
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
		file_proto_v1_authorization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_proto_v1_authorization_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_proto_v1_authorization_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthMessageRequest); i {
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
		file_proto_v1_authorization_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthMessageResponse); i {
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
		file_proto_v1_authorization_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshMessageRequest); i {
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
		file_proto_v1_authorization_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshMessageResponse); i {
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
		file_proto_v1_authorization_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse_AuthorizationDataContainer); i {
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
			RawDescriptor: file_proto_v1_authorization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_v1_authorization_proto_goTypes,
		DependencyIndexes: file_proto_v1_authorization_proto_depIdxs,
		MessageInfos:      file_proto_v1_authorization_proto_msgTypes,
	}.Build()
	File_proto_v1_authorization_proto = out.File
	file_proto_v1_authorization_proto_rawDesc = nil
	file_proto_v1_authorization_proto_goTypes = nil
	file_proto_v1_authorization_proto_depIdxs = nil
}