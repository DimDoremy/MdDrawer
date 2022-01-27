package bll

import (
	"github.com/DimDoremy/MdDrawer/common"
	"github.com/DimDoremy/MdDrawer/dal"
	"github.com/DimDoremy/MdDrawer/dto"
)

func GetAllDatas() dto.ResultDto {
	var jsonResult dto.ResultDto
	data, err := dal.GetAllDatas()
	if err != nil {
		common.App.Logger().Warn("Engine load failed", err)
		jsonResult.RtnMsg = err.Error()
		jsonResult.RtnCode = 1
		return jsonResult
	}
	jsonResult.RtnCode = 0
	if len(data) == 0 {
		jsonResult.RtnMsg = "查询结果为空"
		jsonResult.RtnData = nil
	} else {
		jsonResult.RtnMsg = "查询成功"
		jsonResult.RtnData = dto.ResultData{
			Data:  data,
			Count: len(data),
		}
	}
	return jsonResult
}

func GetDatasByWhere(column string, args []string) dto.ResultDto {
	var jsonResult dto.ResultDto
	var tmps []interface{}
	for _, v := range args {
		tmps = append(tmps, v)
	}
	data, err := dal.GetDatasByWhere(column, tmps...)
	if err != nil {
		common.App.Logger().Warn("Engine load failed:", err)
		jsonResult.RtnMsg = err.Error()
		jsonResult.RtnCode = 1
		return jsonResult
	}
	jsonResult.RtnCode = 0
	if len(data) == 0 {
		jsonResult.RtnMsg = "查询结果为空"
		jsonResult.RtnData = nil
	} else {
		jsonResult.RtnMsg = "查询成功"
		jsonResult.RtnData = dto.ResultData{
			Data:  data,
			Count: len(data),
		}
	}
	return jsonResult
}
