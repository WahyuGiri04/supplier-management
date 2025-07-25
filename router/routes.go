package router

import (
	"supplier-management/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine, path string) {
	supplierGroupHandler := handler.NewSupplierGroupHandler()
	supplierGroupRoute := route.Group(path + "/supplier-groups")
	{
		supplierGroupRoute.POST("/", supplierGroupHandler.Create)
		supplierGroupRoute.GET("/", supplierGroupHandler.GetAll)
		supplierGroupRoute.GET("/pagination", supplierGroupHandler.GetPagination)
		supplierGroupRoute.GET("/:uuid", supplierGroupHandler.GetByUUID)
		supplierGroupRoute.PUT("/:uuid", supplierGroupHandler.Update)
	}
	supplierHanlder := handler.NewSupplierHandler()
	supplierRoute := route.Group(path + "/suppliers")
	{
		supplierRoute.POST("/", supplierHanlder.Create)
		supplierRoute.GET("/", supplierHanlder.GetAll)
		supplierRoute.GET("/pagination", supplierHanlder.GetPagination)
		supplierRoute.GET("/:uuid", supplierHanlder.GetByUUID)
		supplierRoute.PUT("/:uuid", supplierHanlder.Update)
	}
	addressHandler := handler.NewAddressHandler()
	addressRoute := route.Group(path + "/address")
	{
		addressRoute.POST("/", addressHandler.Create)
		addressRoute.GET("/", addressHandler.GetAll)
		addressRoute.GET("/pagination", addressHandler.GetPagination)
		addressRoute.GET("/:uuid", addressHandler.GetByUUID)
		addressRoute.PUT("/:uuid", addressHandler.Update)
	}

	contactHandler := handler.NewContactsHandler()
	contactRoute := route.Group(path + "/contacts")
	{
		contactRoute.POST("/", contactHandler.Create)
		contactRoute.GET("/", contactHandler.GetAll)
		contactRoute.GET("/pagination", contactHandler.GetPagination)
		contactRoute.GET("/:uuid", contactHandler.GetByUUID)
		contactRoute.PUT("/:uuid", contactHandler.Update)
	}
}
