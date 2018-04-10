package repo

import (
	"os/exec"
)

type GitCommand struct {
	path string
}

func buildGitCommand(args ...string) *exec.Cmd {
	return exec.Command("git", args...)
}

func buildGitCloneCommand(origin, path string) *exec.Cmd {
	args := []string{"clone", origin, path}
	return buildGitCommand(args...)
}

func buildGitPullCommand(path string) *exec.Cmd {
	args := []string{"-C", path, "pull"}
	return buildGitCommand(args...)
}

func executeCommand(cmd *exec.Cmd) {
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()
	go scanOutput(stdout)
	go scanOutput(stderr)
	cmd.Wait()

}
func (c *GitCommand) Clone(origin string) {
	cmd := buildGitCloneCommand(origin, c.path)
	executeCommand(cmd)
}

func (c *GitCommand) Pull() {
	cmd := buildGitPullCommand(c.path)
	executeCommand(cmd)
}

func NewGitCommand(p string) *GitCommand {
	return &GitCommand{
		path: p,
	}
}
