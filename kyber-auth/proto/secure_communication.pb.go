// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: secure_communication.proto

package kyber_auth

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

type KeyExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Kem       []byte `protobuf:"bytes,2,opt,name=kem,proto3" json:"kem,omitempty"`
}

func (x *KeyExchangeRequest) Reset() {
	*x = KeyExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secure_communication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyExchangeRequest) ProtoMessage() {}

func (x *KeyExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secure_communication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyExchangeRequest.ProtoReflect.Descriptor instead.
func (*KeyExchangeRequest) Descriptor() ([]byte, []int) {
	return file_secure_communication_proto_rawDescGZIP(), []int{0}
}

func (x *KeyExchangeRequest) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *KeyExchangeRequest) GetKem() []byte {
	if x != nil {
		return x.Kem
	}
	return nil
}

type KeyExchangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncryptedSharedSecret []byte `protobuf:"bytes,1,opt,name=encrypted_shared_secret,json=encryptedSharedSecret,proto3" json:"encrypted_shared_secret,omitempty"`
}

func (x *KeyExchangeResponse) Reset() {
	*x = KeyExchangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secure_communication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyExchangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyExchangeResponse) ProtoMessage() {}

func (x *KeyExchangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secure_communication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyExchangeResponse.ProtoReflect.Descriptor instead.
func (*KeyExchangeResponse) Descriptor() ([]byte, []int) {
	return file_secure_communication_proto_rawDescGZIP(), []int{1}
}

func (x *KeyExchangeResponse) GetEncryptedSharedSecret() []byte {
	if x != nil {
		return x.EncryptedSharedSecret
	}
	return nil
}

var File_secure_communication_proto protoreflect.FileDescriptor

var file_secure_communication_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x45, 0x0a, 0x12, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x6d, 0x22, 0x4d, 0x0a, 0x13, 0x4b, 0x65, 0x79, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x17, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x15, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x32, 0x76, 0x0a, 0x12, 0x4b, 0x65, 0x79, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a,
	0x0b, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x27, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x73, 0x6c, 0x69, 0x69, 0x74, 0x2e,
	0x6c, 0x6b, 0x2f, 0x72, 0x32, 0x34, 0x2d, 0x30, 0x35, 0x35, 0x2f, 0x72, 0x32, 0x34, 0x2d, 0x30,
	0x35, 0x35, 0x2f, 0x6b, 0x79, 0x62, 0x65, 0x72, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_secure_communication_proto_rawDescOnce sync.Once
	file_secure_communication_proto_rawDescData = file_secure_communication_proto_rawDesc
)

func file_secure_communication_proto_rawDescGZIP() []byte {
	file_secure_communication_proto_rawDescOnce.Do(func() {
		file_secure_communication_proto_rawDescData = protoimpl.X.CompressGZIP(file_secure_communication_proto_rawDescData)
	})
	return file_secure_communication_proto_rawDescData
}

var file_secure_communication_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_secure_communication_proto_goTypes = []interface{}{
	(*KeyExchangeRequest)(nil),  // 0: securecommunication.KeyExchangeRequest
	(*KeyExchangeResponse)(nil), // 1: securecommunication.KeyExchangeResponse
}
var file_secure_communication_proto_depIdxs = []int32{
	0, // 0: securecommunication.KeyExchangeService.KeyExchange:input_type -> securecommunication.KeyExchangeRequest
	1, // 1: securecommunication.KeyExchangeService.KeyExchange:output_type -> securecommunication.KeyExchangeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_secure_communication_proto_init() }
func file_secure_communication_proto_init() {
	if File_secure_communication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_secure_communication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyExchangeRequest); i {
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
		file_secure_communication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyExchangeResponse); i {
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
			RawDescriptor: file_secure_communication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_secure_communication_proto_goTypes,
		DependencyIndexes: file_secure_communication_proto_depIdxs,
		MessageInfos:      file_secure_communication_proto_msgTypes,
	}.Build()
	File_secure_communication_proto = out.File
	file_secure_communication_proto_rawDesc = nil
	file_secure_communication_proto_goTypes = nil
	file_secure_communication_proto_depIdxs = nil
}