package goutil

import (
	"testing"
)

var nowDate int = 201409

func TestIncreaseDate(t *testing.T) {
	list := IncreaseData(nowDate)
	if 12 != len(list) {
		t.Errorf("generate error:%v", list)
	}
}

func BenchmarkIncreateDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IncreaseData(nowDate)
	}
}

func TestDecreaseDate(t *testing.T) {
	list := DecreaseData(nowDate)
	if 12 != len(list) {
		t.Errorf("generate error:%v", list)
	}
}

func BenchmarkDecreateDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecreaseData(nowDate)
	}
}
