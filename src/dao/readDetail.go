package dao

import (
	"fmt"
	"myBlog/model"
	"time"
)

var DataInterface SqlData

type SqlData struct {
}

// 查询7天的用户访问记录
func (d *SqlData) VisitsInfos(day int) (resp []model.Readdetail, err error) {

	times := fmt.Sprintf("-%dh", day*24)
	m, _ := time.ParseDuration(times)
	startTime := time.Now().Add(m)

	err = Db.Where("updated_at between ? and ? ", startTime, time.Now()).Find(&resp).Error

	return
}

// 查询站点总访问次数（每个文章加起来之和）
func (d *SqlData) TotalCount() (total int, err error) {

	var resp []model.Article
	err = Db.Find(&resp).Error

	for i := 0; i < len(resp); i++ {
		total += resp[i].ReadCount
	}

	return
}

// 新增文章。需要插入一条记录，文章ID需要关联起来
func (d *SqlData) CreateRecord(data *model.Readdetail) (err error) {
	err = Db.Create(&data).Error
	return
}
