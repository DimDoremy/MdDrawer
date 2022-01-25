package dal

import (
	"github.com/DimDoremy/MdDrawer/common"
	"github.com/DimDoremy/MdDrawer/dto"
)

//获取全部Texts数据库信息
func GetAllTexts() []dto.Texts {
	texts := make([]dto.Texts, 0)
	err := common.Engine.Find(&texts)
	if err != nil {
		common.Logger.Error("Engine load failed", common.ErrPacker(err))
	}
	return texts
}

//通过where进行查询
func GetTextsByWhere(where string, args ...interface{}) []dto.Texts {
	texts := make([]dto.Texts, 0)
	err := common.Engine.Where(where, args...).Find(&texts)
	if err != nil {
		common.Logger.Error("Engine load failed", common.ErrPacker(err))
	}
	return texts
}
