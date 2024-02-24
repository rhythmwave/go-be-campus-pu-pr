package models

type GetLearningAchievement struct {
	Id                                     string `db:"id"`
	LearningAchievementCategoryId          string `db:"learning_achievement_category_id"`
	LearningAchievementCategoryName        string `db:"learning_achievement_category_name"`
	LearningAchievementCategoryEnglishName string `db:"learning_achievement_category_english_name"`
	Name                                   string `db:"name"`
	EnglishName                            string `db:"english_name"`
}

type CreateLearningAchievement struct {
	LearningAchievementCategoryId string `db:"learning_achievement_category_id"`
	Name                          string `db:"name"`
	EnglishName                   string `db:"english_name"`
	CreatedBy                     string `db:"created_by"`
}

type UpdateLearningAchievement struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	EnglishName string `db:"english_name"`
	UpdatedBy   string `db:"updated_by"`
}
