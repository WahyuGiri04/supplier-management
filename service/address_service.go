package service

import (
	"supplier-management/model"
	"supplier-management/repo"
	baseService "supplier-management/service/base"
)

type AddressServiceInterface interface {
	baseService.BaseServiceInterface[model.Address]
}

type AddressService struct {
	*baseService.BaseService[model.Address]
	repo repo.AddressRepoInterface
}

func NewAddressService(repo repo.AddressRepoInterface) AddressServiceInterface {
	return &AddressService{
		BaseService: baseService.NewBaseService(repo),
		repo:        repo,
	}
}
