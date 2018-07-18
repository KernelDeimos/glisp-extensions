package glispextensions

import glisp "github.com/zhemao/glisp/interpreter"

func GlispForkAndRun(
	env *glisp.Glisp, _ string, args []glisp.Sexp,
) (glisp.Sexp, error) {
	code := args[0].SexpString()
	envc := env.Clone()
	go envc.EvalString(code)
	return glisp.SexpNull, nil
}
