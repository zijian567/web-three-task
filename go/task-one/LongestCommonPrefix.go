package task_one

// LongestCommonPrefix 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
func LongestCommonPrefix(strs []string) string {
	s0 := strs[0]
	for j, c := range s0 { // 从左到右
		for _, s := range strs { // 从上到下
			if j == len(s) || s[j] != byte(c) { // 这一列有字母缺失或者不同
				return s0[:j] // 0 到 j-1 是公共前缀
			}
		}
	}
	return s0
}
