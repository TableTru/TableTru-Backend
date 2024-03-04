package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type OpenTimeController struct {
	service service.OpenTimeService
}

func NewOpenTimeController(s service.OpenTimeService) OpenTimeController {
	return OpenTimeController{
		service: s,
	}
}

func (c OpenTimeController) GetAllOpenTime(ctx *gin.Context) {
	var opentimes models.OpenTime

	keyword := ctx.Query("keyword")

	data, total, err := c.service.FindAllOpenTime(opentimes, keyword)

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
		Message: "OpenTime result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *OpenTimeController) GetOpenTime(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var opentime models.OpenTime
	opentime.ID = id
	foundOpenTime, err := p.service.FindOpenTime(opentime)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding OpenTime")
		return
	}
	response := foundOpenTime.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of OpenTime",
		Data:    &response})

}

func (p *OpenTimeController) AddOpenTime(ctx *gin.Context) {
	var opentime models.OpenTime
	ctx.ShouldBindJSON(&opentime)

	if opentime.Day == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}
	err := p.service.CreateOpenTime(opentime)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create opentime")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created OpenTime")
}

func (p OpenTimeController) UpdateOpenTime(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var opentime models.OpenTime
	opentime.ID = id

	opentimeRecord, err := p.service.FindOpenTime(opentime)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "OpenTime with given id not found")
		return
	}
	ctx.ShouldBindJSON(&opentimeRecord)

	if opentimeRecord.Day == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	if err := p.service.UpdateOpenTime(opentimeRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store OpenTime")
		return
	}
	response := opentimeRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated OpenTime",
		Data:    response,
	})
}


func (p *OpenTimeController) DeleteOpenTime(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.DeleteOpenTime(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete OpenTime")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	ctx.JSON(http.StatusOK, response)
}