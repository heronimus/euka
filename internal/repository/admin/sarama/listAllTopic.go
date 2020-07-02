package admin

import (
	"github.com/Shopify/sarama"
	"sort"
	config 		"github.com/eumnDev/euka/internal/global/config/sarama"
	response	"github.com/eumnDev/euka/internal/global/responses/admin/general"
)

//ListAllTopic ....
func ListAllTopic(eukaConfig config.EukaConfig) ([]response.ResponseAdminTopicDetail, error) {

	clusterAdmin, err := sarama.NewClusterAdmin(eukaConfig.BrokerList, eukaConfig.SaramaConfig)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := clusterAdmin.Close(); err != nil {
			panic(err)
		}
	}()

	listTopicDetails, _ := clusterAdmin.ListTopics()

	var sortedTopicKeys []string
	for topicKey := range listTopicDetails {
		sortedTopicKeys = append(sortedTopicKeys, topicKey)
	}
	sort.Strings(sortedTopicKeys)

	listTopicDetailResponses := make([]response.ResponseAdminTopicDetail, len(sortedTopicKeys))
	for i, topicName := range sortedTopicKeys {

		topicDetail := response.NewResponseAdminTopicDetail()

		topicDetail.TopicName = topicName
		topicDetail.NumPartitions = listTopicDetails[topicName].NumPartitions
		topicDetail.ReplicationFactor = listTopicDetails[topicName].ReplicationFactor
		topicDetail.ReplicaAssignment = listTopicDetails[topicName].ReplicaAssignment
		topicDetail.ConfigEntries = listTopicDetails[topicName].ConfigEntries

		listTopicDetailResponses[i] = *topicDetail
	}

	return listTopicDetailResponses, err
}
