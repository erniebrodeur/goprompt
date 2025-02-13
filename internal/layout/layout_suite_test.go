package layout_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLayoutBDD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Layout BDD Suite")
}

