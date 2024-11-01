package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-blog-golang/server/model"
	"my-blog-golang/server/utils"
	"my-blog-golang/server/utils/errmsg"
	"net/http"
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
