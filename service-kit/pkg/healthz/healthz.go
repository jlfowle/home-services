package healthz

type Test interface {
	Run() error
}

type Check struct {
	Name string
	Test Test
}

type HealthService interface {
	GetHealth() error
	GetReady() error
	GetLive() error
}

type Report struct {
	TestResults []Result `json:"testResults,omitempty"`
	Healthy     bool     `json:"healthy"`
}

type Result struct {
	Name string `json:"name"`
	Pass bool   `json:"pass"`
	Message string `json:"message,omitempty"`
}

func NewHealthService(readyChecks, liveChecks []Check) HealthService {
	return &healthService{
		readyChecks: readyChecks,
		liveChecks:  liveChecks,
	}
}

type healthService struct {
	readyChecks []Check
	liveChecks  []Check
}

func (h *healthService) GetHealth() Report {
	return runChecks(h.liveChecks..., h.readyChecks...)
}

func (healthService) GetReady() Report {
	return runChecks(h.readyChecks...)
}

func (healthService) GetLive() Report {
	return runChecks(h.liveChecks...)
}

func runChecks(checks []Check) Report {
	results := []Result{}
	report := &Report{
		TestResults: results,
		Healthy: true,
	}
	for _, c := range checks {
		res := &Result{
			Name: c.Name,
			Pass: true,
		}
		if err := c.Test.Run();err != nil {
			res.Message = err.Error()
			res.Pass = false
			report.Healthy = false
		}
		results = append(results, res)
	}
	return report
}
