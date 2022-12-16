package config

import "gitlab.cloud.gcm/i.ippolitov/go-microconfig/microconfig"

//HiveClientCfg - тип данных хранящий информацию о конфигурации клиента theHive
type HiveClientCfg struct {
	microconfig.ClientHttpCfg `yaml:",inline"`
}

func (cfg *HiveClientCfg) SetValuesFromEnv(envPrefix string) {
	cfg.ClientHttpCfg.SetValuesFromEnv(envPrefix)
}
