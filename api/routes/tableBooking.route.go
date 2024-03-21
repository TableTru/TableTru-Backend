package routes

import (
	"TableTru/api/controller"
	"TableTru/infrastructure"
)

type TableBookingRoute struct {
	Controller controller.TableBookingController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewTableBookingRoute(
	controller controller.TableBookingController,
	handler infrastructure.GinRouter,

) TableBookingRoute {
	return TableBookingRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (u TableBookingRoute) Setup() {
	tableBooking := u.Handler.Gin.Group("/tableBookings") //Router group
	{
		tableBooking.GET("/", u.Controller.GetAllTableBooking)
		tableBooking.POST("/", u.Controller.AddTableBooking)
		tableBooking.GET("/:id", u.Controller.GetTableBooking)
		tableBooking.DELETE("/:id", u.Controller.DeleteTableBooking)
		tableBooking.PUT("/:id", u.Controller.UpdateTableBooking)

		tableBooking.GET("/GetUserBookingByStatus", u.Controller.GetUserBookingByStatus)
		tableBooking.GET("/GetStoreBookingByStatus", u.Controller.GetStoreBookingByStatus)
		tableBooking.GET("/CheckBookingTime", u.Controller.CheckBookingTime)
	}
}
