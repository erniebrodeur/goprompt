package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LeftSegment", func() {
	Describe("Output()", func() {
		It("is expected to be:┤ ", func() {
			Expect(leftSegment{}.output()).To(Equal("┤ "))
		})
	})

	Describe("Len()", func() {
		It("is expected to be 2", func() {
			Expect(leftSegment{}.len()).To(Equal(2))
		})
	})
})
