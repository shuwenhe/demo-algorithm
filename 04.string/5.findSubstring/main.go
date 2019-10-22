// 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
//
// 示例 1：
// 输入：
//   s = "barfoothefoobarman",
//   words = ["foo","bar"]
// 输出：[0,9]
// 解释：
// 从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
// 输出的顺序不重要, [9,0] 也是有效答案。
// 示例 2：
// 输入：
//   s = "wordgoodgoodgoodbestword",
//   words = ["word","good","best","word"]
// 输出：[]

package main

import (
	"fmt"
)

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	res := findSubstring(s, words)
	fmt.Println(res)
}

func findSubstring(s string, words []string) []int {
	lens := len(s)
	res := make([]int, 0, lens)

	lenws := len(words)
	if lens == 0 || lenws == 0 || lens < lenws*len(words[0]) {
		return res
	}

	lenw := len(words[0])

	// record 记录 words 中每个单词总共出现的次数
	record := make(map[string]int, lenws)
	for _, w := range words {
		if len(w) != lenw {
			// 新添加的 example 2 出现了 words 中单词长度不一致的情况。
			// 这个违反了假设
			// 直接返回 res
			return res
		}
		record[w]++
	}

	// remain 记录 words 中每个单词还能出现的次数
	remain := make(map[string]int, lenws)
	// count 记录符合要求的单词的连续出现次数
	count := 1 // count 的初始值只要不为 0 就可以
	left, right := 0, 0

	/**
	 * s[left:right] 作为一个窗口存在
	 * 假设 word:= s[right:right+lenw]
	 * 如果 remain[word]>0 ，那么移动窗口 右边，同时更新移动后 s[left:right] 的统计信息
	 * 		remain[word]--
	 * 		right += lenw
	 * 		count++
	 * 否则，移动窗口 左边，同时更新移动后 s[left:right] 的统计信息
	 * 		remain[s[left:left+lenw]]++
	 * 		count--
	 * 		left += lenw
	 *
	 * 每次移动窗口 右边 后，如果 count = lenws ，那么
	 * 		说明找到了一个符合条件的解
	 * 		append(res, left)，然后
	 * 		移动窗口 左边
	 */

	// reset 重置 remain 和 count
	reset := func() {
		if count == 0 {
			// 统计记录没有被修改，不需要重置
			// 因为有这个预判，所以需要第一次执行 reset 时
			// count 的值不能为 0
			// 即 count 的初始值不能为 0
			return
		}
		for k, v := range record {
			remain[k] = v
		}
		count = 0
	}

	// moveLeft 让 left 指向下一个单词
	moveLeft := func() {
		// left 指向下一个单词前，需要修改统计记录
		// 增加 left 指向的 word 可出现次数一次，
		remain[s[left:left+lenw]]++
		// 连续符合条件的单词数减少一个
		count--
		// left 后移一个 word 的长度
		left += lenw
	}

	// left 需要分别从{0,1,...,lenw-1}这些位置开始检查，才能不遗漏
	for i := 0; i < lenw; i++ {
		left, right = i, i
		reset()

		// s[left:] 的长度 >= words 中所有 word 组成的字符串的长度时，
		// s[left:] 中才有可能存在要找的字符串
		for lens-left >= lenws*lenw {
			word := s[right : right+lenw]
			remainTimes, ok := remain[word]

			switch {
			case !ok:
				// word 不在 words 中
				// 从 right+lenw 处，作为新窗口，重新开始统计
				left, right = right+lenw, right+lenw
				reset()
			case remainTimes == 0:
				// word 的出现次数上次就用完了
				// 说明 word 在 s[left:right] 中出现次数过多
				moveLeft()
				// 这个case会连续出现
				// 直到 s[left:right] 中的统计结果是 remain[word] == 1
				// 这个过程中，right 一直不变
			default:
				// ok && remainTimes > 0，word 符合出现的条件
				// moveRight
				remain[word]--
				count++
				right += lenw
				// 检查 words 能否排列组合成 s[left:right]
				if count == lenws {
					res = append(res, left)
					// moveLeft 可以避免重复统计 s[left+lenw:right] 中的信息
					moveLeft()
				}
			}
		}
	}

	return res
}
