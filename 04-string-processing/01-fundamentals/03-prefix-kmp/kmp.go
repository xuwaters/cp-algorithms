package kmp

//
// Prefix[i] = max(k: s[0..k-1] = s[i-(k-1)..i], k = 0..i)
//
// length of max prefix that matches the suffix of s[0..i]
//
//
//
//
//

func PrefixFunction(s string) []int {
	n := len(s)
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

//
func ComputeAutomation(s string) [][]int {
	s += "#"
	n := len(s)
	pi := PrefixFunction(s)
	aut := make([][]int, n)
	for i := 0; i < n; i++ {
		aut[i] = make([]int, 26)
		for c := 0; c < 26; c++ {
			if i > 0 && byte(c+'a') != s[i] {
				aut[i][c] = aut[pi[i-1]][c]
			} else {
				aut[i][c] = i
				if byte(c+'a') == s[i] {
					aut[i][c]++
				}
			}
		}
	}
	return aut
}
