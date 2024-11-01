package v1

import (
	"go-blog-server/model"
	"go-blog-server/utils/errmsg"
	"go-blog-server/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

// 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	_ = c.ShouldBindJSON(&data)

	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}

	code = model.CheckUser(data.Username)
	if code == errmsg.ERROR_USERNAME_USED {
		// 用户已经存在
		code = errmsg.ERROR_USERNAME_USED
	}
	if code == errmsg.SUCCSE {
		// 新增用户
		code = model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		//"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")
	if pageSize == 0 {
		pageSize = 10 // Limit 10
	}
	if pageNum == 0 {
		pageNum = -1 // 不分页
	}
	data, total := model.GetUsers(username, pageSize, pageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	// 校验前端传递过来的字段
	var msg string
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}

	// 先检查要修改的用户名称在数据库是否存在
	code = model.EditCheckUser(id, data.Username)
	if code == errmsg.SUCCSE {
		code = model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
