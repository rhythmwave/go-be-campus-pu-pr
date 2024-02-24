package objects

import "github.com/sccicitb/pupr-backend/objects/common"

type GetLocation struct {
	Id   uint64
	Name string
}

type LocationListWithPagination struct {
	Pagination common.Pagination
	Data       []GetLocation
}

type TempGetData struct {
	Id    string
	Title string
	Body  string
}

type TempGetDataWithPagination struct {
	Pagination common.Pagination
	Data       []TempGetData
}

type TempCreateData struct {
	Title string
	Body  string
}
