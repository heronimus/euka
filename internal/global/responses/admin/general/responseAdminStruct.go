package admin

//ResponseAdminTopicDetail ...
type ResponseAdminTopicDetail struct {
	TopicName         string
	NumPartitions     int32
	ReplicationFactor int16
	ReplicaAssignment map[int32][]int32
	ConfigEntries     map[string]*string
}
