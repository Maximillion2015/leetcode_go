package leetcode_go

import "fmt"

func twoSum(nums []int, target int)  []int {
	var result []int
	for i, num := range nums {
		target2 := target - num
		if isContain, index := contains(nums[i+1:], target2); isContain {
			result = append(result, i)
			result = append(result, index+i+1)
			return result
		}
	}
	return result
}

func contains(nums []int, target int) (bool, int) {
	for i, num := range nums{
		if  target == num {
			return true, i
		}
	}
	return false, -1
}

func main() {
	nums := []int{3, 3, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
