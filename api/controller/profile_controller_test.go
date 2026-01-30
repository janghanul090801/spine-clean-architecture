package controller_test

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/api/controller"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/domain"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/domain/mocks"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setUserID(userID domain.ID) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("id", userID)
		return c.Next()
	}
}

func TestFetch(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockProfile := &domain.Profile{
			Name:  "Test Name",
			Email: "test@gmail.com",
		}

		userID := domain.NewID()

		mockProfileUsecase := new(mocks.ProfileUsecase)

		mockProfileUsecase.On("GetProfileByID", mock.Anything, &userID).Return(mockProfile, nil)

		app := fiber.New()

		pc := controller.NewProfileController(mockProfileUsecase)

		app.Use(setUserID(userID))
		app.Get("/profile", pc.Fetch)

		body, err := json.Marshal(mockProfile)
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		bodyBytes, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		assert.Equal(t, bodyString, string(bodyBytes))

		mockProfileUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		userID := domain.NewID()

		mockProfileUsecase := new(mocks.ProfileUsecase)

		customErr := errors.New("unexpected")

		mockProfileUsecase.On("GetProfileByID", mock.Anything, &userID).Return(nil, customErr)

		app := fiber.New()

		pc := controller.NewProfileController(mockProfileUsecase)

		app.Use(setUserID(userID))
		app.Get("/profile", pc.Fetch)

		body, err := json.Marshal(domain.ErrorResponse{Message: customErr.Error()})
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		resp, err := app.Test(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		bodyBytes, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		assert.Equal(t, bodyString, string(bodyBytes))

		mockProfileUsecase.AssertExpectations(t)
	})

}
