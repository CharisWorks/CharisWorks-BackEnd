package utils

type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	return e.Message
}
func (e *InternalError) SetError(msg string) {
	e.Message = msg
}
