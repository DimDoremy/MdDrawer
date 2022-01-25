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

func TestDalTexts(t *testing.T) {
	t.Run("TestGetAllTexts", func(t *testing.T) {
		got := dal.GetAllTexts()
		if got == nil || len(got) <= 0 {
			t.Error("TestGetAllTexts failed")
		}
		t.Log(len(got))
	})
	t.Run("TestGetTextsByWhere", func(t *testing.T) {
		got := dal.GetTextsByWhere("desc LIKE ?", "%PSY%")
		if got == nil || len(got) <= 0 {
			t.Error("TestGetTextsByWhere failed")
		}
		t.Log(len(got))
	})
}

func TestDalDatas(t *testing.T) {
	t.Run("TestGetAllDatas", func(t *testing.T) {
		got := dal.GetAllDatas()
		if got == nil || len(got) <= 0 {
			t.Error("TestGetAllDatas failed")
		}
		t.Log(len(got))
	})
	t.Run("TestGetDatasByWhere", func(t *testing.T) {
		got := dal.GetDatasByWhere("type = ?", 33)
		if got == nil || len(got) <= 0 {
			t.Error("TestGetDatasByWhere failed")
		}
		t.Log(len(got))
	})
}

func TestCommonNumberId(t *testing.T) {
	t.Run("TestGetNumberId", func(t *testing.T) {
		got := common.GetNumberId("doremy")
		if got != uint32(1332064733) {
			t.Error("TestGetNumberId failed")
		}
	})
}
