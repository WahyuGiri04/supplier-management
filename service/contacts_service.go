package service

import (
	"supplier-management/model"
	"supplier-management/repo"
	baseService "supplier-management/service/base"
)

type ContactsServiceInterface interface {
	baseService.BaseServiceInterface[model.Contact]
}

type ContactsService struct {
	*baseService.BaseService[model.Contact]
	repo repo.ContactsRepoInterface
}

func NewContactsService(repo repo.ContactsRepoInterface) ContactsServiceInterface {
	return &ContactsService{
		BaseService: baseService.NewBaseService(repo),
		repo:        repo,
	}
}
