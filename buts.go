package buts

type Kind int

const (
	Invalid Kind = iota
	Element
	Structure
	Table
)

type TypeSystem interface {
	// Creates a Type corresponding to the definition
	New(kind Kind, name string) Type
}

type Modifier interface {
	//
	//
	Modify(t Type, v interface{})
}
