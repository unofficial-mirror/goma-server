// Copyright 2018 The Goma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: auth/acl.proto

package auth

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

// Group defines a group of users that shares the same service account.
// Different groups may share a service account.
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// if audience is not empty, group member should have this audience
	// in the token.
	Audience string `protobuf:"bytes,3,opt,name=audience,proto3" json:"audience,omitempty"`
	// matches mail in the token.
	Emails []string `protobuf:"bytes,4,rep,name=emails,proto3" json:"emails,omitempty"`
	// matches domain part of email in the token.
	Domains []string `protobuf:"bytes,5,rep,name=domains,proto3" json:"domains,omitempty"`
	// use this service account for this group.
	// it is used to read json in service-accounts volume.
	// If service_account is "default", use server's default service account.
	// If service_account is empty, pass EUC as is.
	ServiceAccount string `protobuf:"bytes,6,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// If reject is true, deny access from this group.
	Reject bool `protobuf:"varint,7,opt,name=reject,proto3" json:"reject,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_acl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_auth_acl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_auth_acl_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Group) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

func (x *Group) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *Group) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *Group) GetServiceAccount() string {
	if x != nil {
		return x.ServiceAccount
	}
	return ""
}

func (x *Group) GetReject() bool {
	if x != nil {
		return x.Reject
	}
	return false
}

type ACL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// First matched group will be used.
	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *ACL) Reset() {
	*x = ACL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_acl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACL) ProtoMessage() {}

func (x *ACL) ProtoReflect() protoreflect.Message {
	mi := &file_auth_acl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACL.ProtoReflect.Descriptor instead.
func (*ACL) Descriptor() ([]byte, []int) {
	return file_auth_acl_proto_rawDescGZIP(), []int{1}
}

func (x *ACL) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_auth_acl_proto protoreflect.FileDescriptor

var file_auth_acl_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0xc8, 0x01, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x22, 0x2a, 0x0a, 0x03, 0x41, 0x43, 0x4c, 0x12, 0x23, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x67, 0x6f, 0x6d, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_acl_proto_rawDescOnce sync.Once
	file_auth_acl_proto_rawDescData = file_auth_acl_proto_rawDesc
)

func file_auth_acl_proto_rawDescGZIP() []byte {
	file_auth_acl_proto_rawDescOnce.Do(func() {
		file_auth_acl_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_acl_proto_rawDescData)
	})
	return file_auth_acl_proto_rawDescData
}

var file_auth_acl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_acl_proto_goTypes = []interface{}{
	(*Group)(nil), // 0: auth.Group
	(*ACL)(nil),   // 1: auth.ACL
}
var file_auth_acl_proto_depIdxs = []int32{
	0, // 0: auth.ACL.groups:type_name -> auth.Group
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_acl_proto_init() }
func file_auth_acl_proto_init() {
	if File_auth_acl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_acl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_auth_acl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACL); i {
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
			RawDescriptor: file_auth_acl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_acl_proto_goTypes,
		DependencyIndexes: file_auth_acl_proto_depIdxs,
		MessageInfos:      file_auth_acl_proto_msgTypes,
	}.Build()
	File_auth_acl_proto = out.File
	file_auth_acl_proto_rawDesc = nil
	file_auth_acl_proto_goTypes = nil
	file_auth_acl_proto_depIdxs = nil
}
