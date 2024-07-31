// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: resume.proto

package genprotos

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

type ResumeCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Education string `protobuf:"bytes,2,opt,name=education,proto3" json:"education,omitempty"`
	Summary   string `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *ResumeCreateReq) Reset() {
	*x = ResumeCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeCreateReq) ProtoMessage() {}

func (x *ResumeCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeCreateReq.ProtoReflect.Descriptor instead.
func (*ResumeCreateReq) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{0}
}

func (x *ResumeCreateReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ResumeCreateReq) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *ResumeCreateReq) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

type ResumeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Education string `protobuf:"bytes,3,opt,name=education,proto3" json:"education,omitempty"`
	Summary   string `protobuf:"bytes,4,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *ResumeRes) Reset() {
	*x = ResumeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeRes) ProtoMessage() {}

func (x *ResumeRes) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeRes.ProtoReflect.Descriptor instead.
func (*ResumeRes) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{1}
}

func (x *ResumeRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResumeRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ResumeRes) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *ResumeRes) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

type ResumeUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Education string `protobuf:"bytes,2,opt,name=education,proto3" json:"education,omitempty"`
	Summary   string `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
}

func (x *ResumeUpdateReq) Reset() {
	*x = ResumeUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeUpdateReq) ProtoMessage() {}

func (x *ResumeUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeUpdateReq.ProtoReflect.Descriptor instead.
func (*ResumeUpdateReq) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{2}
}

func (x *ResumeUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResumeUpdateReq) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *ResumeUpdateReq) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

type ResumeGetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Education string      `protobuf:"bytes,2,opt,name=education,proto3" json:"education,omitempty"`
	Filter    *Pagination `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ResumeGetAllReq) Reset() {
	*x = ResumeGetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeGetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeGetAllReq) ProtoMessage() {}

func (x *ResumeGetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeGetAllReq.ProtoReflect.Descriptor instead.
func (*ResumeGetAllReq) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{3}
}

func (x *ResumeGetAllReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ResumeGetAllReq) GetEducation() string {
	if x != nil {
		return x.Education
	}
	return ""
}

func (x *ResumeGetAllReq) GetFilter() *Pagination {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ResumeGetByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resume *ResumeRes `protobuf:"bytes,1,opt,name=resume,proto3" json:"resume,omitempty"`
}

func (x *ResumeGetByIdRes) Reset() {
	*x = ResumeGetByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeGetByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeGetByIdRes) ProtoMessage() {}

func (x *ResumeGetByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeGetByIdRes.ProtoReflect.Descriptor instead.
func (*ResumeGetByIdRes) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{4}
}

func (x *ResumeGetByIdRes) GetResume() *ResumeRes {
	if x != nil {
		return x.Resume
	}
	return nil
}

type ResumeGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resumes []*ResumeRes `protobuf:"bytes,1,rep,name=resumes,proto3" json:"resumes,omitempty"`
}

func (x *ResumeGetAllRes) Reset() {
	*x = ResumeGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resume_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeGetAllRes) ProtoMessage() {}

func (x *ResumeGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_resume_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeGetAllRes.ProtoReflect.Descriptor instead.
func (*ResumeGetAllRes) Descriptor() ([]byte, []int) {
	return file_resume_proto_rawDescGZIP(), []int{5}
}

func (x *ResumeGetAllRes) GetResumes() []*ResumeRes {
	if x != nil {
		return x.Resumes
	}
	return nil
}

var File_resume_proto protoreflect.FileDescriptor

var file_resume_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x73, 0x74, 0x61, 0x66, 0x66, 0x1a, 0x0a, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x62, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x6c, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x22, 0x59, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x64, 0x75, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x73,
	0x0a, 0x0f, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x64,
	0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x73,
	0x32, 0x86, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x0b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x42, 0x79, 0x69, 0x64, 0x1a, 0x17,
	0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x16, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x0b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x42, 0x79, 0x69, 0x64, 0x1a, 0x0b, 0x2e, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resume_proto_rawDescOnce sync.Once
	file_resume_proto_rawDescData = file_resume_proto_rawDesc
)

func file_resume_proto_rawDescGZIP() []byte {
	file_resume_proto_rawDescOnce.Do(func() {
		file_resume_proto_rawDescData = protoimpl.X.CompressGZIP(file_resume_proto_rawDescData)
	})
	return file_resume_proto_rawDescData
}

var file_resume_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_resume_proto_goTypes = []interface{}{
	(*ResumeCreateReq)(nil),  // 0: staff.ResumeCreateReq
	(*ResumeRes)(nil),        // 1: staff.ResumeRes
	(*ResumeUpdateReq)(nil),  // 2: staff.ResumeUpdateReq
	(*ResumeGetAllReq)(nil),  // 3: staff.ResumeGetAllReq
	(*ResumeGetByIdRes)(nil), // 4: staff.ResumeGetByIdRes
	(*ResumeGetAllRes)(nil),  // 5: staff.ResumeGetAllRes
	(*Pagination)(nil),       // 6: staff.Pagination
	(*Byid)(nil),             // 7: staff.Byid
	(*Void)(nil),             // 8: staff.Void
}
var file_resume_proto_depIdxs = []int32{
	6, // 0: staff.ResumeGetAllReq.filter:type_name -> staff.Pagination
	1, // 1: staff.ResumeGetByIdRes.resume:type_name -> staff.ResumeRes
	1, // 2: staff.ResumeGetAllRes.resumes:type_name -> staff.ResumeRes
	0, // 3: staff.ResumeService.Create:input_type -> staff.ResumeCreateReq
	7, // 4: staff.ResumeService.GetById:input_type -> staff.Byid
	3, // 5: staff.ResumeService.GetAll:input_type -> staff.ResumeGetAllReq
	2, // 6: staff.ResumeService.Update:input_type -> staff.ResumeUpdateReq
	7, // 7: staff.ResumeService.Delete:input_type -> staff.Byid
	1, // 8: staff.ResumeService.Create:output_type -> staff.ResumeRes
	4, // 9: staff.ResumeService.GetById:output_type -> staff.ResumeGetByIdRes
	5, // 10: staff.ResumeService.GetAll:output_type -> staff.ResumeGetAllRes
	1, // 11: staff.ResumeService.Update:output_type -> staff.ResumeRes
	8, // 12: staff.ResumeService.Delete:output_type -> staff.Void
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_resume_proto_init() }
func file_resume_proto_init() {
	if File_resume_proto != nil {
		return
	}
	file_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_resume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeCreateReq); i {
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
		file_resume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeRes); i {
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
		file_resume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeUpdateReq); i {
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
		file_resume_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeGetAllReq); i {
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
		file_resume_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeGetByIdRes); i {
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
		file_resume_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeGetAllRes); i {
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
			RawDescriptor: file_resume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resume_proto_goTypes,
		DependencyIndexes: file_resume_proto_depIdxs,
		MessageInfos:      file_resume_proto_msgTypes,
	}.Build()
	File_resume_proto = out.File
	file_resume_proto_rawDesc = nil
	file_resume_proto_goTypes = nil
	file_resume_proto_depIdxs = nil
}
