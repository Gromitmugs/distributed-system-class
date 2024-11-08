package main

import (
	"log"

	"github.com/Gromitmugs/distribued-system-class/tcp_echo_server_client/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd.AddCommand(
		cmd.ServerCmd,
		cmd.ClientCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}

var rootCmd = &cobra.Command{
	Short: "Simple TCP Echo Server/Client",
}
