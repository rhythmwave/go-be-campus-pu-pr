package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetLearningAchievement struct {
	Id          string
	Name        string
	EnglishName string
}

type LearningAchievementListWithPagination struct {
	Pagination common.Pagination
	Data       []GetLearningAchievement
}

type CreateLearningAchievement struct {
	LearningAchievementCategoryId string
	Name                          string
	EnglishName                   string
}

type UpdateLearningAchievement struct {
	Id          string
	Name        string
	EnglishName string
}
