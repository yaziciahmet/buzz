package buzz

import "fmt"

const (
	invalidTypeMsg = "invalid type. expected: %s received: %T"
	notNullableMsg = "%s is not nullable"
	nonEmptyMsg    = "%s can not be empty"
)

type FieldError struct {
	Name       string
	Constraint string
	Message    string
}

func MakeFieldError(name, constraint, message string) FieldError {
	return FieldError{
		Name:       name,
		Constraint: constraint,
		Message:    message,
	}
}

func (e FieldError) Error() string {
	return e.Message
}

func notNullableFieldErr(name string) FieldError {
	return MakeFieldError(name, "Nonnil", fmt.Sprintf(notNullableMsg, name))
}

func nonEmptyFieldErr(name string) FieldError {
	return MakeFieldError(name, "Nonempty", fmt.Sprintf(nonEmptyMsg, name))
}

type FieldErrorAggregator struct {
	Errors []FieldError
}

func NewFieldErrorAggregator() *FieldErrorAggregator {
	return &FieldErrorAggregator{}
}

func (a *FieldErrorAggregator) Error() string {
	if a.Empty() {
		return ""
	}

	return a.Errors[0].Error()
}

func (a *FieldErrorAggregator) Handle(err error) error {
	switch e := err.(type) {
	case FieldError:
		a.Add(e)
		break
	case *FieldErrorAggregator:
		a.Merge(e)
		break
	default:
		return err
	}

	return nil
}

func (a *FieldErrorAggregator) Add(err FieldError) {
	a.Errors = append(a.Errors, err)
}

func (a *FieldErrorAggregator) Merge(aggr *FieldErrorAggregator) {
	if !aggr.Empty() {
		a.Errors = append(a.Errors, aggr.Errors...)
	}
}

func (a *FieldErrorAggregator) Empty() bool {
	return len(a.Errors) == 0
}

func (a *FieldErrorAggregator) OrNil() error {
	if a.Empty() {
		return nil
	}

	return a
}
