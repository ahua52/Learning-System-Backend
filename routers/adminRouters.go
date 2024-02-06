package routers

import (
	"HMS/payment/controllers/signIn"
	// "HMS/payment/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/")
	//middlewares.JWTAuth
	{
		// adminRouters.GET("/api/signin", admin.BasePaymentControl{}.SignIn)
		adminRouters.POST("/auth/signin", signIn.BaseSignInControl{}.GetAuth)
		// adminRouters.GET("/auth/signin", signIn.BaseSignInControl{}.SignIn)
	}
}
