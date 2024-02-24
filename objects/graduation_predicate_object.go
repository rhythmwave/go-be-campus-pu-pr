package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetGraduationPredicate struct {
	Id                          string
	Predicate                   string
	MinimumGpa                  float64
	MaximumStudySemester        uint32
	RepeatCourseLimit           uint32
	BelowMinimumGradePointLimit uint32
}

type GraduationPredicateListWithPagination struct {
	Pagination common.Pagination
	Data       []GetGraduationPredicate
}

type CreateGraduationPredicate struct {
	Predicate                   string
	MinimumGpa                  float64
	MaximumStudySemester        uint32
	RepeatCourseLimit           uint32
	BelowMinimumGradePointLimit uint32
}

type UpdateGraduationPredicate struct {
	Id                          string
	Predicate                   string
	MinimumGpa                  float64
	MaximumStudySemester        uint32
	RepeatCourseLimit           uint32
	BelowMinimumGradePointLimit uint32
}
