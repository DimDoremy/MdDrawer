package dal

import (
	"github.com/DimDoremy/MdDrawer/common"
	"github.com/DimDoremy/MdDrawer/dto"
)

func GetAllDatas() ([]dto.Datas, error) {
	datas := make([]dto.Datas, 0)
	err := common.Engine.Find(&datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}

func GetDatasByWhere(column string, args ...interface{}) ([]dto.Datas, error) {
	datas := make([]dto.Datas, 0)
	err := common.Engine.In(column, args...).Find(&datas)
	if err != nil {
		return nil, err
	}
	return datas, nil
}
