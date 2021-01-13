package buts

type Kind int

const (
	Invalid Kind = iota
	Element
	Structure
	DbTable
	MemTable
)

type GoType int

const (
	GoInvalid GoType = iota
	GoString
	GoInt
	GoStructure
)

func (g GoType) String() string {
	switch g {
	case GoInvalid:
		return "GoInvalid"
	case GoInt:
		return "GoInt"
	case GoString:
		return "GoString"
	case GoStructure:
		return "GoStructure"
	default:
		return "GoInvalid"
	}
}

type DbType int

const (
	DbInvalid DbType = iota
	DbVarchar
	DbInt
)

func (k Kind) String() string {
	switch k {
	case Invalid:
		return "invalid"
	case Element:
		return "element"
	case Structure:
		return "structure"
	case DbTable:
		return "dbtable"
	case MemTable:
		return "memtable"
	}
	return "undefined"
}

type TypeSystem interface {
	// Creates a Type corresponding to the definition
	New(kind Kind, name string) Type
	Register([]ElementReg, []StructureReg, []TableReg) error
}

type ElementReg struct {
	Name           string
	Description    string
	GoType         GoType
	DbType         DbType
	DbLength       int
	DbDecimals     int
	Tags           string
	Domain         string
	DomainTable    string
	DomainGoColumn string
	DomainDbColumn string
	Conversion     string
}

type StructureReg struct {
	Name        string
	Description string
	Tags        string
	Items       []FeldReg
}

type FeldReg struct {
	Name        string
	Description string
	Kind        Kind
	Type        string
}

type TableReg struct {
	Name        string
	Description string
	Tags        string
	TableType   Kind
	Kind        Kind
	Type        string
	SQLName     string
	Indizes     []TableIndex
}

type TableIndex struct {
	Name   string
	PK     bool
	Unique bool
	Fields []string
}

type Modifier interface {
	//
	//
	Modify(t Type, v interface{})
}
