package asciidocgo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the Document type", func() {
	var d *Document

	BeforeEach(func() {
		d = NewDocument([]string{}, map[string]string{})
	})

	Context("a Document can be monitored", func() {
		It("should be unmonitored by default", func() {
			Expect(d.IsMonitored()).To(BeFalse())
		})

		It("should be monitored when monitoring is turned on", func() {
			d.Monitor()

			Expect(d.IsMonitored()).To(BeTrue())
		})
	})

	Context("time reporting", func() {
		var load, read, parse, load_render, total, write int

		BeforeEach(func() {
			load, _ = d.LoadTime()
			read, _ = d.ReadTime()
			parse, _ = d.ParseTime()
			load_render, _ = d.LoadRenderTime()
			total, _ = d.TotalTime()
			write, _ = d.WriteTime()
		})

		It("should calculate load time by adding read time and parse time", func() {
			Expect(load).To(Equal(read + parse))
		})

		It("should calculate render time by adding load time and render time", func() {
			Expect(parse).To(Equal(read + parse))
		})

		It("should calculate total time by adding loadRender time and write time", func() {
			Expect(total).To(Equal(load_render + write))
		})
	})
})
