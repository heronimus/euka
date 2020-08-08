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
  messageLine := 0
	for scanner.Scan() {
		messageLine = messageLine + 1
		if len(scanner.Text()) > 0 {
			messages = append(messages, scanner.Text())
		} else {
			log.Printf("WARNING: empty line at line (%d) will be skipped.\n\n", messageLine)
		}
	}

	err = euka.SetConfigFile(configPath)
	if err != nil {
    log.Fatal(err)
  }
	responses, _ := euka.ProduceSyncMessage(topic, messages)
	for i, response := range responses {
		log.Printf("Message Content : %.80s ...\n", messages[i])
		log.Printf("Message Stored  : topic(%s)/partition(%d)/offset(%d)\n\n", response.Topic, response.Partition, response.Offset)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
