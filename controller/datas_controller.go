package controller

import (
	"encoding/json"

	"github.com/DimDoremy/MdDrawer/bll"
	"github.com/DimDoremy/MdDrawer/dto"
	"github.com/kataras/iris/v12"
)

func GetAllDatas(ctx iris.Context) {
	bllreader := bll.GetAllDatas()
	data, err := json.Marshal(bllreader)
	if err != nil {
		ctx.StopWithError(400, err)
	}
	ctx.Binary(data)
}

func PostDatasByWhere(ctx iris.Context) {
	bllreader := dto.ResultDto{
		RtnCode: 1,
		RtnMsg:  "Request read failed",
		RtnData: nil,
	}
	req := ctx.FormValues()
	bllreader = bll.GetDatasByWhere(req["column"][0], req["args"])
	data, err := json.Marshal(bllreader)
	if err != nil {
		ctx.StopWithError(400, err)
	}
	ctx.Binary(data)
}
