package models

type GetLocationList struct {
	Id   uint64 `db:"id"`
	Name string `db:"name"`
}

type TempGetDataList struct {
	Id    string `db:"id"`
	Title string `db:"title"`
	Body  string `db:"body"`
}

type TempCreateData struct {
	Title string `db:"title"`
	Body  string `db:"body"`
}
