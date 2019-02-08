package congruence

//
// a * x = b (mod n)
//
// Solution 1:
//   if gcd(a, n) = 1
//   we can find a^-1 = a^(phi(n)-1)  (mod n)
//   then x = b * a^-1  (mod n)
//
// If a, n are not coprime (gcd(a, n) > 1)
//    a*x + n*y = b
// we can solve this equation by using extended gcd algorithm.
//    a*xg + n*yg = g
// if g | b
//   a * xg*b/g + n * yg*b/g = b
//
// => x = (xg + i * n/g) * b/g
//    y = (yg - i * a/g) * b/g
//
//
//
//
//
