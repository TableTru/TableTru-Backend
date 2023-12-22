package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewController struct {
	service service.ReviewService
}

func NewReviewController(s service.ReviewService) ReviewController {
	return ReviewController{
		service: s,
	}
}

func (c ReviewController) GetAllReview(ctx *gin.Context) {
	var reviews models.Review

	keyword := ctx.Query("keyword")

	data, total, err := c.service.FindAllReview(reviews, keyword)

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
		Message: "Review result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *ReviewController) GetReview(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var review models.Review
	review.ID = id
	foundReview, err := p.service.FindReview(review)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding Review")
		return
	}
	response := foundReview.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Review",
		Data:    &response})

}

func (p *ReviewController) AddReview(ctx *gin.Context) {
	var review models.Review
	ctx.ShouldBindJSON(&review)

	if review.Comment == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Comment is required")
		return
	}
	err := p.service.CreateReview(review)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create review")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Review")
}

func (p ReviewController) UpdateReview(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var review models.Review
	review.ID = id

	reviewRecord, err := p.service.FindReview(review)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Review with given id not found")
		return
	}
	ctx.ShouldBindJSON(&reviewRecord)

	if reviewRecord.Comment == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Comment is required")
		return
	}

	if err := p.service.UpdateReview(reviewRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Review")
		return
	}
	response := reviewRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Review",
		Data:    response,
	})
}


func (p *ReviewController) DeleteReview(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.DeleteReview(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete Review")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	ctx.JSON(http.StatusOK, response)
}