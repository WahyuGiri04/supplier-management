package repo

import (
	"supplier-management/model"
	baseRepository "supplier-management/repo/base"
)

type SupplierGroupRepoInterface interface {
	baseRepository.BaseRepositoryInterface[model.SupplierGroup]
	GetByValue(value string) (*model.SupplierGroup, error)
}

type SupplierGroupRepo struct {
	*baseRepository.BaseRepository[model.SupplierGroup]
}

func NewSupplierGroupRepo() SupplierGroupRepoInterface {
	return &SupplierGroupRepo{
		BaseRepository: baseRepository.NewBaseRepository[model.SupplierGroup](),
	}
}


func (r *SupplierGroupRepo) GetByValue(value string) (*model.SupplierGroup, error) {
	var supplierGroup model.SupplierGroup
	err := r.GetDB().Where("value = ? AND is_deleted = ? AND is_active = ?", value, false, true).First(&supplierGroup).Error
	if err != nil {
		return nil, err
	}
	return &supplierGroup, nil
}