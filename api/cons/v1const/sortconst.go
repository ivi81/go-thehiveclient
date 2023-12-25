//sortconst.go - содежит константы необходимые для указания в запросе к api параметров сортировки результата запроса
package v1const

//go:generate stringer -type=SortConst
//go:generate enummethods -type=SortConst
type SortConst int

const (
	desc = SortConst(iota)
	asc
	_fields
)

const (
	DESC   = desc
	ASC    = asc
	FIELDS = _fields
)
