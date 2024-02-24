package models

type GetCreditQuota struct {
	Id                string  `db:"id"`
	MinimumGradePoint float64 `db:"minimum_grade_point"`
	MaximumGradePoint float64 `db:"maximum_grade_point"`
	MaximumCredit     uint32  `db:"maximum_credit"`
}

type CreateCreditQuota struct {
	MinimumGradePoint float64 `db:"minimum_grade_point"`
	MaximumGradePoint float64 `db:"maximum_grade_point"`
	MaximumCredit     uint32  `db:"maximum_credit"`
	CreatedBy         string  `db:"created_by"`
}

type UpdateCreditQuota struct {
	Id                string  `db:"id"`
	MinimumGradePoint float64 `db:"minimum_grade_point"`
	MaximumGradePoint float64 `db:"maximum_grade_point"`
	MaximumCredit     uint32  `db:"maximum_credit"`
	UpdatedBy         string  `db:"updated_by"`
}
