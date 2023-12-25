//resolutionstatus.go - содержит константы описывающие подтверждение успешность зарегистрированного КВ в карточке сущьности Case в IRP the Hive.
package caseconst

//go:generate stringer -type=ResolutionStatus
//go:generate enummethods -type=ResolutionStatus
type ResolutionStatus int

const (
	Indeterminate = ResolutionStatus(iota)
	FalsePositive
	TruePositive
	Other
	Duplicated
)
