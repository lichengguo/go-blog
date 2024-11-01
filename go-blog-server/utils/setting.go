package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

// 声明配置文件参数变量
var (
	AppMode  string
	HttpPort string
	JwtKey   string
	ImgPath  string
	LogPath  string
	LogFile  string
	//LinkName string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

// init 只要包被导入就执行
func init() {
	// 获取程序所在的绝对目录
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("获取当前程序执行目录失败")
		return
	}
	file, err := ini.Load(fmt.Sprintf("%s/%s", dir, "config/config.ini"))
	if err != nil {
		fmt.Println("配置文件读取有误,请检查配置文件.", err)
		return
	}
	LoadServer(file)
	LoadData(file)

}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug") //默认值为debug
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
	ImgPath = file.Section("server").Key("ImgPath").MustString("uploads")
	LogPath = file.Section("server").Key("LogPath").MustString("log")
	LogFile = file.Section("server").Key("LogFile").MustString("test.log")
	//LinkName = file.Section("server").Key("LinkName").MustString("latest_log.log")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("root123")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
