package services

import (
	"os/exec"
)

type gitService struct{}

func NewGitService() *gitService {
	return &gitService{}
}

func (g gitService) Changes() ([]byte, error) {
	cmd := exec.Command("git", "diff", "--cached")
	return cmd.Output()
}

func (g gitService) Commit(commit string) *exec.Cmd {
	return exec.Command("git", "commit", "-m", commit)
}
