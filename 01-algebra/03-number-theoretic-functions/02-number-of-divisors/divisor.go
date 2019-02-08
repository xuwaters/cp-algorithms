package divisor

//
// let n = p1^e1 * p2^e2 * p3^e3 * ... * pk^ek
//
// Number of divisors d(n):
//   d(n) = (e1+1)*(e2+1)*(e3+1)*...*(ek+1)
//
//
// Sum of divisors s(n):
//   s(n) = (p1^(e1+1)-1)/(p1-1) *
//          (p2^(e2+1)-1)/(p2-1) *
//           ...
//          (pk^(ek+1)-1)/(pk-1)
//
//
// both d(n) and s(n) are multiplicative functions
//    f(a * b) = f(a) * f(b)
//
// 

