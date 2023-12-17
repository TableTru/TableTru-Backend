package routes

import (
	"TableTru/api/controller"
	"TableTru/infrastructure"
)

type UserRoute struct {
	Controller controller.UserController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewUserRoute(
	controller controller.UserController,
	handler infrastructure.GinRouter,

) UserRoute {
	return UserRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (u UserRoute) Setup() {
	user := u.Handler.Gin.Group("/users") //Router group
	{
		user.GET("/", u.Controller.GetAllUser)
		user.POST("/", u.Controller.AddUser)
		// user.GET("/:id", u.Controller.GetPost)
		// user.DELETE("/:id", u.Controller.DeletePost)
		// user.PUT("/:id", u.Controller.UpdatePost)
	}
}
