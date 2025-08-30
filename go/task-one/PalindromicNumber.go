package task_one

// IsPalindrome checks if a number is palindrome or not
func IsPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	if n < 10 {
		return true
	}
	return n == reverse(n)
}

func reverse(n int) int {
	var rev int
	for n > 0 {
		rev = rev*10 + n%10
		n = n / 10
	}
	return rev
}
