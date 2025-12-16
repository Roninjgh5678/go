package main

import (
	"golangproject/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("HTML/*")
	routers.UserRoutersInit(r)
	r.Run(":8081")
}
