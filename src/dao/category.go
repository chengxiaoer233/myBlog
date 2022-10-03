package dao

import (
	"errors"
	"myBlog/model"
)

var CategoryInterface SqlCategory
type SqlCategory struct {

}

// CheckUser 查询分类是否存在
func (c *SqlCategory) CheckCategoryByName(name string) (int, error) {

	var cate []model.Category
	err := Db.Select("id").Where("name=?", name).Find(&cate).Error
	if err != nil {
		return model.ErrInner, err
	}

	if len(cate) > 0 {
		return model.ErrCateNameUsed, errors.New(model.GetErrMsg(model.ErrCateNameUsed))
	}

	return model.Success, nil
}

// CheckUser 查询分类是否存在
func (c *SqlCategory) CheckCategoryById(id int) (int, error) {

	var cate []model.Category
	err := Db.Select("id").Where("id=?", id).Find(&cate).Error
	if err != nil {
		return model.ErrInner, err
	}

	if len(cate) > 0 {
		return model.ErrCateNameUsed, errors.New(model.GetErrMsg(model.ErrCateNameUsed))
	}

	return model.Success, nil
}


// CreateUser 新增分类
func (c *SqlCategory) CreateCategory(data *model.Category) (int, error) {

	err := Db.Create(&data).Error
	if err != nil {
		return model.Error, err
	}

	return model.Success, nil
}

// DeleteCategory 删除一个分类，硬删除
func (c *SqlCategory) DeleteCategory(id int)(err error){

	var cate model.Category
	err = Db.Where("id =?",id).Unscoped().Delete(&cate).Error
	return
}

// EditCategory 修改一个分类
func (c *SqlCategory) EditCategory(id int,data *model.Category)(code int , err error){

	var cate model.Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = Db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return model.ErrInner,err
	}

	return model.Success,nil
}

// GetOneCateGory 返回一个文章分类
func (c *SqlCategory) GetOneCateGory(id int)(cate model.Category,err error){
	err = Db.Where("id = ?",id).First(&cate).Error
	return
}

//  返回所有的分类
func (c *SqlCategory) GetAllCateS(pageSize int,pageNum int)(dataS *[]model.Category,total int64,err error){

	err = Db.Find(&dataS).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error

	err = Db.Model(&dataS).Count(&total).Error
	if err != nil {
		return nil, 0,err
	}

	return dataS, total,nil
}
