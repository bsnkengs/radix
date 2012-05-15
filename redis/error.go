package redis

type ErrorFlag uint8

const (
	ErrorRedis ErrorFlag = 1 << iota
	ErrorConnection
	ErrorLoading
	ErrorAuth
	ErrorParse
	ErrorInvalidReply
	ErrorTimeout
)

type Error struct {
	msg   string
	flags ErrorFlag
}

// newError creates a new Error.
func newError(msg string, flags ...ErrorFlag) *Error {
	err := new(Error)
	err.msg = msg
	for _, f := range flags {
		err.flags |= f
	}
	return err
}

// newErrorExt creates a new Error with flags of the given error.
func newErrorExt(msg string, err *Error, flags ...ErrorFlag) *Error {
	return newError(msg, append(flags, err.flags)...)
}

// Error returns a string representation of the error.
func (e *Error) Error() string {
	return radixError(e.msg)
}

// Test returns true, if any of the given error flags is set in the error, otherwise false.
func (e *Error) Test(flags ...ErrorFlag) bool {
	for _, f := range flags {
		if e.flags&f > 0 {
			return true
		}
	}
	return false
}

func radixError(msg string) string {
	return "redis: " + msg
}