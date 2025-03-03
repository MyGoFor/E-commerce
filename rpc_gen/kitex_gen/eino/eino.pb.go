// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: eino.proto

package eino

import (
	context "context"
	order "github.com/MyGoFor/E-commerce/rpc_gen/kitex_gen/order"
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

type SearchOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question string `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *SearchOrderReq) Reset() {
	*x = SearchOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eino_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchOrderReq) ProtoMessage() {}

func (x *SearchOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_eino_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchOrderReq.ProtoReflect.Descriptor instead.
func (*SearchOrderReq) Descriptor() ([]byte, []int) {
	return file_eino_proto_rawDescGZIP(), []int{0}
}

func (x *SearchOrderReq) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

type SearchOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order []*order.Order `protobuf:"bytes,1,rep,name=order,proto3" json:"order,omitempty"`
}

func (x *SearchOrderResp) Reset() {
	*x = SearchOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eino_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchOrderResp) ProtoMessage() {}

func (x *SearchOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_eino_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchOrderResp.ProtoReflect.Descriptor instead.
func (*SearchOrderResp) Descriptor() ([]byte, []int) {
	return file_eino_proto_rawDescGZIP(), []int{1}
}

func (x *SearchOrderResp) GetOrder() []*order.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type PlaceOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int32  `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Question string `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *PlaceOrderReq) Reset() {
	*x = PlaceOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eino_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderReq) ProtoMessage() {}

func (x *PlaceOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_eino_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderReq.ProtoReflect.Descriptor instead.
func (*PlaceOrderReq) Descriptor() ([]byte, []int) {
	return file_eino_proto_rawDescGZIP(), []int{2}
}

func (x *PlaceOrderReq) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PlaceOrderReq) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

type PlaceOrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp string `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
}

func (x *PlaceOrderResp) Reset() {
	*x = PlaceOrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eino_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderResp) ProtoMessage() {}

func (x *PlaceOrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_eino_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderResp.ProtoReflect.Descriptor instead.
func (*PlaceOrderResp) Descriptor() ([]byte, []int) {
	return file_eino_proto_rawDescGZIP(), []int{3}
}

func (x *PlaceOrderResp) GetResp() string {
	if x != nil {
		return x.Resp
	}
	return ""
}

var File_eino_proto protoreflect.FileDescriptor

var file_eino_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x69,
	0x6e, 0x6f, 0x1a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2c, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a,
	0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x22, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70, 0x32, 0x86, 0x01, 0x0a, 0x0b, 0x45, 0x69,
	0x6e, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x65, 0x69, 0x6e, 0x6f, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x15,
	0x2e, 0x65, 0x69, 0x6e, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x65, 0x69, 0x6e, 0x6f, 0x2e, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x65, 0x69, 0x6e,
	0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4d, 0x79, 0x47, 0x6f, 0x46, 0x6f, 0x72, 0x2f, 0x45, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6b, 0x69, 0x74, 0x65,
	0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x69, 0x6e, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eino_proto_rawDescOnce sync.Once
	file_eino_proto_rawDescData = file_eino_proto_rawDesc
)

func file_eino_proto_rawDescGZIP() []byte {
	file_eino_proto_rawDescOnce.Do(func() {
		file_eino_proto_rawDescData = protoimpl.X.CompressGZIP(file_eino_proto_rawDescData)
	})
	return file_eino_proto_rawDescData
}

var file_eino_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eino_proto_goTypes = []interface{}{
	(*SearchOrderReq)(nil),  // 0: eino.SearchOrderReq
	(*SearchOrderResp)(nil), // 1: eino.SearchOrderResp
	(*PlaceOrderReq)(nil),   // 2: eino.PlaceOrderReq
	(*PlaceOrderResp)(nil),  // 3: eino.PlaceOrderResp
	(*order.Order)(nil),     // 4: order.Order
}
var file_eino_proto_depIdxs = []int32{
	4, // 0: eino.SearchOrderResp.order:type_name -> order.Order
	0, // 1: eino.EinoService.SearchOrder:input_type -> eino.SearchOrderReq
	2, // 2: eino.EinoService.PlaceOrder:input_type -> eino.PlaceOrderReq
	1, // 3: eino.EinoService.SearchOrder:output_type -> eino.SearchOrderResp
	3, // 4: eino.EinoService.PlaceOrder:output_type -> eino.PlaceOrderResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eino_proto_init() }
func file_eino_proto_init() {
	if File_eino_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eino_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchOrderReq); i {
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
		file_eino_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchOrderResp); i {
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
		file_eino_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderReq); i {
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
		file_eino_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderResp); i {
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
			RawDescriptor: file_eino_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eino_proto_goTypes,
		DependencyIndexes: file_eino_proto_depIdxs,
		MessageInfos:      file_eino_proto_msgTypes,
	}.Build()
	File_eino_proto = out.File
	file_eino_proto_rawDesc = nil
	file_eino_proto_goTypes = nil
	file_eino_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.1. DO NOT EDIT.

type EinoService interface {
	SearchOrder(ctx context.Context, req *SearchOrderReq) (res *SearchOrderResp, err error)
	PlaceOrder(ctx context.Context, req *PlaceOrderReq) (res *PlaceOrderResp, err error)
}
