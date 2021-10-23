package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

func ensureConfigDir() (string, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	configPath = path.Join(configPath, "stl")
	// sshPrivatePath := path.Join(configPath, "ssh.private")

	stat, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		// mkdir
		os.Mkdir(configPath, 0775)
	}else if !stat.IsDir() {
		// remove
		if err := os.Remove(configPath); err != nil {
			return "", err
		}
		// mkdir
		os.Mkdir(configPath, 0775)
	}
	return configPath, nil
}

// func ensureDotSshDir() (string, error){
//     sshPath, err := os.UserHomeDir()
//
// }
//
func keyCmd(cmd *cobra.Command, args []string) error {
	configPath, err := ensureConfigDir()
	if err != nil {
		return err
	}

	sshKeyPath := path.Join(configPath, "stlKey")
	// find ssh key
	if _, err := os.Stat(sshKeyPath + ".pub"); os.IsNotExist(err) {
		// generate a new one
		if _, err := exec.Command("ssh-keygen", "-f", sshKeyPath).CombinedOutput(); err != nil {
			return err
		}
	}

	// read and print
	key, err := os.ReadFile(sshKeyPath + ".pub")
	if err != nil {
		return err
	}
	fmt.Printf("%s", key)
	return nil
}

func keyRenewCmd(cmd *cobra.Command, args []string) error {
	configPath, err := ensureConfigDir()
	if err != nil {
		return err
	}

	sshKeyPath := path.Join(configPath, "stlKey")

	// remove old ones
	if err := os.Remove(sshKeyPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.Remove(sshKeyPath+".pub"); err != nil && !os.IsNotExist(err) {
		return err
	}

	// generate new
	if _, err := exec.Command("ssh-keygen", "-f", sshKeyPath).CombinedOutput(); err != nil {
		return err
	}

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
		Use:   `renew`,
		Short: `Renew your ssh key(this will generate a new ID)(for client)`,
		RunE:  keyRenewCmd,
	}, &cobra.Command{
		Use:   `add <key>`,
		Short: `Add a user's ssh public key(for server)`,
		RunE:  keyAddCmd,
	})

	rootCmd.AddCommand(cmd)
}
