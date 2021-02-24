// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opentelemetry/proto/resource/v1/resource.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/influxdata/telegraf/plugins/common/opentelemetry/common/v1"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Resource information.
type Resource struct {
	// Set of labels that describe the resource.
	Attributes []*v1.KeyValue `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0, then
	// no attributes were dropped.
	DroppedAttributesCount uint32   `protobuf:"varint,2,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_446f73eacf88f3f5, []int{0}
}
func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetAttributes() []*v1.KeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Resource) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Resource)(nil), "opentelemetry.proto.resource.v1.Resource")
}

func init() {
	proto.RegisterFile("opentelemetry/proto/resource/v1/resource.proto", fileDescriptor_446f73eacf88f3f5)
}

var fileDescriptor_446f73eacf88f3f5 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcb, 0x2f, 0x48, 0xcd,
	0x2b, 0x49, 0xcd, 0x49, 0xcd, 0x4d, 0x2d, 0x29, 0xaa, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7,
	0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0xd5, 0x2f, 0x33, 0x84, 0xb3, 0xf5, 0xc0, 0x52,
	0x42, 0xf2, 0x28, 0xea, 0x21, 0x82, 0x7a, 0x70, 0x35, 0x65, 0x86, 0x52, 0x5a, 0xd8, 0x0c, 0x4c,
	0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x03, 0x19, 0x07, 0x61, 0x41, 0xf4, 0x29, 0xf5, 0x32, 0x72, 0x71,
	0x04, 0x41, 0xf5, 0x0a, 0xb9, 0x73, 0x71, 0x25, 0x96, 0x94, 0x14, 0x65, 0x26, 0x95, 0x96, 0xa4,
	0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0xa9, 0xeb, 0x61, 0xb3, 0x0e, 0x6a, 0x46, 0x99,
	0xa1, 0x9e, 0x77, 0x6a, 0x65, 0x58, 0x62, 0x4e, 0x69, 0x6a, 0x10, 0x92, 0x56, 0x21, 0x0b, 0x2e,
	0x89, 0x94, 0xa2, 0xfc, 0x82, 0x82, 0xd4, 0x94, 0x78, 0x84, 0x68, 0x7c, 0x72, 0x7e, 0x69, 0x5e,
	0x89, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6f, 0x90, 0x18, 0x54, 0xde, 0x11, 0x2e, 0xed, 0x0c, 0x92,
	0x75, 0x2a, 0xe7, 0x52, 0xca, 0xcc, 0xd7, 0x23, 0xe0, 0x43, 0x27, 0x5e, 0x98, 0x93, 0x03, 0x40,
	0x52, 0x01, 0x8c, 0x51, 0x0e, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x20, 0x77, 0xe9, 0x83, 0x34,
	0xeb, 0x22, 0xbc, 0x8f, 0x62, 0x96, 0x2e, 0x24, 0x30, 0xd2, 0x53, 0xf3, 0xf4, 0xd3, 0x51, 0x02,
	0x39, 0x89, 0x0d, 0x2c, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xba, 0x7f, 0x2f, 0x93, 0x8e,
	0x01, 0x00, 0x00,
}
