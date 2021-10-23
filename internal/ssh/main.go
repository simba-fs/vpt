package ssh

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/meow55555/stl/internal/util"
)

type Tunnel struct {
	mode       string // client or host
	localPort  string
	serverIP   string
	serverPort string // ssh port, default: 22
	startTime  time.Time
}

// New establish a ssh tunnel
func New(mode string, localPort string, serverIP string, serverPort string) *Tunnel {
	tunnel := &Tunnel{
		mode:       mode,
		localPort:  localPort,
		serverIP:   serverIP,
		serverPort: serverPort,
		startTime:  time.Now(),
	}

	return tunnel
}

// Connect will exec ssh command depand on mode. 
func (t *Tunnel) Connect() error {
	keyPath, err := util.SSHKeyPath()
	if err != nil {
		return err
	}

	cmd := &exec.Cmd{}
	if t.mode == "host" {
		cmd = exec.Command(
			"ssh",
			"-i", keyPath,
			"-NR", fmt.Sprintf("/tmp/stl:localhost:%s", t.localPort),
			t.serverIP,
			"-p", t.serverPort,
		)


	} else {
		cmd = exec.Command(
			"ssh",
			"-i", keyPath,
			"-NL", fmt.Sprintf("%s:/tmp/stl", t.localPort),
			t.serverIP,
			"-p", t.serverPort,
		)
	}

	a, _ := cmd.CombinedOutput()
	fmt.Println(string(a))
	return nil
}
