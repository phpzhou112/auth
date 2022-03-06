package alg

import "fmt"

// 01背包问题
func Bag01() {
	var (
		m = 10 // 背包容量
		n = 4  // 物品数量
	)

	var w = []int{0, 2, 3, 4, 7} // 每个物品重量
	var c = []int{0, 1, 3, 5, 9} // 每个物品价值
	//var dp [35][105]int
	//for i := 1; i <= n; i++ {
	//	for j := 1; j <= m; j++ {
	//		if j < w[i] {
	//			dp[i][j] = dp[i-1][j]
	//		} else {
	//			dp[i][j] = Max(dp[i-1][j], dp[i-1][j-w[i]]+c[i])
	//		}
	//	}
	//}
	var dp [35]int
	for i := 1; i <= n; i++ {
		for j := m; j >= 1; j-- {
			if j >= w[i] {
				dp[j] = Max(dp[j], dp[j-w[i]]+c[i])
			}
		}
		for k := 0; k <= m; k++ {
			fmt.Print(dp[k])
			fmt.Print(" ")
		}
		fmt.Println("")
	}

	//for i := 0; i <= n; i++ {
	//	for j := 0; j <= m; j++ {
	//		fmt.Print(dp[i][j])
	//		fmt.Print(" ")
	//	}
	//	fmt.Println("")
	//}
}
