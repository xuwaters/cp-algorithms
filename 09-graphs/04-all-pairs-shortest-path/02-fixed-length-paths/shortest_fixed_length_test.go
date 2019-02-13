package fixed

import (
	"testing"
)

func TestShortestFixedLength(t *testing.T) {
	const F = -1
	dataList := []struct {
		k   int
		g   [][]int
		ans [][]int
	}{
		{
			k: 1,
			g: [][]int{
				{F, 2, 3, 1},
				{2, F, 1, F},
				{3, 1, F, 5},
				{1, F, 5, F},
			},
			ans: [][]int{
				{F, 2, 3, 1},
				{2, F, 1, F},
				{3, 1, F, 5},
				{1, F, 5, F},
			},
		},
		{
			k: 2,
			g: [][]int{
				{F, 2, 3, 1},
				{2, F, 1, F},
				{3, 1, F, 5},
				{1, F, 5, F},
			},
			ans: [][]int{
				{2, 4, 3, 8},
				{4, 2, 5, 3},
				{3, 5, 2, 4},
				{8, 3, 4, 2},
			},
		},
		{
			k: 5,
			g: [][]int{
				{F, 2, 3, 1},
				{2, F, 1, F},
				{3, 1, F, 5},
				{1, F, 5, F},
			},
			ans: [][]int{
				{8, 6, 7, 5},
				{6, 8, 5, 7},
				{7, 5, 8, 6},
				{5, 7, 6, 8},
			},
		},
	}

	for _, data := range dataList {
		got := ShortestFixedLength(data.g, data.k)
		if matEquals(got, data.ans) {
			t.Logf(" OK: k=%d, got: \n%s", data.k, matPrint(got))
		} else {
			t.Fatalf("ERR: k=%d, got: \n%s\nans:\n%s", data.k, matPrint(got), matPrint(data.ans))
		}
	}
}
