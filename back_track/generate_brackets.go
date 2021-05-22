package back_track

import (
    "strings"
)

// GenerateBrackets 生成合法的括号组合
// N对括号组成对字符串Pattern且满足合法条件
// 合法条件1: 左括号数量 = 右括号数量
// 合法条件2: 合法Pattern的字串的 左括号数量 >= 右括号数量
func GenerateBrackets(N int) []string {
    var res []string
    BracketsBackTrack(&res, N, N, []string{})
    return res
}

// BracketsBackTrack 回溯
// left 左括号剩余数量
// right 右括号剩余数量
func BracketsBackTrack(res *[]string, left, right int, track []string) {
    // 左括号剩余比较多，不合法
    if left > right {
        return
    }

    // 数量小于0
    if left < 0 || right < 0 {
        return
    }

    // 所有括号正好分配完
    if left == 0 && right == 0 {
        *res = append(*res, strings.Join(track, ","))
        return
    }

    // 放置左括号
    track = append(track, "(")
    BracketsBackTrack(res, left-1, right, track)
    track = track[:len(track)-1]

    // 放置右括号
    track = append(track, ")")
    BracketsBackTrack(res, left, right-1, track)
    track = track[:len(track)-1]
}
