package common

type ObjectType int

var (
	hiveObjects = [...]string{"case", "Case", "case_artifact", "Artifact", "case_task", "Task", "case_task_log", "Log", "alert", "Alert", "organisation", "Organisation"}
)

const (
	HIVECASE = ObjectType(iota)
	HIVECASEAPI1
	HIVEARTIFACT
	HIVEARTIFACTAPI1
	HIVETASK
	HIVETASKAPI1
	HIVETASKLOG
	HIVETASKLOGAPI1
	HIVEALERT
	HIVEALERTAPI1
	HIVEORG
	HIVEORGAPI1
)

func (m *ObjectType) IsValid() bool {

	switch *m {
	case
		HIVECASE,
		HIVEARTIFACT,
		HIVETASK,
		HIVETASKLOG,
		HIVEALERT,
		HIVEORG:

		return true
	}
	return false
}

func (m ObjectType) String() string {
	if m.IsValid() {
		return hiveObjects[m]
	}
	return ""
}

//SetValue конвертация строки в значение типа
func (m *ObjectType) SetValue(s string) bool {
	for i, v := range hiveObjects {
		if v == s {
			*m = ObjectType(i)
			return true
		}
	}
	return false
}

func (m ObjectType) MarshalJSON() ([]byte, error) {
	return MarshalConstantJSON(m)
}

func (m *ObjectType) UnmarshalJSON(data []byte) error {
	return UnmarshalConstantJSON(m, data)
}
