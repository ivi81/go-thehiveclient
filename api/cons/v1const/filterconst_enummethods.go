
//filterconst_enummethods.go is auto generated by utility enummethods. DO NOT EDIT
//This file have some type methods implementing the Enumerator and json.Marshaler, json.Unmarshaler interfaces.
package v1const

import (
	
	"strings"
	"github.com/ivi81/enummethods/enumerator"
)


//MarshalJSON - реализует метод интерфейса json.Marshaler
func (m FilterConst) MarshalJSON() ([]byte, error) {
	return enumerator.MarshalJSON(m)
}

//UnmarshalJSON - реализует метод интерфейса json.UnMarshaler
func (m *FilterConst) UnmarshalJSON(data []byte) error {
	return enumerator.UnmarshalJSON(m, data)
}


//IsValid проверка корректности значения
//Реализует интерфейс Validator
func (m FilterConst) IsValid() bool {

	switch m {
	case
		_lt,
		_lte,
		_gt,
		_gte,
		_ne,
		_not,
		_and,
		_or,
		_between,
		_like,
		_contains,
		_field,
		_value,
		_from,
		_to:
		return true
	}
	return false
}


//SetValue конвертация строки в значение типа
//Реализует интерфейс Unstringer
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
