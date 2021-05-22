package sliding_window

// FindAllAnagramsInAString 找出所有字母的异位词
// 输入一个字符串S，输入的P，在S中找子串，子串字母与S相同，但是排列不一样
// 也就是还是根据目标串从源串找子排列的问题
func FindAllAnagramsInAString(source, pattern string) []int {
    // 初始化字母需求计数器
    need := make(map[rune]int)
    for _, v := range pattern {
        need[v]++
    }

    // 初始化窗口计数器
    window := make(map[rune]int)

    // 窗口区间[left, right)
    left, right := 0, 0
    // 所有字母需求（==len(need)）达成计数
    var valid int

    // 结果数组
    var res []int

    s := []rune(source)
    // 开始滑动右边窗口
    for right < len(s) {
        var c = s[right]
        right++
        if _, ok := need[c]; ok {
            window[c]++
            // 满足当前字符需求
            if window[c] == need[c] {
                valid++
            }
        }

        for right - left >= len(pattern) {
            if valid == len(need) {
                res = append(res, left)
            }
            var d = s[left]
            left++
            if _, ok := need[d]; ok {
                if window[d] == need[d] {
                    valid--
                }
                window[d]--
            }
        }
    }
    return res
}
