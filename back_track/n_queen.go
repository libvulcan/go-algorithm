package back_track

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
