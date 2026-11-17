package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	casos := []struct {
		nums []int
		alvo int
		quer []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}},
	}
	for _, c := range casos {
		got := TwoSum(c.nums, c.alvo)
		if !reflect.DeepEqual(got, c.quer) {
			t.Errorf("TwoSum(%v, %d) = %v; quer %v", c.nums, c.alvo, got, c.quer)
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i
	}
	for b.Loop() {
		TwoSum(nums, 19997)
	}
}
