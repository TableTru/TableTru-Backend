package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type RatingController struct {
	service service.RatingService
}

func NewRatingController(s service.RatingService) RatingController {
	return RatingController{
		service: s,
	}
}

func (c RatingController) GetAllRating(ctx *gin.Context) {
	var ratings models.Rating

	keyword := ctx.Query("keyword")

	data, total, err := c.service.FindAllRating(ratings, keyword)

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
		Message: "Rating result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *RatingController) GetRating(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var rating models.Rating
	rating.ID = id
	foundRating, err := p.service.FindRating(rating)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding Rating")
		return
	}
	response := foundRating.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Rating",
		Data:    &response})

}

func (p *RatingController) AddRating(ctx *gin.Context) {
	var rating models.Rating
	ctx.ShouldBindJSON(&rating)

	if rating.Score < 1 || rating.Score > 5 {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Score must be in the range of 1-5.")
		return
	}
	err := p.service.CreateRating(rating)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create rating")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Rating")
}

func (p RatingController) UpdateRating(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var rating models.Rating
	rating.ID = id

	ratingRecord, err := p.service.FindRating(rating)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Rating with given id not found")
		return
	}
	ctx.ShouldBindJSON(&ratingRecord)

	if rating.Score < 1 || rating.Score > 5 {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Score must be in the range of 1-5.")
		return
	}

	if err := p.service.UpdateRating(ratingRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Rating")
		return
	}
	response := ratingRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Rating",
		Data:    response,
	})
}


func (p *RatingController) DeleteRating(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.DeleteRating(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete Rating")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	ctx.JSON(http.StatusOK, response)
}