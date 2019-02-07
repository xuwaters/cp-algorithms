package gcd

//
// Proof that: gcd(a, b) == gcd(b, a mod b) for all a >= 0 and b > 0
//
// let d = gcd(a, b)
//  => d | a, and d | b
// let a mod b = a - b * floor(a/b)
//  => d | (a mod b)
//  => d | gcd(b, a mod b)
//  => gcd(a, b) | gcd(b, a mod b)
//
// by the same way, we have
//   gcd(b, a mod b) | gcd(a, b)
//
//  so gcd(a, b) == gcd(b, a mod b)
//

//
// Time Complexity
//   if a > b >= 1 and b < Fn (n-th fibonacci number) for some n,
//   the Euclidean algorithm performs at most n-2 recursive calls.
//
// Since Fn grows exponentially,
//  so gcd time complexity is O(log(min(a, b)))
//

//
// We can also prove that any two consecutive terms of the Fibonacci sequence
// are relatively prime.
//
// First:
//   gcd(F0, F1) = 1
// Suppose gcd(Fn-1, Fn) = 1,
//   then gcd(Fn, Fn+1) = gcd(Fn, Fn + Fn-1) = gcd(Fn-1, Fn) = 1
// DONE
//

func GreatestCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	}
	return GreatestCommonDivisor(b, a%b)
}

func GreatestCommonDivisorLoop(a, b int) int {
	for b != 0 {
		a %= b
		a, b = b, a
	}
	return a
}

func LeastCommonMultiple(a, b int) int {
	return a / GreatestCommonDivisorLoop(a, b) * b
}
