package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetLearningAchievementCategory struct {
	Id          string
	Name        string
	EnglishName string
}

type LearningAchievementCategoryListWithPagination struct {
	Pagination common.Pagination
	Data       []GetLearningAchievementCategory
}

type CreateLearningAchievementCategory struct {
	CurriculumId string
	Name         string
	EnglishName  string
}

type UpdateLearningAchievementCategory struct {
	Id          string
	Name        string
	EnglishName string
}
