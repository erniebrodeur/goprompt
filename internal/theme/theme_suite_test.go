package theme_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestThemeBDD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Theme BDD Suite")
}

