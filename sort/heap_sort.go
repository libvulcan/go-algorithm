package sort

type Heap struct {
    // 数组可以表示堆结构
    data []int
}

// Sort 堆排序
func (heap *Heap) Sort() {
    // 构建顶堆
    for i := len(heap.data)/2 - 1; i >= 0; i-- {
        heap.Adjust(i, len(heap.data))
    }

    for j := len(heap.data) - 1; j > 0; j-- {
        heap.data[0], heap.data[j] = heap.data[j], heap.data[0]
        heap.Adjust(0, j)
    }
}

// Adjust 调整
func (heap *Heap) Adjust(i int, length int) {
    // 当前元素arr[i]
    var arrI = heap.data[i]
    // i节点的左子节点i*2+1节点
    for k := i*2 + 1; k < length; k = k*2 + 1 {
        // 左子节点小于右子节点，k指向右节点
        if k+1 < length && heap.data[k] < heap.data[k+1] {
            k++
        }
        // 子节点大于父节点，子节点赋值给父节点
        if heap.data[k] > arrI {
            heap.data[i] = heap.data[k]
            i = k
        }
    }
    // 当前元素放到最终位置
    heap.data[i] = arrI
}
