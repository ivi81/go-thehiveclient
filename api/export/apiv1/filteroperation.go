//filteroperation.go - содержит функции конструкторы операций применяемых при построении фильтра выборки документов theHive
package apiv1

import "gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/v1const"

//And - конструирование оператора "and"
//пример {"_and":[...]}
func And(elem ...E) E {
	return naryConstructor(v1const.AND, elem...)
}

//Or - конструирование оператора "or"
//пример {"_or":[...]}
func Or(elem ...E) E {
	return naryConstructor(v1const.OR, elem...)
}

//Like - конструирование оператора "like"
//пример: "_like": {"_field": "title","_value": "IVI"}
func Like(field string, value interface{}) E {
	return E{v1const.LIKE, M{
		v1const.FIELD.String(): field,
		v1const.VALUE.String(): value,
	}}
}

//Between - конструирование оператора "between"
//пример: "_between": {	"_field": "_updatedAt",	"_from": 1666216800963,	"_to": 1666907999963 }
func Beetwen(field string, from interface{}, to interface{}) E {
	if field == "" || from == nil || to == nil {
		return E{}
	}

	return E{v1const.BETWEEN, M{
		v1const.FIELD.String(): field,
		v1const.FROM.String():  from,
		v1const.TO.String():    to,
	}}
}

//Not - конструирование оператора "not"
//пример: "_not":{"number":111111}
func Not(value interface{}) E {
	return E{v1const.NOT, value}
}

//Eq -  конструирование оператора "eq" (equal), в web-интерфейсе theHive он же "=", он же "Yes"
//пример: {"_field":"flag","_value":true}
//{"_name":"filter","_field":"number","_value":0}
func Eq(field string, value interface{}) E {
	return E{field, value}
}

//Lt - конструирование оператора "lt" (less then)
//пример: {"_lt":{"number":111111}}
func Lt(field string, value interface{}) E {
	return unaryContructor(v1const.LT, field, value)
}

//Lte - конструирование оператора "lte" (less then equal)
//пример: {"_lte":{"number":111111}}
func Lte(field string, value interface{}) E {
	return unaryContructor(v1const.LTE, field, value)
}

//Gt - конструирование оператора "gt" (grate then)
//пример: {"_gt":{"number":111111}}
func Gt(field string, value interface{}) E {
	return unaryContructor(v1const.GT, field, value)
}

//Gte - конструирование оператора "gte" (grate then equal)
//пример: {"_gte":{"number":111111}}
func Gte(field string, value interface{}) E {
	return unaryContructor(v1const.GTE, field, value)
}

//Ne - конструирование оператора "ne" (!=)
//пример: {"_ne":{"number":111111}}
func Ne(field string, value interface{}) E {
	return unaryContructor(v1const.NE, field, value)
}
//unaryContructor - конструктор 
func unaryContructor(op v1const.FilterConst, key string, value interface{}) E {
	return E{op, E{key, value}}
}

func naryConstructor(op v1const.FilterConst, elem ...E) E {
	if len(elem) == 1 {
		return elem[0]
	}
	return E{op, D(elem)}
}
