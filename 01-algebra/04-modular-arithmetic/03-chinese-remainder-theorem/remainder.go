package chinese_remainder

//
// let p = p1*p2*p3*...*pk where pi are pairwise relative prime,
// we are also given a set of congruence equations:
//
//   a = a1 (mod p1)
//   a = a2 (mod p2)
//   ...
//   a = ak (mod pk)
//
// where a are some given constants. The original Chinese Reminder Theorem (CRT) states that
// the given set of congruence equations always has one and exactly one solution modulo p.
//
// let Mi = p/pi, and ti = inv(Mi, pi) or ti*Mi = 1 (mod pi)
// solution a = a1*t1*M1 + a2*t2*M2 + ... + ak*tk*Mk  (mod p)
//
//
//
// Corollary:
// A consequence of the CRT is that the equation:
//     x = a (mod p)
// is equivalent to the system of equations
//     x = a (mod p1)
//     x = a (mod p2)
//     ...
//     x = a (mod pk)
//
//
// Garner's Algorithm:
// Another consequence of CRT is that we can represent big numbers using an array of small integers.
// let p = p1*p2*...*pk, any number a less than p can be represented as array a1, a2, ..., ak,
// where ai = a (mod pi)
//   a = x1 + x2*p1 + x3*p1*p2 + ... + xk*p1*p2*...pk_1
// which is called mixed radix representation of a.
//
//   let r[i,j] = inv(p[i])  (mod p[j])
//
//       a1 = a % p1 = x1  (mod p1)
//
//       a2 = a % p2 = x1 + x2*p1   (mod p2)
//    => a2 - x1 = x2*p1 (mod p2)
//    => (a2 - x1) * inv(p1) = x2 (mod p2)
//
//       a3 = a % p3 = x1 + x2*p1 + x3*p1*p2   (mod p3)
//       (a3 - x1) * r[1,3] = x2 + x3*p2     (mod p3)
//       ((a3-x1)*r[1,3]-x2)*r[2,3] = x3    (mod p3)
//
//

func GarnerAlgorithm(a []int, p []int) []int {
	// a[i] = value % p[i]
	n := len(a)
	r := CalculateInverse(p)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = a[i]
		for j := 0; j < i; j++ {
			x[i] = (x[i] - x[j]) * r[j][i]
			x[i] = (x[i] + p[i]) % p[i]
		}
	}
	return x
}

func CalculateInverse(p []int) [][]int {
	n := len(p)
	r := make([][]int, n)
	// r[i][j] = Pi^-1  (mod Pj)
	for i := 0; i < n; i++ {
		r[i] = make([]int, n)
		for j := 0; j < n; j++ {
			r[i][j] = ModInverse(p[i], p[j])
		}
	}
	return r
}

func ModInverse(a, m int) int {
	// calculate inv(a) = a^-1 (mod m)
	// a * x = 1 (mod m)
	// a * x + m * y = 1
	_, x, _ := ExtGCD(a, m)
	return (x%m + m) % m
}

func ExtGCD(a, b int) (g, x, y int) {
	// x0, x1 = [1, 0]
	// y0, y1 = [0, 1]
	// M(q) = [[0, 1],
	//         [1, -q]]

	// Identity Matrix
	x0, x1 := 1, 0
	y0, y1 := 0, 1
	for b != 0 {
		q := a / b
		a, b = b, a%b
		x0, x1 = x1, x0-x1*q
		y0, y1 = y1, x0-y1*q
	}
	return a, x0, y0
}
