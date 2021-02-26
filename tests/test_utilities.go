package tests

type BuilderFlags []string

func (f *BuilderFlags) String() string {
	var resultString string
	for _, builder := range *f {
		resultString += builder
		resultString += ", "
	}
	return resultString
}

func (f *BuilderFlags) Set(value string) error {
	*f = append(*f, value)
	return nil
}
