package producer

import (
	"bufio"
	"log"
	"os"

	euka  "github.com/eumnDev/euka/pkg/euka"
)

//FromFile ...
func FromFile(topic string, filePath string, configPath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var messages []string
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}

	err = euka.SetConfigFile(configPath)
	if err != nil {
    log.Fatal(err)
  }
	responses, _ := euka.ProduceSyncMessage(topic, messages)
	for i, response := range responses {
		log.Printf("Message Content : %s ...\n", messages[i][:35])
		log.Printf("Message Stored  : topic(%s)/partition(%d)/offset(%d)\n\n", response.Topic, response.Partition, response.Offset)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
