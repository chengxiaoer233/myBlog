package model

var ServerConf ServerConfStruct
var DbConf DbConfStruct
var ObsConf ObsConfStruct
var LogConf LogConfStruct

// 服务配置相关信息
type ServerConfStruct struct {
	AppMode  string `json:"appMode"`
	HttpPort string `json:"httpPort"`
	JwtKey   string `json:"jwtKey"`
}

// 数据库相关配置
type DbConfStruct struct {
	DbHost     string `json:"dbHost"`
	DbPort     string `json:"dbPort"`
	DbUser     string `json:"dbUser"`
	DbPassWord string `json:"dbPassWord"`
	DbName     string `json:"dbName"`
}

// 华为对象存储相关配置
type ObsConfStruct struct {
	AccessKey string `json:"accessKey"`
	ScriptKey string `json:"scriptKey"`
	EndPoint  string `json:"endPoint"`
	Bucket    string `json:"bucket"`
	Dir       string `json:"dir"`
	Host      string `json:"host"`
}

type LogConfStruct struct {
	LogPath    string `json:"logPath"`
	MaxAge     int    `json:"maxAge"`
	RotateTime int    `json:"rotateTime"`
	Level      int    `json:"level"`
}
