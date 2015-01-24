package asciidocgo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/bjornerik/asciidocgo/consts/contentModel"
	"github.com/bjornerik/asciidocgo/consts/context"
)

var _ = Describe("the AbstractBlock type", func() {
	var a *abstractBlock

	BeforeEach(func() {
		a = newAbstractBlock(nil, context.Document)
	})

	Context("can be initialized with sane defaults", func() {
		It("should be able to be created by default", func() {
			Expect(&abstractBlock{}).ToNot(BeNil())
			Expect(a).ToNot(BeNil())
		})

		It("should have a 'compound' content model by default", func() {
			Expect(a.ContentModel()).To(Equal(contentmodel.Compound))
		})

		It("should be able to switch compound models", func() {
			a.SetContentModel(contentmodel.Simple)
			Expect(a.ContentModel()).To(Equal(contentmodel.Simple))
		})

		It("should have 0 subs by default", func() {
			Expect(len(a.Subs())).To(Equal(0))
		})

		It("should have a template name of block_#{context}", func() {
			Expect(a.TemplateName()).To(Equal("block_" + a.Context().String()))
		})

		It("should be able to set the template name", func() {
			a.SetTemplateName("my_template")
			Expect(a.TemplateName()).To(Equal("my_template"))
		})

		It("should have 0 blocks by default", func() {
			Expect(len(a.Blocks())).To(Equal(0))
		})

		It("should have a default document context level of 0", func() {
			Expect(a.Level()).To(Equal(0))
		})

		It("should have a level of -1 with no context or parent", func() {
			Expect(newAbstractBlock(nil, context.Section).Level()).To(Equal(-1))
		})

		It("should inherit the block level if child of non-section context", func() {
			parent := newAbstractBlock(nil, context.Document)
			parent.SetLevel(2)
			child := newAbstractBlock(parent, context.Paragraph)

			Expect(child.Level()).To(Equal(2))
		})

		It("should have an empty title by default", func() {
			Expect(a.title).To(Equal(""))
		})

		It("should be able to set the title", func() {
			a.setTitle("foobar")
			Expect(a.title).To(Equal("foobar"))
		})

		It("should have an empty style by default", func() {
			Expect(a.Style()).To(Equal(""))
		})

		It("should be able to set the style", func() {
			a.SetStyle("foobar")
			Expect(a.Style()).To(Equal("foobar"))
		})

		It("should have an empty caption by default", func() {
			Expect(a.Caption()).To(Equal(""))
		})

		It("should be able to set the caption", func() {
			a.SetCaption("foobar")
			Expect(a.Caption()).To(Equal("foobar"))
		})
	})

	Context("can manipulate its context", func() {
		It("should be able to get its context", func() {
			Expect(a.Context()).To(Equal(context.Document))
		})

		It("should be able to set its context", func() {
			a.SetContext(context.Paragraph)
			Expect(a.Context()).To(Equal(context.Paragraph))
		})

		It("should update its template name when context changes", func() {
			Expect(a.TemplateName()).To(Equal("block_document"))
			a.SetContext(context.Paragraph)
			Expect(a.TemplateName()).To(Equal("block_paragraph"))
		})

		It("should be able to render its content", func() {
			a.SetContext(context.Paragraph)

			Expect(a.Render()).To(Equal(""))
		})
	})

	Context("can add and delete sub-blocks", func() {
		It("should return false on a HasBlocks check when there are no blocks", func() {
			Expect(a.HasBlocks()).To(BeFalse())
		})

		It("should return true on a HasBlocks check when there are blocks", func() {
			a.AppendBlock(newAbstractBlock(nil, context.Document))
			Expect(a.HasBlocks()).To(BeTrue())
		})

		It("should be able to add blocks", func() {
			Expect(len(a.Blocks())).To(Equal(0))

			a.AppendBlock(newAbstractBlock(nil, context.Document))

			Expect(len(a.Blocks())).To(Equal(1))
		})

		It("should respond correctly to Sub() queries", func() {
			Expect(a.HasSub("test")).To(BeFalse())

			a.subs = []string{"a", "test", "c"}

			Expect(a.HasSub("test")).To(BeTrue())
		})
	})

	Context("captioned titles", func() {
		It("should be empty by default", func() {
			Expect(a.CaptionedTitle()).To(Equal(""))
		})

		It("should be able to be set", func() {
			a.setTitle("title")
			Expect(a.CaptionedTitle()).To(Equal("title"))
		})

		It("should be based on both the title and the caption", func() {
			a.setTitle("title")
			a.SetCaption("caption")
			Expect(a.CaptionedTitle()).To(Equal("captiontitle"))
		})

		It("should not change the caption if there is already one there", func() {
			a.SetCaption("foo")
			a.AssignCaption("bar", "key")

			Expect(a.CaptionedTitle()).To(Equal("foo"))
		})

		It("should change the caption if one is passed and a title is there", func() {
			a.SetCaption("foo")
			a.setTitle("title")
			a.AssignCaption("bar", "key")

			Expect(a.CaptionedTitle()).To(Equal("bartitle"))
		})

		XIt("should not assign a caption if only a title exists", func() {
			// what the fuck -- this is incomplete b/c I don't understand the
			// original test case. will convert when I figure it out (or redo
			// the whole thing)
		})
	})

	Context("the relationship between blocks and sections", func() {
		It("should be empty by default", func() {
			Expect(len(a.Sections())).To(Equal(0))
		})

		// FIXME: why is this even broken?
		XIt("should be able to add new sections", func() {
			a.AppendBlock(newAbstractBlock(nil, context.Document))
			Expect(len(a.Sections())).To(Equal(1))
		})
	})

	Context("abstract block subs", func() {
		It("should be able to remove subs", func() {
			a.subs = append(a.subs, "foo", "bar", "baz")

			Expect(len(a.subs)).To(Equal(3))

			a.RemoveSub("foo")
			a.RemoveSub("bar")

			Expect(a.HasSub("foo")).To(BeFalse())
			Expect(a.HasSub("bar")).To(BeFalse())
			Expect(a.HasSub("baz")).To(BeTrue())

			Expect(len(a.subs)).To(Equal(1))
		})
	})
})
