// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

package session

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type SuccessMessage struct {
	Successfull          bool     `protobuf:"varint,1,opt,name=Successfull,proto3" json:"Successfull,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessMessage) Reset()         { *m = SuccessMessage{} }
func (m *SuccessMessage) String() string { return proto.CompactTextString(m) }
func (*SuccessMessage) ProtoMessage()    {}
func (*SuccessMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{0}
}

func (m *SuccessMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessMessage.Unmarshal(m, b)
}
func (m *SuccessMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessMessage.Marshal(b, m, deterministic)
}
func (m *SuccessMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessMessage.Merge(m, src)
}
func (m *SuccessMessage) XXX_Size() int {
	return xxx_messageInfo_SuccessMessage.Size(m)
}
func (m *SuccessMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessMessage proto.InternalMessageInfo

func (m *SuccessMessage) GetSuccessfull() bool {
	if m != nil {
		return m.Successfull
	}
	return false
}

type CreateSessionMessage struct {
	Ttl                  int64    `protobuf:"varint,1,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSessionMessage) Reset()         { *m = CreateSessionMessage{} }
func (m *CreateSessionMessage) String() string { return proto.CompactTextString(m) }
func (*CreateSessionMessage) ProtoMessage()    {}
func (*CreateSessionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{1}
}

func (m *CreateSessionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSessionMessage.Unmarshal(m, b)
}
func (m *CreateSessionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSessionMessage.Marshal(b, m, deterministic)
}
func (m *CreateSessionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSessionMessage.Merge(m, src)
}
func (m *CreateSessionMessage) XXX_Size() int {
	return xxx_messageInfo_CreateSessionMessage.Size(m)
}
func (m *CreateSessionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSessionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSessionMessage proto.InternalMessageInfo

func (m *CreateSessionMessage) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type GetSessionMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionMessage) Reset()         { *m = GetSessionMessage{} }
func (m *GetSessionMessage) String() string { return proto.CompactTextString(m) }
func (*GetSessionMessage) ProtoMessage()    {}
func (*GetSessionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{2}
}

func (m *GetSessionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionMessage.Unmarshal(m, b)
}
func (m *GetSessionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionMessage.Marshal(b, m, deterministic)
}
func (m *GetSessionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionMessage.Merge(m, src)
}
func (m *GetSessionMessage) XXX_Size() int {
	return xxx_messageInfo_GetSessionMessage.Size(m)
}
func (m *GetSessionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionMessage proto.InternalMessageInfo

func (m *GetSessionMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AddValueToSessionMessage struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string         `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                *_struct.Value `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddValueToSessionMessage) Reset()         { *m = AddValueToSessionMessage{} }
func (m *AddValueToSessionMessage) String() string { return proto.CompactTextString(m) }
func (*AddValueToSessionMessage) ProtoMessage()    {}
func (*AddValueToSessionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{3}
}

func (m *AddValueToSessionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddValueToSessionMessage.Unmarshal(m, b)
}
func (m *AddValueToSessionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddValueToSessionMessage.Marshal(b, m, deterministic)
}
func (m *AddValueToSessionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddValueToSessionMessage.Merge(m, src)
}
func (m *AddValueToSessionMessage) XXX_Size() int {
	return xxx_messageInfo_AddValueToSessionMessage.Size(m)
}
func (m *AddValueToSessionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddValueToSessionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddValueToSessionMessage proto.InternalMessageInfo

func (m *AddValueToSessionMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddValueToSessionMessage) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AddValueToSessionMessage) GetValue() *_struct.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

type AddValuesToSessionMessage struct {
	Id                   string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Values               map[string]*_struct.Value `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AddValuesToSessionMessage) Reset()         { *m = AddValuesToSessionMessage{} }
func (m *AddValuesToSessionMessage) String() string { return proto.CompactTextString(m) }
func (*AddValuesToSessionMessage) ProtoMessage()    {}
func (*AddValuesToSessionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{4}
}

func (m *AddValuesToSessionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddValuesToSessionMessage.Unmarshal(m, b)
}
func (m *AddValuesToSessionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddValuesToSessionMessage.Marshal(b, m, deterministic)
}
func (m *AddValuesToSessionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddValuesToSessionMessage.Merge(m, src)
}
func (m *AddValuesToSessionMessage) XXX_Size() int {
	return xxx_messageInfo_AddValuesToSessionMessage.Size(m)
}
func (m *AddValuesToSessionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AddValuesToSessionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AddValuesToSessionMessage proto.InternalMessageInfo

func (m *AddValuesToSessionMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddValuesToSessionMessage) GetValues() map[string]*_struct.Value {
	if m != nil {
		return m.Values
	}
	return nil
}

type SessionResponse struct {
	Id                   string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Values               map[string]*_struct.Value `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *SessionResponse) Reset()         { *m = SessionResponse{} }
func (m *SessionResponse) String() string { return proto.CompactTextString(m) }
func (*SessionResponse) ProtoMessage()    {}
func (*SessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{5}
}

func (m *SessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionResponse.Unmarshal(m, b)
}
func (m *SessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionResponse.Marshal(b, m, deterministic)
}
func (m *SessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionResponse.Merge(m, src)
}
func (m *SessionResponse) XXX_Size() int {
	return xxx_messageInfo_SessionResponse.Size(m)
}
func (m *SessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SessionResponse proto.InternalMessageInfo

func (m *SessionResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SessionResponse) GetValues() map[string]*_struct.Value {
	if m != nil {
		return m.Values
	}
	return nil
}

type InvalidateSessionMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvalidateSessionMessage) Reset()         { *m = InvalidateSessionMessage{} }
func (m *InvalidateSessionMessage) String() string { return proto.CompactTextString(m) }
func (*InvalidateSessionMessage) ProtoMessage()    {}
func (*InvalidateSessionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{6}
}

func (m *InvalidateSessionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvalidateSessionMessage.Unmarshal(m, b)
}
func (m *InvalidateSessionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvalidateSessionMessage.Marshal(b, m, deterministic)
}
func (m *InvalidateSessionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvalidateSessionMessage.Merge(m, src)
}
func (m *InvalidateSessionMessage) XXX_Size() int {
	return xxx_messageInfo_InvalidateSessionMessage.Size(m)
}
func (m *InvalidateSessionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InvalidateSessionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InvalidateSessionMessage proto.InternalMessageInfo

func (m *InvalidateSessionMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type InvalidateSessionValueMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvalidateSessionValueMessage) Reset()         { *m = InvalidateSessionValueMessage{} }
func (m *InvalidateSessionValueMessage) String() string { return proto.CompactTextString(m) }
func (*InvalidateSessionValueMessage) ProtoMessage()    {}
func (*InvalidateSessionValueMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{7}
}

func (m *InvalidateSessionValueMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvalidateSessionValueMessage.Unmarshal(m, b)
}
func (m *InvalidateSessionValueMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvalidateSessionValueMessage.Marshal(b, m, deterministic)
}
func (m *InvalidateSessionValueMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvalidateSessionValueMessage.Merge(m, src)
}
func (m *InvalidateSessionValueMessage) XXX_Size() int {
	return xxx_messageInfo_InvalidateSessionValueMessage.Size(m)
}
func (m *InvalidateSessionValueMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InvalidateSessionValueMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InvalidateSessionValueMessage proto.InternalMessageInfo

func (m *InvalidateSessionValueMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InvalidateSessionValueMessage) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type InvalidateSessionValuesMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Keys                 []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvalidateSessionValuesMessage) Reset()         { *m = InvalidateSessionValuesMessage{} }
func (m *InvalidateSessionValuesMessage) String() string { return proto.CompactTextString(m) }
func (*InvalidateSessionValuesMessage) ProtoMessage()    {}
func (*InvalidateSessionValuesMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{8}
}

func (m *InvalidateSessionValuesMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvalidateSessionValuesMessage.Unmarshal(m, b)
}
func (m *InvalidateSessionValuesMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvalidateSessionValuesMessage.Marshal(b, m, deterministic)
}
func (m *InvalidateSessionValuesMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvalidateSessionValuesMessage.Merge(m, src)
}
func (m *InvalidateSessionValuesMessage) XXX_Size() int {
	return xxx_messageInfo_InvalidateSessionValuesMessage.Size(m)
}
func (m *InvalidateSessionValuesMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_InvalidateSessionValuesMessage.DiscardUnknown(m)
}

var xxx_messageInfo_InvalidateSessionValuesMessage proto.InternalMessageInfo

func (m *InvalidateSessionValuesMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InvalidateSessionValuesMessage) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterType((*SuccessMessage)(nil), "hobord.session.SuccessMessage")
	proto.RegisterType((*CreateSessionMessage)(nil), "hobord.session.CreateSessionMessage")
	proto.RegisterType((*GetSessionMessage)(nil), "hobord.session.GetSessionMessage")
	proto.RegisterType((*AddValueToSessionMessage)(nil), "hobord.session.AddValueToSessionMessage")
	proto.RegisterType((*AddValuesToSessionMessage)(nil), "hobord.session.AddValuesToSessionMessage")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "hobord.session.AddValuesToSessionMessage.ValuesEntry")
	proto.RegisterType((*SessionResponse)(nil), "hobord.session.SessionResponse")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "hobord.session.SessionResponse.ValuesEntry")
	proto.RegisterType((*InvalidateSessionMessage)(nil), "hobord.session.InvalidateSessionMessage")
	proto.RegisterType((*InvalidateSessionValueMessage)(nil), "hobord.session.InvalidateSessionValueMessage")
	proto.RegisterType((*InvalidateSessionValuesMessage)(nil), "hobord.session.InvalidateSessionValuesMessage")
}

func init() { proto.RegisterFile("session.proto", fileDescriptor_3a6be1b361fa6f14) }

var fileDescriptor_3a6be1b361fa6f14 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x5d, 0x12, 0x36, 0xe8, 0xaf, 0x5a, 0x47, 0x2d, 0x34, 0x42, 0x05, 0xa3, 0x18, 0x0e, 0xe1,
	0x9f, 0x27, 0x15, 0x21, 0x21, 0x6e, 0x63, 0x43, 0x88, 0xc3, 0x0e, 0xb8, 0x08, 0x21, 0x2e, 0xd0,
	0x24, 0x5e, 0x09, 0x8d, 0xe2, 0x29, 0x76, 0x2a, 0xf5, 0xa3, 0x71, 0xe1, 0x1b, 0xf0, 0x9d, 0x50,
	0x9c, 0x04, 0x56, 0x3b, 0x69, 0x7d, 0xe2, 0x66, 0xfd, 0xf2, 0x7b, 0xef, 0xf5, 0xd9, 0xef, 0x15,
	0xf6, 0x05, 0x13, 0x22, 0xe1, 0x19, 0xb9, 0xcc, 0xb9, 0xe4, 0x68, 0xf0, 0x9d, 0x87, 0x3c, 0x8f,
	0x49, 0x3d, 0x1d, 0xdd, 0x9d, 0x73, 0x3e, 0x4f, 0xd9, 0xb1, 0xfa, 0x1a, 0x16, 0x17, 0xc7, 0x42,
	0xe6, 0x45, 0x24, 0xab, 0x6d, 0x3c, 0x81, 0xc1, 0xb4, 0x88, 0x22, 0x26, 0xc4, 0x39, 0x13, 0x62,
	0x36, 0x67, 0x68, 0x0c, 0xfd, 0x7a, 0x72, 0x51, 0xa4, 0xa9, 0xef, 0x8c, 0x9d, 0xe0, 0x06, 0xbd,
	0x3a, 0xc2, 0x01, 0xdc, 0x3a, 0xcd, 0xd9, 0x4c, 0xb2, 0x69, 0x25, 0xd1, 0x20, 0x6f, 0x82, 0x27,
	0x65, 0x85, 0xf0, 0x68, 0x79, 0xc4, 0x0f, 0x61, 0xf8, 0x8e, 0x49, 0x6d, 0x6d, 0x00, 0x6e, 0x12,
	0xab, 0xad, 0x1e, 0x75, 0x93, 0x18, 0xff, 0x00, 0xff, 0x24, 0x8e, 0x3f, 0xcd, 0xd2, 0x82, 0x7d,
	0xe4, 0x9b, 0x77, 0x4b, 0x89, 0x05, 0x5b, 0xf9, 0xae, 0x1a, 0x94, 0x47, 0xf4, 0x0c, 0x76, 0x97,
	0x25, 0xd4, 0xf7, 0xc6, 0x4e, 0xd0, 0x9f, 0x1c, 0x92, 0xca, 0x2e, 0x69, 0xec, 0x12, 0x45, 0x4c,
	0xab, 0x25, 0xfc, 0xdb, 0x81, 0x3b, 0x8d, 0x98, 0xd8, 0xaa, 0x76, 0x0e, 0x7b, 0x0a, 0x26, 0x7c,
	0x6f, 0xec, 0x05, 0xfd, 0xc9, 0x4b, 0xb2, 0x7e, 0xb7, 0xa4, 0x93, 0xaa, 0x52, 0x15, 0x6f, 0x33,
	0x99, 0xaf, 0x68, 0x4d, 0x32, 0xfa, 0x00, 0xfd, 0x2b, 0xe3, 0xc6, 0x8b, 0xd3, 0xe2, 0xc5, 0xb5,
	0xf0, 0xf2, 0xda, 0x7d, 0xe5, 0xe0, 0x9f, 0x0e, 0x1c, 0xd4, 0xca, 0x94, 0x89, 0x4b, 0x9e, 0x09,
	0xd3, 0xc5, 0xe9, 0x5f, 0x17, 0xae, 0x72, 0xf1, 0x54, 0x77, 0xa1, 0x11, 0xfc, 0xaf, 0xdf, 0xfe,
	0x04, 0xfc, 0xf7, 0xd9, 0x72, 0x96, 0x26, 0xb1, 0x19, 0x25, 0x3d, 0x23, 0x27, 0x70, 0xcf, 0xd8,
	0x55, 0x84, 0xd6, 0x41, 0xc1, 0x67, 0x70, 0xd4, 0x4e, 0x21, 0xba, 0x38, 0x10, 0x5c, 0x5b, 0xb0,
	0x55, 0x75, 0x6d, 0x3d, 0xaa, 0xce, 0x93, 0x5f, 0xbb, 0x70, 0x70, 0x56, 0xa3, 0xa7, 0x2c, 0x5f,
	0x26, 0x11, 0x43, 0x14, 0xe0, 0x5f, 0xca, 0xd1, 0x03, 0xfd, 0x7a, 0x8d, 0x06, 0x8c, 0xee, 0x6f,
	0x79, 0x01, 0xbc, 0x83, 0x3e, 0xc3, 0xfe, 0x5a, 0xc7, 0xd0, 0x23, 0x1d, 0xd3, 0x56, 0x41, 0x1b,
	0xe6, 0x6f, 0x30, 0x34, 0xea, 0x86, 0x82, 0xae, 0x64, 0xeb, 0xc1, 0xb6, 0x51, 0x08, 0x01, 0x99,
	0xc5, 0x40, 0x8f, 0xad, 0xcb, 0x63, 0xa3, 0x91, 0xc0, 0x61, 0xfb, 0x6b, 0xa2, 0xe7, 0x3a, 0x78,
	0x63, 0x70, 0x46, 0x47, 0x86, 0xd6, 0xda, 0xdf, 0x21, 0xde, 0x41, 0x0b, 0xb8, 0xdd, 0x11, 0x1c,
	0x44, 0xec, 0xb4, 0x84, 0xbd, 0xd8, 0x57, 0x18, 0x1a, 0x1c, 0xe6, 0xeb, 0x74, 0xf5, 0x66, 0xbb,
	0xc0, 0x9b, 0xde, 0x97, 0xeb, 0xf5, 0xb7, 0x70, 0x4f, 0x75, 0xf3, 0xc5, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc8, 0x9b, 0xf9, 0x25, 0x41, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DSessionServiceClient is the client API for DSessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DSessionServiceClient interface {
	GetSession(ctx context.Context, in *GetSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error)
	CreateSession(ctx context.Context, in *CreateSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error)
	AddValueToSession(ctx context.Context, in *AddValueToSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error)
	AddValuesToSession(ctx context.Context, in *AddValuesToSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error)
	InvalidateSessionValue(ctx context.Context, in *InvalidateSessionValueMessage, opts ...grpc.CallOption) (*SuccessMessage, error)
	InvalidateSessionValues(ctx context.Context, in *InvalidateSessionValuesMessage, opts ...grpc.CallOption) (*SuccessMessage, error)
	InvalidateSession(ctx context.Context, in *InvalidateSessionMessage, opts ...grpc.CallOption) (*SuccessMessage, error)
}

type dSessionServiceClient struct {
	cc *grpc.ClientConn
}

func NewDSessionServiceClient(cc *grpc.ClientConn) DSessionServiceClient {
	return &dSessionServiceClient{cc}
}

func (c *dSessionServiceClient) GetSession(ctx context.Context, in *GetSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := c.cc.Invoke(ctx, "/hobord.session.DSessionService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSessionServiceClient) CreateSession(ctx context.Context, in *CreateSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := c.cc.Invoke(ctx, "/hobord.session.DSessionService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSessionServiceClient) AddValueToSession(ctx context.Context, in *AddValueToSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := c.cc.Invoke(ctx, "/hobord.session.DSessionService/AddValueToSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSessionServiceClient) AddValuesToSession(ctx context.Context, in *AddValuesToSessionMessage, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := c.cc.Invoke(ctx, "/hobord.session.DSessionService/AddValuesToSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSessionServiceClient) InvalidateSessionValue(ctx context.Context, in *InvalidateSessionValueMessage, opts ...grpc.CallOption) (*SuccessMessage, error) {
	out := new(SuccessMessage)
	err := c.cc.Invoke(ctx, "/hobord.session.DSessionService/InvalidateSessionValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSessionServiceClient) InvalidateSessionValues(ctx context.Context, in *InvalidateSessionValuesMessage, opts ...grpc.CallOption) (*SuccessMessage, error) {
	out := new(SuccessMessage)
	err := c.cc.Invoke(ctx, "/hobord.session.DSessionService/InvalidateSessionValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSessionServiceClient) InvalidateSession(ctx context.Context, in *InvalidateSessionMessage, opts ...grpc.CallOption) (*SuccessMessage, error) {
	out := new(SuccessMessage)
	err := c.cc.Invoke(ctx, "/hobord.session.DSessionService/InvalidateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DSessionServiceServer is the server API for DSessionService service.
type DSessionServiceServer interface {
	GetSession(context.Context, *GetSessionMessage) (*SessionResponse, error)
	CreateSession(context.Context, *CreateSessionMessage) (*SessionResponse, error)
	AddValueToSession(context.Context, *AddValueToSessionMessage) (*SessionResponse, error)
	AddValuesToSession(context.Context, *AddValuesToSessionMessage) (*SessionResponse, error)
	InvalidateSessionValue(context.Context, *InvalidateSessionValueMessage) (*SuccessMessage, error)
	InvalidateSessionValues(context.Context, *InvalidateSessionValuesMessage) (*SuccessMessage, error)
	InvalidateSession(context.Context, *InvalidateSessionMessage) (*SuccessMessage, error)
}

// UnimplementedDSessionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDSessionServiceServer struct {
}

func (*UnimplementedDSessionServiceServer) GetSession(ctx context.Context, req *GetSessionMessage) (*SessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (*UnimplementedDSessionServiceServer) CreateSession(ctx context.Context, req *CreateSessionMessage) (*SessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (*UnimplementedDSessionServiceServer) AddValueToSession(ctx context.Context, req *AddValueToSessionMessage) (*SessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddValueToSession not implemented")
}
func (*UnimplementedDSessionServiceServer) AddValuesToSession(ctx context.Context, req *AddValuesToSessionMessage) (*SessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddValuesToSession not implemented")
}
func (*UnimplementedDSessionServiceServer) InvalidateSessionValue(ctx context.Context, req *InvalidateSessionValueMessage) (*SuccessMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateSessionValue not implemented")
}
func (*UnimplementedDSessionServiceServer) InvalidateSessionValues(ctx context.Context, req *InvalidateSessionValuesMessage) (*SuccessMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateSessionValues not implemented")
}
func (*UnimplementedDSessionServiceServer) InvalidateSession(ctx context.Context, req *InvalidateSessionMessage) (*SuccessMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InvalidateSession not implemented")
}

func RegisterDSessionServiceServer(s *grpc.Server, srv DSessionServiceServer) {
	s.RegisterService(&_DSessionService_serviceDesc, srv)
}

func _DSessionService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSessionServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hobord.session.DSessionService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSessionServiceServer).GetSession(ctx, req.(*GetSessionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSessionService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSessionServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hobord.session.DSessionService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSessionServiceServer).CreateSession(ctx, req.(*CreateSessionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSessionService_AddValueToSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddValueToSessionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSessionServiceServer).AddValueToSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hobord.session.DSessionService/AddValueToSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSessionServiceServer).AddValueToSession(ctx, req.(*AddValueToSessionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSessionService_AddValuesToSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddValuesToSessionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSessionServiceServer).AddValuesToSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hobord.session.DSessionService/AddValuesToSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSessionServiceServer).AddValuesToSession(ctx, req.(*AddValuesToSessionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSessionService_InvalidateSessionValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateSessionValueMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSessionServiceServer).InvalidateSessionValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hobord.session.DSessionService/InvalidateSessionValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSessionServiceServer).InvalidateSessionValue(ctx, req.(*InvalidateSessionValueMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSessionService_InvalidateSessionValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateSessionValuesMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSessionServiceServer).InvalidateSessionValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hobord.session.DSessionService/InvalidateSessionValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSessionServiceServer).InvalidateSessionValues(ctx, req.(*InvalidateSessionValuesMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSessionService_InvalidateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvalidateSessionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSessionServiceServer).InvalidateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hobord.session.DSessionService/InvalidateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSessionServiceServer).InvalidateSession(ctx, req.(*InvalidateSessionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _DSessionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hobord.session.DSessionService",
	HandlerType: (*DSessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSession",
			Handler:    _DSessionService_GetSession_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _DSessionService_CreateSession_Handler,
		},
		{
			MethodName: "AddValueToSession",
			Handler:    _DSessionService_AddValueToSession_Handler,
		},
		{
			MethodName: "AddValuesToSession",
			Handler:    _DSessionService_AddValuesToSession_Handler,
		},
		{
			MethodName: "InvalidateSessionValue",
			Handler:    _DSessionService_InvalidateSessionValue_Handler,
		},
		{
			MethodName: "InvalidateSessionValues",
			Handler:    _DSessionService_InvalidateSessionValues_Handler,
		},
		{
			MethodName: "InvalidateSession",
			Handler:    _DSessionService_InvalidateSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}
