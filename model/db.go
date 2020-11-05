package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"myblog/utils"
	"time"
)

// 配置数据库连接的参数

var (
	db  *gorm.DB
	err error
)

func InitDb() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)
	db, err = gorm.Open(utils.Db, dsn)
	if err != nil {
		fmt.Println("连接数据库错误, err: ", err)
	}
	//defer db.Close()

	// 防止数据库建表时使用 复数 （users）
	db.SingularTable(true)

	_ = db.AutoMigrate(&User{}, &Article{}, &Category{})

	// 配置连接参数
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。 数据库连接时间不能大于框架的过期时间
	sqlDB.SetConnMaxLifetime(time.Second * 10)
	return db
}
