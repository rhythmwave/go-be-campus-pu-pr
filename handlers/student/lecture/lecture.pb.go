// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: student/lecture.proto

package lecture

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
		mi := &file_student_lecture_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_student_lecture_proto_msgTypes[0]
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
	return file_student_lecture_proto_rawDescGZIP(), []int{0}
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
		mi := &file_student_lecture_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_student_lecture_proto_msgTypes[1]
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
	return file_student_lecture_proto_rawDescGZIP(), []int{1}
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

type GetHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Limit     uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
	StartDate string `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate   string `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date"`
}

func (x *GetHistoryRequest) Reset() {
	*x = GetHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_lecture_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryRequest) ProtoMessage() {}

func (x *GetHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_lecture_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetHistoryRequest) Descriptor() ([]byte, []int) {
	return file_student_lecture_proto_rawDescGZIP(), []int{2}
}

func (x *GetHistoryRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetHistoryRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetHistoryRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetHistoryRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type GetHistoryResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	LectureDate string `protobuf:"bytes,2,opt,name=lecture_date,json=lectureDate,proto3" json:"lecture_date"`
	SubjectName string `protobuf:"bytes,3,opt,name=subject_name,json=subjectName,proto3" json:"subject_name"`
	AttendTime  uint32 `protobuf:"varint,4,opt,name=attend_time,json=attendTime,proto3" json:"attend_time"`
}

func (x *GetHistoryResponseData) Reset() {
	*x = GetHistoryResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_lecture_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoryResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryResponseData) ProtoMessage() {}

func (x *GetHistoryResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_student_lecture_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryResponseData.ProtoReflect.Descriptor instead.
func (*GetHistoryResponseData) Descriptor() ([]byte, []int) {
	return file_student_lecture_proto_rawDescGZIP(), []int{3}
}

func (x *GetHistoryResponseData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetHistoryResponseData) GetLectureDate() string {
	if x != nil {
		return x.LectureDate
	}
	return ""
}

func (x *GetHistoryResponseData) GetSubjectName() string {
	if x != nil {
		return x.SubjectName
	}
	return ""
}

func (x *GetHistoryResponseData) GetAttendTime() uint32 {
	if x != nil {
		return x.AttendTime
	}
	return 0
}

type GetHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta       *Meta                     `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta"`
	Pagination *Pagination               `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination"`
	Data       []*GetHistoryResponseData `protobuf:"bytes,3,rep,name=data,proto3" json:"data"`
}

func (x *GetHistoryResponse) Reset() {
	*x = GetHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_lecture_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoryResponse) ProtoMessage() {}

func (x *GetHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_student_lecture_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetHistoryResponse) Descriptor() ([]byte, []int) {
	return file_student_lecture_proto_rawDescGZIP(), []int{4}
}

func (x *GetHistoryResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *GetHistoryResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *GetHistoryResponse) GetData() []*GetHistoryResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type AttendAutonomousLectureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LectureId     string `protobuf:"bytes,1,opt,name=lecture_id,json=lectureId,proto3" json:"lecture_id"`
	PhotoPath     string `protobuf:"bytes,2,opt,name=photo_path,json=photoPath,proto3" json:"photo_path"`
	PhotoPathType string `protobuf:"bytes,3,opt,name=photo_path_type,json=photoPathType,proto3" json:"photo_path_type"`
}

func (x *AttendAutonomousLectureRequest) Reset() {
	*x = AttendAutonomousLectureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_lecture_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttendAutonomousLectureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttendAutonomousLectureRequest) ProtoMessage() {}

func (x *AttendAutonomousLectureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_student_lecture_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttendAutonomousLectureRequest.ProtoReflect.Descriptor instead.
func (*AttendAutonomousLectureRequest) Descriptor() ([]byte, []int) {
	return file_student_lecture_proto_rawDescGZIP(), []int{5}
}

func (x *AttendAutonomousLectureRequest) GetLectureId() string {
	if x != nil {
		return x.LectureId
	}
	return ""
}

func (x *AttendAutonomousLectureRequest) GetPhotoPath() string {
	if x != nil {
		return x.PhotoPath
	}
	return ""
}

func (x *AttendAutonomousLectureRequest) GetPhotoPathType() string {
	if x != nil {
		return x.PhotoPathType
	}
	return ""
}

type AttendAutonomousLectureResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AttendAutonomousLectureResponseData) Reset() {
	*x = AttendAutonomousLectureResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_lecture_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttendAutonomousLectureResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttendAutonomousLectureResponseData) ProtoMessage() {}

func (x *AttendAutonomousLectureResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_student_lecture_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttendAutonomousLectureResponseData.ProtoReflect.Descriptor instead.
func (*AttendAutonomousLectureResponseData) Descriptor() ([]byte, []int) {
	return file_student_lecture_proto_rawDescGZIP(), []int{6}
}

type AttendAutonomousLectureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Meta                                `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta"`
	Data *AttendAutonomousLectureResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *AttendAutonomousLectureResponse) Reset() {
	*x = AttendAutonomousLectureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_student_lecture_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttendAutonomousLectureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttendAutonomousLectureResponse) ProtoMessage() {}

func (x *AttendAutonomousLectureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_student_lecture_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttendAutonomousLectureResponse.ProtoReflect.Descriptor instead.
func (*AttendAutonomousLectureResponse) Descriptor() ([]byte, []int) {
	return file_student_lecture_proto_rawDescGZIP(), []int{7}
}

func (x *AttendAutonomousLectureResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *AttendAutonomousLectureResponse) GetData() *AttendAutonomousLectureResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_student_lecture_proto protoreflect.FileDescriptor

var file_student_lecture_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22, 0x4c, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x72, 0x65, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70,
	0x72, 0x65, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x77, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x86, 0x01, 0x0a, 0x1e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x41,
	0x75, 0x74, 0x6f, 0x6e, 0x6f, 0x6d, 0x6f, 0x75, 0x73, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x50, 0x61, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x22, 0x25, 0x0a,
	0x23, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x6f, 0x6e, 0x6f, 0x6d, 0x6f, 0x75,
	0x73, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x96, 0x01, 0x0a, 0x1f, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x41,
	0x75, 0x74, 0x6f, 0x6e, 0x6f, 0x6d, 0x6f, 0x75, 0x73, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x48, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x6f, 0x6e, 0x6f,
	0x6d, 0x6f, 0x75, 0x73, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xf0, 0x01,
	0x0a, 0x15, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x7e, 0x0a, 0x17, 0x41, 0x74, 0x74, 0x65, 0x6e,
	0x64, 0x41, 0x75, 0x74, 0x6f, 0x6e, 0x6f, 0x6d, 0x6f, 0x75, 0x73, 0x4c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x2f, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x6f, 0x6e,
	0x6f, 0x6d, 0x6f, 0x75, 0x73, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x6f,
	0x6e, 0x6f, 0x6d, 0x6f, 0x75, 0x73, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x22, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x1a, 0x5a, 0x18, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_student_lecture_proto_rawDescOnce sync.Once
	file_student_lecture_proto_rawDescData = file_student_lecture_proto_rawDesc
)

func file_student_lecture_proto_rawDescGZIP() []byte {
	file_student_lecture_proto_rawDescOnce.Do(func() {
		file_student_lecture_proto_rawDescData = protoimpl.X.CompressGZIP(file_student_lecture_proto_rawDescData)
	})
	return file_student_lecture_proto_rawDescData
}

var file_student_lecture_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_student_lecture_proto_goTypes = []interface{}{
	(*Meta)(nil),                                // 0: student_lecture.Meta
	(*Pagination)(nil),                          // 1: student_lecture.Pagination
	(*GetHistoryRequest)(nil),                   // 2: student_lecture.GetHistoryRequest
	(*GetHistoryResponseData)(nil),              // 3: student_lecture.GetHistoryResponseData
	(*GetHistoryResponse)(nil),                  // 4: student_lecture.GetHistoryResponse
	(*AttendAutonomousLectureRequest)(nil),      // 5: student_lecture.AttendAutonomousLectureRequest
	(*AttendAutonomousLectureResponseData)(nil), // 6: student_lecture.AttendAutonomousLectureResponseData
	(*AttendAutonomousLectureResponse)(nil),     // 7: student_lecture.AttendAutonomousLectureResponse
}
var file_student_lecture_proto_depIdxs = []int32{
	0, // 0: student_lecture.GetHistoryResponse.meta:type_name -> student_lecture.Meta
	1, // 1: student_lecture.GetHistoryResponse.pagination:type_name -> student_lecture.Pagination
	3, // 2: student_lecture.GetHistoryResponse.data:type_name -> student_lecture.GetHistoryResponseData
	0, // 3: student_lecture.AttendAutonomousLectureResponse.meta:type_name -> student_lecture.Meta
	6, // 4: student_lecture.AttendAutonomousLectureResponse.data:type_name -> student_lecture.AttendAutonomousLectureResponseData
	5, // 5: student_lecture.StudentLectureHandler.AttendAutonomousLecture:input_type -> student_lecture.AttendAutonomousLectureRequest
	2, // 6: student_lecture.StudentLectureHandler.GetHistory:input_type -> student_lecture.GetHistoryRequest
	7, // 7: student_lecture.StudentLectureHandler.AttendAutonomousLecture:output_type -> student_lecture.AttendAutonomousLectureResponse
	4, // 8: student_lecture.StudentLectureHandler.GetHistory:output_type -> student_lecture.GetHistoryResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_student_lecture_proto_init() }
func file_student_lecture_proto_init() {
	if File_student_lecture_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_student_lecture_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_student_lecture_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_student_lecture_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoryRequest); i {
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
		file_student_lecture_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoryResponseData); i {
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
		file_student_lecture_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoryResponse); i {
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
		file_student_lecture_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttendAutonomousLectureRequest); i {
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
		file_student_lecture_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttendAutonomousLectureResponseData); i {
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
		file_student_lecture_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttendAutonomousLectureResponse); i {
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
			RawDescriptor: file_student_lecture_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_student_lecture_proto_goTypes,
		DependencyIndexes: file_student_lecture_proto_depIdxs,
		MessageInfos:      file_student_lecture_proto_msgTypes,
	}.Build()
	File_student_lecture_proto = out.File
	file_student_lecture_proto_rawDesc = nil
	file_student_lecture_proto_goTypes = nil
	file_student_lecture_proto_depIdxs = nil
}
