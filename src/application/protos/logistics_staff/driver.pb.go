// Code generated by protoc-gen-go. DO NOT EDIT.
// source: driver.proto

package logistics_staff

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type DriverProfile struct {
	DriverId             string   `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,3,opt,name=phone_number,proto3" json:"phone_number,omitempty"`
	ImageUrl             string   `protobuf:"bytes,4,opt,name=image_url,proto3" json:"image_url,omitempty"`
	Active               bool     `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DriverProfile) Reset()         { *m = DriverProfile{} }
func (m *DriverProfile) String() string { return proto.CompactTextString(m) }
func (*DriverProfile) ProtoMessage()    {}
func (*DriverProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{0}
}

func (m *DriverProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverProfile.Unmarshal(m, b)
}
func (m *DriverProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverProfile.Marshal(b, m, deterministic)
}
func (m *DriverProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverProfile.Merge(m, src)
}
func (m *DriverProfile) XXX_Size() int {
	return xxx_messageInfo_DriverProfile.Size(m)
}
func (m *DriverProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverProfile.DiscardUnknown(m)
}

var xxx_messageInfo_DriverProfile proto.InternalMessageInfo

func (m *DriverProfile) GetDriverId() string {
	if m != nil {
		return m.DriverId
	}
	return ""
}

func (m *DriverProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DriverProfile) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *DriverProfile) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *DriverProfile) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *DriverProfile) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *DriverProfile) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type RegisterDriverRequest struct {
	PhoneNumber          string   `protobuf:"bytes,1,opt,name=phone_number,proto3" json:"phone_number,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterDriverRequest) Reset()         { *m = RegisterDriverRequest{} }
func (m *RegisterDriverRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterDriverRequest) ProtoMessage()    {}
func (*RegisterDriverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{1}
}

func (m *RegisterDriverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterDriverRequest.Unmarshal(m, b)
}
func (m *RegisterDriverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterDriverRequest.Marshal(b, m, deterministic)
}
func (m *RegisterDriverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterDriverRequest.Merge(m, src)
}
func (m *RegisterDriverRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterDriverRequest.Size(m)
}
func (m *RegisterDriverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterDriverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterDriverRequest proto.InternalMessageInfo

func (m *RegisterDriverRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *RegisterDriverRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterDriverResponse struct {
	DriverId             string   `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterDriverResponse) Reset()         { *m = RegisterDriverResponse{} }
func (m *RegisterDriverResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterDriverResponse) ProtoMessage()    {}
func (*RegisterDriverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{2}
}

func (m *RegisterDriverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterDriverResponse.Unmarshal(m, b)
}
func (m *RegisterDriverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterDriverResponse.Marshal(b, m, deterministic)
}
func (m *RegisterDriverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterDriverResponse.Merge(m, src)
}
func (m *RegisterDriverResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterDriverResponse.Size(m)
}
func (m *RegisterDriverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterDriverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterDriverResponse proto.InternalMessageInfo

func (m *RegisterDriverResponse) GetDriverId() string {
	if m != nil {
		return m.DriverId
	}
	return ""
}

type LoginDriverRequest struct {
	PhoneNumber          string   `protobuf:"bytes,1,opt,name=phone_number,proto3" json:"phone_number,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginDriverRequest) Reset()         { *m = LoginDriverRequest{} }
func (m *LoginDriverRequest) String() string { return proto.CompactTextString(m) }
func (*LoginDriverRequest) ProtoMessage()    {}
func (*LoginDriverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{3}
}

func (m *LoginDriverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginDriverRequest.Unmarshal(m, b)
}
func (m *LoginDriverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginDriverRequest.Marshal(b, m, deterministic)
}
func (m *LoginDriverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginDriverRequest.Merge(m, src)
}
func (m *LoginDriverRequest) XXX_Size() int {
	return xxx_messageInfo_LoginDriverRequest.Size(m)
}
func (m *LoginDriverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginDriverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginDriverRequest proto.InternalMessageInfo

func (m *LoginDriverRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *LoginDriverRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginDriverResponse struct {
	DriverId             string   `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginDriverResponse) Reset()         { *m = LoginDriverResponse{} }
func (m *LoginDriverResponse) String() string { return proto.CompactTextString(m) }
func (*LoginDriverResponse) ProtoMessage()    {}
func (*LoginDriverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{4}
}

func (m *LoginDriverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginDriverResponse.Unmarshal(m, b)
}
func (m *LoginDriverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginDriverResponse.Marshal(b, m, deterministic)
}
func (m *LoginDriverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginDriverResponse.Merge(m, src)
}
func (m *LoginDriverResponse) XXX_Size() int {
	return xxx_messageInfo_LoginDriverResponse.Size(m)
}
func (m *LoginDriverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginDriverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginDriverResponse proto.InternalMessageInfo

func (m *LoginDriverResponse) GetDriverId() string {
	if m != nil {
		return m.DriverId
	}
	return ""
}

type ConfirmSmsCodeRequest struct {
	DriverId             string   `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	SmsCode              string   `protobuf:"bytes,2,opt,name=sms_code,proto3" json:"sms_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmSmsCodeRequest) Reset()         { *m = ConfirmSmsCodeRequest{} }
func (m *ConfirmSmsCodeRequest) String() string { return proto.CompactTextString(m) }
func (*ConfirmSmsCodeRequest) ProtoMessage()    {}
func (*ConfirmSmsCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{5}
}

func (m *ConfirmSmsCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmSmsCodeRequest.Unmarshal(m, b)
}
func (m *ConfirmSmsCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmSmsCodeRequest.Marshal(b, m, deterministic)
}
func (m *ConfirmSmsCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmSmsCodeRequest.Merge(m, src)
}
func (m *ConfirmSmsCodeRequest) XXX_Size() int {
	return xxx_messageInfo_ConfirmSmsCodeRequest.Size(m)
}
func (m *ConfirmSmsCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmSmsCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmSmsCodeRequest proto.InternalMessageInfo

func (m *ConfirmSmsCodeRequest) GetDriverId() string {
	if m != nil {
		return m.DriverId
	}
	return ""
}

func (m *ConfirmSmsCodeRequest) GetSmsCode() string {
	if m != nil {
		return m.SmsCode
	}
	return ""
}

type ConfirmSmsCodeResponse struct {
	Profile              *DriverProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	AccessToken          string         `protobuf:"bytes,2,opt,name=access_token,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConfirmSmsCodeResponse) Reset()         { *m = ConfirmSmsCodeResponse{} }
func (m *ConfirmSmsCodeResponse) String() string { return proto.CompactTextString(m) }
func (*ConfirmSmsCodeResponse) ProtoMessage()    {}
func (*ConfirmSmsCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{6}
}

func (m *ConfirmSmsCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmSmsCodeResponse.Unmarshal(m, b)
}
func (m *ConfirmSmsCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmSmsCodeResponse.Marshal(b, m, deterministic)
}
func (m *ConfirmSmsCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmSmsCodeResponse.Merge(m, src)
}
func (m *ConfirmSmsCodeResponse) XXX_Size() int {
	return xxx_messageInfo_ConfirmSmsCodeResponse.Size(m)
}
func (m *ConfirmSmsCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmSmsCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmSmsCodeResponse proto.InternalMessageInfo

func (m *ConfirmSmsCodeResponse) GetProfile() *DriverProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *ConfirmSmsCodeResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type GetDriverProfileRequest struct {
	DriverId             string   `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDriverProfileRequest) Reset()         { *m = GetDriverProfileRequest{} }
func (m *GetDriverProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetDriverProfileRequest) ProtoMessage()    {}
func (*GetDriverProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{7}
}

func (m *GetDriverProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDriverProfileRequest.Unmarshal(m, b)
}
func (m *GetDriverProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDriverProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetDriverProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDriverProfileRequest.Merge(m, src)
}
func (m *GetDriverProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetDriverProfileRequest.Size(m)
}
func (m *GetDriverProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDriverProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDriverProfileRequest proto.InternalMessageInfo

func (m *GetDriverProfileRequest) GetDriverId() string {
	if m != nil {
		return m.DriverId
	}
	return ""
}

type GetDriverProfileResponse struct {
	Profile              *DriverProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetDriverProfileResponse) Reset()         { *m = GetDriverProfileResponse{} }
func (m *GetDriverProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetDriverProfileResponse) ProtoMessage()    {}
func (*GetDriverProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{8}
}

func (m *GetDriverProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDriverProfileResponse.Unmarshal(m, b)
}
func (m *GetDriverProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDriverProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetDriverProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDriverProfileResponse.Merge(m, src)
}
func (m *GetDriverProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetDriverProfileResponse.Size(m)
}
func (m *GetDriverProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDriverProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDriverProfileResponse proto.InternalMessageInfo

func (m *GetDriverProfileResponse) GetProfile() *DriverProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type UpdateDriverProfileRequest struct {
	DriverId             string   `protobuf:"bytes,1,opt,name=driver_id,proto3" json:"driver_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl             string   `protobuf:"bytes,3,opt,name=image_url,proto3" json:"image_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateDriverProfileRequest) Reset()         { *m = UpdateDriverProfileRequest{} }
func (m *UpdateDriverProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDriverProfileRequest) ProtoMessage()    {}
func (*UpdateDriverProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{9}
}

func (m *UpdateDriverProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDriverProfileRequest.Unmarshal(m, b)
}
func (m *UpdateDriverProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDriverProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateDriverProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDriverProfileRequest.Merge(m, src)
}
func (m *UpdateDriverProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDriverProfileRequest.Size(m)
}
func (m *UpdateDriverProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDriverProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDriverProfileRequest proto.InternalMessageInfo

func (m *UpdateDriverProfileRequest) GetDriverId() string {
	if m != nil {
		return m.DriverId
	}
	return ""
}

func (m *UpdateDriverProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateDriverProfileRequest) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

type UpdateDriverProfileResponse struct {
	Profile              *DriverProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateDriverProfileResponse) Reset()         { *m = UpdateDriverProfileResponse{} }
func (m *UpdateDriverProfileResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateDriverProfileResponse) ProtoMessage()    {}
func (*UpdateDriverProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_521003751d596b5e, []int{10}
}

func (m *UpdateDriverProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDriverProfileResponse.Unmarshal(m, b)
}
func (m *UpdateDriverProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDriverProfileResponse.Marshal(b, m, deterministic)
}
func (m *UpdateDriverProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDriverProfileResponse.Merge(m, src)
}
func (m *UpdateDriverProfileResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateDriverProfileResponse.Size(m)
}
func (m *UpdateDriverProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDriverProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDriverProfileResponse proto.InternalMessageInfo

func (m *UpdateDriverProfileResponse) GetProfile() *DriverProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func init() {
	proto.RegisterType((*DriverProfile)(nil), "DriverProfile")
	proto.RegisterType((*RegisterDriverRequest)(nil), "RegisterDriverRequest")
	proto.RegisterType((*RegisterDriverResponse)(nil), "RegisterDriverResponse")
	proto.RegisterType((*LoginDriverRequest)(nil), "LoginDriverRequest")
	proto.RegisterType((*LoginDriverResponse)(nil), "LoginDriverResponse")
	proto.RegisterType((*ConfirmSmsCodeRequest)(nil), "ConfirmSmsCodeRequest")
	proto.RegisterType((*ConfirmSmsCodeResponse)(nil), "ConfirmSmsCodeResponse")
	proto.RegisterType((*GetDriverProfileRequest)(nil), "GetDriverProfileRequest")
	proto.RegisterType((*GetDriverProfileResponse)(nil), "GetDriverProfileResponse")
	proto.RegisterType((*UpdateDriverProfileRequest)(nil), "UpdateDriverProfileRequest")
	proto.RegisterType((*UpdateDriverProfileResponse)(nil), "UpdateDriverProfileResponse")
}

func init() {
	proto.RegisterFile("driver.proto", fileDescriptor_521003751d596b5e)
}

var fileDescriptor_521003751d596b5e = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x5f, 0x4f, 0xf2, 0x30,
	0x18, 0xc5, 0xb3, 0x17, 0x5e, 0xfe, 0x3c, 0xa2, 0x89, 0x23, 0xe0, 0x82, 0xc6, 0x90, 0x5e, 0x71,
	0x85, 0x89, 0x24, 0x7a, 0x2f, 0x24, 0xdc, 0x78, 0xa1, 0x53, 0x63, 0xe2, 0xcd, 0x52, 0xb6, 0x67,
	0xd8, 0xb8, 0xad, 0xb3, 0xed, 0xf0, 0xc3, 0xfa, 0x65, 0xcc, 0xba, 0x0d, 0x1c, 0x10, 0x85, 0xc4,
	0x3b, 0x7a, 0x0e, 0x3d, 0xfd, 0xb5, 0x39, 0xcf, 0xa0, 0xe5, 0x09, 0xb6, 0x40, 0x31, 0x8c, 0x05,
	0x57, 0x9c, 0x7c, 0x1a, 0x70, 0x38, 0xd1, 0xc2, 0x9d, 0xe0, 0x3e, 0x0b, 0xd0, 0x3c, 0x83, 0x66,
	0xf6, 0x0f, 0x87, 0x79, 0x96, 0xd1, 0x37, 0x06, 0x4d, 0x7b, 0x25, 0x98, 0x26, 0x54, 0x23, 0x1a,
	0xa2, 0xf5, 0x4f, 0x1b, 0xfa, 0xb7, 0x49, 0xa0, 0x15, 0xbf, 0xf2, 0x08, 0x9d, 0x28, 0x09, 0x67,
	0x28, 0xac, 0x8a, 0xf6, 0x4a, 0x5a, 0x9a, 0xca, 0x42, 0x3a, 0x47, 0x27, 0x11, 0x81, 0x55, 0xcd,
	0x52, 0x97, 0x82, 0xd9, 0x85, 0x1a, 0x75, 0x15, 0x5b, 0xa0, 0xf5, 0xbf, 0x6f, 0x0c, 0x1a, 0x76,
	0xbe, 0x32, 0xcf, 0x01, 0x5c, 0x81, 0x54, 0xa1, 0xe7, 0x50, 0x65, 0xd5, 0xf4, 0xb6, 0x6f, 0x4a,
	0xea, 0x27, 0xb1, 0x57, 0xf8, 0xf5, 0xcc, 0x5f, 0x29, 0xe4, 0x19, 0x3a, 0x36, 0xce, 0x99, 0x54,
	0x28, 0xb2, 0x4b, 0xda, 0xf8, 0x9e, 0xa0, 0x54, 0x1b, 0xc8, 0xc6, 0x16, 0xe4, 0x1e, 0x34, 0x62,
	0x2a, 0xe5, 0x07, 0x17, 0x5e, 0x7e, 0xdd, 0xe5, 0x9a, 0x5c, 0x41, 0x77, 0x3d, 0x58, 0xc6, 0x3c,
	0x92, 0xbf, 0x3c, 0x1f, 0x79, 0x04, 0xf3, 0x96, 0xcf, 0x59, 0xf4, 0xb7, 0x34, 0x23, 0x68, 0x97,
	0x52, 0x77, 0x42, 0xb9, 0x87, 0xce, 0x98, 0x47, 0x3e, 0x13, 0xe1, 0x43, 0x28, 0xc7, 0xdc, 0xc3,
	0x82, 0xe6, 0xe7, 0x02, 0xf4, 0xa0, 0x21, 0x43, 0xe9, 0xb8, 0xdc, 0x2b, 0x4a, 0xb0, 0x5c, 0x13,
	0x1f, 0xba, 0xeb, 0x91, 0x39, 0xca, 0x00, 0xea, 0x71, 0xd6, 0x2f, 0x9d, 0x78, 0x70, 0x79, 0x34,
	0x2c, 0xb5, 0xce, 0x2e, 0xec, 0xf4, 0x2d, 0xa8, 0xeb, 0xa2, 0x94, 0x8e, 0xe2, 0x6f, 0x18, 0xe5,
	0x67, 0x94, 0x34, 0x72, 0x0d, 0x27, 0x53, 0x54, 0xe5, 0x80, 0x5d, 0xe0, 0xc9, 0x04, 0xac, 0xcd,
	0x8d, 0xfb, 0x22, 0x92, 0x00, 0x7a, 0x4f, 0xba, 0x63, 0xfb, 0x13, 0x6c, 0x9d, 0x9f, 0xd2, 0x6c,
	0x54, 0xd6, 0x66, 0x83, 0x4c, 0xe1, 0x74, 0xeb, 0x69, 0xfb, 0x62, 0xdf, 0xb4, 0x5f, 0x8e, 0x87,
	0x17, 0x01, 0x4f, 0x5b, 0xcb, 0x5c, 0xe9, 0x48, 0x45, 0x7d, 0x7f, 0x56, 0xd3, 0x9f, 0x81, 0xd1,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xd4, 0x51, 0x9a, 0x16, 0x04, 0x00, 0x00,
}
