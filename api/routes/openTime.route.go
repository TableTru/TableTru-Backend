package routes

import (
	"TableTru/api/controller"
	"TableTru/infrastructure"
)

type OpenTimeRoute struct {
	Controller controller.OpenTimeController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewOpenTimeRoute(
	controller controller.OpenTimeController,
	handler infrastructure.GinRouter,

) OpenTimeRoute {
	return OpenTimeRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (u OpenTimeRoute) Setup() {
	openTime := u.Handler.Gin.Group("/opentimes") //Router group
	{
		openTime.GET("/", u.Controller.GetAllOpenTime)
		openTime.POST("/", u.Controller.AddOpenTime)
		openTime.GET("/:id", u.Controller.GetOpenTime)
		openTime.DELETE("/:id", u.Controller.DeleteOpenTime)
		openTime.PUT("/:id", u.Controller.UpdateOpenTime)
	}
}
