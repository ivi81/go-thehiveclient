package export

import (
	"encoding/json"
	"time"

	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/caseconst"
)

//HiveCaseReqAttr описывает атрибуты обязательные для Case:
//- Title : заголовок Case
//- Description : описание Case
//- Severity (number) : строгость Case (1: low; 2: medium; 3: high) default=2
//- StartDate (date) : Дата и время начала рассмотрения Case, поумолчанию берется текущее время "now"
//- Owner (string) : пользователь которому назначен Case, поумолчанию назначается пользователь его созадавший
//- Flag (boolean) : flag of the case default=false
//- Tlp (number) : значение по протоколу TLP (0: white; 1: green; 2: amber; 3: red), по умолчанию "2"
//- Tags (multi-string) : список тэгов которыми помечен Case, по умолчанию пустой массив строк
type HiveCaseReqAttr struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Severity    int       `json:"severity,omitempty"`
	StartDate   time.Time `json:"startDate,omitempty"`
	Owner       string    `json:"owner"`
	Flag        bool      `json:"flag"`
	Tags        []string  `json:"tags"`
	Tlp         int       `json:"tlp"`
}

//HiveCaseOptionAttr описывает атрибуты не обязательные для Case:
//- ResolutionStatus (caseResolutionStatus) : решение по Case (Indeterminate, FalsePositive, TruePositive, Other or Duplicated)
//- ImpactStatus (caseImpactStatus) : impact status of the case (NoImpact, WithImpact or NotApplicable)
//- Summary : краткое описание case, заполняется при закрытии Case-а
//- EndDate : дата принятия решения
type HiveCaseOptionAttr struct {
	ImpactStatus     caseconst.ImpactStatus     `json:"impactStatus,omitempty"`
	ResolutionStatus caseconst.ResolutionStatus `json:"resolutionStatus,omitempty"` //Indeterminate, FalsePositive, TruePositive, Other or Duplicated
	Summary          string                     `json:"summary,omitempty"`
	EndDate          time.Time                  `json:"endDate,omitempty"`
}

//HiveCaseBackendAttr описывает атрибуты формируемые Hive-ом и возвращаемые при запросах к нему
//- Status (caseStatus) : состояние Case (Open, Resolved or Deleted), по умолчанию "Open"
//- CaseId (number) : идентификатор Сase (генерируется атоматически )
//- MergeInto (string) : идентификатор Сase созданого в реультате слияния других Case-ов
//- MergeFrom (multi-string) : идентификаторы Case-ов участвовавших в слияние
type HiveCaseBackendAttr struct {
	Status    caseconst.Status `json:"status"` //status of the case (Open, Resolved or Deleted) default=Open
	CaseId    int              `json:"caseId"`
	MergeInfo string           `json:"mergeInto"`
	MergeFrom []string         `json:"mergeFrom"`
}

//HiveCaseReq набор аттрибутов которые могут быть переданы в теле запроса к api
type HiveCaseReq struct {
	HiveCaseReqAttr
	HiveCaseOptionAttr
}

//HiveCase набор аттрибутов которые возвращаются theHive при запросе к api
type HiveCase struct {
	HiveCommonAttr
	HiveCaseReq
	HiveCaseBackendAttr

	/*Pap          int          `json:"pap"`
	CustomFields CustomFields `json:"customFields"`
	Stats        Stats        `json:"stats"`
	Permissions  []string     `json:"permissions"`*/
}

func (h *HiveCase) UnmarshalJSON(data []byte) error {
	var err error
	data, err = preUnmarshalJSONFieldNormalize(data, cons.FieldPrefix, "")
	if err == nil {
		json.Unmarshal(data, h)
	}

	return err
}

/*
type HiveTaskInfo struct {
	Title      string    `json:"title"`
	Group      string    `json:"group,omitempty"`
	Decription string    `json:"description,omitempty"`
	Owner      string    `jsonP:"owner,omitempty"`
	Status     string    `json:"status,omitempty"`
	Flag       bool      `json:"flag"`
	StartDate  time.Time `json:"startDate,omitempty"`

	Order      int       `json:"order,omitempty"`
	DueDate    time.Time `json:"dueDate,omitempty"`
}

//HiveTask - структура
type HiveTask struct {
	HiveCommonObj
	HiveTaskInfo
}

/*
"id": "~4264",
"_id": "~4264",
"createdBy": "[email protected]",
"createdAt": 1630684502715,
"updatedBy": "[email protected]",
  "updatedAt": 1630685486000,
"_type": "case_task",


"title": "Malware analysis",
"group": "identification",
"description": "Analysis of the file to identify the malware",
"owner": "[email protected]",
"status": "InProgress",
"flag": false,
"startDate": 1630683608000,
"endDate": 1630684608000,
"order": 3,
"dueDate": 1630694608000
type Alert struct {
}*/

//type CustomFields struct{}
//type Stats struct{}
