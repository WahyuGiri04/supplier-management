package repo

import (
	"supplier-management/model"
	baseRepository "supplier-management/repo/base"
)

type ContactsRepoInterface interface {
	baseRepository.BaseRepositoryInterface[model.Contact]
}

type ContactsRepo struct {
	*baseRepository.BaseRepository[model.Contact]
}

func NewContactsRepo() ContactsRepoInterface {
	return &ContactsRepo{
		BaseRepository: baseRepository.NewBaseRepository[model.Contact](),
	}
}
