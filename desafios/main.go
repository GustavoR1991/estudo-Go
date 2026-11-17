package main

import (
	"fmt"

	parenteses "desafios/02_parenteses_validos"
)

func main() {
	// fmt.Println(twosum.TwoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(parenteses.EhValido("()[]{}")) // true
	fmt.Println(parenteses.EhValido("([)]"))
}
