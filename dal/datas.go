package dal

import (
	"github.com/DimDoremy/MdDrawer/common"
	"github.com/DimDoremy/MdDrawer/dto"
)

func GetAllDatas() ([]dto.Datas, error) {
	datas := make([]dto.Datas, 0)
	err := common.Engine.Find(&datas)
	if err != nil {
		common.Logger.Warn("Engine load failed", common.ErrPacker(err))
		return nil, err
	}
	return datas, nil
}

func GetDatasByWhere(where string, args ...interface{}) ([]dto.Datas, error) {
	datas := make([]dto.Datas, 0)
	err := common.Engine.Where(where, args...).Find(&datas)
	if err != nil {
		common.Logger.Warn("Engine load failed", common.ErrPacker(err))
		return nil, err
	}
	return datas, nil
}
