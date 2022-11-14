package resp

//Response ответ о theHive прошедший первичную обработку
type Response struct {
	StatusCode int
	Err        *RespError
	Data       []byte
}

//RespError описывает формат полей ошибки возвращаемой при неудачном запросе к theHive
type RespError struct {
	Type        string `json:"type"`
	Message     string `json:"message"`
	Unprocessed bool
}
