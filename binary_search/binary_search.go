// Package binary_search
// FindIndex 查找一个数的索引
// FindLeftBoundIndexFirst 优先查找左边界索引
// FindRightBoundIndexFirst 优先查找右边界索引
package binary_search

// FindIndex 查找一个数的索引
// arr 有序数组
// target 查找目标
func FindIndex(arr []int, target int) int {
    left, right := 0, len(arr)-1

    for left <= right {
        mid := left + (right-left)/2
        if target == arr[mid] {
            // 找到目标
            return mid
        } else if target > arr[mid] {
            // 如果目标大于中位数
            left = mid + 1
        } else if target < arr[mid] {
            // 如果目标小于中位数
            right = mid - 1
        }
    }

    // 遍历完了仍然没有找到
    return -1
}

// FindLeftBoundIndexFirst 优先查找左边界索引
// 例如arr=[1,2,2,2,3,4,5]，target=2
// 优先返回索引1
// arr 有序数组
// target 查找目标
func FindLeftBoundIndexFirst(arr []int, target int) int {
    left, right := 0, len(arr) - 1

    for left <= right {
        mid := left + (right - left) / 2
        if target == arr[mid] {
            // 锁定左边界，右边界移动
            right = mid - 1
        } else if target > arr[mid] {
            left = mid + 1
        } else if target < arr[mid] {
            right = mid - 1
        }
    }

    // 越界检查
    if left >= len(arr) || arr[left] != target {
        return -1
    }

    return left
}

// FindRightBoundIndexFirst 优先查找右边界索引
// 例如arr=[1,2,2,2,3,4,5]，target=2
// 优先返回索引1
// arr 有序数组
// target 查找目标
func FindRightBoundIndexFirst(arr []int, target int) int {
    left, right := 0, len(arr) - 1

    for left <= right {
        mid := left + (right - left) / 2
        if target == arr[mid] {
            // 锁定右边界，左边界移动
            left = mid + 1
        } else if target > arr[mid] {
            left = mid + 1
        } else if target < arr[mid] {
            right = mid - 1
        }
    }

    // 越界检查
    if right < 0 || arr[right] != target {
        return -1
    }

    return right
}
