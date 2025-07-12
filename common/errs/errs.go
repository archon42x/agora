package errs

type ErrorCode uint32

const (
	REGISTER_ERROR ErrorCode = 10001
	LOGIN_ERROR    ErrorCode = 10002
)
