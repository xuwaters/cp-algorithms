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

//
// [b, a%b] = [a, b] * [[0, 1]
//                      [1, -q]]  where q = floor(a/b)
// [gcd(a, b), 0] = [a, b] * [[0,1], [1,-q1]] * [[0,1], [1,-q2]] * ... * [[0,1], [1,-qk]]
// so if we keep track of M = M(q1)*M(q2)*..*M(qk) where M(qi) = [[0, 1], [1, -qi]]
// finally, we get  [g, 0] = [a, b] * M
//
// [[x, s], [y, v], [gcd(a,b), 0]] = [[1, 0], [0, 1], [a, b]] * M(q1) * ... * M(qk)
// x, y  is the result:  a * x + b * y = gcd(a, b)
//
//
func GreatestCommonDivisorExtLoop(a, b int) (g int, x int, y int) {
	// M(q) = [[0, 1], [1, -q]]
	//
	// [x0, x1] * M(q) = [x1, x0-x1*q]
	// [y0, y1] * M(q) = [y1, y0-y1*q]
	// [a, b] * M(q) = [b, a-b*q] = [b, a%b]
	// finally [x0, y0] is the result

	x0, x1 := 1, 0
	y0, y1 := 0, 1
	for b != 0 {
		q := a / b
		a, b = b, a%b
		x0, x1 = x1, x0-x1*q
		y0, y1 = y1, y0-y1*q
	}
	return a, x0, y0
}
