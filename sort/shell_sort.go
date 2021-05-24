package sort

// ShellSort 希尔排序
// 参考插入排序 InsertionSort，但是事先分了多组序列并在同一趟进行插入排序
func ShellSort(arr *[]int, scale int) {
    // arr拆分为step个数组
    // 每一轮的第n个数组arr_n =
    // (arr[n-1], arr[n-1 + step], ..., arr[i + (n-1) * step])
    for step := len(*arr) / scale; step > 0; step /= scale {
        // 每一轮减少step，直到step == 1
        for i := step; i < len(*arr); i++ {
            var j = i
            var arrI = (*arr)[i]

            for j-step >= 0 && arrI < (*arr)[j-step] {
                (*arr)[j] = (*arr)[j-step]
                j -= step
            }
            (*arr)[j] = arrI
        }
        // fmt.Println(*arr)
    }
}
