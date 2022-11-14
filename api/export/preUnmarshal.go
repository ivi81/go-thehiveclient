package export

import (
	"encoding/json"
	"strings"
)

//preUnmarshalJSONFieldNormalize нормирование названий полей json объекта (к сожалению это необходимо т.к. всречаются вариации в api разных версий)
func preUnmarshalJSONFieldNormalize(data []byte, fieldPrefix, fieldSuffix string) ([]byte, error) {

	var err error
	
	if fieldPrefix == "" && fieldSuffix == "" {
		return data, err
	}

	m := map[string]json.RawMessage{}
	mm := map[string]json.RawMessage{}

	if err = json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	for k, v := range m {
		key := k
		if fieldPrefix != "" && strings.HasPrefix(key, fieldPrefix) {
			key = key[len(fieldPrefix):]
		}
		if fieldSuffix != "" && strings.HasSuffix(key, fieldSuffix) {
			key = key[:len(key)-len(fieldSuffix)]
		}
		if _, ok := mm[key]; !ok {
			mm[key] = v
		}
	}

	return json.Marshal(mm)
}
