package handler

import (
	baseHandler "supplier-management/handler/base"
	"supplier-management/model"
	"supplier-management/repo"
	"supplier-management/service"
)

type SupplierHandlerInterface interface {
	baseHandler.BaseHandlerInterface[model.Supplier]
}

type SupplierHandler struct {
	*baseHandler.BaseHandler[model.Supplier]
	service service.SupplierServiceInterface
}

func NewSupplierHandler() SupplierHandlerInterface {
	repo := repo.NewSupplierRepo()
	service := service.NewSupplierService(repo)
	return &SupplierHandler{
		BaseHandler: baseHandler.NewBaseHandler(service),
		service:     service,
	}
}
