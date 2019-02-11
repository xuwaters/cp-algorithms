package suffix_array

import (
	"math"
	"sort"
)

type SuffixArray struct {
	data      []rune
	rank      [][]int
	layer     int          // len(rank)
	alphabets map[rune]int // rank of alphabets
	arr       []int        // suffix array result
}

func NewSuffixArray(data []rune) *SuffixArray {
	return &SuffixArray{
		data: data,
	}
}

func (s *SuffixArray) buildAlphabets() {
	s.alphabets = make(map[rune]int)
	chrList := make([]rune, 0)
	for _, chr := range s.data {
		if _, ok := s.alphabets[chr]; !ok {
			s.alphabets[chr] = 1
			chrList = append(chrList, chr)
		}
	}
	sort.Slice(chrList, func(i, j int) bool { return chrList[i] < chrList[j] })
	for i, chr := range chrList {
		s.alphabets[chr] = i
	}
}

func (s *SuffixArray) Len() int {
	return len(s.data)
}

func (s *SuffixArray) Layer() int {
	return len(s.rank)
}

func (s *SuffixArray) Arr() []int {
	return s.arr
}

func (s *SuffixArray) Build() {
	s.buildAlphabets()

	n := len(s.data)
	layerSize := int(math.Ceil(math.Log2(float64(n)))) + 1
	s.rank = make([][]int, layerSize) // length of the prefix of suffix in layer L is (2^L)
	for i := 0; i < layerSize; i++ {
		s.rank[i] = make([]int, n)
	}

	// layer 0
	for i := 0; i < n; i++ {
		s.rank[0][i] = s.alphabets[s.data[i]]
	}

	//
	type entry struct {
		r0  int
		r1  int
		idx int
	}
	entries := make([]entry, n)

	entryCompare := func(i, j int) int {
		if entries[i].r0 == entries[j].r0 {
			return entries[i].r1 - entries[j].r1
		}
		return entries[i].r0 - entries[j].r0
	}

	layer := 1
	for cnt := 1; cnt < n; cnt <<= 1 {
		//
		for i := 0; i < n; i++ {
			entries[i].idx = i
			entries[i].r0 = s.rank[layer-1][i]
			if i+cnt < n {
				entries[i].r1 = s.rank[layer-1][i+cnt]
			} else {
				entries[i].r1 = -1
			}
		}
		// using radix sort to speed up ( O(n*logn) )
		// quicksort ( O(n*logn*logn))
		sort.Slice(entries, func(i, j int) bool {
			return entryCompare(i, j) < 0
		})
		// update rank of layer
		for i := 0; i < n; i++ {
			idx := entries[i].idx
			if i == 0 {
				s.rank[layer][idx] = i
			} else {
				prevIdx := entries[i-1].idx
				if entryCompare(i, i-1) == 0 {
					s.rank[layer][idx] = s.rank[layer][prevIdx]
				} else {
					s.rank[layer][idx] = i
				}
			}
		}
		//
		layer++
	}

	// result
	s.arr = make([]int, n)
	for i := 0; i < n; i++ {
		s.arr[i] = entries[i].idx
	}
}

func (s *SuffixArray) LongestCommonPrefix(i, j int) int {
	// LCP of Suffix[i], Suffix[j]
	lcp := 0
	n := s.Len()
	if i == j {
		return n - i
	}
	for layer := s.Layer() - 1; layer >= 0 && i < n && j < n; layer-- {
		cnt := (1 << uint(layer))
		if s.rank[layer][i] == s.rank[layer][j] {
			i += cnt
			j += cnt
			lcp += cnt
		}
	}
	return lcp
}


// 
// p is suffix array, p[i] is index of i-th suffix
// 
// vector<int> lcp_construction(string const& s, vector<int> const& p) {
//     int n = s.size();
//     vector<int> rank(n, 0);
//     for (int i = 0; i < n; i++)
//         rank[p[i]] = i;
// 
//     int k = 0;
//     vector<int> lcp(n-1, 0);
//     for (int i = 0; i < n; i++) {
//         if (rank[i] == n - 1) {
//             k = 0;
//             continue;
//         }
//         int j = p[rank[i] + 1];
//         while (i + k < n && j + k < n && s[i+k] == s[j+k])
//             k++;
//         lcp[rank[i]] = k;
//         if (k)
//             k--;
//     }
//     return lcp;
// }
// 
