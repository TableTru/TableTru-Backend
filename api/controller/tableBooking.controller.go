package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type TableBookingController struct {
	service service.TableBookingService
}

func NewTableBookingController(s service.TableBookingService) TableBookingController {
	return TableBookingController{
		service: s,
	}
}

func (c TableBookingController) GetAllTableBooking(ctx *gin.Context) {
	var tableBookings models.TableBooking

	keyword := ctx.Query("keyword")

	data, total, err := c.service.FindAllTableBooking(tableBookings, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "TableBooking result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *TableBookingController) GetTableBooking(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var tableBooking models.TableBooking
	tableBooking.ID = id
	foundTableBooking, err := p.service.FindTableBooking(tableBooking)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding TableBooking")
		return
	}
	response := foundTableBooking.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of TableBooking",
		Data:    &response})

}

func (p *TableBookingController) AddTableBooking(ctx *gin.Context) {
	var tableBooking models.TableBooking
	ctx.ShouldBindJSON(&tableBooking)

	if tableBooking.BookingTime.IsZero() {
		util.ErrorJSON(ctx, http.StatusBadRequest, "BookingTime is required and must be a valid time")
		return
	}
	err := p.service.CreateTableBooking(tableBooking)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create tableBooking")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created TableBooking")
}

func (p TableBookingController) UpdateTableBooking(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var tableBooking models.TableBooking
	tableBooking.ID = id

	tableBookingRecord, err := p.service.FindTableBooking(tableBooking)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "TableBooking with given id not found")
		return
	}
	ctx.ShouldBindJSON(&tableBookingRecord)

	// if tableBooking.BookingTime.IsZero() {
	// 	util.ErrorJSON(ctx, http.StatusBadRequest, "BookingTime is required and must be a valid time")
	// 	return
	// }

	if err := p.service.UpdateTableBooking(tableBookingRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store TableBooking")
		return
	}
	response := tableBookingRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated TableBooking",
		Data:    response,
	})
}


func (p *TableBookingController) DeleteTableBooking(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.DeleteTableBooking(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete TableBooking")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	ctx.JSON(http.StatusOK, response)
}

func (c TableBookingController) GetUserBookingByStatus(ctx *gin.Context) {
	var tableBookings models.TableBooking

	status := ctx.Query("status")
	idParam := ctx.Query("UserId")
	userId, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}

	tableBookings.UserID = userId
	tableBookings.Status = status

	data, total, err := c.service.FindAllUserBookingByStatus(tableBookings)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "TableBooking result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (c TableBookingController) GetStoreBookingByStatus(ctx *gin.Context) {
	var tableBookings models.TableBooking

	status := ctx.Query("status")
	idParam := ctx.Query("StoreId")
	storeId, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}

	tableBookings.StoreID = storeId
	tableBookings.Status = status

	data, total, err := c.service.FindAllUserBookingByStatus(tableBookings)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
		return
	}
	respArr := make([]map[string]interface{}, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "TableBooking result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}