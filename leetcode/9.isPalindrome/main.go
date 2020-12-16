package main

import (
	"fmt"
)

func main() {
	b := isPalindrome(-121)
	fmt.Println("b = ", b)
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) { // 当x<0时x不是回文数
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10 // ?
}
