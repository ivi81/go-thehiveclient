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
func (c *HiveApiClient) AddLog(ctx context.Context, taskId string, reqBody export.HiveLogReq) (result *export.HiveLog, err error) {
	var (
		res *resp.Response
	)

	fName, pName := debugging.GetShortFuncNameAndPckage()

	endPoint := fmt.Sprintf("http://%s/%s/log", c.Url, cons.URITaskId.Replace(taskId))

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, reqBody); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}
	result = &export.HiveLog{}
	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
	}

	return
}

//DeleteLog удаление Log-а по его id
func (c *HiveApiClient) DeleteLog(ctx context.Context, logId string) (result bool, err error) {
	var (
		res *resp.Response
	)

	fName, pName := debugging.GetShortFuncNameAndPckage()

	endPoint := fmt.Sprintf("http://%s/%s/log", c.Url, cons.URILogId.Replace(logId))

	result = false
	if res, err = c.doApiRequest(ctx, endPoint, http.MethodDelete, nil); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}

	if res.Data != nil {
		result = true
	}

	return
}

//ListTaskLogs получения списка Log-ов связанных с Task-ом имеющим указанный taskId
func (c *HiveApiClient) ListTaskLogs(ctx context.Context, taskId string, filter interface{}, sort []apiv1.E, page *apiv1.Page) (result []export.HiveLog, err error) {
	var (
		res *resp.Response
	)

	operation := apiv1.Operation{v1const.GetTask, taskId, v1const.Logs}
	query := apiv1.CreateQueryApiV1Req(operation, filter, sort, page)

	endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URIAPIv1Query)

	fName, pName := debugging.GetShortFuncNameAndPckage()

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, query); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}
	//result := gjson.ParseBytes()

	if err = json.Unmarshal(res.Data, &result); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
	}

	return
}
