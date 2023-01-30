//type.go - содержит определение структуры HiveApiClient описывающую клиент для работы с api theHive
//и ее метод doRequest реализующий базовое взаимодействие с api
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/resp"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/logger"
)

type HiveApiClient struct {
	Client *http.Client
	Header *http.Header
	Url    string
	log    logger.Logger
}

//doApiRequest - выполняет сконфигурированный запрос к api Hive
func (c *HiveApiClient) doApiRequest(ctx context.Context, endPoint, reqMethod string, reqBody interface{}, queryParametrs ...string) (*resp.Response, error) {
	var (
		err      error
		buf      bytes.Buffer
		httpResp *http.Response
		res      *resp.Response
		req      *http.Request
	)

	if reqBody != nil {
		if err = json.NewEncoder(&buf).Encode(reqBody); err != nil {
			return nil, err
		}
	}

	if req, err = http.NewRequestWithContext(ctx, reqMethod, endPoint, &buf); err != nil {
		return nil, err
	}

	c.httpHeaderInject(req)

	if len(queryParametrs) != 0 {
		req.URL.RawQuery = strings.Join(queryParametrs, "&")
	}

	httpResp, err = c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if res, err = resp.Handler(httpResp); err != nil {
		return nil, err
	}

	return res, nil
}

//httpHeaderInject включение пользовательского http-заголовка в тело запроса
func (c *HiveApiClient) httpHeaderInject(req *http.Request) {
	if c.Header != nil {
		req.Header = *c.Header
	}
	/*else if c.Header != nil {
		req.Header = c.Header
	}*/
}
