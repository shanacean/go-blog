package model

import (
	"go-blog/utils/result"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string   `gorm:"type: varchar(100); not null" json:"title"`
	Category    Category `gorm:"foreignKey:Cid;reference:ID" json:"category,omitempty"`
	Cid         int      `gorm:"type: int; not nul" json:"cid"`
	Description string   `gorm:"type: varchar(200)" json:"description"`
	Content     string   `gorm:"type: longtext" json:"content"`
	Img         string   `gorm:"type: varchar(100)" json:"img"`
}

// AddArticle 添加文章
func AddArticle(article *Article) result.Code {
	if err := db.Create(article).Error; err != nil {
		return result.ERROR
	}
	return result.SUCCESS
}

// DeleteArticle 删除文章
func DeleteArticle(id int) result.Code {
	err := db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return result.ERROR
	}
	return result.SUCCESS
}

// EditArticle 编辑文章
func EditArticle(id int, c *Article) result.Code {
	var maps = make(map[string]interface{})
	maps["title"] = c.Title
	maps["cid"] = c.Cid
	maps["description"] = c.Description
	maps["content"] = c.Content
	maps["img"] = c.Img
	err := db.Model(&Article{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return result.ERROR
	}
	return result.SUCCESS
}

// GetSingleArticle 查询单个文章信息
func GetSingleArticle(id int) (Article, result.Code) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, result.ERROR_ARTICLE_NOT_EXIST
	}
	return article, result.SUCCESS
}

// GetArticles 分页查询所有文章
func GetArticles(pageSize, pageNum int) ([]Article, int64, result.Code) {
	var articles []Article
	var total int64
	if err := db.Preload("Category").
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&articles).Count(&total).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, result.ERROR
	}
	return articles, total, result.SUCCESS
}

func GetArticlesByCategory(cid, pageSize, pageNum int) ([]Article, int64, result.Code) {
	var articles []Article
	var total int64
	err := db.Preload("Category").
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Where("cid = ?", cid).
		Find(&articles).
		Count(&total)
	if err != nil {
		return articles, total, result.ERROR_CATEGORY_NOT_EXIST
	}
	return articles, total, result.SUCCESS
}

