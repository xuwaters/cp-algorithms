package totient

//
// phi-function
// phi(n) is the number of integers between 1 and n inclusive, which are coprime to n.
//
// if p is prime
//   phi(p) = p -1
//   phi(p^k) = p^k - p^(k-1) = p^k * (1 - 1/p)
//
// if a, b are coprime
//   phi(a*b) = phi(a) * phi(b)
//
// else
//   phi(a*b) = phi(a) * phi(b) * d / phi(d)
// where d = gcd(a, b)
//
// if n = p1^a1 * p2^a2 * ... * pk^ak
// phi(n) = n * (1 - 1/p1) * (1 - 1/p2) * ... * (1 - 1/pk)
//
//
//
//
// a^phi(m) % m = 1  if gcd(a, m) = 1
//
// if m is prime, phi(m) = m * (1 - 1/m) = m - 1
//
// then a^phi(m) % m = 1 for any a < m
//      a^(m-1) % m = 1
//
// a^n = a^(n % phi(m)) (mod m)
//

//
// Generalization:
//  For arbitrary x, m, and n >= log2(m):
//    x^n = x^(phi(m) + (n % phi(m)))   (mod m)
//

func Phi(n int) int {
	result := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			result -= result / i
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}
