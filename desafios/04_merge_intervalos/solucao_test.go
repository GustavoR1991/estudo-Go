package intervalos

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	casos := []struct {
		in   [][]int
		quer [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
		{[][]int{{5, 6}}, [][]int{{5, 6}}},
		{[][]int{{8, 9}, {1, 2}, {3, 5}, {4, 7}}, [][]int{{1, 2}, {3, 7}, {8, 9}}},
	}
	for _, c := range casos {
		got := Merge(c.in)
		if !reflect.DeepEqual(got, c.quer) {
			t.Errorf("Merge(%v) = %v; quer %v", c.in, got, c.quer)
		}
	}
}
