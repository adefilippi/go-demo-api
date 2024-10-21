package test

import (
	"encoding/json"
	"testing"
	"net/http/httptest"
	"net/http"

	"github.com/gin-gonic/gin"

	"example/web-service-gin/repository"
	"example/web-service-gin/entity"
	"example/web-service-gin/api"
	"example/web-service-gin/service/router"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	r.GET("models", api.GetModels)
	r.GET("/", api.Home)
	return r
}

func TestHomepageHandler(t *testing.T) {
	repository.Setup()
	r := router.SetupRouter()
	recorder := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/", nil)

	r.ServeHTTP(recorder, req)
	responseData := recorder.Body.String()
	t.Run("Get All models", func(t *testing.T) {
		if recorder.Code != http.StatusOK {
			t.Error("Expected 200, got ", recorder.Code, responseData)
		}

		if responseData != "Ok" {
			t.Error("Expected response Ok, got ", responseData)
		}

	})
}

func TestModelsGetHandler(t *testing.T) {
	repository.Setup()
	r := router.SetupRouter()
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/models", nil)

	r.ServeHTTP(recorder, req)
	responseData := recorder.Body.String()
	t.Run("Get All models", func(t *testing.T) {
		if recorder.Code != http.StatusOK {
			t.Error("Expected 200, got ", recorder.Code, responseData)
		}
		var models []entity.Model
		json.Unmarshal(recorder.Body.Bytes(), &models)
		t.Error("Expected 200, got ", len(models))
	})
}
