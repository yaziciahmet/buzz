package buzz

type BuzzError struct {
	*BuzzInterface[error]
}

func Error() *BuzzError {
	return &BuzzError{
		BuzzInterface: Interface[error](),
	}
}
