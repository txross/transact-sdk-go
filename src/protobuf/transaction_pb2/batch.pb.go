// Copyright 2016 Intel Corporation
// Copyright (C) 2020 Theadora Ross
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// -----------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        (unknown)
// source: batch.proto

package transaction_pb2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BatchHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Public key for the client that signed the BatchHeader
	SignerPublicKey string `protobuf:"bytes,1,opt,name=signer_public_key,json=signerPublicKey,proto3" json:"signer_public_key,omitempty"`
	// List of transaction.header_signatures that match the order of
	// transactions required for the batch
	TransactionIds []string `protobuf:"bytes,2,rep,name=transaction_ids,json=transactionIds,proto3" json:"transaction_ids,omitempty"`
}

func (x *BatchHeader) Reset() {
	*x = BatchHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchHeader) ProtoMessage() {}

func (x *BatchHeader) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchHeader.ProtoReflect.Descriptor instead.
func (*BatchHeader) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{0}
}

func (x *BatchHeader) GetSignerPublicKey() string {
	if x != nil {
		return x.SignerPublicKey
	}
	return ""
}

func (x *BatchHeader) GetTransactionIds() []string {
	if x != nil {
		return x.TransactionIds
	}
	return nil
}

type Batch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The serialized version of the BatchHeader
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The signature derived from signing the header
	HeaderSignature string `protobuf:"bytes,2,opt,name=header_signature,json=headerSignature,proto3" json:"header_signature,omitempty"`
	// A list of the transactions that match the list of
	// transaction_ids listed in the batch header
	Transactions []*Transaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
	// A debugging flag which indicates this batch should be traced through the
	// system, resulting in a higher level of debugging output.
	Trace bool `protobuf:"varint,4,opt,name=trace,proto3" json:"trace,omitempty"`
}

func (x *Batch) Reset() {
	*x = Batch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Batch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Batch) ProtoMessage() {}

func (x *Batch) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Batch.ProtoReflect.Descriptor instead.
func (*Batch) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{1}
}

func (x *Batch) GetHeader() []byte {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Batch) GetHeaderSignature() string {
	if x != nil {
		return x.HeaderSignature
	}
	return ""
}

func (x *Batch) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *Batch) GetTrace() bool {
	if x != nil {
		return x.Trace
	}
	return false
}

type BatchList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Batches []*Batch `protobuf:"bytes,1,rep,name=batches,proto3" json:"batches,omitempty"`
}

func (x *BatchList) Reset() {
	*x = BatchList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_batch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchList) ProtoMessage() {}

func (x *BatchList) ProtoReflect() protoreflect.Message {
	mi := &file_batch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchList.ProtoReflect.Descriptor instead.
func (*BatchList) Descriptor() ([]byte, []int) {
	return file_batch_proto_rawDescGZIP(), []int{2}
}

func (x *BatchList) GetBatches() []*Batch {
	if x != nil {
		return x.Batches
	}
	return nil
}

var File_batch_proto protoreflect.FileDescriptor

var file_batch_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x62, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x11, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x05, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x30, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x22, 0x2d, 0x0a, 0x09, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_batch_proto_rawDescOnce sync.Once
	file_batch_proto_rawDescData = file_batch_proto_rawDesc
)

func file_batch_proto_rawDescGZIP() []byte {
	file_batch_proto_rawDescOnce.Do(func() {
		file_batch_proto_rawDescData = protoimpl.X.CompressGZIP(file_batch_proto_rawDescData)
	})
	return file_batch_proto_rawDescData
}

var file_batch_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_batch_proto_goTypes = []interface{}{
	(*BatchHeader)(nil), // 0: BatchHeader
	(*Batch)(nil),       // 1: Batch
	(*BatchList)(nil),   // 2: BatchList
	(*Transaction)(nil), // 3: Transaction
}
var file_batch_proto_depIdxs = []int32{
	3, // 0: Batch.transactions:type_name -> Transaction
	1, // 1: BatchList.batches:type_name -> Batch
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_batch_proto_init() }
func file_batch_proto_init() {
	if File_batch_proto != nil {
		return
	}
	file_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_batch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchHeader); i {
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
		file_batch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Batch); i {
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
		file_batch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchList); i {
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
			RawDescriptor: file_batch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_batch_proto_goTypes,
		DependencyIndexes: file_batch_proto_depIdxs,
		MessageInfos:      file_batch_proto_msgTypes,
	}.Build()
	File_batch_proto = out.File
	file_batch_proto_rawDesc = nil
	file_batch_proto_goTypes = nil
	file_batch_proto_depIdxs = nil
}
