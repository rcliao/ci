package e2etest

type Result struct {
	Pass bool
}

type Status struct {
	ID          string
	State       string
	TargerURL   string
	Description string
	Context     string
}

type Pipeline interface {
	Clone(URL string) error
	Build(command string) error
	Env(defaultEnv []string) []string
	Start(Env []string) error
	Test(Env []string) (Result, error)
}

type StatusDAO interface {
	Update(status Status) error
	Store(status Status) error
	Get(ID string) Status
}
