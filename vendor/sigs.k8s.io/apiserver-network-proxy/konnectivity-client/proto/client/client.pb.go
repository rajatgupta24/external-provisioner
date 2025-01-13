// Copyright The Kubernetes Authors.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: konnectivity-client/proto/client/client.proto

package client

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

type PacketType int32

const (
	PacketType_DIAL_REQ  PacketType = 0
	PacketType_DIAL_RSP  PacketType = 1
	PacketType_CLOSE_REQ PacketType = 2
	PacketType_CLOSE_RSP PacketType = 3
	PacketType_DATA      PacketType = 4
	PacketType_DIAL_CLS  PacketType = 5
	PacketType_DRAIN     PacketType = 6
)

// Enum value maps for PacketType.
var (
	PacketType_name = map[int32]string{
		0: "DIAL_REQ",
		1: "DIAL_RSP",
		2: "CLOSE_REQ",
		3: "CLOSE_RSP",
		4: "DATA",
		5: "DIAL_CLS",
		6: "DRAIN",
	}
	PacketType_value = map[string]int32{
		"DIAL_REQ":  0,
		"DIAL_RSP":  1,
		"CLOSE_REQ": 2,
		"CLOSE_RSP": 3,
		"DATA":      4,
		"DIAL_CLS":  5,
		"DRAIN":     6,
	}
)

func (x PacketType) Enum() *PacketType {
	p := new(PacketType)
	*p = x
	return p
}

func (x PacketType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PacketType) Descriptor() protoreflect.EnumDescriptor {
	return file_konnectivity_client_proto_client_client_proto_enumTypes[0].Descriptor()
}

func (PacketType) Type() protoreflect.EnumType {
	return &file_konnectivity_client_proto_client_client_proto_enumTypes[0]
}

func (x PacketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PacketType.Descriptor instead.
func (PacketType) EnumDescriptor() ([]byte, []int) {
	return file_konnectivity_client_proto_client_client_proto_rawDescGZIP(), []int{0}
}

type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type PacketType `protobuf:"varint,1,opt,name=type,proto3,enum=PacketType" json:"type,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*Packet_DialRequest
	//	*Packet_DialResponse
	//	*Packet_Data
	//	*Packet_CloseRequest
	//	*Packet_CloseResponse
	//	*Packet_CloseDial
	//	*Packet_Drain
	Payload isPacket_Payload `protobuf_oneof:"payload"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_konnectivity_client_proto_client_client_proto_rawDescGZIP(), []int{0}
}

func (x *Packet) GetType() PacketType {
	if x != nil {
		return x.Type
	}
	return PacketType_DIAL_REQ
}

func (m *Packet) GetPayload() isPacket_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Packet) GetDialRequest() *DialRequest {
	if x, ok := x.GetPayload().(*Packet_DialRequest); ok {
		return x.DialRequest
	}
	return nil
}

func (x *Packet) GetDialResponse() *DialResponse {
	if x, ok := x.GetPayload().(*Packet_DialResponse); ok {
		return x.DialResponse
	}
	return nil
}

func (x *Packet) GetData() *Data {
	if x, ok := x.GetPayload().(*Packet_Data); ok {
		return x.Data
	}
	return nil
}

func (x *Packet) GetCloseRequest() *CloseRequest {
	if x, ok := x.GetPayload().(*Packet_CloseRequest); ok {
		return x.CloseRequest
	}
	return nil
}

func (x *Packet) GetCloseResponse() *CloseResponse {
	if x, ok := x.GetPayload().(*Packet_CloseResponse); ok {
		return x.CloseResponse
	}
	return nil
}

func (x *Packet) GetCloseDial() *CloseDial {
	if x, ok := x.GetPayload().(*Packet_CloseDial); ok {
		return x.CloseDial
	}
	return nil
}

func (x *Packet) GetDrain() *Drain {
	if x, ok := x.GetPayload().(*Packet_Drain); ok {
		return x.Drain
	}
	return nil
}

type isPacket_Payload interface {
	isPacket_Payload()
}

type Packet_DialRequest struct {
	DialRequest *DialRequest `protobuf:"bytes,2,opt,name=dialRequest,proto3,oneof"`
}

type Packet_DialResponse struct {
	DialResponse *DialResponse `protobuf:"bytes,3,opt,name=dialResponse,proto3,oneof"`
}

type Packet_Data struct {
	Data *Data `protobuf:"bytes,4,opt,name=data,proto3,oneof"`
}

type Packet_CloseRequest struct {
	CloseRequest *CloseRequest `protobuf:"bytes,5,opt,name=closeRequest,proto3,oneof"`
}

type Packet_CloseResponse struct {
	CloseResponse *CloseResponse `protobuf:"bytes,6,opt,name=closeResponse,proto3,oneof"`
}

type Packet_CloseDial struct {
	CloseDial *CloseDial `protobuf:"bytes,7,opt,name=closeDial,proto3,oneof"`
}

type Packet_Drain struct {
	Drain *Drain `protobuf:"bytes,8,opt,name=drain,proto3,oneof"`
}

func (*Packet_DialRequest) isPacket_Payload() {}

func (*Packet_DialResponse) isPacket_Payload() {}

func (*Packet_Data) isPacket_Payload() {}

func (*Packet_CloseRequest) isPacket_Payload() {}

func (*Packet_CloseResponse) isPacket_Payload() {}

func (*Packet_CloseDial) isPacket_Payload() {}

func (*Packet_Drain) isPacket_Payload() {}

type DialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tcp or udp?
	Protocol string `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// node:port
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// random id for client, maybe should be longer
	Random int64 `protobuf:"varint,3,opt,name=random,proto3" json:"random,omitempty"`
}

func (x *DialRequest) Reset() {
	*x = DialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DialRequest) ProtoMessage() {}

func (x *DialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DialRequest.ProtoReflect.Descriptor instead.
func (*DialRequest) Descriptor() ([]byte, []int) {
	return file_konnectivity_client_proto_client_client_proto_rawDescGZIP(), []int{1}
}

func (x *DialRequest) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *DialRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DialRequest) GetRandom() int64 {
	if x != nil {
		return x.Random
	}
	return 0
}

type DialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// error failed reason; enum?
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// connectID indicates the identifier of the connection
	ConnectID int64 `protobuf:"varint,2,opt,name=connectID,proto3" json:"connectID,omitempty"`
	// random copied from DialRequest
	Random int64 `protobuf:"varint,3,opt,name=random,proto3" json:"random,omitempty"`
}

func (x *DialResponse) Reset() {
	*x = DialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DialResponse) ProtoMessage() {}

func (x *DialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DialResponse.ProtoReflect.Descriptor instead.
func (*DialResponse) Descriptor() ([]byte, []int) {
	return file_konnectivity_client_proto_client_client_proto_rawDescGZIP(), []int{2}
}

func (x *DialResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *DialResponse) GetConnectID() int64 {
	if x != nil {
		return x.ConnectID
	}
	return 0
}

func (x *DialResponse) GetRandom() int64 {
	if x != nil {
		return x.Random
	}
	return 0
}

type CloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// connectID of the stream to close
	ConnectID int64 `protobuf:"varint,1,opt,name=connectID,proto3" json:"connectID,omitempty"`
}

func (x *CloseRequest) Reset() {
	*x = CloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseRequest) ProtoMessage() {}

func (x *CloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseRequest.ProtoReflect.Descriptor instead.
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return file_konnectivity_client_proto_client_client_proto_rawDescGZIP(), []int{3}
}

func (x *CloseRequest) GetConnectID() int64 {
	if x != nil {
		return x.ConnectID
	}
	return 0
}

type CloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// error message
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// connectID indicates the identifier of the connection
	ConnectID int64 `protobuf:"varint,2,opt,name=connectID,proto3" json:"connectID,omitempty"`
}

func (x *CloseResponse) Reset() {
	*x = CloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseResponse) ProtoMessage() {}

func (x *CloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseResponse.ProtoReflect.Descriptor instead.
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return file_konnectivity_client_proto_client_client_proto_rawDescGZIP(), []int{4}
}

func (x *CloseResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CloseResponse) GetConnectID() int64 {
	if x != nil {
		return x.ConnectID
	}
	return 0
}

type CloseDial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// random id of the DialRequest
	Random int64 `protobuf:"varint,1,opt,name=random,proto3" json:"random,omitempty"`
}

func (x *CloseDial) Reset() {
	*x = CloseDial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseDial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseDial) ProtoMessage() {}

func (x *CloseDial) ProtoReflect() protoreflect.Message {
	mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseDial.ProtoReflect.Descriptor instead.
func (*CloseDial) Descriptor() ([]byte, []int) {
	return file_konnectivity_client_proto_client_client_proto_rawDescGZIP(), []int{5}
}

func (x *CloseDial) GetRandom() int64 {
	if x != nil {
		return x.Random
	}
	return 0
}

type Drain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Drain) Reset() {
	*x = Drain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Drain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Drain) ProtoMessage() {}

func (x *Drain) ProtoReflect() protoreflect.Message {
	mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Drain.ProtoReflect.Descriptor instead.
func (*Drain) Descriptor() ([]byte, []int) {
	return file_konnectivity_client_proto_client_client_proto_rawDescGZIP(), []int{6}
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// connectID to connect to
	ConnectID int64 `protobuf:"varint,1,opt,name=connectID,proto3" json:"connectID,omitempty"`
	// error message if error happens
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// stream data
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_konnectivity_client_proto_client_client_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_konnectivity_client_proto_client_client_proto_rawDescGZIP(), []int{7}
}

func (x *Data) GetConnectID() int64 {
	if x != nil {
		return x.ConnectID
	}
	return 0
}

func (x *Data) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Data) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_konnectivity_client_proto_client_client_proto protoreflect.FileDescriptor

var file_konnectivity_client_proto_client_client_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6b, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf1, 0x02, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x64,
	0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x0c, 0x64, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x64, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x33, 0x0a, 0x0c, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0d, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x09,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x09, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x05, 0x64, 0x72, 0x61, 0x69,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x44, 0x72, 0x61, 0x69, 0x6e, 0x48,
	0x00, 0x52, 0x05, 0x64, 0x72, 0x61, 0x69, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x5b, 0x0a, 0x0b, 0x44, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x22, 0x5a, 0x0a, 0x0c, 0x44, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x22, 0x2c, 0x0a, 0x0c,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x44, 0x22, 0x43, 0x0a, 0x0d, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x44, 0x22,
	0x23, 0x0a, 0x09, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x61,
	0x6e, 0x64, 0x6f, 0x6d, 0x22, 0x07, 0x0a, 0x05, 0x44, 0x72, 0x61, 0x69, 0x6e, 0x22, 0x4e, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x69, 0x0a,
	0x0a, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x49, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x41,
	0x4c, 0x5f, 0x52, 0x53, 0x50, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4c, 0x4f, 0x53, 0x45,
	0x5f, 0x52, 0x45, 0x51, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f,
	0x52, 0x53, 0x50, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x41, 0x10, 0x04, 0x12,
	0x0c, 0x0a, 0x08, 0x44, 0x49, 0x41, 0x4c, 0x5f, 0x43, 0x4c, 0x53, 0x10, 0x05, 0x12, 0x09, 0x0a,
	0x05, 0x44, 0x52, 0x41, 0x49, 0x4e, 0x10, 0x06, 0x32, 0x2f, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x12, 0x07, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x07, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x46, 0x5a, 0x44, 0x73, 0x69, 0x67,
	0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x6b, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_konnectivity_client_proto_client_client_proto_rawDescOnce sync.Once
	file_konnectivity_client_proto_client_client_proto_rawDescData = file_konnectivity_client_proto_client_client_proto_rawDesc
)

func file_konnectivity_client_proto_client_client_proto_rawDescGZIP() []byte {
	file_konnectivity_client_proto_client_client_proto_rawDescOnce.Do(func() {
		file_konnectivity_client_proto_client_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_konnectivity_client_proto_client_client_proto_rawDescData)
	})
	return file_konnectivity_client_proto_client_client_proto_rawDescData
}

var file_konnectivity_client_proto_client_client_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_konnectivity_client_proto_client_client_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_konnectivity_client_proto_client_client_proto_goTypes = []interface{}{
	(PacketType)(0),       // 0: PacketType
	(*Packet)(nil),        // 1: Packet
	(*DialRequest)(nil),   // 2: DialRequest
	(*DialResponse)(nil),  // 3: DialResponse
	(*CloseRequest)(nil),  // 4: CloseRequest
	(*CloseResponse)(nil), // 5: CloseResponse
	(*CloseDial)(nil),     // 6: CloseDial
	(*Drain)(nil),         // 7: Drain
	(*Data)(nil),          // 8: Data
}
var file_konnectivity_client_proto_client_client_proto_depIdxs = []int32{
	0, // 0: Packet.type:type_name -> PacketType
	2, // 1: Packet.dialRequest:type_name -> DialRequest
	3, // 2: Packet.dialResponse:type_name -> DialResponse
	8, // 3: Packet.data:type_name -> Data
	4, // 4: Packet.closeRequest:type_name -> CloseRequest
	5, // 5: Packet.closeResponse:type_name -> CloseResponse
	6, // 6: Packet.closeDial:type_name -> CloseDial
	7, // 7: Packet.drain:type_name -> Drain
	1, // 8: ProxyService.Proxy:input_type -> Packet
	1, // 9: ProxyService.Proxy:output_type -> Packet
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_konnectivity_client_proto_client_client_proto_init() }
func file_konnectivity_client_proto_client_client_proto_init() {
	if File_konnectivity_client_proto_client_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_konnectivity_client_proto_client_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet); i {
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
		file_konnectivity_client_proto_client_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DialRequest); i {
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
		file_konnectivity_client_proto_client_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DialResponse); i {
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
		file_konnectivity_client_proto_client_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseRequest); i {
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
		file_konnectivity_client_proto_client_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseResponse); i {
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
		file_konnectivity_client_proto_client_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseDial); i {
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
		file_konnectivity_client_proto_client_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Drain); i {
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
		file_konnectivity_client_proto_client_client_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
	file_konnectivity_client_proto_client_client_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Packet_DialRequest)(nil),
		(*Packet_DialResponse)(nil),
		(*Packet_Data)(nil),
		(*Packet_CloseRequest)(nil),
		(*Packet_CloseResponse)(nil),
		(*Packet_CloseDial)(nil),
		(*Packet_Drain)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_konnectivity_client_proto_client_client_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_konnectivity_client_proto_client_client_proto_goTypes,
		DependencyIndexes: file_konnectivity_client_proto_client_client_proto_depIdxs,
		EnumInfos:         file_konnectivity_client_proto_client_client_proto_enumTypes,
		MessageInfos:      file_konnectivity_client_proto_client_client_proto_msgTypes,
	}.Build()
	File_konnectivity_client_proto_client_client_proto = out.File
	file_konnectivity_client_proto_client_client_proto_rawDesc = nil
	file_konnectivity_client_proto_client_client_proto_goTypes = nil
	file_konnectivity_client_proto_client_client_proto_depIdxs = nil
}
