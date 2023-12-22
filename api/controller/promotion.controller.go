package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type PromotionController struct {
	service service.PromotionService
}

func NewPromotionController(s service.PromotionService) PromotionController {
	return PromotionController{
		service: s,
	}
}

func (c PromotionController) GetAllPromotion(ctx *gin.Context) {
	var promotions models.Promotion

	keyword := ctx.Query("keyword")

	data, total, err := c.service.FindAllPromotion(promotions, keyword)

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
		Message: "Promotion result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *PromotionController) GetPromotion(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var promotion models.Promotion
	promotion.ID = id
	foundPromotion, err := p.service.FindPromotion(promotion)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding Promotion")
		return
	}
	response := foundPromotion.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Promotion",
		Data:    &response})

}

func (p *PromotionController) AddPromotion(ctx *gin.Context) {
	var promotion models.Promotion
	ctx.ShouldBindJSON(&promotion)

	if promotion.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}
	err := p.service.CreatePromotion(promotion)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create promotion")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Promotion")
}

func (p PromotionController) UpdatePromotion(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var promotion models.Promotion
	promotion.ID = id

	promotionRecord, err := p.service.FindPromotion(promotion)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Promotion with given id not found")
		return
	}
	ctx.ShouldBindJSON(&promotionRecord)

	if promotionRecord.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	if err := p.service.UpdatePromotion(promotionRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Promotion")
		return
	}
	response := promotionRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Promotion",
		Data:    response,
	})
}


func (p *PromotionController) DeletePromotion(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.DeletePromotion(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete Promotion")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	ctx.JSON(http.StatusOK, response)
}