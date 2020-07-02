package producer

import (
	"github.com/Shopify/sarama"
	"log"
	config 		"github.com/eumnDev/euka/internal/global/config/sarama"
	response 	"github.com/eumnDev/euka/internal/global/responses/producer/general"
)

//ProduceSyncMessage ....
func ProduceSyncMessage(topic string, messages []string, eukaConfig config.EukaConfig) ([]response.ResponseProducer, error) {
	producer, err := sarama.NewSyncProducer(eukaConfig.BrokerList, eukaConfig.SaramaConfig)

	defer func() {
		if err := producer.Close(); err != nil {
			log.Printf("repository/sarama/producer")
			panic(err)
		}
	}()

	producerResponses := make([]response.ResponseProducer, len(messages))
	for i, message := range messages {
		producerMessage := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(message),
		}

		producerResponse := response.NewResponseProducer()
		producerResponse.Topic = topic
		producerResponse.Partition, producerResponse.Offset, err = producer.SendMessage(producerMessage)

		producerResponses[i] = *producerResponse
	}

	return producerResponses, err

}
