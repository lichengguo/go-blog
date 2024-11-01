package model

import (
	"fmt"
	"go-blog-server/utils"
	"go-blog-server/utils/errmsg"

	"github.com/gin-gonic/gin"

	"mime/multipart"
)

// 保存文件
func UpLoadFile(c *gin.Context, file *multipart.FileHeader) (string, int) {
	// 文件名称md5加密,避免重名
	fileName := utils.Md5Str(file.Filename)
	// 拼接文件存储路径
	fileSavePath := fmt.Sprintf("%s/%s", utils.ImgPath, fileName)
	// 存储文件
	err := c.SaveUploadedFile(file, fileSavePath)
	if err != nil {
		return "", errmsg.ERROR
	}

	return fileSavePath, errmsg.SUCCSE
}
