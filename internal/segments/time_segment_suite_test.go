package segments_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTimeSegmentBDD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Time Segment BDD Suite")
}

