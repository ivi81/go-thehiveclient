package export

import (
	"time"

	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/common"
)

//ResponderReq  описывает атрибуты для запроса на запуск responder
//- ResponderId идентификатор респондера
//- ObjectType тип сущьности hive к которой применяется респондер
//- ObjectId уникальный внутренний идентификатор сущьности "_id"
//- CortexId идетификатор cortex
//Обязательные поля (responderId, objectType and objectId)
type CortexResponderReq struct {
	ResponderId string            `json:"responderId"`
	ObjectType  common.ObjectType `json:"objectType"`
	ObjectId    string            `json:"objectId"`
	CortexId    string            `json:"cortexId,omitempty"`
}

//CortexResponderResult описывает атрибуты ответа на запуск реподера либо получения списка отпработавших респондеров
type CortexResponderResult struct {
	CortexResponderReq
	ResponderName       string        `json:"responderName"`
	ResponderDefinition string        `json:"responderDefinition"`
	CortexJobId         string        `json:"cortexJobId"`
	Status              string        `json:"status"` // "Waiting"
	StartDate           time.Time     `json:"startDate"`
	Operations          []string      `json:"operations:omitempty"`
	Report              []interface{} `json:"report:omitempty"`
}

//CortexRespondersAction список респондеров
type ListCortexResponderAction []CortexResponderResult

/*list available Responders#
Request#
To get the list of Responders available for a Case, based on its TLP and PAP, you can call the following API:


GET /api/connector/cortex/responder/case/{id}
With:

id: Case identifier*/

type AvailableCortexResponder struct {
	Id           string
	Name         string
	Description  string
	Version      string
	DateTypeList []string
	CortexIds    []string
}

type ListAvailableCortexResponders []AvailableCortexResponder

/*

[
    {
        "id": "e33d63082066c739c07d2bbc199bfe7e",
        "name": "MALSPAM_Reply_to_user_1_0",
        "version": "1.0",
        "description": "Reply to user with an email. Applies on tasks",
        "dataTypeList": [
            "thehive:case",
            "thehive:case_task",
            "thehive:case_task_log"
        ],
        "cortexIds": [
            "Demo"
        ]
    }
]*/
