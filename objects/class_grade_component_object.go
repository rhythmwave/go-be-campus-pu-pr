package objects

type GetClassGradeComponent struct {
	Id         string
	Name       string
	Percentage float64
	IsActive   bool
}

type SetClassGradeComponent struct {
	Name       string
	Percentage float64
	IsActive   bool
}
