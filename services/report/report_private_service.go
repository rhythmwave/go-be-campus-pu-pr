package report

import (
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/objects"
)

func mapStudentStatus(data []models.GetReportStudentStatus) []objects.ReportStudentStatus {
	var result []objects.ReportStudentStatus

	resultMap := make(map[string]objects.ReportStudentStatus)
	statusMap := make(map[string][]objects.ReportStudentStatusStatus)

	for _, v := range data {
		statusMap[v.StudyProgramId] = append(statusMap[v.StudyProgramId], objects.ReportStudentStatusStatus{
			Status: v.Status,
			Total:  v.Total,
		})

		resultMap[v.StudyProgramId] = objects.ReportStudentStatus{
			StudyProgramId:        v.StudyProgramId,
			StudyProgramName:      v.StudyProgramName,
			DiktiStudyProgramCode: v.DiktiStudyProgramCode,
			DiktiStudyProgramType: v.DiktiStudyProgramType,
			StudyLevelShortName:   v.StudyLevelShortName,
		}
	}

	for k, v := range resultMap {
		singleResult := v
		singleResult.Statuses = statusMap[k]
		result = append(result, singleResult)
	}

	return result
}

func mapStudentClassGrade(subjectData []models.GetSubject, gradeData []models.GetReportStudentClassGrade) []objects.ReportStudentClassGrade {
	result := []objects.ReportStudentClassGrade{}

	gradeMap := make(map[string][]objects.ReportStudentClassGradeGrade)
	for _, v := range gradeData {
		gradeMap[v.SubjectId] = append(gradeMap[v.SubjectId], objects.ReportStudentClassGradeGrade{
			GradeCode: v.GradeCode,
			Total:     v.Total,
		})
	}

	for _, v := range subjectData {
		result = append(result, objects.ReportStudentClassGrade{
			SubjectId:   v.Id,
			SubjectName: v.Name,
			Grades:      gradeMap[v.Id],
		})
	}

	return result
}

func mapStudentProvince(resultData []models.GetReportStudentProvince) []objects.ReportStudentProvince {
	var result []objects.ReportStudentProvince

	resultMap := make(map[uint32]objects.ReportStudentProvince)
	studentForceMap := make(map[uint32][]objects.ReportStudentProvinceStudentForce)
	for _, v := range resultData {
		resultMap[v.ProvinceId] = objects.ReportStudentProvince{
			ProvinceId:   v.ProvinceId,
			ProvinceName: v.ProvinceName,
		}
		studentForceMap[v.ProvinceId] = append(studentForceMap[v.ProvinceId], objects.ReportStudentProvinceStudentForce{
			StudentForce: v.StudentForce,
			Total:        v.Total,
		})
	}

	for k, v := range resultMap {
		singleValue := v
		singleValue.StudentForces = studentForceMap[k]

		result = append(result, singleValue)
	}

	return result
}

func mapStudentSchoolProvince(resultData []models.GetReportStudentSchoolProvince) []objects.ReportStudentSchoolProvince {
	var result []objects.ReportStudentSchoolProvince

	resultMap := make(map[uint32]objects.ReportStudentSchoolProvince)
	studentForceMap := make(map[uint32][]objects.ReportStudentSchoolProvinceStudentForce)
	for _, v := range resultData {
		resultMap[v.ProvinceId] = objects.ReportStudentSchoolProvince{
			ProvinceId:   v.ProvinceId,
			ProvinceName: v.ProvinceName,
		}
		studentForceMap[v.ProvinceId] = append(studentForceMap[v.ProvinceId], objects.ReportStudentSchoolProvinceStudentForce{
			StudentForce: v.StudentForce,
			Total:        v.Total,
		})
	}

	for k, v := range resultMap {
		singleValue := v
		singleValue.StudentForces = studentForceMap[k]

		result = append(result, singleValue)
	}

	return result
}
