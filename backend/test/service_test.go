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
	sampleSessionToken = "64ed7eb3-826e-4f20-86dd-7a6bb4f85026" //string //
	sampleUserID       = int64(2)                               //int64  //
	sampleUserName     = "yiting"
	sampleUserEmail    = "yiting@travelmaster.com"
	dummySessionToken  = "84db0f70-b3c6-4e18-be80-e68a073f0a49" //string

	// sample activity
	sampleActivityID = int64(1) //int64

	// sample itinerary
	generateItiId = int64(1) //int64

	// sample review
	sampleReviewID = int64(1) //int64 //

	// error response for sign_up endpoint
	ErrMissingUserInfo   = errors.New("either email, username, or hashed password missing")
	ErrUserAlreadyExists = errors.New("user already exists")

	// error response for login request
	ErrInvalidLogin = errors.New("invalid login")

	// error response for create activity
	ErrNullTitle             = errors.New("title cannot be empty")
	ErrActivityAlreadyExists = errors.New("an activity with the same title already exists")
	ErrUserNotFound          = errors.New("user id doesn't exists")

	// error response for get activity
	ErrActivityNotFound = errors.New("activity not found")

	// error response for create review
	ErrUserAlreadyCreatedReview = errors.New("user already created review for the activity")
	ErrNotAllowed               = errors.New("user is not allowed to perform this action")

	// error response for modify review
	ErrReviewNotFound  = errors.New("review not found")
	ErrAlreadyReported = errors.New("user has already reported the activity")

	// error response for undo report
	ErrReportNotFound = errors.New("report not found")
	ErrNoSearchFail   = errors.New("search Name failed")
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

		MockPostForm(ctx, req, []string{"test.png"}, "avatar")

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

		MockPostForm(ctx, req, []string{"test.png"}, "avatar")

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

		MockPostForm(ctx, req, []string{"test.png"}, "avatar")

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

		MockPostForm(ctx, req, []string{"test.txt"}, "avatar")

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
		dummySessionToken = got.SessionToken
	})
	t.Run("signup_success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.SignupForm{
			Username:       sampleUserName,
			Email:          sampleUserEmail,
			HashedPassword: "real_password",
		}

		MockPostForm(ctx, req, []string{"test.png"}, "avatar")

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

		MockPostForm(ctx, req, []string{"test.png"}, "avatar")

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
		fmt.Println(got)
		assertEqual(t, want, canCompare, "login_success")
		sampleSessionToken = got.SessionToken
	})
}

func TestCreateActivity(t *testing.T) {

	t.Run("create_activity_missing_title", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.CreateActivityForm{
			UserId:          sampleUserID,
			Title:           "",
			Paid:            false,
			Categories:      []string{"Outdoor", "Nature"},
			Description:     "Quite Interesting",
			Longitude:       88.1,
			Latitude:        34.2,
			MonOpeningTime:  8,
			TueOpeningTime:  8,
			WedOpeningTime:  8,
			ThurOpeningTime: 8,
			FriOpeningTime:  8,
			SatOpeningTime:  8,
			SunOpeningTime:  8,
			MonClosingTime:  22,
			TueClosingTime:  22,
			WedClosingTime:  22,
			ThurClosingTime: 22,
			FriClosingTime:  22,
			SatClosingTime:  22,
			SunClosingTime:  22,
		}

		MockPostForm(ctx, req, []string{"test.png", "test1.png"}, "image")

		// send request to endpoint
		s.CreateActivity(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrNullTitle.Error(), got.Error, "error")
	})
	t.Run("create_activity_invalid_user", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.CreateActivityForm{
			UserId:          -1,
			Title:           "",
			Paid:            false,
			Categories:      []string{"Outdoor", "Nature"},
			Description:     "Quite Interesting",
			Longitude:       88.1,
			Latitude:        34.2,
			MonOpeningTime:  8,
			TueOpeningTime:  8,
			WedOpeningTime:  8,
			ThurOpeningTime: 8,
			FriOpeningTime:  8,
			SatOpeningTime:  8,
			SunOpeningTime:  8,
			MonClosingTime:  22,
			TueClosingTime:  22,
			WedClosingTime:  22,
			ThurClosingTime: 22,
			FriClosingTime:  22,
			SatClosingTime:  22,
			SunClosingTime:  22,
		}

		MockPostForm(ctx, req, []string{"test.png"}, "image")

		// send request to endpoint
		s.CreateActivity(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrUserNotFound.Error(), got.Error, "error")
	})
	t.Run("create_activity_success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.CreateActivityForm{
			UserId:          sampleUserID,
			Title:           "Singapore Night Safari",
			Paid:            false,
			Categories:      []string{"Outdoor", "Nature"},
			Description:     "Quite Interesting",
			Longitude:       88.1,
			Latitude:        34.2,
			MonOpeningTime:  8,
			TueOpeningTime:  8,
			WedOpeningTime:  8,
			ThurOpeningTime: 8,
			FriOpeningTime:  8,
			SatOpeningTime:  8,
			SunOpeningTime:  8,
			MonClosingTime:  22,
			TueClosingTime:  22,
			WedClosingTime:  22,
			ThurClosingTime: 22,
			FriClosingTime:  22,
			SatClosingTime:  22,
			SunClosingTime:  22,
		}

		MockPostForm(ctx, req, []string{"test.png"}, "image")

		// send request to endpoint
		s.CreateActivity(ctx)

		var got models.CreateActivityResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusCreated, w.Code)
		assert.NotEmpty(t, got.ActivityId)
		sampleActivityID = got.ActivityId
		fmt.Println(got)
	})
	t.Run("create_activity_duplicate_title", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.CreateActivityForm{
			UserId:          sampleUserID,
			Title:           "Singapore Night Safari",
			Rating:          5,
			Paid:            false,
			Categories:      []string{"Outdoor", "Nature"},
			Description:     "Quite Interesting",
			Longitude:       88.1,
			Latitude:        34.2,
			MonOpeningTime:  8,
			TueOpeningTime:  8,
			WedOpeningTime:  8,
			ThurOpeningTime: 8,
			FriOpeningTime:  8,
			SatOpeningTime:  8,
			SunOpeningTime:  8,
			MonClosingTime:  22,
			TueClosingTime:  22,
			WedClosingTime:  22,
			ThurClosingTime: 22,
			FriClosingTime:  22,
			SatClosingTime:  22,
			SunClosingTime:  22,
		}

		MockPostForm(ctx, req, []string{"test.png"}, "image")

		// send request to endpoint
		s.CreateActivity(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrActivityAlreadyExists.Error(), got.Error, "error")
	})
}

func TestGetActivity(t *testing.T) {

	t.Run("get_activity_non_existing_id", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.GetActivityReq{
			ActivityId: -1,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.GetActivity(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrActivityNotFound.Error(), got.Error, "error")
	})
	t.Run("get_activity_success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.GetActivityReq{
			ActivityId: int(sampleActivityID),
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.GetActivity(ctx)

		var got models.GetActivityResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusOK, w.Code)
		assertEqual(t, got.ActivityId, sampleActivityID, "error")
	})
}

func TestCreateReview(t *testing.T) {
	t.Run("create_review_invalid_session_token", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.AddReviewReq{
			SessionToken: "",
			ActivityId:   sampleActivityID,
			Title:        "I like it here!",
			Description:  "Just has great time with m y family and friends!",
			Rating:       5.0,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.AddReview(ctx)

		var msg string
		if err := json.Unmarshal(w.Body.Bytes(), &msg); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusInternalServerError, w.Code)
		assertEqual(t, ErrNotAllowed.Error(), msg, "error")
	})
	t.Run("create_review_invalid_activity", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.AddReviewReq{
			SessionToken: sampleSessionToken,
			ActivityId:   -1,
			Title:        "I like it here!",
			Description:  "Just has great time with m y family and friends!",
			Rating:       5.0,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.AddReview(ctx)

		var msg string
		if err := json.Unmarshal(w.Body.Bytes(), &msg); err != nil {
			log.Fatalln(err)
		}

		assert.EqualValues(t, http.StatusInternalServerError, w.Code)
		assertEqual(t, ErrActivityNotFound.Error(), msg, "error")
	})
	t.Run("create_review_success_rating_5", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.AddReviewReq{
			SessionToken: sampleSessionToken,
			ActivityId:   sampleActivityID,
			Title:        "I like it here!",
			Description:  "Just has great time with m y family and friends!",
			Rating:       5.0,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.AddReview(ctx)

		var got models.GetActivityResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		canCompare := models.GetActivityResp{
			ActivityId:   got.ActivityId,
			Rating:       got.Rating,
			ReviewCounts: got.ReviewCounts,
		}
		want := models.GetActivityResp{
			ActivityId:   sampleActivityID,
			Rating:       5.0,
			ReviewCounts: 1,
		}
		assert.EqualValues(t, http.StatusCreated, w.Code)
		assertEqual(t, want, canCompare, "error")
	})
	t.Run("create_review_success_rating_4", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.AddReviewReq{
			SessionToken: dummySessionToken,
			ActivityId:   sampleActivityID,
			Title:        "Nice Place!",
			Description:  "Just has great time with m y family and friends!",
			Rating:       4.0,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.AddReview(ctx)

		var got models.GetActivityResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		canCompare := models.GetActivityResp{
			ActivityId:   got.ActivityId,
			Rating:       got.Rating,
			ReviewCounts: got.ReviewCounts,
		}
		want := models.GetActivityResp{
			ActivityId:   int64(1),
			Rating:       4.5,
			ReviewCounts: 2,
		}
		assert.EqualValues(t, http.StatusCreated, w.Code)
		assertEqual(t, want, canCompare, "error")
	})
	t.Run("create_review_duplicate", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.AddReviewReq{
			SessionToken: sampleSessionToken,
			ActivityId:   sampleActivityID,
			Title:        "I like it here!",
			Description:  "Just had great time with my family and friends!",
			Rating:       5.0,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.AddReview(ctx)

		var msg string
		if err := json.Unmarshal(w.Body.Bytes(), &msg); err != nil {
			log.Fatalln(err)
		}

		assert.EqualValues(t, http.StatusMethodNotAllowed, w.Code)
		assertEqual(t, ErrUserAlreadyCreatedReview.Error(), msg, "error")
	})
}

func TestModifyReview(t *testing.T) {
	// if ReviewID, ActivityID and UserID are not a pair in database, return the error
	t.Run("modify_review_unfounded", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.UpdateReviewReq{
			ReviewId:    -1,
			ActivityId:  sampleActivityID,
			UserId:      sampleUserID,
			Delete:      false,
			Title:       "Really great!",
			Description: "Not bad",
			NewRating:   3.0,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.UpdateReview(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrReviewNotFound.Error(), got.Error, "error")
	})
	t.Run("modify_activity_unfounded", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.UpdateReviewReq{
			ReviewId:    sampleReviewID,
			ActivityId:  -1,
			UserId:      sampleUserID,
			Delete:      false,
			Title:       "Really great!",
			Description: "Not bad",
			NewRating:   3.0,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.UpdateReview(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrReviewNotFound.Error(), got.Error, "error")
	})
	t.Run("modify_review_success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.UpdateReviewReq{
			ReviewId:    sampleReviewID,
			ActivityId:  sampleActivityID,
			UserId:      sampleUserID,
			Delete:      false,
			Title:       "Really great!",
			Description: "Not bad",
			NewRating:   3.0,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.UpdateReview(ctx)

		var got models.GetActivityResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusAccepted, w.Code)
		assertEqual(t, 3.5, got.Rating, "error")
	})
	t.Run("delete_review_success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.UpdateReviewReq{
			ReviewId:    sampleReviewID,
			ActivityId:  sampleActivityID,
			UserId:      sampleUserID,
			Delete:      true,
			Title:       "Really great!",
			Description: "Not bad",
			NewRating:   0.0,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.UpdateReview(ctx)

		var got models.GetActivityResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		canCompare := models.GetActivityResp{
			ActivityId:   got.ActivityId,
			Rating:       got.Rating,
			ReviewCounts: got.ReviewCounts,
		}
		want := models.GetActivityResp{
			ActivityId:   sampleActivityID,
			Rating:       4.0,
			ReviewCounts: 1,
		}
		assert.EqualValues(t, http.StatusAccepted, w.Code)
		assertEqual(t, want, canCompare, "error")
	})

}

func TestReportActivity(t *testing.T) {
	t.Run("report_unfounded_activity", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.IncrementInactiveCountReq{
			ActivityId: -1,
			UserId:     sampleUserID,
			Reason:     "Don't like",
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.IncrementInactiveCount(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrActivityNotFound.Error(), got.Error, "error")
	})
	t.Run("report_unfounded_user", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.IncrementInactiveCountReq{
			ActivityId: sampleActivityID,
			UserId:     -1,
			Reason:     "Don't like",
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.IncrementInactiveCount(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrUserNotFound.Error(), got.Error, "error")
	})
	t.Run("report_activity_success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.IncrementInactiveCountReq{
			ActivityId: sampleActivityID,
			UserId:     sampleUserID,
			Reason:     "Don't like",
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.IncrementInactiveCount(ctx)

		var got models.ChangeInactiveCountResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		canCompare := models.ChangeInactiveCountResp{
			ActivityId:    got.ActivityId,
			InactiveCount: got.InactiveCount,
			InactiveFlag:  got.InactiveFlag,
		}
		want := models.ChangeInactiveCountResp{
			ActivityId:    sampleActivityID,
			InactiveCount: 1,
			InactiveFlag:  false,
		}
		assert.EqualValues(t, http.StatusAccepted, w.Code)
		assertEqual(t, want, canCompare, "error")
	})
	t.Run("already_reported", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.IncrementInactiveCountReq{
			ActivityId: sampleActivityID,
			UserId:     sampleUserID,
			Reason:     "Don't like",
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.IncrementInactiveCount(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrAlreadyReported.Error(), got.Error, "error")
	})
}

func TestUndoReportActivity(t *testing.T) {

	t.Run("undo_unfounded_activity", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.DecrementInactiveCountReq{
			ActivityId: -1,
			UserId:     sampleUserID,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.DecrementInactiveCount(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrActivityNotFound.Error(), got.Error, "error")
	})
	t.Run("undo_unfounded_user", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.DecrementInactiveCountReq{
			ActivityId: sampleActivityID,
			UserId:     -1,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.DecrementInactiveCount(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrUserNotFound.Error(), got.Error, "error")
	})
	t.Run("undo_report_activity_success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.DecrementInactiveCountReq{
			ActivityId: sampleActivityID,
			UserId:     sampleUserID,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.DecrementInactiveCount(ctx)

		var got models.ChangeInactiveCountResp
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		canCompare := models.ChangeInactiveCountResp{
			ActivityId:    got.ActivityId,
			InactiveCount: got.InactiveCount,
			InactiveFlag:  got.InactiveFlag,
		}
		want := models.ChangeInactiveCountResp{
			ActivityId:    sampleActivityID,
			InactiveCount: 0,
			InactiveFlag:  false,
		}
		assert.EqualValues(t, http.StatusAccepted, w.Code)
		assertEqual(t, want, canCompare, "error")
	})
	t.Run("undo_report_unfound_report", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)

		req := models.DecrementInactiveCountReq{
			ActivityId: sampleActivityID,
			UserId:     sampleUserID,
		}

		MockJsonPost(ctx, req)

		// send request to endpoint
		s.DecrementInactiveCount(ctx)

		var got errorMessage
		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			log.Fatalln(err)
		}
		assert.EqualValues(t, http.StatusBadRequest, w.Code)
		assertEqual(t, ErrReportNotFound.Error(), got.Error, "error")
	})
}

func TestGenerateItinerary(t *testing.T) {
	// we are not concerned with the quality of the actual itinerary generated
	t.Run("successfully generate itinerary", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := GetTestGinContext(w)
		req := models.GenerateItineraryRequest{
			SessionToken:        sampleSessionToken,
			PreferredCategories: []string{},
			StartTime:           100,
			EndTime:             1000,
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
			SessionToken:        sampleSessionToken,
			PreferredCategories: []string{},
			StartTime:           10000,
			EndTime:             1000,
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
			SessionToken:        "bad token",
			PreferredCategories: []string{},
			StartTime:           100,
			EndTime:             1000,
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
			Id:           generateItiId,
			Name:         updatedName,
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
			Id:           1000,
			Name:         "new name",
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
			Id:           generateItiId,
			Name:         "new name",
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
			Id:           generateItiId,
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
			Id:           1000,
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
			Id:           generateItiId,
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

func MockPostForm(ctx *gin.Context, content interface{}, imageFile []string, imageFieldName string) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writeForm(writer, content, imageFile, imageFieldName)
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

func writeForm(writer *multipart.Writer, body interface{}, fileName []string, fieldName string) {

	val := reflect.ValueOf(body)
	for i := 0; i < val.Type().NumField(); i++ {
		var fieldVal string
		if val.Field(i).Type().Kind() != reflect.String {
			a, _ := json.Marshal(val.Field(i).Interface())
			fieldVal = string(a)
		} else {
			fieldVal = val.Field(i).Interface().(string)
		}
		if err := writer.WriteField(val.Type().Field(i).Tag.Get("form"), fieldVal); err != nil {
			log.Fatalln(err)
		}
		//fmt.Println(val.Type().Field(i).Tag.Get("form"))
		//fmt.Println(fieldVal)
	}

	for i := 0; i < len(fileName); i++ {
		file, err := os.Open(fileName[i])
		if err != nil {
			log.Fatalln(err)
		}
		fileWriter, err := writer.CreateFormFile(fieldName, fileName[i])
		if err != nil {
			log.Fatalln(err)
		}
		if _, err = io.Copy(fileWriter, file); err != nil {
			log.Fatalln(err)
		}
	}
}
