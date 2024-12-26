package errors

func NewKeyError(key string, reason error) KeyError {
	return KeyError{
		key,
		reason,
	}
}

type KeyError struct {
	Key    string
	Reason error
}

func (e KeyError) Error() string {
	return e.Reason.Error()
}
