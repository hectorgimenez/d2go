package item

type Type struct {
	ID        int
	Name      string
	Code      string
	Throwable bool
	Beltable  bool
}

// TODO: Refactor to support parent types
func (t Type) IsType(typeName string) bool {
	return t.Code == typeName
}
