package e2etest

// Result represents the test result
type Result struct {
	Pass        bool
	Description string
	Context     string
}

// Status is DTO object to combine commit status with its result
type Status struct {
	ID          string
	State       string
	TargetURL   string
	Description string
	Context     string
}

// Main is the main method to run though pipeline
func Main(pipeline Pipeline, url, hash string) (Result, error) {
	result := Result{}

	err := pipeline.Clone(url)
	if err != nil {
		return result, err
	}

	buildCommand := "./scripts/build.sh"
	err = pipeline.Build(buildCommand)
	if err != nil {
		result.Description = "failed to build repository"
		result.Context = err.Error()
		return result, nil
	}

	env := pipeline.Env()

	stop := make(chan bool, 1)
	startCommand := "./scripts/start.sh"
	err = pipeline.Start(env, startCommand, stop)
	defer func() {
		stop <- true
	}()
	if err != nil {
		result.Description = "failed to start application"
		result.Context = err.Error()
		return result, nil
	}

	result, err = pipeline.Test(env)
	if err != nil {
		result.Description = "failed to start test"
		result.Context = err.Error()
		return result, nil
	}
	return result, nil
}

// Pipeline defines each step in the test
type Pipeline interface {
	Clone(URL string) error
	Build(command string) error
	Env() []string
	Start(Env []string, command string, stop <-chan bool) error
	Test(Env []string) (Result, error)
}

// StatusDAO defines the interaction status with DB
type StatusDAO interface {
	UpdateStatus(status Status) error
	CreateStatus(status Status) error
	Log(ID, data string)
	GetStatus(ID string) Status
}

// TokenDAO defines the services layer need for status
type TokenDAO interface {
	StoreToken(token string) error
	GetToken() string
}
