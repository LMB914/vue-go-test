package main

import (
	"bip/app/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := route.CollerRoute(gin.Default())
	panic(r.Run())
}











