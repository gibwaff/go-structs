package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	Data []int
	Size int
}

func newStack(arr []int) Stack {
	st := Stack{
		Data: arr,
		Size: len(arr),
	}
	return st
}

func (s *Stack) getSize() int {
	return s.Size
}

func (s *Stack) reallocate() {
	if s.Size == 0 {
		arr := make([]int, 1)
		s.Data = arr
		return
	}
	arr := make([]int, s.Size*2)
	for i, el := range s.Data {
		arr[i] = el
	}
	s.Data = arr
}

func (s *Stack) push_back(x int) {
	if s.Size == len(s.Data) {
		s.reallocate()
	}
	s.Data[s.Size] = x
	s.Size++
}

func (s *Stack) lower() {
	arr := make([]int, len(s.Data)/4)
	for i := 0; i < len(arr); i++ {
		arr[i] = s.Data[i]
	}
	s.Data = arr
}

func (s *Stack) pop_back() {
	s.Data[s.Size-1] = 0
	s.Size--
	if s.Size*4 <= len(s.Data) {
		s.lower()
	}
}

func (s *Stack) top() int {
	return s.Data[s.Size-1]
}

func isSign(c rune) bool {
	if c == '*' || c == '+' || c == '-' {
		return true
	}
	return false
}

func calculating(a, b int, sign rune) int {
	if sign == '*' {
		return a * b
	} else if sign == '-' {
		return a - b
	} else {
		return a + b
	}
}

func isSpace(c rune) bool {
	if c == ' ' {
		return true
	}
	return false
}

func calcPolish(s string) (res int) {
	rune_s := []rune(s)
	st_arr := make([]int, 0)
	st := newStack(st_arr)
	isResCnaged := false

	for i := 0; i < len(rune_s); i++ {
		if isSign(rune_s[i]) {
			if !isResCnaged {
				res = st.top()
				st.pop_back()
				isResCnaged = true
			}
			a := st.top()
			st.pop_back()
			res = calculating(a, res, rune_s[i])
			i++
		} else {
			buf_str := strconv.Itoa(int(rune_s[i] - 48))
			for j := i + 1; ; j++ {
				if isSpace(rune_s[j]) {
					i = j
					break
				}
				buf_str += strconv.Itoa(int(rune_s[j] - 48))
			}
			a, _ := strconv.Atoi(buf_str)
			st.push_back(a)
		}
	}

	return
}

func main() {
	fmt.Println(calcPolish("100 2 30 * -"))
}
