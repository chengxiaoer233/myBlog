package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"myBlog/model"
	"time"
)

var Db *gorm.DB

func initDb() {

	user := model.DbConf.DbUser
	pass := model.DbConf.DbPassWord
	host := model.DbConf.DbHost
	port := model.DbConf.DbPort
	db := model.DbConf.DbName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, db)

	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{

		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("initDb() gorm open err = %s", err.Error()))
	}

	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	// 刷新数据库中的表格，使其保持最新，即让数据库之前存储的记录的表格字段和程序中最新使用的表格字段保持一致（只增不减）
	_ = Db.AutoMigrate(&model.User{}, &model.Article{}, &model.Category{}, model.Profile{}, model.Comment{}, model.Readdetail{})

	// *gorm.DB 返回一个通用的数据库接口 *sql.DB
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := Db.DB()
	if err != nil {
		panic(fmt.Sprintf("initDb() change to sqlDB error ,err= %s", err.Error()))
	}

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(60 * time.Second)
}
