package math

func AddBit(a, b int64) int64 {
    sum := a ^ b
    carry := (a & b) << 1
    var tmp1, tmp2 int64
    for carry != 0 {
        tmp1, tmp2 = sum, carry
        sum, carry = tmp1 ^ tmp2, (tmp1 & tmp2) << 1
    }
    return sum
}

func SubtractBit(a, b int64) int64 {
    // b取补码（取反 + 1），结果再加a
    return AddBit(a, AddBit(^b, 1))
}