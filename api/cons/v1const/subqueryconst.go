//subqueryconst.go
package v1const

//go:generate stringer -type=SubQueryConst
//go:generate enummethods -type=SubQueryConst
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
