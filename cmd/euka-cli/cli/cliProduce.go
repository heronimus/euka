package cli

import (
  produce "github.com/eumnDev/euka/cmd/euka-cli/produce"
  "github.com/spf13/cobra"
)

func init() {
  cliProduce.AddCommand(
		cliProduceCLI,
		cliProduceFile)

  cliMain.AddCommand(cliProduce)
}

var cliProduce = &cobra.Command{
  Use: "produce",
  Short: "Send anything to Kafka Topic",
  Long:  `Send any message to specific Kafka Topic`,
}

var cliProduceCLI = &cobra.Command{
  Use:   "from-cli [topic] [messages]",
  Short: "Send anything to Kafka topic",
  Long:  `Send any simple message to specific Kafka Topic`,
  Args:  cobra.ExactArgs(2),
  Run: func(cmd *cobra.Command, args []string) {
      produce.FromParam(args[0], args[1], configPath)
  },
}

var cliProduceFile = &cobra.Command{
  Use:   "from-file [topic] [file-path]",
  Short: "Send anything to Kafka Topic",
  Long:  `Send message from file to specific Kafka Topic`,
  Args:  cobra.ExactArgs(2),
  Run: func(cmd *cobra.Command, args []string) {
      produce.FromFile(args[0], args[1], configPath)
  },
}
