package grade_component

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

type gradeComponentHandler struct {
	*service.ServiceCtx
}

func (f gradeComponentHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Grade Component", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := f.GradeComponentService.GetList(ctx, paginationData, in.GetStudyProgramId(), in.GetSubjectCategoryId())
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
		resultData = append(resultData, &GetListResponseData{
			Id:                  v.Id,
			StudyProgramId:      v.StudyProgramId,
			StudyProgramName:    v.StudyProgramName,
			SubjectCategoryId:   v.SubjectCategoryId,
			SubjectCategoryName: v.SubjectCategoryName,
			Name:                v.Name,
			IsActive:            v.IsActive,
			DefaultPercentage:   v.DefaultPercentage,
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Grade Component",
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

func (f gradeComponentHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Grade Component", "Create", &in)

	data := objects.CreateGradeComponent{
		StudyProgramId:    in.GetStudyProgramId(),
		SubjectCategoryId: in.GetSubjectCategoryId(),
		Name:              in.GetName(),
		IsActive:          in.GetIsActive(),
	}
	errs := f.GradeComponentService.Create(ctx, data)
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
			Message: "Create Grade Component",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (f gradeComponentHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer f.AdminActivityLogService.Create(ctx, r, time.Now(), "Grade Component", "Update", &in)

	data := objects.UpdateGradeComponent{
		Id:                in.GetId(),
		SubjectCategoryId: in.GetSubjectCategoryId(),
		Name:              in.GetName(),
		IsActive:          in.GetIsActive(),
	}
	errs := f.GradeComponentService.Update(ctx, data)
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
			Message: "Update Grade Component",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (a gradeComponentHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
	defer a.AdminActivityLogService.Create(ctx, r, time.Now(), "Grade Component", "Delete", nil)

	errs := a.GradeComponentService.Delete(ctx, in.GetId())
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
			Message: "Delete Grade Component",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (g gradeComponentHandler) GetListBySubjectCategory(w http.ResponseWriter, r *http.Request) {
	var result GetListBySubjectCategoryResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetListBySubjectCategoryRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetListBySubjectCategoryResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer g.AdminActivityLogService.Create(ctx, r, time.Now(), "Grade Component", "Get List By Subject Category", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := g.GradeComponentService.GetListBySubjectCategory(ctx, paginationData, in.GetStudyProgramId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetListBySubjectCategoryResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetListBySubjectCategoryResponseData{}
	for _, v := range data.Data {
		gradeComponents := []*GetListBySubjectCategoryResponseDataGradeComponent{}

		for _, w := range v.GradeComponents {
			gradeComponents = append(gradeComponents, &GetListBySubjectCategoryResponseDataGradeComponent{
				Id:                w.Id,
				Name:              w.Name,
				DefaultPercentage: w.DefaultPercentage,
				IsActive:          w.IsActive,
			})
		}

		resultData = append(resultData, &GetListBySubjectCategoryResponseData{
			StudyProgramId:      v.StudyProgramId,
			StudyProgramName:    v.StudyProgramName,
			SubjectCategoryId:   v.SubjectCategoryId,
			SubjectCategoryName: v.SubjectCategoryName,
			GradeComponents:     gradeComponents,
		})
	}

	result = GetListBySubjectCategoryResponse{
		Meta: &Meta{
			Message: "Get List Grade Component By Subject Category",
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

func (g gradeComponentHandler) BulkUpdatePercentage(w http.ResponseWriter, r *http.Request) {
	var result BulkUpdatePercentageResponse

	ctx := r.Context()
	var in BulkUpdatePercentageRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = BulkUpdatePercentageResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer g.AdminActivityLogService.Create(ctx, r, time.Now(), "Grade Component", "Bulk Update Percentage", &in)

	baseData := objects.BulkUpdatePercentageGradeComponent{
		StudyProgramId:    in.GetStudyProgramId(),
		SubjectCategoryId: in.GetSubjectCategoryId(),
	}
	data := []objects.BulkUpdatePercentageGradeComponentData{}
	for _, v := range in.GetGradeComponents() {
		data = append(data, objects.BulkUpdatePercentageGradeComponentData{
			Id:                v.GetId(),
			DefaultPercentage: v.GetDefaultPercentage(),
			IsActive:          v.GetIsActive(),
		})
	}
	errs := g.GradeComponentService.BulkUpdatePercentage(ctx, baseData, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = BulkUpdatePercentageResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = BulkUpdatePercentageResponse{
		Meta: &Meta{
			Message: "Bulk Update Percentage Grade Component",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
