package thesis

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type thesisHandler struct {
	*service.ServiceCtx
}

func (t thesisHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer t.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Get Detail", nil)

	data, errs := t.ThesisService.GetDetail(ctx, "")
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

	var files []*GetDetailResponseDataFile
	for _, v := range data.Files {
		files = append(files, &GetDetailResponseDataFile{
			Id:              v.Id,
			FileUrl:         v.FileUrl,
			FilePath:        v.FilePath,
			FilePathType:    v.FilePathType,
			FileDescription: utils.NullStringScan(v.FileDescription),
		})
	}

	var thesisSupervisors []*GetDetailResponseDataThesisSupervisor
	for _, v := range data.ThesisSupervisors {
		thesisSupervisors = append(thesisSupervisors, &GetDetailResponseDataThesisSupervisor{
			Id:                       v.Id,
			LecturerId:               v.LecturerId,
			LecturerName:             v.LecturerName,
			LecturerFrontTitle:       utils.NullStringScan(v.LecturerFrontTitle),
			LecturerBackDegree:       utils.NullStringScan(v.LecturerBackDegree),
			ThesisSupervisorRoleId:   v.ThesisSupervisorRoleId,
			ThesisSupervisorRoleName: v.ThesisSupervisorRoleName,
			ThesisSupervisorRoleSort: v.ThesisSupervisorRoleSort,
		})
	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Thesis",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:                        data.Id,
			StudentId:                 data.StudentId,
			StudentName:               data.StudentName,
			StudentNimNumber:          data.StudentNimNumber,
			StartSemesterId:           data.StartSemesterId,
			StartSemesterType:         data.StartSemesterType,
			StartSemesterSchoolYear:   data.StartSemesterSchoolYear,
			FinishSemesterId:          utils.NullStringScan(data.FinishSemesterId),
			FinishSemesterType:        utils.NullStringScan(data.FinishSemesterType),
			FinishSemesterSchoolYear:  data.FinishSemesterSchoolYear,
			Topic:                     data.Topic,
			Title:                     data.Title,
			EnglishTitle:              utils.NullStringScan(data.EnglishTitle),
			StartDate:                 data.StartDate.Format(constants.DateRFC),
			FinishDate:                utils.SafetyDate(data.FinishDate),
			Remarks:                   utils.NullStringScan(data.Remarks),
			IsJointThesis:             data.IsJointThesis,
			Status:                    data.Status,
			ProposalSeminarDate:       utils.SafetyDate(data.ProposalSeminarDate),
			ProposalCertificateNumber: utils.NullStringScan(data.ProposalCertificateNumber),
			ProposalCertificateDate:   utils.SafetyDate(data.ProposalCertificateDate),
			ThesisDefenseCount:        data.ThesisDefenseCount,
			GradePoint:                data.GradePoint,
			GradeCode:                 utils.NullStringScan(data.GradeCode),
			Files:                     files,
			ThesisSupervisors:         thesisSupervisors,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}

func (t thesisHandler) Create(w http.ResponseWriter, r *http.Request) {
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
	defer t.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Create", &in)

	data := objects.CreateThesis{
		// StudentId:                 in.GetStudentId(),
		Topic: in.GetTopic(),
		// Status:                    in.GetStatus(),
		Title:        in.GetTitle(),
		EnglishTitle: in.GetEnglishTitle(),
		// StartSemesterId:           in.GetStartSemesterId(),
		// StartDate:                 startDate,
		Remarks:         in.GetRemarks(),
		IsJointThesis:   in.GetIsJointThesis(),
		FilePath:        in.GetFilePath(),
		FilePathType:    in.GetFilePathType(),
		FileDescription: in.GetFileDescription(),
		// ProposalSeminarDate:       proposalSeminarDate,
		// ProposalCertificateNumber: in.GetProposalCertificateNumber(),
		// ProposalCertificateDate:   proposalCertificateDate,
		// ThesisSupervisors:         thesisSupervisors,
	}
	errs := t.ThesisService.Create(ctx, data)
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
			Message: "Create Thesis",
			Status:  http.StatusCreated,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusCreated, &result)
}

func (t thesisHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	defer t.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Thesis", "Update", &in)

	// startDate, errs := utils.StringToTime(in.GetStartDate())
	// if errs != nil {
	// 	utils.PrintError(*errs)
	// 	result = UpdateResponse{
	// 		Meta: &Meta{
	// 			Message: errs.Err.Error(),
	// 			Status:  uint32(errs.HttpCode),
	// 			Code:    errs.CustomCode,
	// 		},
	// 	}
	// 	utils.JSONResponse(w, errs.HttpCode, &result)
	// 	return
	// }
	// proposalSeminarDate, errs := utils.StringToTime(in.GetProposalSeminarDate())
	// if errs != nil {
	// 	utils.PrintError(*errs)
	// 	result = UpdateResponse{
	// 		Meta: &Meta{
	// 			Message: errs.Err.Error(),
	// 			Status:  uint32(errs.HttpCode),
	// 			Code:    errs.CustomCode,
	// 		},
	// 	}
	// 	utils.JSONResponse(w, errs.HttpCode, &result)
	// 	return
	// }
	// proposalCertificateDate, errs := utils.StringToTime(in.GetProposalCertificateDate())
	// if errs != nil {
	// 	utils.PrintError(*errs)
	// 	result = UpdateResponse{
	// 		Meta: &Meta{
	// 			Message: errs.Err.Error(),
	// 			Status:  uint32(errs.HttpCode),
	// 			Code:    errs.CustomCode,
	// 		},
	// 	}
	// 	utils.JSONResponse(w, errs.HttpCode, &result)
	// 	return
	// }

	var files []objects.UpdateThesisFile
	for _, v := range in.GetFiles() {
		files = append(files, objects.UpdateThesisFile{
			FilePath:        v.GetFilePath(),
			FilePathType:    v.GetFilePathType(),
			FileDescription: v.GetFileDescription(),
		})
	}

	data := objects.UpdateThesis{
		Topic:         in.GetTopic(),
		Title:         in.GetTitle(),
		EnglishTitle:  in.GetEnglishTitle(),
		Remarks:       in.GetRemarks(),
		IsJointThesis: in.GetIsJointThesis(),
		Files:         files,
	}
	errs := t.ThesisService.Update(ctx, data)
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
			Message: "Update Thesis",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
	}
	utils.JSONResponse(w, http.StatusOK, &result)
}
