package caseconst

import (
	"strings"

	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/common"
)

//go:generate stringer -type=ImpactStatus
type ImpactStatus int

const (
	NoImpact = ImpactStatus(iota)
	WithImpact
	NotApplicable
)

func (m ImpactStatus) IsValid() bool {

	switch m {
	case
		NoImpact,
		WithImpact,
		NotApplicable:
		return true
	}
	return false
}

//SetValue конвертация строки в значение типа
func (m *ImpactStatus) SetValue(s string) bool {
	i := strings.Index(_ImpactStatus_name, s)
	if i != -1 {

		for index, v := range _Status_index {
			if i-int(v) == 0 {
				*m = ImpactStatus(index)
				return true
			}
		}
	}
	return false
}

func (m ImpactStatus) MarshalJSON() ([]byte, error) {
	return common.MarshalConstantJSON(m)
}

func (m *ImpactStatus) UnmarshalJSON(data []byte) error {
	return common.UnmarshalConstantJSON(m, data)
}
