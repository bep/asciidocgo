package asciidocgo

import (
	"os"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the path_resolver module", func() {
	Context("initialization", func() {
		It("can be created using the default config", func() {
			Expect(NewPathResolver(0, "")).ToNot(BeNil())
		})

		It("should use a system path separator by default", func() {
			//Expect(NewPathResolver(0, "").FileSeparator()).To(Equal(os.PathSeparator))
			Expect(NewPathResolver('/', "").FileSeparator()).To(BeNumerically("==", '/'))
			Expect(NewPathResolver('\\', "").FileSeparator()).To(BeNumerically("==", '\\'))
		})

		It("should be set to the current working path by default", func() {
			pwd, _ := os.Getwd()

			Expect(NewPathResolver(0, "").WorkingDir()).To(Equal(pwd))
			Expect(NewPathResolver(0, `/`).WorkingDir()).To(Equal(`/`))
			Expect(NewPathResolver(0, "test").WorkingDir()).To(Equal(pwd + string(os.PathSeparator) + "test"))
		})
	})

	Context("web path", func() {
		It("can test for a web path", func() {
			Expect(IsWebRoot("")).To(BeFalse())
			Expect(IsWebRoot("a")).To(BeFalse())
			Expect(IsWebRoot(`\a\b/c`)).To(BeFalse())

			Expect(IsWebRoot(`/a/b/c`)).To(BeTrue())
		})
	})

	Context("posixify", func() {
		It("should not adjust an empty path", func() {
			Expect(Posixfy((""))).To(Equal(""))
		})

		It("should be idempotent", func() {
			Expect(Posixfy(Posixfy(("/hello/world")))).To(Equal("/hello/world"))
		})

		It("should be replace backslashes with slashes", func() {
			Expect(Posixfy((`a\b\c`))).To(Equal("a/b/c"))
		})
	})

	Context("root check", func() {
		// not really sure what this is intended to do, as it's a useless
		// wrapper over the path pkg's isAbs (which doesn't do anything funny
		// with Windows paths)?
		XIt("should return true for root paths on Windows", func() {
			Expect(IsRoot(`C:\`)).To(BeTrue())
			Expect(IsRoot(`C:/`)).To(BeTrue())
			Expect(IsRoot(`C:\a/b/`)).To(BeTrue())
			Expect(IsRoot(`C:/a/b/../c`)).To(BeTrue())
			Expect(IsRoot(`C:\a/b/../c`)).To(BeTrue())
		})

		It("should return false for relative paths on Windows", func() {
			Expect(IsRoot(`a\b\c\`)).To(BeFalse())
			Expect(IsRoot(`..\a\b\c`)).To(BeFalse())
		})

		It("should return true for root paths on a Unixy OS", func() {
			Expect(IsRoot(`/`)).To(BeTrue())
			Expect(IsRoot(`/foo/bar/baz`)).To(BeTrue())
			Expect(IsRoot(`/bar/../foo/baz`)).To(BeTrue())
		})

		It("should return false for relative paths on a Unixy OS", func() {
			Expect(IsRoot(`test`)).To(BeFalse())
			Expect(IsRoot(`a/b/c`)).To(BeFalse())
			Expect(IsRoot(`../a/../b/c`)).To(BeFalse())
		})
	})

	Context("webroot check", func() {
		It("should report false for a path not beginning with /", func() {
			Expect(IsWebRoot("")).To(BeFalse())
			Expect(IsWebRoot("C:\\")).To(BeFalse())
			Expect(IsWebRoot("\\")).To(BeFalse())
		})

		It("should report true for a path beginning with /", func() {
			Expect(IsWebRoot("/")).To(BeTrue())
			Expect(IsWebRoot("/a/b/")).To(BeTrue())
			Expect(IsWebRoot("/a\\b/./c")).To(BeTrue())
			Expect(IsWebRoot("/a/b/./c")).To(BeTrue())
		})
	})

	Context("expand path", func() {
		It("should return an empty string when expanding an empty path", func() {
			Expect(ExpandPath("")).To(Equal(""))
		})

		It("should return a posix path when provided with a weird mash", func() {
			Expect(ExpandPath(`c:\a/.\b/../c`)).To(Equal("c:/a/b/../c"))
		})
	})

	Context("path partitioning", func() {
		//TODO
	})
})
