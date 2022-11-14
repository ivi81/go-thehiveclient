package resp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gitlab.cloud.gcm/i.ippolitov/debugging"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/resp/cons"
)

//Handler обработчик ответа на запрос
func Handler(resp *http.Response) (*Response, error) {

	var (
		err      error
		respBody []byte
	)

	defer resp.Body.Close()

	if respBody, err = io.ReadAll(resp.Body); err != nil {
		err := fmt.Errorf("error Read Response Body : %s : %s", debugging.GetFuncName(), err.Error())
		return nil, err
	}

	response := &Response{
		StatusCode: resp.StatusCode,
	}

	switch resp.StatusCode {
	case cons.HIVEReqErrUthorization,
		cons.HIVEReqErrAuthentication,
		cons.HIVEReqErrNotFound,
		cons.HIVEReqErrAttrChecking:
		response.Err, err = initRespError(resp, respBody)
	case cons.HIVEReqSuccessfully, cons.HIVEReqCreated, cons.HIVEReqDeleted:
		response.Data = respBody
		response.StatusCode = resp.StatusCode
	default:
		response.Err, err = initUnprocessedError(resp, respBody)
	}
	return response, err
}

//initUnprocwssedError сообщение о необрабатываемой ошибке
func initUnprocessedError(httpResp *http.Response, Data []byte) (*RespError, error) {
	var (
		respErr *RespError
		err     error
	)

	if respErr, err = initRespError(httpResp, Data); respErr == nil {
		return respErr, err
	}
	respErr.Unprocessed = true

	return respErr, fmt.Errorf("unprocessed httpError: %s", err)
}

//initRespError обработка сообщения об ощибке
func initRespError(httpResp *http.Response, Data []byte) (*RespError, error) {
	var (
		respErr RespError
		err     error
		errMsg  string
	)
	if err = json.Unmarshal(Data, &respErr); err != nil {
		if err = json.Unmarshal(Data, errMsg); err != nil {
			err = fmt.Errorf("unmarshalError: %s : %s", debugging.GetFuncName(), err.Error())
			return nil, err
		}
		respErr.Message = fmt.Sprintf("%s", errMsg)
	}

	err = fmt.Errorf("%s, statusCode: %s, message: %s", httpResp.Proto, httpResp.Status, respErr.Message)

	return &respErr, err
}
