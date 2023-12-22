package routes

import (
	"TableTru/api/controller"
	"TableTru/infrastructure"
)

type StoreImageRoute struct {
	Controller controller.StoreImageController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewStoreImageRoute(
	controller controller.StoreImageController,
	handler infrastructure.GinRouter,

) StoreImageRoute {
	return StoreImageRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (u StoreImageRoute) Setup() {
	storeImage := u.Handler.Gin.Group("/storeImages") //Router group
	{
		storeImage.GET("/", u.Controller.GetAllStoreImage)
		storeImage.POST("/", u.Controller.AddStoreImage)
		storeImage.GET("/:id", u.Controller.GetStoreImage)
		storeImage.DELETE("/:id", u.Controller.DeleteStoreImage)
		storeImage.PUT("/:id", u.Controller.UpdateStoreImage)
	}
}
