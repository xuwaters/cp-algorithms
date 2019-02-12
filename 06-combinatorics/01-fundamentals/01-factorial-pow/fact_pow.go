package fact_pow

func FactPowPrime(n, k int) int {
	if k <= 1 {
		return -1
	}
	// if k is a prime number
	// calculate max x such that k^x divides n!
	x := 0
	for n > 0 {
		n /= k
		x += n
	}
	return x
}

// if k = k1^p1 * k2^p2 * ... * km^pm
// how to calculate?
// foreach ki, calculate xi, and then
// x = min(xi / pi, i = 1..m)

func FactPow(n, k int) int {
	f := Factors(k)
	x := -1
	for i := 0; i < len(f); i++ {
		xi := FactPowPrime(n, f[i][0]) / f[i][1]
		if x < 0 || xi < x {
			x = xi
		}
	}
	return x
}

// TODO: this factorization is slow, use prime table to accelerate
func Factors(k int) [][2]int {
	f := make([][2]int, 0)
	for i := 2; i*i <= k; i++ {
		if k%i == 0 {
			cnt := 0
			for k%i == 0 {
				k /= i
				cnt++
			}
			f = append(f, [2]int{i, cnt})
		}
	}
	if k > 1 {
		f = append(f, [2]int{k, 1})
	}
	return f
}
