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
			"paymentReferenceId": "ServerError",
		})
		// appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":            "error",
			"paymentReferenceId": "Unauthorized",
		})
		// appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(Username, Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":            "error",
			"paymentReferenceId": "ServerError",
		})
		// appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	user, _ := getUserData(Username)
	fmt.Println(1111, user)
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "success",
		"data": gin.H{
			"accessToken": token,
			"user":        user,
		},
	})
}

func getUserData(username string) ([]models.Role, error) {
	user, err := models.GetUser(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
