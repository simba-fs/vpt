package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func statusCmd(cmd *cobra.Command, args []string) error {
	fmt.Printf("Status: ......\n")

	return nil
}

func init() {
	cmd := &cobra.Command{
		Use:   `status [key]`,
		Short: "Return status of you or specific user if key is provide",
		RunE:  statusCmd,
	}

	rootCmd.AddCommand(cmd)
}
