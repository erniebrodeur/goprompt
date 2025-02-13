package aggregator_test

import (
"testing"

. "github.com/onsi/ginkgo/v2"
. "github.com/onsi/gomega"
)

// Entry point for the aggregator BDD suite.
func TestAggregatorBDD(t *testing.T) {
RegisterFailHandler(Fail)
RunSpecs(t, "Aggregator BDD Suite")
}

