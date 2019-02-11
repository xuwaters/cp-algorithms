package ac

import "testing"

func TestACMachine(t *testing.T) {

	type query struct {
		text    string
		matches [][2]int
	}

	dataList := []struct {
		patterns []string
		queries  []query
	}{
		{
			patterns: []string{
				"ssissippi",
				"ababab",
				"abab",
				"abc",
				"abcd",
				"abd",
				"ssdssd",
				"sspssd",
				"你好你好",
			},
			queries: []query{
				{
					text:    "mississippimississippi",
					matches: [][2]int{{2, 11}, {13, 22}},
				},
				{
					text:    "abcdssdsspssdabdabababssissippim",
					matches: [][2]int{{0, 3}, {0, 4}, {7, 13}, {13, 16}, {16, 20}, {16, 22}, {22, 31}},
				},
				{
					text:    "ab你好你好ssabdssdssd",
					matches: [][2]int{{2, 6}, {8, 11}, {11, 17}},
				},
			},
		},
	}

	for _, data := range dataList {
		t.Logf(">> start query patterns: %#v", data.patterns)
		ac := NewACMachine()
		for _, pat := range data.patterns {
			ac.AddString(pat)
		}
		ac.BuildLinks()
		for _, q := range data.queries {
			matches := ac.FindAll(q.text)
			strList := []string{}
			queryRunes := []rune(q.text)
			for _, m := range matches {
				strList = append(strList, string(queryRunes[m[0]:m[1]]))
			}
			if !ArrayEquals(matches, q.matches) {
				t.Fatalf("ERR: text = %v, strlist = %#v, matches = %+v", q.text, strList, matches)
			} else {
				t.Logf(" OK: text = %v, strlist = %#v, matches = %+v", q.text, strList, matches)
			}
		}
	}
}

func ArrayEquals(a, b [][2]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if b[i][0] != val[0] || b[i][1] != val[1] {
			return false
		}
	}
	return true
}
