package apiv1

import "gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/v1const"

//QueryApiV1Req описание структуры запроса к apiv1
type QueryApiV1Req struct {
	Query []M `json:"query"`
}

//AddGetOp -добавляет  к запросу операцию Get
func (q *QueryApiV1Req) AddGetOp(op v1const.QueryConst, entiryId string) {
	q.Query = append(q.Query, M{"_name": op, v1const.IdOrName.String(): entiryId})
}

//AddOther - добавляет к запросу операцию выборки тиа List, Wait  и прочие, все кроме Get
func (q *QueryApiV1Req) AddOtherOp(op v1const.QueryConst) {
	q.Query = append(q.Query, M{"_name": op})
}

func (q *QueryApiV1Req) AddSubOp(op v1const.SubQueryConst) {
	q.Query = append(q.Query, M{"_name": op})
}

//AddFilter - добавляет к запросу критерий фильтрации
func (q *QueryApiV1Req) AddFilter(filterBody interface{}) {
	m := M{"_name": v1const.Filter}
	m.Attach(filterBody)
	q.Query = append(q.Query, m)
}

//AddSort - добавляет к запросу критерий сортироваки результата
func (q *QueryApiV1Req) AddSort(sort ...E) {
	m := M{"_name": v1const.Sort, v1const.FIELDS.String(): sort}
	q.Query = append(q.Query, m)
}

//AddPage - добавляет к запросу критерий страничного вывода результата
func (q *QueryApiV1Req) AddPage(from, to int, extraData ...string) {
	m := M{
		"_name":                    v1const.Page,
		v1const.From.String():      from,
		v1const.TO.String():        to,
		v1const.ExtraData.String(): extraData,
	}
	q.Query = append(q.Query, m)
}

type Page struct {
	From, To  int
	ExtraData []string
}

//Operation описывает операции в запросе
type Operation struct {
	Op       v1const.QueryConst
	IdOrName string
	SubOp    v1const.SubQueryConst
}

//Sort описывает условия сортировки
type Sort []E

//Filter описывает условия фильтрации
type Filter interface{}

//CreateQueryApiV1Req создать запрос к api v1
func CreateQueryApiV1Req(operation Operation, filter Filter, sort Sort, page *Page) QueryApiV1Req {

	q := QueryApiV1Req{}

	if operation.Op.IsGetOperation() {
		q.AddGetOp(operation.Op, operation.IdOrName)
	} else {
		q.AddOtherOp(operation.Op)
	}

	if operation.SubOp != 0 {
		q.AddSubOp(operation.SubOp)
	}

	if filter != nil {
		q.AddFilter(filter)
	}

	if sort != nil {
		q.AddSort(sort...)
	}

	if page != nil {
		q.AddPage(page.From, page.To, page.ExtraData...)
	}

	return q
}
