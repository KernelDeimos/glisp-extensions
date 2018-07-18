package glispextensions

import "fmt"

type GlispVar struct {
	Value interface{}
}

func (gv GlispVar) SexpString() string {
	return fmt.Sprint(gv.Value)
}
