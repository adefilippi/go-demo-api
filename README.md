# Gin APi Test


## 1. Install
Install TaskFile:
 ```bash
 sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d
 
 ## For MacOS only
 brew install go-task/tap/go-task
 ```
For more info about TaskFile, visit: https://taskfile.dev/#/installation

### Install project
```bash
task tidy # go mod tidy
```

## 2. Commands
```yaml
## Go
task tidy # go mod tidy
task run # go run . 
task test # GIN_MODE=test go test ./test
task build # go build

## App
# Générate Swagger documentation
task doc # swag init --parseDependency --parseInternal

# Add fixtures
task  fixtures # go run . fixtures

## Docker
task up # docker compose up -d
task down # docker compose down
task restart # task down + task up
```

## 3. Stack
### Packages
- gin-gonic/gin for http server
- swaggo/gin-swagger to generate documentation
- gorm.io/gorm for entity management
- golang-jwt/jwt/v5 and icahParks/keyfunc/v3 for jwt/kwks management
- go-testfixtures/testfixtures/v3 to generate fixtures
- stretchr/testify and h2non/gock for testing
<br/>

### Features
Each <span style="color:#BF4342; font-weight:bold">protected</span> route <span style="color:#BF4342; font-weight:bold">must have middleware.SecurityMiddleware()</span>
<br/>
On each request to API, if route is protected, request must have header <span style="color:#BF4342; font-weight:bold">Authorization</span> (Bearer jwt token) or <span style="color:#BF4342; font-weight:bold">X-API-Key</span> 

## 4. Mock external web service
#### 1. Add new interceptor
```go
//main_test.go
func (s *WebServiceGinSuite) SetupTest() {
	....
	gock.New("https://example.com").
	Get("/endpoint").
        Persist().
        Reply(200).
        BodyString("Expected Response")	
	....
}
```
<br/>

#### 2. Add test
```go
//main_test.go
func (s *WebServiceGinSuite) TestEndpointHandler() {
    s.T().Run("Health Check 2", func(t *testing.T) {
    req, _ := http.NewRequest("GET", "https://example.com/endpoint", nil)
    res, err := (&http.Client{}).Do(req)
    body, _ := ioutil.ReadAll(res.Body)
    assert.Nil(s.T(), err)
    assert.Equal(s.T(), res.StatusCode, 200)
    assert.Equal(t, body, "Expected Response")
    })
}
```