package aggregator_test

import (
	"context"
	"errors"
	"time"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"yourmodule/internal/aggregator" // Adjust import path as needed
)

type fakeSegment struct {
	text  string
	fail  bool
	delay time.Duration
}

func (f fakeSegment) Output(ctx context.Context) (string, error) {
	// Simulate segment behavior by sleeping and optionally returning an error
	select {
	case <-ctx.Done():
		// Timed out or cancelled
		return "", errors.New("segment timed out")
	case <-time.After(f.delay):
		if f.fail {
			return "", errors.New("segment error")
		}
		return f.text, nil
	}
}

var _ = ginkgo.Describe("Aggregator", func() {

	ginkgo.Context("when all segments succeed quickly", func() {
		ginkgo.It("returns the combined prompt with no errors", func() {
			segs := []aggregator.Segment{
				fakeSegment{text: "DirSegment"},
				fakeSegment{text: "GitSegment"},
			}

			prompt, err := aggregator.BuildPrompt(context.Background(), segs, 100*time.Millisecond)
			Expect(err).NotTo(HaveOccurred())
			Expect(prompt).To(ContainSubstring("DirSegment"))
			Expect(prompt).To(ContainSubstring("GitSegment"))
		})
	})

	ginkgo.Context("when one segment fails immediately", func() {
		ginkgo.It("includes [ERR] for only the failing segment", func() {
			segs := []aggregator.Segment{
				fakeSegment{text: "DirSegment"},
				fakeSegment{fail: true},
			}

			prompt, err := aggregator.BuildPrompt(context.Background(), segs, 100*time.Millisecond)
			Expect(err).NotTo(HaveOccurred())
			Expect(prompt).To(ContainSubstring("DirSegment"))
			Expect(prompt).To(ContainSubstring("[ERR]"))
		})
	})

	ginkgo.Context("when one segment times out", func() {
		ginkgo.It("includes [ERR] only for the timed-out segment", func() {
			segs := []aggregator.Segment{
				fakeSegment{text: "QuickSegment", delay: 10 * time.Millisecond},
				fakeSegment{text: "SlowSegment", delay: 500 * time.Millisecond},
			}

			prompt, err := aggregator.BuildPrompt(context.Background(), segs, 50*time.Millisecond)
			Expect(err).NotTo(HaveOccurred())
			Expect(prompt).To(ContainSubstring("QuickSegment"))
			Expect(prompt).To(ContainSubstring("[ERR]"))
		})
	})

	ginkgo.Context("when all segments fail or time out", func() {
		ginkgo.It("returns a fallback prompt, e.g. `$dir() %`", func() {
			segs := []aggregator.Segment{
				fakeSegment{fail: true},
				fakeSegment{fail: true},
			}

			prompt, err := aggregator.BuildPrompt(context.Background(), segs, 100*time.Millisecond)
			Expect(err).NotTo(HaveOccurred())
			Expect(prompt).To(Equal("$dir() %"))
		})
	})
})