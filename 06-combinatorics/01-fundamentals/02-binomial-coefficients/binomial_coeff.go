package binomial

//
//               n!
// C(n, k) = ----------
//            k!(n-k)!
//
// Properties:
//
//   C(n, k) = C(n-1, k-1) + C(n-1, k)
//   C(n, k) = C(n-1, k-1) * n / k
//   SUM( C(n, k), k=0..n ) = 2^n
//   SUM( C(m, k), m=0..n ) = C(n+1, k+1)
//   SUM( C(n, k)^2, k=0..n ) = C(2n, n)
//   SUM( C(n, k) * k, k=0..n ) = n*2^(n-1)
//   SUM( C(n-k, k), k=0..n ) = Fibonacci[n+1]
//
//
//

func BinomialCoefficients(n, k int) int {
	res := 1
	for i := 1; i <= k; i++ {
		res = res * (n - k + i) / i
	}
	return res
}
