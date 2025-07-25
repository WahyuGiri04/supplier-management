package handler

import (
	baseHandler "supplier-management/handler/base"
	"supplier-management/model"
	"supplier-management/repo"
	"supplier-management/service"
)

type ContactsHandlerInterface interface {
	baseHandler.BaseHandlerInterface[model.Contact]
}

type ContactsHandler struct {
	*baseHandler.BaseHandler[model.Contact]
	service service.ContactsServiceInterface
}

func NewContactsHandler() ContactsHandlerInterface {
	repo := repo.NewContactsRepo()
	service := service.NewContactsService(repo)
	return &ContactsHandler{
		BaseHandler: baseHandler.NewBaseHandler(service),
		service:     service,
	}
}
