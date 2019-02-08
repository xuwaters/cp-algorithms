package ternary

// -1,0,1  base 3 number system
// use z denote -1
//
//  0 = 0
//  1 = 1
//  2 = 1Z
//  3 = 10
//  4 = 11
//  5 = 1Z1
//  6 = 1Z0
//
//
//

func Int2BalancedTernary(n int) []int {
	if n < 0 {
		ret := Int2BalancedTernary(-n)
		for i := 0; i < len(ret); i++ {
			ret[i] = -ret[i]
		}
		return ret
	}
	if n == 0 {
		return []int{0}
	}

	// fmt.Printf(">> n = %d\n", n)
	digits := make([]int, 0)
	c := 0
	for n > 0 || c > 0 {
		d := n%3 + c
		c = d / 3
		d %= 3
		if d == 2 {
			c++
			d = -1
		}
		// fmt.Printf("n = %d, c = %d, d = %d\n", n, c, d)
		digits = append(digits, d)
		n /= 3
	}
	return digits
}

func Digits2String(digits []int) string {
	n := len(digits)
	s := make([]byte, len(digits))
	for i := 0; i < n; i++ {
		switch digits[n-1-i] {
		case -1:
			s[i] = 'Z'
		case 0:
			s[i] = '0'
		case 1:
			s[i] = '1'
		default:
			panic("invalid input")
		}
	}
	return string(s)
}

func String2Value(str string) int {
	getVal := func(chr byte) int {
		switch chr {
		case '0':
			return 0
		case '1':
			return 1
		case 'Z':
			return -1
		default:
			panic("invalid chr")
		}
	}

	val := 0
	n := len(str)
	for i := 0; i < n; i++ {
		val = val*3 + getVal(str[i])
	}
	return val
}
