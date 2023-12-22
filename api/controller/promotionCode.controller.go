package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type PromotionCodeController struct {
	service service.PromotionCodeService
}

func NewPromotionCodeController(s service.PromotionCodeService) PromotionCodeController {
	return PromotionCodeController{
		service: s,
	}
}

func (c PromotionCodeController) GetAllPromotionCode(ctx *gin.Context) {
	var promotionCodes models.PromotionCode

	keyword := ctx.Query("keyword")

	data, total, err := c.service.FindAllPromotionCode(promotionCodes, keyword)

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
		Message: "PromotionCode result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *PromotionCodeController) GetPromotionCode(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var promotionCode models.PromotionCode
	promotionCode.ID = id
	foundPromotionCode, err := p.service.FindPromotionCode(promotionCode)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding PromotionCode")
		return
	}
	response := foundPromotionCode.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of PromotionCode",
		Data:    &response})

}

func (p *PromotionCodeController) AddPromotionCode(ctx *gin.Context) {
	var promotionCode models.PromotionCode
	ctx.ShouldBindJSON(&promotionCode)

	if promotionCode.Status == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Status is required")
		return
	}
	err := p.service.CreatePromotionCode(promotionCode)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create promotionCode")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created PromotionCode")
}

func (p PromotionCodeController) UpdatePromotionCode(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var promotionCode models.PromotionCode
	promotionCode.ID = id

	promotionCodeRecord, err := p.service.FindPromotionCode(promotionCode)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "PromotionCode with given id not found")
		return
	}
	ctx.ShouldBindJSON(&promotionCodeRecord)

	if promotionCodeRecord.Status == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Status is required")
		return
	}

	if err := p.service.UpdatePromotionCode(promotionCodeRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store PromotionCode")
		return
	}
	response := promotionCodeRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated PromotionCode",
		Data:    response,
	})
}


func (p *PromotionCodeController) DeletePromotionCode(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.DeletePromotionCode(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete PromotionCode")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	ctx.JSON(http.StatusOK, response)
}