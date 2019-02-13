package fixed

import (
	"testing"
)

func TestFixedLengthPaths(t *testing.T) {
	dataList := []struct {
		k   int
		g   [][]int
		ans [][]int
	}{
		{
			k: 0,
			g: [][]int{
				{0, 1, 1, 1},
				{1, 0, 1, 0},
				{1, 1, 0, 1},
				{1, 0, 1, 0},
			},
			ans: [][]int{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
		},
		{
			k: 3,
			g: [][]int{
				{0, 1, 1, 1},
				{1, 0, 1, 0},
				{1, 1, 0, 1},
				{1, 0, 1, 0},
			},
			ans: [][]int{
				{4, 5, 5, 5},
				{5, 2, 5, 2},
				{5, 5, 4, 5},
				{5, 2, 5, 2},
			},
		},
		{
			k: 5,
			g: [][]int{
				{0, 1, 1, 1},
				{1, 0, 1, 0},
				{1, 1, 0, 1},
				{1, 0, 1, 0},
			},
			ans: [][]int{
				{32, 29, 33, 29},
				{29, 18, 29, 18},
				{33, 29, 32, 29},
				{29, 18, 29, 18},
			},
		},
	}

	for _, data := range dataList {
		got := FixedLengthPaths(data.g, data.k)
		if matEquals(got, data.ans) {
			t.Logf("got: \n%s", matPrint(got))
		} else {
			t.Fatalf("got:\n%s\nans:\n%s", matPrint(got), matPrint(data.ans))
		}
	}
}
