package dao

import (
	"errors"
	"gorm.io/gorm"
	"myBlog/model"
)

var ArticleInterface SqlArticle
type SqlArticle struct{

}

// CheckUser 查询文章是否存在
func (a *SqlArticle) CheckArticleById(id int) (int, error) {

	var art []model.Article
	err := Db.Select("id").Where("id=?", id).Find(&art).Error
	if err != nil {
		return model.ErrInner, err
	}

	if len(art) > 0 {
		return model.ErrArtExists, errors.New(model.GetErrMsg(model.ErrArtExists))
	}

	return model.Success, nil
}

// CreateArt 新增文章
func (a *SqlArticle) CreateArt(data *model.Article) (err error) {
	err = Db.Create(&data).Error
	return
}

// DeleteArt 删除文章
func (a *SqlArticle) DeleteArt(id int)(err error){
	var data model.Article
	err = Db.Where("id = ?",id).Unscoped().Delete(&data).Error
	return
}

// EditArticle 编辑文章
func (a *SqlArticle)EditArticle(id int,data *model.Article)(err error){

	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = Db.Model(data).Where("id = ?",id).Updates(&maps).Error
	return
}

// 查询返回单个文章
// GetArtInfo 查询单个文章
func (a *SqlArticle)GetArtInfo(id int) (art model.Article, err error) {

	err = Db.Where("id = ?", id).Preload("Category").First(&art).Error

	err = Db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1)).Error

	return
}

// 查询所有文章
// GetArt 查询文章列表
func (a *SqlArticle)GetAllArts(pageSize int, pageNum int) (articleList *[]model.Article, total int64, err error) {

	err = Db.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("Category").Find(&articleList).Error

	// 单独计数
	err = Db.Model(&articleList).Count(&total).Error

	return
}

// SearchArticle 搜索文章标题
func (a *SqlArticle)SearchArticle(title string, pageSize int, pageNum int) (articleList *[]model.Article, total int64, err error) {

	err = Db.Select("article.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, Category.name").
		Order("Created_At DESC").Joins("Category").
		Where("title LIKE ?", title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error

	//单独计数
	err = Db.Model(&articleList).Where("title LIKE ?",title+"%").Count(&total).Error

	return
}

// GetCateArt 查询分类下的所有文章
func (a *SqlArticle) GetCateArts(id int, pageSize int, pageNum int) (cateArtList *[]model.Article, total int64, err error) {

	err = Db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid =?", id).Find(&cateArtList).Error

	err = Db.Model(&cateArtList).Where("cid =?", id).Count(&total).Error

	return
}