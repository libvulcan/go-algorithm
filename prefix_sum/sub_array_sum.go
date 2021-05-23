// Package prefix_sum
// SubArraySum 求数组中一共有多少个和为K的子数组
// SubArraySumBetter 优化解法
package prefix_sum

// SubArraySum 求数组中一共有多少个和为K的子数组
// arr 数组
func SubArraySum(arr []int, K int) int {
    prefixSumArr := make([]int, len(arr)+1)
    prefixSumArr[0] = 0
    var res int

    // 构造前缀和
    for i, v := range arr {
        prefixSumArr[i+1] = prefixSumArr[i] + v
    }

    for i := 1; i < len(prefixSumArr); i++ {
        for j := 0; j < i; j++ {
            if prefixSumArr[j] == prefixSumArr[i]-K {
                res++
            }
        }
    }

    return res
}

// SubArraySumBetter 优化解法
func SubArraySumBetter(arr []int, K int) int {
    // 前缀和 -> 前缀出现次数
    prefixSumMap := make(map[int]int)
    prefixSumMap[0] = 1
    var sumI, ans int
    for _, v := range arr {
        sumI += v
        var sumJ = sumI - K
        if _, ok := prefixSumMap[sumJ]; ok {
            ans += prefixSumMap[sumJ]
        }
        prefixSumMap[sumI]++
    }

    return ans
}
