package dp_opt

//
// some dynamic programming problems have a recurrence of this form:
//
//    dp(i, j) = min{ dp(i-1, k) + c(k,j) } for k <= j
//
// if opt(i, j) is the vale of k that minimizes the above expression.
// if opt(i, j) <= opt(i, j+1), this is known as monotonically condition,
// we can apply divide and conquer DP
//
//
// The greatest difficulty with Divide and Conquer DP problems is proving the monotonicity of opt.
// Many Divide and Conquer DP problems can also be solved with the Convex Hull trick or vice-versa.
// It is useful to know and understand both!
//
//
//

// Generic Implementation

type DivideConquerDP struct {
	n      int
	cost   func(i, j int) int
	dpPrev []int
	dpCurr []int
}

func (dp *DivideConquerDP) Compute(left, right int, optLeft, optRight int) {
	// compute dp_cur[l], ... dp_cur[r] (inclusive)
	if left > right {
		return
	}
	mid := (left + right) / 2
	best := [2]int{-1, -1} // { value, index }
	for k := optLeft; k <= mid && k <= optRight; k++ {
		currCost := dp.dpPrev[k-1] + dp.cost(k, mid)
		if best[1] < 0 || currCost < best[0] || currCost == best[0] && k < best[1] {
			best = [2]int{currCost, k}
		}
	}
	dp.dpCurr[mid] = best[0]
	opt := best[1]
	dp.Compute(left, mid-1, optLeft, opt)
	dp.Compute(mid+1, right, opt, optRight)
}

// TODO: real usage

// 
// https://www.spoj.com/problems/LARMY/
// 
// dp[i][j]
//  = dp[i-1][k] + cost(k, j)
// 
// Prove: opt(i,j) <= opt(i,j+1)
// 
// [--------------------------|---------------][-]
// 1. add (j+1) to end of row will only increase unhappiness by itself.
// 2. if move opt(i, j) to the left, unhappiness will increase (because opt(i,j) is already opt for (i,j))
// 
// so opt(i, j) <= opt(i, j+1)
// 
// there is another question: how to calculate cost(k, j) efficiently?
// we can still use method similar with RMQ method.
// precalculate cost with [2^k] length of log(n) layers, then we can answer result with log(len) time.
// 
