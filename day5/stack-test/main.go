package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

// Push adds an item to the top
func (s *Stack) Push(x int) {
	s.items = append(s.items, x)
}

func (s *Stack) getHeight() int {
	return len(s.items)
}

func (s *Stack) getItems() []int {
	return s.items
}

// Pop removes and returns the top item (LIFO)
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false // stack is empty
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// PopAt removes and returns an item at a specific position
func (s *Stack) PopAt(pos int) (int, bool) {
	if pos < 0 || pos >= len(s.items) {
		return 0, false // out of range
	}
	val := s.items[pos]
	s.items = append(s.items[:pos], s.items[pos+1:]...)
	return val, true
}

func (s *Stack) PushItemAt(item int, pos int) bool {
	if pos < 0 || pos >= s.getHeight() {
		return false
	}

	s.items = append(s.items[:pos], append([]int{item}, s.items[pos:]...)...)
	return true
}

func (s *Stack) PopOddItem() bool {
	if s.getHeight() == 0 {
		return false
	}

	var filtered []int
	for _, v := range s.items {
		if v%2 != 0 {
			filtered = append(filtered, v)
		}
	}
	s.items = filtered

	return true
}

func (s *Stack) PopEvenItem() bool {
	if s.getHeight() == 0 {
		return false
	}

	var filtered []int
	for _, v := range s.items {
		if v%2 == 0 {
			filtered = append(filtered, v)
		}
	}
	s.items = filtered

	return true
}

func (s *Stack) PopItemsAtEvenPos() bool {
	if s.getHeight() == 0 {
		return false
	}

	var filtered []int
	for i := 0; i < s.getHeight(); i++ {
		if i%2 == 0 {
			filtered = append(filtered, s.items[i])
		}
	}

	s.items = filtered

	return true
}

func (s *Stack) PopItemsAtOddPos() {
	if s.getHeight() == 0 {
		fmt.Println("Stack is empty")
	}

	var filtered []int
	for i := 0; i < s.getHeight(); i++ {
		if i%2 != 0 {
			filtered = append(filtered, s.items[i])
		}
	}

	s.items = filtered
}

func (s *Stack) PopARange(start, end int) {
	if s.getHeight() == 0 {
		fmt.Println("Stack is empty")
	} else if (start < 0 || start >= s.getHeight()) || (end < 0 || end >= s.getHeight()) {
		fmt.Println("Out of range")
	}
	var filtered []int = s.items[:start:end]
	s.items = filtered
}

func main() {
	st := &Stack{}
	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Push(9)
	st.Push(40)
	st.Push(27)

	fmt.Println("Stack:", st.getItems())

	val, ok := st.Pop()
	fmt.Println("Pop top :", val, ok, "=>", st.getItems(), "Height:", st.getHeight())

	val, ok = st.PopAt(1)
	fmt.Println("Pop at idx 1:", val, ok, "=>", st.getItems(), "Height:", st.getHeight())

	ok = st.PushItemAt(50, 1)
	fmt.Println("Push at idx 1:", ok, "=>", st.getItems(), "Height:", st.getHeight())

	ok = st.PopOddItem()
	fmt.Println("Pop odd item:", ok, "=>", st.getItems(), "Height:", st.getHeight())

	ok = st.PopEvenItem()
	fmt.Println("Pop even item:", ok, "=>", st.getItems(), "Height:", st.getHeight())

	st.Push(10)
	st.Push(20)
	st.Push(30)
	st.Push(9)
	st.Push(40)
	st.Push(27)

	ok = st.PopItemsAtEvenPos()
	fmt.Println("Pop items at even pos:", ok, "=>", st.getItems(), "Height:", st.getHeight())

	st.PopItemsAtOddPos()
	fmt.Println("Pop items at odd pos:", "=>", st.getItems(), "Height:", st.getHeight())

	st.PopARange(2, 4)
	fmt.Println("Pop a range:", "=>", st.getItems(), "Height:", st.getHeight())
}
