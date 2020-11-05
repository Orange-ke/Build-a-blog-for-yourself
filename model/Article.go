package model

import (
	"gorm.io/gorm"
	"myblog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"` // 物理外键
	gorm.Model
	Title   string `gorm:"type: varchar(100); not null" json:"title"`
	Cid     int    `gorm:"type: int; not null" json:"cid"`
	Desc    string `gorm:"type: varchar(100); not null" json:"desc"`
	Content string `gorm:"type: longtext" json:"content"`
	Img     string `gorm:"type: varchar(100)" json:"img"`
}

// 新增文章
func CreateArticle(data *Article) (code int) {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

//查询单个文章
func GetArticle(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

//查询文章列表
func GetArticles(title string, cid int, pageSize int, pageNum int) ([]Article, int, int) {
	//需要添加总数 字段 便于前端分页展示
	var articles []Article
	var total int
	var err error
	if title == "" && cid == 0 {
		err = db.Order("Updated_At DESC").Preload("Category").Find(&articles).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	} else if title != "" && cid == 0 {
		err = db.Order("Update_At DESC").Preload("Category").Where("title LIKE ?", title + "%").Find(&articles).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	} else if title == "" && cid != 0 {
		err = db.Order("Updated_At DESC").Preload("Category").Where("cid = ?", cid).Find(&articles).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	} else {
		err = db.Order("Updated_At DESC").Preload("Category").Where("cid = ? AND title LIKE ?", cid, title + "%").Find(&articles).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articles, errmsg.SUCCESS, total
}

// 编辑文章
func EditArticle(id int, data *Article) int {
	// gorm使用struct更新时，零值默认不会更新，因此使用map来更新
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&article).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章(软删除)
func DeleteArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
