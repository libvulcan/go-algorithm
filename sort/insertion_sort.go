package sort

// InsertionSort 插入排序
func InsertionSort(arr *[]int) {
    // 上一个元素索引
    var preIdx int
    // 当前元素指针
    var curr int

    // 初始arr[:1]已排序序列
    // arr[1:]未排序序列
    for i := 1; i < len(*arr); i++ {
        preIdx = i - 1
        curr = (*arr)[i]

        // 如果当前数小于前一个数
        for preIdx >= 0 && (*arr)[preIdx] > curr {
            // 从后往前遍历arr[:i]，不断后移序列
            // 直到arr[prefix] < curr，完成遍历
            (*arr)[preIdx+1] = (*arr)[preIdx]
            preIdx--
        }
        // 如果数组未发生变化，即为arr[preIdx+1] = arr[i] = curr，数组无变化
        // 如果数组发生变化, preIdx = preIdx - n
        // 最后一轮preIdx多减了1，所以插入curr到preIdx + 1位置
        (*arr)[preIdx+1] = curr
    }
}
