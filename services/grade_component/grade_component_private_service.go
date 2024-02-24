package grade_component

import (
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/objects"
)

func mapGetListBySubjectCategory(modelResult []models.GetGradeComponentDistinctSubjectCategory, gradeComponentData []models.GetPercentageBySubjectCategories) []objects.GetGradeComponentBySubjectCategory {
	results := []objects.GetGradeComponentBySubjectCategory{}

	tempGradeComponent := make(map[string][]objects.GetGradeComponentBySubjectCategoryGradeComponent)

	for _, v := range gradeComponentData {
		singleSubjectData := tempGradeComponent[v.SubjectCategoryId]

		singleSubjectData = append(singleSubjectData, objects.GetGradeComponentBySubjectCategoryGradeComponent{
			Id:                v.Id,
			Name:              v.Name,
			DefaultPercentage: v.DefaultPercentage,
			IsActive:          v.IsActive,
		})

		tempGradeComponent[v.SubjectCategoryId] = singleSubjectData
	}

	for _, v := range modelResult {
		results = append(results, objects.GetGradeComponentBySubjectCategory{
			StudyProgramId:      v.StudyProgramId,
			StudyProgramName:    v.StudyProgramName,
			SubjectCategoryId:   v.SubjectCategoryId,
			SubjectCategoryName: v.SubjectCategoryName,
			GradeComponents:     tempGradeComponent[v.SubjectCategoryId],
		})
	}

	return results
}
