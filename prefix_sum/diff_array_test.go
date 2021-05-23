package prefix_sum

import (
    "fmt"
    "reflect"
    "testing"
)

func TestDiffArr(t *testing.T) {
    source := []int{2, 5, 6, 8, 12, 3, 4}

    dif := &DifferenceArray{}

    // 生成差分数组
    diff := dif.InitDiffArray(source)
    fmt.Println(diff)

    // 还原差分数组
    arr := dif.ReduceArray()
    fmt.Println(arr)

    // 是否正确还原
    fmt.Println(reflect.DeepEqual(arr, source))

    // 区间操作
    dif.IncrBetween(0, len(dif.diff), 10)
    fmt.Println(dif.ReduceArray())
}

func TestCorpFlightBookings(t *testing.T) {
    fmt.Println(CorpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))
}
