package cmd

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var ClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start the TCP echo client",
	RunE: func(cmd *cobra.Command, args []string) error {
		conn, err := net.Dial("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			return nil
		}
		defer conn.Close()

		if len(args) == 0 {
			return fmt.Errorf("need at least an argument for client message")
		}

		for _, message := range args {
			sendMessage(conn, message)
			time.Sleep(time.Second)
		}
		return nil
	},
}

func sendMessage(conn net.Conn, message string) error {
	if _, err := conn.Write([]byte(message + "\n")); err != nil {
		return err
	}

	resp, err := bufio.NewReader(conn).ReadString(byte('\n'))
	if err != nil {
		return err
	}
	fmt.Printf("message received from server at time %s: %s",
		time.Now().Format("15:04:05"), resp)

	return nil
}
