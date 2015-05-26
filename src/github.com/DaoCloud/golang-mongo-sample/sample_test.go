package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample", func() {
	Describe("Adding and retreving Person object from MongoDB", func() {
		Context("inspecting their name", func() {
			It("should result 'Lee'", func() {
				Expect(GetResult()).To(Equal("Lee"))
			})
		})
	})
})
