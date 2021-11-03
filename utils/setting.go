package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("Error: read config.ini", err)
	}

	LoadServer(file)

	LoadDataBase(file)

	LoadQiNiu(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")
}

func LoadDataBase(file *ini.File) {
	DBHost = file.Section("database").Key("DBHost").MustString("localhost")
	DBPort = file.Section("database").Key("DBPort").MustString("3306")
	DBUser = file.Section("database").Key("DBUser").MustString("root")
	DBPassword = file.Section("database").Key("DBPassword").MustString("yzsby123")
	DBName = file.Section("database").Key("DBName").MustString("go-blog")

}

func LoadQiNiu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	Domain = file.Section("qiniu").Key("Domain").String()

}
