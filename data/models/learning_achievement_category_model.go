package models

type GetLearningAchievementCategory struct {
	Id             string `db:"id"`
	CurriculumId   string `db:"curriculum_id"`
	CurriculumName string `db:"curriculum_name"`
	Name           string `db:"name"`
	EnglishName    string `db:"english_name"`
}

type CreateLearningAchievementCategory struct {
	CurriculumId string `db:"curriculum_id"`
	Name         string `db:"name"`
	EnglishName  string `db:"english_name"`
	CreatedBy    string `db:"created_by"`
}

type UpdateLearningAchievementCategory struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	EnglishName string `db:"english_name"`
	UpdatedBy   string `db:"updated_by"`
}
