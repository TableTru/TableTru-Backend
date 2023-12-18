package routes

import (
	"TableTru/api/controller"
	"TableTru/infrastructure"
)

type CategoryRoute struct {
	Controller controller.CategoryController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewCategoryRoute(
	controller controller.CategoryController,
	handler infrastructure.GinRouter,

) CategoryRoute {
	return CategoryRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (u CategoryRoute) Setup() {
	category := u.Handler.Gin.Group("/categories") //Router group
	{
		category.GET("/", u.Controller.GetAllCategory)
		category.POST("/", u.Controller.AddCategory)
		category.GET("/:id", u.Controller.GetCategory)
		category.DELETE("/:id", u.Controller.DeleteCategory)
		category.PUT("/:id", u.Controller.UpdateCategory)
	}
}
