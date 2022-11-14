package v1const

import (
	"strings"

	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/common"
)

//go:generate stringer -type=PageConst
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

//IsValid проверка корректности значения
func (m PageConst) IsValid() bool {

	switch m {
	case
		From,
		To,
		ExtraData:

		return true
	}
	return false
}

//SetValue конвертация строки в значение типа
func (m *PageConst) SetValue(s string) bool {
	i := strings.Index(_PageConst_name, s)
	if i != -1 {

		for index, v := range _PageConst_index {
			if i-int(v) == 0 {
				*m = PageConst(index)
				return true
			}
		}
	}
	return false
}

func (m PageConst) MarshalJSON() ([]byte, error) {
	return common.MarshalConstantJSON(m)
}

func (m *PageConst) UnmarshalJSON(data []byte) error {
	return common.UnmarshalConstantJSON(m, data)
}
