package slidingwindow

import (
	"bytes"
	"math"
)

/* 1456. 定长子串中元音的最大数目 */
func maxVowels(s string, k int) int {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	count, ans := 0, 0
	for l, r := 0, 0; r < len(s); r++ {
		if bytes.ContainsRune(vowels, rune(s[r])) {
			count++
		}

		if r-l+1 == k {
			ans = max(ans, count)
			if ans == k {
				break
			}

			if bytes.ContainsRune(vowels, rune(s[l])) {
				count--
			}
			l++
		}
	}

	return ans
}

/* 643. 子数组最大平均数 I */
func findMaxAverage(nums []int, k int) float64 {
	ans, sum := -math.MaxFloat64, 0
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]

		if r-l+1 == k {
			ans = max(ans, float64(sum)/float64(k))

			sum -= nums[l]
			l++
		}
	}

	return ans
}

/* 1343. 大小为 K 且平均值大于等于阈值的子数组数目 */
func numOfSubarrays(arr []int, k int, threshold int) int {
	sum, count := 0, 0
	for l, r := 0, 0; r < len(arr); r++ {
		sum += arr[r]

		if r-l+1 == k {
			if threshold*k <= sum {
				count++
			}

			sum -= arr[l]
			l++
		}
	}

	return count
}

/* 2090. 半径为 k 的子数组平均值 */
func getAverages(nums []int, k int) []int {
	avg := make([]int, len(nums))
	for i := range avg {
		avg[i] = -1
	}

	sum := 0
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]

		if r-l == 2*k {
			avg[(l+r)/2] = sum / (2*k + 1)

			sum -= nums[l]
			l++
		}
	}

	return avg
}

/* 2379. 得到 K 个黑块的最少涂色次数 */
func minimumRecolors(blocks string, k int) int {
	count, ans := 0, math.MaxInt
	for l, r := 0, 0; r < len(blocks); r++ {
		if blocks[r] == 'W' {
			count++
		}

		if r-l+1 == k {
			ans = min(ans, count)

			if blocks[l] == 'W' {
				count--
			}
			l++
		}
	}

	return ans
}

/* 2841. 几乎唯一子数组的最大和 */
func maxSum(nums []int, m int, k int) int64 {
	duplicate := make(map[int]int, len(nums))
	sum, ans := 0, 0
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		duplicate[nums[r]]++

		if r-l+1 == k {
			if len(duplicate) >= m {
				ans = max(ans, sum)
			}

			sum -= nums[l]
			duplicate[nums[l]]--
			if duplicate[nums[l]] == 0 {
				delete(duplicate, nums[l])
			}
			l++
		}
	}

	return int64(ans)
}

/* 2461. 长度为 K 子数组中的最大和 */
func maximumSubarraySum(nums []int, k int) int64 {
	duplicate := make(map[int]int, len(nums))
	sum, ans := 0, 0
	for l, r := 0, 0; r < len(nums); r++ {
		sum += nums[r]
		duplicate[nums[r]]++

		if r-l+1 == k {
			if len(duplicate) == k {
				ans = max(ans, sum)
			}

			sum -= nums[l]
			duplicate[nums[l]]--
			if duplicate[nums[l]] == 0 {
				delete(duplicate, nums[l])
			}
			l++
		}
	}

	return int64(ans)
}

/* 1423. 可获得的最大点数 */
func maxScore(cardPoints []int, k int) int {
	sum, ans := 0, math.MaxInt
	total := 0
	for l, r := 0, 0; r < len(cardPoints); r++ {
		sum += cardPoints[r]
		total += cardPoints[r]

		if r-l+1 == len(cardPoints)-k {
			ans = min(sum, ans)

			sum -= cardPoints[l]
			l++
		}
	}

	if len(cardPoints)-k == 0 {
		ans = 0
	}

	return total - ans
}
