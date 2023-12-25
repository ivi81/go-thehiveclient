//status.go - содержит константы описывающие текцщее состояние Case в IRP the Hive.
package caseconst

//go:generate stringer -type=Status
//go:generate enummethods -type=Status
type Status int

const (
	Open = Status(iota)
	Resolved
	Deleted
)
