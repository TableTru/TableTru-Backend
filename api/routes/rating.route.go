package routes

import (
	"TableTru/api/controller"
	"TableTru/infrastructure"
)

type RatingRoute struct {
	Controller controller.RatingController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewRatingRoute(
	controller controller.RatingController,
	handler infrastructure.GinRouter,

) RatingRoute {
	return RatingRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (u RatingRoute) Setup() {
	rating := u.Handler.Gin.Group("/ratings") //Router group
	{
		rating.GET("/", u.Controller.GetAllRating)
		rating.POST("/", u.Controller.AddRating)
		rating.GET("/:id", u.Controller.GetRating)
		rating.DELETE("/:id", u.Controller.DeleteRating)
		rating.PUT("/:id", u.Controller.UpdateRating)
	}
}
