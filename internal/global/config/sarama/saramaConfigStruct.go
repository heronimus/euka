package config

import (
	"github.com/Shopify/sarama"
)

//EukaConfig ...
type EukaConfig struct {
	BrokerList       []string
	ProviderIdentity string
	SaramaConfig     *sarama.Config
}

//configFile ...
type configFile struct {
	BrokerList       []string `yaml:"brokerList" env:"BROKER_LIST" env-default:"localhost:9092"`
	ProviderIdentity string   `yaml:"providerIdentity" env:"PROVIDER_IDENTITY" env-default:"sarama"`
}
