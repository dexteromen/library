package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	authRoutes(router)    //Working
	libraryRoutes(router) //Working
	bookRoutes(router)    //Working
	issueApprovalRoutes(router)
}
