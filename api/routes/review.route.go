package routes

import (
	"TableTru/api/controller"
	"TableTru/infrastructure"
)

type ReviewRoute struct {
	Controller controller.ReviewController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewReviewRoute(
	controller controller.ReviewController,
	handler infrastructure.GinRouter,

) ReviewRoute {
	return ReviewRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (u ReviewRoute) Setup() {
	review := u.Handler.Gin.Group("/reviews") //Router group
	{
		review.GET("/", u.Controller.GetAllReview)
		review.POST("/", u.Controller.AddReview)
		review.GET("/:id", u.Controller.GetReview)
		review.DELETE("/:id", u.Controller.DeleteReview)
		review.PUT("/:id", u.Controller.UpdateReview)
	}
}
