package dal

import (
	"github.com/DimDoremy/MdDrawer/common"
	"github.com/DimDoremy/MdDrawer/dto"
)

//获取全部Texts数据库信息
func GetAllTexts() ([]dto.Texts, error) {
	texts := make([]dto.Texts, 0)
	err := common.Engine.Find(&texts)
	if err != nil {
		return nil, err
	}
	return texts, nil
}

//通过where进行查询
func GetTextsByWhere(column string, args ...interface{}) ([]dto.Texts, error) {
	texts := make([]dto.Texts, 0)
	err := common.Engine.In(column, args...).Find(&texts)
	if err != nil {
		return nil, err
	}
	return texts, nil
}
