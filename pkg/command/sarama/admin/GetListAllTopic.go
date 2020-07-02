package admin

import (
	admin			"github.com/eumnDev/euka/internal/repository/admin/sarama"
	config 		"github.com/eumnDev/euka/internal/global/config/sarama"
	response 	"github.com/eumnDev/euka/internal/global/responses/admin/general"
)

//GetListAllTopic ....
func GetListAllTopic(eukaConfig config.EukaConfig) ([]response.ResponseAdminTopicDetail, error) {
	adminResponses, err := admin.ListAllTopic(eukaConfig)
	return adminResponses, err
}
