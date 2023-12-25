//status.go - содержит константы описывающие состояния сущьностей Log в IRP the Hive, в Log заносится информация отражающая ход выполнения Задачи (Task).
package logconst

//go:generate stringer -type=Status
//go:generate enummethods -type=Status
type Status int

const (
	Ok = Status(iota)
	Delete
)
