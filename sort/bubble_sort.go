package sort

// BubbleSort 冒泡排序
func BubbleSort(arr *[]int) {
    // 直接遍历
    for i := 0; i < len(*arr); i++ {
        for j := i + 1; j < len(*arr); j++ {
            if (*arr)[i] > (*arr)[j] {
                (*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
            }
        }
    }
}

// BubbleSortWithChangeFlag 冒泡排序优化
// 如果一趟没有发生置换，说明已经排好序了
func BubbleSortWithChangeFlag(arr *[]int)  {
    var isChanged bool
    for i := 0; i < len(*arr); i++ {
        // 新的一趟，初始化标记
        isChanged = false
        for j := i + 1; j < len(*arr); j++ {
            if (*arr)[i] > (*arr)[j] {
                (*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
                // 发生置换，做标记
                isChanged = true
            }
        }
        // 未发生置换
        if !isChanged {
            break
        }
    }
}
