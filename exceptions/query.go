package exceptions

type QueryException struct {
}

func (e *QueryException) Error(errorMessage string) string {
	return errorMessage
}
