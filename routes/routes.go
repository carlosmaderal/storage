package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/pingchk", PingHandler)
	r.POST("/upload", UploadFileHandler)
	r.GET("/storage/*filePath", GetfileHandler) 
}
