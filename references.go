package glispextensions

import (
	"errors"

	glisp "github.com/zhemao/glisp/interpreter"
)

// New reference creates a copy of the value specified, and provides a reference
// to that copy.
func NewReference(
	env *glisp.Glisp, _ string, args []glisp.Sexp,
) (glisp.Sexp, error) {
	var value glisp.Sexp = glisp.SexpNull
	if len(args) > 0 {
		value = args[0]
	}
	ref := &GlispReference{value}
	return ref, nil
}

// Dereference reports the value of a reference
func Dereference(
	env *glisp.Glisp, n string, args []glisp.Sexp,
) (glisp.Sexp, error) {
	if len(args) < 1 {
		return glisp.SexpNull, errors.New(n + " requires a parameter")
	}
	ref, okay := args[0].(*GlispReference)
	if !okay {
		return glisp.SexpNull, errors.New("parameter given is not a reference")
	}
	return ref.Value, nil
}

// Dereference reports the value of a reference
func SetReferenceValue(
	env *glisp.Glisp, n string, args []glisp.Sexp,
) (glisp.Sexp, error) {
	if len(args) < 2 {
		return glisp.SexpNull, errors.New(n + " requires two parameters")
	}
	ref, okay := args[0].(*GlispReference)
	if !okay {
		return glisp.SexpNull, errors.New(
			"first parameter given is not a reference")
	}
	ref.Value = args[1]
	return glisp.SexpNull, nil
}

func ImportReferences(env *glisp.Glisp) {
	// NewReference is named "cpy" instead of "new" to avoid any confusion
	// about the behaviour. (it copies the value into a new reference)
	env.AddFunction("&cpy", NewReference)
	env.AddFunction("&get", Dereference)
	env.AddFunction("&set", SetReferenceValue)
}
