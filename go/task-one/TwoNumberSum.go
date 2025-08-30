package task_one

// TwoNumberSum 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoNumberSum(nums []int, target int) []int {
	for i, x := range nums { // 枚举 i
		for j := i + 1; j < len(nums); j++ { // 枚举 i 右边的 j
			if x+nums[j] == target { // 满足要求
				return []int{i, j} // 返回两个数的下标
			}
		}
	}
	return nil // 题目保证一定有解，代码不会执行到这里
}
