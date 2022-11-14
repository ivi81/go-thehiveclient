package caseconst

import (
	"strings"

	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/common"
)

//go:generate stringer -type=ResolutionStatus
type ResolutionStatus int

const (
	Indeterminate = ResolutionStatus(iota)
	FalsePositive
	TruePositive
	Other
	Duplicated
)

func (m ResolutionStatus) IsValid() bool {

	switch m {
	case
		Indeterminate,
		FalsePositive,
		TruePositive,
		Other,
		Duplicated:
		return true
	}
	return false
}

//SetValue конвертация строки в значение типа
func (m *ResolutionStatus) SetValue(s string) bool {
	i := strings.Index(_ResolutionStatus_name, s)
	if i != -1 {

		for index, v := range _Status_index {
			if i-int(v) == 0 {
				*m = ResolutionStatus(index)
				return true
			}
		}
	}
	return false
}

func (m ResolutionStatus) MarshalJSON() ([]byte, error) {
	return common.MarshalConstantJSON(m)
}

func (m *ResolutionStatus) UnmarshalJSON(data []byte) error {
	return common.UnmarshalConstantJSON(m, data)
}
