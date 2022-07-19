package solve0121

/**
 * @index 121
 * @title 买卖股票的最佳时机
 * @difficulty 简单
 * @tags array,dynamic-programming
 * @draft false
 * @link https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
 * @frontendId 121
 */

func maxProfit(prices []int) int {
	profit := 0
	minPrice := int(1e9)
	// 获取历史上最低价格，同时维护最大值不停的更新最大值
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}

func maxProfit_(prices []int) int {
	minPrice := make([]int, 1000000)
	minPrice[0] = prices[0]
	// 获取所有价格前面的价格最小值
	for i := 1; i < len(prices); i++ {
		if minPrice[i-1] < prices[i] {
			minPrice[i] = minPrice[i-1]
		} else {
			minPrice[i] = prices[i]
		}
	}
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-minPrice[i] > profit {
			profit = prices[i] - minPrice[i]
		}
	}
	return profit
}

// maxProfit 超时
func maxProfitTimeout(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		tempMax := 0
		for j := 0; j < i; j++ {
			if prices[i]-prices[j] > 0 && prices[i]-prices[j] > tempMax {
				tempMax = prices[i] - prices[j]
			}
		}
		if profit < tempMax {
			profit = tempMax
		}
	}
	return profit
}
