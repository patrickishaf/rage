package rage

type ValidationReason struct {
	Value   string
	Message string
}

/*
When you call the Validate function, an object like this will be returned.
This struct will contain details of the error that made an object to fail
the validation test (if the object fails)

Fields:

	IsValid: This field will have a value of true if the object passes the validation test and false otherwise
	Error: This field will be nil if the object passes the test and will contain an error otherwise
	Reasons: If the object you want to validate passes the validation test, this slice will be empty
*/
type ValidationResult struct {
	IsValid bool
	Reasons []ValidationReason
}

func ComposeReason(value string, message string) *ValidationReason {
	return &ValidationReason{
		Value:   value,
		Message: message,
	}
}
