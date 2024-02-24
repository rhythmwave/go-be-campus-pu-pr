package utils

import (
	appConstants "github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
)

func GetMaximumCredit(creditQuotaData []models.GetCreditQuota, studentPreviousGradePoint *float64) uint32 {
	if studentPreviousGradePoint == nil {
		return appConstants.DefaultMaximumCredit
	}

	previousSemesterGradePoint := NullFloatScan(studentPreviousGradePoint)
	for _, v := range creditQuotaData {
		if studentPreviousGradePoint == nil || previousSemesterGradePoint >= v.MinimumGradePoint {
			return v.MaximumCredit
		}
	}

	return 0
}
