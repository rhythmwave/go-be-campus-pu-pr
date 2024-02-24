package objects

type GetSubjectGradeComponent struct {
	Id         string
	Name       string
	Percentage float64
	IsActive   bool
}

type SetSubjectGradeComponent struct {
	Name       string
	Percentage float64
	IsActive   bool
}
