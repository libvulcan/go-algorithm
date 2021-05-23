package math

// GCD 辗转相除法
func GCD(a, b int64) int64 {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

// GCDSub 辗转相减法
func GCDSub(a, b int64) int64 {
    for a != b {
        if a > b {
            a = a - b
        } else {
            b = b - a
        }
    }
    return a
}

// GCDBit 位运算(BUG)
func GCDBit(a, b int64) int64 {
    if a < b {
        return GCDBit(b, a)
    }

    if b == 0 {
        return a
    }

    if (a & 1 == 0) &&  (b & 1 == 0) {
        return GCDBit(a >> 1, b >> 1) << 1
    }

    if a & 1 == 0 {
        return GCDBit(a >> 1, b)
    }

    if b & 1 == 0 {
        return GCDBit(b, a >> 1)
    }

    return GCDBit(AddBit(a, b) >> 1, SubtractBit(a, b) >> 1)
}