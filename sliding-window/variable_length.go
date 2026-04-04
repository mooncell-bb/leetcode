package slidingwindow

import "slices"

/* 3. 无重复字符的最长子串 */
func lengthOfLongestSubstring(s string) int {
	duplicate := make(map[rune]int, len(s))
	count := 0
	for l, r := 0, 0; r < len(s); r++ {
		duplicate[rune(s[r])]++

		for duplicate[rune(s[r])] > 1 {
			duplicate[rune(s[l])]--
			l++
		}

		count = max(count, r-l+1)
	}

	return count
}

/* 3090. 每个字符最多出现两次的最长子字符串 */
func maximumLengthSubstring(s string) int {
	duplicate := make(map[rune]int, len(s))
	count := 0
	for l, r := 0, 0; r < len(s); r++ {
		duplicate[rune(s[r])]++

		for duplicate[rune(s[r])] > 2 {
			duplicate[rune(s[l])]--
			l++
		}

		count = max(count, r-l+1)
	}

	return count
}

/* 1493. 删掉一个元素以后全为 1 的最长子数组 */
func longestSubarray(nums []int) int {
	count, ans := 0, 0
	for l, r := 0, 0; r < len(nums); r++ {
		if nums[r] == 0 {
			count++
		}

		for count > 1 {
			if nums[l] == 0 {
				count--
			}
			l++
		}

		ans = max(ans, r-l)
	}

	return ans
}

/* 3634. 使数组平衡的最少移除数目 */
func minRemoval(nums []int, k int) int {
	slices.Sort(nums)

	maxSave := 0
	for l, r := 0, 0; r < len(nums); r++ {
		for nums[l]*k < nums[r] {
			l++
		}

		maxSave = max(maxSave, r-l+1)
	}

	return len(nums) - maxSave
}

/* 1208. 尽可能使字符串相等 */
func equalSubstring(s string, t string, maxCost int) int {
	costs := make([]int, len(s))
	for i := range s {
		diff := int(s[i]) - int(t[i])
		if diff > 0 {
			costs[i] = diff
		} else {
			costs[i] = -diff
		}
	}

	total, ans := 0, 0
	for l, r := 0, 0; r < len(costs); r++ {
		total += costs[r]

		for total > maxCost {
			total -= costs[l]
			l++
		}

		ans = max(ans, r-l+1)
	}

	return ans
}

/* 904. 水果成篮 */
func totalFruit(fruits []int) int {
	duplicate := make(map[int]int, len(fruits))
	count := 0
	for l, r := 0, 0; r < len(fruits); r++ {
		duplicate[fruits[r]]++

		for len(duplicate) > 2 {
			duplicate[fruits[l]]--
			if duplicate[fruits[l]] == 0 {
				delete(duplicate, fruits[l])
			}
			l++
		}

		count = max(count, r-l+1)
	}
	return count
}

/* 1695. 删除子数组的最大得分 */
func maximumUniqueSubarray(nums []int) int {
	duplicate := make(map[int]int)
	sum, ans := 0, 0
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		duplicate[nums[r]]++

		for duplicate[nums[r]] > 1 {
			sum -= nums[l]
			duplicate[nums[l]]--
			l++
		}

		ans = max(ans, sum)
	}
	return ans
}

/* 2958. 最多 K 个重复元素的最长子数组 */
func maxSubarrayLength(nums []int, k int) int {
	duplicate := make(map[int]int)
	ans := 0
	for l, r := 0, 0; r < len(nums); r++ {
		duplicate[nums[r]]++

		for duplicate[nums[r]] > k {
			duplicate[nums[l]]--
			l++
		}

		ans = max(ans, r-l+1)
	}

	return ans
}

/* 2024. 考试的最大困扰度 */
func maxConsecutiveAnswers(answerKey string, k int) int {
	duplicate := make(map[rune]int, 2)
	count := 0
	for l, r := 0, 0; r < len(answerKey); r++ {
		duplicate[rune(answerKey[r])]++

		for duplicate[rune('T')] > k && duplicate[rune('F')] > k {
			duplicate[rune(answerKey[l])]--
			l++
		}

		count = max(count, r-l+1)
	}

	return count
}

/* 1004. 最大连续1的个数 III */
func longestOnes(nums []int, k int) int {
	count, ans := 0, 0
	for l, r := 0, 0; r < len(nums); r++ {
		if nums[r] == 0 {
			count++
		}

		for count > k {
			if nums[l] == 0 {
				count--
			}
			l++
		}

		ans = max(ans, r-l+1)
	}

	return ans
}
