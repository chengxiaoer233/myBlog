package dao

import (
	"gorm.io/gorm"
	"myBlog/model"
)

var CommentInterface SqlComment

type SqlComment struct {
}

// AddComment 添加评论
func (c *SqlComment)AddComment(data model.Comment) (err error){
	err = Db.Create(&data).Error
	return
}

// DeleteComment 删除评论
func (c *SqlComment)DeleteComment(id int) (err error){
	err = Db.Where("id = ?",id).Delete(&model.Comment{}).Error
	return
}

// 展示不允许编辑评论

// GetOneComment 查询单个评论
func (c *SqlComment)GetOneComment(id int) (data model.Comment,err error){
	err = Db.Where("id = ?",id).First(&data).Error
	return
}

// GetCommentLists 后台获取所有评论列表
func (c *SqlComment)GetCommentLists(pageSize int,pageNum int)(data *[]model.Comment,total int64,err error) {

	// 获取所有的总数
	err = Db.Find(&data).Count(&total).Error

	// 获取需要分页的数据，按最新的时间排序
	// 三标联合查询，评论表，用户表、文章表
	err = Db.Model(&data).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Select("comment.id, article.title,user_id,article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Scan(&data).Error
	return
}

// GetCommentCount 获取谋篇文章的评论数量
func (c *SqlComment)GetCommentCount(id int)(total int64,err error){
	err = Db.Where("article_id = ? and status = 1",id).Count(&total).Error
	return
}

// GetCommentListFront 前端展示页面获取文章评论列表
func (c *SqlComment)GetCommentListFront(id int, pageSize int, pageNum int) (commentList *[]model.Comment, total int64, err error) {

	err = Db.Find(&model.Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Count(&total).Error

	err = Db.Model(&model.Comment{}).Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Select("comment.id, article.title, user_id, article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Where("article_id = ?",
		id).Where("status = ?", 1).Scan(&commentList).Error

	return
}

// 通过评论
func (c *SqlComment)CheckComment(id int, data *model.Comment) (err error) {

	var res model.Comment
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	// 更新状态
	err = Db.Model(&model.Comment{}).Where("id = ?", id).Updates(maps).First(&res).Error

	// 评论数目加一
	err = Db.Model(&model.Article{}).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error

	return
}

// 撤销评论
func (c *SqlComment)UncheckComment(id int, data *model.Comment) (err error) {

	var res model.Comment
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	// 更新评论状态
	err = Db.Model(&model.Comment{}).Where("id = ?", id).Updates(maps).First(&res).Error

	// 更新最新的数目
	err = Db.Model(&model.Article{}).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error

	return
}