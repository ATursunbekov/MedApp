// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/anamnesis.proto

package anamnesis

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SaveSessionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Timestamp     string                 `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Notes         string                 `protobuf:"bytes,3,opt,name=notes,proto3" json:"notes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveSessionRequest) Reset() {
	*x = SaveSessionRequest{}
	mi := &file_proto_anamnesis_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveSessionRequest) ProtoMessage() {}

func (x *SaveSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_anamnesis_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveSessionRequest.ProtoReflect.Descriptor instead.
func (*SaveSessionRequest) Descriptor() ([]byte, []int) {
	return file_proto_anamnesis_proto_rawDescGZIP(), []int{0}
}

func (x *SaveSessionRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SaveSessionRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *SaveSessionRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type SaveSessionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveSessionResponse) Reset() {
	*x = SaveSessionResponse{}
	mi := &file_proto_anamnesis_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveSessionResponse) ProtoMessage() {}

func (x *SaveSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_anamnesis_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveSessionResponse.ProtoReflect.Descriptor instead.
func (*SaveSessionResponse) Descriptor() ([]byte, []int) {
	return file_proto_anamnesis_proto_rawDescGZIP(), []int{1}
}

func (x *SaveSessionResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_anamnesis_proto protoreflect.FileDescriptor

const file_proto_anamnesis_proto_rawDesc = "" +
	"\n" +
	"\x15proto/anamnesis.proto\x12\asession\"a\n" +
	"\x12SaveSessionRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1c\n" +
	"\ttimestamp\x18\x02 \x01(\tR\ttimestamp\x12\x14\n" +
	"\x05notes\x18\x03 \x01(\tR\x05notes\"-\n" +
	"\x13SaveSessionResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2Z\n" +
	"\x0eSessionService\x12H\n" +
	"\vSaveSession\x12\x1b.session.SaveSessionRequest\x1a\x1c.session.SaveSessionResponseB0Z.github.com/ATursunbekov/MedApp/proto/anamnesisb\x06proto3"

var (
	file_proto_anamnesis_proto_rawDescOnce sync.Once
	file_proto_anamnesis_proto_rawDescData []byte
)

func file_proto_anamnesis_proto_rawDescGZIP() []byte {
	file_proto_anamnesis_proto_rawDescOnce.Do(func() {
		file_proto_anamnesis_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_anamnesis_proto_rawDesc), len(file_proto_anamnesis_proto_rawDesc)))
	})
	return file_proto_anamnesis_proto_rawDescData
}

var file_proto_anamnesis_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_anamnesis_proto_goTypes = []any{
	(*SaveSessionRequest)(nil),  // 0: session.SaveSessionRequest
	(*SaveSessionResponse)(nil), // 1: session.SaveSessionResponse
}
var file_proto_anamnesis_proto_depIdxs = []int32{
	0, // 0: session.SessionService.SaveSession:input_type -> session.SaveSessionRequest
	1, // 1: session.SessionService.SaveSession:output_type -> session.SaveSessionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_anamnesis_proto_init() }
func file_proto_anamnesis_proto_init() {
	if File_proto_anamnesis_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_anamnesis_proto_rawDesc), len(file_proto_anamnesis_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_anamnesis_proto_goTypes,
		DependencyIndexes: file_proto_anamnesis_proto_depIdxs,
		MessageInfos:      file_proto_anamnesis_proto_msgTypes,
	}.Build()
	File_proto_anamnesis_proto = out.File
	file_proto_anamnesis_proto_goTypes = nil
	file_proto_anamnesis_proto_depIdxs = nil
}
