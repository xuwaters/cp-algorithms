package suffix_automation

import "testing"

func TestSuffixAutomation(t *testing.T) {
	str := "ababc你好你好abc"
	sa := NewSuffixAutomation()
	for _, chr := range []rune(str) {
		sa.Extend(chr)
	}
	var walk func(root *State, curr []rune)
	walk = func(root *State, curr []rune) {
		t.Logf("s = %s", string(curr))
		for k, v := range root.Next {
			curr = append(curr, k)
			walk(v, curr)
			curr = curr[:len(curr)-1]
		}
	}
	walk(sa.Start, []rune{})
}
