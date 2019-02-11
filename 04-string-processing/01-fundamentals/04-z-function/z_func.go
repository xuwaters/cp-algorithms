package z_func

//
// Z Function:
// In other words, z[i] is the length of the longest common prefix between s and the suffix of s starting at i.
// Note. In this article, to avoid ambiguity, we assume 0-based indexes;
// that is: the first character of s has index 0 and the last one has index n−1.
//
// The first element of Z-function, z[0], is generally not well defined.
// In this article we will assume it is zero (although it doesn't change anything in the algorithm implementation).
//
// Examples:
//
// "aaaaa" - [0,4,3,2,1]
// "aaabaab" - [0,2,1,0,2,1,0]
// "abacaba" - [0,0,1,0,3,0,1]
//

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ZFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			// because s[l..r] matches with s[0..r-l]
			// so s[0..i-l] = s[l:i], we can start from z[i-l] for z[i]
			// but only l..r matches, so z[i] length can be exceed r-i+1
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}
