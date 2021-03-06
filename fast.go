//+build amd64

package loc

//go:noinline

// FastCaller returns information about the calling goroutine's stack. The argument s is the number of frames to ascend, with 0 identifying the caller of Caller.
//
// It's hacked version of runtime.Caller with no allocs.
func FastCaller(s int) (r PC) {
	return fastCaller(s)
}

//go:noinline

// FastFuncentry returns information about the calling goroutine's stack. The argument s is the number of frames to ascend, with 0 identifying the caller of Caller.
//
// It's hacked version of runtime.Callers -> runtime.CallersFrames -> Frames.Next -> Frame.Entry with no allocs.
func FastFuncentry(s int) (r PC) {
	r = fastCaller(s)

	return r.FuncEntry()
}

func fastCaller(s int) (c PC)
