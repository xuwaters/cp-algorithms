package ac

import (
	"container/list"
)

type ACMachine struct {
	root *State
}

type State struct {
	Next       map[rune]*State
	Leaf       bool
	Parent     *State
	JumpRune   rune            // parent.Next[JumpRune] = this
	PrefixLink *State          // longest prefix state (same as KMP's PrefixFunction) that is suffix of this state
	GoJump     map[rune]*State // final jump decision
	Length     int             // path length from root
}

func NewACMachine() *ACMachine {
	return &ACMachine{
		root: NewState(nil, 0, 0),
	}
}

func NewState(p *State, chr rune, length int) *State {
	return &State{
		Next:       make(map[rune]*State),
		Leaf:       false,
		Parent:     p,
		JumpRune:   chr,
		PrefixLink: nil,
		GoJump:     make(map[rune]*State),
		Length:     length,
	}
}

func (ac *ACMachine) AddString(str string) {
	curr := ac.root
	for _, chr := range []rune(str) {
		if curr.Next[chr] == nil {
			next := NewState(curr, chr, curr.Length+1)
			curr.Next[chr] = next
		}
		curr = curr.Next[chr]
	}
	curr.Leaf = true
}

// BuildLinks builds all links after all strings are added.
func (ac *ACMachine) BuildLinks() {
	queue := list.New()
	queue.PushBack(ac.root)
	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(*State)
		ac.setupPrefixLink(curr)
		for _, next := range curr.Next {
			queue.PushBack(next)
		}
	}
}

func (ac *ACMachine) setupPrefixLink(curr *State) *State {
	if curr.PrefixLink == nil {
		if curr == ac.root || curr.Parent == ac.root {
			curr.PrefixLink = ac.root
		} else {
			parentPrefixLink := ac.setupPrefixLink(curr.Parent)
			curr.PrefixLink = ac.setupGo(parentPrefixLink, curr.JumpRune)
		}
		// setup all other goJump for curr state
		for chr := range curr.Next {
			ac.setupGo(curr, chr)
		}
		for chr := range curr.PrefixLink.GoJump {
			ac.setupGo(curr, chr)
		}
	}
	return curr.PrefixLink
}

func (ac *ACMachine) setupGo(curr *State, chr rune) *State {
	next, ok := curr.GoJump[chr]
	if !ok {
		next = curr.Next[chr]
		if next == nil {
			if curr == ac.root {
				next = ac.root
			} else {
				prev := ac.setupPrefixLink(curr)
				next = ac.setupGo(prev, chr)
			}
		}
		if next != ac.root {
			curr.GoJump[chr] = next
		}
	}
	return next
}

func (ac *ACMachine) Go(curr *State, chr rune) *State {
	if next, ok := curr.GoJump[chr]; ok {
		return next
	}
	return ac.root
}

func (ac *ACMachine) FindAll(s string) [][2]int {
	matches := make([][2]int, 0)
	curr := ac.root
	for i, chr := range []rune(s) {
		curr = ac.Go(curr, chr)
		if curr.Leaf {
			matches = append(matches, [2]int{i + 1 - curr.Length, i + 1})
		}
	}
	return matches
}
