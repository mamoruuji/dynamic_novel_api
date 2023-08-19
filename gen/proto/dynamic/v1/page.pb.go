// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/dynamic/v1/page.proto

package dynamicv1

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

type PageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId    int32  `protobuf:"varint,1,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Order     int32  `protobuf:"varint,3,opt,name=order,proto3" json:"order,omitempty"`
	ChapterId int32  `protobuf:"varint,4,opt,name=chapter_id,json=chapterId,proto3" json:"chapter_id,omitempty"`
}

func (x *PageData) Reset() {
	*x = PageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dynamic_v1_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageData) ProtoMessage() {}

func (x *PageData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dynamic_v1_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageData.ProtoReflect.Descriptor instead.
func (*PageData) Descriptor() ([]byte, []int) {
	return file_proto_dynamic_v1_page_proto_rawDescGZIP(), []int{0}
}

func (x *PageData) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *PageData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PageData) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *PageData) GetChapterId() int32 {
	if x != nil {
		return x.ChapterId
	}
	return 0
}

type ListPagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPagesRequest) Reset() {
	*x = ListPagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dynamic_v1_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPagesRequest) ProtoMessage() {}

func (x *ListPagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dynamic_v1_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPagesRequest.ProtoReflect.Descriptor instead.
func (*ListPagesRequest) Descriptor() ([]byte, []int) {
	return file_proto_dynamic_v1_page_proto_rawDescGZIP(), []int{1}
}

type ListPagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pages []*PageData `protobuf:"bytes,1,rep,name=pages,proto3" json:"pages,omitempty"`
}

func (x *ListPagesResponse) Reset() {
	*x = ListPagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dynamic_v1_page_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPagesResponse) ProtoMessage() {}

func (x *ListPagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dynamic_v1_page_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPagesResponse.ProtoReflect.Descriptor instead.
func (*ListPagesResponse) Descriptor() ([]byte, []int) {
	return file_proto_dynamic_v1_page_proto_rawDescGZIP(), []int{2}
}

func (x *ListPagesResponse) GetPages() []*PageData {
	if x != nil {
		return x.Pages
	}
	return nil
}

type AddPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	ChapterId int32  `protobuf:"varint,2,opt,name=chapter_id,json=chapterId,proto3" json:"chapter_id,omitempty"`
	Order     int32  `protobuf:"varint,3,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *AddPageRequest) Reset() {
	*x = AddPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dynamic_v1_page_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPageRequest) ProtoMessage() {}

func (x *AddPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dynamic_v1_page_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPageRequest.ProtoReflect.Descriptor instead.
func (*AddPageRequest) Descriptor() ([]byte, []int) {
	return file_proto_dynamic_v1_page_proto_rawDescGZIP(), []int{3}
}

func (x *AddPageRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddPageRequest) GetChapterId() int32 {
	if x != nil {
		return x.ChapterId
	}
	return 0
}

func (x *AddPageRequest) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type AddPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId int32 `protobuf:"varint,1,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
}

func (x *AddPageResponse) Reset() {
	*x = AddPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dynamic_v1_page_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPageResponse) ProtoMessage() {}

func (x *AddPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dynamic_v1_page_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPageResponse.ProtoReflect.Descriptor instead.
func (*AddPageResponse) Descriptor() ([]byte, []int) {
	return file_proto_dynamic_v1_page_proto_rawDescGZIP(), []int{4}
}

func (x *AddPageResponse) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

type DeletePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId int32 `protobuf:"varint,1,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
}

func (x *DeletePageRequest) Reset() {
	*x = DeletePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dynamic_v1_page_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePageRequest) ProtoMessage() {}

func (x *DeletePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dynamic_v1_page_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePageRequest.ProtoReflect.Descriptor instead.
func (*DeletePageRequest) Descriptor() ([]byte, []int) {
	return file_proto_dynamic_v1_page_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePageRequest) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

type DeletePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePageResponse) Reset() {
	*x = DeletePageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dynamic_v1_page_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePageResponse) ProtoMessage() {}

func (x *DeletePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dynamic_v1_page_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePageResponse.ProtoReflect.Descriptor instead.
func (*DeletePageResponse) Descriptor() ([]byte, []int) {
	return file_proto_dynamic_v1_page_proto_rawDescGZIP(), []int{6}
}

type UpdatePageStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageId int32  `protobuf:"varint,1,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Order  int32  `protobuf:"varint,3,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *UpdatePageStatusRequest) Reset() {
	*x = UpdatePageStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dynamic_v1_page_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePageStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePageStatusRequest) ProtoMessage() {}

func (x *UpdatePageStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dynamic_v1_page_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePageStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdatePageStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_dynamic_v1_page_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePageStatusRequest) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *UpdatePageStatusRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdatePageStatusRequest) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type UpdatePageStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePageStatusResponse) Reset() {
	*x = UpdatePageStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dynamic_v1_page_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePageStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePageStatusResponse) ProtoMessage() {}

func (x *UpdatePageStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dynamic_v1_page_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePageStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdatePageStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_dynamic_v1_page_proto_rawDescGZIP(), []int{8}
}

var File_proto_dynamic_v1_page_proto protoreflect.FileDescriptor

var file_proto_dynamic_v1_page_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x22,
	0x6e, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x22, 0x5b, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xff, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x50, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6d, 0x6f, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x5f, 0x6e, 0x6f, 0x76, 0x65, 0x6c, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x2f, 0x76,
	0x31, 0x3b, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_dynamic_v1_page_proto_rawDescOnce sync.Once
	file_proto_dynamic_v1_page_proto_rawDescData = file_proto_dynamic_v1_page_proto_rawDesc
)

func file_proto_dynamic_v1_page_proto_rawDescGZIP() []byte {
	file_proto_dynamic_v1_page_proto_rawDescOnce.Do(func() {
		file_proto_dynamic_v1_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dynamic_v1_page_proto_rawDescData)
	})
	return file_proto_dynamic_v1_page_proto_rawDescData
}

var file_proto_dynamic_v1_page_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_dynamic_v1_page_proto_goTypes = []interface{}{
	(*PageData)(nil),                 // 0: proto.dynamic.v1.PageData
	(*ListPagesRequest)(nil),         // 1: proto.dynamic.v1.ListPagesRequest
	(*ListPagesResponse)(nil),        // 2: proto.dynamic.v1.ListPagesResponse
	(*AddPageRequest)(nil),           // 3: proto.dynamic.v1.AddPageRequest
	(*AddPageResponse)(nil),          // 4: proto.dynamic.v1.AddPageResponse
	(*DeletePageRequest)(nil),        // 5: proto.dynamic.v1.DeletePageRequest
	(*DeletePageResponse)(nil),       // 6: proto.dynamic.v1.DeletePageResponse
	(*UpdatePageStatusRequest)(nil),  // 7: proto.dynamic.v1.UpdatePageStatusRequest
	(*UpdatePageStatusResponse)(nil), // 8: proto.dynamic.v1.UpdatePageStatusResponse
}
var file_proto_dynamic_v1_page_proto_depIdxs = []int32{
	0, // 0: proto.dynamic.v1.ListPagesResponse.pages:type_name -> proto.dynamic.v1.PageData
	1, // 1: proto.dynamic.v1.PageService.ListPages:input_type -> proto.dynamic.v1.ListPagesRequest
	3, // 2: proto.dynamic.v1.PageService.AddPage:input_type -> proto.dynamic.v1.AddPageRequest
	5, // 3: proto.dynamic.v1.PageService.DeletePage:input_type -> proto.dynamic.v1.DeletePageRequest
	7, // 4: proto.dynamic.v1.PageService.UpdatePageStatus:input_type -> proto.dynamic.v1.UpdatePageStatusRequest
	2, // 5: proto.dynamic.v1.PageService.ListPages:output_type -> proto.dynamic.v1.ListPagesResponse
	4, // 6: proto.dynamic.v1.PageService.AddPage:output_type -> proto.dynamic.v1.AddPageResponse
	6, // 7: proto.dynamic.v1.PageService.DeletePage:output_type -> proto.dynamic.v1.DeletePageResponse
	8, // 8: proto.dynamic.v1.PageService.UpdatePageStatus:output_type -> proto.dynamic.v1.UpdatePageStatusResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_dynamic_v1_page_proto_init() }
func file_proto_dynamic_v1_page_proto_init() {
	if File_proto_dynamic_v1_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dynamic_v1_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageData); i {
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
		file_proto_dynamic_v1_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPagesRequest); i {
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
		file_proto_dynamic_v1_page_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPagesResponse); i {
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
		file_proto_dynamic_v1_page_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPageRequest); i {
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
		file_proto_dynamic_v1_page_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPageResponse); i {
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
		file_proto_dynamic_v1_page_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePageRequest); i {
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
		file_proto_dynamic_v1_page_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePageResponse); i {
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
		file_proto_dynamic_v1_page_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePageStatusRequest); i {
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
		file_proto_dynamic_v1_page_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePageStatusResponse); i {
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
			RawDescriptor: file_proto_dynamic_v1_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_dynamic_v1_page_proto_goTypes,
		DependencyIndexes: file_proto_dynamic_v1_page_proto_depIdxs,
		MessageInfos:      file_proto_dynamic_v1_page_proto_msgTypes,
	}.Build()
	File_proto_dynamic_v1_page_proto = out.File
	file_proto_dynamic_v1_page_proto_rawDesc = nil
	file_proto_dynamic_v1_page_proto_goTypes = nil
	file_proto_dynamic_v1_page_proto_depIdxs = nil
}
