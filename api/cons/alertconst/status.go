//status.go - содержит константы описывающие состояния сущьностей Alert в IRP the Hive.
package alertconst

//go:generate stringer -type=Status
//go:generate enummethods -type=Status
type Status int

const (
	New = Status(iota)
	Updated
	Ignored
	Imported
)
