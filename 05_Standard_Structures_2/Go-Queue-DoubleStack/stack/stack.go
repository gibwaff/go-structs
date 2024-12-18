package stack

import (
	"strconv"
)

type Stack struct {
	Data []int
	Size int
}

func NewStack(arr []int) Stack {
	st := Stack{
		Data: arr,
		Size: 0,
	}
	return st
}

func (s *Stack) IsEmpty() bool {
	return s.Data[0] == 0
}

func (s *Stack) GetSize() int {
	return s.Size
}

func (s *Stack) Reallocate() {
	if s.Size == 0 {
		arr := make([]int, 1)
		s.Data = arr
		return
	}
	arr := make([]int, len(s.Data)*2)
	for i, el := range s.Data {
		arr[i] = el
	}
	s.Data = arr
}

func (s *Stack) Push_back(x int) {
	if s.Size == len(s.Data) {
		s.Reallocate()
	}
	s.Data[s.Size] = x
	s.Size++
}

func (s *Stack) Lower() {
	arr := make([]int, len(s.Data)/4)
	for i := 0; i < len(arr); i++ {
		arr[i] = s.Data[i]
	}
	s.Data = arr
}

func (s *Stack) Pop_back() int {
	num := s.Data[s.Size-1]
	s.Data[s.Size-1] = 0
	s.Size--
	if s.Size*4 <= len(s.Data) && len(s.Data) > 4 {
		s.Lower()
	}
	return num
}

func (s *Stack) Top() int {
	return s.Data[s.Size-1]
}

func IsSign(c rune) bool {
	if c == '*' || c == '+' || c == '-' {
		return true
	}
	return false
}

func Calculating(a, b int, sign rune) int {
	if sign == '*' {
		return a * b
	} else if sign == '-' {
		return a - b
	} else {
		return a + b
	}
}

func IsSpace(c rune) bool {
	return c == ' '
}

func CalcPolish(s string) (res int) {
	rune_s := []rune(s)
	st_arr := make([]int, 0)
	st := NewStack(st_arr)
	isResCnaged := false

	for i := 0; i < len(rune_s); i++ {
		if IsSign(rune_s[i]) {
			if !isResCnaged {
				res = st.Top()
				st.Pop_back()
				isResCnaged = true
			}
			a := st.Top()
			st.Pop_back()
			res = Calculating(a, res, rune_s[i])
			i++
		} else {
			buf_str := strconv.Itoa(int(rune_s[i] - 48))
			for j := i + 1; ; j++ {
				if IsSpace(rune_s[j]) {
					i = j
					break
				}
				buf_str += strconv.Itoa(int(rune_s[j] - 48))
			}
			a, _ := strconv.Atoi(buf_str)
			st.Push_back(a)
		}
	}

	return
}
