//pageconst.go - содежит константы необходимые для указания в запросе к api количесва возвращаемой в ответе информации
package v1const

//go:generate stringer -type=PageConst
//go:generate enummethods -type=PageConst
type PageConst int

const (
	from = PageConst(iota)
	to
	extraData
)

const (
	From      = from
	To        = to
	ExtraData = extraData
)
