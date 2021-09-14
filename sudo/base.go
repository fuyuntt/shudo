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

const (
	emptyNum Num = 0x1ff
	one      Num = 0x001
	nine     Num = 0x100
)

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
	for i := one; i <= nine; i <<= 1 {
		if st.check(r, c, ce, i) {
			s[idx] = i
			st.add(r, c, ce, i)
			if s.deduction(st, idx+1) {
				return true
			}
			st.rm(r, c, ce, i)
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
	for i := 0; i < 81; i++ {
		r := i / 9
		c := i % 9
		n := sd[i]
		if n != emptyNum {
			s.rows[r] |= n
			s.cols[c] |= n
			s.cells[cell(r, c)] |= n
		}
	}
}
func (s *state) check(r, c, cell int, x Num) bool {
	return (s.rows[r]|s.cols[c]|s.cells[cell])&x == 0
}
func (s *state) add(r, c, cell int, n Num) {
	s.rows[r] |= n
	s.cols[c] |= n
	s.cells[cell] |= n
}
func (s *state) rm(r, c, cell int, n Num) {
	s.rows[r] &= ^n
	s.cols[c] &= ^n
	s.cells[cell] &= ^n
}
func cell(i, j int) int {
	return i/3*3 + j/3
}
