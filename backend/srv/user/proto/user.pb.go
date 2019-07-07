// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

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

type UserCreateRequest struct {
	StudentId            uint64   `protobuf:"varint,1,opt,name=studentId,proto3" json:"studentId,omitempty"`
	StudentName          string   `protobuf:"bytes,2,opt,name=studentName,proto3" json:"studentName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCreateRequest) Reset()         { *m = UserCreateRequest{} }
func (m *UserCreateRequest) String() string { return proto.CompactTextString(m) }
func (*UserCreateRequest) ProtoMessage()    {}
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreateRequest.Unmarshal(m, b)
}
func (m *UserCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreateRequest.Marshal(b, m, deterministic)
}
func (m *UserCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreateRequest.Merge(m, src)
}
func (m *UserCreateRequest) XXX_Size() int {
	return xxx_messageInfo_UserCreateRequest.Size(m)
}
func (m *UserCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreateRequest proto.InternalMessageInfo

func (m *UserCreateRequest) GetStudentId() uint64 {
	if m != nil {
		return m.StudentId
	}
	return 0
}

func (m *UserCreateRequest) GetStudentName() string {
	if m != nil {
		return m.StudentName
	}
	return ""
}

type UserCreateResponse struct {
	Status               int32    `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	UserId               int32    `protobuf:"varint,4,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCreateResponse) Reset()         { *m = UserCreateResponse{} }
func (m *UserCreateResponse) String() string { return proto.CompactTextString(m) }
func (*UserCreateResponse) ProtoMessage()    {}
func (*UserCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreateResponse.Unmarshal(m, b)
}
func (m *UserCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreateResponse.Marshal(b, m, deterministic)
}
func (m *UserCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreateResponse.Merge(m, src)
}
func (m *UserCreateResponse) XXX_Size() int {
	return xxx_messageInfo_UserCreateResponse.Size(m)
}
func (m *UserCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreateResponse proto.InternalMessageInfo

func (m *UserCreateResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UserCreateResponse) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*UserCreateRequest)(nil), "UserCreateRequest")
	proto.RegisterType((*UserCreateResponse)(nil), "UserCreateResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x31, 0xef, 0x82, 0x30,
	0x10, 0x47, 0xff, 0xfd, 0x8b, 0x24, 0x9c, 0x93, 0x67, 0x62, 0x1a, 0xe3, 0xd0, 0x30, 0x31, 0x31,
	0xc8, 0xe8, 0xa8, 0x0b, 0x8b, 0x43, 0x8d, 0x1f, 0x00, 0xc3, 0x8d, 0x52, 0xec, 0x5d, 0xbf, 0xbf,
	0xa9, 0x25, 0x91, 0x84, 0xf1, 0xbd, 0xe1, 0xdd, 0xef, 0x00, 0x02, 0x93, 0xaf, 0x47, 0xef, 0xc4,
	0x95, 0x77, 0xd8, 0x3e, 0x98, 0xfc, 0xc5, 0x53, 0x27, 0x64, 0xe9, 0x1d, 0x88, 0x05, 0x8f, 0x50,
	0xb0, 0x84, 0x9e, 0x06, 0x69, 0x7b, 0xad, 0x8c, 0xaa, 0x32, 0xfb, 0x13, 0x68, 0x60, 0x33, 0xc1,
	0xad, 0x7b, 0x91, 0xfe, 0x37, 0xaa, 0x2a, 0xec, 0x5c, 0x95, 0x57, 0xc0, 0x79, 0x94, 0x47, 0x37,
	0x30, 0xe1, 0x1e, 0x72, 0x96, 0x4e, 0x02, 0xeb, 0x95, 0x51, 0xd5, 0xda, 0x4e, 0x14, 0x7d, 0x1c,
	0xd4, 0xf6, 0x3a, 0x4b, 0x3e, 0xd1, 0xe9, 0x0c, 0x59, 0xac, 0x60, 0x03, 0x79, 0x2a, 0x21, 0xd6,
	0x8b, 0xad, 0x87, 0x5d, 0xbd, 0x3c, 0x55, 0xfe, 0x3d, 0xf3, 0xef, 0x7b, 0xcd, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x5a, 0xeb, 0x32, 0x54, 0xec, 0x00, 0x00, 0x00,
}
