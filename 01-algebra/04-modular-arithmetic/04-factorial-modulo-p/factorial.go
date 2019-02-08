package factorial

//
// Wilson Theorem:
//   (p-1)! = -1 = p-1  (mod p)
// if p is prime.
//
//    n! =     1 *    2 *...* (p-1)* p*
//          (p+1)* (p+2)*...*(2p-1)*2p*
//         (2p+1)*(2p+2)*...*(3p-1)*3p*
//
//    n! % p = (p-1)*1 * (p-1)*2 * (p-1)*3 * ... * (2*3*...*(n%p))
//
//
//

// TODO: don't understand this algorithm!!
func factmod(n, p int) int {
	res := 1
	for n > 1 {
		f := 1
		if (n/p)%2 == 1 {
			f = p - 1
		}
		res = (res * f) % p
		for i := 2; i <= n%p; i++ {
			res = (res * i) % p
		}
		n /= p
	}
	return res % p
}
