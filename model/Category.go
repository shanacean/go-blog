package model

import (
	"go-blog/utils/result"
	"gorm.io/gorm"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type: varchar(20); not null" json:"name"`
}

func CheckCategory(name string) result.Code {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return result.ERROR_CATEGORY_NAME_USED //2001
	}
	return result.SUCCESS
}

func AddCategory(c *Category) result.Code {
	if err := db.Create(c).Error; err != nil {
		return result.ERROR
	}
	return result.SUCCESS
}

func DeleteCategory(id int) result.Code {
	err := db.Where("id = ?", id).Delete(&Category{}).Error
	if err != nil {
		return result.ERROR
	}
	return result.SUCCESS
}

func EditCategory(id int, c *Category) result.Code {
	err := db.Model(&Category{}).Where("id = ?", id).Update("name", c.Name).Error
	if err != nil {
		return result.ERROR
	}
	return result.SUCCESS
}

func GetCategory(pageSize, pageNum int) ([]Category, int64) {
	var cs []Category
	var total int64
	if err := db.Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&cs).
		Count(&total).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cs, total
}
