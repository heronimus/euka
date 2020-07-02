package cli

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

var configPath string
var bootstrapServer string

//CLI Main entrance
var cliMain = &cobra.Command{Use: "euka-cli"}

//Execute main cli
func Execute() {
  if err := cliMain.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
	cliMain.PersistentFlags().StringVarP(&configPath, "config", "c", "config.yaml", "config file")
	cliMain.PersistentFlags().StringVarP(&bootstrapServer, "bootstrap-server", "b", "", "Simple provide bootstrap-server without config file.")
}
