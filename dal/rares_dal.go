package dal

import (
	"math/rand"
	"time"

	"github.com/DimDoremy/MdDrawer/common"
	"github.com/DimDoremy/MdDrawer/dto"
)

func GetAllRares() ([]dto.Rares, error) {
	rares := make([]dto.Rares, 0)
	err := common.Engine.Find(&rares)
	if err != nil {
		return nil, err
	}
	return rares, nil
}

func GetRaresByWhere(column string, args ...interface{}) ([]dto.Rares, error) {
	rares := make([]dto.Rares, 0)
	err := common.Engine.In(column, args...).Find(&rares)
	if err != nil {
		return nil, err
	}
	return rares, nil
}

func GetRandomPack10(nid uint32, flag bool) (map[int]map[int64]int, bool) {
	packs := make(map[int]map[int64]int)
	mrc := make(map[int]int)
	mrc[1], mrc[2], mrc[3], mrc[4] = 0, 0, 0, 0
	for i := 0; len(packs) < 10 && i < 100; i++ {
		pack := GetRandomPack(nid)
		mr := getPackMaxRare(pack)
		for k, v := range mr {
			mrc[k] += v
		}
		packs[i+1] = pack
		for len(packs) == 9 && mrc[1] == 0 && mrc[2] == 0 {
			pack = GetRandomPack(nid)
			mr = getPackMaxRare(pack)
			if mr[1] == 0 && mr[2] == 0 {
				continue
			} else if flag && mr[1] > 0 {
				for k, v := range mr {
					mrc[k] += v
				}
				packs[i+2] = pack
				break
			} else if !flag {
				for k, v := range mr {
					mrc[k] += v
				}
				packs[i+2] = pack
				break
			}
		}
	}
	flag = mrc[1] == 0 && mrc[2] == 1
	return packs, flag
}

func GetRandomPack(nid uint32) map[int64]int {
	pack := make(map[int64]int)
	for i := 0; len(pack) < common.CNumber && i < 100; i++ {
		rand.Seed(time.Now().UnixNano() + int64(nid) + int64(i))
		percent := rand.Intn(100) + 1
		where := "rare = ?"
		var rare *dto.Rares
		switch {
		case percent < common.UltimateRare:
			rare = getRandomCard(where, 1)
		case percent < common.SpecialRare+common.UltimateRare:
			rare = getRandomCard(where, 2)
		case percent < common.Rare+common.SpecialRare+common.UltimateRare:
			rare = getRandomCard(where, 3)
		default:
			rare = getRandomCard(where, 4)
		}
		pack[rare.Id] = rare.Rare
	}
	return pack
}

func getRandomCard(where string, args ...interface{}) *dto.Rares {
	rare := new(dto.Rares)
	count, err := common.Engine.Where(where, args...).Count(rare)
	if err != nil {
		common.App.Logger().Warn("Engine load failed", err)
	}
	offset := int(rand.Int63n(count) + 1)
	_, err = common.Engine.Where(where, args...).Limit(1, offset).Get(rare)
	if err != nil {
		common.App.Logger().Warn("Engine load failed", err)
	}
	return rare
}

func getPackMaxRare(pack map[int64]int) map[int]int {
	rare := make(map[int]int)
	rare[1], rare[2], rare[3], rare[4] = 0, 0, 0, 0
	for _, v := range pack {
		rare[v] += 1
	}
	return rare
}
