package ytt

import (
	. "code.cloudfoundry.org/cf-for-k8s-ytt-tests/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Networking", func() {

	var ctx RenderingContext
	var data map[string]string
	var templates []string

	BeforeEach(func() {
		templates = []string{
			pathToFile("config/networking"),
			pathToFile("tests/ytt/networking/networking-values.yml"),
		}
	})

	JustBeforeEach(func() {
		ctx = NewRenderingContext(templates...).WithData(data)
	})

	Context("CRD", func() {

		BeforeEach(func() {
			data = map[string]string{}
		})

		It("have a Gateway", func() {

			Expect(ctx).To(ProduceYAML(
				WithGateway("istio-ingressgateway", "foo"),
			))
		})
	})
})
