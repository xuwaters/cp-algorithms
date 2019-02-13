package kruskal

import (
	"math/rand"
)

type DisjointSet struct {
	parent map[int]int
}

func NewDisjointSet() *DisjointSet {
	return &DisjointSet{
		parent: make(map[int]int),
	}
}

func (s *DisjointSet) MakeSet(v int) {
	if _, ok := s.parent[v]; !ok {
		s.parent[v] = v
	}
}

func (s *DisjointSet) FindSet(v int) int {
	p, ok := s.parent[v]
	if !ok {
		return -1
	}
	if p == v {
		return p
	}
	s.parent[v] = s.FindSet(s.parent[v])
	return s.parent[v]
}

func (s *DisjointSet) UnionSets(a, b int) {
	pa := s.FindSet(a)
	pb := s.FindSet(b)
	if pa != pb {
		// 1. random link; 2. link by rank; 3. link by nodes
		if rand.Intn(2) == 0 {
			s.parent[pa] = pb
		} else {
			s.parent[pb] = pa
		}
	}
}

func (s *DisjointSet) CountSets() int {
	cnt := 0
	for k, v := range s.parent {
		if k == v {
			cnt++
		}
	}
	return cnt
}
