package sliding_window

// PermutationInString 字符串排列
// 给定串S1，S2，判断S2是否包含S1的一个排列
// 解法类似 sliding_window.MinWindowSubStr
func PermutationInString(S1, S2 string) string {
    need := make(map[rune]int)
    for _, v := range S1 {
        need[v]++
    }

    // 窗口计数器
    window := make(map[rune]int)

    // 窗口[left,right)
    left, right := 0, 0
    // 满足条件计数
    var valid int

    s2rune := []rune(S2)
    // 开始滑动
    for right < len(s2rune) {
        var c = s2rune[right]
        right++

        if _, ok := need[c]; ok {
            window[c]++
            if window[c] == need[c] {
                valid++
            }
        }

        for right-left >= len(S1) {
            if valid == len(need) {
                return string(s2rune[left:right])
            }

            var d = s2rune[left]
            left++

            if _, ok := need[d]; ok {
                if window[d] == need[d] {
                    valid--
                }
                window[d]--
            }
        }
    }

    return ""
}
