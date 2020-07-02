package producer

import (
	"log"

	euka  "github.com/eumnDev/euka/pkg/euka"
)

//FromParam ...
func FromParam(topic string, message string, configPath string) {
	err := euka.SetConfigFile(configPath)
	if err != nil {
    log.Fatal(err)
  }
	responses, _ := euka.ProduceSyncMessage(topic, []string{message})
	response := responses[0]
	log.Printf("Message Content : %s ...\n", message)
	log.Printf("Message Stored  : topic(%s)/partition(%d)/offset(%d)\n\n", response.Topic, response.Partition, response.Offset)
}
