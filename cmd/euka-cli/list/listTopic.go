package list

import (
	"log"
	"os"
	euka  "github.com/eumnDev/euka/pkg/euka"
)

//Topic ...
func Topic(configPath string) {
	err := euka.SetConfigFile(configPath)
	if err != nil {
    log.Fatal(err)
  }
	responses, _ := euka.AdminListAllTopic()

	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	log.SetOutput(os.Stdout)
	for _, topic := range responses {
		log.Printf(topic.TopicName)
	}
}
