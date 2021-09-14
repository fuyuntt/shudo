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

type Num struct {
	set   int16
	count int8
	exact int8
}

func (n *Num) IsExact() bool {
	return n.count == 1
}

func (n *Num) SetExact(exact int8) {
	n.set = 1 << (exact - 1)
	n.count = 1
	n.exact = exact
}

func (n *Num) SetAll() {
	n.set = 0x1ff
	n.restore()
}

func (n *Num) Exact() int8 {
	return n.exact
}

func (n *Num) Count() int8 {
	return n.count
}

func (n *Num) Exclude(o *Num) {
	n.set &= ^o.set
	n.restore()
}
func (n *Num) restore() {
	n.count = countBit(n.set)
	if n.count == 1 {
		n.exact = bitMap[n.set]
	}
}
func (n *Num) PrintStr() string {
	if n.IsExact() {
		return strconv.Itoa(int(n.exact))
	} else {
		return "_"
	}
}
func (n *Num) MaxNum() int8 {
	return bitMap[n.set]
}

func countBit(n int16) int8 {
	var i int8
	for n > 0 {
		n &= n - 1
		i++
	}
	return i
}

type Sudo [81]Num

func FromStr(str string) *Sudo {
	res := Sudo{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			idx := i*9 + j
			n := str[idx] - '0'
			if n == 0 {
				res[idx].SetAll()
			} else {
				res[idx].SetExact(int8(n))
			}
		}
	}
	var rows [9]int16
	var cols [9]int16
	var cells [9]int16
	res.Iter(func(r int, c int, cell int, n *Num) bool {
		if n.IsExact() {
			rows[r] |= n.set
			cols[c] |= n.set
			cells[cell] |= n.set
		}
		return true
	})
	res.Iter(func(r int, c int, cell int, n *Num) bool {
		if !n.IsExact() {
			n.set &= ^rows[r]
			n.set &= ^cols[c]
			n.set &= ^cells[cell]
			n.restore()
		}
		return true
	})
	return &res
}

func (s *Sudo) Iter(f func(r int, c int, cell int, n *Num) bool) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !f(i, j, i/3*3+j/3, &s[i*9+j]) {
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

func (s *Sudo) Resolve() {

}

func (s *Sudo) fuzzyDeduction() bool {
	return !s.Iter(func(r int, c int, cell int, n *Num) bool {
		if n.IsExact() {
			return true
		}
		origin := Num{set: n.set, count: n.count, exact: n.exact}
		for i := 0; i < 9; i++ {
			if (origin.set>>i)&1 == 1 {
				n.SetExact(int8(i + 1))
				res := s.strictDeduction(r, c, cell, n)
				if res {
					return false
				}
			}
		}
		n.set = origin.set
		n.count = origin.count
		n.exact = origin.exact
		return true
	})
}
func (s *Sudo) strictDeduction(r int, c int, cell int, n *Num) bool {
	s.Iter(func(r_ int, c_ int, cell_ int, n_ *Num) bool {
		if n_.IsExact() {
			return true
		}
		if r_ == r {

		}
	})
	return true, 0, 0
}
