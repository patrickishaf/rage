package rage

func buildFieldNotFoundError(fieldName string) string {
	return "the field '" + fieldName + "' was not found in the object"
}
