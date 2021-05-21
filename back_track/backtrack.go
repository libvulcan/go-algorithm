// Package back_track
// 回溯算法:
// 本质多叉树遍历
// 基本模版:
// var res []interface{}
// func BackTrack(路径track, 选择列表) {
//     if 满足结束条件 {
//         res = append(res, track)
//         return
//     }
//    for 选择 := range 选择列表 {
//        做选择
//        BackTrack(路径track, 选择列表)
//        撤销选择
//    }
// }
// 回溯算法 - 问题列表:
// FullPermutation 全排列问题
// NQueen N皇后问题
// CanPartitionKSubSet 数组是否能划分为K个元素之和相同子集（从数字的视角）
// CanPartitionKSubSetFromBucket 数组是否能划分为K个元素之和相同子集（从桶/子集的视角）
package back_track

// FullPermutation 全排列
// res 最终所有排列 nums 数组 track 路径
func FullPermutation(res *[][]int, nums, track []int) {
	// 触发结束条件
	if len(nums) == len(track) {
		*res = append(*res, track)
		return
	}

	for _, v := range nums {
		if !Contains(track, v) {
			// 选择
			track = append(track, v)
			// 继续下一步决策
			FullPermutation(res, nums, track)
			// 撤销选择
			track = track[:len(track)-1]
		}
	}
}

// Contains 切片l内是否包含v
func Contains(l []int, v int) bool {
	for _, val := range l {
		if val == v {
			return true
		}
	}
	return false
}

// NQueen n皇后棋盘摆放问题，NxN棋盘，N个皇后
// res 所有摆放方式 track摆放方式 N 边长
// 如果只需要知道其中一种解法，第一次append后直接return true
func NQueen(res *[][][]int, track [][]int, N int) {
	if N == -1 {
		*res = append(*res, track)
		return
	}

	rows := track[N]
	for i := range rows {
		// 第N行, 第i列
		if IsValid(track, N, i) {
			// 放置皇后
			track[N][i] = 1
			// 继续下一步放置
			NQueen(res, track, N-1)
			// 拿回皇后
			track[N][i] = 0
		}
	}
}

func IsValid(track [][]int, row, col int) bool {
	n := len(track)
	// 检查竖直（列）方向是否有皇后
	for i := row + 1; i < n; i++ {
		if track[i][col] == 1 {
			return false
		}
	}

	// 检查左下（/）方向是否有皇后
	// (x, y) -> (x-1, y+1)
	for x, y := col-1, row+1; x >= 0 && y < n; x, y = x-1, y+1 {
		if track[y][x] == 1 {
			return false
		}
	}

	// 检查右下（\）方向是否有皇后
	// (x, y) -> (x+1, y+1)
	for x, y := col+1, row+1; x < n && y < n; x, y = x+1, y+1 {
		if track[y][x] == 1 {
			return false
		}
	}
	return true
}

// CanPartitionKSubSet 数组是否能划分为K个元素和相同子集
// nums 数组 K 划分的子集个数
func CanPartitionKSubSet(nums []int, K int) bool {
	var total int
	for _, v := range nums {
		total += v
	}

	// 每个子集的和
	totalPerSet := total / K

	// K个桶（子集）
	bucket := make([]int, K)

	return BackTrackWithNumber(nums, bucket, 0, totalPerSet)

}

func BackTrackWithNumber(nums, bucket []int, curr, totalPerSet int) bool {
	if curr == len(nums) {
		return IsAllElementsSame(bucket, totalPerSet)
	}

	for i := range bucket {
		if bucket[i]+nums[curr] <= totalPerSet {
			bucket[i] += nums[curr]
			// 立即返回
			if BackTrackWithNumber(nums, bucket, curr+1, totalPerSet) {
				return true
			}
			bucket[i] -= nums[curr]
		}
	}

	return false
}

func IsAllElementsSame(l []int, target int) bool {
	for _, v := range l {
		if v != target {
			return false
		}
	}

	return true
}

// CanPartitionKSubSetFromBucket 数组是否能划分为K个元素和相同子集（从桶/子集的视角）
func CanPartitionKSubSetFromBucket(nums []int, K int) bool {
	var total int
	for _, v := range nums {
		total += v
	}
	if total % K != 0 {
		return false
	}

	totalPerSet := total / K

	used := make([]bool, len(nums))

	return BackTrackWithBucket(K, 0, nums, used, 0, totalPerSet)
}

func BackTrackWithBucket(K, bucket int, nums []int, used []bool, curr, totalPerSet int) bool {
	// 所有桶都装满了
	if K == 0 {
		return true
	}

	// 当前桶装满了
	if bucket == totalPerSet {
		return BackTrackWithBucket(K-1, 0, nums, used, 0, totalPerSet)
	}

	for i := curr; i < len(nums); i++ {
		// nums[i]已经被使用了
		if used[i] {
			continue
		}

		// 当前桶剩余量装不下nums[i]了
		if nums[i] + bucket > totalPerSet {
			continue
		}

		// 装桶
		used[i] = true
		bucket += nums[i]

		// 递归下一个数字是否装入当前桶
		if BackTrackWithBucket(K, bucket, nums, used, curr+1, totalPerSet) {
			return true
		}

		// 撤销装桶
		used[i] = false
		bucket -= nums[i]
	}

	return false
}
