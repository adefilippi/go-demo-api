package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"example/web-service-gin/entity"
	"example/web-service-gin/fixtures"
	"example/web-service-gin/repository"
	"example/web-service-gin/service/env"
	"example/web-service-gin/service/router"

	"example/web-service-gin/test/utils"

	"fmt"
	"reflect"
)

var (
	token string
)

type WebServiceGinSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *WebServiceGinSuite) SetupSuite() {
	env.Init(".env.test")
	repository.Setup()
	s.router = router.SetupRouter()

}

func (s *WebServiceGinSuite) TearDownSuite() {
}

func (s *WebServiceGinSuite) SetupTest() {
	fixtures.SetupFixtures()

	token = "Bearer " + utils.GenerateToken("ROLE_SUPER_ADMIN")

	result := utils.GetJwksString()
	//gock.Observe(gock.DumpRequest)
	gock.New("https://dps-api-auth.herokuapp.com").
		Get("/jwks").
		Persist().
		Reply(200).
		BodyString(result)

	gock.New("https://auth.herokuapp.com").
		Get("/jwks").
		Persist().
		Reply(200).
		BodyString("the response")

}

func (s *WebServiceGinSuite) TearDownTest() {
	// Teardown code for each test
}

func (s *WebServiceGinSuite) TestHomepageHandler() {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/home", nil)
	fmt.Println(reflect.TypeOf(req))
	req.Header.Set("Authorization", token)

	s.router.ServeHTTP(recorder, req)
	responseData := recorder.Body.String()

	s.T().Run("Health Check", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "\"Ok\"", responseData)
		//assert.Equal(t, "a", "b")
	})
}

func (s *WebServiceGinSuite) TestModelsUpdateHandler() {
	recorder := httptest.NewRecorder()
	fmt.Println(reflect.TypeOf(recorder))
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

	headers := map[string]string{
		"Authorization": token,
		"Content-Type":  "application/json",
	}

	req := utils.CreateRequest("PATCH", "/models/1ec5846b-147d-6496-9ee5-0d943c703025", headers, updatedModel)
	s.router.ServeHTTP(recorder, req)

	s.T().Run("Update a model", func(t *testing.T) {
		var modelResponse entity.Model
		err := utils.UnmarshallResponse(recorder, &modelResponse)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(s.T(), updatedModel.ID, modelResponse.ID)
		assert.Equal(s.T(), updatedModel.Name, modelResponse.Name)
	})
}

func (s *WebServiceGinSuite) TestModelsAddFileHandler() {
	recorder := httptest.NewRecorder()
	fmt.Println(reflect.TypeOf(recorder))
	name := "New Model Name"
	title := "New Title"
	subtitle := "New SubTitle"
	description := "New Description"
	newModel := entity.Model{
		Name:        name,
		Title:       &title,
		SubTitle:    &subtitle,
		Description: &description,
		IsNew:       true,
		Price:       29999.99,
		Displayable: true,
	}

	headers := map[string]string{
		"Authorization": token,
		"Content-Type":  "application/json",
	}

	req := utils.CreateRequest("POST", "/models", headers, newModel)
	s.router.ServeHTTP(recorder, req)
	var modelResponse entity.Model
	s.T().Run("Create a model", func(t *testing.T) {

		err := utils.UnmarshallResponse(recorder, &modelResponse)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}
		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.Equal(s.T(), newModel.Name, modelResponse.Name)
	})

	// POST File linked to mode with form-data
	url := fmt.Sprintf("/models/%s/image", modelResponse.ID)

	headers = map[string]string{
		"Authorization": token,
	}
	req = utils.UploadFile(url, headers, []string{"./test/files/test.jpg"})
	recorder = httptest.NewRecorder()
	s.router.ServeHTTP(recorder, req)
	s.T().Run("Create image to model", func(t *testing.T) {
		var mediaObjectResponse entity.MediaObject
		err := utils.UnmarshallResponse(recorder, &mediaObjectResponse)
		if err != nil {
			responseData := recorder.Body.String()
			t.Errorf("Error unmarshalling response: %v - %v", err, responseData)
			return
		}
		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.Equal(s.T(), "test.jpg", *mediaObjectResponse.Name)
		assert.Equal(s.T(), *modelResponse.ID, mediaObjectResponse.ModelID)
		assert.Equal(s.T(), "image/jpeg", *mediaObjectResponse.MimeType)
	})

}

func (s *WebServiceGinSuite) TestModelsCreateHandler() {
	recorder := httptest.NewRecorder()
	fmt.Println(reflect.TypeOf(recorder))
	name := "New Model Name"
	title := "New Title"
	subtitle := "New SubTitle"
	description := "New Description"
	newModel := entity.Model{
		Name:        name,
		Title:       &title,
		SubTitle:    &subtitle,
		Description: &description,
		IsNew:       true,
		Price:       29999.99,
		Displayable: true,
	}

	headers := map[string]string{
		"Authorization": token,
		"Content-Type":  "application/json",
	}

	req := utils.CreateRequest("POST", "/models", headers, newModel)
	s.router.ServeHTTP(recorder, req)

	s.T().Run("Update a model", func(t *testing.T) {
		var modelResponse entity.Model
		err := utils.UnmarshallResponse(recorder, &modelResponse)
		if err != nil {
			t.Errorf("Error unmarshalling response: %v", err)
			return
		}
		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.Equal(s.T(), newModel.Name, modelResponse.Name)
	})

}

func (s *WebServiceGinSuite) TestModelsDeleteHandler() {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/models/1ec5846b-147d-6496-9ee5-0d943c703025", nil)
	req.Header.Set("Authorization", token)
	s.router.ServeHTTP(recorder, req)

	s.T().Run("Delete a model", func(t *testing.T) {
		assert.Equal(t, http.StatusNoContent, recorder.Code)
	})
}

func TestWebServiceGinSuite(t *testing.T) {
	suite.Run(t, new(WebServiceGinSuite))
}
