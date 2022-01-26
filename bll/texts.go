package bll

import (
	"github.com/DimDoremy/MdDrawer/common"
	"github.com/DimDoremy/MdDrawer/dal"
	"github.com/DimDoremy/MdDrawer/dto"
)

func GetAllTexts() dto.ResultDto {
	var jsonResult dto.ResultDto
	data, err := dal.GetAllTexts()
	if err != nil {
		common.Logger.Warn(err.Error())
		jsonResult.RtnMsg = err.Error()
		jsonResult.RtnCode = 1
	}
	if len(data) == 0 {
		jsonResult.RtnData = nil
	} else {
		jsonResult.RtnData = struct {
			Data  interface{} `json:"data"`
			Count int         `json:"count"`
		}{
			Data:  data,
			Count: len(data),
		}
	}
	return jsonResult
}
