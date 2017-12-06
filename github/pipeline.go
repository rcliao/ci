package github

import "github.com/rcliao/e2etest"

// Pipeline implements e2e.Pipeline to do github pipelines for testing
type Pipeline struct {
	tempDir string
}

// Clone will clone the repository down to the temp dir
func (p *Pipeline) Clone(URL string) error {
	panic("not implemented")
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
