package cli

import (
  list "github.com/eumnDev/euka/cmd/euka-cli/list"
  "github.com/spf13/cobra"
)

func init() {
  cliList.AddCommand(
    cliListTopic)

  cliMain.AddCommand(cliList)
}

//CLI List Command
var cliList = &cobra.Command{
  Use: "list",
  Short: "List any kafka data (topic)",
  Long:  `List any kafka data`,
}

var cliListTopic = &cobra.Command{
  Use:   "topic",
  Short: "List All Kafka Topic",
  Long:  `List All Kafka Topic`,
  Args:  cobra.NoArgs,
  Run: func(cmd *cobra.Command, args []string) {
    list.Topic(configPath)
  },
}
