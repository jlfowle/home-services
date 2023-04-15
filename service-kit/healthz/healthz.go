package healthz

import "fmt"

type Test interface {
	Run() error
}

type Check struct {
	Name string
	Test Test
}

type HealthService interface {
	GetHealth() *Report
	GetReady() *Report
	GetLive() *Report
}

type Report struct {
	TestResults []*Result `json:"testResults,omitempty"`
	Healthy     bool     `json:"healthy"`
}

type Result struct {
	Name string `json:"name"`
	Pass bool   `json:"pass"`
	Err error `json:"err,omitempty"`
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

func (h *healthService) GetHealth() *Report {
	return runChecks(append(h.liveChecks, h.readyChecks...))
}

func (h *healthService) GetReady() *Report {
	return runChecks(h.readyChecks)
}

func (h *healthService) GetLive() *Report {
	return runChecks(h.liveChecks)
}

func runChecks(checks []Check) *Report {
	results := []*Result{}
	report := &Report{
		Healthy: true,
	}
	for _, c := range checks {
		res := &Result{
			Name: c.Name,
			Pass: true,
		}
		if err := c.Test.Run();err != nil {
			res.Err = fmt.Errorf("test error: %w", err)
			res.Pass = false
			report.Healthy = false
		}
		results = append(results, res)
	}
	report.TestResults = results
	return report
}
