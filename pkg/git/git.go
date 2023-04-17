package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetStagedFiles()string {
	cmd := exec.Command("git", "diff", "--staged")
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
	
	if err := cmd.Run(); err != nil {
        fmt.Printf("failed to run git command: %s", stderr.String())
    }

	return(stdout.String())
}