package sliding_window

const IntMax = int(^uint(0) >> 1)

// MinWindowSubStr 最小覆盖字串
// 输入字符串S，输入字串T
// 在S中找出包含所有T字母的最小字串
func MinWindowSubStr(source, target string) string {
    // 初始化需求计数器
    need := make(map[rune]int, len(target))
    for _, v := range target {
        need[v]++
    }

    // 初始化窗口计数器
    window := make(map[rune]int)
    // 窗口[left, right)
    left, right := 0, 0
    // 满足need条件的字符个数
    var valid int
    // 最小字串的起始索引和长度
    start, sLen := 0, IntMax

    // 转化为rune方便取字符
    s := []rune(source)
    for right < len(s) {
        // 获取当前字符
        var c = s[right]
        // 右扩窗口
        right++

        // 当需要当前字符时
        if _, ok := need[c]; ok {
            window[c]++
            // 当已经满足当前字符当需求时，valid计数加一
            if window[c] == need[c] {
                valid++
            }
        }

        // 当Target串所有字符都包含
        // 准备左缩窗口
        for valid == len(need) {
            // 先记录或更新最小覆盖串的索引
            // 此时的满足结果是 s[start : start+sLen]
            if right - left < sLen {
                start = left
                sLen = right - left
            }
            // 待移除的字符
            var d = s[left]
            // 左缩窗口
            left++
            // 更新窗口数据
            if _, ok := need[d]; ok {
                if window[d] == need[d] {
                    valid--
                }
                window[d]--
            }
        }
    }

    if sLen == IntMax {
        return ""
    }

    return string(s[start:start+sLen])
}
