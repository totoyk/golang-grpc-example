// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/pet.proto

package pet

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Pet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Species string `protobuf:"bytes,2,opt,name=species,proto3" json:"species,omitempty"`
	Breed   string `protobuf:"bytes,3,opt,name=breed,proto3" json:"breed,omitempty"`
	Color   string `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	Age     string `protobuf:"bytes,5,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *Pet) Reset() {
	*x = Pet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pet) ProtoMessage() {}

func (x *Pet) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pet.ProtoReflect.Descriptor instead.
func (*Pet) Descriptor() ([]byte, []int) {
	return file_proto_pet_proto_rawDescGZIP(), []int{0}
}

func (x *Pet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pet) GetSpecies() string {
	if x != nil {
		return x.Species
	}
	return ""
}

func (x *Pet) GetBreed() string {
	if x != nil {
		return x.Breed
	}
	return ""
}

func (x *Pet) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Pet) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

type GetPetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Species string `protobuf:"bytes,2,opt,name=species,proto3" json:"species,omitempty"`
	Breed   string `protobuf:"bytes,3,opt,name=breed,proto3" json:"breed,omitempty"`
	Color   string `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	Age     string `protobuf:"bytes,5,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *GetPetReply) Reset() {
	*x = GetPetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetReply) ProtoMessage() {}

func (x *GetPetReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetReply.ProtoReflect.Descriptor instead.
func (*GetPetReply) Descriptor() ([]byte, []int) {
	return file_proto_pet_proto_rawDescGZIP(), []int{1}
}

func (x *GetPetReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPetReply) GetSpecies() string {
	if x != nil {
		return x.Species
	}
	return ""
}

func (x *GetPetReply) GetBreed() string {
	if x != nil {
		return x.Breed
	}
	return ""
}

func (x *GetPetReply) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *GetPetReply) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

var File_proto_pet_proto protoreflect.FileDescriptor

var file_proto_pet_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x70, 0x65, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x03, 0x50, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x65, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x65, 0x65, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x79, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x65, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x32, 0x42, 0x0a, 0x0a, 0x50, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x34, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x70, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x74, 0x6f, 0x79, 0x6b, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x72, 0x65, 0x63, 0x61, 0x70, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pet_proto_rawDescOnce sync.Once
	file_proto_pet_proto_rawDescData = file_proto_pet_proto_rawDesc
)

func file_proto_pet_proto_rawDescGZIP() []byte {
	file_proto_pet_proto_rawDescOnce.Do(func() {
		file_proto_pet_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pet_proto_rawDescData)
	})
	return file_proto_pet_proto_rawDescData
}

var file_proto_pet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_pet_proto_goTypes = []interface{}{
	(*Pet)(nil),           // 0: pet.Pet
	(*GetPetReply)(nil),   // 1: pet.GetPetReply
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_proto_pet_proto_depIdxs = []int32{
	2, // 0: pet.PetService.GetMyPet:input_type -> google.protobuf.Empty
	1, // 1: pet.PetService.GetMyPet:output_type -> pet.GetPetReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_pet_proto_init() }
func file_proto_pet_proto_init() {
	if File_proto_pet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pet); i {
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
		file_proto_pet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPetReply); i {
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
			RawDescriptor: file_proto_pet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pet_proto_goTypes,
		DependencyIndexes: file_proto_pet_proto_depIdxs,
		MessageInfos:      file_proto_pet_proto_msgTypes,
	}.Build()
	File_proto_pet_proto = out.File
	file_proto_pet_proto_rawDesc = nil
	file_proto_pet_proto_goTypes = nil
	file_proto_pet_proto_depIdxs = nil
}
