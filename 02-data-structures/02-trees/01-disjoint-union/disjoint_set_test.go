package disjoint_set

import "testing"

func TestDisjointSet(t *testing.T) {

	type Command func(s *DisjointSet) int

	cmdMakeSet := func(val int) Command {
		return func(s *DisjointSet) int {
			s.MakeSet(val)
			return 0
		}
	}

	cmdUnionSets := func(a, b int) Command {
		return func(s *DisjointSet) int {
			s.UnionSets(a, b)
			return 0
		}
	}

	cmdCountSet := func() Command {
		return func(s *DisjointSet) int {
			return s.CountSets()
		}
	}

	commands := []struct {
		ans int
		cmd Command
	}{
		{ans: 0, cmd: cmdMakeSet(1)},
		{ans: 0, cmd: cmdMakeSet(2)},
		{ans: 0, cmd: cmdMakeSet(3)},
		{ans: 0, cmd: cmdMakeSet(4)},
		{ans: 0, cmd: cmdUnionSets(1, 2)},
		{ans: 0, cmd: cmdUnionSets(3, 4)},
		{ans: 0, cmd: cmdMakeSet(5)},
		{ans: 0, cmd: cmdMakeSet(6)},
		{ans: 0, cmd: cmdUnionSets(5, 6)},
		{ans: 0, cmd: cmdUnionSets(1, 6)},
		{ans: 2, cmd: cmdCountSet()},
	}

	s := NewDisjointSet()
	for _, command := range commands {
		got := command.cmd(s)
		if got != command.ans {
			t.Fatalf("ERR: got = %+v, command = %+v", got, command)
		} else {
			t.Logf(" OK: got = %+v, command = %+v", got, command)
		}
	}
}
