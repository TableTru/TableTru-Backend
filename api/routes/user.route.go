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
		user.GET("/:id", u.Controller.GetUser)
		user.DELETE("/:id", u.Controller.DeleteUser)
		user.PUT("/:id", u.Controller.UpdateUser)

		user.GET("/getLoginUser", u.Controller.GetLoginUser)
		user.GET("/checkRegisterUser", u.Controller.CheckRegisterUser)
	}
}
