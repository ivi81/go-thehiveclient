package export

import (
	"encoding/json"
	"time"
)

type CustomTime time.Time

func (m CustomTime) MarshalJSON() ([]byte, error) {
	//var s string
	/*if s = m.String(); s == "" {
		return nil, fmt.Errorf("MarshalError: Value '%b' is not valid for type HiveAlertStatus", int(m))
	}*/
	return json.Marshal(m)
}

func (m CustomTime) UnmarshalJSON(data []byte) error {
	/*var (
		s   string
		err error
	)
	/*if err = json.Unmarshal(data, s); err != nil {
		return fmt.Errorf("UnmarshalError: Value '%s' is not string", string(data))
	}
	if !m.strToValue(s) {
		return fmt.Errorf("UnmarshalError: Value '%s' is not valid", s)
	}*/
	json.Unmarshal(data, m)
	return nil
}
