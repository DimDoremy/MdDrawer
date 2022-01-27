package controller

import (
	"encoding/json"

	"github.com/DimDoremy/MdDrawer/bll"
	"github.com/DimDoremy/MdDrawer/dto"
	"github.com/kataras/iris/v12"
)

func GetAllRares(ctx iris.Context) {
	bllreader := bll.GetAllRares()
	data, err := json.Marshal(bllreader)
	if err != nil {
		ctx.StopWithError(400, err)
	}
	ctx.Binary(data)
}

func PostRaresByWhere(ctx iris.Context) {
	bllreader := dto.ResultDto{
		RtnCode: 1,
		RtnMsg:  "Request read failed",
		RtnData: nil,
	}
	req := ctx.FormValues()
	bllreader = bll.GetRaresByWhere(req["column"][0], req["args"])
	data, err := json.Marshal(bllreader)
	if err != nil {
		ctx.StopWithError(400, err)
	}
	ctx.Binary(data)
}
