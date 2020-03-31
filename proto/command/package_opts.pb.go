// Code generated by protoc-gen-go. DO NOT EDIT.
// source: command/package_opts.proto

package command

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

// PackageOpts is a package option.
// NEXT_ID_TO_USE: 8
type PackageOpts struct {
	// emulation_command specifies which emulation layer is necessary for this package.
	// If empty, it means no emulation layer is necessary.
	EmulationCommand string `protobuf:"bytes,7,opt,name=emulation_command,json=emulationCommand,proto3" json:"emulation_command,omitempty"`
	// output_file_filters is a set of regexp to filter output files.
	OutputFileFilters    []string `protobuf:"bytes,5,rep,name=output_file_filters,json=outputFileFilters,proto3" json:"output_file_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PackageOpts) Reset()         { *m = PackageOpts{} }
func (m *PackageOpts) String() string { return proto.CompactTextString(m) }
func (*PackageOpts) ProtoMessage()    {}
func (*PackageOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_290ec2fb95a58d08, []int{0}
}

func (m *PackageOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageOpts.Unmarshal(m, b)
}
func (m *PackageOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageOpts.Marshal(b, m, deterministic)
}
func (m *PackageOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageOpts.Merge(m, src)
}
func (m *PackageOpts) XXX_Size() int {
	return xxx_messageInfo_PackageOpts.Size(m)
}
func (m *PackageOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageOpts.DiscardUnknown(m)
}

var xxx_messageInfo_PackageOpts proto.InternalMessageInfo

func (m *PackageOpts) GetEmulationCommand() string {
	if m != nil {
		return m.EmulationCommand
	}
	return ""
}

func (m *PackageOpts) GetOutputFileFilters() []string {
	if m != nil {
		return m.OutputFileFilters
	}
	return nil
}

func init() {
	proto.RegisterType((*PackageOpts)(nil), "command.PackageOpts")
}

func init() {
	proto.RegisterFile("command/package_opts.proto", fileDescriptor_290ec2fb95a58d08)
}

var fileDescriptor_290ec2fb95a58d08 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0xd1, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x4c, 0x4f, 0x8d, 0xcf, 0x2f, 0x28, 0x29,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0xca, 0x29, 0x75, 0x30, 0x72, 0x71, 0x07,
	0x40, 0xe4, 0xfd, 0x0b, 0x4a, 0x8a, 0x85, 0xb4, 0xb9, 0x04, 0x53, 0x73, 0x4b, 0x73, 0x12, 0x4b,
	0x32, 0xf3, 0xf3, 0xe2, 0xa1, 0x8a, 0x24, 0xd8, 0x15, 0x18, 0x35, 0x38, 0x83, 0x04, 0xe0, 0x12,
	0xce, 0x10, 0x71, 0x21, 0x3d, 0x2e, 0xe1, 0xfc, 0xd2, 0x92, 0x82, 0xd2, 0x92, 0xf8, 0xb4, 0xcc,
	0x9c, 0x54, 0x10, 0x51, 0x92, 0x5a, 0x54, 0x2c, 0xc1, 0xaa, 0xc0, 0xac, 0xc1, 0x19, 0x24, 0x08,
	0x91, 0x72, 0xcb, 0xcc, 0x49, 0x75, 0x83, 0x48, 0x78, 0xb1, 0x70, 0x30, 0x0a, 0x30, 0x79, 0xb1,
	0x70, 0x30, 0x09, 0x30, 0x7b, 0xb1, 0x70, 0x30, 0x0b, 0xb0, 0x78, 0xb1, 0x70, 0xb0, 0x08, 0xb0,
	0x7a, 0xb1, 0x70, 0xb0, 0x09, 0xb0, 0x27, 0xb1, 0x81, 0x9d, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x33, 0x90, 0x89, 0x37, 0xb8, 0x00, 0x00, 0x00,
}
