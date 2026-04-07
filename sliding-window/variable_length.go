package slidingwindow

import (
	"math"
	"slices"
)

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

/* 209. 长度最小的子数组 */
func minSubArrayLen(target int, nums []int) int {
	sum, minLen := 0, math.MaxInt
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]

		for sum >= target {
			minLen = min(minLen, r-l+1)

			sum -= nums[l]
			l++
		}
	}

	if minLen == math.MaxInt {
		return 0
	}

	return minLen
}

/* 3795. 不同元素和至少为 K 的最短子数组长度 */
func minLength(nums []int, k int) int {
	duplicate := make(map[int]int, len(nums))
	sum, minLen := 0, math.MaxInt
	for l, r := 0, 0; r < len(nums); r++ {
		duplicate[nums[r]]++
		if duplicate[nums[r]] == 1 {
			sum += nums[r]
		}

		for sum >= k {
			minLen = min(minLen, r-l+1)

			duplicate[nums[l]]--
			if duplicate[nums[l]] == 0 {
				sum -= nums[l]
			}

			l++
		}
	}

	if minLen == math.MaxInt {
		return -1
	}

	return minLen
}

/* 2904. 最短且字典序最小的美丽子字符串 */
func shortestBeautifulSubstring(s string, k int) string {
	count, ans := 0, []int{math.MaxInt, -1, -1}
	for l, r := 0, 0; r < len(s); r++ {
		if s[r] == '1' {
			count++
		}

		for count >= k {
			if ans[0] > r-l+1 || ans[0] == r-l+1 && s[ans[1]:ans[2]+1] > s[l:r+1] {
				ans = []int{r - l + 1, l, r}
			}

			if s[l] == '1' {
				count--
			}

			l++
		}
	}

	if ans[0] == math.MaxInt {
		return ""
	}

	return s[ans[1] : ans[2]+1]
}

/* 1234. 替换子串得到平衡字符串 */
func balancedString(s string) int {
	duplicate := make(map[rune]int, 4)
	needLen, ans := len(s)/4, math.MaxInt

	for _, data := range s {
		duplicate[data]++
	}

	if duplicate[rune('Q')] == needLen &&
		duplicate[rune('W')] == needLen &&
		duplicate[rune('E')] == needLen &&
		duplicate[rune('R')] == needLen {
		return 0

	}

	for l, r := 0, 0; r < len(s); r++ {
		duplicate[rune(s[r])]--

		for duplicate[rune('Q')] <= needLen &&
			duplicate[rune('W')] <= needLen &&
			duplicate[rune('E')] <= needLen &&
			duplicate[rune('R')] <= needLen {
			ans = min(ans, r-l+1)
			duplicate[rune(s[l])]++
			l++
		}
	}

	return ans
}

/* 2875. 无限数组的最短子数组 */
func minSizeSubarray(nums []int, target int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	circle := target / totalSum * len(nums)
	target = target % totalSum

	sum, ans := 0, math.MaxInt
	nums = append(nums, nums...)
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]

		for sum > target {
			sum -= nums[l]
			l++
		}

		if sum == target {
			ans = min(ans, r-l+1)
		}
	}

	if ans == math.MaxInt {
		return -1
	}

	return circle + ans
}

/* 76. 最小覆盖子串 */
func minWindow(s string, t string) string {
	need, valid := make(map[rune]int), 0
	for _, b := range t {
		need[rune(b)]++
	}

	total, ans := make(map[rune]int), [2]int{0, math.MaxInt}
	for l, r := 0, 0; r < len(s); r++ {
		total[rune(s[r])]++
		if total[rune(s[r])] == need[rune(s[r])] {
			valid++
		}

		for valid == len(need) && l < len(s) {
			len := min(ans[1]-ans[0], r-l+1)
			if len == r-l+1 {
				ans = [2]int{l, r + 1}
			}

			if total[rune(s[l])] == need[rune(s[l])] {
				valid--
			}
			total[rune(s[l])]--
			l++
		}
	}

	if ans[1] == math.MaxInt {
		return ""
	}

	return s[ans[0]:ans[1]]
}

/* 632. 最小区间 */
func smallestRange(nums [][]int) []int {
	type pair struct{ x, i int }
	pairs := []pair{}
	for i, arr := range nums {
		for _, x := range arr {
			pairs = append(pairs, pair{x, i})
		}
	}
	slices.SortFunc(pairs, func(a, b pair) int { return a.x - b.x })
	ans := [2]pair{
		{x: pairs[0].x, i: 0},
		{x: pairs[len(pairs)-1].x, i: len(pairs) - 1},
	}

	valid := 0
	duplicate := make(map[int]int, len(nums))
	for l, r := 0, 0; r < len(pairs); r++ {
		duplicate[pairs[r].i]++
		if duplicate[pairs[r].i] == 1 {
			valid++
		}

		for valid == len(nums) && l < len(pairs) {
			need := false
			if ans[1].x-ans[0].x == pairs[r].x-pairs[l].x {
				if pairs[l].x < ans[0].x {
					need = true
				}
			} else {
				minLen := min(ans[1].x-ans[0].x, pairs[r].x-pairs[l].x)
				if minLen == pairs[r].x-pairs[l].x {
					need = true
				}
			}

			if need {
				ans = [2]pair{
					{x: pairs[l].x, i: l},
					{x: pairs[r].x, i: r},
				}
			}

			if duplicate[pairs[l].i] == 1 {
				valid--
			}
			duplicate[pairs[l].i]--
			l++
		}
	}

	return []int{ans[0].x, ans[1].x}
}
