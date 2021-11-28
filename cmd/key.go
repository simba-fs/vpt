package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"

	"github.com/simba-fs/vpt/internal/util"
)

const debug = false

func keyCmd(cmd *cobra.Command, args []string) error {
	key, err := util.SSHKey()
	if err != nil {
		return err
	}
	fmt.Printf("%s", key)
	return nil
}

func keyRenewCmd(cmd *cobra.Command, args []string) error {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	configPath, err = util.EnsureDir(path.Join(configPath, "stl"))
	if err != nil {
		return err
	}

	sshKeyPath := path.Join(configPath, "stlKey")

	// remove old ones
	if err := os.Remove(sshKeyPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	if err := os.Remove(sshKeyPath + ".pub"); err != nil && !os.IsNotExist(err) {
		return err
	}

	// generate new
	if _, err := exec.Command("ssh-keygen", "-f", sshKeyPath).CombinedOutput(); err != nil {
		return err
	}

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
	})

	rootCmd.AddCommand(cmd)
}
