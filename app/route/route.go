package route

import (
	"bip/app/controller"
	"github.com/gin-gonic/gin"
)

func CollerRoute(r *gin.Engine) *gin.Engine {
	r.POST("/bip",controller.UserRegister)
	return r
}
