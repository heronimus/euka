package euka

import (
	eukaSaramaConfig "github.com/eumnDev/euka/internal/global/config/sarama"
)

//EukaGlobalConfig ...
var EukaGlobalConfig eukaSaramaConfig.EukaConfig

//SetConfigSimple ..
func SetConfigSimple(brokerList []string) error {
	var err error
	EukaGlobalConfig, err = eukaSaramaConfig.NewSimpleConfig(brokerList)

	return err
}

//SetConfigFile ..
func SetConfigFile(configPath string) error {
	var err error
	EukaGlobalConfig, err = eukaSaramaConfig.NewConfigFromFile(configPath)

	return err
}
