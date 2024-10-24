package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/gin-gonic/gin"

	"example/web-service-gin/entity"
	"example/web-service-gin/fixtures"
	"example/web-service-gin/repository"
	"example/web-service-gin/service/env"
	"example/web-service-gin/service/router"
)

var token string

type WebServiceGinSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *WebServiceGinSuite) SetupSuite() {
	fmt.Println("Setup Suite")
	env.Init(".env.test")
	repository.Setup()
	fixtures.SetupFixtures()
	s.router = router.SetupRouter()
}

func (s *WebServiceGinSuite) TearDownSuite() {
	fmt.Println("TearDown Suite")
}

func (s *WebServiceGinSuite) SetupTest() {
	token := "token"                       // Set token here for each test
	s.T().Logf("Current token: %s", token) // Log token in each test
}

func TestWebServiceGinSuite(t *testing.T) {
	suite.Run(t, new(WebServiceGinSuite))
}

func (s *WebServiceGinSuite) TestHomepageHandler() {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	s.router.ServeHTTP(recorder, req)
	responseData := recorder.Body.String()

	s.T().Run("Health Check", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "\"Ok\"", responseData)
	})
}

func (s *WebServiceGinSuite) TestModelsGetHandler() {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/models", nil)
	s.router.ServeHTTP(recorder, req)

	s.T().Run("Get All models", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, recorder.Code)

		var models []entity.Model
		json.Unmarshal(recorder.Body.Bytes(), &models)
		assert.Equal(t, 9, len(models))

	})
}
