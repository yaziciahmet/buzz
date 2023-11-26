package buzz

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

type FieldErrorAggregator struct {
	Errors []FieldError
}

func NewFieldErrorAggregator() *FieldErrorAggregator {
	return &FieldErrorAggregator{}
}

func (a *FieldErrorAggregator) Add(err FieldError) {
	a.Errors = append(a.Errors, err)
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

func (a *FieldErrorAggregator) Merge(aggr *FieldErrorAggregator) {
	a.Errors = append(a.Errors, aggr.Errors...)
}

func (a *FieldErrorAggregator) Error() string {
	return ""
}
