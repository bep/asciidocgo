package asciidocgo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAsciidocgoSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Asciidocgo Suite")
}
