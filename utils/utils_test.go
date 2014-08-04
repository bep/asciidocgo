package utils

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the utils package", func() {
	Context("the Mult func", func() {
		It("should return an empty string when passed an empty array", func() {
			Expect(Arr{}.Mult("|")).To(Equal(""))
		})

		It("should return the only element of a one element array", func() {
			Expect(Arr{"a"}.Mult("|")).To(Equal("a"))
		})

		It("should return elems separated correctly when passed a multi-elem array", func() {
			Expect(Arr{"a", "b"}.Mult("|")).To(Equal("a|b"))
			Expect(Arr{"a", "b", "c"}.Mult("|")).To(Equal("a|b|c"))
			Expect(Arr{"a", "b", "c", "d"}.Mult("|")).To(Equal("a|b|c|d"))
			Expect(Arr{"a", "b", "c", "d", "e"}.Mult("|")).To(Equal("a|b|c|d|e"))
		})
	})
})
