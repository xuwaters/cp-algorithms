package fft

import (
	"math"
	"math/cmplx"
)

func FastFourierTransform(a []complex128, invert bool) {
	n := len(a)
	// n is power of 2
	if n == 1 {
		return
	}
	a0 := make([]complex128, n/2)
	a1 := make([]complex128, n/2)
	for i := 0; 2*i < n; i++ {
		a0[i] = a[2*i]
		a1[i] = a[2*i+1]
	}

	FastFourierTransform(a0, invert)
	FastFourierTransform(a1, invert)

	ang := float64(2) * math.Pi / float64(n)
	if invert {
		ang = -ang
	}

	w := complex(1.0, 0.0)
	// wn := complex(math.Cos(ang), math.Sin(ang)) // e^(i*ang)
	wn := cmplx.Exp(complex(0.0, ang))
	for i := 0; 2*i < n; i++ {
		a[i] = a0[i] + w*a1[i]
		a[i+n/2] = a0[i] - w*a1[i]
		if invert {
			a[i] /= 2
			a[i+n/2] /= 2
		}
		w *= wn
	}
}

func MultiplyArray(a, b []int) []int {
	n := 1
	for n < len(a)+len(b) {
		n <<= 1
	}
	ca := make([]complex128, n)
	cb := make([]complex128, n)
	for i := 0; i < n && (i < len(a) || i < len(b)); i++ {
		if i < len(a) {
			ca[i] = complex(float64(a[i]), 0.0)
		}
		if i < len(b) {
			cb[i] = complex(float64(b[i]), 0.0)
		}
	}

	// begin
	FastFourierTransform(ca, false)
	FastFourierTransform(cb, false)

	// product
	for i := 0; i < n; i++ {
		ca[i] *= cb[i]
	}

	// result
	FastFourierTransform(ca, true)

	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = int(math.Round(real(ca[i])))
	}
	return result
}

// // Iterative version of fft
// 
// using cd = complex<double>;
// const double PI = acos(-1);
// 
// int reverse(int num, int lg_n) {
//     int res = 0;
//     for (int i = 0; i < lg_n; i++) {
//         if (num & (1 << i))
//             res |= 1 << (lg_n - 1 - i);
//     }
//     return res;
// }
// 
// 
// 
// void fft(vector<cd> & a, bool invert) {
//     int n = a.size();
//     int lg_n = 0;
//     while ((1 << lg_n) < n)
//         lg_n++;

//     for (int i = 0; i < n; i++) {
//         if (i < reverse(i, lg_n))
//             swap(a[i], a[reverse(i, lg_n)]);
//     }
// 
//     for (int len = 2; len <= n; len <<= 1) {
//         double ang = 2 * PI / len * (invert ? -1 : 1);
//         cd wlen(cos(ang), sin(ang));
//         for (int i = 0; i < n; i += len) {
//             cd w(1);
//             for (int j = 0; j < len / 2; j++) {
//                 cd u = a[i+j], v = a[i+j+len/2] * w;
//                 a[i+j] = u + v;
//                 a[i+j+len/2] = u - v;
//                 w *= wlen;
//             }
//         }
//     }
// 
//     if (invert) {
//         for (cd & x : a)
//             x /= n;
//     }
// }
// 


// 
//  Another implementation
// 
// using cd = complex<double>;
// const double PI = acos(-1);
// 
// void fft(vector<cd> & a, bool invert) {
//     int n = a.size();
// 
//     for (int i = 1, j = 0; i < n; i++) {
//         int bit = n >> 1;
//         for (; j & bit; bit >>= 1)
//             j ^= bit;
//         j ^= bit;
// 
//         if (i < j)
//             swap(a[i], a[j]);
//     }
// 
//     for (int len = 2; len <= n; len <<= 1) {
//         double ang = 2 * PI / len * (invert ? -1 : 1);
//         cd wlen(cos(ang), sin(ang));
//         for (int i = 0; i < n; i += len) {
//             cd w(1);
//             for (int j = 0; j < len / 2; j++) {
//                 cd u = a[i+j], v = a[i+j+len/2] * w;
//                 a[i+j] = u + v;
//                 a[i+j+len/2] = u - v;
//                 w *= wlen;
//             }
//         }
//     }
// 
//     if (invert) {
//         for (cd & x : a)
//             x /= n;
//     }
// }
// 
