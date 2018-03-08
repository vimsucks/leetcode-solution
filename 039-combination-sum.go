package main

func main() {
	combinationSum([]int{2, 3, 6, 7}, 7)
	combinationSum([]int{1, 2}, 3)
	combinationSum([]int{1, 2}, 4)
}

func combinationSum(candidates []int, target int) [][]int {
	combinations := [][]int{}
	for i, c := range candidates {
		sum := c
		combination := []int{c}
		combinationSumIter(candidates[i:], target, combination, sum, &combinations)
	}
	return combinations
}

func combinationSumIter(candidates []int, target int, combination []int, sum int, combinations *[][]int) {
	if sum == target {
		// 为什么这里要拷贝一个 slice ？
		comb := make([]int, len(combination))
		copy(comb, combination)
		*combinations = append(*combinations, comb)
	} else if sum < target {
		for i, c := range candidates {
			combinationSumIter(candidates[i:], target, append(combination, c), sum+c, combinations)
		}
	}
}
