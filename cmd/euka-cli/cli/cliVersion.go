package cli

import (
  "fmt"
  "github.com/spf13/cobra"
)

func init() {
  cliMain.AddCommand(cliVersion)
}

var cliVersion = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Euka-CLI",
  Long:  `All software has versions.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Euka-CLI v0.0.1")
  },
}
