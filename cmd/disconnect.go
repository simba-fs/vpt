package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func disconnectCmd(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("Args must be one")
	}

	port, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}
	fmt.Printf("port: %v\n", port)

	return nil
}

func init() {
	cmd := &cobra.Command{
		Use:   "disconnect <port>",
		Short: "disconnect ssh tunnel on <port>",
		RunE:  disconnectCmd,
	}

	rootCmd.AddCommand(cmd)
}
