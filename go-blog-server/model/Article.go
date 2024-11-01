package model

import (
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
	"my-blog-golang/server/utils/errmsg"
)

// 文章详细

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"` // 文章标题
	Cid     int    `gorm:"type:int;not null" json:"cid"`            //
	Desc    string `gorm:"type:varchar(200)" json:"desc"`           // 说明
	Content string `gorm:"type:longtext" json:"content"`            // 文章类容
	Img     string `gorm:"type:varchar(100)" json:"img"`            // 图片
}

// 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询分类下的所有文章
func GetCateArt(id int, pageSize, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	// 预加载 https://v1.gorm.io/zh_CN/docs/preload.html
	db.Preload("Category").Where("cid = ?", id).Find(&cateArtList).Count(&total)
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtList).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCSE, total
}

// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCSE
}

// 查询文章列表
func GetArt(title string, pageSize, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	if title != "" {
		db.Preload("Category").Where("title LIKE ?", "%"+title+"%").Find(&articleList).Count(&total)
		db.Order("Updated_At DESC").Preload("Category").Where("title LIKE ?", "%"+title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR, 0
		}
	} else {
		db.Find(&articleList).Count(&total)
		err = db.Order("Updated_At DESC").Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR, 0
		}
	}
	return articleList, errmsg.SUCCSE, total
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArt(id int) int {
	var art Article
	err = db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
