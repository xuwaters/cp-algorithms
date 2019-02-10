package hashing

//
//  string hash:
//    Sum(s[i]*p) % m
//   we can choose p = 31 or 53, and m = 10e9+9 (this is a prime number)
//

// good for lower case letters
func strHash(str string) int {
	const p = 31
	const m = 1000000009
	hash := 0
	factor := 1
	for _, c := range str {
		hash = (hash + int(c-'a'+1)*factor) % m
		factor = (factor * p) % m
	}
	return hash
}

// 
//  fast hash calculation of substrings of given string
// hash(s[i..j]) = sum(s[i]*p^(k-i) % m, k=i..j)
// 
// hash(s[i..j]) * p^i = (hash(s[0..j]) - hash(s[0..i-1])) % m
// 


// TODO: not understand
