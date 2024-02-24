package subject

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type subjectHandler struct {
	*service.ServiceCtx
}

func (a subjectHandler) GetList(w http.ResponseWriter, r *http.Request) {
	var result GetListResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetListRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetListResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Subject", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}

	isThesis, errs := utils.StringToBoolPointer(in.GetIsThesis())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetListResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	requestData := objects.GetSubjectRequest{
		CurriculumIds:            in.GetCurriculumId(),
		PrerequisiteOfSubjectId:  in.GetPrerequisiteOfSubjectId(),
		EquivalentToCurriculumId: in.GetEquivalentToCurriculumId(),
		SemesterPackage:          in.GetSemesterPackage(),
		ClassSemesterId:          in.GetClassSemesterId(),
		IsThesis:                 isThesis,
		IsMbkm:                   in.GetIsMbkm(),
	}
	data, errs := a.SubjectService.GetList(ctx, paginationData, requestData)
	if errs != nil {
		utils.PrintError(*errs)
		result = GetListResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetListResponseData{}
	for _, v := range data.Data {
		classes := []*GetListResponseDataClass{}
		for _, w := range v.Classes {
			classes = append(classes, &GetListResponseDataClass{
				Id:   w.Id,
				Name: w.Name,
			})
		}

		resultData = append(resultData, &GetListResponseData{
			Id:                            v.Id,
			StudyProgramId:                v.StudyProgramId,
			StudyProgramName:              v.StudyProgramName,
			CurriculumId:                  v.CurriculumId,
			CurriculumName:                v.CurriculumName,
			Code:                          v.Code,
			Name:                          v.Name,
			IsMandatory:                   v.IsMandatory,
			SemesterPackage:               v.SemesterPackage,
			TheoryCredit:                  v.TheoryCredit,
			PracticumCredit:               v.PracticumCredit,
			FieldPracticumCredit:          v.FieldPracticumCredit,
			SubjectPrerequisiteId:         utils.NullStringScan(v.SubjectPrerequisiteId),
			PrerequisiteType:              utils.NullStringScan(v.PrerequisiteType),
			PrerequisiteMinimumGradePoint: utils.NullFloatScan(v.PrerequisiteMinimumGradePoint),
			EquivalentStudyProgramId:      utils.NullStringScan(v.EquivalentStudyProgramId),
			EquivalentStudyProgramName:    utils.NullStringScan(v.EquivalentStudyProgramName),
			EquivalentCurriculumId:        utils.NullStringScan(v.EquivalentCurriculumId),
			EquivalentCurriculumName:      utils.NullStringScan(v.EquivalentCurriculumName),
			EquivalentSubjectId:           utils.NullStringScan(v.EquivalentSubjectId),
			EquivalentSubjectCode:         utils.NullStringScan(v.EquivalentSubjectCode),
			EquivalentSubjectName:         utils.NullStringScan(v.EquivalentSubjectName),
			SubjectCategoryId:             v.SubjectCategoryId,
			SubjectCategoryName:           v.SubjectCategoryName,
			Classes:                       classes,
			IsMbkm:                        v.IsMbkm,
			TotalLessonPlan:               v.TotalLessonPlan,
			IsThesis:                      utils.NullBooleanScan(v.IsThesis),
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Subject",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Pagination: &Pagination{
			Page:         uint32(data.Pagination.Page),
			Limit:        uint32(data.Pagination.Limit),
			Prev:         uint32(*data.Pagination.Prev),
			Next:         uint32(*data.Pagination.Next),
			TotalPages:   uint32(*data.Pagination.TotalPages),
			TotalRecords: uint32(*data.Pagination.TotalRecords),
		},
		Data: resultData,
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a subjectHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	var result GetDetailResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetDetailRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetDetailResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Subject", "Get Detail", nil)

	data, errs := a.SubjectService.GetDetail(ctx, in.GetId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetDetailResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	prerequisiteSubjects := []*GetDetailResponseDataPrerequisiteSubject{}
	for _, v := range data.SubjectPrerequisites {
		prerequisiteSubjects = append(prerequisiteSubjects, &GetDetailResponseDataPrerequisiteSubject{
			Id:                v.Id,
			Code:              v.Code,
			Name:              v.Name,
			PrerequisiteType:  v.PrerequisiteType,
			MinimumGradePoint: utils.NullFloatScan(v.MinimumGradePoint),
		})
	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Subject",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                           data.Id,
			StudyProgramId:               data.StudyProgramId,
			StudyProgramName:             data.StudyProgramName,
			CurriculumId:                 data.CurriculumId,
			CurriculumName:               data.CurriculumName,
			Code:                         data.Code,
			Name:                         data.Name,
			ShortName:                    utils.NullStringScan(data.ShortName),
			EnglishName:                  utils.NullStringScan(data.EnglishName),
			EnglishShortName:             utils.NullStringScan(data.EnglishShortName),
			IsMandatory:                  data.IsMandatory,
			Trait:                        data.Trait,
			Type:                         utils.NullStringScan(data.Type),
			SubjectCategoryId:            data.SubjectCategoryId,
			SubjectCategoryName:          data.SubjectCategoryName,
			CurriculumType:               data.CurriculumType,
			TheoryCredit:                 data.TheoryCredit,
			PracticumCredit:              data.PracticumCredit,
			FieldPracticumCredit:         data.FieldPracticumCredit,
			SemesterPackage:              data.SemesterPackage,
			RepeatCourseLimit:            data.RepeatCourseLimit,
			IsActive:                     data.IsActive,
			HasLectureUnit:               data.HasLectureUnit,
			HasTeachingMaterial:          data.HasTeachingMaterial,
			HasLectureSummary:            data.HasLectureSummary,
			SupportingLecturerId:         utils.NullStringScan(data.SupportingLecturerId),
			SupportingLecturerName:       utils.NullStringScan(data.SupportingLecturerName),
			StartDate:                    utils.SafetyDate(data.StartDate),
			EndDate:                      utils.SafetyDate(data.EndDate),
			MinimumPassingGradePoint:     data.MinimumPassingGradePoint,
			MinimumMandatoryCreditTaken:  utils.NullUint32Scan(data.MinimumMandatoryCreditTaken),
			MinimumOptionalCreditTaken:   utils.NullUint32Scan(data.MinimumOptionalCreditTaken),
			MinimumTotalCreditTaken:      utils.NullUint32Scan(data.MinimumTotalCreditTaken),
			MinimumMandatoryCreditPassed: utils.NullUint32Scan(data.MinimumMandatoryCreditPassed),
			MinimumOptionalCreditPassed:  utils.NullUint32Scan(data.MinimumOptionalCreditPassed),
			MinimumTotalCreditPassed:     utils.NullUint32Scan(data.MinimumTotalCreditPassed),
			MinimumGpa:                   utils.NullFloatScan(data.MinimumGpa),
			Abstraction:                  utils.NullStringScan(data.Abstraction),
			SyllabusPath:                 utils.NullStringScan(data.SyllabusPath),
			SyllabusPathType:             utils.NullStringScan(data.SyllabusPathType),
			SyllabusUrl:                  data.SyllabusUrl,
			PrerequisiteSubjects:         prerequisiteSubjects,
			IsThesis:                     utils.NullBooleanScan(data.IsThesis),
			IsMbkm:                       data.IsMbkm,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a subjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var result CreateResponse

	ctx := r.Context()
	var in CreateRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = CreateResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Subject", "Create", &in)

	startDate, errs := utils.StringToTime(in.GetStartDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = CreateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	endDate, errs := utils.StringToTime(in.GetEndDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = CreateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	data := objects.CreateSubject{
		CurriculumId:                 in.GetCurriculumId(),
		Code:                         in.GetCode(),
		Name:                         in.GetName(),
		ShortName:                    in.GetShortName(),
		EnglishName:                  in.GetEnglishName(),
		EnglishShortName:             in.GetEnglishShortName(),
		IsMandatory:                  in.GetIsMandatory(),
		Trait:                        in.GetTrait(),
		Type:                         in.GetType(),
		SubjectCategoryId:            in.GetSubjectCategoryId(),
		CurriculumType:               in.GetCurriculumType(),
		TheoryCredit:                 in.GetTheoryCredit(),
		PracticumCredit:              in.GetPracticumCredit(),
		FieldPracticumCredit:         in.GetFieldPracticumCredit(),
		SemesterPackage:              in.GetSemesterPackage(),
		RepeatCourseLimit:            in.GetRepeatCourseLimit(),
		IsActive:                     in.GetIsActive(),
		HasLectureUnit:               in.GetHasLectureUnit(),
		HasTeachingMaterial:          in.GetHasTeachingMaterial(),
		HasLectureSummary:            in.GetHasLectureSummary(),
		SupportingLecturerId:         in.GetSupportingLecturerId(),
		StartDate:                    startDate,
		EndDate:                      endDate,
		MinimumPassingGradePoint:     in.GetMinimumPassingGradePoint(),
		MinimumMandatoryCreditTaken:  in.GetMinimumMandatoryCreditTaken(),
		MinimumOptionalCreditTaken:   in.GetMinimumOptionalCreditTaken(),
		MinimumTotalCreditTaken:      in.GetMinimumTotalCreditTaken(),
		MinimumMandatoryCreditPassed: in.GetMinimumMandatoryCreditPassed(),
		MinimumOptionalCreditPassed:  in.GetMinimumOptionalCreditPassed(),
		MinimumTotalCreditPassed:     in.GetMinimumTotalCreditPassed(),
		MinimumGpa:                   in.GetMinimumGpa(),
		Abstraction:                  in.GetAbstraction(),
		SyllabusPath:                 in.GetSyllabusPath(),
		SyllabusPathType:             in.GetSyllabusPathType(),
		IsThesis:                     in.GetIsThesis(),
		IsMbkm:                       in.GetIsMbkm(),
	}
	errs = a.SubjectService.Create(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = CreateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = CreateResponse{
		Meta: &Meta{
			Message: "Create Subject",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (a subjectHandler) Update(w http.ResponseWriter, r *http.Request) {
	var result UpdateResponse

	ctx := r.Context()
	var in UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = UpdateResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Subject", "Update", &in)

	startDate, errs := utils.StringToTime(in.GetStartDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	endDate, errs := utils.StringToTime(in.GetEndDate())
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}
	data := objects.UpdateSubject{
		Id:                           in.GetId(),
		CurriculumId:                 in.GetCurriculumId(),
		Code:                         in.GetCode(),
		Name:                         in.GetName(),
		ShortName:                    in.GetShortName(),
		EnglishName:                  in.GetEnglishName(),
		EnglishShortName:             in.GetEnglishShortName(),
		IsMandatory:                  in.GetIsMandatory(),
		Trait:                        in.GetTrait(),
		Type:                         in.GetType(),
		SubjectCategoryId:            in.GetSubjectCategoryId(),
		CurriculumType:               in.GetCurriculumType(),
		TheoryCredit:                 in.GetTheoryCredit(),
		PracticumCredit:              in.GetPracticumCredit(),
		FieldPracticumCredit:         in.GetFieldPracticumCredit(),
		SemesterPackage:              in.GetSemesterPackage(),
		RepeatCourseLimit:            in.GetRepeatCourseLimit(),
		IsActive:                     in.GetIsActive(),
		HasLectureUnit:               in.GetHasLectureUnit(),
		HasTeachingMaterial:          in.GetHasTeachingMaterial(),
		HasLectureSummary:            in.GetHasLectureSummary(),
		SupportingLecturerId:         in.GetSupportingLecturerId(),
		StartDate:                    startDate,
		EndDate:                      endDate,
		MinimumPassingGradePoint:     in.GetMinimumPassingGradePoint(),
		MinimumMandatoryCreditTaken:  in.GetMinimumMandatoryCreditTaken(),
		MinimumOptionalCreditTaken:   in.GetMinimumOptionalCreditTaken(),
		MinimumTotalCreditTaken:      in.GetMinimumTotalCreditTaken(),
		MinimumMandatoryCreditPassed: in.GetMinimumMandatoryCreditPassed(),
		MinimumOptionalCreditPassed:  in.GetMinimumOptionalCreditPassed(),
		MinimumTotalCreditPassed:     in.GetMinimumTotalCreditPassed(),
		MinimumGpa:                   in.GetMinimumGpa(),
		Abstraction:                  in.GetAbstraction(),
		SyllabusPath:                 in.GetSyllabusPath(),
		SyllabusPathType:             in.GetSyllabusPathType(),
		IsThesis:                     in.GetIsThesis(),
		IsMbkm:                       in.GetIsMbkm(),
	}
	errs = a.SubjectService.Update(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = UpdateResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = UpdateResponse{
		Meta: &Meta{
			Message: "Update Subject",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a subjectHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var result DeleteResponse

	ctx := r.Context()
	var in DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = DeleteResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Subject", "Delete", nil)

	errs := a.SubjectService.Delete(ctx, in.GetId())
	if errs != nil {
		utils.PrintError(*errs)
		result = DeleteResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = DeleteResponse{
		Meta: &Meta{
			Message: "Delete Subject",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a subjectHandler) SetPrerequisiteSubject(w http.ResponseWriter, r *http.Request) {
	var result SetPrerequisiteSubjectResponse

	ctx := r.Context()
	var in SetPrerequisiteSubjectRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = SetPrerequisiteSubjectResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Subject", "Set Prerequisite Subject", nil)

	prerequisiteData := []objects.SetPrerequisiteSubject{}
	for _, v := range in.GetPrerequisites() {
		prerequisiteData = append(prerequisiteData, objects.SetPrerequisiteSubject{
			Id:                v.GetId(),
			PrerequisiteType:  v.GetPrerequisiteType(),
			MinimumGradePoint: v.GetMinimumGradePoint(),
		})
	}
	errs := a.SubjectService.SetPrerequisiteSubject(ctx, in.GetId(), prerequisiteData)
	if errs != nil {
		utils.PrintError(*errs)
		result = SetPrerequisiteSubjectResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = SetPrerequisiteSubjectResponse{
		Meta: &Meta{
			Message: "Set Prerequisite Subject",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a subjectHandler) SetEquivalentSubject(w http.ResponseWriter, r *http.Request) {
	var result SetEquivalentSubjectResponse

	ctx := r.Context()
	var in SetEquivalentSubjectRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = SetEquivalentSubjectResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Subject", "Set Equivalent Subject", nil)

	data := objects.SetEquivalentSubject{
		SubjectId:           in.GetSubjectId(),
		EquivalentSubjectId: in.GetEquivalentSubjectId(),
		IsViceVersa:         in.GetIsViceVersa(),
	}
	errs := a.SubjectService.SetEquivalentSubject(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = SetEquivalentSubjectResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = SetEquivalentSubjectResponse{
		Meta: &Meta{
			Message: "Set Equivalent Subject",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a subjectHandler) DeleteEquivalentSubject(w http.ResponseWriter, r *http.Request) {
	var result DeleteEquivalentSubjectResponse

	ctx := r.Context()
	var in DeleteEquivalentSubjectRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = DeleteEquivalentSubjectResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Subject", "Delete Equivalent Subject", nil)

	errs := a.SubjectService.DeleteEquivalentSubject(ctx, in.GetSubjectId(), in.GetEquivalentSubjectId())
	if errs != nil {
		utils.PrintError(*errs)
		result = DeleteEquivalentSubjectResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = DeleteEquivalentSubjectResponse{
		Meta: &Meta{
			Message: "Delete Equivalent Subject",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
