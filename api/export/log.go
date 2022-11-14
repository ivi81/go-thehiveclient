package export

import (
	"encoding/json"

	"github.com/tidwall/gjson"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/logconst"
)

//HiveLogAttr описывает атрибуты обязательные для Log
//Обязательные атрибуты:
//- Message : содержимое Log-а
//- StartDate : date of the log submission default=now
//- Status (logStatus) : status of the log (Ok or Deleted) default=Ok
//Не обязательные атрибуты:
//- Owner (string) : Указывает пользователь которому назначен Log
//- Attachment (attachment) : file attached to the log
type logAttr struct {
	Message   string          `json:"message"`
	Status    logconst.Status `json:"status"`
	StartDate int             `json:"startDate"`
	Owner     string          `json:"owner,omitempty"`
	// Attachment  `json:"attachment,omitempty"`
}

//HiveLogReq набор атрибутов которые могут быть переданы в теле запроса к api
type HiveLogReq struct {
	logAttr
}

//HiveLog набор аттрибутов которые возвращаются theHive при запросе к api
type HiveLog struct {
	HiveCommonAttr
	logAttr
}

func (h *HiveLog) UnmarshalJSON(data []byte) error {

	var err error

	data, err = preUnmarshalJSONFieldNormalize(data, cons.FieldPrefix, "")

	if err == nil {
		type typeAlias HiveLog
		h.getReplacedFeildValue(data)
		alias := typeAlias(*h)
		if err = json.Unmarshal(data, &alias); err != nil {
			return err
		}
		*h = HiveLog(alias)
	}
	return err
}

//getReplacedFeildValue извлекает из массива байт значение полей документа
//которые в зависимости от версии api имеют вариации в своих названиях (например owner и assisnee)
func (h *HiveLog) getReplacedFeildValue(data []byte) {
	h.Owner = gjson.GetBytes(data, "assignee").String()
	h.StartDate = int(gjson.GetBytes(data, "date").Int())
}
