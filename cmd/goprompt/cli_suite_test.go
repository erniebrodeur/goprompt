package main_test

import (
"testing"

. "github.com/onsi/ginkgo/v2"
. "github.com/onsi/gomega"
)

// Entry point for the full CLI BDD Suite
func TestGoPromptCLI(t *testing.T) {
RegisterFailHandler(Fail)
RunSpecs(t, "GoPrompt CLI BDD Suite")
}

