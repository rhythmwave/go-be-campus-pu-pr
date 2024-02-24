package models

type GetGraduationPredicate struct {
	Id                          string  `db:"id"`
	Predicate                   string  `db:"predicate"`
	MinimumGpa                  float64 `db:"minimum_gpa"`
	MaximumStudySemester        uint32  `db:"maximum_study_semester"`
	RepeatCourseLimit           uint32  `db:"repeat_course_limit"`
	BelowMinimumGradePointLimit uint32  `db:"below_minimum_grade_point_limit"`
}

type CreateGraduationPredicate struct {
	Predicate                   string  `db:"predicate"`
	MinimumGpa                  float64 `db:"minimum_gpa"`
	MaximumStudySemester        uint32  `db:"maximum_study_semester"`
	RepeatCourseLimit           uint32  `db:"repeat_course_limit"`
	BelowMinimumGradePointLimit uint32  `db:"below_minimum_grade_point_limit"`
	CreatedBy                   string  `db:"created_by"`
}

type UpdateGraduationPredicate struct {
	Id                          string  `db:"id"`
	Predicate                   string  `db:"predicate"`
	MinimumGpa                  float64 `db:"minimum_gpa"`
	MaximumStudySemester        uint32  `db:"maximum_study_semester"`
	RepeatCourseLimit           uint32  `db:"repeat_course_limit"`
	BelowMinimumGradePointLimit uint32  `db:"below_minimum_grade_point_limit"`
	UpdatedBy                   string  `db:"updated_by"`
}
