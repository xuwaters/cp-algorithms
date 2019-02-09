package fenwick

import "testing"

func TestFenwickTree(t *testing.T) {

	type Command func(t *FenwickTree) int

	cmdAdd := func(idx int, val int) Command {
		return func(t *FenwickTree) int {
			t.Add(idx, val)
			return 0
		}
	}

	cmdQuery := func(start, end int) Command {
		return func(t *FenwickTree) int {
			return t.SumOfRange(start, end)
		}
	}

	type entry struct {
		ans int
		cmd Command
	}

	dataList := []struct {
		n       int
		entries []entry
	}{
		{
			n: 10,
			entries: []entry{
				{ans: 0, cmd: cmdAdd(0, 1)},
				{ans: 0, cmd: cmdAdd(1, 1)},
				{ans: 0, cmd: cmdAdd(2, 82)},
				{ans: 84, cmd: cmdQuery(0, 2)},
				{ans: 0, cmd: cmdAdd(9, 181)},
				{ans: 0, cmd: cmdAdd(8, 9)},
				{ans: 274, cmd: cmdQuery(0, 9)},
				{ans: 0, cmd: cmdAdd(8, 9)},
				{ans: 283, cmd: cmdQuery(0, 9)},
			},
		},
	}
	for _, data := range dataList {
		f := NewFenwickTree(data.n)
		for _, e := range data.entries {
			got := e.cmd(f)
			if got != e.ans {
				t.Fatalf("ERR: got = %+v, e = %+v", got, e)
			} else {
				t.Logf(" OK: got = %+v, e = %+v", got, e)
			}
		}
	}
}
