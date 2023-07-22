package rest

import (
	"net/http/httptest"

	"github.com/davidbenavidez/_git/go-gin-microservice-boilerplate.git/test/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/wI2L/fizz"
)

var _ = Describe("Service", func() {
	var (
		svc      *mocks.MockService
		recorder *httptest.ResponseRecorder
		f        *fizz.Fizz
	)

	BeforeEach(func() {
		svc = mocks.NewMockService(GinkgoT())
		recorder = httptest.NewRecorder()
		f = fizz.New()
		New(svc, f)
	})
	Describe("Create", func() {
		It("Should create something", func() {
			mockResp := []byte(`{"id":"someId"}`)
			req := httptest.NewRequest("POST", "/something/v1/someId", nil)

			svc.EXPECT().DoSomething().Return(mockResp, nil)
			f.ServeHTTP(recorder, req)
			Expect(recorder.Body.String()).To(MatchJSON(`{"id":"someId"}`))
		})
	})
})
