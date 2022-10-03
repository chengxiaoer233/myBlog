package dao

import (
	"errors"
	"myBlog/model"
)

var ArticleInterface SqlArticle
type SqlArticle struct{

}

// CheckUser 查询分类是否存在
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