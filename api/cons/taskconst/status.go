//status.go - содержит константы описывающие состояния сущьности Task в IRP the Hive.
package taskconst

//go:generate stringer -type=Status
//go:generate enummethods -type=Status
type Status int

const (
	Waiting = Status(iota)
	InProgress
	Completed
	Cancel
)
