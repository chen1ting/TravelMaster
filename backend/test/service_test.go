package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/chen1ting/TravelMaster/internal/models"
	"github.com/chen1ting/TravelMaster/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type errorMessage struct {
	Error string `json:"error"`
}

var (
	s = service.NewService("testing")
	// sample user credentials
	sampleSessionToken string
	sampleUserID       int64
	sampleUserName     = "yiting"
	sampleUserEmail    = "yiting@travelmaster.com"

	// sample activity

	// error response for sign_up endpoint
	ErrMissingUserInfo   = errors.New("either email, username, or hashed password missing")
	ErrUserAlreadyExists = errors.New("user already exists")

	// error response for login request
	ErrInvalidLogin = errors.New("invalid login")

	// error response for create activity
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

		req := models.SignupForm{
			Username:       sampleUserName,
			Email:          "",
			HashedPassword: "real_password",
		}

		MockJsonPostForm(ctx, req, "test.png")

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

		req := models.SignupForm{
			Username:       "",
			Email:          sampleUserEmail,
			HashedPassword: "real_password",
		}

		MockJsonPostForm(ctx, req, "test.png")

		// send request to endpoint
		s.SignupView(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrMissingUserInfo.Error(), got.Error, "signup_missing_username")
	})

	t.Run("signup_missing_hashed_password", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.SignupForm{
			Username:       sampleUserName,
			Email:          sampleUserEmail,
			HashedPassword: "",
		}

		MockJsonPostForm(ctx, req, "test.png")

		// send request to endpoint
		s.SignupView(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrMissingUserInfo.Error(), got.Error, "signup_missing_email")
	})

	t.Run("signup_wrong_filetype", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.SignupForm{
			Username:       "yiting_no_avatar",
			Email:          "yiting_no_avatar@travelmaster.com",
			HashedPassword: "real_password",
		}

		MockJsonPostForm(ctx, req, "test.txt")

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

		req := models.SignupForm{
			Username:       sampleUserName,
			Email:          sampleUserEmail,
			HashedPassword: "real_password",
		}

		MockJsonPostForm(ctx, req, "test.png")

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
			Username: sampleUserName,
			Email:    sampleUserEmail,
		}
		sampleSessionToken = got.SessionToken
		sampleUserID = got.UserId
		assert.EqualValues(t, http.StatusCreated, w.Code)
		assert.NotEmpty(t, got.AvatarName)
		assertEqual(t, want, canCompare, "signup_success")
	})

	t.Run("signup_user_existed", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.SignupForm{
			Username:       sampleUserName,
			Email:          sampleUserEmail,
			HashedPassword: "real_password",
		}

		MockJsonPostForm(ctx, req, "test.png")

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

func TestValidateToken(t *testing.T) {
	t.Run("invalid_token", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := models.ValidateTokenReq{
			SessionToken: "",
		}
		MockJsonPost(ctx, req)
		s.ValidateToken(ctx)
		var got models.ValidateTokenResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		want := models.ValidateTokenResp{Valid: false, UserId: -1}
		assert.EqualValues(t, http.StatusOK, w.Code)
		assertEqual(t, want, got, "test_invalid_token")
	})
	t.Run("valid_token", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := models.ValidateTokenReq{
			SessionToken: sampleSessionToken,
		}
		MockJsonPost(ctx, req)
		s.ValidateToken(ctx)
		var got models.ValidateTokenResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		want := models.ValidateTokenResp{Valid: true, UserId: sampleUserID}
		assert.EqualValues(t, http.StatusOK, w.Code)
		assertEqual(t, want, got, "test_valid_token")
	})
}

func TestLogout(t *testing.T) {
	t.Run("invalid_user", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := models.LogoutReq{
			SessionToken: "",
		}
		MockJsonPost(ctx, req)
		s.LogoutView(ctx)
		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
	})

	// notice here we have invalidated the previous session token
	t.Run("valid_user", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := models.LogoutReq{
			SessionToken: sampleSessionToken,
		}
		MockJsonPost(ctx, req)
		s.LogoutView(ctx)
		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusOK, w.Code)
	})

	// see if after logout, will sending the request with same token still be valid
	t.Run("multiple logouts", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := models.LogoutReq{
			SessionToken: sampleSessionToken,
		}
		MockJsonPost(ctx, req)
		s.LogoutView(ctx)
		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
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
			Username:       sampleUserName,
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
			Username:       sampleUserName,
			HashedPassword: "real_password",
		}
		MockJsonPost(ctx, req)
		s.LoginView(ctx)
		var got models.LoginResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusOK, w.Code)
		canCompare := models.LoginResp{
			UserId:   got.UserId,
			Username: got.Username,
			Email:    got.Email,
		}
		want := models.LoginResp{
			UserId:   sampleUserID,
			Username: sampleUserName,
			Email:    sampleUserEmail,
		}
		assertEqual(t, canCompare, want, "login_success")
		sampleSessionToken = got.SessionToken
	})
}

func TestCreateActivity(t *testing.T) {

}

var generateItiId int64

func TestGenerateItinerary(t *testing.T) {
	// we are not concerned with the quality of the actual itinerary generated
	t.Run("successfully generate itinerary", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := models.GenerateItineraryRequest{
			SessionToken: sampleSessionToken,
			PreferredCategories: []string{},
			StartTime: 100,
			EndTime: 1000,
		}
		MockJsonPost(ctx, req)
		s.GenerateItinerary(ctx)
		assert.EqualValues(t, http.StatusOK, w.Code)

		var got models.GenerateItineraryResponse
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		generateItiId = got.GeneratedItinerary.Id
	})

	t.Run("bad itinerary start and end times", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := models.GenerateItineraryRequest{
			SessionToken: sampleSessionToken,
			PreferredCategories: []string{},
			StartTime: 10000,
			EndTime: 1000,
		}
		MockJsonPost(ctx, req)
		s.GenerateItinerary(ctx)
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
	})

	// we are not concerned with the quality of the actual itinerary generated
	t.Run("unrecognised session token", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := models.GenerateItineraryRequest{
			SessionToken: "bad token",
			PreferredCategories: []string{},
			StartTime: 100,
			EndTime: 1000,
		}
		MockJsonPost(ctx, req)
		s.GenerateItinerary(ctx)
		assert.EqualValues(t, http.StatusInternalServerError, w.Code)
	})
}

const updatedName = "new name"

func TestUpdateItinerary(t *testing.T) {
	// we are not concerned with the quality of the actual itinerary generated
	t.Run("successfully update existing itinerary", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := &models.SaveItineraryRequest{
			SessionToken: sampleSessionToken,
			Id: generateItiId,
			Name: updatedName,
		}
		MockJsonPost(ctx, req)
		s.UpdateItinerary(ctx)
		assert.EqualValues(t, http.StatusOK, w.Code)
		var got models.SaveItineraryResponse
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, got.Id, generateItiId)
		assert.EqualValues(t, got.Name, req.Name)
	})

	t.Run("fail update non existing itinerary", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := &models.SaveItineraryRequest{
			SessionToken: sampleSessionToken,
			Id: 1000,
			Name: "new name",
		}
		MockJsonPost(ctx, req)
		s.UpdateItinerary(ctx)
		assert.EqualValues(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("fail bad token", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := &models.SaveItineraryRequest{
			SessionToken: "bad token",
			Id: generateItiId,
			Name: "new name",
		}
		MockJsonPost(ctx, req)
		s.UpdateItinerary(ctx)
		assert.EqualValues(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetItinerary(t *testing.T) {
	// we are not concerned with the quality of the actual itinerary generated
	t.Run("successfully get existing itinerary", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := &models.GetItineraryRequest{
			SessionToken: sampleSessionToken,
			Id: generateItiId,
		}
		MockJsonPost(ctx, req)
		s.GetItinerary(ctx)
		assert.EqualValues(t, http.StatusOK, w.Code)
		var got models.GetItineraryResponse
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, generateItiId, got.Itinerary.Id)
		assert.EqualValues(t, updatedName, got.Itinerary.Name)
	})

	t.Run("fail get non existing itinerary", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := &models.GetItineraryRequest{
			SessionToken: sampleSessionToken,
			Id: 1000,
		}
		MockJsonPost(ctx, req)
		s.GetItinerary(ctx)
		assert.EqualValues(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("fail bad token", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := &models.GetItineraryRequest{
			SessionToken: "bad token",
			Id: generateItiId,
		}
		MockJsonPost(ctx, req)
		s.GetItinerary(ctx)
		assert.EqualValues(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetItineraries(t *testing.T) {
	// we are not concerned with the quality of the actual itinerary generated
	t.Run("successfully get existing itineraries", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := &models.GetItinerariesRequest{
			SessionToken: sampleSessionToken,
		}
		MockJsonPost(ctx, req)
		s.GetItineraries(ctx)
		assert.EqualValues(t, http.StatusOK, w.Code)
		var got models.GetItinerariesResponse
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, 1, len(got.Itineraries))
		assert.EqualValues(t, generateItiId, got.Itineraries[0].Id)
		assert.EqualValues(t, updatedName, got.Itineraries[0].Name)
	})

	t.Run("fail bad token", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := &models.GetItinerariesRequest{
			SessionToken: "bad token",
		}
		MockJsonPost(ctx, req)
		s.GetItineraries(ctx)
		assert.EqualValues(t, http.StatusInternalServerError, w.Code)
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

func MockJsonPost(ctx *gin.Context, content interface{}) {
	ctx.Request.Method = "POST"
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Set("user_id", 1)

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// wrap it in a no-op closer
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func MockJsonPostForm(ctx *gin.Context, content interface{}, imageFile string) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writeForm(writer, content, imageFile)
	if err := writer.Close(); err != nil {
		log.Fatalln(err)
	}
	ctx.Request.Body = io.NopCloser(body)

	// write form header
	ctx.Request.Method = "POST"
	ctx.Request.Header.Set("Content-Type", writer.FormDataContentType())

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

func writeForm(writer *multipart.Writer, body interface{}, fileName string) {

	val := reflect.ValueOf(body)
	for i := 0; i < val.Type().NumField(); i++ {
		if err := writer.WriteField(val.Type().Field(i).Tag.Get("form"), val.Field(i).String()); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(val.Type().Field(i).Tag.Get("json"))
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	fileWriter, err := writer.CreateFormFile("avatar", fileName)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = io.Copy(fileWriter, file); err != nil {
		log.Fatalln(err)
	}
}
