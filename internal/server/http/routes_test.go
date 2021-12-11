package http_test

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/jlgallego99/OSTfind/internal/server/http"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Routes", func() {
	var router *gin.Engine
	var w *httptest.ResponseRecorder

	BeforeEach(func() {
		router = http.SetupRoutes()
		w = httptest.NewRecorder()
	})
})
