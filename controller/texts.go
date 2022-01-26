package controller

import (
	"github.com/DimDoremy/MdDrawer/bll"
	"github.com/kataras/iris/v12"
)

func GetAllTexts(ctx iris.Context) {
	bll := bll.GetAllTexts()
	ctx.JSON(bll)
}
