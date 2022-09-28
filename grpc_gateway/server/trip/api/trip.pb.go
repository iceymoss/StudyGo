// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: trip.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTripRequest) Reset() {
	*x = GetTripRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripRequest) ProtoMessage() {}

func (x *GetTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripRequest.ProtoReflect.Descriptor instead.
func (*GetTripRequest) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{0}
}

func (x *GetTripRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PathData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *PathData) Reset() {
	*x = PathData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathData) ProtoMessage() {}

func (x *PathData) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathData.ProtoReflect.Descriptor instead.
func (*PathData) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{1}
}

func (x *PathData) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type GetTripResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripId  int32       `protobuf:"varint,1,opt,name=tripId,proto3" json:"tripId,omitempty"`
	UserId  int32       `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Start   string      `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	Path    []*PathData `protobuf:"bytes,4,rep,name=path,proto3" json:"path,omitempty"`
	End     string      `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	Distans int32       `protobuf:"varint,6,opt,name=distans,proto3" json:"distans,omitempty"`
	Fee     float32     `protobuf:"fixed32,7,opt,name=fee,proto3" json:"fee,omitempty"`
	Time    int32       `protobuf:"varint,8,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GetTripResponse) Reset() {
	*x = GetTripResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripResponse) ProtoMessage() {}

func (x *GetTripResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripResponse.ProtoReflect.Descriptor instead.
func (*GetTripResponse) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{2}
}

func (x *GetTripResponse) GetTripId() int32 {
	if x != nil {
		return x.TripId
	}
	return 0
}

func (x *GetTripResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetTripResponse) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *GetTripResponse) GetPath() []*PathData {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *GetTripResponse) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *GetTripResponse) GetDistans() int32 {
	if x != nil {
		return x.Distans
	}
	return 0
}

func (x *GetTripResponse) GetFee() float32 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *GetTripResponse) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

type CreateTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripId int32  `protobuf:"varint,1,opt,name=tripId,proto3" json:"tripId,omitempty"`
	UserId int32  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Start  string `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End    string `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *CreateTripRequest) Reset() {
	*x = CreateTripRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTripRequest) ProtoMessage() {}

func (x *CreateTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTripRequest.ProtoReflect.Descriptor instead.
func (*CreateTripRequest) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTripRequest) GetTripId() int32 {
	if x != nil {
		return x.TripId
	}
	return 0
}

func (x *CreateTripRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateTripRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *CreateTripRequest) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type CreateTripRespnse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TripId  int32   `protobuf:"varint,1,opt,name=tripId,proto3" json:"tripId,omitempty"`
	UserId  int32   `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Start   string  `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End     string  `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	Distans int32   `protobuf:"varint,5,opt,name=distans,proto3" json:"distans,omitempty"`
	Fee     float32 `protobuf:"fixed32,6,opt,name=fee,proto3" json:"fee,omitempty"`
	Time    int32   `protobuf:"varint,7,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *CreateTripRespnse) Reset() {
	*x = CreateTripRespnse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTripRespnse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTripRespnse) ProtoMessage() {}

func (x *CreateTripRespnse) ProtoReflect() protoreflect.Message {
	mi := &file_trip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTripRespnse.ProtoReflect.Descriptor instead.
func (*CreateTripRespnse) Descriptor() ([]byte, []int) {
	return file_trip_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTripRespnse) GetTripId() int32 {
	if x != nil {
		return x.TripId
	}
	return 0
}

func (x *CreateTripRespnse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateTripRespnse) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *CreateTripRespnse) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *CreateTripRespnse) GetDistans() int32 {
	if x != nil {
		return x.Distans
	}
	return 0
}

func (x *CreateTripRespnse) GetFee() float32 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *CreateTripRespnse) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_trip_proto protoreflect.FileDescriptor

var file_trip_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x72,
	0x69, 0x70, 0x2e, 0x76, 0x31, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x08, 0x50, 0x61, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xd0, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x72, 0x69, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x72, 0x69,
	0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x25, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x6b, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0xab, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74,
	0x72, 0x69, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x66, 0x65,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x32, 0x91, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x69, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70,
	0x12, 0x17, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72,
	0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x69, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69,
	0x70, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x74, 0x72, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x2e, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trip_proto_rawDescOnce sync.Once
	file_trip_proto_rawDescData = file_trip_proto_rawDesc
)

func file_trip_proto_rawDescGZIP() []byte {
	file_trip_proto_rawDescOnce.Do(func() {
		file_trip_proto_rawDescData = protoimpl.X.CompressGZIP(file_trip_proto_rawDescData)
	})
	return file_trip_proto_rawDescData
}

var file_trip_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_trip_proto_goTypes = []interface{}{
	(*GetTripRequest)(nil),    // 0: trip.v1.GetTripRequest
	(*PathData)(nil),          // 1: trip.v1.PathData
	(*GetTripResponse)(nil),   // 2: trip.v1.GetTripResponse
	(*CreateTripRequest)(nil), // 3: trip.v1.CreateTripRequest
	(*CreateTripRespnse)(nil), // 4: trip.v1.CreateTripRespnse
}
var file_trip_proto_depIdxs = []int32{
	1, // 0: trip.v1.GetTripResponse.path:type_name -> trip.v1.PathData
	0, // 1: trip.v1.TripService.GetTrip:input_type -> trip.v1.GetTripRequest
	3, // 2: trip.v1.TripService.CreateTrip:input_type -> trip.v1.CreateTripRequest
	2, // 3: trip.v1.TripService.GetTrip:output_type -> trip.v1.GetTripResponse
	4, // 4: trip.v1.TripService.CreateTrip:output_type -> trip.v1.CreateTripRespnse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_trip_proto_init() }
func file_trip_proto_init() {
	if File_trip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripRequest); i {
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
		file_trip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathData); i {
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
		file_trip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripResponse); i {
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
		file_trip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTripRequest); i {
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
		file_trip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTripRespnse); i {
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
			RawDescriptor: file_trip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trip_proto_goTypes,
		DependencyIndexes: file_trip_proto_depIdxs,
		MessageInfos:      file_trip_proto_msgTypes,
	}.Build()
	File_trip_proto = out.File
	file_trip_proto_rawDesc = nil
	file_trip_proto_goTypes = nil
	file_trip_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TripServiceClient is the client API for TripService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TripServiceClient interface {
	GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*GetTripResponse, error)
	CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*CreateTripRespnse, error)
}

type tripServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTripServiceClient(cc grpc.ClientConnInterface) TripServiceClient {
	return &tripServiceClient{cc}
}

func (c *tripServiceClient) GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*GetTripResponse, error) {
	out := new(GetTripResponse)
	err := c.cc.Invoke(ctx, "/trip.v1.TripService/GetTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tripServiceClient) CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*CreateTripRespnse, error) {
	out := new(CreateTripRespnse)
	err := c.cc.Invoke(ctx, "/trip.v1.TripService/CreateTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TripServiceServer is the server API for TripService service.
type TripServiceServer interface {
	GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error)
	CreateTrip(context.Context, *CreateTripRequest) (*CreateTripRespnse, error)
}

// UnimplementedTripServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTripServiceServer struct {
}

func (*UnimplementedTripServiceServer) GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrip not implemented")
}
func (*UnimplementedTripServiceServer) CreateTrip(context.Context, *CreateTripRequest) (*CreateTripRespnse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrip not implemented")
}

func RegisterTripServiceServer(s *grpc.Server, srv TripServiceServer) {
	s.RegisterService(&_TripService_serviceDesc, srv)
}

func _TripService_GetTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).GetTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trip.v1.TripService/GetTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).GetTrip(ctx, req.(*GetTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TripService_CreateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TripServiceServer).CreateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trip.v1.TripService/CreateTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TripServiceServer).CreateTrip(ctx, req.(*CreateTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TripService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trip.v1.TripService",
	HandlerType: (*TripServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTrip",
			Handler:    _TripService_GetTrip_Handler,
		},
		{
			MethodName: "CreateTrip",
			Handler:    _TripService_CreateTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trip.proto",
}
