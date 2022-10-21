package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chen1ting/TravelMaster/internal/models"
	"github.com/chen1ting/TravelMaster/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

var s = service.NewService("testing")

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()

	ctx := GetTestGinContext(w)
	MockJsonGet(ctx, nil, nil)

	s.Ping(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code)

	got, _ := strconv.Atoi(w.Body.String())

	fmt.Println(got)
}

func TestSignUp(t *testing.T) {
	
}

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()

	ctx := GetTestGinContext(w)

	// configure query params
	u := url.Values{}
	u.Add("foo", "bar")

	content := models.LoginReq{
		Username:       "yiting",
		HashedPassword: "123123",
	}

	MockJsonPost(ctx, content)

	s.LoginView(ctx)

	assert.EqualValues(t, http.StatusOK, w.Code)

	got, _ := strconv.Atoi(w.Body.String())

	assert.Equal(t, 1, got)
}

// GetTestGinContext mocks gin context to test service
func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

// MockJsonGet mock the gin getrequest
func MockJsonGet(c *gin.Context, params gin.Params, u url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)

	// set path params
	c.Params = params

	// set query params
	c.Request.URL.RawQuery = u.Encode()
}
