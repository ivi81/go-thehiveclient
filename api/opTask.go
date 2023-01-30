//opTask.go - содержит методы структуры HiveApiClient реализующие операций с сущьностями Task
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

//CreateTask создает Task в заданном Case
func (c *HiveApiClient) CreateTask(ctx context.Context, caseId string, reqBody *export.HiveTaskReq) (result *export.HiveTask, err error) {
	var (
		res *resp.Response
	)

	endPoint := fmt.Sprintf("http://%s/%s/task", c.Url, cons.URICaseId.Replace(caseId))

	fName, pName := debugging.GetShortFuncNameAndPckage()

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, reqBody); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}

	if err = json.Unmarshal(res.Data, &result); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
	}

	return
}

//UpdateTask изменение значений полей Task-а
func (c *HiveApiClient) UpdateTask(ctx context.Context, taskId string, reqBody *export.HiveTaskReq) (result *export.HiveTask, err error) {

	var res *resp.Response

	endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URITaskId.Replace(taskId))

	fName, pName := debugging.GetShortFuncNameAndPckage()

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPatch, reqBody); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
	}

	return
}

//GetTask получение полей Task-а
func (c *HiveApiClient) GetTask(ctx context.Context, taskId string) (result *export.HiveTask, err error) {
	var (
		res *resp.Response
	)

	endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URITaskId.Replace(taskId))

	fName, pName := debugging.GetShortFuncNameAndPckage()

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodGet, nil); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
	}

	return
}

//ListCaseTask получение списка Task-ов связанных с указанным Case-ом
func (c *HiveApiClient) ListCaseTask(ctx context.Context, caseId string, filter interface{}, sort []apiv1.E, page *apiv1.Page) (result []export.HiveTask, err error) {

	var (
		res *resp.Response
	)

	operation := apiv1.Operation{v1const.GetCase, caseId, v1const.Tasks}
	query := apiv1.CreateQueryApiV1Req(operation, filter, sort, page)

	endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URIAPIv1Query)

	fName, pName := debugging.GetShortFuncNameAndPckage()

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, query); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}

	if err = json.Unmarshal(res.Data, &result); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
	}

	return
}
