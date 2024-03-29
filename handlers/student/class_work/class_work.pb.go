// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: student/class_work.proto

package class_work

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

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message"`
	Status  uint32 `protobuf:"varint,2,opt,name=status,proto3" json:"status"`
	Code    string `protobuf:"bytes,3,opt,name=code,proto3" json:"code"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_class_work_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_student_class_work_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_student_class_work_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Meta) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Meta) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page         uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Limit        uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
	Prev         uint32 `protobuf:"varint,3,opt,name=prev,proto3" json:"prev"`
	Next         uint32 `protobuf:"varint,4,opt,name=next,proto3" json:"next"`
	TotalPages   uint32 `protobuf:"varint,5,opt,name=total_pages,json=totalPages,proto3" json:"total_pages"`
	TotalRecords uint32 `protobuf:"varint,6,opt,name=total_records,json=totalRecords,proto3" json:"total_records"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_class_work_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_student_class_work_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_student_class_work_proto_rawDescGZIP(), []int{1}
}

func (x *Pagination) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetPrev() uint32 {
	if x != nil {
		return x.Prev
	}
	return 0
}

func (x *Pagination) GetNext() uint32 {
	if x != nil {
		return x.Next
	}
	return 0
}

func (x *Pagination) GetTotalPages() uint32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *Pagination) GetTotalRecords() uint32 {
	if x != nil {
		return x.TotalRecords
	}
	return 0
}

type GetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_class_work_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_class_work_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_student_class_work_proto_rawDescGZIP(), []int{2}
}

func (x *GetListRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetListRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Title                  string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`
	Abstraction            string  `protobuf:"bytes,3,opt,name=abstraction,proto3" json:"abstraction"`
	FileUrl                string  `protobuf:"bytes,4,opt,name=file_url,json=fileUrl,proto3" json:"file_url"`
	FilePath               string  `protobuf:"bytes,5,opt,name=file_path,json=filePath,proto3" json:"file_path"`
	FilePathType           string  `protobuf:"bytes,6,opt,name=file_path_type,json=filePathType,proto3" json:"file_path_type"`
	LecturerId             string  `protobuf:"bytes,7,opt,name=lecturer_id,json=lecturerId,proto3" json:"lecturer_id"`
	LecturerName           string  `protobuf:"bytes,8,opt,name=lecturer_name,json=lecturerName,proto3" json:"lecturer_name"`
	LecturerFrontTitle     string  `protobuf:"bytes,9,opt,name=lecturer_front_title,json=lecturerFrontTitle,proto3" json:"lecturer_front_title"`
	LecturerBackDegree     string  `protobuf:"bytes,10,opt,name=lecturer_back_degree,json=lecturerBackDegree,proto3" json:"lecturer_back_degree"`
	StartTime              string  `protobuf:"bytes,11,opt,name=start_time,json=startTime,proto3" json:"start_time"`
	EndTime                string  `protobuf:"bytes,12,opt,name=end_time,json=endTime,proto3" json:"end_time"`
	TotalSubmission        uint32  `protobuf:"varint,13,opt,name=total_submission,json=totalSubmission,proto3" json:"total_submission"`
	SubmissionFileUrl      string  `protobuf:"bytes,14,opt,name=submission_file_url,json=submissionFileUrl,proto3" json:"submission_file_url"`
	SubmissionFilePath     string  `protobuf:"bytes,15,opt,name=submission_file_path,json=submissionFilePath,proto3" json:"submission_file_path"`
	SubmissionFilePathType string  `protobuf:"bytes,16,opt,name=submission_file_path_type,json=submissionFilePathType,proto3" json:"submission_file_path_type"`
	SubmissionPoint        float64 `protobuf:"fixed64,17,opt,name=submission_point,json=submissionPoint,proto3" json:"submission_point"`
}

func (x *GetListResponseData) Reset() {
	*x = GetListResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_class_work_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponseData) ProtoMessage() {}

func (x *GetListResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_student_class_work_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponseData.ProtoReflect.Descriptor instead.
func (*GetListResponseData) Descriptor() ([]byte, []int) {
	return file_student_class_work_proto_rawDescGZIP(), []int{3}
}

func (x *GetListResponseData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetListResponseData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetListResponseData) GetAbstraction() string {
	if x != nil {
		return x.Abstraction
	}
	return ""
}

func (x *GetListResponseData) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *GetListResponseData) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *GetListResponseData) GetFilePathType() string {
	if x != nil {
		return x.FilePathType
	}
	return ""
}

func (x *GetListResponseData) GetLecturerId() string {
	if x != nil {
		return x.LecturerId
	}
	return ""
}

func (x *GetListResponseData) GetLecturerName() string {
	if x != nil {
		return x.LecturerName
	}
	return ""
}

func (x *GetListResponseData) GetLecturerFrontTitle() string {
	if x != nil {
		return x.LecturerFrontTitle
	}
	return ""
}

func (x *GetListResponseData) GetLecturerBackDegree() string {
	if x != nil {
		return x.LecturerBackDegree
	}
	return ""
}

func (x *GetListResponseData) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *GetListResponseData) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *GetListResponseData) GetTotalSubmission() uint32 {
	if x != nil {
		return x.TotalSubmission
	}
	return 0
}

func (x *GetListResponseData) GetSubmissionFileUrl() string {
	if x != nil {
		return x.SubmissionFileUrl
	}
	return ""
}

func (x *GetListResponseData) GetSubmissionFilePath() string {
	if x != nil {
		return x.SubmissionFilePath
	}
	return ""
}

func (x *GetListResponseData) GetSubmissionFilePathType() string {
	if x != nil {
		return x.SubmissionFilePathType
	}
	return ""
}

func (x *GetListResponseData) GetSubmissionPoint() float64 {
	if x != nil {
		return x.SubmissionPoint
	}
	return 0
}

type GetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta       *Meta                  `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta"`
	Pagination *Pagination            `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination"`
	Data       []*GetListResponseData `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_class_work_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_student_class_work_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse.ProtoReflect.Descriptor instead.
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return file_student_class_work_proto_rawDescGZIP(), []int{4}
}

func (x *GetListResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *GetListResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *GetListResponse) GetData() []*GetListResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type SubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassWorkId  string `protobuf:"bytes,1,opt,name=class_work_id,json=classWorkId,proto3" json:"class_work_id"`
	FilePath     string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path"`
	FilePathType string `protobuf:"bytes,3,opt,name=file_path_type,json=filePathType,proto3" json:"file_path_type"`
}

func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_class_work_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRequest) ProtoMessage() {}

func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_class_work_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRequest.ProtoReflect.Descriptor instead.
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_student_class_work_proto_rawDescGZIP(), []int{5}
}

func (x *SubmitRequest) GetClassWorkId() string {
	if x != nil {
		return x.ClassWorkId
	}
	return ""
}

func (x *SubmitRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *SubmitRequest) GetFilePathType() string {
	if x != nil {
		return x.FilePathType
	}
	return ""
}

type SubmitResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubmitResponseData) Reset() {
	*x = SubmitResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_class_work_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitResponseData) ProtoMessage() {}

func (x *SubmitResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_student_class_work_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitResponseData.ProtoReflect.Descriptor instead.
func (*SubmitResponseData) Descriptor() ([]byte, []int) {
	return file_student_class_work_proto_rawDescGZIP(), []int{6}
}

type SubmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Meta               `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta"`
	Data *SubmitResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *SubmitResponse) Reset() {
	*x = SubmitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_class_work_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitResponse) ProtoMessage() {}

func (x *SubmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_student_class_work_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitResponse.ProtoReflect.Descriptor instead.
func (*SubmitResponse) Descriptor() ([]byte, []int) {
	return file_student_class_work_proto_rawDescGZIP(), []int{7}
}

func (x *SubmitResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *SubmitResponse) GetData() *SubmitResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_student_class_work_proto protoreflect.FileDescriptor

var file_student_class_work_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x4c,
	0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xa4, 0x01, 0x0a,
	0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x72, 0x65, 0x76, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x72, 0x65, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x92, 0x05, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x62, 0x73, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x24, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x30, 0x0a,
	0x14, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x64,
	0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x42, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x19, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xbc, 0x01, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3e,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x76, 0x0a, 0x0d, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0d,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a,
	0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x7a, 0x0a, 0x0e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xc2, 0x01, 0x0a, 0x17, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x12, 0x54, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x12, 0x21, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_student_class_work_proto_rawDescOnce sync.Once
	file_student_class_work_proto_rawDescData = file_student_class_work_proto_rawDesc
)

func file_student_class_work_proto_rawDescGZIP() []byte {
	file_student_class_work_proto_rawDescOnce.Do(func() {
		file_student_class_work_proto_rawDescData = protoimpl.X.CompressGZIP(file_student_class_work_proto_rawDescData)
	})
	return file_student_class_work_proto_rawDescData
}

var file_student_class_work_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_student_class_work_proto_goTypes = []interface{}{
	(*Meta)(nil),                // 0: student_class_work.Meta
	(*Pagination)(nil),          // 1: student_class_work.Pagination
	(*GetListRequest)(nil),      // 2: student_class_work.GetListRequest
	(*GetListResponseData)(nil), // 3: student_class_work.GetListResponseData
	(*GetListResponse)(nil),     // 4: student_class_work.GetListResponse
	(*SubmitRequest)(nil),       // 5: student_class_work.SubmitRequest
	(*SubmitResponseData)(nil),  // 6: student_class_work.SubmitResponseData
	(*SubmitResponse)(nil),      // 7: student_class_work.SubmitResponse
}
var file_student_class_work_proto_depIdxs = []int32{
	0, // 0: student_class_work.GetListResponse.meta:type_name -> student_class_work.Meta
	1, // 1: student_class_work.GetListResponse.pagination:type_name -> student_class_work.Pagination
	3, // 2: student_class_work.GetListResponse.data:type_name -> student_class_work.GetListResponseData
	0, // 3: student_class_work.SubmitResponse.meta:type_name -> student_class_work.Meta
	6, // 4: student_class_work.SubmitResponse.data:type_name -> student_class_work.SubmitResponseData
	2, // 5: student_class_work.StudentClassWorkHandler.GetList:input_type -> student_class_work.GetListRequest
	5, // 6: student_class_work.StudentClassWorkHandler.Submit:input_type -> student_class_work.SubmitRequest
	4, // 7: student_class_work.StudentClassWorkHandler.GetList:output_type -> student_class_work.GetListResponse
	7, // 8: student_class_work.StudentClassWorkHandler.Submit:output_type -> student_class_work.SubmitResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_student_class_work_proto_init() }
func file_student_class_work_proto_init() {
	if File_student_class_work_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_student_class_work_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_student_class_work_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_student_class_work_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRequest); i {
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
		file_student_class_work_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResponseData); i {
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
		file_student_class_work_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResponse); i {
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
		file_student_class_work_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRequest); i {
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
		file_student_class_work_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitResponseData); i {
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
		file_student_class_work_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitResponse); i {
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
			RawDescriptor: file_student_class_work_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_student_class_work_proto_goTypes,
		DependencyIndexes: file_student_class_work_proto_depIdxs,
		MessageInfos:      file_student_class_work_proto_msgTypes,
	}.Build()
	File_student_class_work_proto = out.File
	file_student_class_work_proto_rawDesc = nil
	file_student_class_work_proto_goTypes = nil
	file_student_class_work_proto_depIdxs = nil
}
