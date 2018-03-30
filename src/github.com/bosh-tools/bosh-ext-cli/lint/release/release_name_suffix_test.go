package release_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	check "github.com/bosh-tools/bosh-ext-cli/lint/check"
	. "github.com/bosh-tools/bosh-ext-cli/lint/release"
)

var _ = Describe("ReleaseNameSuffix", func() {
	var (
		chk ReleaseNameSuffix
	)

	BeforeEach(func() {
		chk = NewReleaseNameSuffix(check.Context{}, "", check.Config{})
	})

	Describe("Check", func() {
		type Example struct {
			Name         string
			ProblemPiece string
		}

		examples := []Example{
			// Problematic
			{Name: "namerelease", ProblemPiece: "Name redundantly ends with 'release'"},
			{Name: "nameboshrelease", ProblemPiece: "Name redundantly ends with 'boshrelease'"},
			{Name: "namebosh-release", ProblemPiece: "Name redundantly ends with 'bosh-release'"},

			// Successful
			{Name: "name"},
			{Name: "name-12"},
			{Name: "name-name"},
			{Name: "name-name-name"},
		}

		for _, ex := range examples {
			ex := ex

			It(fmt.Sprintf("returns suggestion if name is '%s'", ex.Name), func() {
				chk = NewReleaseNameSuffix(check.Context{}, ex.Name, check.Config{})

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
