package schema

const (
	Long   = "long"
	Float  = "float"
	String = "string"
	Date   = "date"
	Bool   = "boolean"
	Array  = "array"
	Map    = "map"
	Json   = "json"
)

type Entry struct {
	Key       string  `json:"key"`
	ValueType string  `json:"valtype"`
	ElemType  string  `json:"elemtype,omitempty"`
	Schema    []Entry `json:"schema,omitempty"`
}

type Schemas []Entry

func (s Schemas) Len() int {
	return len(s)
}
func (s Schemas) Less(i, j int) bool {
	return s[i].Key < s[j].Key
}
func (s Schemas) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type Data map[string]interface{}
type Datas []Data
