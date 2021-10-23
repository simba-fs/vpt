package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type tunnel struct {
	LocalPort  int
	ServerIP   string
	ServerPort int
}

func connectCmd(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("Args must be two")
	}

	// first arg
	mode := args[0]
	if mode != "client" && mode != "host" {
		return errors.New(`First arg must be "client" or "host"`)
	}

	// second arg
	part := strings.SplitN(args[1], ":", 3)

	// local port
	localPort, err := strconv.ParseInt(part[0], 10, 64)
	if err != nil {
		return err
	}

	serverIP := part[1]

	// server port
	serverPort, err := strconv.ParseInt(part[2], 10, 64)
	if err != nil {
		return err
	}

	fmt.Printf("mode: %v\n", mode)
	fmt.Printf("localPort: %v\n", localPort)
	fmt.Printf("serverIP: %v\n", serverIP)
	fmt.Printf("serverPort: %v\n", serverPort)

	return nil
}

func init() {
	cmd := &cobra.Command{
		Use:   `connect <mode> <localPort:serverIP:serverPort>`,
		Long:  `<mode> must be "host" or "client"`,
		Short: "Establish a ssh tunnel to configured server",
		RunE:  connectCmd,
	}

	rootCmd.AddCommand(cmd)
}
