package test

import (
	"io"
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"

	"example/web-service-gin/entity"
	"example/web-service-gin/fixtures"
	"example/web-service-gin/repository"
	"example/web-service-gin/service/env"
	"example/web-service-gin/service/router"

	"example/web-service-gin/test/utils"
)

var token string

type WebServiceGinSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *WebServiceGinSuite) SetupSuite() {
	env.Init(".env.test")
	repository.Setup()
	s.router = router.SetupRouter()
	token = utils.GetToken()
}

func (s *WebServiceGinSuite) TearDownSuite() {
}

func (s *WebServiceGinSuite) SetupTest() {
	fixtures.SetupFixtures()
}

func (s *WebServiceGinSuite) TearDownTest() {
	// Teardown code for each test
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

func (s *WebServiceGinSuite) TestModelsUpdateHandler() {
	recorder := httptest.NewRecorder()

	id := uuid.MustParse("1ec5846b-0068-621c-82a9-0d943c703025")
	name := "Updated Model Name"
	title := "Updated Title"
	subtitle := "Updated SubTitle"
	description := "Updated Description"
	updatedModel := entity.Model{
		ID:          &id,
		Name:        name,
		Title:       &title,
		SubTitle:    &subtitle,
		Description: &description,
		IsNew:       true,
		Price:       29999.99,
		Slug:        "updated-model-name",
		Displayable: true,
	}

	body, err := json.Marshal(updatedModel)
	if err != nil {
		s.T().Fatal(err)
	}

	req := httptest.NewRequest("PATCH", "/models/1ec5846b-147d-6496-9ee5-0d943c703025", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(recorder, req)

	s.T().Run("Update a model", func(t *testing.T) {

		var modelResponse entity.Model
		responseBody, err := io.ReadAll(recorder.Body)
		if err != nil {
			s.T().Fatal(err)
		}

		if err := json.Unmarshal(responseBody, &modelResponse); err != nil {
			s.T().Fatal(err)
		}

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(s.T(), updatedModel.ID, modelResponse.ID)
		assert.Equal(s.T(), updatedModel.Name, modelResponse.Name)
	})
}

func (s *WebServiceGinSuite) TestModelsDeleteHandler() {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/models/1ec5846b-147d-6496-9ee5-0d943c703025", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	s.router.ServeHTTP(recorder, req)

	s.T().Run("Delete a model", func(t *testing.T) {
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})
}

func TestWebServiceGinSuite(t *testing.T) {
	suite.Run(t, new(WebServiceGinSuite))
}
