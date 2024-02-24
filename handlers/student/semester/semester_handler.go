package semester

import (
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects/common"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type semesterHandler struct {
	*service.ServiceCtx
}

func (a semesterHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer a.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Semester", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit: in.GetLimit(),
		Page:  in.GetPage(),
	}
	data, errs := a.SemesterService.GetList(ctx, paginationData, in.GetStudyProgramId(), "")
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
			Id:                v.Id,
			SemesterStartYear: v.SemesterStartYear,
			SchoolYear:        v.SchoolYear,
			SemesterType:      v.SemesterType,
			IsActive:          v.IsActive,
			StartDate:         v.StartDate.Format(constants.DateRFC),
			EndDate:           v.EndDate.Format(constants.DateRFC),
			MidtermStartDate:  utils.SafetyDate(v.MidtermStartDate),
			MidtermEndDate:    utils.SafetyDate(v.MidtermEndDate),
			EndtermStartDate:  utils.SafetyDate(v.EndtermStartDate),
			EndtermEndDate:    utils.SafetyDate(v.EndtermEndDate),
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Semester",
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
