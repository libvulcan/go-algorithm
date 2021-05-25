package sort

// BucketSort 桶排序
func BucketSort(arr *[]int, bucketSize int) {
    if len(*arr) < 2 {
        return
    }

    var maxVal = (*arr)[0]
    var minVal = (*arr)[0]

    // 找到最大最小值
    for _, v := range *arr {
        if v < minVal {
            minVal = v
        } else if v > maxVal {
            maxVal = v
        }
    }

    // 创建桶
    buckets := make([][]int, (maxVal-minVal)/bucketSize+1)

    // 分配到桶中
    for _, v := range *arr {
        buckets[(v-minVal)/bucketSize] = append(buckets[(v-minVal)/bucketSize], v)
    }

    var arrIdx = 0
    for i, v := range buckets {
        InsertionSort(&buckets[i])
        for _, vj := range v {
            (*arr)[arrIdx] = vj
            arrIdx++
        }
    }
}
