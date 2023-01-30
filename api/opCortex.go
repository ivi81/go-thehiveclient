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
func (c *HiveApiClient) RunResponder(ctx context.Context, reqBody export.CortexResponderReq) (result *export.CortexResponderResult, err error) {
	var (
		res *resp.Response
	)

	fName, pName := debugging.GetShortFuncNameAndPckage()

	endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URICortexAction)

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, reqBody); err != nil {
		err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		return
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("function: %s , %s", fName, err.Error())
	}

	return
}

//ListResonderAction список респондеров отработавших для сущьности с указанным entityId
func (c *HiveApiClient) ListResonderAction(ctx context.Context, hiveEntity common.ObjectType, entityId string) (result []export.CortexResponderResult, err error) {
	var (
		res *resp.Response
	)

	fName, pName := debugging.GetShortFuncNameAndPckage()

	switch hiveEntity {
	case common.HIVECASE, common.HIVEALERT, common.HIVETASK, common.HIVETASKLOG:
		endPoint := fmt.Sprintf("http://%s/%s", c.Url, cons.URICortexActionEntityId.Replace(hiveEntity.String(), entityId))

		if res, err = c.doApiRequest(ctx, endPoint, http.MethodGet, nil); err != nil {
			err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
			return
		}

		if err = json.Unmarshal(res.Data, result); err != nil {
			err = fmt.Errorf("package: %s, function: %s, %s", fName, pName, err.Error())
		}

		return
	}

	err = fmt.Errorf("package: %s, function: %s, not valid hive type object: %s", fName, pName, hiveEntity)
	return
	// /api/connector/cortex/action/case_task_log/{id} id: Log identifier
	// /api/connector/cortex/action/case/{id} id: Case identifier
	// /api/connector/cortex/action/responder/alert/{id} id: Alert identifier

	// /api/connector/cortex/responder/alert/{id}
	// /api/connector/cortex/responder/case/{id}

}
