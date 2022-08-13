// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: Protos/RestaurantProtos/RestaurantProtos.proto

package RestaurantProtos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	OrderProtos "github.com/jainam240101/zomato-clone/Protos/OrderProtos"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_Protos_RestaurantProtos_RestaurantProtos_proto protoreflect.FileDescriptor

var file_Protos_RestaurantProtos_RestaurantProtos_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x42, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x0e, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_Protos_RestaurantProtos_RestaurantProtos_proto_goTypes = []interface{}{
	(*OrderProtos.OrderResponse)(nil), // 0: OrderResponse
}
var file_Protos_RestaurantProtos_RestaurantProtos_proto_depIdxs = []int32{
	0, // 0: RestaurantService.AcceptOrder:input_type -> OrderResponse
	0, // 1: RestaurantService.AcceptOrder:output_type -> OrderResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Protos_RestaurantProtos_RestaurantProtos_proto_init() }
func file_Protos_RestaurantProtos_RestaurantProtos_proto_init() {
	if File_Protos_RestaurantProtos_RestaurantProtos_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Protos_RestaurantProtos_RestaurantProtos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Protos_RestaurantProtos_RestaurantProtos_proto_goTypes,
		DependencyIndexes: file_Protos_RestaurantProtos_RestaurantProtos_proto_depIdxs,
	}.Build()
	File_Protos_RestaurantProtos_RestaurantProtos_proto = out.File
	file_Protos_RestaurantProtos_RestaurantProtos_proto_rawDesc = nil
	file_Protos_RestaurantProtos_RestaurantProtos_proto_goTypes = nil
	file_Protos_RestaurantProtos_RestaurantProtos_proto_depIdxs = nil
}
