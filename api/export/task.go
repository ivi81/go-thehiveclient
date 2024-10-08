package export

import (
	"encoding/json"

	"github.com/tidwall/gjson"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/taskconst"
)

// HiveTaskAttr описывает атрибуты сущьности Task
// Обязательные атрибуты:
// - Title : заголовок Task-а
// - Status : состояние Task-a  (HIVETASKWAITING, HIVETASKINPROGRESS, HIVETASKCOMPLITED, HIVETASKCANCEL), по умолчанию HIVETASKWAITING
// - Flag : flag of the task default=false
// Не обязательные атрибуты:
// - Owner (string) : Указывает пользователь которому назначен Task, когда Status InProgress то назначается текущий пользователь
// - Description (text) : task details
// - StartDate (date) : date of the beginning of the task. This is automatically set when status is set to Open
// - EndDate (date) : date of the end of the task. This is automatically set when status is set to Completed
type HiveTaskAttr struct {
	Title       string           `json:"title,omitempty"`
	Status      taskconst.Status `json:"status"`
	Flag        bool             `json:"flag"`
	Owner       string           `json:"owner,omitempty"`
	Description string           `json:"description,omitempty"`
	StartDate   int              `json:"startDate,omitempty"` //*CustomTime `json:"startDate,omitempty"` //time.Time `json:"startDate,omitempty"`
	EndDate     int              `json:"endDate,omitempty"`   //*CustomTime `json:"endDate,omitempty"`   //time.Time `json:"endDate,omitempty"`
	DueDate     int              `json:"dueDate,omitempty"`
	Order       int              `json:"order,omitempty"`
}

// HiveTaskReq набор аттрибутов которые могут быть переданы в теле запроса к api
type HiveTaskReq struct {
	HiveTaskAttr
}

// HiveTask набор аттрибутов которые возвращаются theHive при запросе к api
type HiveTask struct {
	HiveCommonAttr
	HiveTaskAttr
}

func (h *HiveTask) UnmarshalJSON(data []byte) error {

	type typeAlias HiveTask
	var err error

	data, err = preUnmarshalJSONFieldNormalize(data, cons.FieldPrefix, "")

	if err == nil {
		h.getReplacedFeildValue(data)
		alias := typeAlias(*h)
		if err = json.Unmarshal(data, &alias); err != nil {
			return err
		}
		*h = HiveTask(alias)
	}
	return err
}

// getReplacedFeildValue извлекает из массива байт значение полей документа
// которые в зависимости от версии api имеют вариации в своих названиях (например owner и assignee)
func (h *HiveTask) getReplacedFeildValue(data []byte) {
	h.Owner = gjson.GetBytes(data, "assignee").String()
	//	h.StartDate = int(gjson.GetBytes(data, "date").Int())
}

//Если что то будем городить еще HiveTaskBackendAttr
