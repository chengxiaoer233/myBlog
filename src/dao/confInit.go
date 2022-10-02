package dao

import (
	"fmt"
	"gopkg.in/ini.v1"
	"myBlog/model"
)

func Init() {

	// 读取配置文件，如果没有读到就直接panic
	iniFile, err := ini.Load("./etc/config.ini")
	if err != nil {
		panic(fmt.Sprintf("注意：配置文件读取错误，请检查文件路径:%s", err.Error()))
	}

	loadServerConf(iniFile)

	loadDbConf(iniFile)
}

// 加载服务器配置文件
func loadServerConf(iniFile *ini.File) {
	model.ServerConf.AppMode = iniFile.Section("server").Key("AppMode").MustString("debug")
	model.ServerConf.HttpPort = iniFile.Section("server").Key("HttpPort").MustString(":3000")
	if model.ServerConf.HttpPort[0] != ':' { // 允许业务部携带：，直接传端口号
		model.ServerConf.HttpPort = fmt.Sprintf(":%s", model.ServerConf.HttpPort)
	}
	model.ServerConf.JwtKey = iniFile.Section("server").Key("JwtKey").MustString("89js82js72")
}

// 加载数据库配置文件
func loadDbConf(iniFile *ini.File) {
	model.DbConf.DbHost = iniFile.Section("database").Key("DbHost").MustString("localhost")
	model.DbConf.DbPort = iniFile.Section("database").Key("DbPort").MustString("3306")
	model.DbConf.DbUser = iniFile.Section("database").Key("DbUser").MustString("root")
	model.DbConf.DbPassWord = iniFile.Section("database").Key("DbPassWord").MustString("12345678")
	model.DbConf.DbName = iniFile.Section("database").Key("DbName").MustString("ginblog")
}
