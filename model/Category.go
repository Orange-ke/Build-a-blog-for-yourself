package model

import (
	"github.com/jinzhu/gorm"
	"myblog/utils/errmsg"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20); not null" json:"name"`
}

// 数据库操作
// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED // 2001
	}
	return errmsg.SUCCESS
}

// 新增分类
func CreateCategory(data *Category) (code int) {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

// 查询单个分类
func GetCateInfo(id int) (Category, int) {
	var category Category
	err= db.Where("ID = ?", id).First(&category).Error
	if err != nil {
		return category, errmsg.ERROR
	}
	return category, errmsg.SUCCESS
}

// 查询分类列表
func GetCategories(name string, pageSize int, pageNum int) ([]Category, int) {
	var categories []Category
	var total int
	if name == "" {
		err := db.Find(&categories).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0
		}
		return categories, total
	}
	err := db.Where("name = ?", name).Find(&categories).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return categories, total
}

// 编辑分类 (修改密码单独设定，需要验证)
func EditCategory(id int, data *Category) int {
	// gorm使用struct更新时，零值默认不会更新，因此使用map来更新
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&category).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除分类
func DeleteCategory(id int) int {
	var category Category
	err := db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
