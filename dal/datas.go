package dal

import (
	"github.com/DimDoremy/MdDrawer/common"
	"github.com/DimDoremy/MdDrawer/dto"
)

func GetAllDatas() []dto.Datas {
	datas := make([]dto.Datas, 0)
	err := common.Engine.Find(&datas)
	if err != nil {
		common.Logger.Error("Engine load failed", common.ErrPacker(err))
	}
	return datas
}

func GetDatasByWhere(where string, args ...interface{}) []dto.Datas {
	datas := make([]dto.Datas, 0)
	err := common.Engine.Where(where, args...).Find(&datas)
	if err != nil {
		common.Logger.Error("Engine load failed", common.ErrPacker(err))
	}
	return datas
}
