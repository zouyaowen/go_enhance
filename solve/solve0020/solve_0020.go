package solve0020

import (
	"fmt"
	"sync"
)

/**
 * @index 20
 * @title 有效的括号
 * @difficulty 简单
 * @tags stack,string
 * @draft false
 * @link https://leetcode-cn.com/problems/valid-parentheses/
 * @frontendId 20
 */

type StackElement interface{}

type Stack struct {
	elements []StackElement
	lock     sync.RWMutex
}

func NewStack() *Stack {
	stack := &Stack{}
	stack.elements = make([]StackElement, 0)
	return stack
}

func (s *Stack) Push(element interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.elements = append(s.elements, element)
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Pop() StackElement {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.elements) == 0 {
		return nil
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[0 : len(s.elements)-1]
	return element
}

func (s *Stack) Peek() StackElement {
	if len(s.elements) == 0 {
		return nil
	}
	element := s.elements[len(s.elements)-1]
	// s.elements = s.elements[0 : len(s.elements)-1]
	return element
}

func isValid(s string) bool {
	stack := NewStack()
	for _, char := range s {
		// fmt.Println("char", char)
		// fmt.Printf("Unicode: %c  %d\n", char, char)
		// go range 遍历字符串,默认遍历数值，%c 获取单个字符
		newElement := fmt.Sprintf("%c", char)
		if stack.IsEmpty() {
			stack.Push(newElement)
			continue
		}
		if newElement == ")" || newElement == "]" || newElement == "}" {
			peekElement := fmt.Sprintf("%s", stack.Peek())
			if newElement == ")" && peekElement == "(" {
				stack.Pop()
			} else if newElement == "]" && peekElement == "[" {
				stack.Pop()
			} else if newElement == "}" && peekElement == "{" {
				stack.Pop()
			} else {
				return false
			}
		} else {
			stack.Push(newElement)
		}
	}
	return stack.IsEmpty()
}
