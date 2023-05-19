package rage

type schemaField struct {
	name       string
	dataType   string
	isRequired bool
	min        float64
	max        float64
	minLength  int
	maxLength  int
}

func (s *schemaField) setName(name string) {
	s.name = name
}

func (s *schemaField) setDataType(dataType string) {
	s.dataType = dataType
}

func (s *schemaField) setIsRequired(isRequired bool) {
	s.isRequired = isRequired
}

func (s *schemaField) setMin(min float64) {
	s.min = min
}

func (s *schemaField) setMax(max float64) {
	s.max = max
}

func (s *schemaField) setMinLen(minLength int) {
	s.minLength = minLength
}

func (s *schemaField) setMaxLen(maxLength int) {
	s.maxLength = maxLength
}
