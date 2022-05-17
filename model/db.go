package model

import (
	"day02/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	connet := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	// 下面写法 := 局部变量覆盖了全局变量 var db *gorm.DB 会导致空指针错误
	/*db, err := gorm.Open(utils.Db, fmt.Sprintf(connet*/
	db, err = gorm.Open(utils.Db, fmt.Sprintf(connet,
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		log.Fatal("db connect fail!", err, fmt.Sprintf(connet,
			utils.DbUser,
			utils.DbPassWord,
			utils.DbHost,
			utils.DbPort,
			utils.DbName,
		))
	}
	// 表名字单复数处理
	db.SingularTable(true)
	// 连接池最大闲置
	db.DB().SetMaxIdleConns(10)
	// 连接池最大连接
	db.DB().SetMaxOpenConns(100)
	// 连接超时时间
	db.DB().SetConnMaxLifetime(time.Second * 10)

	// db.AutoMigrate(&User{}, &Category{}, &Article{})
	// 上线是 要有关闭处理

	// defer db.Close()

}
