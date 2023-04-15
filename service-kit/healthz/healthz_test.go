package healthz_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jlfowle/home-services/service-kit/healthz"
)

type mockTest struct {
	shouldFail bool
}

func (m *mockTest) Run() error {
	if m.shouldFail {
		return errors.New("failed")
	}
	return nil
}

var _ = Describe("Healthz", func() {
	var (
		dummyService healthz.HealthService
		tests []map[string][]bool = []map[string][]bool{
			{
				"ready_checks": []bool{false},
				"live_checks": []bool{false},
			},
			{
				"ready_checks": []bool{true},
				"live_checks": []bool{false},
			},
			{
				"ready_checks": []bool{false},
				"live_checks": []bool{true},
			},
			{
				"ready_checks": []bool{false, true},
				"live_checks": []bool{false},
			},
			{
				"ready_checks": []bool{false},
				"live_checks": []bool{false, true},
			},
		}
	)

	for _, test := range tests {
		var (
			readyChecks, liveChecks []healthz.Check
			count, expectedPasses, expectedFails int
		)
		for _, checkType := range []string{"health", "ready", "live"} {
			Describe("test health checks", func(){
				BeforeEach(func() {
					for i, c := range test["ready_checks"] {
						if checkType != "live" {
							count++
							if c {
								expectedFails++
							} else {
								expectedPasses++
							}
						}
						readyChecks = append(readyChecks, healthz.Check{Name: fmt.Sprintf("ready%v", i), Test: &mockTest{shouldFail: c}})
					}
					for i, c := range test["live_checks"] {
						if checkType != "ready" {
							count++
							if c {
								expectedFails++
							} else {
								expectedPasses++
							}
						}
						liveChecks = append(liveChecks, healthz.Check{Name: fmt.Sprintf("live%v", i), Test: &mockTest{shouldFail: c}})
					}
					dummyService = healthz.NewHealthService(
						readyChecks,
						liveChecks,
					)
				})
				Describe("check report", func() {
					It("reports correctly", func() {
						var report *healthz.Report
						if checkType == "health" {
							report = dummyService.GetHealth()
						} else if checkType == "ready" {
							report = dummyService.GetReady()
						} else if checkType == "live" {
							report = dummyService.GetLive()
						}
						if expectedFails > 0 {
							Expect(report.Healthy).To(BeFalse())
						} else {
							Expect(report.Healthy).To(BeTrue())
						}
						
						Expect(len(report.TestResults)).To(BeEquivalentTo(count))
						passes := 0
						fails := 0
						for _, t := range report.TestResults {
							if t.Pass {
								Expect(t.Err).ToNot(HaveOccurred())
								passes++
							} else {
								Expect(t.Err).To(HaveOccurred())
								Expect(errors.Unwrap(t.Err)).To(HaveOccurred())
								fails++
							}
						}
						Expect(passes).To(BeEquivalentTo(expectedPasses))
						Expect(fails).To(BeEquivalentTo(expectedFails))
					})
				})
			})
		}
	}
})
