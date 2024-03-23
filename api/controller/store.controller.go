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
	foundStore, err := p.service.FindStore(store)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding Store")
		return
	}
	response := foundStore.ResponseMap()

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

	storeRecord, err := p.service.FindStore(store)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Store with given id not found")
		return
	}
	ctx.ShouldBindJSON(&storeRecord)

	response := storeRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated Store",
		Data:    response,
	})
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

func (p *StoreController) CheckStoreByName(ctx *gin.Context) {
	storeName := ctx.Query("storeName")

	var store models.Store
	store.Name = storeName
	foundStore, err := p.service.FindStore(store)
	if err != nil {
		ctx.JSON(http.StatusOK, &util.Response{
			Success: false,
			Message: "not found",
		})
	} else {
		response := foundStore.ResponseMap()

		ctx.JSON(http.StatusOK, &util.Response{
			Success: true,
			Message: "Result set of Store",
			Data:    &response,
		})
	}
}

func (c StoreController) GetStoreByNum(ctx *gin.Context) {
	var stores models.Store

	keyword := ctx.Query("keyword")
	numString := ctx.Query("num")
	num, err := strconv.Atoi(numString)
	if err != nil {
		// กรณีเกิด error ในการแปลง
	}

	data, total, err := c.service.FindStoreByNum(stores, keyword, num)

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

func (c StoreController) SearchStoreSortRating(ctx *gin.Context) {
	var searchInput models.SearchInput
	var stores models.Store

	// keyword := ctx.Query("keyword")
	ctx.ShouldBindJSON(&searchInput)

	keyword := searchInput.Search

	respArr := make([]map[string]interface{}, 0)

	//เรียงตาม rating
	if searchInput.CategoryID != 0 {
		stores.CategoryID = searchInput.CategoryID
		data, total, err := c.service.SearchStoreRatingSort(stores, keyword)
		if err != nil {
			util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
			return
		}

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
	} else {
		data, total, err := c.service.FindAllStore(stores, keyword)
		if err != nil {
			util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
			return
		}

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

}

func (c StoreController) SearchStoreSortDistance(ctx *gin.Context) {
	var searchInput models.SearchInput
	var stores models.Store

	// keyword := ctx.Query("keyword")
	ctx.ShouldBindJSON(&searchInput)

	keyword := searchInput.Search

	respArr := make([]map[string]interface{}, 0)

	//เรียงตาม rating
	if searchInput.CategoryID != 0 {
		stores.CategoryID = searchInput.CategoryID
		data, total, err := c.service.SearchStoreRatingSort(stores, keyword)
		if err != nil {
			util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
			return
		}

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
	} else {
		data, total, err := c.service.FindAllStore(stores, keyword)
		if err != nil {
			util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find questions")
			return
		}

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

}
