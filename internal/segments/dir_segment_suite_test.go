package segments_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Entry point for DirSegment BDD tests
func TestDirSegmentBDD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dir Segment BDD Suite")
}

