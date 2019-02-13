package floyd

import "testing"

func TestFloyd(t *testing.T) {
	const F = INF
	dataList := []struct {
		d   [][]int
		ans [][]int
	}{
		{
			d: [][]int{
				//0,1, 2, 3, 4, 5, 6, 7
				{0, 2, 4, 6, 1, F, F, F}, // 0
				{2, 0, 3, F, F, F, F, F}, // 1
				{4, 3, 0, 1, F, F, 1, F}, // 2
				{6, F, 1, 0, 7, F, 1, F}, // 3
				{1, F, F, 7, 0, F, 9, F}, // 4
				{F, F, F, F, F, 0, 3, 4}, // 5
				{F, F, 1, 1, 9, 3, 0, 6}, // 6
				{F, F, F, F, F, 4, 6, 0}, // 7
			},
			ans: [][]int{
				{0, 2, 4, 5, 1, 8, 5, 11},
				{2, 0, 3, 4, 3, 7, 4, 10},
				{4, 3, 0, 1, 5, 4, 1, 7},
				{5, 4, 1, 0, 6, 4, 1, 7},
				{1, 3, 5, 6, 0, 9, 6, 12},
				{8, 7, 4, 4, 9, 0, 3, 4},
				{5, 4, 1, 1, 6, 3, 0, 6},
				{11, 10, 7, 7, 12, 4, 6, 0},
			},
		},
	}

	for _, data := range dataList {
		Floyd(data.d)
		for r := 0; r < len(data.d); r++ {
			if arrayEquals(data.ans[r], data.d[r]) {
				t.Logf(" OK: %2d: %+2v", r, data.d[r])
			} else {
				t.Fatalf("ERR: %2d: %+2v", r, data.d[r])
			}
		}
	}
}

func arrayEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
