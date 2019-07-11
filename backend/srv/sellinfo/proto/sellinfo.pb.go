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

type SellInfoQueryResponse_Status int32

const (
	SellInfoQueryResponse_UNKNOWN     SellInfoQueryResponse_Status = 0
	SellInfoQueryResponse_EMPTY_PARAM SellInfoQueryResponse_Status = -1
	SellInfoQueryResponse_SUCCESS     SellInfoQueryResponse_Status = 1
	SellInfoQueryResponse_NOT_EXIST   SellInfoQueryResponse_Status = 2
)

var SellInfoQueryResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "EMPTY_PARAM",
	1:  "SUCCESS",
	2:  "NOT_EXIST",
}

var SellInfoQueryResponse_Status_value = map[string]int32{
	"UNKNOWN":     0,
	"EMPTY_PARAM": -1,
	"SUCCESS":     1,
	"NOT_EXIST":   2,
}

func (x SellInfoQueryResponse_Status) String() string {
	return proto.EnumName(SellInfoQueryResponse_Status_name, int32(x))
}

func (SellInfoQueryResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{1, 0}
}

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

type ContentCreateRequest_Type int32

const (
	ContentCreateRequest_UNKNOWN ContentCreateRequest_Type = 0
	ContentCreateRequest_PICTURE ContentCreateRequest_Type = 1
	ContentCreateRequest_VIDEO   ContentCreateRequest_Type = 2
)

var ContentCreateRequest_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "PICTURE",
	2: "VIDEO",
}

var ContentCreateRequest_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"PICTURE": 1,
	"VIDEO":   2,
}

func (x ContentCreateRequest_Type) String() string {
	return proto.EnumName(ContentCreateRequest_Type_name, int32(x))
}

func (ContentCreateRequest_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{4, 0}
}

type ContentCreateResponse_Status int32

const (
	ContentCreateResponse_UNKNOWN       ContentCreateResponse_Status = 0
	ContentCreateResponse_INVALID_PARAM ContentCreateResponse_Status = -1
	ContentCreateResponse_SUCCESS       ContentCreateResponse_Status = 1
	ContentCreateResponse_INVALID_TOKEN ContentCreateResponse_Status = 2
)

var ContentCreateResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "INVALID_TOKEN",
}

var ContentCreateResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"INVALID_TOKEN": 2,
}

func (x ContentCreateResponse_Status) String() string {
	return proto.EnumName(ContentCreateResponse_Status_name, int32(x))
}

func (ContentCreateResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{5, 0}
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
	return fileDescriptor_9c322b32f573704b, []int{0}
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

type SellInfoQueryResponse struct {
	Status               SellInfoQueryResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=SellInfoQueryResponse_Status" json:"status,omitempty"`
	SellInfoId           int32                        `protobuf:"varint,2,opt,name=sellInfoId,proto3" json:"sellInfoId,omitempty"`
	ValidTime            int64                        `protobuf:"varint,3,opt,name=validTime,proto3" json:"validTime,omitempty"`
	GoodName             string                       `protobuf:"bytes,4,opt,name=goodName,proto3" json:"goodName,omitempty"`
	Description          string                       `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	ContentId            string                       `protobuf:"bytes,6,opt,name=contentId,proto3" json:"contentId,omitempty"`
	Tag                  []string                     `protobuf:"bytes,7,rep,name=tag,proto3" json:"tag,omitempty"`
	ContentToken         string                       `protobuf:"bytes,8,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SellInfoQueryResponse) Reset()         { *m = SellInfoQueryResponse{} }
func (m *SellInfoQueryResponse) String() string { return proto.CompactTextString(m) }
func (*SellInfoQueryResponse) ProtoMessage()    {}
func (*SellInfoQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{1}
}

func (m *SellInfoQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SellInfoQueryResponse.Unmarshal(m, b)
}
func (m *SellInfoQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SellInfoQueryResponse.Marshal(b, m, deterministic)
}
func (m *SellInfoQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellInfoQueryResponse.Merge(m, src)
}
func (m *SellInfoQueryResponse) XXX_Size() int {
	return xxx_messageInfo_SellInfoQueryResponse.Size(m)
}
func (m *SellInfoQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SellInfoQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SellInfoQueryResponse proto.InternalMessageInfo

func (m *SellInfoQueryResponse) GetStatus() SellInfoQueryResponse_Status {
	if m != nil {
		return m.Status
	}
	return SellInfoQueryResponse_UNKNOWN
}

func (m *SellInfoQueryResponse) GetSellInfoId() int32 {
	if m != nil {
		return m.SellInfoId
	}
	return 0
}

func (m *SellInfoQueryResponse) GetValidTime() int64 {
	if m != nil {
		return m.ValidTime
	}
	return 0
}

func (m *SellInfoQueryResponse) GetGoodName() string {
	if m != nil {
		return m.GoodName
	}
	return ""
}

func (m *SellInfoQueryResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SellInfoQueryResponse) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *SellInfoQueryResponse) GetTag() []string {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *SellInfoQueryResponse) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
}

type SellInfoCreateRequest struct {
	ValidTime            int64    `protobuf:"varint,1,opt,name=validTime,proto3" json:"validTime,omitempty"`
	GoodName             string   `protobuf:"bytes,2,opt,name=goodName,proto3" json:"goodName,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ContentId            string   `protobuf:"bytes,4,opt,name=contentId,proto3" json:"contentId,omitempty"`
	Tag                  []string `protobuf:"bytes,5,rep,name=tag,proto3" json:"tag,omitempty"`
	ContentToken         string   `protobuf:"bytes,6,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
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

func (m *SellInfoCreateRequest) GetTag() []string {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *SellInfoCreateRequest) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
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

type ContentCreateRequest struct {
	ContentId            string                    `protobuf:"bytes,1,opt,name=contentId,proto3" json:"contentId,omitempty"`
	ContentToken         string                    `protobuf:"bytes,2,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
	Content              []byte                    `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Type                 ContentCreateRequest_Type `protobuf:"varint,4,opt,name=type,proto3,enum=ContentCreateRequest_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ContentCreateRequest) Reset()         { *m = ContentCreateRequest{} }
func (m *ContentCreateRequest) String() string { return proto.CompactTextString(m) }
func (*ContentCreateRequest) ProtoMessage()    {}
func (*ContentCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{4}
}

func (m *ContentCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentCreateRequest.Unmarshal(m, b)
}
func (m *ContentCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentCreateRequest.Marshal(b, m, deterministic)
}
func (m *ContentCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentCreateRequest.Merge(m, src)
}
func (m *ContentCreateRequest) XXX_Size() int {
	return xxx_messageInfo_ContentCreateRequest.Size(m)
}
func (m *ContentCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContentCreateRequest proto.InternalMessageInfo

func (m *ContentCreateRequest) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *ContentCreateRequest) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
}

func (m *ContentCreateRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ContentCreateRequest) GetType() ContentCreateRequest_Type {
	if m != nil {
		return m.Type
	}
	return ContentCreateRequest_UNKNOWN
}

type ContentCreateResponse struct {
	Status               ContentCreateResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=ContentCreateResponse_Status" json:"status,omitempty"`
	ContentId            string                       `protobuf:"bytes,2,opt,name=contentId,proto3" json:"contentId,omitempty"`
	ContentToken         string                       `protobuf:"bytes,3,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ContentCreateResponse) Reset()         { *m = ContentCreateResponse{} }
func (m *ContentCreateResponse) String() string { return proto.CompactTextString(m) }
func (*ContentCreateResponse) ProtoMessage()    {}
func (*ContentCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c322b32f573704b, []int{5}
}

func (m *ContentCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentCreateResponse.Unmarshal(m, b)
}
func (m *ContentCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentCreateResponse.Marshal(b, m, deterministic)
}
func (m *ContentCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentCreateResponse.Merge(m, src)
}
func (m *ContentCreateResponse) XXX_Size() int {
	return xxx_messageInfo_ContentCreateResponse.Size(m)
}
func (m *ContentCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContentCreateResponse proto.InternalMessageInfo

func (m *ContentCreateResponse) GetStatus() ContentCreateResponse_Status {
	if m != nil {
		return m.Status
	}
	return ContentCreateResponse_UNKNOWN
}

func (m *ContentCreateResponse) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *ContentCreateResponse) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
}

func init() {
	proto.RegisterEnum("SellInfoQueryResponse_Status", SellInfoQueryResponse_Status_name, SellInfoQueryResponse_Status_value)
	proto.RegisterEnum("SellInfoCreateResponse_Status", SellInfoCreateResponse_Status_name, SellInfoCreateResponse_Status_value)
	proto.RegisterEnum("ContentCreateRequest_Type", ContentCreateRequest_Type_name, ContentCreateRequest_Type_value)
	proto.RegisterEnum("ContentCreateResponse_Status", ContentCreateResponse_Status_name, ContentCreateResponse_Status_value)
	proto.RegisterType((*SellInfoQueryRequest)(nil), "SellInfoQueryRequest")
	proto.RegisterType((*SellInfoQueryResponse)(nil), "SellInfoQueryResponse")
	proto.RegisterType((*SellInfoCreateRequest)(nil), "SellInfoCreateRequest")
	proto.RegisterType((*SellInfoCreateResponse)(nil), "SellInfoCreateResponse")
	proto.RegisterType((*ContentCreateRequest)(nil), "ContentCreateRequest")
	proto.RegisterType((*ContentCreateResponse)(nil), "ContentCreateResponse")
}

func init() { proto.RegisterFile("sellinfo.proto", fileDescriptor_9c322b32f573704b) }

var fileDescriptor_9c322b32f573704b = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x6f, 0xda, 0x30,
	0x14, 0xc7, 0xe7, 0x04, 0x02, 0x3c, 0x5a, 0x94, 0x59, 0x85, 0x45, 0xd1, 0x56, 0xa1, 0x9c, 0x90,
	0x26, 0xe5, 0xc0, 0x34, 0xa6, 0x1d, 0x19, 0xcd, 0x21, 0x62, 0x0d, 0xd4, 0x09, 0xdd, 0x76, 0xaa,
	0xb2, 0xc6, 0xad, 0xa2, 0xa5, 0x71, 0x46, 0xcc, 0x24, 0x0e, 0xfb, 0x74, 0x3b, 0xec, 0xb8, 0x0f,
	0xb1, 0xe3, 0x3e, 0xc4, 0x26, 0x42, 0x68, 0x4a, 0x6a, 0x8a, 0x26, 0x95, 0x13, 0xef, 0xf9, 0x6f,
	0x3b, 0xef, 0xf7, 0x7f, 0x7e, 0xd0, 0x4a, 0x69, 0x14, 0x85, 0xf1, 0x15, 0x33, 0x93, 0x39, 0xe3,
	0xcc, 0x18, 0xc0, 0x91, 0x4b, 0xa3, 0xc8, 0x8e, 0xaf, 0xd8, 0xd9, 0x82, 0xce, 0x97, 0x84, 0x7e,
	0x5d, 0xd0, 0x94, 0xe3, 0x63, 0x80, 0x34, 0xcf, 0xdb, 0x81, 0x86, 0xba, 0xa8, 0x57, 0x25, 0x77,
	0x32, 0xc6, 0x1f, 0x09, 0xda, 0xa5, 0x8d, 0x69, 0xc2, 0xe2, 0x94, 0xe2, 0xd7, 0xa0, 0xa4, 0xdc,
	0xe7, 0x8b, 0x34, 0xdb, 0xd5, 0xea, 0xbf, 0x30, 0x85, 0x3a, 0xd3, 0xcd, 0x44, 0x24, 0x17, 0x97,
	0x2e, 0x94, 0xca, 0x17, 0xe2, 0xe7, 0xd0, 0xf8, 0xe6, 0x47, 0x61, 0xe0, 0x85, 0x37, 0x54, 0x93,
	0xbb, 0xa8, 0x27, 0x93, 0x22, 0x81, 0x75, 0xa8, 0x5f, 0x33, 0x16, 0x38, 0xfe, 0x0d, 0xd5, 0x2a,
	0x5d, 0xd4, 0x6b, 0x90, 0xdb, 0x18, 0x77, 0xa1, 0x19, 0xd0, 0xf4, 0x72, 0x1e, 0x26, 0x3c, 0x64,
	0xb1, 0x56, 0xcd, 0x96, 0xef, 0xa6, 0x56, 0x67, 0x5f, 0xb2, 0x98, 0xd3, 0x98, 0xdb, 0x81, 0xa6,
	0x64, 0xeb, 0x45, 0x02, 0xab, 0x20, 0x73, 0xff, 0x5a, 0xab, 0x75, 0xe5, 0x5e, 0x83, 0xac, 0xfe,
	0x62, 0x03, 0x0e, 0xf2, 0x65, 0x8f, 0x7d, 0xa1, 0xb1, 0x56, 0xcf, 0xb6, 0x6c, 0xe5, 0x8c, 0x31,
	0x28, 0xeb, 0x0a, 0x71, 0x13, 0x6a, 0x33, 0x67, 0xec, 0x4c, 0x3e, 0x38, 0xea, 0x13, 0xac, 0x41,
	0xd3, 0x3a, 0x9d, 0x7a, 0x9f, 0x2e, 0xa6, 0x43, 0x32, 0x3c, 0x55, 0xff, 0x6e, 0x7e, 0x68, 0x25,
	0x73, 0x67, 0xa3, 0x91, 0xe5, 0xba, 0x2a, 0xc2, 0x87, 0xd0, 0x70, 0x26, 0xde, 0x85, 0xf5, 0xd1,
	0x76, 0x3d, 0x55, 0x32, 0x7e, 0xa2, 0x82, 0xf6, 0x68, 0x4e, 0x7d, 0x4e, 0x37, 0x3e, 0x6d, 0x61,
	0x41, 0x0f, 0x61, 0x91, 0x1e, 0xc6, 0x22, 0xef, 0xc1, 0x52, 0xd9, 0x81, 0xa5, 0xba, 0x1b, 0x8b,
	0x22, 0xc0, 0xf2, 0x03, 0x41, 0xa7, 0x5c, 0x49, 0xde, 0x38, 0x83, 0x52, 0xe3, 0x1c, 0x9b, 0x62,
	0xe1, 0x7f, 0x76, 0x8e, 0x71, 0x26, 0x76, 0x42, 0x87, 0x43, 0xdb, 0x39, 0x1f, 0xbe, 0xb7, 0x4f,
	0xf6, 0x78, 0xf1, 0xb4, 0x10, 0x7a, 0x93, 0xb1, 0xe5, 0xa8, 0x92, 0xf1, 0x0b, 0xc1, 0xd1, 0x68,
	0x5d, 0xd6, 0x3d, 0x3b, 0x0a, 0x64, 0xa8, 0x8c, 0xac, 0x0c, 0x48, 0xba, 0x0f, 0x08, 0x6b, 0x50,
	0xcb, 0xe3, 0xcc, 0x92, 0x03, 0xb2, 0x09, 0xb1, 0x09, 0x15, 0xbe, 0x4c, 0xd6, 0xfd, 0xdd, 0xea,
	0xeb, 0xa6, 0xe8, 0x03, 0x4c, 0x6f, 0x99, 0x50, 0x92, 0xe9, 0x8c, 0x97, 0x50, 0x59, 0x45, 0xdb,
	0x55, 0x37, 0xa1, 0x36, 0xb5, 0x47, 0xde, 0x8c, 0x58, 0x2a, 0xc2, 0x0d, 0xa8, 0x9e, 0xdb, 0x27,
	0xd6, 0x44, 0x95, 0x8c, 0xdf, 0x08, 0xda, 0xa5, 0x03, 0x77, 0xbe, 0x67, 0xa1, 0xae, 0xec, 0xca,
	0x16, 0x09, 0x69, 0x1f, 0x09, 0x59, 0xd0, 0x2a, 0x8f, 0xef, 0x5b, 0xff, 0x3b, 0xd4, 0x37, 0x3d,
	0x85, 0x07, 0x50, 0xcd, 0x06, 0x12, 0x6e, 0x9b, 0xa2, 0x09, 0xa8, 0x77, 0xc4, 0x73, 0x0b, 0xbf,
	0x05, 0x65, 0x5d, 0x39, 0xee, 0x98, 0xc2, 0x37, 0xa9, 0x3f, 0xdb, 0xd1, 0xb8, 0xfd, 0x77, 0x50,
	0xcb, 0xd9, 0xe1, 0x37, 0xb7, 0xa7, 0xb4, 0x85, 0x46, 0xea, 0x1d, 0x31, 0xe6, 0xcf, 0x4a, 0x36,
	0xb7, 0x5f, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x02, 0xf4, 0xa6, 0x5d, 0xc9, 0x05, 0x00, 0x00,
}
