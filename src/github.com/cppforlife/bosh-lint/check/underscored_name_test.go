package check_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	check "github.com/cppforlife/bosh-lint/check"
)

var _ = Describe("UnderscoredName", func() {
	var (
		chk check.UnderscoredName
	)

	BeforeEach(func() {
		chk = check.NewUnderscoredName(check.Context{}, "")
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
			{Name: "_name", ProblemPiece: "Name does not match suggested regexp"},
			{Name: "_name_", ProblemPiece: "Name does not match suggested regexp"},
			{Name: "name-name", ProblemPiece: "Name does not match suggested regexp"},
			{Name: "name.name", ProblemPiece: "Name does not match suggested regexp"},

			// Successful
			{Name: "name"},
			{Name: "app_17"},
			{Name: "app_17_name"},
		}

		for _, ex := range examples {
			ex := ex

			It(fmt.Sprintf("returns suggestion if name is '%s'", ex.Name), func() {
				chk = check.NewUnderscoredName(check.Context{}, ex.Name)

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
