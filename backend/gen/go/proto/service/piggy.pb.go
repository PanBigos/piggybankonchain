// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/service/piggy.proto

package service

import (
	v1 "github.com/Exca-DK/pegism/gen/go/proto/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_service_piggy_proto protoreflect.FileDescriptor

var file_proto_service_piggy_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x69, 0x67, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x65, 0x67,
	0x69, 0x73, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x69, 0x67, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd8, 0x03, 0x0a,
	0x0c, 0x50, 0x69, 0x67, 0x67, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x75, 0x0a,
	0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x69, 0x67, 0x67, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x2e, 0x70, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x69, 0x67, 0x67, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x69, 0x67, 0x67, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22,
	0x13, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x67, 0x67, 0x79, 0x2f, 0x7b, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x7d, 0x12, 0x60, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x69, 0x67, 0x67, 0x79,
	0x12, 0x1a, 0x2e, 0x70, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x69, 0x67, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x67, 0x67,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x67, 0x67, 0x79, 0x2f, 0x7b, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x12, 0x73, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x69, 0x67,
	0x67, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x2e,
	0x70, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x67,
	0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x65, 0x67, 0x69,
	0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x67, 0x67, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x67, 0x67, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2f, 0x7b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x7d, 0x12, 0x7a, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x69, 0x67, 0x67, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x2e, 0x70, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x69, 0x67, 0x67, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x69, 0x67, 0x67, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x12, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x69, 0x67, 0x67, 0x79, 0x2f, 0x6e, 0x61, 0x6d, 0x65,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0xa9, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0a,
	0x50, 0x69, 0x67, 0x67, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x78, 0x63, 0x61, 0x2d, 0x44, 0x4b,
	0x2f, 0x70, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x03, 0x50,
	0x53, 0x58, 0xaa, 0x02, 0x0e, 0x50, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0xca, 0x02, 0x0e, 0x50, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x1a, 0x50, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0f, 0x50, 0x65, 0x67, 0x69, 0x73, 0x6d, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_service_piggy_proto_goTypes = []interface{}{
	(*v1.UpdatePiggyNameRequest)(nil),   // 0: pegism.v1.UpdatePiggyNameRequest
	(*v1.GetPiggyRequest)(nil),          // 1: pegism.v1.GetPiggyRequest
	(*v1.GetPiggyFromNameRequest)(nil),  // 2: pegism.v1.GetPiggyFromNameRequest
	(*v1.UpdatePiggyNameResponse)(nil),  // 3: pegism.v1.UpdatePiggyNameResponse
	(*v1.GetPiggyResponse)(nil),         // 4: pegism.v1.GetPiggyResponse
	(*v1.GetPiggyFromNameResponse)(nil), // 5: pegism.v1.GetPiggyFromNameResponse
}
var file_proto_service_piggy_proto_depIdxs = []int32{
	0, // 0: pegism.service.PiggyService.UpdatePiggyName:input_type -> pegism.v1.UpdatePiggyNameRequest
	1, // 1: pegism.service.PiggyService.GetPiggy:input_type -> pegism.v1.GetPiggyRequest
	1, // 2: pegism.service.PiggyService.GetPiggyFromProfile:input_type -> pegism.v1.GetPiggyRequest
	2, // 3: pegism.service.PiggyService.GetPiggyFromName:input_type -> pegism.v1.GetPiggyFromNameRequest
	3, // 4: pegism.service.PiggyService.UpdatePiggyName:output_type -> pegism.v1.UpdatePiggyNameResponse
	4, // 5: pegism.service.PiggyService.GetPiggy:output_type -> pegism.v1.GetPiggyResponse
	4, // 6: pegism.service.PiggyService.GetPiggyFromProfile:output_type -> pegism.v1.GetPiggyResponse
	5, // 7: pegism.service.PiggyService.GetPiggyFromName:output_type -> pegism.v1.GetPiggyFromNameResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_service_piggy_proto_init() }
func file_proto_service_piggy_proto_init() {
	if File_proto_service_piggy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_service_piggy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_service_piggy_proto_goTypes,
		DependencyIndexes: file_proto_service_piggy_proto_depIdxs,
	}.Build()
	File_proto_service_piggy_proto = out.File
	file_proto_service_piggy_proto_rawDesc = nil
	file_proto_service_piggy_proto_goTypes = nil
	file_proto_service_piggy_proto_depIdxs = nil
}
