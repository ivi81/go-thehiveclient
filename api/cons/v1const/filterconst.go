package v1const

import (
	"strings"

	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/common"
)

////go:generate jsonenums -type=FilterConst
//go:generate stringer -type=FilterConst
type FilterConst int

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

//IsValid проверка корректности значения
func (m FilterConst) IsValid() bool {

	switch m {
	case
		LT,
		LTE,
		GT,
		GTE,
		NE,
		NOT,
		AND,
		OR,
		BETWEEN,
		LIKE,
		CONTAINS,
		FIELD,
		VALUE,
		FROM,
		TO:

		return true
	}
	return false
}

//SetValue конвертация строки в значение типа
func (m *FilterConst) SetValue(s string) bool {
	i := strings.Index(_FilterConst_name, s)
	if i != -1 {

		for index, v := range _FilterConst_index {
			if i-int(v) == 0 {
				*m = FilterConst(index)
				return true
			}
		}
	}
	return false
}

func (m FilterConst) MarshalJSON() ([]byte, error) {
	return common.MarshalConstantJSON(m)
}

func (m *FilterConst) UnmarshalJSON(data []byte) error {
	return common.UnmarshalConstantJSON(m, data)
}
