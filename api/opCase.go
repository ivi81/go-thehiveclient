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
func (c *HiveApiClient) CreateCase(ctx context.Context, reqBody export.HiveCaseReq) (*export.HiveCase, error) {

	var (
		err error
		//req    *http.Request
		res    *resp.Response
		result export.HiveCase
	)

	endPoint := fmt.Sprintf("http://%s/%s/", c.Url, cons.URICase)
	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, reqBody); err != nil {
		return nil, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}
	/*	var buf bytes.Buffer
		if err = json.NewEncoder(&buf).Encode(reqBody); err != nil {
			return nil, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
		}

		if req, err = http.NewRequestWithContext(ctx, http.MethodPost, endPoint, &buf); err != nil {
			return nil, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
		}

		response, err := c.Client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
		}

		if res, err = resp.Handler(response); err != nil {
			return nil, err
		}
	*/
	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("unmarshal Error : %s : %s", debugging.GetFuncName(), err.Error())
	}

	return &result, err
}

//UpdateCase операция изменения Case
func (c *HiveApiClient) UpdateCase(ctx context.Context, caseId string, reqBody export.HiveCaseReq) (*export.HiveCase, error) {
	var (
		err error
		//	req    *http.Request
		res    *resp.Response
		result export.HiveCase
	)

	endPoint := fmt.Sprintf("http://%s/%s/", c.Url, cons.URICaseId.Replace(caseId))
	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPatch, reqBody); err != nil {
		return nil, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("unmarshal Error : %s : %s", debugging.GetFuncName(), err.Error())
	}

	return &result, err
}

//DeleteCase операция удаления Case
func (c *HiveApiClient) DeleteCase(ctx context.Context, caseId string) (bool, error) {
	var (
		err error
		//req *http.Request
		res *resp.Response
	)

	endPoint := fmt.Sprintf("http://%s/%s/?force=1", c.Url, cons.URICaseId.Replace(caseId))
	if res, err = c.doApiRequest(ctx, endPoint, http.MethodDelete, nil); err != nil {
		return false, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}
	if res.Data == nil {
		return false, nil
	}

	return true, err
}

//MergeCase операция объединения двух Case-ов в одни
func (c *HiveApiClient) MergeCase(ctx context.Context, caseId1, caseId2 string) (*export.HiveCase, error) {

	var (
		err error
		//	req    *http.Request
		res    *resp.Response
		result export.HiveCase
	)

	endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URIMergeCase.Replace(caseId1, caseId2))

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, nil); err != nil {
		return nil, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("unmarshal Error : %s : %s", debugging.GetFuncName(), err.Error())
	}

	return &result, err
}

//ExportCaseToMISP операция экспортирования Case-а в MISP
func (c *HiveApiClient) ExportCaseToMISP(ctx context.Context, caseId string, mispServer string) (bool, error) {
	var (
		err error
		res *resp.Response
		//req *http.Request
	)

	endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URIMispExport.Replace(caseId, mispServer))

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, nil); err != nil {
		return false, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}
	if res.Data == nil {
		return false, nil
	}

	return true, err
}

//ListRelatedCases операция получения Case-ов связанных с текущим
func (c *HiveApiClient) ListRelatedCases(ctx context.Context, caseId string) ([]export.HiveCase, error) {
	var (
		err error
		//req    *http.Request
		res    *resp.Response
		result []export.HiveCase
	)

	endPoint := path.Join("http://"+c.Url, cons.URICaseId.Replace(caseId), "links")

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodGet, nil); err != nil {
		return nil, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("unmarshal Error : %s : %s", debugging.GetFuncName(), err.Error())
	}

	return result, err
}

//ListRelatedAlerts
func (c *HiveApiClient) ListRelatedAlerts(reqParama apiv0.RequestApiV0Body) ([]export.HiveAlert, error) {
	return nil, nil
}

//List case tasks
// /api/v0/query
