// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: homeit.proto

package homeit

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homeit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_homeit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_homeit_proto_rawDescGZIP(), []int{0}
}

type Supplies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Supplies []*Supply `protobuf:"bytes,1,rep,name=supplies,proto3" json:"supplies,omitempty"`
}

func (x *Supplies) Reset() {
	*x = Supplies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homeit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Supplies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supplies) ProtoMessage() {}

func (x *Supplies) ProtoReflect() protoreflect.Message {
	mi := &file_homeit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Supplies.ProtoReflect.Descriptor instead.
func (*Supplies) Descriptor() ([]byte, []int) {
	return file_homeit_proto_rawDescGZIP(), []int{1}
}

func (x *Supplies) GetSupplies() []*Supply {
	if x != nil {
		return x.Supplies
	}
	return nil
}

type Supply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created  string `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated  string `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	UserId   string `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	Name     string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Priority int32  `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *Supply) Reset() {
	*x = Supply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_homeit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Supply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supply) ProtoMessage() {}

func (x *Supply) ProtoReflect() protoreflect.Message {
	mi := &file_homeit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Supply.ProtoReflect.Descriptor instead.
func (*Supply) Descriptor() ([]byte, []int) {
	return file_homeit_proto_rawDescGZIP(), []int{2}
}

func (x *Supply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Supply) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *Supply) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *Supply) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Supply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Supply) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

var File_homeit_proto protoreflect.FileDescriptor

var file_homeit_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x68, 0x6f, 0x6d, 0x65, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x35,
	0x0a, 0x08, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x08, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x08, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x06, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x32, 0x3d, 0x0a, 0x0b,
	0x46, 0x6f, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x00, 0x42, 0x4a, 0x0a, 0x14, 0x69,
	0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x42, 0x0b, 0x48, 0x6f, 0x6d, 0x65, 0x49, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x70, 0x69, 0x6f, 0x72, 0x6f, 0x77, 0x73, 0x6b, 0x69, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_homeit_proto_rawDescOnce sync.Once
	file_homeit_proto_rawDescData = file_homeit_proto_rawDesc
)

func file_homeit_proto_rawDescGZIP() []byte {
	file_homeit_proto_rawDescOnce.Do(func() {
		file_homeit_proto_rawDescData = protoimpl.X.CompressGZIP(file_homeit_proto_rawDescData)
	})
	return file_homeit_proto_rawDescData
}

var file_homeit_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_homeit_proto_goTypes = []interface{}{
	(*Empty)(nil),    // 0: proto.Empty
	(*Supplies)(nil), // 1: proto.Supplies
	(*Supply)(nil),   // 2: proto.Supply
}
var file_homeit_proto_depIdxs = []int32{
	2, // 0: proto.Supplies.supplies:type_name -> proto.Supply
	0, // 1: proto.FoodService.GetSupplies:input_type -> proto.Empty
	1, // 2: proto.FoodService.GetSupplies:output_type -> proto.Supplies
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_homeit_proto_init() }
func file_homeit_proto_init() {
	if File_homeit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_homeit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_homeit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Supplies); i {
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
		file_homeit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Supply); i {
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
			RawDescriptor: file_homeit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_homeit_proto_goTypes,
		DependencyIndexes: file_homeit_proto_depIdxs,
		MessageInfos:      file_homeit_proto_msgTypes,
	}.Build()
	File_homeit_proto = out.File
	file_homeit_proto_rawDesc = nil
	file_homeit_proto_goTypes = nil
	file_homeit_proto_depIdxs = nil
}