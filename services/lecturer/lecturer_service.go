package lecturer

import (
	"context"
	"math"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	appUtils "github.com/sccicitb/pupr-backend/utils"
)

type lecturerService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (p lecturerService) GetList(ctx context.Context, paginationData common.PaginationRequest, req objects.GetLecturerRequest) (objects.LecturerListWithPagination, *constants.ErrorResponse) {
	var result objects.LecturerListWithPagination

	tx, err := p.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := p.LecturerRepo.GetList(ctx, tx, paginationData, req)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLecturer{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLecturer{
			Id:                              v.Id,
			Name:                            v.Name,
			PhoneNumber:                     v.PhoneNumber,
			MobilePhoneNumber:               v.MobilePhoneNumber,
			OfficePhoneNumber:               v.OfficePhoneNumber,
			IdNationalLecturer:              v.IdNationalLecturer,
			FrontTitle:                      v.FrontTitle,
			BackDegree:                      v.BackDegree,
			DiktiStudyProgramCode:           v.DiktiStudyProgramCode,
			StudyProgramName:                v.StudyProgramName,
			EmploymentStatus:                v.EmploymentStatus,
			Status:                          v.Status,
			AuthenticationId:                v.AuthenticationId,
			AuthenticationIsActive:          v.AuthenticationIsActive,
			AuthenticationSuspensionRemarks: v.AuthenticationSuspensionRemarks,
			AcademicGuidanceId:              v.AcademicGuidanceId,
			AcademicGuidanceTotalStudent:    v.AcademicGuidanceTotalStudent,
			AcademicGuidanceDecisionNumber:  v.AcademicGuidanceDecisionNumber,
			AcademicGuidanceDecisionDate:    v.AcademicGuidanceDecisionDate,
		})
	}

	result = objects.LecturerListWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (p lecturerService) GetSchedule(ctx context.Context, paginationData common.PaginationRequest, studyProgramId, idNationalLecturer, semesterId string) (objects.LecturerScheduleWithPagination, *constants.ErrorResponse) {
	var result objects.LecturerScheduleWithPagination

	tx, err := p.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := p.LecturerRepo.GetSchedule(ctx, tx, paginationData, studyProgramId, idNationalLecturer, semesterId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLecturerSchedule{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLecturerSchedule{
			Id:                 v.Id,
			IdNationalLecturer: v.IdNationalLecturer,
			Name:               v.Name,
			FrontTitle:         v.FrontTitle,
			BackDegree:         v.BackDegree,
			StudyProgramName:   v.StudyProgramName,
			SubjectName:        v.SubjectName,
			ClassName:          v.ClassName,
			TotalSubjectCredit: v.TotalSubjectCredit,
			LecturePlanDate:    v.LecturePlanDate,
			StartTime:          v.StartTime,
			EndTime:            v.EndTime,
			RoomName:           v.RoomName,
			TotalParticipant:   v.TotalParticipant,
		})
	}

	result = objects.LecturerScheduleWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (f lecturerService) GetDetail(ctx context.Context, id string) (objects.GetLecturerDetail, *constants.ErrorResponse) {
	var result objects.GetLecturerDetail

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	resultData, errs := f.LecturerRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GetLecturerDetail{
		Id:                        resultData.Id,
		IdNationalLecturer:        resultData.IdNationalLecturer,
		Name:                      resultData.Name,
		FrontTitle:                resultData.FrontTitle,
		BackDegree:                resultData.BackDegree,
		StudyProgramId:            resultData.StudyProgramId,
		StudyProgramName:          resultData.StudyProgramName,
		IdNumber:                  resultData.IdNumber,
		BirthDate:                 resultData.BirthDate,
		BirthRegencyId:            resultData.BirthRegencyId,
		BirthRegencyName:          resultData.BirthRegencyName,
		BirthCountryId:            resultData.BirthCountryId,
		BirthCountryName:          resultData.BirthCountryName,
		IdEmployee:                resultData.IdEmployee,
		Stambuk:                   resultData.Stambuk,
		Sex:                       resultData.Sex,
		BloodType:                 resultData.BloodType,
		Religion:                  resultData.Religion,
		MaritalStatus:             resultData.MaritalStatus,
		Address:                   resultData.Address,
		RegencyId:                 resultData.RegencyId,
		RegencyName:               resultData.RegencyName,
		CountryId:                 resultData.CountryId,
		CountryName:               resultData.CountryName,
		PostalCode:                resultData.PostalCode,
		PhoneNumber:               resultData.PhoneNumber,
		Fax:                       resultData.Fax,
		MobilePhoneNumber:         resultData.MobilePhoneNumber,
		OfficePhoneNumber:         resultData.OfficePhoneNumber,
		EmployeeType:              resultData.EmployeeType,
		EmployeeStatus:            resultData.EmployeeStatus,
		SkCpnsNumber:              resultData.SkCpnsNumber,
		SkCpnsDate:                resultData.SkCpnsDate,
		TmtCpnsDate:               resultData.TmtCpnsDate,
		CpnsCategory:              resultData.CpnsCategory,
		CpnsDurationMonth:         resultData.CpnsDurationMonth,
		PrePositionDate:           resultData.PrePositionDate,
		SkPnsNumber:               resultData.SkPnsNumber,
		SkPnsDate:                 resultData.SkPnsDate,
		TmtPnsDate:                resultData.TmtPnsDate,
		PnsCategory:               resultData.PnsCategory,
		PnsOathDate:               resultData.PnsOathDate,
		JoinDate:                  resultData.JoinDate,
		EndDate:                   resultData.EndDate,
		TaspenNumber:              resultData.TaspenNumber,
		FormerInstance:            resultData.FormerInstance,
		Remarks:                   resultData.Remarks,
		LecturerNumber:            resultData.LecturerNumber,
		AcademicPosition:          resultData.AcademicPosition,
		EmploymentStatus:          resultData.EmploymentStatus,
		Expertise:                 resultData.Expertise,
		HighestDegree:             resultData.HighestDegree,
		InstanceCode:              resultData.InstanceCode,
		TeachingCertificateNumber: resultData.TeachingCertificateNumber,
		TeachingPermitNumber:      resultData.TeachingPermitNumber,
		Status:                    resultData.Status,
		ResignSemester:            resultData.ResignSemester,
		ExpertiseGroupId:          resultData.ExpertiseGroupId,
		ExpertiseGroupName:        resultData.ExpertiseGroupName,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a lecturerService) Create(ctx context.Context, data objects.CreateLecturer) *constants.ErrorResponse {
	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	createData := models.CreateLecturer{
		IdNationalLecturer:        data.IdNationalLecturer,
		Name:                      data.Name,
		FrontTitle:                utils.NewNullString(data.FrontTitle),
		BackDegree:                utils.NewNullString(data.BackDegree),
		StudyProgramId:            utils.NewNullString(data.StudyProgramId),
		IdNumber:                  utils.NewNullString(data.IdNumber),
		BirthDate:                 utils.NewNullTime(data.BirthDate),
		BirthRegencyId:            utils.NewNullInt32(int32(data.BirthRegencyId)),
		IdEmployee:                utils.NewNullString(data.IdEmployee),
		Stambuk:                   utils.NewNullString(data.Stambuk),
		Sex:                       utils.NewNullString(data.Sex),
		BloodType:                 utils.NewNullString(data.BloodType),
		Religion:                  utils.NewNullString(data.Religion),
		MaritalStatus:             utils.NewNullString(data.MaritalStatus),
		Address:                   utils.NewNullString(data.Address),
		RegencyId:                 utils.NewNullInt32(int32(data.RegencyId)),
		PostalCode:                utils.NewNullString(data.PostalCode),
		PhoneNumber:               utils.NewNullString(data.PhoneNumber),
		Fax:                       utils.NewNullString(data.Fax),
		MobilePhoneNumber:         utils.NewNullString(data.MobilePhoneNumber),
		OfficePhoneNumber:         utils.NewNullString(data.OfficePhoneNumber),
		EmployeeType:              utils.NewNullString(data.EmployeeType),
		EmployeeStatus:            utils.NewNullString(data.EmployeeStatus),
		SkCpnsNumber:              utils.NewNullString(data.SkCpnsNumber),
		SkCpnsDate:                utils.NewNullTime(data.SkCpnsDate),
		TmtCpnsDate:               utils.NewNullTime(data.TmtCpnsDate),
		CpnsCategory:              utils.NewNullString(data.CpnsCategory),
		CpnsDurationMonth:         utils.NewNullInt32(int32(data.CpnsDurationMonth)),
		PrePositionDate:           utils.NewNullTime(data.PrePositionDate),
		SkPnsNumber:               utils.NewNullString(data.SkPnsNumber),
		SkPnsDate:                 utils.NewNullTime(data.SkPnsDate),
		TmtPnsDate:                utils.NewNullTime(data.TmtPnsDate),
		PnsCategory:               utils.NewNullString(data.PnsCategory),
		PnsOathDate:               utils.NewNullTime(data.PnsOathDate),
		JoinDate:                  utils.NewNullTime(data.JoinDate),
		EndDate:                   utils.NewNullTime(data.EndDate),
		TaspenNumber:              utils.NewNullString(data.TaspenNumber),
		FormerInstance:            utils.NewNullString(data.FormerInstance),
		Remarks:                   utils.NewNullString(data.Remarks),
		LecturerNumber:            utils.NewNullString(data.LecturerNumber),
		AcademicPosition:          utils.NewNullString(data.AcademicPosition),
		EmploymentStatus:          utils.NewNullString(data.EmploymentStatus),
		Expertise:                 utils.NewNullString(data.Expertise),
		HighestDegree:             utils.NewNullString(data.HighestDegree),
		InstanceCode:              utils.NewNullString(data.InstanceCode),
		TeachingCertificateNumber: utils.NewNullString(data.TeachingCertificateNumber),
		TeachingPermitNumber:      utils.NewNullString(data.TeachingPermitNumber),
		Status:                    utils.NewNullString(data.Status),
		ResignSemester:            utils.NewNullString(data.ResignSemester),
		ExpertiseGroupId:          utils.NewNullString(data.ExpertiseGroupId),
		CreatedBy:                 claims.ID,
	}
	errs = a.LecturerRepo.Create(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a lecturerService) Update(ctx context.Context, data objects.UpdateLecturer) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.LecturerRepo.GetDetail(ctx, tx, data.Id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	updateData := models.UpdateLecturer{
		Id:                        data.Id,
		IdNationalLecturer:        data.IdNationalLecturer,
		Name:                      data.Name,
		FrontTitle:                utils.NewNullString(data.FrontTitle),
		BackDegree:                utils.NewNullString(data.BackDegree),
		StudyProgramId:            utils.NewNullString(data.StudyProgramId),
		IdNumber:                  utils.NewNullString(data.IdNumber),
		BirthDate:                 utils.NewNullTime(data.BirthDate),
		BirthRegencyId:            utils.NewNullInt32(int32(data.BirthRegencyId)),
		IdEmployee:                utils.NewNullString(data.IdEmployee),
		Stambuk:                   utils.NewNullString(data.Stambuk),
		Sex:                       utils.NewNullString(data.Sex),
		BloodType:                 utils.NewNullString(data.BloodType),
		Religion:                  utils.NewNullString(data.Religion),
		MaritalStatus:             utils.NewNullString(data.MaritalStatus),
		Address:                   utils.NewNullString(data.Address),
		RegencyId:                 utils.NewNullInt32(int32(data.RegencyId)),
		PostalCode:                utils.NewNullString(data.PostalCode),
		PhoneNumber:               utils.NewNullString(data.PhoneNumber),
		Fax:                       utils.NewNullString(data.Fax),
		MobilePhoneNumber:         utils.NewNullString(data.MobilePhoneNumber),
		OfficePhoneNumber:         utils.NewNullString(data.OfficePhoneNumber),
		EmployeeType:              utils.NewNullString(data.EmployeeType),
		EmployeeStatus:            utils.NewNullString(data.EmployeeStatus),
		SkCpnsNumber:              utils.NewNullString(data.SkCpnsNumber),
		SkCpnsDate:                utils.NewNullTime(data.SkCpnsDate),
		TmtCpnsDate:               utils.NewNullTime(data.TmtCpnsDate),
		CpnsCategory:              utils.NewNullString(data.CpnsCategory),
		CpnsDurationMonth:         utils.NewNullInt32(int32(data.CpnsDurationMonth)),
		PrePositionDate:           utils.NewNullTime(data.PrePositionDate),
		SkPnsNumber:               utils.NewNullString(data.SkPnsNumber),
		SkPnsDate:                 utils.NewNullTime(data.SkPnsDate),
		TmtPnsDate:                utils.NewNullTime(data.TmtPnsDate),
		PnsCategory:               utils.NewNullString(data.PnsCategory),
		PnsOathDate:               utils.NewNullTime(data.PnsOathDate),
		JoinDate:                  utils.NewNullTime(data.JoinDate),
		EndDate:                   utils.NewNullTime(data.EndDate),
		TaspenNumber:              utils.NewNullString(data.TaspenNumber),
		FormerInstance:            utils.NewNullString(data.FormerInstance),
		Remarks:                   utils.NewNullString(data.Remarks),
		LecturerNumber:            utils.NewNullString(data.LecturerNumber),
		AcademicPosition:          utils.NewNullString(data.AcademicPosition),
		EmploymentStatus:          utils.NewNullString(data.EmploymentStatus),
		Expertise:                 utils.NewNullString(data.Expertise),
		HighestDegree:             utils.NewNullString(data.HighestDegree),
		InstanceCode:              utils.NewNullString(data.InstanceCode),
		TeachingCertificateNumber: utils.NewNullString(data.TeachingCertificateNumber),
		TeachingPermitNumber:      utils.NewNullString(data.TeachingPermitNumber),
		Status:                    utils.NewNullString(data.Status),
		ResignSemester:            utils.NewNullString(data.ResignSemester),
		ExpertiseGroupId:          utils.NewNullString(data.ExpertiseGroupId),
		UpdatedBy:                 claims.ID,
	}
	errs = a.LecturerRepo.Update(ctx, tx, updateData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a lecturerService) Delete(ctx context.Context, id string) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	_, errs := a.LecturerRepo.GetDetail(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	errs = a.LecturerRepo.Delete(ctx, tx, id)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (a lecturerService) GetSemesterSummary(ctx context.Context) (objects.GetLecturerSemesterSummary, *constants.ErrorResponse) {
	var result objects.GetLecturerSemesterSummary

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	semesterData, errs := a.SemesterRepo.GetActive(ctx, tx)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	academicGuidanceData, errs := a.AcademicGuidanceRepo.GetDetailBySemesterIdLecturerId(ctx, tx, semesterData.Id, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	classData, errs := a.ClassRepo.GetClassLecturersBySemesterIdLecturerId(ctx, tx, semesterData.Id, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GetLecturerSemesterSummary{
		SemesterId:                   semesterData.Id,
		StudyPlanApprovalStartDate:   semesterData.StudyPlanApprovalStartDate,
		StudyPlanApprovalEndDate:     semesterData.StudyPlanApprovalEndDate,
		AcademicGuidanceTotalStudent: academicGuidanceData.TotalStudent,
		TotalClass:                   uint32(len(classData)),
		SchoolYear:                   appUtils.GenerateSchoolYear(semesterData.SemesterStartYear),
		SemesterType:                 semesterData.SemesterType,
		GradingStartDate:             semesterData.GradingStartDate,
		GradingEndDate:               semesterData.GradingEndDate,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a lecturerService) GetProfile(ctx context.Context) (objects.GetLecturerProfile, *constants.ErrorResponse) {
	var result objects.GetLecturerProfile

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return result, errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	lecturerData, errs := a.LecturerRepo.GetDetail(ctx, tx, claims.ID)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	result = objects.GetLecturerProfile{
		Id:                 lecturerData.Id,
		IdNationalLecturer: lecturerData.IdNationalLecturer,
		Name:               lecturerData.Name,
		FrontTitle:         lecturerData.FrontTitle,
		BackDegree:         lecturerData.BackDegree,
		StudyProgramId:     lecturerData.StudyProgramId,
		StudyProgramName:   lecturerData.StudyProgramName,
		BirthDate:          lecturerData.BirthDate,
		BirthRegencyId:     lecturerData.BirthRegencyId,
		BirthRegencyName:   lecturerData.BirthRegencyName,
		BirthCountryId:     lecturerData.BirthCountryId,
		BirthCountryName:   lecturerData.BirthCountryName,
		Sex:                lecturerData.Sex,
		Religion:           lecturerData.Religion,
		Address:            lecturerData.Address,
		RegencyId:          lecturerData.RegencyId,
		RegencyName:        lecturerData.RegencyName,
		CountryId:          lecturerData.CountryId,
		CountryName:        lecturerData.CountryName,
		PostalCode:         lecturerData.PostalCode,
		PhoneNumber:        lecturerData.PhoneNumber,
		Fax:                lecturerData.Fax,
		MobilePhoneNumber:  lecturerData.MobilePhoneNumber,
		OfficePhoneNumber:  lecturerData.OfficePhoneNumber,
		AcademicPosition:   lecturerData.AcademicPosition,
		EmploymentStatus:   lecturerData.EmploymentStatus,
		Status:             lecturerData.Status,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a lecturerService) GetAssignedClass(ctx context.Context, semesterId, lecturerId string, classIsActive *bool) ([]objects.GetLecturerAssignedClass, *constants.ErrorResponse) {
	results := []objects.GetLecturerAssignedClass{}

	claims, errs := utils.GetJWTClaimsFromContext(ctx)
	if errs != nil {
		return results, errs
	}

	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return results, constants.ErrUnknown
	}

	if lecturerId == "" {
		lecturerId = claims.ID
	}

	resultData, errs := a.LecturerRepo.GetAssignedClass(ctx, tx, lecturerId, semesterId, "", classIsActive)
	if errs != nil {
		_ = tx.Rollback()
		return results, errs
	}

	for _, v := range resultData {
		var attendancePercentage float64
		totalLecture := v.TotalLectureDone
		if totalLecture != 0 {
			attendancePercentage = math.Round((float64(v.TotalAttendance)/float64(totalLecture))*10000) / 100
		}
		results = append(results, objects.GetLecturerAssignedClass{
			Id:                   v.Id,
			Name:                 v.Name,
			SubjectCode:          v.SubjectCode,
			SubjectName:          v.SubjectName,
			TheoryCredit:         v.TheoryCredit,
			PracticumCredit:      v.PracticumCredit,
			FieldPracticumCredit: v.FieldPracticumCredit,
			IsGradingResponsible: v.IsGradingResponsible,
			StudyProgramId:       v.StudyProgramId,
			StudyProgramName:     v.StudyProgramName,
			TotalAttendance:      v.TotalAttendance,
			TotalLectureDone:     v.TotalLectureDone,
			AttendancePercentage: attendancePercentage,
		})
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return results, constants.ErrUnknown
	}

	return results, nil
}
