// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: lecturer/general.proto

package general

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
		mi := &file_lecturer_general_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_lecturer_general_proto_msgTypes[0]
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
	return file_lecturer_general_proto_rawDescGZIP(), []int{0}
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
		mi := &file_lecturer_general_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_lecturer_general_proto_msgTypes[1]
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
	return file_lecturer_general_proto_rawDescGZIP(), []int{1}
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

type GetSemesterSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSemesterSummaryRequest) Reset() {
	*x = GetSemesterSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lecturer_general_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSemesterSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSemesterSummaryRequest) ProtoMessage() {}

func (x *GetSemesterSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lecturer_general_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSemesterSummaryRequest.ProtoReflect.Descriptor instead.
func (*GetSemesterSummaryRequest) Descriptor() ([]byte, []int) {
	return file_lecturer_general_proto_rawDescGZIP(), []int{2}
}

type GetSemesterSummaryResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudyPlanApprovalStartDate   string `protobuf:"bytes,1,opt,name=study_plan_approval_start_date,json=studyPlanApprovalStartDate,proto3" json:"study_plan_approval_start_date"`
	StudyPlanApprovalEndDate     string `protobuf:"bytes,2,opt,name=study_plan_approval_end_date,json=studyPlanApprovalEndDate,proto3" json:"study_plan_approval_end_date"`
	AcademicGuidanceTotalStudent uint32 `protobuf:"varint,3,opt,name=academic_guidance_total_student,json=academicGuidanceTotalStudent,proto3" json:"academic_guidance_total_student"`
	TotalClass                   uint32 `protobuf:"varint,4,opt,name=total_class,json=totalClass,proto3" json:"total_class"`
	SchoolYear                   string `protobuf:"bytes,5,opt,name=school_year,json=schoolYear,proto3" json:"school_year"`
	SemesterType                 string `protobuf:"bytes,6,opt,name=semester_type,json=semesterType,proto3" json:"semester_type"`
	GradingStartDate             string `protobuf:"bytes,7,opt,name=grading_start_date,json=gradingStartDate,proto3" json:"grading_start_date"`
	GradingEndDate               string `protobuf:"bytes,8,opt,name=grading_end_date,json=gradingEndDate,proto3" json:"grading_end_date"`
	SemesterId                   string `protobuf:"bytes,9,opt,name=semester_id,json=semesterId,proto3" json:"semester_id"`
}

func (x *GetSemesterSummaryResponseData) Reset() {
	*x = GetSemesterSummaryResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lecturer_general_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSemesterSummaryResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSemesterSummaryResponseData) ProtoMessage() {}

func (x *GetSemesterSummaryResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_lecturer_general_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSemesterSummaryResponseData.ProtoReflect.Descriptor instead.
func (*GetSemesterSummaryResponseData) Descriptor() ([]byte, []int) {
	return file_lecturer_general_proto_rawDescGZIP(), []int{3}
}

func (x *GetSemesterSummaryResponseData) GetStudyPlanApprovalStartDate() string {
	if x != nil {
		return x.StudyPlanApprovalStartDate
	}
	return ""
}

func (x *GetSemesterSummaryResponseData) GetStudyPlanApprovalEndDate() string {
	if x != nil {
		return x.StudyPlanApprovalEndDate
	}
	return ""
}

func (x *GetSemesterSummaryResponseData) GetAcademicGuidanceTotalStudent() uint32 {
	if x != nil {
		return x.AcademicGuidanceTotalStudent
	}
	return 0
}

func (x *GetSemesterSummaryResponseData) GetTotalClass() uint32 {
	if x != nil {
		return x.TotalClass
	}
	return 0
}

func (x *GetSemesterSummaryResponseData) GetSchoolYear() string {
	if x != nil {
		return x.SchoolYear
	}
	return ""
}

func (x *GetSemesterSummaryResponseData) GetSemesterType() string {
	if x != nil {
		return x.SemesterType
	}
	return ""
}

func (x *GetSemesterSummaryResponseData) GetGradingStartDate() string {
	if x != nil {
		return x.GradingStartDate
	}
	return ""
}

func (x *GetSemesterSummaryResponseData) GetGradingEndDate() string {
	if x != nil {
		return x.GradingEndDate
	}
	return ""
}

func (x *GetSemesterSummaryResponseData) GetSemesterId() string {
	if x != nil {
		return x.SemesterId
	}
	return ""
}

type GetSemesterSummaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Meta                           `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta"`
	Data *GetSemesterSummaryResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *GetSemesterSummaryResponse) Reset() {
	*x = GetSemesterSummaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lecturer_general_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSemesterSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSemesterSummaryResponse) ProtoMessage() {}

func (x *GetSemesterSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lecturer_general_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSemesterSummaryResponse.ProtoReflect.Descriptor instead.
func (*GetSemesterSummaryResponse) Descriptor() ([]byte, []int) {
	return file_lecturer_general_proto_rawDescGZIP(), []int{4}
}

func (x *GetSemesterSummaryResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *GetSemesterSummaryResponse) GetData() *GetSemesterSummaryResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lecturer_general_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lecturer_general_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_lecturer_general_proto_rawDescGZIP(), []int{5}
}

type GetProfileResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	IdNationalLecturer string `protobuf:"bytes,2,opt,name=id_national_lecturer,json=idNationalLecturer,proto3" json:"id_national_lecturer"`
	Name               string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	FrontTitle         string `protobuf:"bytes,4,opt,name=front_title,json=frontTitle,proto3" json:"front_title"`
	BackDegree         string `protobuf:"bytes,5,opt,name=back_degree,json=backDegree,proto3" json:"back_degree"`
	StudyProgramId     string `protobuf:"bytes,6,opt,name=study_program_id,json=studyProgramId,proto3" json:"study_program_id"`
	StudyProgramName   string `protobuf:"bytes,7,opt,name=study_program_name,json=studyProgramName,proto3" json:"study_program_name"`
	BirthDate          string `protobuf:"bytes,8,opt,name=birth_date,json=birthDate,proto3" json:"birth_date"`
	BirthRegencyId     uint32 `protobuf:"varint,9,opt,name=birth_regency_id,json=birthRegencyId,proto3" json:"birth_regency_id"`
	BirthRegencyName   string `protobuf:"bytes,10,opt,name=birth_regency_name,json=birthRegencyName,proto3" json:"birth_regency_name"`
	BirthCountryId     uint32 `protobuf:"varint,11,opt,name=birth_country_id,json=birthCountryId,proto3" json:"birth_country_id"`
	BirthCountryName   string `protobuf:"bytes,12,opt,name=birth_country_name,json=birthCountryName,proto3" json:"birth_country_name"`
	Sex                string `protobuf:"bytes,13,opt,name=sex,proto3" json:"sex"`
	Religion           string `protobuf:"bytes,14,opt,name=religion,proto3" json:"religion"`
	Address            string `protobuf:"bytes,15,opt,name=address,proto3" json:"address"`
	RegencyId          uint32 `protobuf:"varint,16,opt,name=regency_id,json=regencyId,proto3" json:"regency_id"`
	RegencyName        string `protobuf:"bytes,17,opt,name=regency_name,json=regencyName,proto3" json:"regency_name"`
	CountryId          uint32 `protobuf:"varint,18,opt,name=country_id,json=countryId,proto3" json:"country_id"`
	CountryName        string `protobuf:"bytes,19,opt,name=country_name,json=countryName,proto3" json:"country_name"`
	PostalCode         string `protobuf:"bytes,20,opt,name=postal_code,json=postalCode,proto3" json:"postal_code"`
	PhoneNumber        string `protobuf:"bytes,21,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number"`
	Fax                string `protobuf:"bytes,22,opt,name=fax,proto3" json:"fax"`
	MobilePhoneNumber  string `protobuf:"bytes,23,opt,name=mobile_phone_number,json=mobilePhoneNumber,proto3" json:"mobile_phone_number"`
	OfficePhoneNumber  string `protobuf:"bytes,24,opt,name=office_phone_number,json=officePhoneNumber,proto3" json:"office_phone_number"`
	AcademicPosition   string `protobuf:"bytes,25,opt,name=academic_position,json=academicPosition,proto3" json:"academic_position"`
	EmploymentStatus   string `protobuf:"bytes,26,opt,name=employment_status,json=employmentStatus,proto3" json:"employment_status"`
	Status             string `protobuf:"bytes,27,opt,name=status,proto3" json:"status"`
}

func (x *GetProfileResponseData) Reset() {
	*x = GetProfileResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lecturer_general_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponseData) ProtoMessage() {}

func (x *GetProfileResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_lecturer_general_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponseData.ProtoReflect.Descriptor instead.
func (*GetProfileResponseData) Descriptor() ([]byte, []int) {
	return file_lecturer_general_proto_rawDescGZIP(), []int{6}
}

func (x *GetProfileResponseData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProfileResponseData) GetIdNationalLecturer() string {
	if x != nil {
		return x.IdNationalLecturer
	}
	return ""
}

func (x *GetProfileResponseData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProfileResponseData) GetFrontTitle() string {
	if x != nil {
		return x.FrontTitle
	}
	return ""
}

func (x *GetProfileResponseData) GetBackDegree() string {
	if x != nil {
		return x.BackDegree
	}
	return ""
}

func (x *GetProfileResponseData) GetStudyProgramId() string {
	if x != nil {
		return x.StudyProgramId
	}
	return ""
}

func (x *GetProfileResponseData) GetStudyProgramName() string {
	if x != nil {
		return x.StudyProgramName
	}
	return ""
}

func (x *GetProfileResponseData) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *GetProfileResponseData) GetBirthRegencyId() uint32 {
	if x != nil {
		return x.BirthRegencyId
	}
	return 0
}

func (x *GetProfileResponseData) GetBirthRegencyName() string {
	if x != nil {
		return x.BirthRegencyName
	}
	return ""
}

func (x *GetProfileResponseData) GetBirthCountryId() uint32 {
	if x != nil {
		return x.BirthCountryId
	}
	return 0
}

func (x *GetProfileResponseData) GetBirthCountryName() string {
	if x != nil {
		return x.BirthCountryName
	}
	return ""
}

func (x *GetProfileResponseData) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *GetProfileResponseData) GetReligion() string {
	if x != nil {
		return x.Religion
	}
	return ""
}

func (x *GetProfileResponseData) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetProfileResponseData) GetRegencyId() uint32 {
	if x != nil {
		return x.RegencyId
	}
	return 0
}

func (x *GetProfileResponseData) GetRegencyName() string {
	if x != nil {
		return x.RegencyName
	}
	return ""
}

func (x *GetProfileResponseData) GetCountryId() uint32 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *GetProfileResponseData) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *GetProfileResponseData) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *GetProfileResponseData) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *GetProfileResponseData) GetFax() string {
	if x != nil {
		return x.Fax
	}
	return ""
}

func (x *GetProfileResponseData) GetMobilePhoneNumber() string {
	if x != nil {
		return x.MobilePhoneNumber
	}
	return ""
}

func (x *GetProfileResponseData) GetOfficePhoneNumber() string {
	if x != nil {
		return x.OfficePhoneNumber
	}
	return ""
}

func (x *GetProfileResponseData) GetAcademicPosition() string {
	if x != nil {
		return x.AcademicPosition
	}
	return ""
}

func (x *GetProfileResponseData) GetEmploymentStatus() string {
	if x != nil {
		return x.EmploymentStatus
	}
	return ""
}

func (x *GetProfileResponseData) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Meta                   `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta"`
	Data *GetProfileResponseData `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lecturer_general_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lecturer_general_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_lecturer_general_proto_rawDescGZIP(), []int{7}
}

func (x *GetProfileResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *GetProfileResponse) GetData() *GetProfileResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_lecturer_general_proto protoreflect.FileDescriptor

var file_lecturer_general_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x72, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x22, 0x4c, 0x0a, 0x04, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x72, 0x65, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x70, 0x72, 0x65, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22,
	0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xcb, 0x03, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x42, 0x0a, 0x1e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x73, 0x74, 0x75, 0x64, 0x79, 0x50, 0x6c,
	0x61, 0x6e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x1c, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x70, 0x6c, 0x61,
	0x6e, 0x5f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x50, 0x6c, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x45, 0x0a, 0x1f, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x5f,
	0x67, 0x75, 0x69, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1c, 0x61, 0x63,
	0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x47, 0x75, 0x69, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x2c, 0x0a, 0x12, 0x67, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x67,
	0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x28, 0x0a, 0x10, 0x67, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x72, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6d,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x72, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x44, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x13, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xcb, 0x07, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x69,
	0x64, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x64, 0x4e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x65,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x67,
	0x72, 0x65, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x74, 0x75, 0x64, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x12, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x62, 0x69, 0x72, 0x74, 0x68, 0x52, 0x65, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x62, 0x69, 0x72, 0x74, 0x68, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x69, 0x72, 0x74, 0x68, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65,
	0x78, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x63, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x61,
	0x78, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x61, 0x78, 0x12, 0x2e, 0x0a, 0x13,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13,
	0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11,
	0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69,
	0x63, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x7e,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x12, 0x3c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xe6,
	0x01, 0x0a, 0x16, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x71, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x2b, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x73, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lecturer_general_proto_rawDescOnce sync.Once
	file_lecturer_general_proto_rawDescData = file_lecturer_general_proto_rawDesc
)

func file_lecturer_general_proto_rawDescGZIP() []byte {
	file_lecturer_general_proto_rawDescOnce.Do(func() {
		file_lecturer_general_proto_rawDescData = protoimpl.X.CompressGZIP(file_lecturer_general_proto_rawDescData)
	})
	return file_lecturer_general_proto_rawDescData
}

var file_lecturer_general_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_lecturer_general_proto_goTypes = []interface{}{
	(*Meta)(nil),                           // 0: lecturer_general.Meta
	(*Pagination)(nil),                     // 1: lecturer_general.Pagination
	(*GetSemesterSummaryRequest)(nil),      // 2: lecturer_general.GetSemesterSummaryRequest
	(*GetSemesterSummaryResponseData)(nil), // 3: lecturer_general.GetSemesterSummaryResponseData
	(*GetSemesterSummaryResponse)(nil),     // 4: lecturer_general.GetSemesterSummaryResponse
	(*GetProfileRequest)(nil),              // 5: lecturer_general.GetProfileRequest
	(*GetProfileResponseData)(nil),         // 6: lecturer_general.GetProfileResponseData
	(*GetProfileResponse)(nil),             // 7: lecturer_general.GetProfileResponse
}
var file_lecturer_general_proto_depIdxs = []int32{
	0, // 0: lecturer_general.GetSemesterSummaryResponse.meta:type_name -> lecturer_general.Meta
	3, // 1: lecturer_general.GetSemesterSummaryResponse.data:type_name -> lecturer_general.GetSemesterSummaryResponseData
	0, // 2: lecturer_general.GetProfileResponse.meta:type_name -> lecturer_general.Meta
	6, // 3: lecturer_general.GetProfileResponse.data:type_name -> lecturer_general.GetProfileResponseData
	2, // 4: lecturer_general.LecturerGeneralHandler.GetSemesterSummary:input_type -> lecturer_general.GetSemesterSummaryRequest
	5, // 5: lecturer_general.LecturerGeneralHandler.GetProfile:input_type -> lecturer_general.GetProfileRequest
	4, // 6: lecturer_general.LecturerGeneralHandler.GetSemesterSummary:output_type -> lecturer_general.GetSemesterSummaryResponse
	7, // 7: lecturer_general.LecturerGeneralHandler.GetProfile:output_type -> lecturer_general.GetProfileResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_lecturer_general_proto_init() }
func file_lecturer_general_proto_init() {
	if File_lecturer_general_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lecturer_general_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_lecturer_general_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_lecturer_general_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSemesterSummaryRequest); i {
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
		file_lecturer_general_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSemesterSummaryResponseData); i {
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
		file_lecturer_general_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSemesterSummaryResponse); i {
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
		file_lecturer_general_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileRequest); i {
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
		file_lecturer_general_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileResponseData); i {
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
		file_lecturer_general_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileResponse); i {
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
			RawDescriptor: file_lecturer_general_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lecturer_general_proto_goTypes,
		DependencyIndexes: file_lecturer_general_proto_depIdxs,
		MessageInfos:      file_lecturer_general_proto_msgTypes,
	}.Build()
	File_lecturer_general_proto = out.File
	file_lecturer_general_proto_rawDesc = nil
	file_lecturer_general_proto_goTypes = nil
	file_lecturer_general_proto_depIdxs = nil
}
