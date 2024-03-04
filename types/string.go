package types

type StringType struct {
	Value string
}

func (s StringType) ToString() string {
	return s.Value
}
