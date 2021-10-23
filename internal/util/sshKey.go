package util

import (
	"os"
	"os/exec"
	"path"
)

// SSHKey return public ssh key
func SSHKey() (string, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	configPath, err = EnsureDir(path.Join(configPath, "stl"))
	if err != nil {
		return "", err
	}

	sshKeyPath := path.Join(configPath, "stlKey")
	// find ssh key
	if _, err := os.Stat(sshKeyPath + ".pub"); os.IsNotExist(err) {
		// generate a new one
		if _, err := exec.Command("ssh-keygen", "-f", sshKeyPath).CombinedOutput(); err != nil {
			return "", err
		}
	}

	// read and print
	key, err := os.ReadFile(sshKeyPath + ".pub")
	if err != nil {
		return "", err
	}
	return string(key), nil

}

// SSHKeyPath return path to ssh public key
func SSHKeyPath() (string, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	configPath, err = EnsureDir(path.Join(configPath, "stl"))
	if err != nil {
		return "", err
	}

	sshKeyPath := path.Join(configPath, "stlKey.pub")
	return sshKeyPath, nil
}

