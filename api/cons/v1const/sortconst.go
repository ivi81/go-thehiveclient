package v1const

import (
	"strings"

	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/common"
)

//go:generate stringer -type=SortConst
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

//IsValid проверка корректности значения
func (m SortConst) IsValid() bool {

	switch m {
	case
		DESC,
		ASC:

		return true
	}
	return false
}

//SetValue конвертация строки в значение типа
func (m *SortConst) SetValue(s string) bool {
	i := strings.Index(_SortConst_name, s)
	if i != -1 {

		for index, v := range _SortConst_index {
			if i-int(v) == 0 {
				*m = SortConst(index)
				return true
			}
		}
	}
	return false
}

func (m SortConst) MarshalJSON() ([]byte, error) {
	return common.MarshalConstantJSON(m)
}

func (m *SortConst) UnmarshalJSON(data []byte) error {
	return common.UnmarshalConstantJSON(m, data)
}
