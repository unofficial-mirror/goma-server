// Copyright 2018 The Goma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: command/setup.proto

package command

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PrefixStyle represents a style of prefix.
// if it is ASIS, prefix parameter is read as-is.
// if it is REGEXP, prefix parameter is read as a regular expression.
// the install is automatically extended to all matching directories.
// e.g. if there is directory named clang-r123 and clang-r234, and
// clang-r[0-9]+ is set in prefix, install for both clang-r123 and clang-r234
// are made.
// in this case, you can use grouping i.e. parentheses, and you can use
// the variables in binary_hash_from field in ForSelector.
// if it is REGEXP_MAY_OMIT_BINARY_HASH_MISSING, then apply REGEXP rule
// but skip if a file specified with binary_hash_from is missing.
type Install_PrefixStyle int32

const (
	Install_ASIS                                Install_PrefixStyle = 0
	Install_REGEXP                              Install_PrefixStyle = 1
	Install_REGEXP_MAY_SKIP_BINARY_HASH_MISSING Install_PrefixStyle = 2
)

// Enum value maps for Install_PrefixStyle.
var (
	Install_PrefixStyle_name = map[int32]string{
		0: "ASIS",
		1: "REGEXP",
		2: "REGEXP_MAY_SKIP_BINARY_HASH_MISSING",
	}
	Install_PrefixStyle_value = map[string]int32{
		"ASIS":                                0,
		"REGEXP":                              1,
		"REGEXP_MAY_SKIP_BINARY_HASH_MISSING": 2,
	}
)

func (x Install_PrefixStyle) Enum() *Install_PrefixStyle {
	p := new(Install_PrefixStyle)
	*p = x
	return p
}

func (x Install_PrefixStyle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Install_PrefixStyle) Descriptor() protoreflect.EnumDescriptor {
	return file_command_setup_proto_enumTypes[0].Descriptor()
}

func (Install_PrefixStyle) Type() protoreflect.EnumType {
	return &file_command_setup_proto_enumTypes[0]
}

func (x Install_PrefixStyle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Install_PrefixStyle.Descriptor instead.
func (Install_PrefixStyle) EnumDescriptor() ([]byte, []int) {
	return file_command_setup_proto_rawDescGZIP(), []int{0, 0}
}

// Install describes which command to be installed.
// path is represented as server path (posix style),
// and converted to client path in CmdDescriptor.Setup.
//
// NEXT ID TO USE: 9
type Install struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// command key (e.g. "gcc", "clang++")
	// note that this value will be Selector name in command.proto.
	// in other words, this value must be the same with matching CommandSpec name
	// in api/goma_data.proto if the install is a compiler.
	// i.e. equivalent CompilerFlags::GetCompilerName.
	//
	// followings are known compiler keys (as of Jun 2017):
	// gcc, g++, clang, clang++, cl.exe, clang-cl, javac, ps3ppusnc.exe
	//
	// for subprogram etc, it should be basename
	// e.g. "libFindBadConstructs.so".
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// prefix dir (or toolchain root).
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// prefix_style represents how prefix need to be handled.
	PrefixStyle Install_PrefixStyle `protobuf:"varint,7,opt,name=prefix_style,json=prefixStyle,proto3,enum=command.Install_PrefixStyle" json:"prefix_style,omitempty"`
	// path in prefix dir.
	// actual binary should be found in "prefix/path".
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// glob for additional files (i.e. wrapper scripts, libs etc).
	// glob syntax is filepath.Match's.
	// relative to prefix dir.
	// if it ends with '/', all files under this dir will be added.
	Files       []string             `protobuf:"bytes,4,rep,name=files,proto3" json:"files,omitempty"`
	ForSelector *Install_ForSelector `protobuf:"bytes,5,opt,name=for_selector,json=forSelector,proto3" json:"for_selector,omitempty"`
	// clang_need_target is true if --target option is always needed
	// to run command.
	// https://clang.llvm.org/docs/CrossCompilation.html#target-triple
	// this can be also true if exec_server want to use client's raw target.
	// e.g. set --target x86_64-apple-darwin-10.8.0 sent by client.
	// note: this is clang/clang++ only.
	ClangNeedTarget bool `protobuf:"varint,6,opt,name=clang_need_target,json=clangNeedTarget,proto3" json:"clang_need_target,omitempty"`
	// windows_cross is true for windows cross compile on linux.
	WindowsCross bool `protobuf:"varint,8,opt,name=windows_cross,json=windowsCross,proto3" json:"windows_cross,omitempty"`
}

func (x *Install) Reset() {
	*x = Install{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_setup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Install) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Install) ProtoMessage() {}

func (x *Install) ProtoReflect() protoreflect.Message {
	mi := &file_command_setup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Install.ProtoReflect.Descriptor instead.
func (*Install) Descriptor() ([]byte, []int) {
	return file_command_setup_proto_rawDescGZIP(), []int{0}
}

func (x *Install) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Install) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *Install) GetPrefixStyle() Install_PrefixStyle {
	if x != nil {
		return x.PrefixStyle
	}
	return Install_ASIS
}

func (x *Install) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Install) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Install) GetForSelector() *Install_ForSelector {
	if x != nil {
		return x.ForSelector
	}
	return nil
}

func (x *Install) GetClangNeedTarget() bool {
	if x != nil {
		return x.ClangNeedTarget
	}
	return false
}

func (x *Install) GetWindowsCross() bool {
	if x != nil {
		return x.WindowsCross
	}
	return false
}

type Setup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Installs []*Install `protobuf:"bytes,1,rep,name=installs,proto3" json:"installs,omitempty"`
}

func (x *Setup) Reset() {
	*x = Setup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_setup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Setup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Setup) ProtoMessage() {}

func (x *Setup) ProtoReflect() protoreflect.Message {
	mi := &file_command_setup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Setup.ProtoReflect.Descriptor instead.
func (*Setup) Descriptor() ([]byte, []int) {
	return file_command_setup_proto_rawDescGZIP(), []int{1}
}

func (x *Setup) GetInstalls() []*Install {
	if x != nil {
		return x.Installs
	}
	return nil
}

// ForSelector is used for setting selector for cross compile.
// selector's name and version is set by compiler specified by
// prefix and path above.
// note that setup_cmd in cmd_server is responsible for setting version and
// equivalent hash, and setting cross option in CmdDescriptor.
// for compilers, the result of --version option must be the same between
// ForSelector and Install.
// for subprograms, the behavior must be the same between ForSelector and
// Install.
// ForSelector and Install are expected to be made from the same source code.
type Install_ForSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// target of this cross compile.
	// this is used for dispatching a compiler as a build target.
	// note that target must be normalized.
	// e.g. x86_64-darwin
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// binary_hash_from represents a filename to calculate hash when selector is
	// made.
	// e.g. darwin/third_party/llvm-build/Release+Asserts/bin/clang
	//
	// if prefix_style above is either of REGEXP or
	// REGEXP_MAY_SKIP_BINARY_HASH_MISSING, you can use regexp grouping variable
	// here.
	// e.g. if prefix_style is REGEXP, prefix is linux/clang-([0-9]+) and
	// found clang-12345, binary_hash_from darwin/clang-$1/bin/clang is
	// automatically converted to darwin/clang-12345/bin/clang when this
	// field is used.
	// if prefix_style is REGEXP_MAY_SKIP_BINARY_HASH_MISSING and
	// darwin/clang-12345/bin/clang does not exist, setup_cmd just skip
	// to proceed setup of the command without error.
	BinaryHashFrom string `protobuf:"bytes,2,opt,name=binary_hash_from,json=binaryHashFrom,proto3" json:"binary_hash_from,omitempty"`
}

func (x *Install_ForSelector) Reset() {
	*x = Install_ForSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_setup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Install_ForSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Install_ForSelector) ProtoMessage() {}

func (x *Install_ForSelector) ProtoReflect() protoreflect.Message {
	mi := &file_command_setup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Install_ForSelector.ProtoReflect.Descriptor instead.
func (*Install_ForSelector) Descriptor() ([]byte, []int) {
	return file_command_setup_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Install_ForSelector) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Install_ForSelector) GetBinaryHashFrom() string {
	if x != nil {
		return x.BinaryHashFrom
	}
	return ""
}

var File_command_setup_proto protoreflect.FileDescriptor

var file_command_setup_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0xcf,
	0x03, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x3f, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x5f, 0x73,
	0x74, 0x79, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x2e, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12,
	0x3f, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x2e, 0x46, 0x6f, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x6c, 0x61,
	0x6e, 0x67, 0x4e, 0x65, 0x65, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x5f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x43, 0x72, 0x6f, 0x73,
	0x73, 0x1a, 0x4f, 0x0a, 0x0b, 0x46, 0x6f, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x48, 0x61, 0x73, 0x68, 0x46, 0x72,
	0x6f, 0x6d, 0x22, 0x4c, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x53, 0x74, 0x79, 0x6c,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x53, 0x49, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x45, 0x47, 0x45, 0x58, 0x50, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x52, 0x45, 0x47, 0x45, 0x58,
	0x50, 0x5f, 0x4d, 0x41, 0x59, 0x5f, 0x53, 0x4b, 0x49, 0x50, 0x5f, 0x42, 0x49, 0x4e, 0x41, 0x52,
	0x59, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x22, 0x48, 0x0a, 0x05, 0x53, 0x65, 0x74, 0x75, 0x70, 0x12, 0x2c, 0x0a, 0x08, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x52, 0x08, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x73, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x0b, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x6f,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x6f,
	0x6d, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_setup_proto_rawDescOnce sync.Once
	file_command_setup_proto_rawDescData = file_command_setup_proto_rawDesc
)

func file_command_setup_proto_rawDescGZIP() []byte {
	file_command_setup_proto_rawDescOnce.Do(func() {
		file_command_setup_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_setup_proto_rawDescData)
	})
	return file_command_setup_proto_rawDescData
}

var file_command_setup_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_command_setup_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_command_setup_proto_goTypes = []interface{}{
	(Install_PrefixStyle)(0),    // 0: command.Install.PrefixStyle
	(*Install)(nil),             // 1: command.Install
	(*Setup)(nil),               // 2: command.Setup
	(*Install_ForSelector)(nil), // 3: command.Install.ForSelector
}
var file_command_setup_proto_depIdxs = []int32{
	0, // 0: command.Install.prefix_style:type_name -> command.Install.PrefixStyle
	3, // 1: command.Install.for_selector:type_name -> command.Install.ForSelector
	1, // 2: command.Setup.installs:type_name -> command.Install
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_command_setup_proto_init() }
func file_command_setup_proto_init() {
	if File_command_setup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_command_setup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Install); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_setup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Setup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_setup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Install_ForSelector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_command_setup_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_command_setup_proto_goTypes,
		DependencyIndexes: file_command_setup_proto_depIdxs,
		EnumInfos:         file_command_setup_proto_enumTypes,
		MessageInfos:      file_command_setup_proto_msgTypes,
	}.Build()
	File_command_setup_proto = out.File
	file_command_setup_proto_rawDesc = nil
	file_command_setup_proto_goTypes = nil
	file_command_setup_proto_depIdxs = nil
}
