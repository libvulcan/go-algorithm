package sort

// SelectionSort 选择排序
// 有序区arr[:i], 无序区arr[i:]
func SelectionSort(arr *[]int) {
    // 最小数的索引
    var minIndex int

    for i := 0; i < len(*arr)-1; i++ {
        // 初始化
        minIndex = i
        // fmt.Println(minIndex, (*arr)[:minIndex], (*arr)[minIndex:])
        for j := i; j < len(*arr); j++ {
            if (*arr)[j] < (*arr)[minIndex] {
                minIndex = j
            }
        }
        // 交换
        (*arr)[i], (*arr)[minIndex] = (*arr)[minIndex], (*arr)[i]
    }
}
