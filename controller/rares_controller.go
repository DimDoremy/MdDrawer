package controller

import (
	"github.com/DimDoremy/MdDrawer/bll"
	"github.com/DimDoremy/MdDrawer/dto"
	"github.com/kataras/iris/v12"
)

func GetAllRares(ctx iris.Context) {
	bllreader := bll.GetAllRares()
	ctx.JSON(bllreader)
}

func PostRaresByWhere(ctx iris.Context) {
	bllreader := dto.ResultDto{
		RtnCode: 1,
		RtnMsg:  "Request read failed",
		RtnData: nil,
	}
	req := ctx.FormValues()
	bllreader = bll.GetRaresByWhere(req["column"][0], req["args"])
	ctx.JSON(bllreader)
}
