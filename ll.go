package main

import "sync/atomic"

type node struct {
	val  int
	next *node
}

type Stack struct {
	head atomic.Pointer[node]
}

func (s *Stack) Push(val int) {
	newNode := &node{val: val}
	for {
		oldHead := s.head.Load()
		newNode.next = oldHead
		if s.head.CompareAndSwap(oldHead, newNode) {
			return
		}
	}
}

func (s *Stack) Pop() (int, bool) {
	for {
		oldHead := s.head.Load()
		if oldHead == nil {
			return 0, false
		}
		if s.head.CompareAndSwap(oldHead, oldHead.next) {
			return oldHead.val, true
		}
	}
}
