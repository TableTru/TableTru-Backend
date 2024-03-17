package routes

import (
	"TableTru/api/controller"
	"TableTru/infrastructure"
)

type PromotionRoute struct {
	Controller controller.PromotionController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewPromotionRoute(
	controller controller.PromotionController,
	handler infrastructure.GinRouter,

) PromotionRoute {
	return PromotionRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (u PromotionRoute) Setup() {
	promotion := u.Handler.Gin.Group("/promotions") //Router group
	{
		promotion.GET("/", u.Controller.GetAllPromotion)
		promotion.POST("/", u.Controller.AddPromotion)
		promotion.GET("/:id", u.Controller.GetPromotion)
		promotion.DELETE("/:id", u.Controller.DeletePromotion)
		promotion.PUT("/:id", u.Controller.UpdatePromotion)
		
		promotion.GET("/GetAllPromotionByStoreId", u.Controller.GetAllPromotionByStoreId)
		promotion.GET("/GetAllPromotionByUserId", u.Controller.GetAllPromotionByUserId)
	}
}
