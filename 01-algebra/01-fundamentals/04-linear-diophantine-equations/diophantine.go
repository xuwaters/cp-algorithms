package diophantine

//
// Linear Diophantine Equation
//   a * x + b * y = c
// where a, b, c are given integers, and x, y are unknown integers.
//
//   if a == 0 , b == 0, we either have zero or infinite solutions.
//
// let g = gcd(a, b)
//   if g | c  then this equation have solutions, otherwise no solution
//
// Proof:
//   let A = { a*x + b*y }
//   let d0 = a*x0 + b*y0 be the smallest positive number in A
//   so for any other positive number
//     d1 = a*x1 + b*y1 = p*d0 + r (where r >=0 and r < d0)
//   that is r = a*x1 + b*y1 - p*(a*x0 + b*y0) = a*(x1-p*x0) + b*(y1-p*y0)
//     so r in A, but r < d0 and d0 is smallest positive number in A
//     so r == 0
//     so for any other positive num di, we have d0 | di
//   since a in A, and b in A, so  d0 | a and d0 | b
//
//   and for any positive comon divisor d of (a, b)
//     let a = k*d, b = m*d
//     so d0 = a*x0 + b*y0 = d*(k*x0 + m*y0)
//     so d0 is the greatest common divisor
//
//   so if c is not a multiple of d0, there is no solution for a*x + b*y = c
//
//
//    if we have  a * x0 + b * y0 = g
//    we have a * (x0*c/g) + b * (y0*c/g) = c
//

//
// FindAnyDiophantineSolution returns
// found: has solution
// g = gcd(a, b)
// a * x + b * y = c
//
func FindAnyDiophantineSolution(a, b, c int) (found bool, g, x, y int) {
	// special cases
	if a == 0 && b == 0 {
		if c == 0 {
			found = true
		} else {
			found = false
		}
		return
	}

	abs := func(v int) int {
		if v < 0 {
			v = -v
		}
		return v
	}

	//
	g, x0, y0 := GreatestCommonDivisorExt(abs(a), abs(b))
	if c%g != 0 {
		found = false
		return
	}
	found = true
	x = x0 * c / g
	y = y0 * c / g
	// fix sign
	if a < 0 {
		x = -x
	}
	if b < 0 {
		y = -y
	}
	return
}

func GreatestCommonDivisorExt(a, b int) (g int, x int, y int) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := GreatestCommonDivisorExt(b, a%b)
	x = y1
	y = x1 - (a/b)*y1
	return
}
