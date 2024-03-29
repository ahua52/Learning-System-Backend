package signIn

import (
	"fmt"
	"net/http"

	"HMS/payment/models"
	"HMS/payment/pkg/util"
	"HMS/payment/service/auth_service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

type Info struct {
	models.User
	permissions []models.Permission
	// Permissions []Permission
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [POST]
func (con BaseSignInControl) GetAuth(c *gin.Context) {
	valid := validation.Validation{}
	json := auth{}
	c.ShouldBind(&json)
	fmt.Println(1111, json)
	Username, Password := json.Username, json.Password

	a := auth{Username: Username, Password: Password}
	ok, _ := valid.Valid(&a)

	if !ok {
		// app.MarkErrors(valid.Errors)
		c.JSON(http.StatusBadRequest, gin.H{
			"message":            "error",
			"status":             -1,
			"paymentReferenceId": "StatusBadRequest",
		})
		// appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: Username, Password: Password}
	isExist, err := authService.Check()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":            "error",
			"status":             -1,
			"paymentReferenceId": "ServerError",
		})
		// appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":            "error",
			"status":             -1,
			"paymentReferenceId": "Unauthorized",
		})
		// appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(Username, Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":            "error",
			"status":             -1,
			"paymentReferenceId": "ServerError",
		})
		// appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	user, role, err := models.GetUser(Username)
	user.Role.Permission = role.Permission
	//to-do: permissions as a field of user

	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"status":  0,
		"data": gin.H{
			"accessToken": token,
			"user":        user,
		},
	})
}

// type ReturnUser struct {
// 	User
// 	// permissions []Permission
// 	Permissions []models.Permission `json:"permissions"`
// }

// func getUserData(username string) (*models.User, *models.Role, error) {
// 	user, role, err := models.GetUser(username)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	return user, role, nil
// }
