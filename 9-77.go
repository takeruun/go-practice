package main

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      int
	DbName    string
	SqlDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SqlDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Println("%T, %v\n", Config.Port, Config.Port)
	fmt.Println("%T, %v\n", Config.DbName, Config.DbName)
	fmt.Println("%T, %v\n", Config.SqlDriver, Config.SqlDriver)
}
