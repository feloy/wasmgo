package specs

const (
	BOOL                 = "bool"
	INTEGER              = "int"
	FLOAT                = "float32"
	SPACE_SEPARATED_LIST = "[]string"
	STRING               = "string"
)

type Category struct {
	Name     string
	Elements []Element
}

type Element struct {
	Name       string
	HasInner   bool
	NoEndTag   bool
	Attributes []Attribute
}

type Attribute struct {
	Name     string
	Property string
	Key      string
	Typ      string
	Values   []string
}
