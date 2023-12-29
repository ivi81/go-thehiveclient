//impactstatus_enummethods.go is auto generated by utility enummethods. DO NOT EDIT
//This file have some type methods implementing the Enumerator and json.Marshaler, json.Unmarshaler interfaces.
package caseconst

import (
	"strings"

	"github.com/ivi81/enummethods/enumerator"
)

//MarshalJSON - реализует метод интерфейса json.Marshaler
func (m ImpactStatus) MarshalJSON() ([]byte, error) {
	return enumerator.MarshalJSON(m)
}

//UnmarshalJSON - реализует метод интерфейса json.UnMarshaler
func (m *ImpactStatus) UnmarshalJSON(data []byte) error {
	return enumerator.UnmarshalJSON(m, data)
}

//IsValid проверка корректности значения
//Реализует интерфейс Validator
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
//Реализует интерфейс Unstringer
func (m *ImpactStatus) SetValue(s string) bool {
	i := strings.Index(_ImpactStatus_name, s)
	if i != -1 {

		for index, v := range _ImpactStatus_index {
			if i-int(v) == 0 {
				*m = ImpactStatus(index)
				return true
			}
		}
	}
	return false
}