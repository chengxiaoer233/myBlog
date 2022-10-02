package model

var ServerConf ServerConfStruct
var DbConf DbConfStruct

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
