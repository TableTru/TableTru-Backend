package routes

import (
	"TableTru/api/controller"
	"TableTru/infrastructure"
)

type ReviewImageRoute struct {
	Controller controller.ReviewImageController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewReviewImageRoute(
	controller controller.ReviewImageController,
	handler infrastructure.GinRouter,

) ReviewImageRoute {
	return ReviewImageRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (u ReviewImageRoute) Setup() {
	reviewImage := u.Handler.Gin.Group("/reviewImages") //Router group
	{
		reviewImage.GET("/", u.Controller.GetAllReviewImage)
		reviewImage.POST("/", u.Controller.AddReviewImage)
		reviewImage.GET("/:id", u.Controller.GetReviewImage)
		reviewImage.DELETE("/:id", u.Controller.DeleteReviewImage)
		reviewImage.PUT("/:id", u.Controller.UpdateReviewImage)
	}
}
