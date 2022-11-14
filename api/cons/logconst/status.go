package logconst

import (
	"strings"

	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/common"
)

//go:generate stringer -type=Status
type Status int

const (
	Ok = Status(iota)
	Delete
)

func (m Status) IsValid() bool {

	switch m {
	case
		Ok,
		Delete:

		return true
	}
	return false
}

//SetValue конвертация строки в значение типа
func (m *Status) SetValue(s string) bool {
	i := strings.Index(_Status_name, s)
	if i != -1 {

		for index, v := range _Status_index {
			if i-int(v) == 0 {
				*m = Status(index)
				return true
			}
		}
	}
	return false
}

func (m Status) MarshalJSON() ([]byte, error) {
	return common.MarshalConstantJSON(m)
}

func (m *Status) UnmarshalJSON(data []byte) error {
	return common.UnmarshalConstantJSON(m, data)
}
