package helpers

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"
)

func SaveFile(c *gin.Context, file *multipart.FileHeader, path string) string {
	prefix := strconv.Itoa(rand.Int()) + "_"

	filename := prefix + strconv.FormatInt(time.Now().Unix(), 10) + filepath.Ext(file.Filename)

	dst := path + filename
	err := c.SaveUploadedFile(file, dst)
	HandleError(err)

	return filename
}
