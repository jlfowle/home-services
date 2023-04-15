package healthz_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jlfowle/home-services/service-kit/healthz"
)

var (
	dummyService healthz.HealthService
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
	Context("with all passing checks", func() {
		BeforeEach(func() {
			dummyService = healthz.NewHealthService(
				[]healthz.Check{
					{Name: "ready", Test: &mockTest{}},
				},
				[]healthz.Check{
					{Name: "live", Test: &mockTest{}},
				},
			)
		})
		It("Is healthy", func() {
			report := dummyService.GetHealth()
			Expect(report.Healthy).To(BeTrue())
			Expect(len(report.TestResults)).To(BeEquivalentTo(2))
			for _, t := range report.TestResults {
				Expect(t.Err).ToNot(HaveOccurred())
				Expect(t.Pass).To(BeTrue())
			}
		})
		It("Is ready", func() {
			report := dummyService.GetReady()
			Expect(report.Healthy).To(BeTrue())
			Expect(len(report.TestResults)).To(BeEquivalentTo(1))
			for _, t := range report.TestResults {
				Expect(t.Err).ToNot(HaveOccurred())
				Expect(t.Pass).To(BeTrue())
			}
		})
		It("Is live", func() {
			report := dummyService.GetLive()
			Expect(report.Healthy).To(BeTrue())
			Expect(len(report.TestResults)).To(BeEquivalentTo(1))
			for _, t := range report.TestResults {
				Expect(t.Err).ToNot(HaveOccurred())
				Expect(t.Pass).To(BeTrue())
			}
		})
	})
	Context("with all failing checks", func() {
		BeforeEach(func() {
			dummyService = healthz.NewHealthService(
				[]healthz.Check{
					{Name: "ready", Test: &mockTest{shouldFail: true}},
				},
				[]healthz.Check{
					{Name: "live", Test: &mockTest{shouldFail: true}},
				},
			)
		})
		It("Is healthy", func() {
			report := dummyService.GetHealth()
			Expect(report.Healthy).To(BeFalse())
			Expect(len(report.TestResults)).To(BeEquivalentTo(2))
			for _, t := range report.TestResults {
				Expect(t.Err).To(HaveOccurred())
				Expect(t.Pass).To(BeFalse())
				Expect(errors.Unwrap(t.Err)).To(HaveOccurred())
			}
		})
		It("Is ready", func() {
			report := dummyService.GetReady()
			Expect(report.Healthy).To(BeFalse())
			Expect(len(report.TestResults)).To(BeEquivalentTo(1))
			for _, t := range report.TestResults {
				Expect(t.Err).To(HaveOccurred())
				Expect(t.Pass).To(BeFalse())
				Expect(errors.Unwrap(t.Err)).To(HaveOccurred())
			}
		})
		It("Is live", func() {
			report := dummyService.GetLive()
			Expect(report.Healthy).To(BeFalse())
			Expect(len(report.TestResults)).To(BeEquivalentTo(1))
			for _, t := range report.TestResults {
				Expect(t.Err).To(HaveOccurred())
				Expect(t.Pass).To(BeFalse())
				Expect(errors.Unwrap(t.Err)).To(HaveOccurred())
			}
		})
	})
	Context("with failing ready but passing live", func() {
		BeforeEach(func() {
			dummyService = healthz.NewHealthService(
				[]healthz.Check{
					{Name: "ready", Test: &mockTest{shouldFail: true}},
				},
				[]healthz.Check{
					{Name: "live", Test: &mockTest{}},
				},
			)
		})
		It("Is healthy", func() {
			report := dummyService.GetHealth()
			Expect(report.Healthy).To(BeFalse())
			Expect(len(report.TestResults)).To(BeEquivalentTo(2))
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
			Expect(passes).To(BeEquivalentTo(1))
			Expect(fails).To(BeEquivalentTo(1))
		})
		It("Is ready", func() {
			report := dummyService.GetReady()
			Expect(report.Healthy).To(BeFalse())
			Expect(len(report.TestResults)).To(BeEquivalentTo(1))
			for _, t := range report.TestResults {
				Expect(t.Err).To(HaveOccurred())
				Expect(t.Pass).To(BeFalse())
				Expect(errors.Unwrap(t.Err)).To(HaveOccurred())
			}
		})
		It("Is live", func() {
			report := dummyService.GetLive()
			Expect(report.Healthy).To(BeTrue())
			Expect(len(report.TestResults)).To(BeEquivalentTo(1))
			for _, t := range report.TestResults {
				Expect(t.Err).ToNot(HaveOccurred())
				Expect(t.Pass).To(BeTrue())
			}
		})
	})
	Context("with failing live but passing ready", func() {
		BeforeEach(func() {
			dummyService = healthz.NewHealthService(
				[]healthz.Check{
					{Name: "ready", Test: &mockTest{}},
				},
				[]healthz.Check{
					{Name: "live", Test: &mockTest{shouldFail: true}},
				},
			)
		})
		It("Is healthy", func() {
			report := dummyService.GetHealth()
			Expect(report.Healthy).To(BeFalse())
			Expect(len(report.TestResults)).To(BeEquivalentTo(2))
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
			Expect(passes).To(BeEquivalentTo(1))
			Expect(fails).To(BeEquivalentTo(1))
		})
		It("Is ready", func() {
			report := dummyService.GetReady()
			Expect(report.Healthy).To(BeTrue())
			Expect(len(report.TestResults)).To(BeEquivalentTo(1))
			for _, t := range report.TestResults {
				Expect(t.Err).ToNot(HaveOccurred())
				Expect(t.Pass).To(BeTrue())
			}
		})
		It("Is live", func() {
			report := dummyService.GetLive()
			Expect(report.Healthy).To(BeFalse())
			Expect(len(report.TestResults)).To(BeEquivalentTo(1))
			for _, t := range report.TestResults {
				Expect(t.Err).To(HaveOccurred())
				Expect(t.Pass).To(BeFalse())
				Expect(errors.Unwrap(t.Err)).To(HaveOccurred())
			}
		})
	})
	Context("with some failing health", func() {
		BeforeEach(func() {
			dummyService = healthz.NewHealthService(
				[]healthz.Check{
					{Name: "ready", Test: &mockTest{}},
					{Name: "ready", Test: &mockTest{shouldFail: true}},
				},
				[]healthz.Check{
					{Name: "live", Test: &mockTest{}},
				},
			)
		})
		It("Is healthy", func() {
			report := dummyService.GetHealth()
			Expect(report.Healthy).To(BeFalse())
			Expect(len(report.TestResults)).To(BeEquivalentTo(3))
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
			Expect(passes).To(BeEquivalentTo(2))
			Expect(fails).To(BeEquivalentTo(1))
		})
		It("Is ready", func() {
			report := dummyService.GetReady()
			Expect(report.Healthy).To(BeFalse())
			Expect(len(report.TestResults)).To(BeEquivalentTo(2))
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
			Expect(passes).To(BeEquivalentTo(1))
			Expect(fails).To(BeEquivalentTo(1))
		})
		It("Is live", func() {
			report := dummyService.GetLive()
			Expect(report.Healthy).To(BeTrue())
			Expect(len(report.TestResults)).To(BeEquivalentTo(1))
			for _, t := range report.TestResults {
				Expect(t.Err).ToNot(HaveOccurred())
				Expect(t.Pass).To(BeTrue())
			}
		})
	})
	Context("with some failing health", func() {
		BeforeEach(func() {
			dummyService = healthz.NewHealthService(
				[]healthz.Check{
					{Name: "ready", Test: &mockTest{}},
				},
				[]healthz.Check{
					{Name: "live1", Test: &mockTest{}},
					{Name: "live2", Test: &mockTest{shouldFail: true}},
				},
			)
		})
		It("Is healthy", func() {
			report := dummyService.GetHealth()
			Expect(report.Healthy).To(BeFalse())
			Expect(len(report.TestResults)).To(BeEquivalentTo(3))
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
			Expect(passes).To(BeEquivalentTo(2))
			Expect(fails).To(BeEquivalentTo(1))
		})
		It("Is ready", func() {
			report := dummyService.GetReady()
			Expect(report.Healthy).To(BeTrue())
			Expect(len(report.TestResults)).To(BeEquivalentTo(1))
			for _, t := range report.TestResults {
				Expect(t.Err).ToNot(HaveOccurred())
				Expect(t.Pass).To(BeTrue())
			}
		})
		It("Is live", func() {
			report := dummyService.GetLive()
			Expect(report.Healthy).To(BeFalse())
			Expect(len(report.TestResults)).To(BeEquivalentTo(2))
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
			Expect(passes).To(BeEquivalentTo(1))
			Expect(fails).To(BeEquivalentTo(1))
		})
	})
})
