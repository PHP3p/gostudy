package model

import (
	"day02/utils/errmsg"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20); not null" json:"username"`
	Password string `gorm:"type:varchar(20); not null" json:"password"`
	Role     int    `gorm:"type:int";json:"role"`
}

func ExitsUser(name string) (code int) {
	var u User
	fmt.Sprintf("名字",name)
	db.Select("id").Where("username=?", name).First(&u)
	if u.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

func CreatUser(data *User) int {
	fmt.Println("shuju",&data)

	err := db.Create(&data).Error
	if err != nil {
		log.Fatal("创建用户失败",err)
		return errmsg.ERROR
	}
	return errmsg.SUCCSE

}
