package service

import (
	"supplier-management/model"
	"supplier-management/repo"
	baseService "supplier-management/service/base"
)

type SupplierGroupServiceInterface interface {
	baseService.BaseServiceInterface[model.SupplierGroup]
}

type SupplierGroupService struct {
	*baseService.BaseService[model.SupplierGroup]
	repo repo.SupplierGroupRepoInterface
}

func NewSupplierGroupService(repo repo.SupplierGroupRepoInterface) SupplierGroupServiceInterface {
	return &SupplierGroupService{
		BaseService: baseService.NewBaseService(repo),
		repo:        repo,
	}
}
