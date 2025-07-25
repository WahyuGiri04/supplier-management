package handler

import (
	baseHandler "supplier-management/handler/base"
	"supplier-management/model"
	"supplier-management/repo"
	"supplier-management/service"
)

type SupplierGroupHandlerInterface interface {
	baseHandler.BaseHandlerInterface[model.SupplierGroup]
}

type SupplierGroupHandler struct {
	*baseHandler.BaseHandler[model.SupplierGroup]
	service service.SupplierGroupServiceInterface
}

func NewSupplierGroupHandler() SupplierGroupHandlerInterface {
	repo := repo.NewSupplierGroupRepo()
	service := service.NewSupplierGroupService(repo)
	return &SupplierGroupHandler{
		BaseHandler: baseHandler.NewBaseHandler(service),
		service:     service,
	}
}
