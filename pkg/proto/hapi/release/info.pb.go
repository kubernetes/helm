// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hapi/release/info.proto

package release

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Info describes release information.
type Info struct {
	Status        *Status              `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	FirstDeployed *timestamp.Timestamp `protobuf:"bytes,2,opt,name=first_deployed,json=firstDeployed,proto3" json:"first_deployed,omitempty"`
	LastDeployed  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_deployed,json=lastDeployed,proto3" json:"last_deployed,omitempty"`
	// Deleted tracks when this object was deleted.
	Deleted *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// Description is human-friendly "log entry" about this release.
	Description          string   `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_info_24e32c33d678a107, []int{0}
}
func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (dst *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(dst, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Info) GetFirstDeployed() *timestamp.Timestamp {
	if m != nil {
		return m.FirstDeployed
	}
	return nil
}

func (m *Info) GetLastDeployed() *timestamp.Timestamp {
	if m != nil {
		return m.LastDeployed
	}
	return nil
}

func (m *Info) GetDeleted() *timestamp.Timestamp {
	if m != nil {
		return m.Deleted
	}
	return nil
}

func (m *Info) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Info)(nil), "hapi.release.Info")
}

func init() { proto.RegisterFile("hapi/release/info.proto", fileDescriptor_info_24e32c33d678a107) }

var fileDescriptor_info_24e32c33d678a107 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0x52, 0x5a, 0xd5, 0x6d, 0x19, 0x2c, 0x24, 0x42, 0x16, 0x22, 0xa6, 0x0e, 0xc8,
	0x91, 0x80, 0x1d, 0x81, 0xba, 0xb0, 0x06, 0x26, 0x16, 0xe4, 0xe2, 0x73, 0xb1, 0xe4, 0xe6, 0x2c,
	0xfb, 0x3a, 0xf0, 0x2f, 0xf8, 0xc9, 0xa8, 0xb6, 0x83, 0xd2, 0xa9, 0xab, 0xbf, 0xf7, 0x3e, 0xbf,
	0x63, 0x57, 0xdf, 0xd2, 0x99, 0xc6, 0x83, 0x05, 0x19, 0xa0, 0x31, 0x9d, 0x46, 0xe1, 0x3c, 0x12,
	0xf2, 0xc5, 0x01, 0x88, 0x0c, 0xaa, 0x9b, 0x2d, 0xe2, 0xd6, 0x42, 0x13, 0xd9, 0x66, 0xaf, 0x1b,
	0x32, 0x3b, 0x08, 0x24, 0x77, 0x2e, 0xc5, 0xab, 0xeb, 0x23, 0x4f, 0x20, 0x49, 0xfb, 0x90, 0xd0,
	0xed, 0xef, 0x88, 0x8d, 0x5f, 0x3b, 0x8d, 0xfc, 0x8e, 0x4d, 0x12, 0x28, 0x8b, 0xba, 0x58, 0xcd,
	0xef, 0x2f, 0xc5, 0xf0, 0x0f, 0xf1, 0x16, 0x59, 0x9b, 0x33, 0xfc, 0x99, 0x5d, 0x68, 0xe3, 0x03,
	0x7d, 0x2a, 0x70, 0x16, 0x7f, 0x40, 0x95, 0xa3, 0xd8, 0xaa, 0x44, 0xda, 0x22, 0xfa, 0x2d, 0xe2,
	0xbd, 0xdf, 0xd2, 0x2e, 0x63, 0x63, 0x9d, 0x0b, 0xfc, 0x89, 0x2d, 0xad, 0x1c, 0x1a, 0xce, 0x4e,
	0x1a, 0x16, 0x87, 0xc2, 0xbf, 0xe0, 0x91, 0x4d, 0x15, 0x58, 0x20, 0x50, 0xe5, 0xf8, 0x64, 0xb5,
	0x8f, 0xf2, 0x9a, 0xcd, 0xd7, 0x10, 0xbe, 0xbc, 0x71, 0x64, 0xb0, 0x2b, 0xcf, 0xeb, 0x62, 0x35,
	0x6b, 0x87, 0x4f, 0x2f, 0xb3, 0x8f, 0x69, 0xbe, 0x7a, 0x33, 0x89, 0xa6, 0x87, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x1a, 0x52, 0x8f, 0x9c, 0x89, 0x01, 0x00, 0x00,
}
