package exec

import (
	"os/exec"
)

// Run runs a shell command and return stdout/stderr
func Run(cmd ...string) []byte {
	exec.Command([]string{"a"}...)

	return make([]byte, 0)
}
