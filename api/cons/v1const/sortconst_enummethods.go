
//sortconst_enummethods.go is auto generated by utility enummethods. DO NOT EDIT
//This file have some type methods implementing the Enumerator and json.Marshaler, json.Unmarshaler interfaces.
package v1const

import (
	
	"strings"
	"github.com/ivi81/enummethods/enumerator"
)


//MarshalJSON - реализует метод интерфейса json.Marshaler
func (m SortConst) MarshalJSON() ([]byte, error) {
	return enumerator.MarshalJSON(m)
}

//UnmarshalJSON - реализует метод интерфейса json.UnMarshaler
func (m *SortConst) UnmarshalJSON(data []byte) error {
	return enumerator.UnmarshalJSON(m, data)
}


//IsValid проверка корректности значения
//Реализует интерфейс Validator
func (m SortConst) IsValid() bool {

	switch m {
	case
		desc,
		asc,
		_fields:
		return true
	}
	return false
}


//SetValue конвертация строки в значение типа
//Реализует интерфейс Unstringer
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

