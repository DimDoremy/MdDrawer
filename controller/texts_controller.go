package controller

import (
	"github.com/DimDoremy/MdDrawer/bll"
	"github.com/DimDoremy/MdDrawer/dto"
	"github.com/kataras/iris/v12"
)

func GetAllTexts(ctx iris.Context) {
	bllreader := bll.GetAllTexts()
	ctx.JSON(bllreader)
}

func PostTextsByWhere(ctx iris.Context) {
	bllreader := dto.ResultDto{
		RtnCode: 1,
		RtnMsg:  "Request read failed",
		RtnData: nil,
	}
	req := ctx.FormValues()
	bllreader = bll.GetTextsByWhere(req["column"][0], req["args"])
	ctx.JSON(bllreader)
}
