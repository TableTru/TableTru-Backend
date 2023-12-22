package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service service.CategoryService
}

func NewCategoryController(s service.CategoryService) CategoryController {
	return CategoryController{
		service: s,
	}
}

func (c CategoryController) GetAllCategory(ctx *gin.Context) {
	var categories models.Category

	keyword := ctx.Query("keyword")

	data, total, err := c.service.FindAllCategory(categories, keyword)

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
		Message: "Category result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *CategoryController) GetCategory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var category models.Category
	category.ID = id
	foundCategory, err := p.service.FindCategory(category)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding Category")
		return
	}
	response := foundCategory.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Category",
		Data:    &response})

}

func (p *CategoryController) AddCategory(ctx *gin.Context) {
	var category models.Category
	ctx.ShouldBindJSON(&category)

	if category.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}
	err := p.service.CreateCategory(category)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create category")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Category")
}

func (p CategoryController) UpdateCategory(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var category models.Category
	category.ID = id

	categoryRecord, err := p.service.FindCategory(category)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Category with given id not found")
		return
	}
	ctx.ShouldBindJSON(&categoryRecord)

	if categoryRecord.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	if err := p.service.UpdateCategory(categoryRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Category")
		return
	}
	response := categoryRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Category",
		Data:    response,
	})
}


func (p *CategoryController) DeleteCategory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.DeleteCategory(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete Category")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	ctx.JSON(http.StatusOK, response)
}