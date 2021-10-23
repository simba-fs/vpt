package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func keyCmd(cmd *cobra.Command, args []string) error {
	fmt.Printf("Return ssh public key\n")

	return nil
}

func keyRenewCmd(cmd *cobra.Command, args []string) error {
	fmt.Printf("Renew ssh key\n")

	return nil
}

func keyAddCmd(cmd *cobra.Command, args []string) error {
	fmt.Printf("Valid and add ssh public key\n")

	return nil
}

func init() {
	cmd := &cobra.Command{
		Use:   `key`,
		Short: "Return ssh public key",
		RunE:  keyCmd,
	}

	cmd.AddCommand(&cobra.Command{
		Use: `renew`,
		Short: `Renew your ssh key(this will generate a new ID)(for client)`,
		RunE: keyRenewCmd,
	}, &cobra.Command{
		Use: `add <key>`,
		Short: `Add a user's ssh public key(for server)`,
		RunE: keyAddCmd,
	})

	rootCmd.AddCommand(cmd)
}
