package sort

// CountSort 计数排序
// maxNumber 最大数，通过 MaxNumber 获取
func CountSort(arr *[]int, maxNumber int) {
    // 额外数组空间
    ext := make([]int, maxNumber+1)

    // 遍历数组，开始计数
    for _, v := range *arr {
        ext[v]++
    }

    // 排序
    var arrIdx int
    for i, v := range ext {
        // 如果不为0则存在arr中数i，切共v个i
        for v > 0 {
            (*arr)[arrIdx] = i
            arrIdx++
            v--
        }
    }
}

// MaxNumber 最大数组
func MaxNumber(arr []int) int {
    var tmp = arr[0]
    for _, v := range arr {
        if tmp < v {
            tmp = v
        }
    }
    return tmp
}
