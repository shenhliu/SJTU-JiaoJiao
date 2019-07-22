// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sellinfo.proto

package sellinfo

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

type SellInfoCreateResponse_Status int32

const (
	SellInfoCreateResponse_UNKNOWN       SellInfoCreateResponse_Status = 0
	SellInfoCreateResponse_INVALID_PARAM SellInfoCreateResponse_Status = -1
	SellInfoCreateResponse_SUCCESS       SellInfoCreateResponse_Status = 1
	SellInfoCreateResponse_INVALID_TOKEN SellInfoCreateResponse_Status = 2
)

var SellInfoCreateResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "INVALID_TOKEN",
}

var SellInfoCreateResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"INVALID_TOKEN": 2,
}

func (x SellInfoCreateResponse_Status) String() string {
	return proto.EnumName(SellInfoCreateResponse_Status_name, int32(x))
}

func (SellInfoCreateResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{3, 0}
}

type SellInfoMsg struct {
	SellInfoId           int32    `protobuf:"varint,1,opt,name=sellInfoId,proto3" json:"sellInfoId,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	ReleaseTime          int64    `protobuf:"varint,3,opt,name=releaseTime,proto3" json:"releaseTime,omitempty"`
	ValidTime            int64    `protobuf:"varint,4,opt,name=validTime,proto3" json:"validTime,omitempty"`
	GoodName             string   `protobuf:"bytes,5,opt,name=goodName,proto3" json:"goodName,omitempty"`
	Price                float64  `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	Description          string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	ContentId            string   `protobuf:"bytes,8,opt,name=contentId,proto3" json:"contentId,omitempty"`
	UserId               int32    `protobuf:"varint,9,opt,name=userId,proto3" json:"userId,omitempty"`
	ClearEmpty           bool     `protobuf:"varint,99,opt,name=clearEmpty,proto3" json:"clearEmpty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SellInfoMsg) Reset()         { *m = SellInfoMsg{} }
func (m *SellInfoMsg) String() string { return proto.CompactTextString(m) }
func (*SellInfoMsg) ProtoMessage()    {}
func (*SellInfoMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{0}
}

func (m *SellInfoMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoMsg.Unmarshal(m, b)
}
func (m *SellInfoMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoMsg.Marshal(b, m, deterministic)
}
func (m *SellInfoMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoMsg.Merge(m, src)
}
func (m *SellInfoMsg) XXX_Size() int {
	return xxx_messageInfo_SellInfoMsg.Size(m)
}
func (m *SellInfoMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoMsg proto.InternalMessageInfo

func (m *SellInfoMsg) GetSellInfoId() int32 {
	if m != nil {
		return m.SellInfoId
	}
	return 0
}

func (m *SellInfoMsg) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SellInfoMsg) GetReleaseTime() int64 {
	if m != nil {
		return m.ReleaseTime
	}
	return 0
}

func (m *SellInfoMsg) GetValidTime() int64 {
	if m != nil {
		return m.ValidTime
	}
	return 0
}

func (m *SellInfoMsg) GetGoodName() string {
	if m != nil {
		return m.GoodName
	}
	return ""
}

func (m *SellInfoMsg) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SellInfoMsg) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SellInfoMsg) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *SellInfoMsg) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SellInfoMsg) GetClearEmpty() bool {
	if m != nil {
		return m.ClearEmpty
	}
	return false
}

type SellInfoQueryRequest struct {
	SellInfoId           int32    `protobuf:"varint,1,opt,name=sellInfoId,proto3" json:"sellInfoId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SellInfoQueryRequest) Reset()         { *m = SellInfoQueryRequest{} }
func (m *SellInfoQueryRequest) String() string { return proto.CompactTextString(m) }
func (*SellInfoQueryRequest) ProtoMessage()    {}
func (*SellInfoQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{1}
}

func (m *SellInfoQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoQueryRequest.Unmarshal(m, b)
}
func (m *SellInfoQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoQueryRequest.Marshal(b, m, deterministic)
}
func (m *SellInfoQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoQueryRequest.Merge(m, src)
}
func (m *SellInfoQueryRequest) XXX_Size() int {
	return xxx_messageInfo_SellInfoQueryRequest.Size(m)
}
func (m *SellInfoQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoQueryRequest proto.InternalMessageInfo

func (m *SellInfoQueryRequest) GetSellInfoId() int32 {
	if m != nil {
		return m.SellInfoId
	}
	return 0
}

type SellInfoCreateRequest struct {
	ValidTime            int64    `protobuf:"varint,1,opt,name=validTime,proto3" json:"validTime,omitempty"`
	GoodName             string   `protobuf:"bytes,2,opt,name=goodName,proto3" json:"goodName,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ContentId            string   `protobuf:"bytes,5,opt,name=contentId,proto3" json:"contentId,omitempty"`
	ContentToken         string   `protobuf:"bytes,6,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
	UserId               int32    `protobuf:"varint,7,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SellInfoCreateRequest) Reset()         { *m = SellInfoCreateRequest{} }
func (m *SellInfoCreateRequest) String() string { return proto.CompactTextString(m) }
func (*SellInfoCreateRequest) ProtoMessage()    {}
func (*SellInfoCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{2}
}

func (m *SellInfoCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoCreateRequest.Unmarshal(m, b)
}
func (m *SellInfoCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoCreateRequest.Marshal(b, m, deterministic)
}
func (m *SellInfoCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoCreateRequest.Merge(m, src)
}
func (m *SellInfoCreateRequest) XXX_Size() int {
	return xxx_messageInfo_SellInfoCreateRequest.Size(m)
}
func (m *SellInfoCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoCreateRequest proto.InternalMessageInfo

func (m *SellInfoCreateRequest) GetValidTime() int64 {
	if m != nil {
		return m.ValidTime
	}
	return 0
}

func (m *SellInfoCreateRequest) GetGoodName() string {
	if m != nil {
		return m.GoodName
	}
	return ""
}

func (m *SellInfoCreateRequest) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *SellInfoCreateRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SellInfoCreateRequest) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *SellInfoCreateRequest) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
}

func (m *SellInfoCreateRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type SellInfoCreateResponse struct {
	Status               SellInfoCreateResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=SellInfoCreateResponse_Status" json:"status,omitempty"`
	SellInfoId           int32                         `protobuf:"varint,2,opt,name=sellInfoId,proto3" json:"sellInfoId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SellInfoCreateResponse) Reset()         { *m = SellInfoCreateResponse{} }
func (m *SellInfoCreateResponse) String() string { return proto.CompactTextString(m) }
func (*SellInfoCreateResponse) ProtoMessage()    {}
func (*SellInfoCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{3}
}

func (m *SellInfoCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoCreateResponse.Unmarshal(m, b)
}
func (m *SellInfoCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoCreateResponse.Marshal(b, m, deterministic)
}
func (m *SellInfoCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoCreateResponse.Merge(m, src)
}
func (m *SellInfoCreateResponse) XXX_Size() int {
	return xxx_messageInfo_SellInfoCreateResponse.Size(m)
}
func (m *SellInfoCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoCreateResponse proto.InternalMessageInfo

func (m *SellInfoCreateResponse) GetStatus() SellInfoCreateResponse_Status {
	if m != nil {
		return m.Status
	}
	return SellInfoCreateResponse_UNKNOWN
}

func (m *SellInfoCreateResponse) GetSellInfoId() int32 {
	if m != nil {
		return m.SellInfoId
	}
	return 0
}

type SellInfoFindRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	GoodName             string   `protobuf:"bytes,3,opt,name=goodName,proto3" json:"goodName,omitempty"`
	LowPrice             float64  `protobuf:"fixed64,4,opt,name=lowPrice,proto3" json:"lowPrice,omitempty"`
	HighPrice            float64  `protobuf:"fixed64,5,opt,name=highPrice,proto3" json:"highPrice,omitempty"`
	Limit                uint32   `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint32   `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SellInfoFindRequest) Reset()         { *m = SellInfoFindRequest{} }
func (m *SellInfoFindRequest) String() string { return proto.CompactTextString(m) }
func (*SellInfoFindRequest) ProtoMessage()    {}
func (*SellInfoFindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{4}
}

func (m *SellInfoFindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoFindRequest.Unmarshal(m, b)
}
func (m *SellInfoFindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoFindRequest.Marshal(b, m, deterministic)
}
func (m *SellInfoFindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoFindRequest.Merge(m, src)
}
func (m *SellInfoFindRequest) XXX_Size() int {
	return xxx_messageInfo_SellInfoFindRequest.Size(m)
}
func (m *SellInfoFindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoFindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoFindRequest proto.InternalMessageInfo

func (m *SellInfoFindRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SellInfoFindRequest) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SellInfoFindRequest) GetGoodName() string {
	if m != nil {
		return m.GoodName
	}
	return ""
}

func (m *SellInfoFindRequest) GetLowPrice() float64 {
	if m != nil {
		return m.LowPrice
	}
	return 0
}

func (m *SellInfoFindRequest) GetHighPrice() float64 {
	if m != nil {
		return m.HighPrice
	}
	return 0
}

func (m *SellInfoFindRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SellInfoFindRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type SellInfoFindResponse struct {
	SellInfo             []*SellInfoMsg `protobuf:"bytes,1,rep,name=sellInfo,proto3" json:"sellInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SellInfoFindResponse) Reset()         { *m = SellInfoFindResponse{} }
func (m *SellInfoFindResponse) String() string { return proto.CompactTextString(m) }
func (*SellInfoFindResponse) ProtoMessage()    {}
func (*SellInfoFindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{5}
}

func (m *SellInfoFindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoFindResponse.Unmarshal(m, b)
}
func (m *SellInfoFindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoFindResponse.Marshal(b, m, deterministic)
}
func (m *SellInfoFindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoFindResponse.Merge(m, src)
}
func (m *SellInfoFindResponse) XXX_Size() int {
	return xxx_messageInfo_SellInfoFindResponse.Size(m)
}
func (m *SellInfoFindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoFindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoFindResponse proto.InternalMessageInfo

func (m *SellInfoFindResponse) GetSellInfo() []*SellInfoMsg {
	if m != nil {
		return m.SellInfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("SellInfoCreateResponse_Status", SellInfoCreateResponse_Status_name, SellInfoCreateResponse_Status_value)
	proto.RegisterType((*SellInfoMsg)(nil), "SellInfoMsg")
	proto.RegisterType((*SellInfoQueryRequest)(nil), "SellInfoQueryRequest")
	proto.RegisterType((*SellInfoCreateRequest)(nil), "SellInfoCreateRequest")
	proto.RegisterType((*SellInfoCreateResponse)(nil), "SellInfoCreateResponse")
	proto.RegisterType((*SellInfoFindRequest)(nil), "SellInfoFindRequest")
	proto.RegisterType((*SellInfoFindResponse)(nil), "SellInfoFindResponse")
}

func init() { proto.RegisterFile("sellinfo.proto", fileDescriptor_9c322b32f573704b) }

var fileDescriptor_9c322b32f573704b = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x6e, 0xda, 0x4a,
	0x14, 0x7d, 0x03, 0x18, 0xcc, 0x25, 0x44, 0xbc, 0x29, 0x50, 0xcb, 0xaa, 0x22, 0xcb, 0x2b, 0x2f,
	0x2a, 0x2f, 0x88, 0x14, 0xa9, 0xbb, 0x22, 0x4a, 0x25, 0x2b, 0x8d, 0x93, 0x0c, 0xa4, 0x5d, 0x56,
	0x2e, 0x1e, 0x88, 0x55, 0xe3, 0xa1, 0x9e, 0xa1, 0x55, 0xfe, 0xa8, 0xbb, 0x7e, 0x44, 0x57, 0xfd,
	0x8e, 0x7e, 0x48, 0xab, 0x19, 0x63, 0x6c, 0x28, 0x24, 0xec, 0xce, 0x39, 0xf7, 0xe2, 0xb9, 0xe7,
	0xdc, 0x19, 0x38, 0xe5, 0x34, 0x8e, 0xa3, 0x64, 0xce, 0xdc, 0x55, 0xca, 0x04, 0xb3, 0x7f, 0x54,
	0xa0, 0x35, 0xa1, 0x71, 0xec, 0x25, 0x73, 0x76, 0xc5, 0x17, 0xf8, 0x0c, 0x80, 0x6f, 0xa0, 0x17,
	0x1a, 0xc8, 0x42, 0x8e, 0x46, 0x4a, 0x0c, 0xee, 0x43, 0x9d, 0x8b, 0x40, 0xac, 0xb9, 0x51, 0x51,
	0xda, 0x06, 0x61, 0x0b, 0x5a, 0x29, 0x8d, 0x69, 0xc0, 0xe9, 0x34, 0x5a, 0x52, 0xa3, 0x6a, 0x21,
	0xa7, 0x4a, 0xca, 0x14, 0x7e, 0x01, 0xcd, 0xaf, 0x41, 0x1c, 0x85, 0x4a, 0xaf, 0x29, 0xbd, 0x20,
	0xb0, 0x09, 0xfa, 0x82, 0xb1, 0xd0, 0x0f, 0x96, 0xd4, 0xd0, 0x2c, 0xe4, 0x34, 0xc9, 0x16, 0xe3,
	0x2e, 0x68, 0xab, 0x34, 0x9a, 0x51, 0xa3, 0x6e, 0x21, 0x07, 0x91, 0x0c, 0xc8, 0x2f, 0x86, 0x94,
	0xcf, 0xd2, 0x68, 0x25, 0x22, 0x96, 0x18, 0x0d, 0xd5, 0x54, 0xa6, 0xe4, 0x17, 0x67, 0x2c, 0x11,
	0x34, 0x11, 0x5e, 0x68, 0xe8, 0x4a, 0x2f, 0x08, 0x39, 0xc9, 0x9a, 0xd3, 0xd4, 0x0b, 0x8d, 0x66,
	0x36, 0x49, 0x86, 0xa4, 0x03, 0xb3, 0x98, 0x06, 0xe9, 0x78, 0xb9, 0x12, 0x0f, 0xc6, 0xcc, 0x42,
	0x8e, 0x4e, 0x4a, 0x8c, 0x7d, 0x01, 0xdd, 0xdc, 0xb0, 0xdb, 0x35, 0x4d, 0x1f, 0x08, 0xfd, 0xb2,
	0xa6, 0x5c, 0x3c, 0xe5, 0x9c, 0xfd, 0x1b, 0x41, 0x2f, 0x6f, 0x1c, 0xa5, 0x34, 0x10, 0x34, 0xef,
	0xdc, 0x71, 0x06, 0x3d, 0xe6, 0x4c, 0xe5, 0x98, 0x33, 0xd5, 0x47, 0x9c, 0xa9, 0x3d, 0xe1, 0x8c,
	0xb6, 0xef, 0x8c, 0x0d, 0x27, 0x1b, 0x30, 0x65, 0x9f, 0x69, 0xa2, 0x6c, 0x6f, 0x92, 0x1d, 0xae,
	0xe4, 0x5e, 0xa3, 0xec, 0x9e, 0xfd, 0x13, 0x41, 0x7f, 0x7f, 0x4a, 0xbe, 0x62, 0x09, 0xa7, 0xf8,
	0x62, 0xbb, 0x3a, 0x72, 0xc6, 0xd3, 0xc1, 0x99, 0x7b, 0xb8, 0xd0, 0x9d, 0xa8, 0xaa, 0xed, 0x6a,
	0xed, 0x1a, 0x5b, 0xf9, 0xc7, 0xd8, 0x5b, 0xa8, 0x67, 0x1d, 0xb8, 0x05, 0x8d, 0x3b, 0xff, 0xd2,
	0xbf, 0xfe, 0xe0, 0x77, 0xfe, 0xc3, 0x26, 0xb4, 0x3d, 0xff, 0xfd, 0xf0, 0x9d, 0xf7, 0xe6, 0xe3,
	0xcd, 0x90, 0x0c, 0xaf, 0x3a, 0x7f, 0xf2, 0x1f, 0x92, 0x85, 0x93, 0xbb, 0xd1, 0x68, 0x3c, 0x99,
	0x74, 0x10, 0xfe, 0xbf, 0x28, 0x9c, 0x5e, 0x5f, 0x8e, 0xfd, 0x4e, 0xc5, 0xfe, 0x85, 0xe0, 0x59,
	0x7e, 0xb8, 0xb7, 0x51, 0x12, 0xe6, 0x49, 0x15, 0x53, 0xa3, 0x9d, 0x9d, 0x39, 0x76, 0x2b, 0xca,
	0xd9, 0x55, 0xf7, 0xb2, 0x33, 0x41, 0x8f, 0xd9, 0xb7, 0x1b, 0x15, 0x5f, 0x4d, 0xc5, 0xb7, 0xc5,
	0x32, 0x9f, 0xfb, 0x68, 0x71, 0x9f, 0x89, 0x9a, 0x12, 0x0b, 0x42, 0xa6, 0x1e, 0x47, 0xcb, 0x48,
	0xa8, 0x60, 0xda, 0x24, 0x03, 0xf2, 0x0c, 0x6c, 0x3e, 0xe7, 0x54, 0xa8, 0x44, 0xda, 0x64, 0x83,
	0xec, 0xd7, 0xc5, 0xbe, 0x66, 0xa3, 0x6c, 0xe2, 0x70, 0x40, 0xcf, 0x4d, 0x34, 0x90, 0x55, 0x75,
	0x5a, 0x83, 0x13, 0xb7, 0xf4, 0x12, 0x90, 0xad, 0x3a, 0xf8, 0x8e, 0x40, 0xcf, 0x15, 0xfc, 0x12,
	0x34, 0xb5, 0xf6, 0xb8, 0xe7, 0x1e, 0xba, 0x06, 0xe6, 0xce, 0x9f, 0xe0, 0x57, 0x50, 0xcf, 0xc2,
	0xc5, 0x7d, 0xf7, 0xe0, 0xf2, 0x9b, 0xcf, 0x8f, 0x6c, 0x01, 0x3e, 0x87, 0x9a, 0x3c, 0x2f, 0xee,
	0xba, 0x07, 0x92, 0x30, 0x7b, 0xee, 0xa1, 0xa1, 0x3e, 0xd5, 0xd5, 0xab, 0x76, 0xfe, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x8b, 0x2b, 0x26, 0x95, 0xe7, 0x04, 0x00, 0x00,
}
