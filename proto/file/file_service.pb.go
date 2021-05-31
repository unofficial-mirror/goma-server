// Copyright 2010 The Goma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//
// goma file service API

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: file/file_service.proto

package file

import (
	api "go.chromium.org/goma/server/proto/api"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_file_file_service_proto protoreflect.FileDescriptor

var file_file_file_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x65, 0x76, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x5f, 0x67, 0x6f, 0x6d, 0x61, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f,
	0x6d, 0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa4, 0x01,
	0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a,
	0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x5f, 0x67, 0x6f, 0x6d, 0x61, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x5f, 0x67, 0x6f, 0x6d, 0x61, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0a, 0x4c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x5f, 0x67, 0x6f, 0x6d, 0x61, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x5f, 0x67,
	0x6f, 0x6d, 0x61, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x26, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x6f, 0x6d, 0x61, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x80, 0x01,
	0x00, 0x88, 0x01, 0x00, 0x90, 0x01, 0x00,
}

var file_file_file_service_proto_goTypes = []interface{}{
	(*api.StoreFileReq)(nil),   // 0: devtools_goma.StoreFileReq
	(*api.LookupFileReq)(nil),  // 1: devtools_goma.LookupFileReq
	(*api.StoreFileResp)(nil),  // 2: devtools_goma.StoreFileResp
	(*api.LookupFileResp)(nil), // 3: devtools_goma.LookupFileResp
}
var file_file_file_service_proto_depIdxs = []int32{
	0, // 0: devtools_goma.FileService.StoreFile:input_type -> devtools_goma.StoreFileReq
	1, // 1: devtools_goma.FileService.LookupFile:input_type -> devtools_goma.LookupFileReq
	2, // 2: devtools_goma.FileService.StoreFile:output_type -> devtools_goma.StoreFileResp
	3, // 3: devtools_goma.FileService.LookupFile:output_type -> devtools_goma.LookupFileResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_file_service_proto_init() }
func file_file_file_service_proto_init() {
	if File_file_file_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file_file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_file_service_proto_goTypes,
		DependencyIndexes: file_file_file_service_proto_depIdxs,
	}.Build()
	File_file_file_service_proto = out.File
	file_file_file_service_proto_rawDesc = nil
	file_file_file_service_proto_goTypes = nil
	file_file_file_service_proto_depIdxs = nil
}
