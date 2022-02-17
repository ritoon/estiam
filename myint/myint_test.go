package myint

import (
	"math"
	"testing"
)

func TestMyIntAdd(t *testing.T) {
	data := []struct {
		title  string
		init   MyInt
		param  int
		Should MyInt
	}{
		{"A", 1, 1, 2},
		{"B", 2, 3, 5},
		{"C", math.MaxInt32, 1, math.MinInt32},
	}
	for _, v := range data {
		res, _ := v.init.Add(v.param)
		if res != v.Should {
			t.Errorf("for test %v waiting for %v and got %v", v.title, v.Should, res)
		}
	}
}
