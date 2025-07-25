package repo

import (
	"supplier-management/model"
	baseRepository "supplier-management/repo/base"
)

type AddressRepoInterface interface {
	baseRepository.BaseRepositoryInterface[model.Address]
}

type AddressRepo struct {
	*baseRepository.BaseRepository[model.Address]
}

func NewAddressRepo() AddressRepoInterface {
	return &AddressRepo{
		BaseRepository: baseRepository.NewBaseRepository[model.Address](),
	}
}
