package disjoint_set

import "testing"

func TestOfflineRMQ(t *testing.T) {
	dataList := []struct {
		arr     []int
		queries [][2]int
		ans     []int
	}{
		{
			arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			queries: [][2]int{
				{1, 14}, {2, 14}, {3, 14}, {4, 7},
			},
			ans: []int{
				1, 2, 3, 4,
			},
		},
		{
			arr: []int{0, 1, 2, 3, 4, 2, 6, 7, 1, 9, 10, 11, 12, 4, 14},
			queries: [][2]int{
				{1, 14}, {2, 14}, {3, 14}, {4, 7}, {7, 12}, {10, 13},
			},
			ans: []int{
				1, 1, 1, 2, 1, 4,
			},
		},
	}

	for _, data := range dataList {
		o := NewOfflineRMQ(data.arr, data.queries)
		got := o.Solve()
		if arrayEqual(got, data.ans) {
			t.Logf(" OK: query ok, got = %+v, data = %+v", got, data)
		} else {
			t.Fatalf("ERR: query ok, got = %+v, data = %+v", got, data)
		}
	}
}

func arrayEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
