package api

import (
	"net/http"

	"strconv"
	"strings"

	"gitlab.cloud.gcm/i.ippolitov/debugging"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/config"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/logger"
)

//NewApiClient создает и настраивает клиент theHive Rupor-a
func NewApiClient(cfg *config.HiveClientCfg, logg logger.Logger) (*HiveApiClient, error) {

	//url := "http://" + cfg.Host
	url := cfg.Host

	if strings.HasSuffix(url, "/") {
		url = strings.TrimSuffix(url, "/")
	}

	if cfg.Port != 0 {
		url = url + ":" + strconv.Itoa(cfg.Port)
	}

	/*tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}*/

	c := &HiveApiClient{
		Url:    url,
		Header: createHeader(cfg),
		Client: &http.Client{}, //Transport: tr},
		log:    logg,
	}
	pName, fName := debugging.GetShortFuncNameAndPckage()
	c.log.Debugf("package: %s, function: %s, NEW CLIENT INITIALIZATION COMPLITE", pName, fName)
	return c, nil
}

//createHeader - создает заголовок HTTP-запроса
func createHeader(cfg *config.HiveClientCfg) *http.Header {
	header := http.Header{}
	if cfg.BearerToken != "" {
		header.Add("Authorization", "Bearer "+cfg.BearerToken)
	}

	if cfg.ApiKey != "" {
		header.Add("x-token", cfg.ApiKey)
	}

	if cfg.UserAgent != "" {
		header.Add("User-Agent", cfg.UserAgent)
	}

	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json")
	header.Add("Connection", "keep-alive")
	return &header
}
