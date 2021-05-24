package sort

// QuickSort 快速排序
func QuickSort(arr *[]int, left, right int) {
    l, r := left, right

    // 支点
    pivot := (*arr)[(l+r)/2]

    for l < r {
        // 比支点小的数
        for pivot > (*arr)[l] {
            l++
        }

        // 比支点大的数
        for pivot < (*arr)[r] {
            r--
        }

        // arr[r] <= pivot <= arr[l
        // 交换位置
        if l <= r {
            (*arr)[l], (*arr)[r] = (*arr)[r], (*arr)[l]
            l++
            r--
        }
    }

    // 左边继续排序，直到左边剩下一个数
    if left < r {
        QuickSort(arr, left, r)
    }

    // 右边继续排序，直到右边剩下一个数
    if right > l {
        QuickSort(arr, l, right)
    }
}
