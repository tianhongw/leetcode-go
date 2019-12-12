package main

import (
	"errors"
)

type stack struct {
	items []interface{}
	len   int
	top   int
}

func newStack() *stack {
	s := &stack{
		top: -1,
	}

	return s
}

func (s *stack) push(item interface{}) {
	s.items = append(s.items, item)
	s.top++
	s.len++
}

func (s *stack) pop() (interface{}, error) {
	if s.len == 0 {
		return nil, errors.New("empty stack")
	}

	item := s.items[s.top]
	s.items = s.items[0 : s.len-1]

	s.top--
	s.len--

	return item, nil
}

func (s *stack) length() int {
	return s.len
}

func isValid(s string) bool {
	stack := newStack()

	for _, c := range s {
		sc := string(c)
		switch sc {
		case "(":
			stack.push(sc)
		case "{":
			stack.push(sc)
		case "[":
			stack.push(sc)
		case ")":
			fv, err := stack.pop()
			if err != nil {
				return false
			}
			if fv.(string) != "(" {
				return false
			}
		case "}":
			fv, err := stack.pop()
			if err != nil {
				return false
			}
			if fv.(string) != "{" {
				return false
			}
		case "]":
			fv, err := stack.pop()
			if err != nil {
				return false
			}
			if fv.(string) != "[" {
				return false
			}
		}
	}

	return stack.length() == 0
}
