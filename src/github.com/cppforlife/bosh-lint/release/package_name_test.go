package release_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	check "github.com/cppforlife/bosh-lint/check"
	. "github.com/cppforlife/bosh-lint/release"
)

var _ = Describe("PackageName", func() {
	var (
		chk PackageName
	)

	BeforeEach(func() {
		chk = NewPackageName(check.Context{}, "", check.Config{})
	})

	Describe("Check", func() {
		type Example struct {
			Name         string
			ProblemPiece string
		}

		examples := []Example{
			// Problematic
			{Name: "", ProblemPiece: "Name does not match suggested regexp"},
			{Name: "name-", ProblemPiece: "Name does not match suggested regexp"},
			{Name: "name_", ProblemPiece: "Name does not match suggested regexp"},
			{Name: "name-name", ProblemPiece: "Name does not match suggested regexp"},
			{Name: "name.name", ProblemPiece: "Name does not match suggested regexp"},

			// Successful
			{Name: "name"},
			{Name: "golang_17"},
			{Name: "golang_1_7"},
			{Name: "golang_17_name"},
		}

		for _, ex := range examples {
			ex := ex

			It(fmt.Sprintf("returns suggestion if name is '%s'", ex.Name), func() {
				chk = NewPackageName(check.Context{}, ex.Name, check.Config{})

				sugs, err := chk.Check()
				Expect(err).ToNot(HaveOccurred())
				if len(ex.ProblemPiece) > 0 {
					Expect(sugs).To(HaveLen(1))
					Expect(sugs[0].Problem()).To(ContainSubstring(ex.ProblemPiece))
				} else {
					Expect(sugs).To(HaveLen(0))
				}
			})
		}
	})
})
