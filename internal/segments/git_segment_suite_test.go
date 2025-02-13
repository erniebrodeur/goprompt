package segments_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Entry point for the Git segment BDD suite.
func TestGitSegmentBDD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Git Segment BDD Suite")
}

