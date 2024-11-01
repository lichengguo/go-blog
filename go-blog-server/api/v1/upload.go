package v1

import (
	"fmt"
	"go-blog-server/model"
	"go-blog-server/utils"
	"go-blog-server/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpLoad(c *gin.Context) {
	file, _ := c.FormFile("file")
	urlPath, code := model.UpLoadFile(c, file)
	// 拼接URL
	url := fmt.Sprintf("%s%s/%s", "http://", utils.HttpPort, urlPath)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
