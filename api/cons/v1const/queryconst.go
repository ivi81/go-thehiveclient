//queryconst.go - содежит константы необходимые для указания в запросе к api действия c данными
package v1const

//go:generate stringer -type=QueryConst
//go:generate enummethods -type=QueryConst

type QueryConst int

const (
	listObservable = QueryConst(iota + 1)
	getLog
	listTag
	getProfile
	listAnalyzerTemplate
	listCustomField
	getTask
	getObservable
	getAlert
	getObservableType
	waitingTask
	getOrganisation
	myTasks
	listTask
	listJob
	getCase
	listDashboard
	getAudit
	listCase
	listAudit
	listCaseTemplate
	tagAutoComplete
	getCaseTemplate
	waitingTasks
	listAction
	getTag
	getUser
	listUser
	listPage
	listLog
	getJob
	listProfile
	listObservableType
	getReportTemplate
	getDashboard
	getAction
	getPage
	listOrganisation
	getCustomField
	listAlert
	idOrName
)

const (
	ListObservable       = listObservable
	GetLog               = getLog
	ListTag              = listTag
	GetProfile           = getProfile
	ListAnalyzerTemplate = listAnalyzerTemplate
	ListCustomField      = listCustomField
	GetTask              = getTask
	GetObservable        = getObservable
	GetAlert             = getAlert
	GetObservableType    = getObservableType
	WaitingTask          = waitingTask
	GetOrganisation      = getOrganisation
	MyTasks              = myTasks
	ListTask             = listTask
	ListJob              = listJob
	GetCase              = getCase
	ListDashboard        = listDashboard
	GetAudit             = getAudit
	ListCase             = listCase
	ListAudit            = listAudit
	ListCaseTemplate     = listCaseTemplate
	TagAutoComplete      = tagAutoComplete
	GetCaseTemplate      = getCaseTemplate
	WaitingTasks         = waitingTasks
	ListAction           = listAction
	GetTag               = getTag
	GetUser              = getUser
	ListUser             = listUser
	ListPage             = listPage
	ListLog              = listLog
	GetJob               = getJob
	ListProfile          = listProfile
	ListObservableType   = listObservableType
	GetReportTemplate    = getReportTemplate
	GetAction            = getAction
	GetPage              = getPage
	ListOrganisation     = listOrganisation
	GetCustomField       = getCustomField
	ListAlert            = listAlert
	IdOrName             = idOrName
)

//IsGetOperation проверка на get операцию
func (m QueryConst) IsGetOperation() bool {
	switch m {
	case
		GetLog,
		GetProfile,
		GetTask,
		GetObservable,
		GetAlert,
		GetObservableType,
		GetOrganisation,
		GetCase,
		GetAudit,
		GetCaseTemplate,
		GetTag,
		GetUser,
		GetJob,
		GetReportTemplate,
		GetAction,
		GetPage,
		GetCustomField:
		return true
	}
	return false
}

//IsListOperation проверка на list операцию
func (m QueryConst) IsListOperation() bool {
	switch m {
	case
		ListObservable,
		ListTag,
		ListAnalyzerTemplate,
		ListCustomField,
		ListTask,
		ListJob,
		ListDashboard,
		ListCase,
		ListAudit,
		ListCaseTemplate,
		ListAction,
		ListUser,
		ListPage,
		ListLog,
		ListProfile,
		ListObservableType,
		ListOrganisation,
		ListAlert:

		return true
	}
	return false
}

//IsOtherOperation проверка на просие операции не являющиеся get и list
func (m QueryConst) IsOtherOperation() bool {
	switch m {
	case
		WaitingTask,
		MyTasks,
		TagAutoComplete,
		WaitingTasks:
		return true
	}
	return false
}

//IsValid проверка корректности значения
/*func (m QueryConst) IsValid() bool {
	if m.IsGetOperation() || m.IsListOperation() || m.IsOtherOperation() {
		return true
	}
	return false
}
*/
