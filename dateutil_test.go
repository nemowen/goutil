package goutil

import (
	"testing"
)

var nowDate int = 201409

func TestIncreaseDate(t *testing.T) {
	list := IncreaseDate(nowDate)
	if 12 != len(list) {
		t.Errorf("generate error:%v", list)
	}
}

func BenchmarkIncreateDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IncreaseDate(nowDate)
	}
}

func TestDecreaseDate(t *testing.T) {
	list := DecreaseDate(nowDate)
	if 12 != len(list) {
		t.Errorf("generate error:%v", list)
	}
}

func BenchmarkDecreateDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecreaseDate(nowDate)
	}
}
