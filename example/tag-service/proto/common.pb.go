// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type Pager struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int64    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	TotalRows            int64    `protobuf:"varint,3,opt,name=total_rows,json=totalRows,proto3" json:"total_rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pager) Reset()         { *m = Pager{} }
func (m *Pager) String() string { return proto.CompactTextString(m) }
func (*Pager) ProtoMessage()    {}
func (*Pager) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

func (m *Pager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pager.Unmarshal(m, b)
}
func (m *Pager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pager.Marshal(b, m, deterministic)
}
func (m *Pager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pager.Merge(m, src)
}
func (m *Pager) XXX_Size() int {
	return xxx_messageInfo_Pager.Size(m)
}
func (m *Pager) XXX_DiscardUnknown() {
	xxx_messageInfo_Pager.DiscardUnknown(m)
}

var xxx_messageInfo_Pager proto.InternalMessageInfo

func (m *Pager) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Pager) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Pager) GetTotalRows() int64 {
	if m != nil {
		return m.TotalRows
	}
	return 0
}

type Error struct {
	Code                 int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details              *anypb.Any `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetDetails() *anypb.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Pager)(nil), "Pager")
	proto.RegisterType((*Error)(nil), "Error")
}

func init() { proto.RegisterFile("proto/common.proto", fileDescriptor_1747d3070a2311a0) }

var fileDescriptor_1747d3070a2311a0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0x31, 0x4b, 0xc0, 0x30,
	0x10, 0x85, 0xa9, 0xb5, 0xd6, 0x9e, 0x88, 0x10, 0x1c, 0xa2, 0x22, 0x48, 0x27, 0xa7, 0x14, 0x74,
	0x74, 0x52, 0x70, 0x97, 0x38, 0x08, 0x2e, 0x25, 0xb6, 0x67, 0x28, 0xb4, 0xbd, 0x92, 0x44, 0x4a,
	0xfb, 0xeb, 0x25, 0x17, 0x3a, 0xdd, 0xbd, 0x7b, 0x8f, 0x2f, 0x2f, 0x20, 0x16, 0x47, 0x81, 0x9a,
	0x8e, 0xa6, 0x89, 0x66, 0xc5, 0xe2, 0xf6, 0xc6, 0x12, 0xd9, 0x11, 0x1b, 0x56, 0x3f, 0x7f, 0xbf,
	0x8d, 0x99, 0xb7, 0x64, 0xd5, 0x5f, 0x50, 0x7c, 0x18, 0x8b, 0x4e, 0x08, 0x38, 0x5d, 0x8c, 0x45,
	0x99, 0x3d, 0x64, 0x8f, 0xb9, 0xe6, 0x5d, 0xdc, 0x41, 0x15, 0x67, 0xeb, 0x87, 0x1d, 0xe5, 0x09,
	0x1b, 0xe7, 0xf1, 0xf0, 0x39, 0xec, 0x28, 0xee, 0x01, 0x02, 0x05, 0x33, 0xb6, 0x8e, 0x56, 0x2f,
	0x73, 0x76, 0x2b, 0xbe, 0x68, 0x5a, 0x7d, 0x8d, 0x50, 0xbc, 0x3b, 0x47, 0x0c, 0xee, 0xa8, 0x4f,
	0xe0, 0x42, 0xf3, 0x2e, 0x24, 0x94, 0x13, 0x7a, 0x1f, 0xdf, 0x8b, 0xd8, 0x4a, 0x1f, 0x52, 0x28,
	0x28, 0x7b, 0x0c, 0x66, 0x18, 0x13, 0xf2, 0xe2, 0xe9, 0x5a, 0xa5, 0xf2, 0xea, 0x28, 0xaf, 0x5e,
	0xe7, 0x4d, 0x1f, 0xa1, 0xb7, 0xab, 0xef, 0x4b, 0x95, 0xfe, 0xf5, 0x92, 0x02, 0x67, 0x3c, 0x9e,
	0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x24, 0x8d, 0x46, 0x88, 0x08, 0x01, 0x00, 0x00,
}
