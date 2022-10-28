package service

import "testing"

// 创建表
func TestCreateRow(t *testing.T) {
	CreateRow()
}

func TestSelectCreateRow(t *testing.T) {
	SelectCreateRow()
}

func TestCreateInBatches(t *testing.T) {
	CreateInBatches()
}

func TestCreateByMap(t *testing.T) {
	CreateByMap()
}
