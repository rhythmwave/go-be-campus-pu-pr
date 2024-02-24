// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pmb/student.proto

package student

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
		mi := &file_pmb_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_pmb_student_proto_msgTypes[0]
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
	return file_pmb_student_proto_rawDescGZIP(), []int{0}
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
		mi := &file_pmb_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_pmb_student_proto_msgTypes[1]
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
	return file_pmb_student_proto_rawDescGZIP(), []int{1}
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

type BulkCreateRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NimNumber                       int64  `protobuf:"varint,1,opt,name=nim_number,json=nimNumber,proto3" json:"nim_number"`
	Name                            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Sex                             string `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex"`
	MaritalStatus                   string `protobuf:"bytes,4,opt,name=marital_status,json=maritalStatus,proto3" json:"marital_status"`
	BirthRegencyId                  uint32 `protobuf:"varint,5,opt,name=birth_regency_id,json=birthRegencyId,proto3" json:"birth_regency_id"`
	BirthDate                       string `protobuf:"bytes,6,opt,name=birth_date,json=birthDate,proto3" json:"birth_date"`
	Religion                        string `protobuf:"bytes,7,opt,name=religion,proto3" json:"religion"`
	Address                         string `protobuf:"bytes,8,opt,name=address,proto3" json:"address"`
	Rt                              string `protobuf:"bytes,9,opt,name=rt,proto3" json:"rt"`
	Rw                              string `protobuf:"bytes,10,opt,name=rw,proto3" json:"rw"`
	VillageId                       uint32 `protobuf:"varint,11,opt,name=village_id,json=villageId,proto3" json:"village_id"`
	PostalCode                      string `protobuf:"bytes,12,opt,name=postal_code,json=postalCode,proto3" json:"postal_code"`
	IdNumber                        string `protobuf:"bytes,13,opt,name=id_number,json=idNumber,proto3" json:"id_number"`
	NisnNumber                      string `protobuf:"bytes,14,opt,name=nisn_number,json=nisnNumber,proto3" json:"nisn_number"`
	MobilePhoneNumber               string `protobuf:"bytes,15,opt,name=mobile_phone_number,json=mobilePhoneNumber,proto3" json:"mobile_phone_number"`
	Nationality                     string `protobuf:"bytes,16,opt,name=nationality,proto3" json:"nationality"`
	DiktiStudyProgramCode           string `protobuf:"bytes,17,opt,name=dikti_study_program_code,json=diktiStudyProgramCode,proto3" json:"dikti_study_program_code"`
	SchoolName                      string `protobuf:"bytes,18,opt,name=school_name,json=schoolName,proto3" json:"school_name"`
	SchoolAddress                   string `protobuf:"bytes,19,opt,name=school_address,json=schoolAddress,proto3" json:"school_address"`
	SchoolProvinceId                uint32 `protobuf:"varint,20,opt,name=school_province_id,json=schoolProvinceId,proto3" json:"school_province_id"`
	SchoolMajor                     string `protobuf:"bytes,21,opt,name=school_major,json=schoolMajor,proto3" json:"school_major"`
	SchoolType                      string `protobuf:"bytes,22,opt,name=school_type,json=schoolType,proto3" json:"school_type"`
	SchoolGraduationYear            string `protobuf:"bytes,23,opt,name=school_graduation_year,json=schoolGraduationYear,proto3" json:"school_graduation_year"`
	FatherName                      string `protobuf:"bytes,24,opt,name=father_name,json=fatherName,proto3" json:"father_name"`
	FatherIdNumber                  string `protobuf:"bytes,25,opt,name=father_id_number,json=fatherIdNumber,proto3" json:"father_id_number"`
	FatherBirthDate                 string `protobuf:"bytes,26,opt,name=father_birth_date,json=fatherBirthDate,proto3" json:"father_birth_date"`
	FatherFinalAcademicBackground   string `protobuf:"bytes,27,opt,name=father_final_academic_background,json=fatherFinalAcademicBackground,proto3" json:"father_final_academic_background"`
	FatherOccupation                string `protobuf:"bytes,28,opt,name=father_occupation,json=fatherOccupation,proto3" json:"father_occupation"`
	MotherName                      string `protobuf:"bytes,29,opt,name=mother_name,json=motherName,proto3" json:"mother_name"`
	MotherIdNumber                  string `protobuf:"bytes,30,opt,name=mother_id_number,json=motherIdNumber,proto3" json:"mother_id_number"`
	MotherBirthDate                 string `protobuf:"bytes,31,opt,name=mother_birth_date,json=motherBirthDate,proto3" json:"mother_birth_date"`
	MotherFinalAcademicBackground   string `protobuf:"bytes,32,opt,name=mother_final_academic_background,json=motherFinalAcademicBackground,proto3" json:"mother_final_academic_background"`
	MotherOccupation                string `protobuf:"bytes,33,opt,name=mother_occupation,json=motherOccupation,proto3" json:"mother_occupation"`
	GuardianName                    string `protobuf:"bytes,34,opt,name=guardian_name,json=guardianName,proto3" json:"guardian_name"`
	GuardianIdNumber                string `protobuf:"bytes,35,opt,name=guardian_id_number,json=guardianIdNumber,proto3" json:"guardian_id_number"`
	GuardianBirthDate               string `protobuf:"bytes,36,opt,name=guardian_birth_date,json=guardianBirthDate,proto3" json:"guardian_birth_date"`
	GuardianFinalAcademicBackground string `protobuf:"bytes,37,opt,name=guardian_final_academic_background,json=guardianFinalAcademicBackground,proto3" json:"guardian_final_academic_background"`
	GuardianOccupation              string `protobuf:"bytes,38,opt,name=guardian_occupation,json=guardianOccupation,proto3" json:"guardian_occupation"`
}

func (x *BulkCreateRequestData) Reset() {
	*x = BulkCreateRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmb_student_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkCreateRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkCreateRequestData) ProtoMessage() {}

func (x *BulkCreateRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_pmb_student_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkCreateRequestData.ProtoReflect.Descriptor instead.
func (*BulkCreateRequestData) Descriptor() ([]byte, []int) {
	return file_pmb_student_proto_rawDescGZIP(), []int{2}
}

func (x *BulkCreateRequestData) GetNimNumber() int64 {
	if x != nil {
		return x.NimNumber
	}
	return 0
}

func (x *BulkCreateRequestData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BulkCreateRequestData) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *BulkCreateRequestData) GetMaritalStatus() string {
	if x != nil {
		return x.MaritalStatus
	}
	return ""
}

func (x *BulkCreateRequestData) GetBirthRegencyId() uint32 {
	if x != nil {
		return x.BirthRegencyId
	}
	return 0
}

func (x *BulkCreateRequestData) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *BulkCreateRequestData) GetReligion() string {
	if x != nil {
		return x.Religion
	}
	return ""
}

func (x *BulkCreateRequestData) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BulkCreateRequestData) GetRt() string {
	if x != nil {
		return x.Rt
	}
	return ""
}

func (x *BulkCreateRequestData) GetRw() string {
	if x != nil {
		return x.Rw
	}
	return ""
}

func (x *BulkCreateRequestData) GetVillageId() uint32 {
	if x != nil {
		return x.VillageId
	}
	return 0
}

func (x *BulkCreateRequestData) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *BulkCreateRequestData) GetIdNumber() string {
	if x != nil {
		return x.IdNumber
	}
	return ""
}

func (x *BulkCreateRequestData) GetNisnNumber() string {
	if x != nil {
		return x.NisnNumber
	}
	return ""
}

func (x *BulkCreateRequestData) GetMobilePhoneNumber() string {
	if x != nil {
		return x.MobilePhoneNumber
	}
	return ""
}

func (x *BulkCreateRequestData) GetNationality() string {
	if x != nil {
		return x.Nationality
	}
	return ""
}

func (x *BulkCreateRequestData) GetDiktiStudyProgramCode() string {
	if x != nil {
		return x.DiktiStudyProgramCode
	}
	return ""
}

func (x *BulkCreateRequestData) GetSchoolName() string {
	if x != nil {
		return x.SchoolName
	}
	return ""
}

func (x *BulkCreateRequestData) GetSchoolAddress() string {
	if x != nil {
		return x.SchoolAddress
	}
	return ""
}

func (x *BulkCreateRequestData) GetSchoolProvinceId() uint32 {
	if x != nil {
		return x.SchoolProvinceId
	}
	return 0
}

func (x *BulkCreateRequestData) GetSchoolMajor() string {
	if x != nil {
		return x.SchoolMajor
	}
	return ""
}

func (x *BulkCreateRequestData) GetSchoolType() string {
	if x != nil {
		return x.SchoolType
	}
	return ""
}

func (x *BulkCreateRequestData) GetSchoolGraduationYear() string {
	if x != nil {
		return x.SchoolGraduationYear
	}
	return ""
}

func (x *BulkCreateRequestData) GetFatherName() string {
	if x != nil {
		return x.FatherName
	}
	return ""
}

func (x *BulkCreateRequestData) GetFatherIdNumber() string {
	if x != nil {
		return x.FatherIdNumber
	}
	return ""
}

func (x *BulkCreateRequestData) GetFatherBirthDate() string {
	if x != nil {
		return x.FatherBirthDate
	}
	return ""
}

func (x *BulkCreateRequestData) GetFatherFinalAcademicBackground() string {
	if x != nil {
		return x.FatherFinalAcademicBackground
	}
	return ""
}

func (x *BulkCreateRequestData) GetFatherOccupation() string {
	if x != nil {
		return x.FatherOccupation
	}
	return ""
}

func (x *BulkCreateRequestData) GetMotherName() string {
	if x != nil {
		return x.MotherName
	}
	return ""
}

func (x *BulkCreateRequestData) GetMotherIdNumber() string {
	if x != nil {
		return x.MotherIdNumber
	}
	return ""
}

func (x *BulkCreateRequestData) GetMotherBirthDate() string {
	if x != nil {
		return x.MotherBirthDate
	}
	return ""
}

func (x *BulkCreateRequestData) GetMotherFinalAcademicBackground() string {
	if x != nil {
		return x.MotherFinalAcademicBackground
	}
	return ""
}

func (x *BulkCreateRequestData) GetMotherOccupation() string {
	if x != nil {
		return x.MotherOccupation
	}
	return ""
}

func (x *BulkCreateRequestData) GetGuardianName() string {
	if x != nil {
		return x.GuardianName
	}
	return ""
}

func (x *BulkCreateRequestData) GetGuardianIdNumber() string {
	if x != nil {
		return x.GuardianIdNumber
	}
	return ""
}

func (x *BulkCreateRequestData) GetGuardianBirthDate() string {
	if x != nil {
		return x.GuardianBirthDate
	}
	return ""
}

func (x *BulkCreateRequestData) GetGuardianFinalAcademicBackground() string {
	if x != nil {
		return x.GuardianFinalAcademicBackground
	}
	return ""
}

func (x *BulkCreateRequestData) GetGuardianOccupation() string {
	if x != nil {
		return x.GuardianOccupation
	}
	return ""
}

type BulkCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*BulkCreateRequestData `protobuf:"bytes,1,rep,name=data,proto3" json:"data"`
}

func (x *BulkCreateRequest) Reset() {
	*x = BulkCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmb_student_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkCreateRequest) ProtoMessage() {}

func (x *BulkCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pmb_student_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkCreateRequest.ProtoReflect.Descriptor instead.
func (*BulkCreateRequest) Descriptor() ([]byte, []int) {
	return file_pmb_student_proto_rawDescGZIP(), []int{3}
}

func (x *BulkCreateRequest) GetData() []*BulkCreateRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type BulkCreateResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	NimNumber int64  `protobuf:"varint,2,opt,name=nim_number,json=nimNumber,proto3" json:"nim_number"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password"`
}

func (x *BulkCreateResponseData) Reset() {
	*x = BulkCreateResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmb_student_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkCreateResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkCreateResponseData) ProtoMessage() {}

func (x *BulkCreateResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_pmb_student_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkCreateResponseData.ProtoReflect.Descriptor instead.
func (*BulkCreateResponseData) Descriptor() ([]byte, []int) {
	return file_pmb_student_proto_rawDescGZIP(), []int{4}
}

func (x *BulkCreateResponseData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BulkCreateResponseData) GetNimNumber() int64 {
	if x != nil {
		return x.NimNumber
	}
	return 0
}

func (x *BulkCreateResponseData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type BulkCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Meta                     `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta"`
	Data []*BulkCreateResponseData `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
}

func (x *BulkCreateResponse) Reset() {
	*x = BulkCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pmb_student_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkCreateResponse) ProtoMessage() {}

func (x *BulkCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pmb_student_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkCreateResponse.ProtoReflect.Descriptor instead.
func (*BulkCreateResponse) Descriptor() ([]byte, []int) {
	return file_pmb_student_proto_rawDescGZIP(), []int{5}
}

func (x *BulkCreateResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *BulkCreateResponse) GetData() []*BulkCreateResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pmb_student_proto protoreflect.FileDescriptor

var file_pmb_student_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6d, 0x62, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x6d, 0x62, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x22, 0x4c, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xa4,
	0x01, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x72, 0x65, 0x76, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x72, 0x65, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xf6, 0x0b, 0x0a, 0x15, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1d, 0x0a, 0x0a, 0x6e, 0x69, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x69, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x65, 0x78, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x69, 0x74, 0x61, 0x6c, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61,
	0x72, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x62, 0x69, 0x72, 0x74, 0x68, 0x52, 0x65, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x77,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x72, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69,
	0x6c, 0x6c, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x76, 0x69, 0x6c, 0x6c, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x64,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x69, 0x73, 0x6e, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x69,
	0x73, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x18, 0x64, 0x69,
	0x6b, 0x74, 0x69, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x64, 0x69,
	0x6b, 0x74, 0x69, 0x53, 0x74, 0x75, 0x64, 0x79, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a,
	0x16, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x47, 0x72, 0x61, 0x64, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59,
	0x65, 0x61, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x49, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a,
	0x0a, 0x11, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x42, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x47, 0x0a, 0x20, 0x66, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x61, 0x64, 0x65,
	0x6d, 0x69, 0x63, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x1d, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x61,
	0x6c, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6f, 0x63,
	0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x66, 0x61, 0x74, 0x68, 0x65, 0x72, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x49, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x6d,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x42, 0x69,
	0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x47, 0x0a, 0x20, 0x6d, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63,
	0x5f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x1d, 0x6d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x63,
	0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x12, 0x2b, 0x0a, 0x11, 0x6d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x70,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6d, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x22,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x5f, 0x69,
	0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x49, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x2e, 0x0a, 0x13, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x5f, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x24, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x67,
	0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x42, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x4b, 0x0a, 0x22, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x5f, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x61, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x5f, 0x62, 0x61, 0x63, 0x6b,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1f, 0x67, 0x75,
	0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x61, 0x64, 0x65,
	0x6d, 0x69, 0x63, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x2f, 0x0a,
	0x13, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x67, 0x75, 0x61, 0x72,
	0x64, 0x69, 0x61, 0x6e, 0x4f, 0x63, 0x63, 0x75, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b,
	0x0a, 0x11, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6d, 0x62, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e,
	0x42, 0x75, 0x6c, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x67, 0x0a, 0x16, 0x42,
	0x75, 0x6c, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x69, 0x6d,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e,
	0x69, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x74, 0x0a, 0x12, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6d, 0x62, 0x5f, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x70, 0x6d, 0x62, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x75,
	0x6c, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x64, 0x0a, 0x11, 0x50, 0x6d,
	0x62, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12,
	0x4f, 0x0a, 0x0a, 0x42, 0x75, 0x6c, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e,
	0x70, 0x6d, 0x62, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x75, 0x6c, 0x6b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x70, 0x6d, 0x62, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x75, 0x6c, 0x6b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x16, 0x5a, 0x14, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x6d, 0x62,
	0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pmb_student_proto_rawDescOnce sync.Once
	file_pmb_student_proto_rawDescData = file_pmb_student_proto_rawDesc
)

func file_pmb_student_proto_rawDescGZIP() []byte {
	file_pmb_student_proto_rawDescOnce.Do(func() {
		file_pmb_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_pmb_student_proto_rawDescData)
	})
	return file_pmb_student_proto_rawDescData
}

var file_pmb_student_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pmb_student_proto_goTypes = []interface{}{
	(*Meta)(nil),                   // 0: pmb_student.Meta
	(*Pagination)(nil),             // 1: pmb_student.Pagination
	(*BulkCreateRequestData)(nil),  // 2: pmb_student.BulkCreateRequestData
	(*BulkCreateRequest)(nil),      // 3: pmb_student.BulkCreateRequest
	(*BulkCreateResponseData)(nil), // 4: pmb_student.BulkCreateResponseData
	(*BulkCreateResponse)(nil),     // 5: pmb_student.BulkCreateResponse
}
var file_pmb_student_proto_depIdxs = []int32{
	2, // 0: pmb_student.BulkCreateRequest.data:type_name -> pmb_student.BulkCreateRequestData
	0, // 1: pmb_student.BulkCreateResponse.meta:type_name -> pmb_student.Meta
	4, // 2: pmb_student.BulkCreateResponse.data:type_name -> pmb_student.BulkCreateResponseData
	3, // 3: pmb_student.PmbStudentHandler.BulkCreate:input_type -> pmb_student.BulkCreateRequest
	5, // 4: pmb_student.PmbStudentHandler.BulkCreate:output_type -> pmb_student.BulkCreateResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pmb_student_proto_init() }
func file_pmb_student_proto_init() {
	if File_pmb_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pmb_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pmb_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_pmb_student_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkCreateRequestData); i {
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
		file_pmb_student_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkCreateRequest); i {
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
		file_pmb_student_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkCreateResponseData); i {
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
		file_pmb_student_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkCreateResponse); i {
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
			RawDescriptor: file_pmb_student_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pmb_student_proto_goTypes,
		DependencyIndexes: file_pmb_student_proto_depIdxs,
		MessageInfos:      file_pmb_student_proto_msgTypes,
	}.Build()
	File_pmb_student_proto = out.File
	file_pmb_student_proto_rawDesc = nil
	file_pmb_student_proto_goTypes = nil
	file_pmb_student_proto_depIdxs = nil
}
