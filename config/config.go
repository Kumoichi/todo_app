package config

import (
	"log"
	"todo_app/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config ConfigList

// インポートされたときに実行される。
func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	// ini.Loadは"gopkg.in/go-ini/ini.v1"パッケージに存在する関数
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	// Config変数に新しい構造体を作っている
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
