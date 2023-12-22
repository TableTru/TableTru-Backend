package routes

import (
	"TableTru/api/controller"
	"TableTru/infrastructure"
)

type LocationRoute struct {
	Controller controller.LocationController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewLocationRoute(
	controller controller.LocationController,
	handler infrastructure.GinRouter,

) LocationRoute {
	return LocationRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (u LocationRoute) Setup() {
	location := u.Handler.Gin.Group("/locations") //Router group
	{
		location.GET("/", u.Controller.GetAllLocation)
		location.POST("/", u.Controller.AddLocation)
		location.GET("/:id", u.Controller.GetLocation)
		location.DELETE("/:id", u.Controller.DeleteLocation)
		location.PUT("/:id", u.Controller.UpdateLocation)
	}
}
