package utils

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	// todo 文件加载路径 相对路径
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("ini文件加载失败", err)
	}
	LoadServer(file)
	LoadData(file)
}
func LoadServer(f *ini.File) {
	server := f.Section("server")
	AppMode = server.Key("AppMode").MustString("debug")
	HttpPort = server.Key("HttpPort").MustString("8000")

}
func LoadData(f *ini.File) {
	database := f.Section("database")
	Db = database.Key("Db").MustString("debug")
	DbHost = database.Key("DbHost").MustString("debug")
	DbPort = database.Key("DbPort").MustString("debug")
	DbUser = database.Key("DbUser").MustString("debug")
	DbPassWord = database.Key("DbPassWord").MustString("debug")
	DbName = database.Key("DbName").MustString("debug")
}
