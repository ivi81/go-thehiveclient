//opLog.go - содержит методы структуры HiveApiClient реализующие операций с сущьностями Log
package api

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"

	"gitlab.cloud.gcm/i.ippolitov/debugging"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/v1const"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/export"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/export/apiv1"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/resp"
)

//AddLog добавление Log-а в Task
func (c *HiveApiClient) AddLog(ctx context.Context, taskId string, reqBody export.HiveLogReq) (*export.HiveLog, error) {
	var (
		err error
		//req    *http.Request
		res    *resp.Response
		result *export.HiveLog
	)

	endPoint := fmt.Sprintf("http://%s/%s/log", c.Url, cons.URITaskId.Replace(taskId))

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, reqBody); err != nil {
		return nil, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("unmarshal Error : %s : %s", debugging.GetFuncName(), err.Error())
	}

	return result, err
}

//DeleteLog удаление Log-а по его id
func (c *HiveApiClient) DeleteLog(ctx context.Context, logId string) (bool, error) {
	var (
		err error
		//req    *http.Request
		res *resp.Response
	)

	endPoint := fmt.Sprintf("http://%s/%s/log", c.Url, cons.URILogId.Replace(logId))

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodDelete, nil); err != nil {
		return false, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}

	if res.Data == nil {
		return false, nil
	}

	return true, err
}

//ListTaskLogs получения списка Log-ов связанных с Task-ом имеющим указанный taskId
func (c *HiveApiClient) ListTaskLogs(ctx context.Context, taskId string, filter interface{}, sort []apiv1.E, page *apiv1.Page) ([]export.HiveLog, error) {
	var (
		err error
		//req    *http.Request
		res    *resp.Response
		result []export.HiveLog
	)

	operation := apiv1.Operation{v1const.GetTask, taskId, v1const.Logs}
	query := apiv1.CreateQueryApiV1Req(operation, filter, sort, page)

	endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URIAPIv1Query)
	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, query); err != nil {
		return nil, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}
	//result := gjson.ParseBytes()

	if err = json.Unmarshal(res.Data, &result); err != nil {
		err = fmt.Errorf("Error in function: %s : %s", debugging.GetFuncName(), err.Error())
	}

	return result, err
}
