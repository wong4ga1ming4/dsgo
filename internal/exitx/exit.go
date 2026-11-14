package exitx

// Error is a command failure with a process exit code.
type Error struct {
	Code int
	Msg  string
}

func (e *Error) Error() string { return e.Msg }

func Usage(msg string) *Error { return &Error{Code: 2, Msg: msg} }

func Fail(msg string) *Error { return &Error{Code: 1, Msg: msg} }
