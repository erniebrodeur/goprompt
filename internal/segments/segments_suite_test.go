package segments_test

import (
	"testing"

	"github.com/capitancambio/reporter"
	"github.com/novln/macchiato"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSegments(t *testing.T) {
	RegisterFailHandler(Fail)
	reporter := reporter.NewGoTestCompatibleReporter()
	RunSpecsWithCustomReporters(t, "Suit", []Reporter{reporter})
	macchiato.RunSpecs(t, "Segments Suite")

	// RunSpecs(t, "Segments Suite")
}
