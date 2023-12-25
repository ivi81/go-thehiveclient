//objecttype.go - содержит описание названий сущьностей IRP the Hive  с учетом из специфики для различных версий api (выявлено эмперически)
package common

//go:generate enummethods -type=HiveObjectType
type HiveObjectType int

var (
	hiveObjectType = [...]string{"case", "Case", "case_artifact", "Artifact", "case_task", "Task", "case_task_log", "Log", "alert", "Alert", "organisation", "Organisation"}
)

const (
	HIVECASE = HiveObjectType(iota)
	HIVECASEAPI1
	HIVEARTIFACT
	HIVEARTIFACTAPI1
	HIVETASK
	HIVETASKAPI1
	HIVETASKLOG
	HIVETASKLOGAPI1
	HIVEALERT
	HIVEALERTAPI1
	HIVEORG
	HIVEORGAPI1
)