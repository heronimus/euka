package producer

import (
	config 		"github.com/eumnDev/euka/internal/global/config/sarama"
	producer	"github.com/eumnDev/euka/internal/repository/producer/sarama"
	response	"github.com/eumnDev/euka/internal/global/responses/producer/general"
)

//NewSyncMessage ....
func NewSyncMessage(topic string, messages []string, eukaConfig config.EukaConfig) ([]response.ResponseProducer, error) {
	producerResponses, err := producer.ProduceSyncMessage(topic, messages, eukaConfig)
	return producerResponses, err
}
