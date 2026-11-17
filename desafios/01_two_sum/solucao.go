package twosum

// Two Sum (Fácil)
//
// Dado um slice de inteiros `nums` e um inteiro `alvo`, retorne os ÍNDICES
// dos dois números que somados dão `alvo`.
// Existe exatamente uma solução, e você não pode usar o mesmo elemento duas vezes.
// Retorne os índices em ordem crescente.
//
// Exemplo: nums = [2,7,11,15], alvo = 9  ->  [0,1]
//
// Desafio extra: resolva em O(n).
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
