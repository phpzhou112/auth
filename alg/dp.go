package alg

import (
	"fmt"
	"math"
)

/**
  解决动态规划问题
*/

// 1. 硬币兑换问题

/**
	给定不同面额的硬币 coins 和一个总金额 amount。
编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
如果没有任何一种硬币组合能组成总金额，返回 -1。
https://leetcode-cn.com/problems/gaM7Ch/
*/

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func CoinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return 0
	}
	var maxValue = math.MaxInt64 / 4
	var dp = make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = maxValue
	}

	for i := 0; i < amount; i++ {
		for _, y := range coins {
			if i+y >= 0 && i+y < amount+1 && y <= amount {
				dp[i+y] = Min(dp[i+y], dp[i]+1)
			}
		}
	}
	fmt.Println(dp)
	if dp[amount] >= maxValue {
		return -1
	}
	return dp[amount]
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func LargeNumberMultiplication(a string, b string) (reuslt string) {
	a = Reverse(a)
	b = Reverse(b)
	c := make([]byte, len(a)+len(b))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			c[i+j] += (a[i] - '0') * (b[j] - '0')
		}
	}

	var plus byte = 0
	for i := 0; i < len(c); i++ {
		if c[i] == 0 {
			break
		}
		temp := c[i] + plus
		plus = 0
		if temp > 9 {
			plus = temp / 10
			reuslt += string(temp - plus*10 + '0')
		} else {
			reuslt += string(temp + '0')
		}

	}
	return Reverse(reuslt)
}

func FindSubT(arr []int) int {
	l, r := 0, len(arr)

	for l < r {
		m := l + ((r - l) >> 1)
		if arr[m] >= m {
			r = m
		} else {
			l = m + 1
		}
	}
	fmt.Println(l, r)
	if arr[l] == l {
		return l
	}
	return -1
}

// 计算最大子序列之和
func ComputeSeqSum(arr []int) int {

	if len(arr) == 0 {
		return 0
	}
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	ret := dp[0]
	i := 0
	for i = 1; i < len(arr); i++ {
		dp[i] = Max(arr[i], dp[i-1]+arr[i])
		ret = Max(ret, dp[i])
	}
	return ret
}

func ComputeSeqSum3(arr []int) int {

	if len(arr) == 0 {
		return 0
	}
	maxV := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]+arr[i-1] > arr[i] {
			arr[i] += arr[i-1]
		}
		if arr[i] > maxV {
			maxV = arr[i]
		}
	}
	return maxV
}

func ComputeSeqSum2(arr []int) int {

	if len(arr) == 0 {
		return 0
	}
	sum := 0
	size := len(arr)
	ret := math.MinInt64
	for i := 0; i < size; i++ {
		sum += arr[i]
		ret = Max(ret, sum)
		if sum < 0 {
			sum = 0
		}
	}
	return ret
}

// 斐波那契数
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func Fib2(n int) int {
	if n <= 1 {
		return n
	}
	p := 0
	q := 0
	r := 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func Rob(nums []int) int {
	N := len(nums)
	if N <= 0 {
		return 0
	}

	if N == 1 {
		return nums[0]
	}
	dp := make([]int, N)
	dp[0] = Max(0, nums[0])
	dp[1] = Max(0, Max(nums[0], nums[1]))

	for i := 2; i < N; i++ {
		dp[i] = Max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[N-1]
}

func _rob(nums []int) int {
	first, second := nums[0], Max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, Max(first+v, second)
	}
	return second
}

func Rob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return Max(nums[0], nums[1])
	}
	return Max(Rob(nums[:n-1]), Rob(nums[1:]))
}
