//filterconst.go - содежит константы необходимые для конструирования в коде фильтров применяемых в запросах
//к api IRP the Hive
package v1const

//go:generate stringer -type=FilterConst
//go:generate enummethods -type=FilterConst
type FilterConst int

//названия которые используются для конструирования текста запросов
const (
	_lt = FilterConst(iota)
	_lte
	_gt
	_gte
	_ne
	_not
	_and
	_or
	_between
	_like
	_contains
	_field
	_value
	_from
	_to
)

//Экспортируемые названия которые применяются в непосредственно в коде
const (
	LT       = _lt
	LTE      = _lte
	GT       = _gt
	GTE      = _gte
	NE       = _ne
	NOT      = _not
	AND      = _and
	OR       = _or
	BETWEEN  = _between
	LIKE     = _like
	CONTAINS = _contains
	FIELD    = _field
	VALUE    = _value
	FROM     = _from
	TO       = _to
)
