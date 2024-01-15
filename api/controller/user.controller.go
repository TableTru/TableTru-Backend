package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

func (c UserController) GetAllUser(ctx *gin.Context) {
	var users models.User

	keyword := ctx.Query("keyword")

	data, total, err := c.service.FindAllUser(users, keyword)

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
		Message: "User result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

func (p *UserController) GetUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to int64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var user models.User
	user.ID = id
	foundUser, err := p.service.FindUser(user)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding User")
		return
	}
	response := foundUser.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of User",
		Data:    &response})

}

func (p *UserController) AddUser(ctx *gin.Context) {
	var user models.User
	ctx.ShouldBindJSON(&user)

	if user.Username == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}
	err := p.service.CreateUser(user)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create user")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created User")
}

func (p UserController) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	var user models.User
	user.ID = id

	userRecord, err := p.service.FindUser(user)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "User with given id not found")
		return
	}
	ctx.ShouldBindJSON(&userRecord)

	if userRecord.Username == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	if err := p.service.UpdateUser(userRecord); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to store User")
		return
	}
	response := userRecord.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Successfully Updated User",
		Data:    response,
	})
}


func (p *UserController) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64) //type conversion string to uint64
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "id invalid")
		return
	}
	err = p.service.DeleteUser(id)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to delete User")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Deleted Sucessfully"}
	ctx.JSON(http.StatusOK, response)
}

func (p *UserController) GetLoginUser(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	foundUser, err := p.service.FindLoginUser(username, password)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Finding User")
		return
	}
	response := foundUser.ResponseMap()

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of User",
		Data:    &response})

}