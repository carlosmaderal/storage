package routes

import (
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"storage/models"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": time.Now().Unix(),
		"host":   os.Getenv("SERVER_HOST") + "/storage/",
	})
}

func UploadFileHandler(c *gin.Context) {
	var filename string
	var err error

	switch {
	case c.Query("download") == "1" && c.Query("file") != "":
		filename, err = models.StoreFileFromURL(c.Query("file"))
	case c.Query("fromurl") == "1" && c.Query("file") != "":
		filename, err = models.CopyFileFromURL(c.Query("file"))
	case c.Query("base64") == "1" && c.Query("file") != "":
		filename, err = models.StoreBase64File(c.Query("file"))
	case c.Query("getfile") == "1" && c.Query("file") != "":
		filename, err = models.StoreDirectFile(c.Query("file"))
	default:
		file, _ := c.FormFile("file")
		filename, err = models.StoreUploadedFile(file)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": os.Getenv("SERVER_HOST") + filename})
}
