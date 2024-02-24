package class_discussion

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

type classDiscussionHandler struct {
	*service.ServiceCtx
}

func (f classDiscussionHandler) GetList(w http.ResponseWriter, r *http.Request) {
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
	defer f.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Discussion", "Get List", nil)

	paginationData := common.PaginationRequest{
		Limit:  in.GetLimit(),
		Page:   in.GetPage(),
		Search: in.GetSearch(),
	}
	data, errs := f.ClassDiscussionService.GetList(ctx, paginationData, in.GetClassId())
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
			Id:                 v.Id,
			Title:              v.Title,
			Abstraction:        v.Abstraction,
			LecturerId:         v.LecturerId,
			LecturerName:       v.LecturerName,
			LecturerFrontTitle: utils.NullStringScan(v.LecturerFrontTitle),
			LecturerBackDegree: utils.NullStringScan(v.LecturerBackDegree),
			TotalComment:       v.TotalComment,
			LastComment:        utils.NullStringScan(v.LastComment),
		})
	}

	result = GetListResponse{
		Meta: &Meta{
			Message: "Get List Class Discussion",
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

func (f classDiscussionHandler) GetComment(w http.ResponseWriter, r *http.Request) {
	var result GetCommentResponse

	ctx := r.Context()
	var decoder = schema.NewDecoder()
	var in GetCommentRequest
	err := decoder.Decode(&in, r.URL.Query())
	if err != nil {
		logrus.Errorln(err)
		result = GetCommentResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer f.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Discussion", "Get Comment", nil)

	paginationData := common.PaginationRequest{
		Limit: in.GetLimit(),
		Page:  in.GetPage(),
	}
	data, errs := f.ClassDiscussionService.GetComment(ctx, paginationData, in.GetClassDiscussionId())
	if errs != nil {
		utils.PrintError(*errs)
		result = GetCommentResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	resultData := []*GetCommentResponseData{}
	for _, v := range data.Data {
		resultData = append(resultData, &GetCommentResponseData{
			Id:                 v.Id,
			StudentId:          utils.NullStringScan(v.StudentId),
			StudentNimNumber:   utils.NullInt64Scan(v.StudentNimNumber),
			StudentName:        utils.NullStringScan(v.StudentName),
			LecturerId:         utils.NullStringScan(v.LecturerId),
			LecturerName:       utils.NullStringScan(v.LecturerName),
			LecturerFrontTitle: utils.NullStringScan(v.LecturerFrontTitle),
			LecturerBackDegree: utils.NullStringScan(v.LecturerBackDegree),
			Comment:            v.Comment,
			SelfComment:        v.SelfComment,
		})
	}

	result = GetCommentResponse{
		Meta: &Meta{
			Message: "Get Comment Class Discussion",
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

func (c classDiscussionHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var result CreateCommentResponse

	ctx := r.Context()
	var in CreateCommentRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = CreateCommentResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Discussion", "Create Comment", nil)

	data := objects.CreateClassDiscussionComment{
		ClassDiscussionId: in.GetClassDiscussionId(),
		Comment:           in.GetComment(),
	}
	errs := c.ClassDiscussionService.CreateComment(ctx, data)
	if errs != nil {
		utils.PrintError(*errs)
		result = CreateCommentResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = CreateCommentResponse{
		Meta: &Meta{
			Message: "Create Class Discussion Comment",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (c classDiscussionHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	var result DeleteCommentResponse

	ctx := r.Context()
	var in DeleteCommentRequest
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		logrus.Errorln(err)
		result = DeleteCommentResponse{
			Meta: &Meta{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
				Code:    constants.DefaultCustomErrorCode,
			},
		}
		utils.JSONResponse(w, http.StatusInternalServerError, &result)
		return
	}
	defer c.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Class Discussion", "Delete Comment", nil)

	errs := c.ClassDiscussionService.DeleteComment(ctx, in.GetId())
	if errs != nil {
		utils.PrintError(*errs)
		result = DeleteCommentResponse{
			Meta: &Meta{
				Message: errs.Err.Error(),
				Status:  uint32(errs.HttpCode),
				Code:    errs.CustomCode,
			},
		}
		utils.JSONResponse(w, errs.HttpCode, &result)
		return
	}

	result = DeleteCommentResponse{
		Meta: &Meta{
			Message: "Delete Comment Class Discussion",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
