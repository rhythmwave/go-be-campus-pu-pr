package thesis

import (
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/objects"
	appUtils "github.com/sccicitb/pupr-backend/utils"
)

func (t thesisService) mapGetDetail(data models.GetDetailThesis, fileData []models.GetThesisFile, supervisorData []models.GetThesisSupervisor) (objects.GetDetailThesis, *constants.ErrorResponse) {
	var result objects.GetDetailThesis

	var files []objects.GetDetailThesisFile
	for _, v := range fileData {
		fileUrl, errs := t.Storage.GetURL(v.FilePath, v.FilePathType, nil)
		if errs != nil {
			return result, errs
		}

		files = append(files, objects.GetDetailThesisFile{
			Id:              v.Id,
			FileUrl:         fileUrl,
			FilePath:        v.FilePath,
			FilePathType:    v.FilePathType,
			FileDescription: v.FileDescription,
		})
	}

	var supervisors []objects.GetDetailThesisSupervisor
	for _, v := range supervisorData {
		supervisors = append(supervisors, objects.GetDetailThesisSupervisor{
			Id:                       v.Id,
			LecturerId:               v.LecturerId,
			LecturerName:             v.LecturerName,
			LecturerFrontTitle:       v.LecturerFrontTitle,
			LecturerBackDegree:       v.LecturerBackDegree,
			ThesisSupervisorRoleId:   v.ThesisSupervisorRoleId,
			ThesisSupervisorRoleName: v.ThesisSupervisorRoleName,
			ThesisSupervisorRoleSort: v.ThesisSupervisorRoleSort,
		})
	}

	var finishSemesterSchoolYear string
	if data.FinishSemesterStartYear != nil {
		finishSemesterSchoolYear = appUtils.GenerateSchoolYear(*data.FinishSemesterStartYear)
	}

	result = objects.GetDetailThesis{
		Id:                        data.Id,
		StudyProgramId:            data.StudyProgramId,
		StudentId:                 data.StudentId,
		StudentName:               data.StudentName,
		StudentNimNumber:          data.StudentNimNumber,
		StartSemesterId:           data.StartSemesterId,
		StartSemesterType:         data.StartSemesterType,
		StartSemesterSchoolYear:   appUtils.GenerateSchoolYear(data.StartSemesterStartYear),
		FinishSemesterId:          data.FinishSemesterId,
		FinishSemesterType:        data.FinishSemesterType,
		FinishSemesterSchoolYear:  finishSemesterSchoolYear,
		Topic:                     data.Topic,
		Title:                     data.Title,
		EnglishTitle:              data.EnglishTitle,
		StartDate:                 data.StartDate,
		FinishDate:                data.FinishDate,
		Remarks:                   data.Remarks,
		IsJointThesis:             data.IsJointThesis,
		Status:                    data.Status,
		ProposalSeminarDate:       data.ProposalSeminarDate,
		ProposalCertificateNumber: data.ProposalCertificateNumber,
		ProposalCertificateDate:   data.ProposalCertificateDate,
		ThesisDefenseCount:        data.ThesisDefenseCount,
		GradePoint:                data.GradePoint,
		GradeCode:                 data.GradeCode,
		Files:                     files,
		ThesisSupervisors:         supervisors,
	}

	return result, nil
}

func mapGetThesisSupervisorLog(lecturerData []models.GetLecturerList, logData []models.GetThesisSupervisorLog) []objects.GetThesisSupervisorLog {
	var result []objects.GetThesisSupervisorLog

	logMap := make(map[string][]objects.GetThesisSupervisorLogThesisSupervisorRole)
	activeThesisMap := make(map[string]uint32)
	for _, v := range logData {
		logMap[v.LecturerId] = append(logMap[v.LecturerId], objects.GetThesisSupervisorLogThesisSupervisorRole{
			Id:    v.ThesisSupervisorRoleId,
			Name:  v.ThesisSupervisorRoleName,
			Total: v.Total,
		})
		activeThesisMap[v.LecturerId] += v.Total
	}

	for _, v := range lecturerData {
		result = append(result, objects.GetThesisSupervisorLog{
			Id:                     v.Id,
			IdNationalLecturer:     v.IdNationalLecturer,
			Name:                   v.Name,
			TotalSupervisedThesis:  v.TotalSupervisedThesis,
			ActiveSupervisedThesis: activeThesisMap[v.Id],
			ThesisSupervisorRoles:  logMap[v.Id],
		})
	}

	return result
}
