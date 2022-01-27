package main

import (
	"fmt"

	"github.com/DimDoremy/MdDrawer/common"
	"github.com/DimDoremy/MdDrawer/controller"
	"github.com/kataras/iris/v12"
)

var err error

func main() {
	//初始化共通部分
	common.Init()
	defer common.Defer()

	//设置路由组，起始uri为/draws
	drawsAPI := common.App.Party(common.Kconf.String("Router.Root"))
	{
		drawsAPI.Use(iris.Compression) // 使用Compression中间件，进行404等的处理

		drawsAPI.Get("/GetAllTexts", controller.GetAllTexts) // 主页访问
		drawsAPI.Get("/GetAllDatas", controller.GetAllDatas)
		drawsAPI.Get("/GetAllRares", controller.GetAllRares)
		drawsAPI.Post("/sweepstake", controller.PostTextsByWhere) // 抽奖
	}

	port := common.Kconf.Int("Web.Port")
	err = common.App.Listen(fmt.Sprintf(":%d", port)) // 监听8080端口
	if err != nil {
		common.App.Logger().Error(fmt.Sprintf("Listen port %d failed", port), err)
	}
}
