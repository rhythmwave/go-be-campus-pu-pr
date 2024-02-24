package location

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/models"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	"github.com/sccicitb/pupr-backend/infra/context/repository"
	"github.com/sccicitb/pupr-backend/objects"
	"github.com/sccicitb/pupr-backend/objects/common"
)

type locationService struct {
	*repository.RepoCtx
	*infra.InfraCtx
}

func (p locationService) GetListCountry(ctx context.Context, paginationData common.PaginationRequest) (objects.LocationListWithPagination, *constants.ErrorResponse) {
	var result objects.LocationListWithPagination

	tx, err := p.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := p.LocationRepo.GetListCountry(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLocation{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLocation{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	result = objects.LocationListWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (p locationService) GetListProvince(ctx context.Context, paginationData common.PaginationRequest, countryId uint32) (objects.LocationListWithPagination, *constants.ErrorResponse) {
	var result objects.LocationListWithPagination

	tx, err := p.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := p.LocationRepo.GetListProvince(ctx, tx, paginationData, countryId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLocation{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLocation{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	result = objects.LocationListWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (p locationService) GetListRegency(ctx context.Context, paginationData common.PaginationRequest, provinceId uint32) (objects.LocationListWithPagination, *constants.ErrorResponse) {
	var result objects.LocationListWithPagination

	tx, err := p.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := p.LocationRepo.GetListRegency(ctx, tx, paginationData, provinceId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLocation{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLocation{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	result = objects.LocationListWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (p locationService) GetListDistrict(ctx context.Context, paginationData common.PaginationRequest, regencyId uint32) (objects.LocationListWithPagination, *constants.ErrorResponse) {
	var result objects.LocationListWithPagination

	tx, err := p.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := p.LocationRepo.GetListDistrict(ctx, tx, paginationData, regencyId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLocation{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLocation{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	result = objects.LocationListWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (p locationService) GetListVillage(ctx context.Context, paginationData common.PaginationRequest, districtId uint32) (objects.LocationListWithPagination, *constants.ErrorResponse) {
	var result objects.LocationListWithPagination

	tx, err := p.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := p.LocationRepo.GetListVillage(ctx, tx, paginationData, districtId)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.GetLocation{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.GetLocation{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	result = objects.LocationListWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}

func (a locationService) TempCreateData(ctx context.Context, data objects.TempCreateData) *constants.ErrorResponse {
	tx, err := a.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	createData := models.TempCreateData{
		Title: data.Title,
		Body:  data.Body,
	}
	errs := a.LocationRepo.TempCreateData(ctx, tx, createData)
	if errs != nil {
		_ = tx.Rollback()
		return errs
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	return nil
}

func (p locationService) TempGetData(ctx context.Context, paginationData common.PaginationRequest) (objects.TempGetDataWithPagination, *constants.ErrorResponse) {
	var result objects.TempGetDataWithPagination

	tx, err := p.DB.Begin(ctx)
	if err != nil {
		return result, constants.ErrUnknown
	}

	modelResult, paginationResult, errs := p.LocationRepo.TempGetData(ctx, tx, paginationData)
	if errs != nil {
		_ = tx.Rollback()
		return result, errs
	}

	resultData := []objects.TempGetData{}
	for _, v := range modelResult {
		resultData = append(resultData, objects.TempGetData{
			Id:    v.Id,
			Title: v.Title,
			Body:  v.Body,
		})
	}

	result = objects.TempGetDataWithPagination{
		Pagination: paginationResult,
		Data:       resultData,
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return result, constants.ErrUnknown
	}

	return result, nil
}
