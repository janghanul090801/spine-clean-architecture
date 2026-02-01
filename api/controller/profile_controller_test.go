package controller_test

//
//import (
//	"encoding/json"
//	"errors"
//	"github.com/NARUBROWN/spine"
//	"github.com/NARUBROWN/spine/core"
//	"github.com/NARUBROWN/spine/pkg/route"
//	"github.com/gofiber/fiber/v2"
//	"github.com/janghanul090801/spine-clean-architecture/api/controller"
//	"github.com/janghanul090801/spine-clean-architecture/domain"
//	"github.com/janghanul090801/spine-clean-architecture/domain/mocks"
//	"github.com/stretchr/testify/require"
//	"io"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//)
//
//type FakeAuthInterceptor struct {
//	UserID string
//}
//
//func (f *FakeAuthInterceptor) PreHandle(
//	ctx core.ExecutionContext,
//	meta core.HandlerMeta,
//) error {
//	ctx.Set("userID", f.UserID)
//	return nil
//}
//
//func (f *FakeAuthInterceptor) PostHandle(
//	ctx core.ExecutionContext,
//	meta core.HandlerMeta,
//) {
//}
//
//func (f *FakeAuthInterceptor) AfterCompletion(
//	ctx core.ExecutionContext,
//	meta core.HandlerMeta,
//	err error,
//) {
//}
//
//func TestFetch(t *testing.T) {
//
//	t.Run("success", func(t *testing.T) {
//		mockProfile := &domain.Profile{
//			Name:  "Test Name",
//			Email: "test@gmail.com",
//		}
//
//		userID := domain.NewID()
//
//		mockProfileUsecase := new(mocks.ProfileUsecase)
//
//		mockProfileUsecase.On("GetProfileByID", mock.Anything, &userID).Return(mockProfile, nil)
//
//		fakeAUthInterceptor := &FakeAuthInterceptor{}
//
//		app := spine.New()
//
//		app.Route("GET", "/profile", (*controller.ProfileController).Fetch, route.WithInterceptors(fakeAUthInterceptor))
//
//		body, err := json.Marshal(mockProfile)
//		assert.NoError(t, err)
//
//		bodyString := string(body)
//
//		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
//		resp, err := app.Test(req)
//		require.NoError(t, err)
//
//		assert.Equal(t, http.StatusOK, resp.StatusCode)
//
//		bodyBytes, err := io.ReadAll(resp.Body)
//		require.NoError(t, err)
//
//		assert.Equal(t, bodyString, string(bodyBytes))
//
//		mockProfileUsecase.AssertExpectations(t)
//	})
//
//	t.Run("error", func(t *testing.T) {
//		userID := domain.NewID()
//
//		mockProfileUsecase := new(mocks.ProfileUsecase)
//
//		customErr := errors.New("unexpected")
//
//		mockProfileUsecase.On("GetProfileByID", mock.Anything, &userID).Return(nil, customErr)
//
//		app := fiber.New()
//
//		pc := controller.NewProfileController(mockProfileUsecase)
//
//		app.Use(setUserID(userID))
//		app.Get("/profile", pc.Fetch)
//
//		body, err := json.Marshal(domain.ErrorResponse{Message: customErr.Error()})
//		assert.NoError(t, err)
//
//		bodyString := string(body)
//
//		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
//		resp, err := app.Test(req)
//		require.NoError(t, err)
//
//		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
//
//		bodyBytes, err := io.ReadAll(resp.Body)
//		require.NoError(t, err)
//
//		assert.Equal(t, bodyString, string(bodyBytes))
//
//		mockProfileUsecase.AssertExpectations(t)
//	})
//
//}
