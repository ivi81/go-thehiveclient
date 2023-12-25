//const.go - содержит описание интерфейсов которому должны соответствовать описываемые константами пречислимые типы
//DEPRICATED. Было перенесено в github.com/ivi81/enummethods
package common

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//NotStringConstanter интерфейсный тип объявляющий функции для работы с нестроковыми константами
type NotStringConstanter interface {
	ConstantStringer
	ConstantValidator
	ConstantUnstringer
}

//ConstantStringer  интерфейсный тип объявляющий функции для преобразования
//константного (не строкового) типа в строку
type ConstantStringer interface {
	fmt.Stringer
}

//ConstantStringer  интерфейсный тип объявляющий функции для преобразования строки
//в значение константного (не строкового) типа
type ConstantUnstringer interface {
	SetValue(s string) bool
}

//ConstantValidator интерфейсный тип объявляющий функции для проверки допустимости значения
//константного (не строкового) типа
type ConstantValidator interface {
	IsValid() bool
}

//MarshalConstantJSON маршалинг значений не строковых константных типов в
//строковое значение json
func MarshalConstantJSON(c ConstantStringer) ([]byte, error) {
	var s string

	if s = c.String(); s == "" {

		err := &json.UnsupportedValueError{
			Value: reflect.ValueOf(c),
			Str:   fmt.Sprintf("\"%b\"", c),
		}
		return nil, err
	}

	return json.Marshal(s)
}

//UnmarshalConstantJSON анмаршалинг строкового значания поля в json-документе в значение не строкового константного типа
func UnmarshalConstantJSON(c ConstantUnstringer, data []byte) error {
	var (
		s   string
		err error
	)
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}
	if !c.SetValue(s) {

		err = &json.UnsupportedValueError{
			Value: reflect.ValueOf(c),
			Str:   fmt.Sprintf("\"%s\" is not valid value", s),
		}
		return err
	}
	return err
}
