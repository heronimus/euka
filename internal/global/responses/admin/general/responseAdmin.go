package admin

//NewResponseAdminTopicDetail ....
func NewResponseAdminTopicDetail() *ResponseAdminTopicDetail {
	newResponse := &ResponseAdminTopicDetail{}

	newResponse.TopicName = "__euka_topic_invalid"
	newResponse.NumPartitions = -999
	newResponse.ReplicationFactor = -999

	return newResponse
}
