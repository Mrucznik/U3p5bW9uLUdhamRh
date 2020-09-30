// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: proto/urls/urls.proto

package urls

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Message that represents an URL data.
type Url struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// URL Address.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Interval determining how often data should be fetched.
	Interval int32 `protobuf:"varint,3,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *Url) Reset() {
	*x = Url{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urls_urls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url) ProtoMessage() {}

func (x *Url) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urls_urls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url.ProtoReflect.Descriptor instead.
func (*Url) Descriptor() ([]byte, []int) {
	return file_proto_urls_urls_proto_rawDescGZIP(), []int{0}
}

func (x *Url) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Url) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Url) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

// Message that represents a response from URL.
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Response received from an URL call.
	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	// How long does it take to get a response.
	Duration float64 `protobuf:"fixed64,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// Creation timestamp.
	CreatedAt int64 `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urls_urls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urls_urls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_urls_urls_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *Response) GetDuration() float64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Response) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

// Request message for rpc `CreateUrl`.
type CreateUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL Address.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Interval determining how often data should be fetched.
	Interval int32 `protobuf:"varint,2,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *CreateUrlRequest) Reset() {
	*x = CreateUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urls_urls_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUrlRequest) ProtoMessage() {}

func (x *CreateUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urls_urls_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUrlRequest.ProtoReflect.Descriptor instead.
func (*CreateUrlRequest) Descriptor() ([]byte, []int) {
	return file_proto_urls_urls_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUrlRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateUrlRequest) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

// Response message for rpc `CreateUrl`.
type CreateUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of created URL address.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateUrlResponse) Reset() {
	*x = CreateUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urls_urls_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUrlResponse) ProtoMessage() {}

func (x *CreateUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urls_urls_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUrlResponse.ProtoReflect.Descriptor instead.
func (*CreateUrlResponse) Descriptor() ([]byte, []int) {
	return file_proto_urls_urls_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUrlResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Request message for rpc `DeleteUrl`.
type DeleteUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of URL Address.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUrlRequest) Reset() {
	*x = DeleteUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urls_urls_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUrlRequest) ProtoMessage() {}

func (x *DeleteUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urls_urls_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUrlRequest.ProtoReflect.Descriptor instead.
func (*DeleteUrlRequest) Descriptor() ([]byte, []int) {
	return file_proto_urls_urls_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUrlRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Response message for rpc `DeleteUrl`.
type DeleteUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUrlResponse) Reset() {
	*x = DeleteUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urls_urls_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUrlResponse) ProtoMessage() {}

func (x *DeleteUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urls_urls_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUrlResponse.ProtoReflect.Descriptor instead.
func (*DeleteUrlResponse) Descriptor() ([]byte, []int) {
	return file_proto_urls_urls_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteUrlResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Request message for rpc `GetUrls`.
type GetUrlsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUrlsRequest) Reset() {
	*x = GetUrlsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urls_urls_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUrlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUrlsRequest) ProtoMessage() {}

func (x *GetUrlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urls_urls_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUrlsRequest.ProtoReflect.Descriptor instead.
func (*GetUrlsRequest) Descriptor() ([]byte, []int) {
	return file_proto_urls_urls_proto_rawDescGZIP(), []int{6}
}

// Response message for rpc `GetUrls`.
type GetUrlsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*Url `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *GetUrlsResponse) Reset() {
	*x = GetUrlsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urls_urls_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUrlsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUrlsResponse) ProtoMessage() {}

func (x *GetUrlsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urls_urls_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUrlsResponse.ProtoReflect.Descriptor instead.
func (*GetUrlsResponse) Descriptor() ([]byte, []int) {
	return file_proto_urls_urls_proto_rawDescGZIP(), []int{7}
}

func (x *GetUrlsResponse) GetUrls() []*Url {
	if x != nil {
		return x.Urls
	}
	return nil
}

// Request message for rpc `GetUrlHistory`.
type GetUrlHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of URL Address.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUrlHistoryRequest) Reset() {
	*x = GetUrlHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urls_urls_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUrlHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUrlHistoryRequest) ProtoMessage() {}

func (x *GetUrlHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urls_urls_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUrlHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetUrlHistoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_urls_urls_proto_rawDescGZIP(), []int{8}
}

func (x *GetUrlHistoryRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Response message for rpc `GetUrlHistory`.
type GetUrlHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	History []*Response `protobuf:"bytes,1,rep,name=history,proto3" json:"history,omitempty"`
}

func (x *GetUrlHistoryResponse) Reset() {
	*x = GetUrlHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urls_urls_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUrlHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUrlHistoryResponse) ProtoMessage() {}

func (x *GetUrlHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urls_urls_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUrlHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetUrlHistoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_urls_urls_proto_rawDescGZIP(), []int{9}
}

func (x *GetUrlHistoryResponse) GetHistory() []*Response {
	if x != nil {
		return x.History
	}
	return nil
}

var File_proto_urls_urls_proto protoreflect.FileDescriptor

var file_proto_urls_urls_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x2f, 0x75, 0x72, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x03, 0x55,
	0x72, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x22, 0x61, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x40, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x23, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x2e, 0x55, 0x72,
	0x6c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x72,
	0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x41, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x72, 0x6c, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x32, 0xf8, 0x02, 0x0a, 0x0b, 0x55, 0x72, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x55, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12,
	0x16, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x57, 0x0a, 0x09, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x75, 0x72, 0x6c, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a,
	0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x4c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x14, 0x2e,
	0x75, 0x72, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x72,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x6b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x1a, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x75, 0x72, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x0c, 0x5a,
	0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_urls_urls_proto_rawDescOnce sync.Once
	file_proto_urls_urls_proto_rawDescData = file_proto_urls_urls_proto_rawDesc
)

func file_proto_urls_urls_proto_rawDescGZIP() []byte {
	file_proto_urls_urls_proto_rawDescOnce.Do(func() {
		file_proto_urls_urls_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_urls_urls_proto_rawDescData)
	})
	return file_proto_urls_urls_proto_rawDescData
}

var file_proto_urls_urls_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_urls_urls_proto_goTypes = []interface{}{
	(*Url)(nil),                   // 0: urls.Url
	(*Response)(nil),              // 1: urls.Response
	(*CreateUrlRequest)(nil),      // 2: urls.CreateUrlRequest
	(*CreateUrlResponse)(nil),     // 3: urls.CreateUrlResponse
	(*DeleteUrlRequest)(nil),      // 4: urls.DeleteUrlRequest
	(*DeleteUrlResponse)(nil),     // 5: urls.DeleteUrlResponse
	(*GetUrlsRequest)(nil),        // 6: urls.GetUrlsRequest
	(*GetUrlsResponse)(nil),       // 7: urls.GetUrlsResponse
	(*GetUrlHistoryRequest)(nil),  // 8: urls.GetUrlHistoryRequest
	(*GetUrlHistoryResponse)(nil), // 9: urls.GetUrlHistoryResponse
}
var file_proto_urls_urls_proto_depIdxs = []int32{
	0, // 0: urls.GetUrlsResponse.urls:type_name -> urls.Url
	1, // 1: urls.GetUrlHistoryResponse.history:type_name -> urls.Response
	2, // 2: urls.UrlsService.CreateUrl:input_type -> urls.CreateUrlRequest
	4, // 3: urls.UrlsService.DeleteUrl:input_type -> urls.DeleteUrlRequest
	6, // 4: urls.UrlsService.GetUrls:input_type -> urls.GetUrlsRequest
	8, // 5: urls.UrlsService.GetUrlHistory:input_type -> urls.GetUrlHistoryRequest
	3, // 6: urls.UrlsService.CreateUrl:output_type -> urls.CreateUrlResponse
	5, // 7: urls.UrlsService.DeleteUrl:output_type -> urls.DeleteUrlResponse
	7, // 8: urls.UrlsService.GetUrls:output_type -> urls.GetUrlsResponse
	9, // 9: urls.UrlsService.GetUrlHistory:output_type -> urls.GetUrlHistoryResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_urls_urls_proto_init() }
func file_proto_urls_urls_proto_init() {
	if File_proto_urls_urls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_urls_urls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Url); i {
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
		file_proto_urls_urls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_urls_urls_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUrlRequest); i {
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
		file_proto_urls_urls_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUrlResponse); i {
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
		file_proto_urls_urls_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUrlRequest); i {
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
		file_proto_urls_urls_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUrlResponse); i {
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
		file_proto_urls_urls_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUrlsRequest); i {
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
		file_proto_urls_urls_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUrlsResponse); i {
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
		file_proto_urls_urls_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUrlHistoryRequest); i {
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
		file_proto_urls_urls_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUrlHistoryResponse); i {
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
			RawDescriptor: file_proto_urls_urls_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_urls_urls_proto_goTypes,
		DependencyIndexes: file_proto_urls_urls_proto_depIdxs,
		MessageInfos:      file_proto_urls_urls_proto_msgTypes,
	}.Build()
	File_proto_urls_urls_proto = out.File
	file_proto_urls_urls_proto_rawDesc = nil
	file_proto_urls_urls_proto_goTypes = nil
	file_proto_urls_urls_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UrlsServiceClient is the client API for UrlsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UrlsServiceClient interface {
	// Create an URL and runs worker, that fetch data from url by specified interval.
	CreateUrl(ctx context.Context, in *CreateUrlRequest, opts ...grpc.CallOption) (*CreateUrlResponse, error)
	// Delete an URL and stops worker.
	DeleteUrl(ctx context.Context, in *DeleteUrlRequest, opts ...grpc.CallOption) (*DeleteUrlResponse, error)
	// Gets all stored URLs.
	GetUrls(ctx context.Context, in *GetUrlsRequest, opts ...grpc.CallOption) (*GetUrlsResponse, error)
	// Get URL fetching history.
	GetUrlHistory(ctx context.Context, in *GetUrlHistoryRequest, opts ...grpc.CallOption) (*GetUrlHistoryResponse, error)
}

type urlsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlsServiceClient(cc grpc.ClientConnInterface) UrlsServiceClient {
	return &urlsServiceClient{cc}
}

func (c *urlsServiceClient) CreateUrl(ctx context.Context, in *CreateUrlRequest, opts ...grpc.CallOption) (*CreateUrlResponse, error) {
	out := new(CreateUrlResponse)
	err := c.cc.Invoke(ctx, "/urls.UrlsService/CreateUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlsServiceClient) DeleteUrl(ctx context.Context, in *DeleteUrlRequest, opts ...grpc.CallOption) (*DeleteUrlResponse, error) {
	out := new(DeleteUrlResponse)
	err := c.cc.Invoke(ctx, "/urls.UrlsService/DeleteUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlsServiceClient) GetUrls(ctx context.Context, in *GetUrlsRequest, opts ...grpc.CallOption) (*GetUrlsResponse, error) {
	out := new(GetUrlsResponse)
	err := c.cc.Invoke(ctx, "/urls.UrlsService/GetUrls", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlsServiceClient) GetUrlHistory(ctx context.Context, in *GetUrlHistoryRequest, opts ...grpc.CallOption) (*GetUrlHistoryResponse, error) {
	out := new(GetUrlHistoryResponse)
	err := c.cc.Invoke(ctx, "/urls.UrlsService/GetUrlHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlsServiceServer is the server API for UrlsService service.
type UrlsServiceServer interface {
	// Create an URL and runs worker, that fetch data from url by specified interval.
	CreateUrl(context.Context, *CreateUrlRequest) (*CreateUrlResponse, error)
	// Delete an URL and stops worker.
	DeleteUrl(context.Context, *DeleteUrlRequest) (*DeleteUrlResponse, error)
	// Gets all stored URLs.
	GetUrls(context.Context, *GetUrlsRequest) (*GetUrlsResponse, error)
	// Get URL fetching history.
	GetUrlHistory(context.Context, *GetUrlHistoryRequest) (*GetUrlHistoryResponse, error)
}

// UnimplementedUrlsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUrlsServiceServer struct {
}

func (*UnimplementedUrlsServiceServer) CreateUrl(context.Context, *CreateUrlRequest) (*CreateUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUrl not implemented")
}
func (*UnimplementedUrlsServiceServer) DeleteUrl(context.Context, *DeleteUrlRequest) (*DeleteUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUrl not implemented")
}
func (*UnimplementedUrlsServiceServer) GetUrls(context.Context, *GetUrlsRequest) (*GetUrlsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrls not implemented")
}
func (*UnimplementedUrlsServiceServer) GetUrlHistory(context.Context, *GetUrlHistoryRequest) (*GetUrlHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrlHistory not implemented")
}

func RegisterUrlsServiceServer(s *grpc.Server, srv UrlsServiceServer) {
	s.RegisterService(&_UrlsService_serviceDesc, srv)
}

func _UrlsService_CreateUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlsServiceServer).CreateUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urls.UrlsService/CreateUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlsServiceServer).CreateUrl(ctx, req.(*CreateUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlsService_DeleteUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlsServiceServer).DeleteUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urls.UrlsService/DeleteUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlsServiceServer).DeleteUrl(ctx, req.(*DeleteUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlsService_GetUrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlsServiceServer).GetUrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urls.UrlsService/GetUrls",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlsServiceServer).GetUrls(ctx, req.(*GetUrlsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlsService_GetUrlHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlsServiceServer).GetUrlHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/urls.UrlsService/GetUrlHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlsServiceServer).GetUrlHistory(ctx, req.(*GetUrlHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UrlsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "urls.UrlsService",
	HandlerType: (*UrlsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUrl",
			Handler:    _UrlsService_CreateUrl_Handler,
		},
		{
			MethodName: "DeleteUrl",
			Handler:    _UrlsService_DeleteUrl_Handler,
		},
		{
			MethodName: "GetUrls",
			Handler:    _UrlsService_GetUrls_Handler,
		},
		{
			MethodName: "GetUrlHistory",
			Handler:    _UrlsService_GetUrlHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/urls/urls.proto",
}
