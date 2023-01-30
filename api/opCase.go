//opCase.go - содержит методы структуры HiveApiClient реализующие операций с сущьностями Case
package api

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"

	"path"

	"gitlab.cloud.gcm/i.ippolitov/debugging"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/export"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/export/apiv0"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/resp"
)

//CreateCase операция создания Case
func (c *HiveApiClient) CreateCase(ctx context.Context, reqBody export.HiveCaseReq) (result *export.HiveCase, err error) {

	var (
		res *resp.Response
	)

	fName, pName := debugging.GetShortFuncNameAndPckage()

	endPoint := fmt.Sprintf("http://%s/%s/", c.Url, cons.URICase)

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, reqBody); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
	}

	return
}

//UpdateCase операция изменения Case
func (c *HiveApiClient) UpdateCase(ctx context.Context, caseId string, reqBody export.HiveCaseReq) (result *export.HiveCase, err error) {
	var (
		res *resp.Response
	)

	fName, pName := debugging.GetShortFuncNameAndPckage()

	endPoint := fmt.Sprintf("http://%s/%s/", c.Url, cons.URICaseId.Replace(caseId))
	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPatch, reqBody); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("function: %s , %s", fName, err.Error())
	}

	return
}

//DeleteCase операция удаления Case
func (c *HiveApiClient) DeleteCase(ctx context.Context, caseId string) (result bool, err error) {
	var (
		res *resp.Response
	)

	fName, pName := debugging.GetShortFuncNameAndPckage()

	endPoint := fmt.Sprintf("http://%s/%s/?force=1", c.Url, cons.URICaseId.Replace(caseId))

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

//MergeCase операция объединения двух Case-ов в одни
func (c *HiveApiClient) MergeCase(ctx context.Context, caseId1, caseId2 string) (result *export.HiveCase, err error) {

	var (
		res *resp.Response
	)

	fName, pName := debugging.GetShortFuncNameAndPckage()

	endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URIMergeCase.Replace(caseId1, caseId2))

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, nil); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("function: %s , %s", fName, err.Error())
	}

	return
}

//ExportCaseToMISP операция экспортирования Case-а в MISP
func (c *HiveApiClient) ExportCaseToMISP(ctx context.Context, caseId string, mispServer string) (result bool, err error) {
	var (
		res *resp.Response
		//req *http.Request
	)
	fName, pName := debugging.GetShortFuncNameAndPckage()

	endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URIMispExport.Replace(caseId, mispServer))
	result = false

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, nil); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}
	if res.Data != nil {
		result = true
	}

	return
}

//ListRelatedCases операция получения Case-ов связанных с текущим
func (c *HiveApiClient) ListRelatedCases(ctx context.Context, caseId string) (result []export.HiveCase, err error) {
	var (
		res *resp.Response
	)
	fName, pName := debugging.GetShortFuncNameAndPckage()

	endPoint := path.Join("http://"+c.Url, cons.URICaseId.Replace(caseId), "links")

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodGet, nil); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("function: %s , %s", fName, err.Error())
	}

	return
}

//ListRelatedAlerts
func (c *HiveApiClient) ListRelatedAlerts(reqParama apiv0.RequestApiV0Body) ([]export.HiveAlert, error) {
	return nil, nil
}

//List case tasks
// /api/v0/query
