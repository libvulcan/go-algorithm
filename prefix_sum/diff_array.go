// Package prefix_sum
// DifferenceArray.InitDiffArray 初始化差分数组
// DifferenceArray.ReduceArray 还原差分数组到原始数组
// DifferenceArray.IncrBetween 给区间增加val
// CorpFlightBookings 差分问题：航班预定统计
package prefix_sum

type DifferenceArray struct {
    diff []int
}

// InitDiffArray 差分数组
func (arr *DifferenceArray) InitDiffArray(source []int) []int {
    if len(source) == 0 {
        return []int{}
    }
    arr.diff = make([]int, len(source))
    arr.diff[0] = source[0]
    for i := 1; i < len(source); i++ {
        arr.diff[i] = source[i] - source[i-1]
    }
    return arr.diff
}

// ReduceArray 还原差分数组到原始数组
func (arr *DifferenceArray) ReduceArray() []int {
    if len(arr.diff) == 0 {
        return []int{}
    }
    res := make([]int, len(arr.diff))
    res[0] = arr.diff[0]
    for i := 1; i < len(arr.diff); i++ {
        res[i] = arr.diff[i] + res[i-1]
    }
    return res
}

// IncrBetween 给区间增加val
func (arr *DifferenceArray) IncrBetween(i, j, val int) {
    arr.diff[i] += val
    if j+1 < len(arr.diff) {
        arr.diff[j+1] -= val
    }
}

// CorpFlightBookings 航班预定统计
// bookings[x] = 第x条记录表示：[i,j,k] 从i到j到每个航班预定了k个座位
// n 共n个航班，返回一个长度为n的数组，按航班序号顺序返回每个航班的预定数
// eg: bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
func CorpFlightBookings(bookings [][]int, n int) []int {
    nums := make([]int, n)
    df := &DifferenceArray{}
    df.InitDiffArray(nums)

    for _, v := range bookings {
        var i = v[0] - 1
        var j = v[1] - 1
        var val = v[2]
        df.IncrBetween(i, j, val)
    }

    return df.ReduceArray()
}
