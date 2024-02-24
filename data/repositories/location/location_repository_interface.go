package location

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type LocationRepositoryInterface interface {
	GetListCountry(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.GetLocationList, common.Pagination, *constants.ErrorResponse)
	GetListProvince(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, countryId uint32) ([]models.GetLocationList, common.Pagination, *constants.ErrorResponse)
	GetListRegency(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, provinceId uint32) ([]models.GetLocationList, common.Pagination, *constants.ErrorResponse)
	GetListDistrict(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, regencyId uint32) ([]models.GetLocationList, common.Pagination, *constants.ErrorResponse)
	GetListVillage(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest, districtId uint32) ([]models.GetLocationList, common.Pagination, *constants.ErrorResponse)
	TempCreateData(ctx context.Context, tx *sqlx.Tx, data models.TempCreateData) *constants.ErrorResponse
	TempGetData(ctx context.Context, tx *sqlx.Tx, pagination common.PaginationRequest) ([]models.TempGetDataList, common.Pagination, *constants.ErrorResponse)
}

func NewLocationRepository(db *db.DB) LocationRepositoryInterface {
	return &locationRepository{
		db,
	}
}
