package segments_test

import (
	"testing"

	"github.com/novln/macchiato"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSegments(t *testing.T) {
	RegisterFailHandler(Fail)
	macchiato.RunSpecs(t, "Segments Suite")

	// RunSpecs(t, "Segments Suite")
}
