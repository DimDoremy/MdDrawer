package common

import (
	"os"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

var Engine *xorm.Engine
var Logger *zap.Logger
var err error

//全局koanf实例。使用"."作为关键路径分隔符。这可以是“/”或任何字符。
var Kconf = koanf.New(".")

var UltimateRare int
var SpecialRare int
var Rare int
var Normal int
var CNumber int

func Init() {
	//使用zap日志库，如果没能初始化则直接panic
	Logger, _ = zap.NewProduction()
	defer Logger.Sync() // 如果可以，刷新缓冲区

	//使用koanf配置读取读取toml配置文件
	if err = Kconf.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		Logger.Error("failed load config", zap.Error(err))
	}

	//初始化xorm数据库框架的连接
	Engine, err = xorm.NewEngine(Kconf.String("SQL.Driver"), Kconf.String("SQL.ConnString"))
	if err != nil {
		Logger.Error("failed to create SQL connection", zap.Error(err))
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
	Logger.Sync()
}

func ErrPacker(err error) zapcore.Field {
	return zap.Error(err)
}
