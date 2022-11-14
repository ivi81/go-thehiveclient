//opLog.go - содержит методы структуры HiveApiClient реализующие операций функциями cortex (респондеры и ...)
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.cloud.gcm/i.ippolitov/debugging"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/common"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/export"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/resp"
)

//RunResponder запуск респондера применительно к сущьности hive
func (c *HiveApiClient) RunResponder(ctx context.Context, reqBody export.CortexResponderReq) (*export.CortexResponderResult, error) {
	var (
		err error
		//req    *http.Request
		res    *resp.Response
		result export.CortexResponderResult
	)

	endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URICortexAction)

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, reqBody); err != nil {
		return nil, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("unmarshal Error : %s : %s", debugging.GetFuncName(), err.Error())
	}

	return &result, err
}

//ListResonderAction список респондеров отработавших для сущьности с указанным entityId
func (c *HiveApiClient) ListResonderAction(ctx context.Context, hiveEntity common.ObjectType, entityId string) ([]export.CortexResponderResult, error) {
	var (
		err error
		//req    *http.Request
		res    *resp.Response
		result []export.CortexResponderResult
	)

	switch hiveEntity {
	case common.HIVECASE, common.HIVEALERT, common.HIVETASK, common.HIVETASKLOG:
		endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URICortexActionEntityId.Replace(hiveEntity.String(), entityId))

		if res, err = c.doApiRequest(ctx, endPoint, http.MethodGet, nil); err != nil {
			return nil, fmt.Errorf("%s : %s", debugging.GetFuncName(), err.Error())
		}

		if err = json.Unmarshal(res.Data, result); err != nil {
			err = fmt.Errorf("unmarshal Error : %s : %s", debugging.GetFuncName(), err.Error())
		}

		return result, err
	}

	return nil, fmt.Errorf("Hive type object %s is not valid for %s", hiveEntity, debugging.GetFuncName())
	// /api/connector/cortex/action/case_task_log/{id} id: Log identifier
	// /api/connector/cortex/action/case/{id} id: Case identifier
	// /api/connector/cortex/action/responder/alert/{id} id: Alert identifier

	// /api/connector/cortex/responder/alert/{id}
	// /api/connector/cortex/responder/case/{id}

}
