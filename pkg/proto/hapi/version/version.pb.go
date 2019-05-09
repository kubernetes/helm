// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hapi/version/version.proto

package version

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Version struct {
	// Sem ver string for the version
	SemVer               string   `protobuf:"bytes,1,opt,name=sem_ver,json=semVer,proto3" json:"sem_ver,omitempty"`
	GitCommit            string   `protobuf:"bytes,2,opt,name=git_commit,json=gitCommit,proto3" json:"git_commit,omitempty"`
	GitTreeState         string   `protobuf:"bytes,3,opt,name=git_tree_state,json=gitTreeState,proto3" json:"git_tree_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_version_fc47a726bdff3b83, []int{0}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (dst *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(dst, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetSemVer() string {
	if m != nil {
		return m.SemVer
	}
	return ""
}

func (m *Version) GetGitCommit() string {
	if m != nil {
		return m.GitCommit
	}
	return ""
}

func (m *Version) GetGitTreeState() string {
	if m != nil {
		return m.GitTreeState
	}
	return ""
}

func init() {
	proto.RegisterType((*Version)(nil), "hapi.version.Version")
}

func init() { proto.RegisterFile("hapi/version/version.proto", fileDescriptor_version_fc47a726bdff3b83) }

var fileDescriptor_version_fc47a726bdff3b83 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x48, 0x2c, 0xc8,
	0xd4, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x83, 0xd1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x3c, 0x20, 0x39, 0x3d, 0xa8, 0x98, 0x52, 0x3a, 0x17, 0x7b, 0x18, 0x84, 0x29, 0x24, 0xce,
	0xc5, 0x5e, 0x9c, 0x9a, 0x1b, 0x5f, 0x96, 0x5a, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4,
	0x56, 0x9c, 0x9a, 0x1b, 0x96, 0x5a, 0x24, 0x24, 0xcb, 0xc5, 0x95, 0x9e, 0x59, 0x12, 0x9f, 0x9c,
	0x9f, 0x9b, 0x9b, 0x59, 0x22, 0xc1, 0x04, 0x96, 0xe3, 0x4c, 0xcf, 0x2c, 0x71, 0x06, 0x0b, 0x08,
	0xa9, 0x70, 0xf1, 0x81, 0xa4, 0x4b, 0x8a, 0x52, 0x53, 0xe3, 0x8b, 0x4b, 0x12, 0x4b, 0x52, 0x25,
	0x98, 0xc1, 0x4a, 0x78, 0xd2, 0x33, 0x4b, 0x42, 0x8a, 0x52, 0x53, 0x83, 0x41, 0x62, 0x4e, 0x9c,
	0x51, 0xec, 0x50, 0x3b, 0x93, 0xd8, 0xc0, 0x0e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x20,
	0xcc, 0x0e, 0x1b, 0xa6, 0x00, 0x00, 0x00,
}
