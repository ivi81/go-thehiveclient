// Code generated by "stringer -type=QueryConst"; DO NOT EDIT.

package v1const

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[listObservable-1]
	_ = x[getLog-2]
	_ = x[listTag-3]
	_ = x[getProfile-4]
	_ = x[listAnalyzerTemplate-5]
	_ = x[listCustomField-6]
	_ = x[getTask-7]
	_ = x[getObservable-8]
	_ = x[getAlert-9]
	_ = x[getObservableType-10]
	_ = x[waitingTask-11]
	_ = x[getOrganisation-12]
	_ = x[myTasks-13]
	_ = x[listTask-14]
	_ = x[listJob-15]
	_ = x[getCase-16]
	_ = x[listDashboard-17]
	_ = x[getAudit-18]
	_ = x[listCase-19]
	_ = x[listAudit-20]
	_ = x[listCaseTemplate-21]
	_ = x[tagAutoComplete-22]
	_ = x[getCaseTemplate-23]
	_ = x[waitingTasks-24]
	_ = x[listAction-25]
	_ = x[getTag-26]
	_ = x[getUser-27]
	_ = x[listUser-28]
	_ = x[listPage-29]
	_ = x[listLog-30]
	_ = x[getJob-31]
	_ = x[listProfile-32]
	_ = x[listObservableType-33]
	_ = x[getReportTemplate-34]
	_ = x[getDashboard-35]
	_ = x[getAction-36]
	_ = x[getPage-37]
	_ = x[listOrganisation-38]
	_ = x[getCustomField-39]
	_ = x[listAlert-40]
	_ = x[idOrName-41]
}

const _QueryConst_name = "listObservablegetLoglistTaggetProfilelistAnalyzerTemplatelistCustomFieldgetTaskgetObservablegetAlertgetObservableTypewaitingTaskgetOrganisationmyTaskslistTasklistJobgetCaselistDashboardgetAuditlistCaselistAuditlistCaseTemplatetagAutoCompletegetCaseTemplatewaitingTaskslistActiongetTaggetUserlistUserlistPagelistLoggetJoblistProfilelistObservableTypegetReportTemplategetDashboardgetActiongetPagelistOrganisationgetCustomFieldlistAlertidOrName"

var _QueryConst_index = [...]uint16{0, 14, 20, 27, 37, 57, 72, 79, 92, 100, 117, 128, 143, 150, 158, 165, 172, 185, 193, 201, 210, 226, 241, 256, 268, 278, 284, 291, 299, 307, 314, 320, 331, 349, 366, 378, 387, 394, 410, 424, 433, 441}

func (i QueryConst) String() string {
	i -= 1
	if i < 0 || i >= QueryConst(len(_QueryConst_index)-1) {
		return "QueryConst(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _QueryConst_name[_QueryConst_index[i]:_QueryConst_index[i+1]]
}