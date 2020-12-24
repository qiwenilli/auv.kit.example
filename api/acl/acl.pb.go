// Code generated by protoc-gen-go. DO NOT EDIT.
// source: acl.proto

package acl

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

type List struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *List) Reset()         { *m = List{} }
func (m *List) String() string { return proto.CompactTextString(m) }
func (*List) ProtoMessage()    {}
func (*List) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{0}
}

func (m *List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_List.Unmarshal(m, b)
}
func (m *List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_List.Marshal(b, m, deterministic)
}
func (m *List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_List.Merge(m, src)
}
func (m *List) XXX_Size() int {
	return xxx_messageInfo_List.Size(m)
}
func (m *List) XXX_DiscardUnknown() {
	xxx_messageInfo_List.DiscardUnknown(m)
}

var xxx_messageInfo_List proto.InternalMessageInfo

type List_Request struct {
	AssetsId             string   `protobuf:"bytes,1,opt,name=assets_id,json=assetsId,proto3" json:"assets_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *List_Request) Reset()         { *m = List_Request{} }
func (m *List_Request) String() string { return proto.CompactTextString(m) }
func (*List_Request) ProtoMessage()    {}
func (*List_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{0, 0}
}

func (m *List_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_List_Request.Unmarshal(m, b)
}
func (m *List_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_List_Request.Marshal(b, m, deterministic)
}
func (m *List_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_List_Request.Merge(m, src)
}
func (m *List_Request) XXX_Size() int {
	return xxx_messageInfo_List_Request.Size(m)
}
func (m *List_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_List_Request.DiscardUnknown(m)
}

var xxx_messageInfo_List_Request proto.InternalMessageInfo

func (m *List_Request) GetAssetsId() string {
	if m != nil {
		return m.AssetsId
	}
	return ""
}

type List_Response struct {
	AssetsList           []string `protobuf:"bytes,1,rep,name=assets_list,json=assetsList,proto3" json:"assets_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *List_Response) Reset()         { *m = List_Response{} }
func (m *List_Response) String() string { return proto.CompactTextString(m) }
func (*List_Response) ProtoMessage()    {}
func (*List_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{0, 1}
}

func (m *List_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_List_Response.Unmarshal(m, b)
}
func (m *List_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_List_Response.Marshal(b, m, deterministic)
}
func (m *List_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_List_Response.Merge(m, src)
}
func (m *List_Response) XXX_Size() int {
	return xxx_messageInfo_List_Response.Size(m)
}
func (m *List_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_List_Response.DiscardUnknown(m)
}

var xxx_messageInfo_List_Response proto.InternalMessageInfo

func (m *List_Response) GetAssetsList() []string {
	if m != nil {
		return m.AssetsList
	}
	return nil
}

func init() {
	proto.RegisterType((*List)(nil), "acl.List")
	proto.RegisterType((*List_Request)(nil), "acl.List.Request")
	proto.RegisterType((*List_Response)(nil), "acl.List.Response")
}

func init() { proto.RegisterFile("acl.proto", fileDescriptor_a452f070aeef01eb) }

var fileDescriptor_a452f070aeef01eb = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x4c, 0xce, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x4c, 0xce, 0x51, 0x8a, 0xe6, 0x62, 0xf1, 0xc9,
	0x2c, 0x2e, 0x91, 0x52, 0xe3, 0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe6,
	0xe2, 0x4c, 0x2c, 0x2e, 0x4e, 0x2d, 0x29, 0x8e, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0xe2, 0x80, 0x08, 0x78, 0xa6, 0x48, 0x69, 0x73, 0x71, 0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7,
	0x15, 0xa7, 0x0a, 0xc9, 0x73, 0x71, 0x43, 0x15, 0xe6, 0x64, 0x16, 0x97, 0x48, 0x30, 0x2a, 0x30,
	0x6b, 0x70, 0x06, 0x71, 0x41, 0x84, 0x40, 0x86, 0x1a, 0x99, 0x70, 0x31, 0x3b, 0x26, 0xe7, 0x08,
	0xe9, 0x42, 0xec, 0x10, 0x12, 0xd4, 0x03, 0x59, 0x0e, 0x62, 0xea, 0x41, 0xed, 0x92, 0x12, 0x42,
	0x16, 0x82, 0x18, 0xeb, 0xc4, 0x1b, 0xc5, 0x9d, 0x58, 0x90, 0xa9, 0x9f, 0x98, 0x9c, 0x63, 0x9d,
	0x98, 0x9c, 0x93, 0xc4, 0x06, 0x76, 0xad, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x98, 0xa9, 0xe5,
	0x0c, 0xba, 0x00, 0x00, 0x00,
}