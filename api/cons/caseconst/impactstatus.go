//impactstatus.go - содержит константы описывающие наличие влияния в зарегистрированном КВ в карточке сущьности Case в IRP the Hive.
package caseconst

//go:generate stringer -type=ImpactStatus
//go:generate enummethods -type=ImpactStatus
type ImpactStatus int

const (
	NoImpact = ImpactStatus(iota)
	WithImpact
	NotApplicable
)
