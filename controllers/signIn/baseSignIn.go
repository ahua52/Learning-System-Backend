package signIn

import (
	"HMS/payment/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseSignInControl struct {
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (con BaseSignInControl) SignIn(ctx *gin.Context) {
	json := User{}
	ctx.ShouldBind(&json)

	var user *models.User
	var routePermition *models.Permission
	var userrole *models.Role

	fmt.Println(user, routePermition, userrole)
	ctx.JSON(http.StatusOK, gin.H{
		"message":            "success",
		"paymentReferenceId": ctx.Request,
	})
}
