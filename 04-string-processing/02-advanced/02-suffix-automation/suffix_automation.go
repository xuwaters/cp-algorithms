package suffix_automation

//
// endpos(t) = set of all positions in the string s, in which the occurrences of t end.
// e.g. endpos("bc") = [2, 4] for string "abcbc"
//
// in suffix machine, endpos-equivalent substrings correspond to the same state.
//
//
// Lemma 1:
//   Two non-empty substrings u and w (with len(u) < len(w)) are endpos-equivalent,
// if and only if the string u occurs in s in the form of a suffix of w.
//
// Lemma 2:
//   Consider two non-empty substrings u and w (with len(u) < len(w)). Then their sets
// endpos either don't intersect at all, or endpos(w) is a subset of enpos(u). And it
// depends on if u is a suffix of w or not.
//
//
// Lemma 3:
//   Consider an endpos-equivalent class. Sort all the substrings in this class by non-increasing
// length. Then in the resulting sequence each substring will be one shorter than the previous one,
// and at the same time will be a suffix of the previous one.
//
//
//
//
// Recap:
//  - substrings of s can be decomposed into equivalent classes according to their endpos.
//  - suffix automation consists of the initial state t0, as well as of one state for each
//    endpos-equivalent class.
//  - for each state v, len(v) = length of longest substring in v, minlen(v) = length of shortest
//    substring in v.
//  - for each state v != t0, suffix link leads to a state that corresponds to the suffix string
//    longest(v) of length minlen(v)-1.
//  - for state v != t0, minlen(v) = len(link(v)) + 1
//  - If we start from an arbitrary state v0 and follow the suffix links, then sooner or later we will
//    reach the initial state t0. In this case we obtain a sequence of disjoint intervals
//    [minlen(vi);len(vi)], which in union forms the continuous interval [0;len(v0)].
//
//
// TODO: not quit understand this process
//

type State struct {
	Len  int
	Link *State
	Next map[rune]*State
}

func NewState(length int) *State {
	return &State{
		Len:  length,
		Next: make(map[rune]*State),
	}
}

type SuffixAutomation struct {
	Start *State
	Last  *State
}

func NewSuffixAutomation() *SuffixAutomation {
	sa := &SuffixAutomation{}
	sa.init()
	return sa
}

func (sa *SuffixAutomation) init() {
	sa.Start = NewState(0)
	sa.Last = sa.Start
}

func (sa *SuffixAutomation) Extend(chr rune) *State {
	copymap := func(dst, src map[rune]*State) {
		for k, v := range src {
			dst[k] = v
		}
	}

	// u = new string end
	u := NewState(sa.Last.Len + 1)
	v := sa.Last
	for ; v != nil && v.Next[chr] == nil; v = v.Link {
		v.Next[chr] = u
	}
	if v == nil {
		u.Link = sa.Start
	} else if v.Next[chr].Len == v.Len+1 {
		u.Link = v.Next[chr]
	} else {
		o := v.Next[chr]
		// copy o to n
		n := NewState(v.Len + 1)
		copymap(n.Next, o.Next) // copy to n.Next from o.Next
		n.Link = o.Link
		//
		o.Link = n
		u.Link = n
		for ; v != nil && v.Next[chr] == o; v = v.Link {
			v.Next[chr] = n
		}
	}
	sa.Last = u
	return u
}
