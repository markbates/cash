package cash

import (
	"fmt"
	"strings"
)

type Numerator int64

func (n Numerator) Format(sep string) string {
	ss := []string{}
	a := n
	for a != 0 {
		d := a % 1000
		if d != 0 {
			ss = append(ss, fmt.Sprint(d))
		} else {
			ss = append(ss, "000")
		}
		a = a / 1000
	}
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
	return strings.TrimPrefix(strings.Join(ss, sep), "000,")
}

type Denominator int64

func (d Denominator) String() string {
	if d < 0 {
		d = d * -1
	}
	if d < 10 {
		return fmt.Sprintf("0%d", d)
	}
	return fmt.Sprintf("%d", d)
}
