package sliding_window

// LongestSubstr 最长无重复字符的子串长度
// source 原串
func LongestSubstr(source string) int {
    window := make(map[rune]int)

    left, right := 0, 0
    res := 0
    s := []rune(source)
    for right < len(s) {
        var c = s[right]
        right++
        window[c]++

        for window[c] > 1 {
            var d = s[left]
            left++
            window[d]--
        }
        res = max(res, right-left)
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
