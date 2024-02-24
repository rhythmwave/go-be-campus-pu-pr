package transcript

import (
	"net/http"
	"time"

	"github.com/gorilla/schema"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/utils"
	"github.com/sirupsen/logrus"
)

type adminTranscriptHandler struct {
	*service.ServiceCtx
}

func (l adminTranscriptHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
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
	defer l.AdminActivityLogService.Create(ctx, r, time.Now(), "Transcript", "Get Detail Transcript", nil)

	data, errs := l.StudentSubjectService.GetPdfDetail(ctx, in.GetStudentId())
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

	semesters := []*GetDetailResponseDataSemester{}
	for _, v := range data.Semesters {
		subjects := []*GetDetailResponseDataSemesterSubject{}
		for _, w := range v.Subjects {
			subjects = append(subjects, &GetDetailResponseDataSemesterSubject{
				SubjectCode:        w.SubjectCode,
				SubjectName:        w.SubjectName,
				SubjectEnglishName: utils.NullStringScan(w.SubjectEnglishName),
				TheoryCredit:       w.TheoryCredit,
				PracticumCredit:    w.PracticumCredit,
				GradeCode:          utils.NullStringScan(w.GradeCode),
			})
		}

		semesters = append(semesters, &GetDetailResponseDataSemester{
			SemesterPackage: v.SemesterPackage,
			Subjects:        subjects,
		})
	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Transcript",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			NimNumber:           data.NimNumber,
			Name:                data.Name,
			BirthRegencyName:    utils.NullStringScan(data.BirthRegencyName),
			BirthDate:           utils.SafetyDate(data.BirthDate),
			GraduationDate:      utils.SafetyDate(data.GraduationDate),
			DiplomaNumber:       utils.NullStringScan(data.DiplomaNumber),
			StudyProgramName:    utils.NullStringScan(data.StudyProgramName),
			StudyLevelName:      utils.NullStringScan(data.StudyLevelName),
			StudyLevelShortName: utils.NullStringScan(data.StudyLevelShortName),
			TotalCredit:         data.TotalCredit,
			Gpa:                 utils.NullFloatScan(data.Gpa),
			GraduationPredicate: utils.NullStringScan(data.GraduationPredicate),
			TheoryCredit:        data.TheoryCredit,
			PracticumCredit:     data.PracticumCredit,
			ThesisTitle:         utils.NullStringScan(data.ThesisTitle),
			ThesisEnglishTitle:  utils.NullStringScan(data.ThesisEnglishTitle),
			Semesters:           semesters,
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}
