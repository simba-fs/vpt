package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/meow55555/stl/internal/ssh"
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
	_, err := strconv.ParseInt(part[0], 10, 64)
	if err != nil {
		return err
	}


	// server port
	_, err = strconv.ParseInt(part[2], 10, 64)
	if err != nil {
		return err
	}

	fmt.Printf("mode: %v\n", mode)
	fmt.Printf("localPort: %v\n", part[0])
	fmt.Printf("serverIP: %v\n", part[1])
	fmt.Printf("serverPort: %v\n", part[2])

	return ssh.New(mode, part[0], part[1], part[2]).Connect()
}

func init() {
	cmd := &cobra.Command{
		Use:   `connect <mode> <localPort:serverIP:serverPort>`,
		Long:  `<mode> must be "host" or "client"`,
		Short: "Establish a ssh tunnel to configured server",
		Example: `host:   stl connect host 3000:example.com:22
client: stl connect client 8888:example.com:22`,
		RunE:  connectCmd,
	}

	rootCmd.AddCommand(cmd)
}
