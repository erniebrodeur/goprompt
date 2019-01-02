package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoPrompt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoPrompt Suite")
}
