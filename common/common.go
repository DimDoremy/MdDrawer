package common

import (
	"os"

	"github.com/kataras/iris/v12"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

var Engine *xorm.Engine
var err error

//全局koanf实例。使用"."作为关键路径分隔符。这可以是“/”或任何字符。
var Kconf = koanf.New(".")

//定义全局变量
var App *iris.Application

var UltimateRare int
var SpecialRare int
var Rare int
var Normal int
var CNumber int

func Init() {
	//创建iris应用
	App = iris.Default()

	//使用koanf配置读取读取toml配置文件
	if err = Kconf.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		App.Logger().Error("failed load config", err)
	}
	App.HandleDir(Kconf.String("Router.Index"), "./frontend")
	App.HandleDir(Kconf.String("Router.Build"), "./frontend/Build")
	App.HandleDir(Kconf.String("Router.TemplateData"), "./frontend/TemplateData")

	//初始化xorm数据库框架的连接
	Engine, err = xorm.NewEngine(Kconf.String("SQL.Driver"), Kconf.String("SQL.ConnString"))
	if err != nil {
		App.Logger().Error("failed to create SQL connection", err)
		os.Exit(1)
	}
	Engine.Logger().SetLevel(xlog.DEFAULT_LOG_LEVEL)

	UltimateRare = Kconf.MustInt("Percent.UR")
	SpecialRare = Kconf.MustInt("Percent.SR")
	Rare = Kconf.MustInt("Percent.R")
	Normal = Kconf.MustInt("Percent.N")
	CNumber = Kconf.MustInt("Percent.CardNumbers")
}

func Defer() {
	Engine.Close()
}
