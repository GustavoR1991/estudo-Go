package main

import (
	"errors"
	"fmt"
	"sort"
)

func unicos(nums []int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("a lista de números está vazia")
	}

	counts := map[int]int{}
	for _, n := range nums {
		counts[n]++
	}

	var out []int
	for n, c := range counts {
		if c == 1 {
			out = append(out, n)
		}
	}
	sort.Ints(out)
	return out, nil
}

func main() {
	nums := []int{1, 1, 1, 2, 6, 7, 2, 3, 3}
	fmt.Println("numeros", nums)
	fmt.Scanln(&nums)

	res, err := unicos(nums)
	if err != nil {
		fmt.Println("erro:", err)
		return
	}
	fmt.Println("unicos", res)

}
