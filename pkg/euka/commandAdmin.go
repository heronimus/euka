package euka

import (
	admin 		"github.com/eumnDev/euka/pkg/command/sarama/admin"
	response 	"github.com/eumnDev/euka/internal/global/responses/admin/general"
)

//AdminListAllTopic ..
func AdminListAllTopic() ([]response.ResponseAdminTopicDetail, error) {
  return admin.GetListAllTopic(EukaGlobalConfig)
}
