package gcd_ext

//
// GreatestCommonDivisorExt returns (g, x, y)
// so that GCD(a, b) = g = a * x + b * y
//
// if we have gcd(b, a mod b) = (g, x1, y1)
// since a mod b = a - floor(a/b) * b
//   g = b * x1 + (a - floor(a/b)*b) * y1
//     = a * y1 + b * (x1 - floor(a/b) * y1)
//   so
//     x = y1, y = x1 - floor(a/b) * y1
//   such that
//     g = a * x + b * y
//
func GreatestCommonDivisorExt(a, b int) (g int, x int, y int) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := GreatestCommonDivisorExt(b, a%b)
	x = y1
	y = x1 - (a/b)*y1
	return
}
