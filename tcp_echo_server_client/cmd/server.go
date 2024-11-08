package cmd

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the TCP echo server",
	RunE: func(cmd *cobra.Command, args []string) error {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
		if err != nil {
			return err
		}
		fmt.Println("TCP listening port", port)

		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}

			go echoMessage(conn)
		}
	},
}

func echoMessage(conn net.Conn) {
	defer conn.Close()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		fmt.Printf("message received from client at time %s: %s",
			time.Now().Format("15:04:05"), message)

		if _, err = conn.Write([]byte(message + "\n")); err != nil {
			return
		}
	}
}
