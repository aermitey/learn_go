// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: types.proto

package apis

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

type PersonalInformationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*PersonalInformation `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *PersonalInformationList) Reset() {
	*x = PersonalInformationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonalInformationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalInformationList) ProtoMessage() {}

func (x *PersonalInformationList) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalInformationList.ProtoReflect.Descriptor instead.
func (*PersonalInformationList) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *PersonalInformationList) GetItems() []*PersonalInformation {
	if x != nil {
		return x.Items
	}
	return nil
}

type PersonalInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: gorm:"primaryKey;column:id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;column:id"`
	// @gotags: gorm:"column:name"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" gorm:"column:name"`
	// @gotags: gorm:"column:sex"
	Sex string `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex,omitempty" gorm:"column:sex"`
	// @gotags: gorm:"column:tall"
	Tall float32 `protobuf:"fixed32,4,opt,name=tall,proto3" json:"tall,omitempty" gorm:"column:tall"`
	// @gotags: gorm:"column:weight"
	Weight float32 `protobuf:"fixed32,5,opt,name=weight,proto3" json:"weight,omitempty" gorm:"column:weight"`
	// @gotags: gorm:"column:age"
	Age int64 `protobuf:"varint,6,opt,name=age,proto3" json:"age,omitempty" gorm:"column:age"`
}

func (x *PersonalInformation) Reset() {
	*x = PersonalInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonalInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalInformation) ProtoMessage() {}

func (x *PersonalInformation) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalInformation.ProtoReflect.Descriptor instead.
func (*PersonalInformation) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *PersonalInformation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PersonalInformation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonalInformation) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *PersonalInformation) GetTall() float32 {
	if x != nil {
		return x.Tall
	}
	return 0
}

func (x *PersonalInformation) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *PersonalInformation) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type FriendCircle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: gorm:"primaryKey;column:id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;column:id"`
	// @gotags: gorm:"column:personId"
	PersonId int64 `protobuf:"varint,2,opt,name=personId,proto3" json:"personId,omitempty" gorm:"column:personId"`
	// @gotags: gorm:"column:name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" gorm:"column:name"`
	// @gotags: gorm:"column:sex"
	Sex string `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty" gorm:"column:sex"`
	// @gotags: gorm:"column:tall"
	Tall float32 `protobuf:"fixed32,5,opt,name=tall,proto3" json:"tall,omitempty" gorm:"column:tall"`
	// @gotags: gorm:"column:weight"
	Weight float32 `protobuf:"fixed32,6,opt,name=weight,proto3" json:"weight,omitempty" gorm:"column:weight"`
	// @gotags: gorm:"column:age"
	Age int64 `protobuf:"varint,7,opt,name=age,proto3" json:"age,omitempty" gorm:"column:age"`
	// @gotags: gorm:"column:fatRate"
	FatRate float32 `protobuf:"fixed32,8,opt,name=fatRate,proto3" json:"fatRate,omitempty" gorm:"column:fatRate"`
	// @gotags: gorm:"column:content"
	Content string `protobuf:"bytes,9,opt,name=content,proto3" json:"content,omitempty" gorm:"column:content"`
	// @gotags: gorm:"column:createTime"
	CreateTime string `protobuf:"bytes,10,opt,name=createTime,proto3" json:"createTime,omitempty" gorm:"column:createTime"`
	// @gotags: gorm:"column:isDeleted"
	IsDeleted int64 `protobuf:"varint,11,opt,name=isDeleted,proto3" json:"isDeleted,omitempty" gorm:"column:isDeleted"`
}

func (x *FriendCircle) Reset() {
	*x = FriendCircle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendCircle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendCircle) ProtoMessage() {}

func (x *FriendCircle) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendCircle.ProtoReflect.Descriptor instead.
func (*FriendCircle) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

func (x *FriendCircle) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FriendCircle) GetPersonId() int64 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *FriendCircle) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FriendCircle) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *FriendCircle) GetTall() float32 {
	if x != nil {
		return x.Tall
	}
	return 0
}

func (x *FriendCircle) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *FriendCircle) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *FriendCircle) GetFatRate() float32 {
	if x != nil {
		return x.FatRate
	}
	return 0
}

func (x *FriendCircle) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *FriendCircle) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *FriendCircle) GetIsDeleted() int64 {
	if x != nil {
		return x.IsDeleted
	}
	return 0
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a,
	0x17, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x22, 0x90, 0x02, 0x0a, 0x0c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x69, 0x72, 0x63, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61, 0x74, 0x52, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x66, 0x61, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_types_proto_goTypes = []interface{}{
	(*PersonalInformationList)(nil), // 0: PersonalInformationList
	(*PersonalInformation)(nil),     // 1: PersonalInformation
	(*FriendCircle)(nil),            // 2: FriendCircle
}
var file_types_proto_depIdxs = []int32{
	1, // 0: PersonalInformationList.items:type_name -> PersonalInformation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonalInformationList); i {
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
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonalInformation); i {
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
		file_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendCircle); i {
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
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
