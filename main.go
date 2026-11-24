package main

import "fmt"

// MaxOfMins returns, for every window size w = 1..n, the maximum over all
// windows of that size of the window's minimum. Each element is the minimum
// of windows up to a size set by its nearest smaller neighbors; fill those
// buckets, then cascade down.
func MaxOfMins(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	left := make([]int, n)  // index of previous strictly smaller, else -1
	right := make([]int, n) // index of next strictly smaller, else n
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = stack[:0]
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	const minInt = -int(^uint(0)>>1) - 1
	ans := make([]int, n+1) // ans[w] for window size w, 1-based
	for w := range ans {
		ans[w] = minInt
	}
	for i := 0; i < n; i++ {
		w := right[i] - left[i] - 1 // largest window where nums[i] is the minimum
		if nums[i] > ans[w] {
			ans[w] = nums[i]
		}
	}
	for w := n - 1; w >= 1; w-- {
		if ans[w+1] > ans[w] {
			ans[w] = ans[w+1] // a min of a bigger window also fits smaller ones
		}
	}
	return ans[1:]
}

func main() {
	fmt.Println(MaxOfMins([]int{10, 20, 30, 50, 10, 70, 30}))
}
