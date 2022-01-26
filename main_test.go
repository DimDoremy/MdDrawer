package main

import (
	"os"
	"testing"

	"github.com/DimDoremy/MdDrawer/common"
	"github.com/DimDoremy/MdDrawer/dal"
)

func TestMain(m *testing.M) {
	common.Init()
	defer common.Defer()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCommonNumberId(t *testing.T) {
	t.Run("TestGetNumberId", func(t *testing.T) {
		got := common.GetNumberId("doremy")
		if got != uint32(1332064733) {
			t.Error("TestGetNumberId failed")
		}
	})
}

func TestDalTexts(t *testing.T) {
	t.Run("TestGetAllTexts", func(t *testing.T) {
		got, err := dal.GetAllTexts()
		if got == nil || len(got) <= 0 || err != nil {
			t.Error("TestGetAllTexts failed")
		}
		t.Log(len(got))
	})
	t.Run("TestGetTextsByWhere", func(t *testing.T) {
		got, err := dal.GetTextsByWhere("desc LIKE ?", "%PSY%")
		if got == nil || len(got) <= 0 || err != nil {
			t.Error("TestGetTextsByWhere failed")
		}
		t.Log(len(got))
	})
}

func TestDalDatas(t *testing.T) {
	t.Run("TestGetAllDatas", func(t *testing.T) {
		got, err := dal.GetAllDatas()
		if got == nil || len(got) <= 0 || err != nil {
			t.Error("TestGetAllDatas failed")
		}
		t.Log(len(got))
	})
	t.Run("TestGetDatasByWhere", func(t *testing.T) {
		got, err := dal.GetDatasByWhere("type = ?", 33)
		if got == nil || len(got) <= 0 || err != nil {
			t.Error("TestGetDatasByWhere failed")
		}
		t.Log(len(got))
	})
}

func TestDalRares(t *testing.T) {
	t.Run("TestGetRandomPack", func(t *testing.T) {
		got := dal.GetRandomPack(uint32(1332064733))
		t.Log(got)
	})
	t.Run("TestGetRandomPack10", func(t *testing.T) {
		got, flag := dal.GetRandomPack10(uint32(1332064733), false)
		t.Log(got)
		if flag {
			t.Log("lowwer")
		}
	})
}
