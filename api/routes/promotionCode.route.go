package routes

import (
	"TableTru/api/controller"
	"TableTru/infrastructure"
)

type PromotionCodeRoute struct {
	Controller controller.PromotionCodeController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewPromotionCodeRoute(
	controller controller.PromotionCodeController,
	handler infrastructure.GinRouter,

) PromotionCodeRoute {
	return PromotionCodeRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (u PromotionCodeRoute) Setup() {
	promotionCode := u.Handler.Gin.Group("/promotionCodes") //Router group
	{
		promotionCode.GET("/", u.Controller.GetAllPromotionCode)
		promotionCode.POST("/", u.Controller.AddPromotionCode)
		promotionCode.GET("/:id", u.Controller.GetPromotionCode)
		promotionCode.DELETE("/:id", u.Controller.DeletePromotionCode)
		promotionCode.PUT("/:id", u.Controller.UpdatePromotionCode)
	}
}
