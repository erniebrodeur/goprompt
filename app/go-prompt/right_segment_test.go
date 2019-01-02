package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RightSegment", func() {
	Describe("Output()", func() {
		It("is expected to be: ├", func() {
			Expect(rightSegment{}.output()).To(Equal(" ├"))
		})
	})

	Describe("Len()", func() {
		It("is expected to be 2", func() {
			Expect(rightSegment{}.len()).To(Equal(2))
		})
	})
})
