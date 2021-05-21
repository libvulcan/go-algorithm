package back_track

// KNumberCombine 1～N中, K个数字的所有组合
// K 相当于多叉树的树高
func KNumberCombine(N, K int) [][]int {
	if K == 0 || N <= 0 {
		return nil
	}

	var res [][]int
	var track []int

	// 从1开始, 所以curr=1
	KNumberCombineBackTrack(&res, 1, N, K, track)

	return res
}

func KNumberCombineBackTrack(res *[][]int, curr, N, K int, track []int) {
	// track中数元素为K个，到达树底
	if len(track)  == K {
		*res = append(*res, append([]int{}, track...))
		return
	}

	for i := curr; i <= N; i++ {
		track = append(track, i)
		KNumberCombineBackTrack(res, i+1, N, K, track)
		track = track[:len(track)-1]
	}
}
