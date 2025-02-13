package aggregator_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAggregatorBDD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Aggregator BDD Suite")
}

