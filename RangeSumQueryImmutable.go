package main

import "fmt"

type NumArray struct {
	nums []int
}


func Constructor(nums []int) NumArray {
	return NumArray{nums: nums}
}


func (this *NumArray) SumRange(left int, right int) int {
	var result int
	for _, num := range this.nums[left:right+1] {
		result += num
	}
	return result
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	param_1 := obj.SumRange(1,5)
	fmt.Print(param_1)
}
/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
