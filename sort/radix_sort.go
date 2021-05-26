package sort

// RadixSort 基数排序
func RadixSort(arr *[]int, maxNumber int) {
    buckets := make([][]int, 10)

    mod, dev := 10, 1
    for i := 0; i < maxNumber; i, dev, mod = i+1, dev*10, mod*10 {
        for j := 0; j < len(*arr); j++ {
            bkPos := (*arr)[j] % mod / dev
            buckets[bkPos] = append(buckets[bkPos], (*arr)[j])
        }

        var pos int
        for j := 0; j < len(buckets);j++ {
            for bj := range buckets[j] {
                v := buckets[j][bj]
                (*arr)[pos] = v
                pos++
            }
            if len(buckets[j]) > 0 {
                buckets[j] = []int{}
            }
        }
    }
}
