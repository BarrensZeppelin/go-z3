package z3

/*
#cgo LDFLAGS: -lz3
#include <z3.h>
#include <stdlib.h>
*/
import "C"
import (
	"runtime"
	"unsafe"
)

func (ctx *Context) ForAll(vars []Value, body Bool) Bool {
	cvars := make([]C.Z3_app, len(vars))
	for i, c := range vars {
		cvars[i] = *(*C.Z3_app)(unsafe.Pointer(&c.impl().c))
	}

	val := wrapValue(ctx, func() C.Z3_ast {
		var cvp *C.Z3_app
		if len(cvars) > 0 {
			cvp = &cvars[0]
		}

		return C.Z3_mk_forall_const(
			ctx.c, C.uint(0),
			C.uint(len(vars)),
			cvp,
			C.uint(0),
			nil,
			body.c,
		)
	})

	runtime.KeepAlive(vars)
	runtime.KeepAlive(body)
	return Bool(val)
}
