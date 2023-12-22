package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewImageController struct {
	service service.ReviewImageService
}

func NewReviewImageController(s service.ReviewImageService) ReviewImageController {
	return ReviewImageController{
		service: s,
	}
}

func (c ReviewImageController) GetAllReviewImage(ctx *gin.Context) {
	var reviewImages models.ReviewImage

	keyword := ctx.Query("keyword")

	data, total, err := c.service.FindAllReviewImage(reviewImages, keyword)

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
		Message: "ReviewImage result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *ReviewImageController) GetReviewImage(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var reviewImage models.ReviewImage
	reviewImage.ID = id
	foundReviewImage, err := p.service.FindReviewImage(reviewImage)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding ReviewImage")
		return
	}
	response := foundReviewImage.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of ReviewImage",
		Data:    &response})

}

func (p *ReviewImageController) AddReviewImage(ctx *gin.Context) {
	var reviewImage models.ReviewImage
	ctx.ShouldBindJSON(&reviewImage)

	if reviewImage.ImageName == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "ImageName is required")
		return
	}
	err := p.service.CreateReviewImage(reviewImage)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create reviewImage")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created ReviewImage")
}

func (p ReviewImageController) UpdateReviewImage(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var reviewImage models.ReviewImage
	reviewImage.ID = id

	reviewImageRecord, err := p.service.FindReviewImage(reviewImage)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "ReviewImage with given id not found")
		return
	}
	ctx.ShouldBindJSON(&reviewImageRecord)

	if reviewImageRecord.ImageName == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "ImageName is required")
		return
	}

	if err := p.service.UpdateReviewImage(reviewImageRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store ReviewImage")
		return
	}
	response := reviewImageRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated ReviewImage",
		Data:    response,
	})
}


func (p *ReviewImageController) DeleteReviewImage(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.DeleteReviewImage(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete ReviewImage")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	ctx.JSON(http.StatusOK, response)
}