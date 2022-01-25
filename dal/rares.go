package dal

import (
	"math/rand"
	"time"

	"github.com/DimDoremy/MdDrawer/common"
	"github.com/DimDoremy/MdDrawer/dto"
)

type void struct{}

var member void

func GetAllRares() []dto.Rares {
	rares := make([]dto.Rares, 0)
	err := common.Engine.Find(&rares)
	if err != nil {
		common.Logger.Error("Engine load failed", common.ErrPacker(err))
	}
	return rares
}

func GetRaresByWhere(where string, args ...interface{}) []dto.Rares {
	rares := make([]dto.Rares, 0)
	err := common.Engine.Where(where, args...).Find(&rares)
	if err != nil {
		common.Logger.Error("Engine load failed", common.ErrPacker(err))
	}
	return rares
}

//私有：随机获取1张
func getRandomCard(nid uint32) int64 {
	rand.Seed(time.Now().UnixNano() + int64(nid))
	sql := "SELECT * FROM Rares WHERE rare = ?"
	percent := rand.Intn(100) + 1
	rare := new(dto.Rares)
	switch {
	case percent < common.UltimateRare:
		_, err := common.Engine.SQL(sql, 1).Get(rare)
		if err != nil {
			common.Logger.Error("Engine load failed", common.ErrPacker(err))
		}
	case percent < common.SpecialRare+common.UltimateRare:
		_, err := common.Engine.SQL(sql, 2).Get(rare)
		if err != nil {
			common.Logger.Error("Engine load failed", common.ErrPacker(err))
		}
	case percent < common.Rare+common.SpecialRare+common.UltimateRare:
		_, err := common.Engine.SQL(sql, 3).Get(rare)
		if err != nil {
			common.Logger.Error("Engine load failed", common.ErrPacker(err))
		}
	default:
		_, err := common.Engine.SQL(sql, 4).Get(rare)
		if err != nil {
			common.Logger.Error("Engine load failed", common.ErrPacker(err))
		}
	}
	return rare.Id
}

func GetRandomPack(nid uint32) []int64 {
	pack := map[int64]void{}
	for {
		// pack[]
	}
}
