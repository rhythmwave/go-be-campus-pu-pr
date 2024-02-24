package common

import "math"

// Pagination Pagination
type Pagination struct {
	Page         int32
	Limit        int32
	Search       string
	Sort         string
	SortBy       string
	Next         *int32
	Prev         *int32
	TotalPages   *int32
	TotalRecords *int32
}

// PaginationRequest PaginationRequest
type PaginationRequest struct {
	Search         string
	Page           uint32
	Limit          uint32
	UserID         uint32
	EmployeeID     uint32
	Sort           string
	SortBy         string
	Category       string
	Type           string
	Process        string
	Next           uint32
	Prev           uint32
	TotalPages     uint32
	TotalRecords   uint32
	StartDate      string
	EndDate        string
	PackageID      string
	CategoryID     string
	ClassroomID    string
	QuizQuestionID string
	AssignmentID   string
	MemberID       string
}

// GetPagination GetPagination
func (p *PaginationRequest) GetPagination(countRow int, page uint32, limit uint32) Pagination {
	p.TotalPages = uint32(math.Ceil(float64(countRow) / float64(limit)))
	var prev uint32 = 1
	if page > 1 {
		prev = page - 1
	}
	var next uint32 = uint32(p.TotalPages)
	if uint32(p.TotalPages) != page {
		next = page + 1
	}
	p.Prev = prev
	p.Next = next
	p.TotalRecords = uint32(countRow)

	nextPage := int32(p.Next)
	prevPage := int32(p.Prev)
	totalPages := int32(p.TotalPages)
	totalRecords := int32(p.TotalRecords)
	result := Pagination{
		Page:         int32(p.Page),
		Limit:        int32(p.Limit),
		Sort:         p.Sort,
		SortBy:       p.SortBy,
		Next:         &nextPage,
		Prev:         &prevPage,
		TotalPages:   &totalPages,
		TotalRecords: &totalRecords,
	}

	return result
}
