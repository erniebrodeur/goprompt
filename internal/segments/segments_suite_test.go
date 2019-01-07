package segments_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSegments(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Segments Suite")
}
