package service

import (
	"github.com/davidbenavidez/_git/go-gin-microservice-boilerplate.git/test/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("service", func() {
	var (
		mockDb *mocks.MockDBClient
		svc    Service
	)

	BeforeEach(func() {
		mockDb = mocks.NewMockDBClient(GinkgoT())
		svc = New(mockDb)
	})

	Describe("DoSomething", func() {
		When("some condition", func() {
			It("Should do something because of condition", func() {
				mockResp := []byte(`{"id":"someId"}`)
				mockDb.EXPECT().Create().Return(mockResp, nil)
				result, err := svc.DoSomething()
				Expect(err).To(BeNil())
				Expect(string(result)).To(Equal(string(mockResp)))
			})
		})
	})
})
