package asciidocgo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Renderer", func() {
	It("can be initialized", func() {
		Expect(&Renderer{}).ToNot(BeNil())
	})

	It("can render a template", func() {
		r := &Renderer{}
		Expect(r.Render("", nil, nil)).To(Equal(""))
	})
})
