package glispextensions

import (
	"fmt"

	glisp "github.com/zhemao/glisp/interpreter"
)

type GlispVar struct {
	Value interface{}
}

func (gv GlispVar) SexpString() string {
	return fmt.Sprint(gv.Value)
}

type GlispReference struct {
	Value glisp.Sexp
}

func (gr *GlispReference) SexpString() string {
	return fmt.Sprint(gr.Value.SexpString())
}
