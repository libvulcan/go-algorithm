package sliding_window

import (
    "fmt"
    "testing"
)

func TestPermutationInString(t *testing.T) {
    // 包含: 排列为"ba"
    fmt.Println(PermutationInString("ab", "eidbaoii"))
    // 不包含: 输出空
    fmt.Println(PermutationInString("ab", "eidboaii"))
    // 包含: 排列为"ii"，支持重复字符
    fmt.Println(PermutationInString("ii", "eidboaiiui"))
}
