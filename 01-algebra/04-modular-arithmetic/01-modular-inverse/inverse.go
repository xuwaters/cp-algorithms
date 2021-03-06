package modular

//
// a * x = 1 (mod m)
//
// x is inverse of a modular m, written as a^-1 ;
// if and only if gcd(a, m) = 1, we can find x

//
// consider linear Diophantine function
//     a * x + m * y = 1
//  => a * x = -y * m + 1
//  => a * x = 1 (mod m)
//
// we can use extended gcd to solve this equation
//
//

// Another method:
// Finding modular inverse using binary exponentiation
//
//  if gcd(a, m) = 1 then:
//    a^phi(m) = 1 (mod m)
//  so x = a^(phi(m)-1)  , that is a^(-1)
//
//  phi(m) = m * (1-1/p1) * (1-1/p2) * .. * (1-1/pk)
//
//
//

//  Find the modular inverse for every number modulo m
//
//   inv(i) = - floor(m/i) * inv[m % i] (mod m)
//
//  Proof:
//      m%i = m - floor(m/i) * i
//      m%i = -[m/i] * i    (mod m)
//  multiply both side by inv(i) * inv(m%i)
//  =>  (m%i) * inv(i) * inv(m%i) = -[m/i] * i * inv(i) * inv(m%i)  (mod m)
//  =>  inv(i) = -[m/i] * inv(m%i)   (mod m)
//

func CalcInverse(m int) []int {
	inv := make([]int, m)
	inv[1] = 1
	for i := 2; i < m; i++ {
		inv[i] = (m - (m/i)*inv[m%i]%m) % m
	}
	return inv
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
	//
	for b != 0 {
		q := a / b
		a, b = b, a%b
		x0, x1 = x1, x0-x1*q
		y0, y1 = y1, x0-y1*q
	}
	return a, x0, y0
}
