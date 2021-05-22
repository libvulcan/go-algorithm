package double_pointer

// RemoveElement 删除所有值为val的元素
func RemoveElement(nums []int, val int) []int {
    fast, slow := 0, 0
    for fast < len(nums) {
        if nums[fast] != val {
            nums[slow] = nums[fast]
            slow++
        }
        fast++
    }
    return nums[:slow]
}