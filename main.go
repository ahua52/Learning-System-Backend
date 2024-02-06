package main

import (
	"HMS/payment/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// routers.GenerateToken()

	routers.AdminRoutersInit(r)
	// r.Run() // listen and serve on 0.0.0.0:8080
	r.Run(":9091")
}
