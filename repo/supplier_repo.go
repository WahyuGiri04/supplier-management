package repo

import (
	"supplier-management/model"
	baseRepository "supplier-management/repo/base"
)

type SupplierRepoInterface interface {
	baseRepository.BaseRepositoryInterface[model.Supplier]
}

type SupplierRepo struct {
	*baseRepository.BaseRepository[model.Supplier]
}

func NewSupplierRepo() SupplierRepoInterface {
	return &SupplierRepo{
		BaseRepository: baseRepository.NewBaseRepository[model.Supplier](),
	}
}