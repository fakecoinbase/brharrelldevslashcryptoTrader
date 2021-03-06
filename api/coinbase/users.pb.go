// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/coinbase/users.proto

package api_coinbase

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	ProfileName          string   `protobuf:"bytes,4,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	ProfileBio           string   `protobuf:"bytes,5,opt,name=profile_bio,json=profileBio,proto3" json:"profile_bio,omitempty"`
	ProfileUrl           string   `protobuf:"bytes,6,opt,name=profile_url,json=profileUrl,proto3" json:"profile_url,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,7,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Resource             string   `protobuf:"bytes,8,opt,name=resource,proto3" json:"resource,omitempty"`
	ResourcePath         string   `protobuf:"bytes,9,opt,name=resource_path,json=resourcePath,proto3" json:"resource_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_42bdba14146df480, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *User) GetProfileName() string {
	if m != nil {
		return m.ProfileName
	}
	return ""
}

func (m *User) GetProfileBio() string {
	if m != nil {
		return m.ProfileBio
	}
	return ""
}

func (m *User) GetProfileUrl() string {
	if m != nil {
		return m.ProfileUrl
	}
	return ""
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *User) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *User) GetResourcePath() string {
	if m != nil {
		return m.ResourcePath
	}
	return ""
}

type UserInfoResponse struct {
	TimeZone             string            `protobuf:"bytes,1,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	NativeCurrency       string            `protobuf:"bytes,2,opt,name=native_currency,json=nativeCurrency,proto3" json:"native_currency,omitempty"`
	BitcoinUnit          string            `protobuf:"bytes,3,opt,name=bitcoin_unit,json=bitcoinUnit,proto3" json:"bitcoin_unit,omitempty"`
	Country              map[string]string `protobuf:"bytes,4,rep,name=country,proto3" json:"country,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedAt            string            `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42bdba14146df480, []int{1}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetTimeZone() string {
	if m != nil {
		return m.TimeZone
	}
	return ""
}

func (m *UserInfoResponse) GetNativeCurrency() string {
	if m != nil {
		return m.NativeCurrency
	}
	return ""
}

func (m *UserInfoResponse) GetBitcoinUnit() string {
	if m != nil {
		return m.BitcoinUnit
	}
	return ""
}

func (m *UserInfoResponse) GetCountry() map[string]string {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *UserInfoResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type ShowUser struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowUser) Reset()         { *m = ShowUser{} }
func (m *ShowUser) String() string { return proto.CompactTextString(m) }
func (*ShowUser) ProtoMessage()    {}
func (*ShowUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_42bdba14146df480, []int{2}
}

func (m *ShowUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowUser.Unmarshal(m, b)
}
func (m *ShowUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowUser.Marshal(b, m, deterministic)
}
func (m *ShowUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowUser.Merge(m, src)
}
func (m *ShowUser) XXX_Size() int {
	return xxx_messageInfo_ShowUser.Size(m)
}
func (m *ShowUser) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowUser.DiscardUnknown(m)
}

var xxx_messageInfo_ShowUser proto.InternalMessageInfo

func (m *ShowUser) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ShowUserResponse struct {
	Data                 *User    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowUserResponse) Reset()         { *m = ShowUserResponse{} }
func (m *ShowUserResponse) String() string { return proto.CompactTextString(m) }
func (*ShowUserResponse) ProtoMessage()    {}
func (*ShowUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_42bdba14146df480, []int{3}
}

func (m *ShowUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowUserResponse.Unmarshal(m, b)
}
func (m *ShowUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowUserResponse.Marshal(b, m, deterministic)
}
func (m *ShowUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowUserResponse.Merge(m, src)
}
func (m *ShowUserResponse) XXX_Size() int {
	return xxx_messageInfo_ShowUserResponse.Size(m)
}
func (m *ShowUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShowUserResponse proto.InternalMessageInfo

func (m *ShowUserResponse) GetData() *User {
	if m != nil {
		return m.Data
	}
	return nil
}

type ShowAuthorizationInfo struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Scopes               []string `protobuf:"bytes,2,rep,name=scopes,proto3" json:"scopes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowAuthorizationInfo) Reset()         { *m = ShowAuthorizationInfo{} }
func (m *ShowAuthorizationInfo) String() string { return proto.CompactTextString(m) }
func (*ShowAuthorizationInfo) ProtoMessage()    {}
func (*ShowAuthorizationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_42bdba14146df480, []int{4}
}

func (m *ShowAuthorizationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowAuthorizationInfo.Unmarshal(m, b)
}
func (m *ShowAuthorizationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowAuthorizationInfo.Marshal(b, m, deterministic)
}
func (m *ShowAuthorizationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowAuthorizationInfo.Merge(m, src)
}
func (m *ShowAuthorizationInfo) XXX_Size() int {
	return xxx_messageInfo_ShowAuthorizationInfo.Size(m)
}
func (m *ShowAuthorizationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowAuthorizationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ShowAuthorizationInfo proto.InternalMessageInfo

func (m *ShowAuthorizationInfo) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ShowAuthorizationInfo) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

type UpdateCurrentUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Time_Zone            string   `protobuf:"bytes,2,opt,name=time_Zone,json=timeZone,proto3" json:"time_Zone,omitempty"`
	NativeCurrency       string   `protobuf:"bytes,3,opt,name=native_currency,json=nativeCurrency,proto3" json:"native_currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCurrentUserRequest) Reset()         { *m = UpdateCurrentUserRequest{} }
func (m *UpdateCurrentUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCurrentUserRequest) ProtoMessage()    {}
func (*UpdateCurrentUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42bdba14146df480, []int{5}
}

func (m *UpdateCurrentUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCurrentUserRequest.Unmarshal(m, b)
}
func (m *UpdateCurrentUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCurrentUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCurrentUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCurrentUserRequest.Merge(m, src)
}
func (m *UpdateCurrentUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCurrentUserRequest.Size(m)
}
func (m *UpdateCurrentUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCurrentUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCurrentUserRequest proto.InternalMessageInfo

func (m *UpdateCurrentUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateCurrentUserRequest) GetTime_Zone() string {
	if m != nil {
		return m.Time_Zone
	}
	return ""
}

func (m *UpdateCurrentUserRequest) GetNativeCurrency() string {
	if m != nil {
		return m.NativeCurrency
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "api.coinbase.User")
	proto.RegisterType((*UserInfoResponse)(nil), "api.coinbase.UserInfoResponse")
	proto.RegisterMapType((map[string]string)(nil), "api.coinbase.UserInfoResponse.CountryEntry")
	proto.RegisterType((*ShowUser)(nil), "api.coinbase.ShowUser")
	proto.RegisterType((*ShowUserResponse)(nil), "api.coinbase.ShowUserResponse")
	proto.RegisterType((*ShowAuthorizationInfo)(nil), "api.coinbase.ShowAuthorizationInfo")
	proto.RegisterType((*UpdateCurrentUserRequest)(nil), "api.coinbase.UpdateCurrentUserRequest")
}

func init() { proto.RegisterFile("api/coinbase/users.proto", fileDescriptor_42bdba14146df480) }

var fileDescriptor_42bdba14146df480 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x8b, 0x13, 0x31,
	0x14, 0xb5, 0xd3, 0x6e, 0x3f, 0x6e, 0xbb, 0xbb, 0x25, 0xe8, 0x1a, 0x2a, 0x6a, 0x9d, 0x05, 0x5d,
	0x10, 0x66, 0xa1, 0xbe, 0x48, 0xdf, 0xd6, 0xb2, 0xc8, 0x22, 0x88, 0x54, 0xfa, 0x22, 0xc2, 0x90,
	0xce, 0x64, 0x99, 0x60, 0x9b, 0x8c, 0xc9, 0x9d, 0x4a, 0xf7, 0x37, 0xf8, 0x17, 0xfc, 0xaf, 0x92,
	0x4c, 0x52, 0x5a, 0x5d, 0xf6, 0xa5, 0xe4, 0x9e, 0x73, 0xd2, 0xdc, 0x73, 0xee, 0x1d, 0xa0, 0xac,
	0x14, 0x97, 0x99, 0x12, 0x72, 0xc9, 0x0c, 0xbf, 0xac, 0x0c, 0xd7, 0x26, 0x29, 0xb5, 0x42, 0x45,
	0x06, 0xac, 0x14, 0x49, 0x60, 0xe2, 0xdf, 0x11, 0xb4, 0x16, 0x86, 0x6b, 0x72, 0x02, 0x91, 0xc8,
	0x69, 0x63, 0xdc, 0xb8, 0xe8, 0xcd, 0x23, 0x91, 0x13, 0x02, 0x2d, 0xc9, 0xd6, 0x9c, 0x46, 0x0e,
	0x71, 0x67, 0xf2, 0x0c, 0x7a, 0xf6, 0x9f, 0x52, 0x47, 0x34, 0x1d, 0xd1, 0xb5, 0xc0, 0x67, 0x4b,
	0xbe, 0x82, 0x41, 0xa9, 0xd5, 0xad, 0x58, 0xf1, 0x9a, 0x6f, 0x39, 0xbe, 0xef, 0x31, 0x27, 0x79,
	0x09, 0xa1, 0x4c, 0x97, 0x42, 0xd1, 0x23, 0xa7, 0x00, 0x0f, 0x7d, 0x10, 0x6a, 0x5f, 0x50, 0xe9,
	0x15, 0x6d, 0x1f, 0x08, 0x16, 0x7a, 0x45, 0x9e, 0x03, 0xb0, 0x0d, 0x43, 0xa6, 0x1d, 0xdf, 0x71,
	0x7c, 0xaf, 0x46, 0x2c, 0x3d, 0x82, 0xae, 0xe6, 0x46, 0x55, 0x3a, 0xe3, 0xb4, 0x5b, 0xf7, 0x17,
	0x6a, 0x72, 0x0e, 0xc7, 0xe1, 0x9c, 0x96, 0x0c, 0x0b, 0xda, 0x73, 0x82, 0x41, 0x00, 0xbf, 0x30,
	0x2c, 0xe2, 0x3f, 0x11, 0x0c, 0x6d, 0x1c, 0x37, 0xf2, 0x56, 0xcd, 0xb9, 0x29, 0x95, 0x34, 0xce,
	0x36, 0x8a, 0x35, 0x4f, 0xef, 0x94, 0xe4, 0x3e, 0xa1, 0xae, 0x05, 0xbe, 0x29, 0xc9, 0xc9, 0x1b,
	0x38, 0x95, 0x0c, 0xc5, 0x86, 0xa7, 0x59, 0xa5, 0x35, 0x97, 0xd9, 0xd6, 0x47, 0x76, 0x52, 0xc3,
	0x33, 0x8f, 0xda, 0x7c, 0x96, 0x02, 0x6d, 0xf0, 0x69, 0x25, 0x05, 0xfa, 0xfc, 0xfa, 0x1e, 0x5b,
	0x48, 0x81, 0xe4, 0x1a, 0x3a, 0x99, 0xaa, 0x24, 0xea, 0x2d, 0x6d, 0x8d, 0x9b, 0x17, 0xfd, 0xc9,
	0xdb, 0x64, 0x7f, 0x58, 0xc9, 0xbf, 0x9d, 0x25, 0xb3, 0x5a, 0x7d, 0x6d, 0x7f, 0xe6, 0xe1, 0xae,
	0x0d, 0x29, 0xd3, 0x9c, 0x21, 0xcf, 0x53, 0x86, 0x3e, 0xe5, 0x9e, 0x47, 0xae, 0x70, 0x34, 0x85,
	0xc1, 0xfe, 0x3d, 0x32, 0x84, 0xe6, 0x0f, 0xbe, 0xf5, 0xc6, 0xec, 0x91, 0x3c, 0x86, 0xa3, 0x0d,
	0x5b, 0x55, 0x61, 0xf8, 0x75, 0x31, 0x8d, 0xde, 0x37, 0xe2, 0x73, 0xe8, 0x7e, 0x2d, 0xd4, 0x2f,
	0xb7, 0x31, 0x4f, 0xa1, 0xe3, 0xb6, 0x61, 0xb7, 0x36, 0x6d, 0x5b, 0xde, 0xe4, 0xf1, 0x14, 0x86,
	0x41, 0xb4, 0xcb, 0xf0, 0x35, 0xb4, 0x72, 0x86, 0xcc, 0x29, 0xfb, 0x13, 0xf2, 0xbf, 0xaf, 0xb9,
	0xe3, 0xe3, 0x8f, 0xf0, 0xc4, 0xde, 0xbd, 0xaa, 0xb0, 0x50, 0x5a, 0xdc, 0x31, 0x14, 0x4a, 0x5a,
	0xcb, 0xe4, 0x0c, 0xda, 0x6b, 0x8e, 0x85, 0xda, 0x3d, 0x56, 0x57, 0x16, 0x37, 0x99, 0x2a, 0xb9,
	0xa1, 0xd1, 0xb8, 0x69, 0xf1, 0xba, 0x8a, 0x11, 0xe8, 0xa2, 0xcc, 0x19, 0xfa, 0x01, 0x60, 0xdd,
	0xcd, 0xcf, 0x8a, 0x1b, 0xdc, 0xed, 0x76, 0xe3, 0x70, 0xb7, 0xdd, 0x90, 0xed, 0x50, 0xbd, 0xef,
	0x07, 0x87, 0xdc, 0xbc, 0x6f, 0xc8, 0x93, 0xef, 0x70, 0x3c, 0xf3, 0xae, 0xec, 0x83, 0x86, 0x7c,
	0x82, 0x53, 0xeb, 0x67, 0xaf, 0x09, 0x72, 0x76, 0x68, 0x3e, 0x44, 0x35, 0x7a, 0x71, 0x3f, 0x1e,
	0x22, 0x8c, 0x1f, 0x2d, 0xdb, 0xee, 0x0b, 0x7e, 0xf7, 0x37, 0x00, 0x00, 0xff, 0xff, 0x31, 0x67,
	0x4f, 0x72, 0xdd, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoinbaseUsersClient is the client API for CoinbaseUsers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoinbaseUsersClient interface {
	ShowCurrentUser(ctx context.Context, in *ShowUser, opts ...grpc.CallOption) (*ShowUserResponse, error)
}

type coinbaseUsersClient struct {
	cc *grpc.ClientConn
}

func NewCoinbaseUsersClient(cc *grpc.ClientConn) CoinbaseUsersClient {
	return &coinbaseUsersClient{cc}
}

func (c *coinbaseUsersClient) ShowCurrentUser(ctx context.Context, in *ShowUser, opts ...grpc.CallOption) (*ShowUserResponse, error) {
	out := new(ShowUserResponse)
	err := c.cc.Invoke(ctx, "/api.coinbase.CoinbaseUsers/ShowCurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoinbaseUsersServer is the server API for CoinbaseUsers service.
type CoinbaseUsersServer interface {
	ShowCurrentUser(context.Context, *ShowUser) (*ShowUserResponse, error)
}

// UnimplementedCoinbaseUsersServer can be embedded to have forward compatible implementations.
type UnimplementedCoinbaseUsersServer struct {
}

func (*UnimplementedCoinbaseUsersServer) ShowCurrentUser(ctx context.Context, req *ShowUser) (*ShowUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowCurrentUser not implemented")
}

func RegisterCoinbaseUsersServer(s *grpc.Server, srv CoinbaseUsersServer) {
	s.RegisterService(&_CoinbaseUsers_serviceDesc, srv)
}

func _CoinbaseUsers_ShowCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinbaseUsersServer).ShowCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.coinbase.CoinbaseUsers/ShowCurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinbaseUsersServer).ShowCurrentUser(ctx, req.(*ShowUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoinbaseUsers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.coinbase.CoinbaseUsers",
	HandlerType: (*CoinbaseUsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShowCurrentUser",
			Handler:    _CoinbaseUsers_ShowCurrentUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/coinbase/users.proto",
}
