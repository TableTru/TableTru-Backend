package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type StoreController struct {
	service service.StoreService
}

func NewStoreController(s service.StoreService) StoreController {
	return StoreController{
		service: s,
	}
}

func (c StoreController) GetAllStore(ctx *gin.Context) {
	var stores models.Store

	keyword := ctx.Query("keyword")

	data, total, err := c.service.FindAllStore(stores, keyword)

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
		Message: "Store result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *StoreController) GetStore(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var store models.Store
	store.ID = id
	foundPost, err := p.service.FindStore(store)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding Store")
		return
	}
	response := foundPost.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Store",
		Data:    &response})

}

func (p *StoreController) AddStore(ctx *gin.Context) {
	var store models.Store
	ctx.ShouldBindJSON(&store)

	if store.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}
	err := p.service.CreateStore(store)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create store")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Store")
}

func (p StoreController) UpdateStore(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var store models.Store
	store.ID = id

	storeRecord, err := p.service.FindStore(store)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Store with given id not found")
		return
	}
	ctx.ShouldBindJSON(&storeRecord)

	if storeRecord.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	if err := p.service.UpdateStore(storeRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store Store")
		return
	}
	response := storeRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Store",
		Data:    response,
	})
}


func (p *StoreController) DeleteStore(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.DeleteStore(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete Store")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	ctx.JSON(http.StatusOK, response)
}