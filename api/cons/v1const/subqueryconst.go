package v1const

import (
	"strings"

	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/common"
)

//go:generate stringer -type=SubQueryConst
type SubQueryConst int

const (
	count = SubQueryConst(iota + 1)
	page
	observables
	assignableUsers
	filter
	linkedCases
	tasks
	logs
	aggregation
	actions
	organisations
	output
	sort
	alerts
	limitedCount
)

const (
	Count           = count
	Page            = page
	Observables     = observables
	AssignableUsers = assignableUsers
	Filter          = filter
	LinkedCases     = linkedCases
	Tasks           = tasks
	Logs            = logs
	Aggregation     = aggregation
	Actions         = actions
	Organisations   = organisations
	Output          = output
	Sort            = sort
	Alerts          = alerts
	LimitedCount    = limitedCount
)

//IsValid проверка корректности значения
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

func (m SubQueryConst) MarshalJSON() ([]byte, error) {
	return common.MarshalConstantJSON(m)
}

func (m *SubQueryConst) UnmarshalJSON(data []byte) error {
	return common.UnmarshalConstantJSON(m, data)
}
