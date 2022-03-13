// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: iamnande/cardmod/item/v1/api.proto

package itemv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Request schema for the GetItem request.
type GetItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the item to get.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetItemRequest) Reset() {
	*x = GetItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamnande_cardmod_item_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemRequest) ProtoMessage() {}

func (x *GetItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iamnande_cardmod_item_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemRequest.ProtoReflect.Descriptor instead.
func (*GetItemRequest) Descriptor() ([]byte, []int) {
	return file_iamnande_cardmod_item_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetItemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request schema for the ListItems request.
type ListItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListItemsRequest) Reset() {
	*x = ListItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamnande_cardmod_item_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsRequest) ProtoMessage() {}

func (x *ListItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iamnande_cardmod_item_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsRequest.ProtoReflect.Descriptor instead.
func (*ListItemsRequest) Descriptor() ([]byte, []int) {
	return file_iamnande_cardmod_item_v1_api_proto_rawDescGZIP(), []int{1}
}

// Response schema for the ListItems request.
type ListItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The collection of items
	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListItemsResponse) Reset() {
	*x = ListItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamnande_cardmod_item_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsResponse) ProtoMessage() {}

func (x *ListItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iamnande_cardmod_item_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsResponse.ProtoReflect.Descriptor instead.
func (*ListItemsResponse) Descriptor() ([]byte, []int) {
	return file_iamnande_cardmod_item_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *ListItemsResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_iamnande_cardmod_item_v1_api_proto protoreflect.FileDescriptor

var file_iamnande_cardmod_item_v1_api_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x6d,
	0x6f, 0x64, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2e, 0x63,
	0x61, 0x72, 0x64, 0x6d, 0x6f, 0x64, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x23,
	0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x6d, 0x6f, 0x64,
	0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xfa, 0x42,
	0x12, 0x72, 0x10, 0x10, 0x03, 0x18, 0x19, 0x32, 0x0a, 0x5e, 0x5b, 0x2d, 0x2c, 0x20, 0x5c, 0x77,
	0x5d, 0x2b, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2e, 0x63, 0x61, 0x72,
	0x64, 0x6d, 0x6f, 0x64, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xc4, 0x01, 0x0a, 0x07, 0x49, 0x74, 0x65,
	0x6d, 0x41, 0x50, 0x49, 0x12, 0x53, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x28, 0x2e, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x6d,
	0x6f, 0x64, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x61, 0x6d, 0x6e,
	0x61, 0x6e, 0x64, 0x65, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x6d, 0x6f, 0x64, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x64, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2a, 0x2e, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64,
	0x65, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x6d, 0x6f, 0x64, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x69, 0x61, 0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2e, 0x63, 0x61,
	0x72, 0x64, 0x6d, 0x6f, 0x64, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x61,
	0x6d, 0x6e, 0x61, 0x6e, 0x64, 0x65, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x6d, 0x6f, 0x64, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x76, 0x31, 0x3b, 0x69, 0x74,
	0x65, 0x6d, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iamnande_cardmod_item_v1_api_proto_rawDescOnce sync.Once
	file_iamnande_cardmod_item_v1_api_proto_rawDescData = file_iamnande_cardmod_item_v1_api_proto_rawDesc
)

func file_iamnande_cardmod_item_v1_api_proto_rawDescGZIP() []byte {
	file_iamnande_cardmod_item_v1_api_proto_rawDescOnce.Do(func() {
		file_iamnande_cardmod_item_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_iamnande_cardmod_item_v1_api_proto_rawDescData)
	})
	return file_iamnande_cardmod_item_v1_api_proto_rawDescData
}

var file_iamnande_cardmod_item_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_iamnande_cardmod_item_v1_api_proto_goTypes = []interface{}{
	(*GetItemRequest)(nil),    // 0: iamnande.cardmod.item.v1.GetItemRequest
	(*ListItemsRequest)(nil),  // 1: iamnande.cardmod.item.v1.ListItemsRequest
	(*ListItemsResponse)(nil), // 2: iamnande.cardmod.item.v1.ListItemsResponse
	(*Item)(nil),              // 3: iamnande.cardmod.item.v1.Item
}
var file_iamnande_cardmod_item_v1_api_proto_depIdxs = []int32{
	3, // 0: iamnande.cardmod.item.v1.ListItemsResponse.items:type_name -> iamnande.cardmod.item.v1.Item
	0, // 1: iamnande.cardmod.item.v1.ItemAPI.GetItem:input_type -> iamnande.cardmod.item.v1.GetItemRequest
	1, // 2: iamnande.cardmod.item.v1.ItemAPI.ListItems:input_type -> iamnande.cardmod.item.v1.ListItemsRequest
	3, // 3: iamnande.cardmod.item.v1.ItemAPI.GetItem:output_type -> iamnande.cardmod.item.v1.Item
	2, // 4: iamnande.cardmod.item.v1.ItemAPI.ListItems:output_type -> iamnande.cardmod.item.v1.ListItemsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_iamnande_cardmod_item_v1_api_proto_init() }
func file_iamnande_cardmod_item_v1_api_proto_init() {
	if File_iamnande_cardmod_item_v1_api_proto != nil {
		return
	}
	file_iamnande_cardmod_item_v1_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_iamnande_cardmod_item_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemRequest); i {
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
		file_iamnande_cardmod_item_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsRequest); i {
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
		file_iamnande_cardmod_item_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsResponse); i {
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
			RawDescriptor: file_iamnande_cardmod_item_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iamnande_cardmod_item_v1_api_proto_goTypes,
		DependencyIndexes: file_iamnande_cardmod_item_v1_api_proto_depIdxs,
		MessageInfos:      file_iamnande_cardmod_item_v1_api_proto_msgTypes,
	}.Build()
	File_iamnande_cardmod_item_v1_api_proto = out.File
	file_iamnande_cardmod_item_v1_api_proto_rawDesc = nil
	file_iamnande_cardmod_item_v1_api_proto_goTypes = nil
	file_iamnande_cardmod_item_v1_api_proto_depIdxs = nil
}
