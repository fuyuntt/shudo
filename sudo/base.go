package sudo

import (
	"strconv"
	"strings"
)

var bitMap = make([]int8, 257)

func init() {
	var bit int8
	for i := 0; i < 257; i++ {
		if i >= 1<<bit {
			bit++
		}
		bitMap[i] = bit
	}
}

const emptyNum Num = 0x1ff

type Num int16

func toNum(i int8) Num {
	return 1 << (i - 1)
}

func (n Num) Exact() int8 {
	return bitMap[n]
}

func (n Num) PrintStr() string {
	if n != emptyNum {
		return strconv.Itoa(int(n.Exact()))
	} else {
		return "_"
	}
}

type Sudo [81]Num

func FromStr(str string) *Sudo {
	res := Sudo{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			idx := i*9 + j
			n := str[idx] - '0'
			if n == 0 {
				res[idx] = emptyNum
			} else {
				res[idx] = toNum(int8(n))
			}
		}
	}
	return &res
}

func (s *Sudo) Iter(fromR, fromC int, f func(r int, c int, n Num) bool) bool {
	for i := fromR; i < 9; i++ {
		for j := fromC; j < 9; j++ {
			if !f(i, j, s[i*9+j]) {
				return false
			}
		}
	}
	return true
}

func (s *Sudo) PrintStr() string {
	sb := strings.Builder{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				sb.WriteString("  ")
				for l := 0; l < 3; l++ {
					idx := i*27 + j*9 + k*3 + l
					sb.WriteString(" ")
					sb.WriteString(s[idx].PrintStr())
				}
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
func (s *Sudo) ToStr() string {
	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		sb.WriteString(s[i].PrintStr())
	}
	return sb.String()
}

func (s *Sudo) Resolve() bool {
	var st state
	st.initState(*s)
	return s.deduction(&st, 0)
}

func (s *Sudo) deduction(st *state, idx int) bool {
	if idx >= 81 {
		return true
	}
	if s[idx] != emptyNum {
		return s.deduction(st, idx+1)
	}
	r := idx / 9
	c := idx % 9
	ce := cell(r, c)
	for i := int8(0); i < 9; i++ {
		if st.check(r, c, ce, i+1) {
			st.add(r, c, ce, i+1)
			s[idx] = toNum(i + 1)
			if s.deduction(st, idx+1) {
				return true
			}
			st.rm(r, c, ce, i+1)
			s[idx] = emptyNum
		}
	}
	return false
}

type state struct {
	rows  [9]Num
	cols  [9]Num
	cells [9]Num
}

func (s *state) initState(sd Sudo) {
	sd.Iter(0, 0, func(r int, c int, n Num) bool {
		if n != emptyNum {
			s.rows[r] |= n
			s.cols[c] |= n
			s.cells[cell(r, c)] |= n
		}
		return true
	})
}
func (s *state) check(r, c, cell int, x int8) bool {
	n := Num(1 << (x - 1))
	return (s.rows[r]|s.cols[c]|s.cells[cell])&n == 0
}
func (s *state) add(r, c, cell int, x int8) {
	n := Num(1 << (x - 1))
	s.rows[r] |= n
	s.cols[c] |= n
	s.cells[cell] |= n
}
func (s *state) rm(r, c, cell int, x int8) {
	n := Num(1 << (x - 1))
	s.rows[r] &= ^n
	s.cols[c] &= ^n
	s.cells[cell] &= ^n
}
func cell(i, j int) int {
	return i/3*3 + j/3
}
