package config

import (
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"path/filepath"
)

//NewConfig ....
func NewConfig(brokerList []string, saramaConfig *sarama.Config) (EukaConfig, error) {

	newConfig := EukaConfig{}
	newConfig.BrokerList = brokerList
	newConfig.ProviderIdentity = "sarama"

	newConfig.SaramaConfig = saramaConfig

	return newConfig, nil
}

//NewSimpleConfig ....
func NewSimpleConfig(brokerList []string) (EukaConfig, error) {

	newConfig := EukaConfig{}
	newConfig.BrokerList = brokerList
	newConfig.ProviderIdentity = "sarama"

	newConfig.SaramaConfig = sarama.NewConfig()
	newConfig.SaramaConfig.Version = sarama.V1_0_0_0
	newConfig.SaramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	newConfig.SaramaConfig.Producer.Retry.Max = 5
	newConfig.SaramaConfig.Producer.Return.Successes = true
	newConfig.SaramaConfig.Consumer.Return.Errors = true

	return newConfig, nil
}

//NewConfigFromFile ....
func NewConfigFromFile(configPath string) (EukaConfig, error) {

	newConfig := EukaConfig{}

	// Read configuration from .yaml file using viper
	viper.SetConfigName(filepath.Base(configPath))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Dir(configPath))
	
	err := viper.ReadInConfig()
	if err != nil {
		return newConfig, err
	}

	// Update EukaConfig value from configuration file
	newConfig.BrokerList = viper.GetStringSlice("brokerList")
	newConfig.ProviderIdentity = viper.GetString("providerIdentity")

	newConfig.SaramaConfig = sarama.NewConfig()
	newConfig.SaramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	newConfig.SaramaConfig.Producer.Retry.Max = 5
	newConfig.SaramaConfig.Producer.Return.Successes = true
	newConfig.SaramaConfig.Consumer.Return.Errors = true

	//Define Kafka Version Compability
	if viper.GetString("kafkaVersion") == "1.0.0.0" {
		newConfig.SaramaConfig.Version = sarama.V1_0_0_0
	}


	return newConfig, err
}
