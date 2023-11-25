package buzz

type ValidationError struct {
	Field      string
	Constraint string
	Message    string
}

func makeValidationError(field, constraint, message string) ValidationError {
	return ValidationError{
		Field:      field,
		Constraint: constraint,
		Message:    message,
	}
}

func (e ValidationError) Error() string {
	return e.Message
}
