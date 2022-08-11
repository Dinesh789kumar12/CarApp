// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: routing.proto

package routingpb

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

type RoutingAvailabilityCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *Location `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *RoutingAvailabilityCarRequest) Reset() {
	*x = RoutingAvailabilityCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingAvailabilityCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingAvailabilityCarRequest) ProtoMessage() {}

func (x *RoutingAvailabilityCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingAvailabilityCarRequest.ProtoReflect.Descriptor instead.
func (*RoutingAvailabilityCarRequest) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{0}
}

func (x *RoutingAvailabilityCarRequest) GetSource() *Location {
	if x != nil {
		return x.Source
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  int32 `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetLatitude() int32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() int32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type RoutingAvailabilityCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car      *Car   `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=Location,proto3" json:"Location,omitempty"`
	Distance int32  `protobuf:"varint,3,opt,name=Distance,proto3" json:"Distance,omitempty"`
}

func (x *RoutingAvailabilityCarResponse) Reset() {
	*x = RoutingAvailabilityCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingAvailabilityCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingAvailabilityCarResponse) ProtoMessage() {}

func (x *RoutingAvailabilityCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingAvailabilityCarResponse.ProtoReflect.Descriptor instead.
func (*RoutingAvailabilityCarResponse) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{2}
}

func (x *RoutingAvailabilityCarResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

func (x *RoutingAvailabilityCarResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *RoutingAvailabilityCarResponse) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type RoutingAvailabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car      *Car   `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=Location,proto3" json:"Location,omitempty"`
}

func (x *RoutingAvailabilityRequest) Reset() {
	*x = RoutingAvailabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingAvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingAvailabilityRequest) ProtoMessage() {}

func (x *RoutingAvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingAvailabilityRequest.ProtoReflect.Descriptor instead.
func (*RoutingAvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{3}
}

func (x *RoutingAvailabilityRequest) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

func (x *RoutingAvailabilityRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId   int32  `protobuf:"varint,1,opt,name=CarId,proto3" json:"CarId,omitempty"`
	CarType string `protobuf:"bytes,2,opt,name=CarType,proto3" json:"CarType,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{4}
}

func (x *Car) GetCarId() int32 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *Car) GetCarType() string {
	if x != nil {
		return x.CarType
	}
	return ""
}

type RoutingAvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car      *Car    `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
	Location string  `protobuf:"bytes,2,opt,name=Location,proto3" json:"Location,omitempty"`
	Price    float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *RoutingAvailabilityResponse) Reset() {
	*x = RoutingAvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingAvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingAvailabilityResponse) ProtoMessage() {}

func (x *RoutingAvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingAvailabilityResponse.ProtoReflect.Descriptor instead.
func (*RoutingAvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{5}
}

func (x *RoutingAvailabilityResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

func (x *RoutingAvailabilityResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *RoutingAvailabilityResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Null struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Null) Reset() {
	*x = Null{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Null) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Null) ProtoMessage() {}

func (x *Null) ProtoReflect() protoreflect.Message {
	mi := &file_routing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Null.ProtoReflect.Descriptor instead.
func (*Null) Descriptor() ([]byte, []int) {
	return file_routing_proto_rawDescGZIP(), []int{6}
}

var File_routing_proto protoreflect.FileDescriptor

var file_routing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x22, 0x4c, 0x0a, 0x1d, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x7a,
	0x0a, 0x1e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x20, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63,
	0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x5a, 0x0a, 0x1a, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70,
	0x62, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x03, 0x43, 0x61, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x43, 0x61, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x61,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x71, 0x0a,
	0x1b, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x03,
	0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x22, 0x06, 0x0a, 0x04, 0x4e, 0x75, 0x6c, 0x6c, 0x32, 0xbc, 0x02, 0x0a, 0x0e, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x64, 0x6f, 0x6e, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x25, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x6d,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x28, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x48, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x70, 0x62, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x69, 0x6e, 0x65, 0x73, 0x68, 0x37, 0x38, 0x39, 0x6b,
	0x75, 0x6d, 0x61, 0x72, 0x31, 0x32, 0x2f, 0x43, 0x61, 0x72, 0x41, 0x70, 0x70, 0x2f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_routing_proto_rawDescOnce sync.Once
	file_routing_proto_rawDescData = file_routing_proto_rawDesc
)

func file_routing_proto_rawDescGZIP() []byte {
	file_routing_proto_rawDescOnce.Do(func() {
		file_routing_proto_rawDescData = protoimpl.X.CompressGZIP(file_routing_proto_rawDescData)
	})
	return file_routing_proto_rawDescData
}

var file_routing_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_routing_proto_goTypes = []interface{}{
	(*RoutingAvailabilityCarRequest)(nil),  // 0: routingpb.RoutingAvailabilityCarRequest
	(*Location)(nil),                       // 1: routingpb.Location
	(*RoutingAvailabilityCarResponse)(nil), // 2: routingpb.RoutingAvailabilityCarResponse
	(*RoutingAvailabilityRequest)(nil),     // 3: routingpb.RoutingAvailabilityRequest
	(*Car)(nil),                            // 4: routingpb.Car
	(*RoutingAvailabilityResponse)(nil),    // 5: routingpb.RoutingAvailabilityResponse
	(*Null)(nil),                           // 6: routingpb.Null
}
var file_routing_proto_depIdxs = []int32{
	1, // 0: routingpb.RoutingAvailabilityCarRequest.source:type_name -> routingpb.Location
	4, // 1: routingpb.RoutingAvailabilityCarResponse.car:type_name -> routingpb.Car
	4, // 2: routingpb.RoutingAvailabilityRequest.car:type_name -> routingpb.Car
	4, // 3: routingpb.RoutingAvailabilityResponse.car:type_name -> routingpb.Car
	3, // 4: routingpb.RoutingService.GetRateBasedonAvailability:input_type -> routingpb.RoutingAvailabilityRequest
	0, // 5: routingpb.RoutingService.GetCarAvailability:input_type -> routingpb.RoutingAvailabilityCarRequest
	0, // 6: routingpb.RoutingService.GetLocation:input_type -> routingpb.RoutingAvailabilityCarRequest
	5, // 7: routingpb.RoutingService.GetRateBasedonAvailability:output_type -> routingpb.RoutingAvailabilityResponse
	2, // 8: routingpb.RoutingService.GetCarAvailability:output_type -> routingpb.RoutingAvailabilityCarResponse
	6, // 9: routingpb.RoutingService.GetLocation:output_type -> routingpb.Null
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_routing_proto_init() }
func file_routing_proto_init() {
	if File_routing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingAvailabilityCarRequest); i {
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
		file_routing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_routing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingAvailabilityCarResponse); i {
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
		file_routing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingAvailabilityRequest); i {
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
		file_routing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_routing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingAvailabilityResponse); i {
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
		file_routing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Null); i {
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
			RawDescriptor: file_routing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routing_proto_goTypes,
		DependencyIndexes: file_routing_proto_depIdxs,
		MessageInfos:      file_routing_proto_msgTypes,
	}.Build()
	File_routing_proto = out.File
	file_routing_proto_rawDesc = nil
	file_routing_proto_goTypes = nil
	file_routing_proto_depIdxs = nil
}
