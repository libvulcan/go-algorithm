package back_track

// FullPermutation 全排列
// res 最终所有排列 nums 数组 track 路径
func FullPermutation(res *[][]int, nums, track []int) {
    // 触发结束条件
    if len(nums) == len(track) {
        *res = append(*res, track)
        return
    }

    for _, v := range nums {
        if !Contains(track, v) {
            // 选择
            track = append(track, v)
            // 继续下一步决策
            FullPermutation(res, nums, track)
            // 撤销选择
            track = track[:len(track)-1]
        }
    }
}

// Contains 切片l内是否包含v
func Contains(l []int, v int) bool {
    for _, val := range l {
        if val == v {
            return true
        }
    }
    return false
}
