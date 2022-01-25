package main

import (
	"fmt"

	"github.com/DimDoremy/MdDrawer/common"
	"github.com/kataras/iris/v12"
)

//定义全局变量
var app *iris.Application
var err error

func main() {
	//初始化共通部分
	common.Init()
	defer common.Defer()

	//创建iris应用
	app = iris.New()

	//设置路由组，起始uri为/draws
	drawsAPI := app.Party(common.Kconf.String("Router.Root"))
	{
		drawsAPI.Use(iris.Compression) // 使用Compression中间件，进行404等的处理

		drawsAPI.Get(common.Kconf.String("Router.Index"), func(ctx iris.Context) {}) // 主页访问
		drawsAPI.Post("/sweepstake", func(ctx iris.Context) {})                      // 抽奖
	}

	port := common.Kconf.Int("Web.Port")
	err = app.Listen(fmt.Sprintf(":%d", port)) // 监听8080端口
	if err != nil {
		common.Logger.Error(fmt.Sprintf("Listen port %d failed", port), common.ErrPacker(err))
	}
}
