package euka

import (
	producer "github.com/eumnDev/euka/pkg/command/sarama/producer"
	response	"github.com/eumnDev/euka/internal/global/responses/producer/general"
)

//ProduceSyncMessage ..
func ProduceSyncMessage(topic string, messages []string) ([]response.ResponseProducer, error) {
  return producer.NewSyncMessage(topic, messages, EukaGlobalConfig)
}
