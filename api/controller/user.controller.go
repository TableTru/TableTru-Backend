package controller

import (
	"TableTru/api/service"
	"TableTru/models"
	"TableTru/util"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
)

// PostController -> PostController
type UserController struct {
	service service.UserService
}

// NewPostController : NewPostController
func NewUserController(s service.UserService) UserController {
	return UserController{
		service: s,
	}
}

func (u UserController) GetAllUser(ctx *gin.Context) {
	var users models.User

	keyword := ctx.Query("keyword")

	data, total, err := u.service.FindAllUser(users, keyword)

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

// AddPost : AddPost controller
func (u *UserController) AddUser(ctx *gin.Context) {
	var user models.User
	ctx.ShouldBindJSON(&user)

	if user.Username == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}
	err := u.service.CreateUser(user)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create user")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created User")
}
