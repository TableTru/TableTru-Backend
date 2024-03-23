package routes

import (
	"TableTru/api/controller"
	"TableTru/infrastructure"
)

type StoreRoute struct {
	Controller controller.StoreController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewStoreRoute(
	controller controller.StoreController,
	handler infrastructure.GinRouter,

) StoreRoute {
	return StoreRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (u StoreRoute) Setup() {
	store := u.Handler.Gin.Group("/stores") //Router group
	{
		store.GET("/", u.Controller.GetAllStore)
		store.POST("/", u.Controller.AddStore)
		store.GET("/:id", u.Controller.GetStore)
		store.DELETE("/:id", u.Controller.DeleteStore)
		store.PUT("/:id", u.Controller.UpdateStore)

		store.GET("/checkStoreByName", u.Controller.CheckStoreByName)
		store.GET("/getStorePreview", u.Controller.GetStoreByNum)
		store.GET("/SearchStoreSortRating", u.Controller.SearchStoreSortRating)
		store.GET("/SearchStoreSortDistance", u.Controller.SearchStoreSortDistance)
		
	}
}
