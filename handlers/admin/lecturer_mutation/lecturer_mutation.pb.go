// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: admin/lecturer_mutation.proto

package lecturer_mutation

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
		mi := &file_admin_lecturer_mutation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_admin_lecturer_mutation_proto_msgTypes[0]
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
	return file_admin_lecturer_mutation_proto_rawDescGZIP(), []int{0}
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
		mi := &file_admin_lecturer_mutation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_admin_lecturer_mutation_proto_msgTypes[1]
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
	return file_admin_lecturer_mutation_proto_rawDescGZIP(), []int{1}
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

	Page               uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	Limit              uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit"`
	Search             string `protobuf:"bytes,3,opt,name=search,proto3" json:"search"`
	StudyProgramId     string `protobuf:"bytes,4,opt,name=study_program_id,json=studyProgramId,proto3" json:"study_program_id"`
	IdNationalLecturer string `protobuf:"bytes,5,opt,name=id_national_lecturer,json=idNationalLecturer,proto3" json:"id_national_lecturer"`
	SemesterId         string `protobuf:"bytes,6,opt,name=semester_id,json=semesterId,proto3" json:"semester_id"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_lecturer_mutation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_lecturer_mutation_proto_msgTypes[2]
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
	return file_admin_lecturer_mutation_proto_rawDescGZIP(), []int{2}
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

func (x *GetListRequest) GetStudyProgramId() string {
	if x != nil {
		return x.StudyProgramId
	}
	return ""
}

func (x *GetListRequest) GetIdNationalLecturer() string {
	if x != nil {
		return x.IdNationalLecturer
	}
	return ""
}

func (x *GetListRequest) GetSemesterId() string {
	if x != nil {
		return x.SemesterId
	}
	return ""
}

type GetListResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name                  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	IdNationalLecturer    string `protobuf:"bytes,3,opt,name=id_national_lecturer,json=idNationalLecturer,proto3" json:"id_national_lecturer"`
	FrontTitle            string `protobuf:"bytes,4,opt,name=front_title,json=frontTitle,proto3" json:"front_title"`
	BackDegree            string `protobuf:"bytes,5,opt,name=back_degree,json=backDegree,proto3" json:"back_degree"`
	SemesterSchoolYear    string `protobuf:"bytes,6,opt,name=semester_school_year,json=semesterSchoolYear,proto3" json:"semester_school_year"`
	SemesterType          string `protobuf:"bytes,7,opt,name=semester_type,json=semesterType,proto3" json:"semester_type"`
	DiktiStudyProgramCode string `protobuf:"bytes,8,opt,name=dikti_study_program_code,json=diktiStudyProgramCode,proto3" json:"dikti_study_program_code"`
	StudyProgramName      string `protobuf:"bytes,9,opt,name=study_program_name,json=studyProgramName,proto3" json:"study_program_name"`
	StudyLevelShortName   string `protobuf:"bytes,10,opt,name=study_level_short_name,json=studyLevelShortName,proto3" json:"study_level_short_name"`
	DiktiStudyProgramType string `protobuf:"bytes,11,opt,name=dikti_study_program_type,json=diktiStudyProgramType,proto3" json:"dikti_study_program_type"`
	MutationDate          string `protobuf:"bytes,12,opt,name=mutation_date,json=mutationDate,proto3" json:"mutation_date"`
	DecisionNumber        string `protobuf:"bytes,13,opt,name=decision_number,json=decisionNumber,proto3" json:"decision_number"`
	Destination           string `protobuf:"bytes,14,opt,name=destination,proto3" json:"destination"`
}

func (x *GetListResponseData) Reset() {
	*x = GetListResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_lecturer_mutation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponseData) ProtoMessage() {}

func (x *GetListResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_admin_lecturer_mutation_proto_msgTypes[3]
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
	return file_admin_lecturer_mutation_proto_rawDescGZIP(), []int{3}
}

func (x *GetListResponseData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetListResponseData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetListResponseData) GetIdNationalLecturer() string {
	if x != nil {
		return x.IdNationalLecturer
	}
	return ""
}

func (x *GetListResponseData) GetFrontTitle() string {
	if x != nil {
		return x.FrontTitle
	}
	return ""
}

func (x *GetListResponseData) GetBackDegree() string {
	if x != nil {
		return x.BackDegree
	}
	return ""
}

func (x *GetListResponseData) GetSemesterSchoolYear() string {
	if x != nil {
		return x.SemesterSchoolYear
	}
	return ""
}

func (x *GetListResponseData) GetSemesterType() string {
	if x != nil {
		return x.SemesterType
	}
	return ""
}

func (x *GetListResponseData) GetDiktiStudyProgramCode() string {
	if x != nil {
		return x.DiktiStudyProgramCode
	}
	return ""
}

func (x *GetListResponseData) GetStudyProgramName() string {
	if x != nil {
		return x.StudyProgramName
	}
	return ""
}

func (x *GetListResponseData) GetStudyLevelShortName() string {
	if x != nil {
		return x.StudyLevelShortName
	}
	return ""
}

func (x *GetListResponseData) GetDiktiStudyProgramType() string {
	if x != nil {
		return x.DiktiStudyProgramType
	}
	return ""
}

func (x *GetListResponseData) GetMutationDate() string {
	if x != nil {
		return x.MutationDate
	}
	return ""
}

func (x *GetListResponseData) GetDecisionNumber() string {
	if x != nil {
		return x.DecisionNumber
	}
	return ""
}

func (x *GetListResponseData) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
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
		mi := &file_admin_lecturer_mutation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_lecturer_mutation_proto_msgTypes[4]
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
	return file_admin_lecturer_mutation_proto_rawDescGZIP(), []int{4}
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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LecturerId     string `protobuf:"bytes,1,opt,name=lecturer_id,json=lecturerId,proto3" json:"lecturer_id"`
	SemesterId     string `protobuf:"bytes,2,opt,name=semester_id,json=semesterId,proto3" json:"semester_id"`
	MutationDate   string `protobuf:"bytes,3,opt,name=mutation_date,json=mutationDate,proto3" json:"mutation_date"`
	DecisionNumber string `protobuf:"bytes,4,opt,name=decision_number,json=decisionNumber,proto3" json:"decision_number"`
	Destination    string `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_lecturer_mutation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_lecturer_mutation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_admin_lecturer_mutation_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRequest) GetLecturerId() string {
	if x != nil {
		return x.LecturerId
	}
	return ""
}

func (x *CreateRequest) GetSemesterId() string {
	if x != nil {
		return x.SemesterId
	}
	return ""
}

func (x *CreateRequest) GetMutationDate() string {
	if x != nil {
		return x.MutationDate
	}
	return ""
}

func (x *CreateRequest) GetDecisionNumber() string {
	if x != nil {
		return x.DecisionNumber
	}
	return ""
}

func (x *CreateRequest) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type CreateResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResponseData) Reset() {
	*x = CreateResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_lecturer_mutation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponseData) ProtoMessage() {}

func (x *CreateResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_admin_lecturer_mutation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponseData.ProtoReflect.Descriptor instead.
func (*CreateResponseData) Descriptor() ([]byte, []int) {
	return file_admin_lecturer_mutation_proto_rawDescGZIP(), []int{6}
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Meta               `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta"`
	Data *CreateResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_lecturer_mutation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_lecturer_mutation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_admin_lecturer_mutation_proto_rawDescGZIP(), []int{7}
}

func (x *CreateResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *CreateResponse) GetData() *CreateResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_admin_lecturer_mutation_proto protoreflect.FileDescriptor

var file_admin_lecturer_mutation_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72,
	0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f,
	0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61,
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
	0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xcf, 0x01,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x74,
	0x75, 0x64, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14,
	0x69, 0x64, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x64, 0x4e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xc9, 0x04, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x69,
	0x64, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x64, 0x4e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x12,
	0x30, 0x0a, 0x14, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73,
	0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61,
	0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x64, 0x69, 0x6b, 0x74, 0x69, 0x5f,
	0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x64, 0x69, 0x6b, 0x74, 0x69, 0x53,
	0x74, 0x75, 0x64, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x2c, 0x0a, 0x12, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x75,
	0x64, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a,
	0x16, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73,
	0x74, 0x75, 0x64, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x64, 0x69, 0x6b, 0x74, 0x69, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x79, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x64, 0x69, 0x6b, 0x74, 0x69, 0x53, 0x74, 0x75, 0x64, 0x79,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xcb, 0x01, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x43, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc1, 0x01, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x14, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x84, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xdb, 0x01, 0x0a, 0x1c, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x4d, 0x75, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72,
	0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x6d,
	0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_lecturer_mutation_proto_rawDescOnce sync.Once
	file_admin_lecturer_mutation_proto_rawDescData = file_admin_lecturer_mutation_proto_rawDesc
)

func file_admin_lecturer_mutation_proto_rawDescGZIP() []byte {
	file_admin_lecturer_mutation_proto_rawDescOnce.Do(func() {
		file_admin_lecturer_mutation_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_lecturer_mutation_proto_rawDescData)
	})
	return file_admin_lecturer_mutation_proto_rawDescData
}

var file_admin_lecturer_mutation_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_admin_lecturer_mutation_proto_goTypes = []interface{}{
	(*Meta)(nil),                // 0: admin_lecturer_mutation.Meta
	(*Pagination)(nil),          // 1: admin_lecturer_mutation.Pagination
	(*GetListRequest)(nil),      // 2: admin_lecturer_mutation.GetListRequest
	(*GetListResponseData)(nil), // 3: admin_lecturer_mutation.GetListResponseData
	(*GetListResponse)(nil),     // 4: admin_lecturer_mutation.GetListResponse
	(*CreateRequest)(nil),       // 5: admin_lecturer_mutation.CreateRequest
	(*CreateResponseData)(nil),  // 6: admin_lecturer_mutation.CreateResponseData
	(*CreateResponse)(nil),      // 7: admin_lecturer_mutation.CreateResponse
}
var file_admin_lecturer_mutation_proto_depIdxs = []int32{
	0, // 0: admin_lecturer_mutation.GetListResponse.meta:type_name -> admin_lecturer_mutation.Meta
	1, // 1: admin_lecturer_mutation.GetListResponse.pagination:type_name -> admin_lecturer_mutation.Pagination
	3, // 2: admin_lecturer_mutation.GetListResponse.data:type_name -> admin_lecturer_mutation.GetListResponseData
	0, // 3: admin_lecturer_mutation.CreateResponse.meta:type_name -> admin_lecturer_mutation.Meta
	6, // 4: admin_lecturer_mutation.CreateResponse.data:type_name -> admin_lecturer_mutation.CreateResponseData
	2, // 5: admin_lecturer_mutation.AdminLecturerMutationHandler.GetList:input_type -> admin_lecturer_mutation.GetListRequest
	5, // 6: admin_lecturer_mutation.AdminLecturerMutationHandler.Create:input_type -> admin_lecturer_mutation.CreateRequest
	4, // 7: admin_lecturer_mutation.AdminLecturerMutationHandler.GetList:output_type -> admin_lecturer_mutation.GetListResponse
	7, // 8: admin_lecturer_mutation.AdminLecturerMutationHandler.Create:output_type -> admin_lecturer_mutation.CreateResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_admin_lecturer_mutation_proto_init() }
func file_admin_lecturer_mutation_proto_init() {
	if File_admin_lecturer_mutation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_lecturer_mutation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_admin_lecturer_mutation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_admin_lecturer_mutation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_admin_lecturer_mutation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_admin_lecturer_mutation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_admin_lecturer_mutation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_admin_lecturer_mutation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponseData); i {
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
		file_admin_lecturer_mutation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
			RawDescriptor: file_admin_lecturer_mutation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_lecturer_mutation_proto_goTypes,
		DependencyIndexes: file_admin_lecturer_mutation_proto_depIdxs,
		MessageInfos:      file_admin_lecturer_mutation_proto_msgTypes,
	}.Build()
	File_admin_lecturer_mutation_proto = out.File
	file_admin_lecturer_mutation_proto_rawDesc = nil
	file_admin_lecturer_mutation_proto_goTypes = nil
	file_admin_lecturer_mutation_proto_depIdxs = nil
}