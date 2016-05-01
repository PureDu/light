package persist

type Interface struct {
	Name    string
	Imports []string
	Methods []*Func
}

type Func struct {
	Name    string
	Doc     string
	Params  []*Param
	Returns []*Param
}

type Param struct {
	Package string
	Name    string
	Type    string
	Props   []*Param
}
