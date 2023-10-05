package exception

type NotFoundError struct {
	Message string
}

func (n NotFoundError) Error() string {
	return n.Message
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{Message: message}
}
