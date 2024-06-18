package generate_examples_test

import (
	"testing"

	generate_examples "github.com/wasawat.ko/documentgenerator/docs/examples"
)

func TestGenerate(t *testing.T) {
	generate_examples.Generate()
	generate_examples.GenerateMultiplePages()
}
