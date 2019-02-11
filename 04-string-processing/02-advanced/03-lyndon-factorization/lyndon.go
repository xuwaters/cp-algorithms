package lyndon

// TODO: not understand

func LyndonFactorization(s string) []string {
	// Duval algorithm
	n := len(s)
	i := 0
	factorization := []string{}
	for i < n {
		j := i + 1
		k := i
		for j < n {
			if s[k] <= s[j] {
				break
			} else if s[k] < s[j] {
				k = i
				j++
			} else { // s[k] > s[j]
				k++
				j++
			}
		}
		for i <= k {
			factorization = append(factorization, s[i:i+j-k])
			i += (j - k)
		}
	}
	return factorization
}
