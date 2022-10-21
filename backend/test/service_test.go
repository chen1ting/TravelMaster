package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/chen1ting/TravelMaster/internal/models"
	"github.com/chen1ting/TravelMaster/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type errorMessage struct {
	Error string `json:"error"`
}

var (
	s                  = service.NewService("testing")
	sampleSessionToken string
	sampleUserID       int64
	// error response for sign_up endpoint
	ErrMissingUserInfo   = errors.New("eight email, username, or hashed password missing")
	ErrUserAlreadyExists = errors.New("user already exists")

	// error response for login request
	ErrInvalidLogin = errors.New("invalid login")

	ErrUserNotExist = errors.New("user not exist")

	ErrUserAlreadyCreatedReview = errors.New("user already created review for the activity")
	ErrNotAllowed               = errors.New("user is not allowed to perform this action")
	ErrGenericServerError       = errors.New("generic server error")
	ErrDatabase                 = errors.New("database error")
	ErrBadRequest               = errors.New("bad request")
	ErrActivityAlreadyExists    = errors.New("an activity with the same title already exists")
	ErrActivityNotFound         = errors.New("activity not found")
	ErrNullTitle                = errors.New("title cannot be empty")
	ErrUserNotFound             = errors.New("user id doesn't exists")
	ErrReviewNotFound           = errors.New("review not found")
	ErrReportNotFound           = errors.New("report not found")
	ErrInvalidUpdateUser        = errors.New("user id doesn't match the activity's user id")
	ErrNoSearchFail             = errors.New("search Name failed")
	ErrAlreadyReported          = errors.New("user has already reported the activity")
	ErrUnknownFileType          = errors.New("unknown file type uploaded")
	ErrImageNoMatch             = errors.New("image not found in the list of the activity")
	ErrImageNotFound            = errors.New("image not found on server, removed file name in the database")
)

func TestSignup(t *testing.T) {
	t.Run("signup_missing_email", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		// write form body
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		assert.NoError(
			t, writer.WriteField("username", "yiting"))
		assert.NoError(
			t, writer.WriteField("hashed_password", "real_password"))
		assert.NoError(
			t, writer.WriteField("email", ""))

		fileWriter, err := writer.CreateFormFile("avatar", "test.png")
		if assert.NoError(t, err) {
			_, err = fileWriter.Write([]byte("test.png"))
			assert.NoError(t, err)
		}
		if _, err = t, writer.Close(); err != nil {
			log.Fatalln(err)
		}
		ctx.Request.Body = io.NopCloser(body)

		// write form header
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", writer.FormDataContentType())

		// send request to endpoint
		s.SignupView(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrMissingUserInfo.Error(), got.Error, "signup_missing_email")
	})

	t.Run("signup_missing_username", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		// write form body
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		assert.NoError(
			t, writer.WriteField("username", ""))
		assert.NoError(
			t, writer.WriteField("hashed_password", "real_password"))
		assert.NoError(
			t, writer.WriteField("email", "yiting@travelmaster.com"))

		fileWriter, err := writer.CreateFormFile("avatar", "test.png")
		if assert.NoError(t, err) {
			_, err = fileWriter.Write([]byte("test.png"))
			assert.NoError(t, err)
		}
		if _, err = t, writer.Close(); err != nil {
			log.Fatalln(err)
		}
		ctx.Request.Body = io.NopCloser(body)

		// write form header
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", writer.FormDataContentType())

		// send request to endpoint
		s.SignupView(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrMissingUserInfo.Error(), got.Error, "signup_missing_username")
	})

	t.Run("signup_wrong_filetype", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		// write form body
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		assert.NoError(
			t, writer.WriteField("username", "yiting_no_avatar"))
		assert.NoError(
			t, writer.WriteField("hashed_password", "real_password"))
		assert.NoError(
			t, writer.WriteField("email", "yiting_no_avatar@travelmaster.com"))

		fileWriter, err := writer.CreateFormFile("avatar", "test.txt")
		if assert.NoError(t, err) {
			_, err = fileWriter.Write([]byte("test.txt"))
			assert.NoError(t, err)
		}
		if _, err = t, writer.Close(); err != nil {
			log.Fatalln(err)
		}
		ctx.Request.Body = io.NopCloser(body)

		// write form header
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", writer.FormDataContentType())

		// send request to endpoint
		s.SignupView(ctx)

		var got models.SignupResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		canCompare := models.SignupResp{
			Username:   got.Username,
			Email:      got.Email,
			AvatarName: got.AvatarName,
		}
		want := models.SignupResp{
			Username:   "yiting_no_avatar",
			Email:      "yiting_no_avatar@travelmaster.com",
			AvatarName: "",
		}
		assert.EqualValues(t, http.StatusCreated, w.Code)
		assertEqual(t, want, canCompare, "signup_success_avatar_unsaved")
	})

	t.Run("signup_success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		// write form body
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		assert.NoError(
			t, writer.WriteField("username", "yiting"))
		assert.NoError(
			t, writer.WriteField("hashed_password", "real_password"))
		assert.NoError(
			t, writer.WriteField("email", "yiting@travelmaster.com"))

		fileWriter, err := writer.CreateFormFile("avatar", "test.png")
		if assert.NoError(t, err) {
			_, err = fileWriter.Write([]byte("test.png"))
			assert.NoError(t, err)
		}
		if _, err = t, writer.Close(); err != nil {
			log.Fatalln(err)
		}
		ctx.Request.Body = io.NopCloser(body)

		// write form header
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", writer.FormDataContentType())

		// send request to endpoint
		s.SignupView(ctx)

		var got models.SignupResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		canCompare := models.SignupResp{
			Username: got.Username,
			Email:    got.Email,
		}
		want := models.SignupResp{
			Username: "yiting",
			Email:    "yiting@travelmaster.com",
		}
		sampleSessionToken = got.SessionToken
		sampleUserID = got.UserId
		assert.EqualValues(t, http.StatusCreated, w.Code)
		assertEqual(t, want, canCompare, "signup_success")
	})

	t.Run("signup_user_existed", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		// write form body
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		assert.NoError(
			t, writer.WriteField("username", "yiting"))
		assert.NoError(
			t, writer.WriteField("hashed_password", "real_password"))
		assert.NoError(
			t, writer.WriteField("email", "yiting@travelmaster.com"))

		fileWriter, err := writer.CreateFormFile("avatar", "test.png")
		if assert.NoError(t, err) {
			_, err = fileWriter.Write([]byte("test.png"))
			assert.NoError(t, err)
		}
		if _, err = t, writer.Close(); err != nil {
			log.Fatalln(err)
		}
		ctx.Request.Body = io.NopCloser(body)

		// write form header
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", writer.FormDataContentType())

		// send request to endpoint
		s.SignupView(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrUserAlreadyExists.Error(), got.Error, "signup_success")
	})

}

func TestLogin(t *testing.T) {
	t.Run("login_user_not_existed", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := models.LoginReq{
			Username:       "jane_doe",
			HashedPassword: "123123",
		}
		MockJsonPost(ctx, req)
		s.LoginView(ctx)
		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrInvalidLogin.Error(), got.Error, "login_user_not_existed_error")
	})

	t.Run("login_user_wrong_pwd", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := models.LoginReq{
			Username:       "yiting",
			HashedPassword: "123123",
		}
		MockJsonPost(ctx, req)
		s.LoginView(ctx)
		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrInvalidLogin.Error(), got.Error, "login_user_wrong_pwd_error")
	})

	t.Run("login_success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := models.LoginReq{
			Username:       "yiting",
			HashedPassword: "real_password",
		}
		MockJsonPost(ctx, req)
		s.LoginView(ctx)
		var got models.LoginResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusOK, w.Code)
		//assertEqual(t, ErrInvalidLogin, got, "login_user_wrong_pwd_error")
	})
}

// Below are utility functions

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

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

// MockJsonGet mock the get request wrapped in gin context
func MockJsonGet(c *gin.Context, params gin.Params, u url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)

	// set path params
	c.Params = params

	// set query params
	c.Request.URL.RawQuery = u.Encode()
}

func assertEqual(t *testing.T, want, got interface{}, msg string) {
	t.Helper()

	if !assert.EqualValues(t, want, got) {
		t.Fatalf("%s: want %v; got %v", msg, want, got)
	}
}
