package transcript

import (
	"net/http"
	"time"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/service"
	"github.com/sccicitb/pupr-backend/utils"
)

type studentTranscriptHandler struct {
	*service.ServiceCtx
}

func (l studentTranscriptHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	var result GetDetailResponse

	ctx := r.Context()

	defer l.LecturerStudentActivityLogService.Create(ctx, r, time.Now(), "Transcript", "Get Detail Transcript", nil)

	data, errs := l.StudentSubjectService.GetDetail(ctx, "")
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

	subjects := []*GetDetailResponseDataSubject{}
	for _, v := range data.Subjects {
		subjects = append(subjects, &GetDetailResponseDataSubject{
			SemesterType:                v.SemesterType,
			SemesterStartYear:           v.SemesterStartYear,
			SemesterSchoolYear:          v.SemesterSchoolYear,
			SubjectCode:                 v.SubjectCode,
			SubjectName:                 v.SubjectName,
			SubjectTheoryCredit:         v.SubjectTheoryCredit,
			SubjectPracticumCredit:      v.SubjectPracticumCredit,
			SubjectFieldPracticumCredit: v.SubjectFieldPracticumCredit,
			GradePoint:                  v.GradePoint,
			GradeCode:                   utils.NullStringScan(v.GradeCode),
		})
	}

	result = GetDetailResponse{
		Meta: &Meta{
			Message: "Get Detail Transcript",
			Status:  http.StatusOK,
			Code:    constants.SuccessCode,
		},
		Data: &GetDetailResponseData{
			Id:               data.Id,
			Name:             data.Name,
			NimNumber:        data.NimNumber,
			StudyProgramName: utils.NullStringScan(data.StudyProgramName),
			TotalCredit:      utils.NullUint32Scan(data.TotalCredit),
			Gpa:              utils.NullFloatScan(data.Gpa),
			Subjects:         subjects,
		},
	}

	utils.JSONResponse(w, http.StatusOK, &result)
}
