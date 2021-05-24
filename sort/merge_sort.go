package sort

// MergeSort 归并排序
// 时间复杂度O(nLogN)
// 与 SelectionSort 一样，不受输入数据影响
func MergeSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }

    // 分为两组
    var mid = len(arr) / 2
    left := arr[:mid]
    right := arr[mid:]

    l := MergeSort(left)
    r := MergeSort(right)

    return Merge(l, r)
    // return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left, right []int) []int {
    var res []int

    for len(left) > 0 && len(right) > 0 {
        if left[0] <= right[0] {
            res = append(res, left[0])
            left = left[1:]
        } else {
            res = append(res, right[0])
            right = right[1:]
        }
    }

    if len(left) > 0 {
        // 左边还剩元素
        res = append(res, left...)
    }

    if len(right) > 0 {
        // 右边还剩元素
        res = append(res, right...)
    }

    return res
}