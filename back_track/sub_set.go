// Package back_track
// CanPartitionKSubSet 数组是否能划分为K个元素和相同子集（从数字的视角）
// CanPartitionKSubSetFromBucket 数组是否能划分为K个元素和相同子集（从桶/子集的视角）
package back_track

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
