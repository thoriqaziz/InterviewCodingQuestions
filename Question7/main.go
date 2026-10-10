// sum two

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println("Indices:", result) // Output: Indices: [0 1]
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		fmt.Printf("index %d num %d\n", index, num)
		complement := target - num

		if complementIndex, exists := m[complement]; exists {
			return []int{complementIndex, index}
		}
		m[num] = index
	}

	return []int{}
}
