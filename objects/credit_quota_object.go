package objects

import "github.com/sccicitb/pupr-backend/objects/common"

type GetCreditQuota struct {
	Id                string
	MinimumGradePoint float64
	MaximumGradePoint float64
	MaximumCredit     uint32
}

type CreditQuotaListWithPagination struct {
	Pagination common.Pagination
	Data       []GetCreditQuota
}

type CreateCreditQuota struct {
	MinimumGradePoint float64
	MaximumGradePoint float64
	MaximumCredit     uint32
}

type UpdateCreditQuota struct {
	Id                string
	MinimumGradePoint float64
	MaximumGradePoint float64
	MaximumCredit     uint32
}
