package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type StoreImageController struct {
	service service.StoreImageService
}

func NewStoreImageController(s service.StoreImageService) StoreImageController {
	return StoreImageController{
		service: s,
	}
}

func (c StoreImageController) GetAllStoreImage(ctx *gin.Context) {
	var storeImages models.StoreImage

	keyword := ctx.Query("keyword")

	data, total, err := c.service.FindAllStoreImage(storeImages, keyword)

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
		Message: "StoreImage result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *StoreImageController) GetStoreImage(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var storeImage models.StoreImage
	storeImage.ID = id
	foundStoreImage, err := p.service.FindStoreImage(storeImage)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding StoreImage")
		return
	}
	response := foundStoreImage.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of StoreImage",
		Data:    &response})

}

func (p *StoreImageController) AddStoreImage(ctx *gin.Context) {
	var storeImage models.StoreImage
	ctx.ShouldBindJSON(&storeImage)

	if storeImage.ImageName == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "ImageName is required")
		return
	}
	err := p.service.CreateStoreImage(storeImage)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create storeImage")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created StoreImage")
}

func (p StoreImageController) UpdateStoreImage(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var storeImage models.StoreImage
	storeImage.ID = id

	storeImageRecord, err := p.service.FindStoreImage(storeImage)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "StoreImage with given id not found")
		return
	}
	ctx.ShouldBindJSON(&storeImageRecord)

	if storeImageRecord.ImageName == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "ImageName is required")
		return
	}

	if err := p.service.UpdateStoreImage(storeImageRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store StoreImage")
		return
	}
	response := storeImageRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated StoreImage",
		Data:    response,
	})
}


func (p *StoreImageController) DeleteStoreImage(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.DeleteStoreImage(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete StoreImage")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	ctx.JSON(http.StatusOK, response)
}