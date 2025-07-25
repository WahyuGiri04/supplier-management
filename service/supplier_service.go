package service

import (
	"supplier-management/model"
	"supplier-management/repo"
	baseService "supplier-management/service/base"
)

type SupplierServiceInterface interface {
	baseService.BaseServiceInterface[model.Supplier]
}

type SupplierService struct {
	*baseService.BaseService[model.Supplier]
	repo repo.SupplierRepoInterface
}

func NewSupplierService(repo repo.SupplierRepoInterface) SupplierServiceInterface {
	return &SupplierService{
		BaseService: baseService.NewBaseService(repo),
		repo:        repo,
	}
}
