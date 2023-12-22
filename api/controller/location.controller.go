package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type LocationController struct {
	service service.LocationService
}

func NewLocationController(s service.LocationService) LocationController {
	return LocationController{
		service: s,
	}
}

func (c LocationController) GetAllLocation(ctx *gin.Context) {
	var locations models.Location

	keyword := ctx.Query("keyword")

	data, total, err := c.service.FindAllLocation(locations, keyword)

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
		Message: "Location result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *LocationController) GetLocation(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var location models.Location
	location.ID = id
	foundLocation, err := p.service.FindLocation(location)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding Location")
		return
	}
	response := foundLocation.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Location",
		Data:    &response})

}

func (p *LocationController) AddLocation(ctx *gin.Context) {
	var location models.Location
	ctx.ShouldBindJSON(&location)

	if location.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}
	err := p.service.CreateLocation(location)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create location")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Location")
}

func (p LocationController) UpdateLocation(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var location models.Location
	location.ID = id

	locationRecord, err := p.service.FindLocation(location)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Location with given id not found")
		return
	}
	ctx.ShouldBindJSON(&locationRecord)

	if locationRecord.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	if err := p.service.UpdateLocation(locationRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Location")
		return
	}
	response := locationRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Location",
		Data:    response,
	})
}


func (p *LocationController) DeleteLocation(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.DeleteLocation(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete Location")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	ctx.JSON(http.StatusOK, response)
}