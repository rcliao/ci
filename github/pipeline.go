package github

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"

	"github.com/rcliao/e2etest"
)

// Pipeline implements e2e.Pipeline to do github pipelines for testing
type Pipeline struct {
	ID        string
	tempDir   string
	statusDao e2etest.StatusDAO
}

// NewPipeline is constructor for creating new pipeline
func NewPipeline(statusDao e2etest.StatusDAO, ID string) *Pipeline {
	return &Pipeline{
		ID:        ID,
		tempDir:   randString(32),
		statusDao: statusDao,
	}
}

// Clone will clone the repository down to the temp dir
func (p *Pipeline) Clone(owner, name string) error {
	commands := []string{
		"git",
		"clone",
		fmt.Sprintf("git@github.com:%s/%s.git", owner, name),
		p.tempDir,
	}
	cmd := exec.Command(commands[0], commands[1:]...)
	output, err := cmd.Output()
	if err == nil {
		log.Println("Git clone command output", output)
		p.statusDao.Log(p.ID, string(output))
	}
	exitErr, isExitErr := err.(*exec.ExitError)
	if isExitErr {
		log.Println("has issue storing error output for Clone", string(exitErr.Stderr))
		p.statusDao.Log(p.ID, string(exitErr.Stderr))
	}
	return err
}

// Build uses `build.sh` command to build the repository
func (p *Pipeline) Build(command string) error {
	panic("not implemented")
}

// Env provides the environment variables for Start method below
func (p *Pipeline) Env() []string {
	panic("not implemented")
}

// Start uses `start.sh` to start command
func (p *Pipeline) Start(Env []string, command string, stop <-chan bool) error {
	panic("not implemented")
}

// Test runs `test.sh` to test
func (p *Pipeline) Test(Env []string) (e2etest.Result, error) {
	panic("not implemented")
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
