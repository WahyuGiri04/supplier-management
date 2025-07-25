package handler

import (
	baseHandler "supplier-management/handler/base"
	"supplier-management/model"
	"supplier-management/repo"
	"supplier-management/service"
)

type AddressHandlerInterface interface {
	baseHandler.BaseHandlerInterface[model.Address]
}

type AddressHandler struct {
	*baseHandler.BaseHandler[model.Address]
	service service.AddressServiceInterface
}

func NewAddressHandler() AddressHandlerInterface {
	repo := repo.NewAddressRepo()
	service := service.NewAddressService(repo)
	return &AddressHandler{
		BaseHandler: baseHandler.NewBaseHandler(service),
		service:     service,
	}
}