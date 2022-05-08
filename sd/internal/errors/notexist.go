package errors

type NotExist interface {
	NotExist() bool
}

type errorNotExist struct {
	message string
}

var ErrNotExist = NewNotExistError("")

func NewNotExistError(message string) error {
	return &errorNotExist{
		message: message,
	}
}

func (e *errorNotExist) NotExist() bool {
	return true
}

func (e *errorNotExist) Error() string {
	return e.message
}
