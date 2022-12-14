// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.21.5
// source: source/filters/http/proto/oidc-validation.proto

package proto

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

type ValidationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider            string                            `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Enforce             bool                              `protobuf:"varint,2,opt,name=enforce,proto3" json:"enforce,omitempty"`
	EnforceResponseCode int32                             `protobuf:"varint,3,opt,name=enforceResponseCode,proto3" json:"enforceResponseCode,omitempty"`
	AccessToken         *ValidationConfig_AccessToken     `protobuf:"bytes,4,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	UserInfo            *ValidationConfig_UserInfo        `protobuf:"bytes,5,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	TLSConfig           *ValidationConfig_FilterTLSConfig `protobuf:"bytes,6,opt,name=TLSConfig,proto3" json:"TLSConfig,omitempty"`
}

func (x *ValidationConfig) Reset() {
	*x = ValidationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_filters_http_proto_oidc_validation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationConfig) ProtoMessage() {}

func (x *ValidationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_source_filters_http_proto_oidc_validation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationConfig.ProtoReflect.Descriptor instead.
func (*ValidationConfig) Descriptor() ([]byte, []int) {
	return file_source_filters_http_proto_oidc_validation_proto_rawDescGZIP(), []int{0}
}

func (x *ValidationConfig) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *ValidationConfig) GetEnforce() bool {
	if x != nil {
		return x.Enforce
	}
	return false
}

func (x *ValidationConfig) GetEnforceResponseCode() int32 {
	if x != nil {
		return x.EnforceResponseCode
	}
	return 0
}

func (x *ValidationConfig) GetAccessToken() *ValidationConfig_AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

func (x *ValidationConfig) GetUserInfo() *ValidationConfig_UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *ValidationConfig) GetTLSConfig() *ValidationConfig_FilterTLSConfig {
	if x != nil {
		return x.TLSConfig
	}
	return nil
}

type ValidationRouteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enforce             bool     `protobuf:"varint,1,opt,name=enforce,proto3" json:"enforce,omitempty"`
	EnforceResponseCode int32    `protobuf:"varint,2,opt,name=enforceResponseCode,proto3" json:"enforceResponseCode,omitempty"`
	UserInfoClaims      []string `protobuf:"bytes,3,rep,name=userInfoClaims,proto3" json:"userInfoClaims,omitempty"`
	OverwriteClaims     bool     `protobuf:"varint,4,opt,name=overwriteClaims,proto3" json:"overwriteClaims,omitempty"`
}

func (x *ValidationRouteConfig) Reset() {
	*x = ValidationRouteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_filters_http_proto_oidc_validation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationRouteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationRouteConfig) ProtoMessage() {}

func (x *ValidationRouteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_source_filters_http_proto_oidc_validation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationRouteConfig.ProtoReflect.Descriptor instead.
func (*ValidationRouteConfig) Descriptor() ([]byte, []int) {
	return file_source_filters_http_proto_oidc_validation_proto_rawDescGZIP(), []int{1}
}

func (x *ValidationRouteConfig) GetEnforce() bool {
	if x != nil {
		return x.Enforce
	}
	return false
}

func (x *ValidationRouteConfig) GetEnforceResponseCode() int32 {
	if x != nil {
		return x.EnforceResponseCode
	}
	return 0
}

func (x *ValidationRouteConfig) GetUserInfoClaims() []string {
	if x != nil {
		return x.UserInfoClaims
	}
	return nil
}

func (x *ValidationRouteConfig) GetOverwriteClaims() bool {
	if x != nil {
		return x.OverwriteClaims
	}
	return false
}

type ValidationConfig_AccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location       LocationType `protobuf:"varint,1,opt,name=location,proto3,enum=greymatter_io.gm_proxy.source.filters.http.LocationType" json:"location,omitempty"`
	Key            string       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	MetadataFilter string       `protobuf:"bytes,3,opt,name=metadataFilter,proto3" json:"metadataFilter,omitempty"`
}

func (x *ValidationConfig_AccessToken) Reset() {
	*x = ValidationConfig_AccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_filters_http_proto_oidc_validation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationConfig_AccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationConfig_AccessToken) ProtoMessage() {}

func (x *ValidationConfig_AccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_source_filters_http_proto_oidc_validation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationConfig_AccessToken.ProtoReflect.Descriptor instead.
func (*ValidationConfig_AccessToken) Descriptor() ([]byte, []int) {
	return file_source_filters_http_proto_oidc_validation_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ValidationConfig_AccessToken) GetLocation() LocationType {
	if x != nil {
		return x.Location
	}
	return LocationType_header
}

func (x *ValidationConfig_AccessToken) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ValidationConfig_AccessToken) GetMetadataFilter() string {
	if x != nil {
		return x.MetadataFilter
	}
	return ""
}

type ValidationConfig_UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Claims        []string       `protobuf:"bytes,1,rep,name=claims,proto3" json:"claims,omitempty"`
	Location      LocationType   `protobuf:"varint,2,opt,name=location,proto3,enum=greymatter_io.gm_proxy.source.filters.http.LocationType" json:"location,omitempty"`
	Key           string         `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	CookieOptions *CookieOptions `protobuf:"bytes,4,opt,name=cookieOptions,proto3" json:"cookieOptions,omitempty"`
}

func (x *ValidationConfig_UserInfo) Reset() {
	*x = ValidationConfig_UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_filters_http_proto_oidc_validation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationConfig_UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationConfig_UserInfo) ProtoMessage() {}

func (x *ValidationConfig_UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_source_filters_http_proto_oidc_validation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationConfig_UserInfo.ProtoReflect.Descriptor instead.
func (*ValidationConfig_UserInfo) Descriptor() ([]byte, []int) {
	return file_source_filters_http_proto_oidc_validation_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ValidationConfig_UserInfo) GetClaims() []string {
	if x != nil {
		return x.Claims
	}
	return nil
}

func (x *ValidationConfig_UserInfo) GetLocation() LocationType {
	if x != nil {
		return x.Location
	}
	return LocationType_header
}

func (x *ValidationConfig_UserInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ValidationConfig_UserInfo) GetCookieOptions() *CookieOptions {
	if x != nil {
		return x.CookieOptions
	}
	return nil
}

type ValidationConfig_FilterTLSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UseTLS             bool       `protobuf:"varint,1,opt,name=useTLS,proto3" json:"useTLS,omitempty"`
	CertPath           string     `protobuf:"bytes,2,opt,name=certPath,proto3" json:"certPath,omitempty"`
	KeyPath            string     `protobuf:"bytes,3,opt,name=keyPath,proto3" json:"keyPath,omitempty"`
	CaPath             string     `protobuf:"bytes,4,opt,name=caPath,proto3" json:"caPath,omitempty"`
	InsecureSkipVerify bool       `protobuf:"varint,5,opt,name=insecureSkipVerify,proto3" json:"insecureSkipVerify,omitempty"`
	ClientAuth         ClientAuth `protobuf:"varint,6,opt,name=clientAuth,proto3,enum=greymatter_io.gm_proxy.source.filters.http.ClientAuth" json:"clientAuth,omitempty"`
}

func (x *ValidationConfig_FilterTLSConfig) Reset() {
	*x = ValidationConfig_FilterTLSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_source_filters_http_proto_oidc_validation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationConfig_FilterTLSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationConfig_FilterTLSConfig) ProtoMessage() {}

func (x *ValidationConfig_FilterTLSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_source_filters_http_proto_oidc_validation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationConfig_FilterTLSConfig.ProtoReflect.Descriptor instead.
func (*ValidationConfig_FilterTLSConfig) Descriptor() ([]byte, []int) {
	return file_source_filters_http_proto_oidc_validation_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ValidationConfig_FilterTLSConfig) GetUseTLS() bool {
	if x != nil {
		return x.UseTLS
	}
	return false
}

func (x *ValidationConfig_FilterTLSConfig) GetCertPath() string {
	if x != nil {
		return x.CertPath
	}
	return ""
}

func (x *ValidationConfig_FilterTLSConfig) GetKeyPath() string {
	if x != nil {
		return x.KeyPath
	}
	return ""
}

func (x *ValidationConfig_FilterTLSConfig) GetCaPath() string {
	if x != nil {
		return x.CaPath
	}
	return ""
}

func (x *ValidationConfig_FilterTLSConfig) GetInsecureSkipVerify() bool {
	if x != nil {
		return x.InsecureSkipVerify
	}
	return false
}

func (x *ValidationConfig_FilterTLSConfig) GetClientAuth() ClientAuth {
	if x != nil {
		return x.ClientAuth
	}
	return ClientAuth_NoClientCert
}

var File_source_filters_http_proto_oidc_validation_proto protoreflect.FileDescriptor

var file_source_filters_http_proto_oidc_validation_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x69, 0x64, 0x63,
	0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x2a, 0x67, 0x72, 0x65, 0x79, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6f,
	0x2e, 0x67, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x1a, 0x30, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x2d,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x33, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x69, 0x64, 0x63, 0x2d,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x08, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x13, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67, 0x72, 0x65, 0x79, 0x6d, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x6f, 0x2e, 0x67, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x61, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x45, 0x2e, 0x67, 0x72, 0x65, 0x79, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6f, 0x2e,
	0x67, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x6a, 0x0a, 0x09, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x67, 0x72, 0x65, 0x79, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x6f, 0x2e, 0x67, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x09, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x9d, 0x01, 0x0a,
	0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x54, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38,
	0x2e, 0x67, 0x72, 0x65, 0x79, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6f, 0x2e, 0x67,
	0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0xeb, 0x01, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x12, 0x54, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x67, 0x72, 0x65, 0x79, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x6f, 0x2e, 0x67, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x5f, 0x0a, 0x0d, 0x63, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x67, 0x72, 0x65, 0x79, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6f,
	0x2e, 0x67, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x43, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x63, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xff, 0x01, 0x0a, 0x0f, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x54, 0x4c, 0x53, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x54, 0x4c, 0x53, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x65, 0x72, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x61, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x53, 0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x12, 0x56, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67, 0x72, 0x65, 0x79, 0x6d,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x6f, 0x2e, 0x67, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x22, 0xb5, 0x01, 0x0a,
	0x15, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65,
	0x12, 0x30, 0x0a, 0x13, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x76,
	0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x79, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_source_filters_http_proto_oidc_validation_proto_rawDescOnce sync.Once
	file_source_filters_http_proto_oidc_validation_proto_rawDescData = file_source_filters_http_proto_oidc_validation_proto_rawDesc
)

func file_source_filters_http_proto_oidc_validation_proto_rawDescGZIP() []byte {
	file_source_filters_http_proto_oidc_validation_proto_rawDescOnce.Do(func() {
		file_source_filters_http_proto_oidc_validation_proto_rawDescData = protoimpl.X.CompressGZIP(file_source_filters_http_proto_oidc_validation_proto_rawDescData)
	})
	return file_source_filters_http_proto_oidc_validation_proto_rawDescData
}

var file_source_filters_http_proto_oidc_validation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_source_filters_http_proto_oidc_validation_proto_goTypes = []interface{}{
	(*ValidationConfig)(nil),                 // 0: greymatter_io.gm_proxy.source.filters.http.ValidationConfig
	(*ValidationRouteConfig)(nil),            // 1: greymatter_io.gm_proxy.source.filters.http.ValidationRouteConfig
	(*ValidationConfig_AccessToken)(nil),     // 2: greymatter_io.gm_proxy.source.filters.http.ValidationConfig.AccessToken
	(*ValidationConfig_UserInfo)(nil),        // 3: greymatter_io.gm_proxy.source.filters.http.ValidationConfig.UserInfo
	(*ValidationConfig_FilterTLSConfig)(nil), // 4: greymatter_io.gm_proxy.source.filters.http.ValidationConfig.FilterTLSConfig
	(LocationType)(0),                        // 5: greymatter_io.gm_proxy.source.filters.http.LocationType
	(*CookieOptions)(nil),                    // 6: greymatter_io.gm_proxy.source.filters.http.CookieOptions
	(ClientAuth)(0),                          // 7: greymatter_io.gm_proxy.source.filters.http.ClientAuth
}
var file_source_filters_http_proto_oidc_validation_proto_depIdxs = []int32{
	2, // 0: greymatter_io.gm_proxy.source.filters.http.ValidationConfig.accessToken:type_name -> greymatter_io.gm_proxy.source.filters.http.ValidationConfig.AccessToken
	3, // 1: greymatter_io.gm_proxy.source.filters.http.ValidationConfig.userInfo:type_name -> greymatter_io.gm_proxy.source.filters.http.ValidationConfig.UserInfo
	4, // 2: greymatter_io.gm_proxy.source.filters.http.ValidationConfig.TLSConfig:type_name -> greymatter_io.gm_proxy.source.filters.http.ValidationConfig.FilterTLSConfig
	5, // 3: greymatter_io.gm_proxy.source.filters.http.ValidationConfig.AccessToken.location:type_name -> greymatter_io.gm_proxy.source.filters.http.LocationType
	5, // 4: greymatter_io.gm_proxy.source.filters.http.ValidationConfig.UserInfo.location:type_name -> greymatter_io.gm_proxy.source.filters.http.LocationType
	6, // 5: greymatter_io.gm_proxy.source.filters.http.ValidationConfig.UserInfo.cookieOptions:type_name -> greymatter_io.gm_proxy.source.filters.http.CookieOptions
	7, // 6: greymatter_io.gm_proxy.source.filters.http.ValidationConfig.FilterTLSConfig.clientAuth:type_name -> greymatter_io.gm_proxy.source.filters.http.ClientAuth
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_source_filters_http_proto_oidc_validation_proto_init() }
func file_source_filters_http_proto_oidc_validation_proto_init() {
	if File_source_filters_http_proto_oidc_validation_proto != nil {
		return
	}
	file_source_filters_http_proto_ensure_variables_proto_init()
	file_source_filters_http_proto_oidc_authentication_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_source_filters_http_proto_oidc_validation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationConfig); i {
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
		file_source_filters_http_proto_oidc_validation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationRouteConfig); i {
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
		file_source_filters_http_proto_oidc_validation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationConfig_AccessToken); i {
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
		file_source_filters_http_proto_oidc_validation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationConfig_UserInfo); i {
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
		file_source_filters_http_proto_oidc_validation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationConfig_FilterTLSConfig); i {
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
			RawDescriptor: file_source_filters_http_proto_oidc_validation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_source_filters_http_proto_oidc_validation_proto_goTypes,
		DependencyIndexes: file_source_filters_http_proto_oidc_validation_proto_depIdxs,
		MessageInfos:      file_source_filters_http_proto_oidc_validation_proto_msgTypes,
	}.Build()
	File_source_filters_http_proto_oidc_validation_proto = out.File
	file_source_filters_http_proto_oidc_validation_proto_rawDesc = nil
	file_source_filters_http_proto_oidc_validation_proto_goTypes = nil
	file_source_filters_http_proto_oidc_validation_proto_depIdxs = nil
}
