
//subqueryconst_enummethods.go is auto generated by utility enummethods. DO NOT EDIT
//This file have some type methods implementing the Enumerator and json.Marshaler, json.Unmarshaler interfaces.
package v1const

import (
	
	"strings"
	"github.com/ivi81/enummethods/enumerator"
)


//MarshalJSON - реализует метод интерфейса json.Marshaler
func (m SubQueryConst) MarshalJSON() ([]byte, error) {
	return enumerator.MarshalJSON(m)
}

//UnmarshalJSON - реализует метод интерфейса json.UnMarshaler
func (m *SubQueryConst) UnmarshalJSON(data []byte) error {
	return enumerator.UnmarshalJSON(m, data)
}


//IsValid проверка корректности значения
//Реализует интерфейс Validator
func (m SubQueryConst) IsValid() bool {

	switch m {
	case
		count,
		page,
		observables,
		assignableUsers,
		filter,
		linkedCases,
		tasks,
		logs,
		aggregation,
		actions,
		organisations,
		output,
		sort,
		alerts,
		limitedCount:
		return true
	}
	return false
}


//SetValue конвертация строки в значение типа
//Реализует интерфейс Unstringer
func (m *SubQueryConst) SetValue(s string) bool {
	i := strings.Index(_SubQueryConst_name, s)
	if i != -1 {

		for index, v := range _SubQueryConst_index {
			if i-int(v) == 0 {
				*m = SubQueryConst(index)
				return true
			}
		}
	}
	return false
}

