package robin_carp

//
// Algorithm:
//   Calculate hash for the pattern s.
//   Calculate hash for all prefix of text t.
//   Compare prefix hash with pattern hash.
//
//
//
//
//
//
//
//
//

// RobinCarp returns all matches of s in t
func RobinCarp(s string, t string) []int {
	const p = 31
	const m = 1000000009

	ns := len(s)
	nt := len(t)
	pLen := ns
	if nt > ns {
		pLen = nt
	}
	pows := make([]int, pLen)
	pows[0] = 1
	for i := 1; i < pLen; i++ {
		pows[i] = (pows[i-1] * p) % m
	}

	// here ht[i] = hash(s[0..i-1])
	ht := make([]int, nt+1)
	for i := 0; i < nt; i++ {
		ht[i+1] = (ht[i] + int(t[i]-'a'+1)*pows[i]) % m
	}
	sHash := 0
	for i := 0; i < ns; i++ {
		sHash = (sHash + int(s[i]-'a'+1)*pows[i]) % m
	}

	occurences := []int{}
	for i := 0; i+ns-1 < nt; i++ {
		currHash := (ht[i+ns] + m - ht[i]) % m
		if currHash == (sHash*pows[i])%m {
			// TODO: why not verify string really equals?
			if t[i:i+ns] == s {
				occurences = append(occurences, i)
			}
		}
	}
	return occurences
}
