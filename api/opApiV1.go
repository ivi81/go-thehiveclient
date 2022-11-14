//opApiV1.go - содержит методы структуры HiveApiClient реализующие операций с api v1
package api

// ApiV1Request операция изменения Case
/*func (c *HiveApiClient) ApiV1Request(ctx context.Context, reqBody apiv1.Request, queryParametrs ...string) (apiv1.Response, error) {
	var (
		err    error
		res    *resp.Response
		result apiv1.Response
	)

	endPoint := path.Join("http://"+c.Url, cons.URIAPIv1Query)

	if res, err = c.doApiRequest(ctx, endPoint, http.MethodPost, reqBody, queryParametrs...); err != nil {
		return nil, fmt.Errorf("%s : %s", support.GetFuncName(), err.Error())
	}

	if err = json.Unmarshal(res.Data, result); err != nil {
		err = fmt.Errorf("unmarshal Error : %s : %s", support.GetFuncName(), err.Error())
	}

	return result, err
}
*/
