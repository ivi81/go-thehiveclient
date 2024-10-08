// reqPrimitive.go  - содержит структуры примитвов данных применяемых при постороении запроса к apiv1
package apiv1

import (
	"encoding/json"
	"fmt"

	"github.com/ivi81/enummethods/enumerator"
	//"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/common"
)

// E - примитив описывающий поле json-документа в виде ключ значение, при Marshal-инге получаем {Key:Value}
type E struct {
	Key   interface{}
	Value interface{}
}

func (fp E) MarshalJSON() ([]byte, error) {
	kb, err := json.Marshal(fp.Key)
	if err != nil || len(kb) == 2 {
		return nil, err
	}
	vb, err := json.Marshal(fp.Value)
	if err != nil || string(vb) == "null" {
		return nil, err
	}

	s := fmt.Sprintf("{%s:%s}", string(kb), string(vb))

	return []byte(s), nil
}

// D - примитив описывающий массив элементов представленных в виде {Key:Value} в json-документе,
// где Value произвольный json-документ при Marshal-инге получаем [{Key1:Value1},{Key2:Value2},...]
type D []E

func (e E) KeyToString() string {
	switch vv := e.Key.(type) {
	case enumerator.Enumerator: //common.ConstantStringer:
		return vv.String()
	case string:
		return vv
	}
	return ""
}

//ToMap преводит D -> M
/*func (d D) ToMap() M {
	m := M{}
	for _, v := range d {
ok,strKey:=v.KeyToString()
if ok
		m[strKey] = v.Value

	return m
}
*/
//M - примитив описывающий совокупность полей json-документе представленных в виде {Key:Value},
//при при Marshal-инге получаем {MapKey1:Value1,MapKey2:Value2,...}
type M map[string]interface{}

// AddE добавить елемент типа Е в M
func (m *M) AddE(elem E) {
	strKey := elem.KeyToString()
	if strKey != "" {
		(*m)[strKey] = elem.Value
	}
}

// Attach добавить в M новые элементы трансформировав их из элементов D, E
func (m *M) Attach(elem ...interface{}) {

	for _, v := range elem {
		switch val := v.(type) {
		case M:
			for k, v := range val {
				(*m)[k] = v
			}
		case D:
			for _, v := range val {
				m.AddE(v)
			}
		case E:
			m.AddE(val)
		}
	}
}

//ToSlice переводит M -> D
/*func (m M) ToSlice() D {
	d := D{}
	for k, v := range m {
		d = append(d, E{k, v})
	}
	return d
}*/

//AddToMap производит добавление структуры данных составленной из произвольных компбинаций типов E,D,M к текущему экземпляру M
//func (m *M) AddToMap(interface{})

// A  - примитив описывающий массив произвольных json-документов
type A []interface{}
