package chinese_remainder

//
// if we use 100 primes which are larger than 10^9,
// than we can represent numbers as large as (10^9)^100 = 10^900
//
// SZ = 100
// primes = make([]int, SZ)
// r[][] = make([][]int, SZ)
//

var (
	sz     = 100
	primes = make([]int, sz)
	r      = make([][]int, sz)
)

func init() {
	// generate primes
	primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	sz = len(primes)
	// calculate r[i][j] * pi = 1 (mod pj)
	r = CalculateInverse(primes)
}

// BigInt uses Garner's algorithm to represent a big integer
type BigInt struct {
	a []int // a[i] = n % pi
}

func NewBigInt(n int) BigInt {
	it := BigInt{
		a: make([]int, sz),
	}
	if n != 0 {
		for i := 0; i < len(it.a); i++ {
			it.a[i] = n % primes[i]
		}
	}
	return it
}

func (it BigInt) Len() int {
	return len(it.a)
}

func (it BigInt) Add(o BigInt) BigInt {
	ret := NewBigInt(0)
	for i := 0; i < sz; i++ {
		ret.a[i] = it.a[i] + o.a[i]
	}
	return ret
}

func (it BigInt) Subtract(o BigInt) BigInt {
	ret := NewBigInt(0)
	for i := 0; i < sz; i++ {
		ret.a[i] = (it.a[i] - o.a[i] + primes[i]) % primes[i]
	}
	return ret
}

func (it BigInt) Multiply(o BigInt) BigInt {
	ret := NewBigInt(0)
	for i := 0; i < sz; i++ {
		ret.a[i] = (it.a[i] * o.a[i]) % primes[i]
	}
	return ret
}

func (it BigInt) ToValue() int {
	// TODO: use big.Int to calculate result
	// calculate x
	x := GarnerAlgorithm(it.a, primes)
	// reconstruct value
	result := 0
	multi := 1
	for i := 0; i < len(x); i++ {
		result += multi * x[i]
		multi *= primes[i]
	}
	return result
}
