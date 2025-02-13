package segments_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGitSegmentBDD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Git Segment BDD Suite")
}

