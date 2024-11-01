package v1

import (
	"github.com/gin-gonic/gin"
	"my-blog-golang/server/model"
	"my-blog-golang/server/utils/errmsg"
	"net/http"
	"strconv"
)

// 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	// 验证前端传递过来的分类
	// 分类名不为空，则进行后续的步骤
	if data.Name != "" {
		code = model.CheckCategory(data.Name)
		if code == errmsg.SUCCSE {
			code = model.CreateCate(&data)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		code = errmsg.ERROR_CATE_NOT_NULL
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = 10 // Limit 10
	}
	if pageNum == 0 {
		pageNum = -1 // 不分页
	}
	data, total := model.GetCate(pageSize, pageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类名
func EditCate(c *gin.Context) {
	var data model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name) // 先检查要修改的分类名称在数据库是否存在
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	if code == errmsg.SUCCSE {
		code = model.EditCate(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
