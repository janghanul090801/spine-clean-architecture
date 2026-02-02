package controller

import (
	"context"
	"fmt"
	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/NARUBROWN/spine/pkg/spine"
	"github.com/janghanul090801/spine-clean-architecture/domain"
	"net/http"
)

type ProfileController struct {
	profileUsecase domain.ProfileUsecase
}

func NewProfileController(usecase domain.ProfileUsecase) *ProfileController {
	return &ProfileController{
		profileUsecase: usecase,
	}
}

func (pc *ProfileController) Fetch(ctx context.Context, spineCtx spine.Ctx) httpx.Response[domain.Profile] {
	v, ok := spineCtx.Get("id")
	if !ok {
		return httpx.Response[domain.Profile]{
			Options: httpx.ResponseOptions{
				Status: http.StatusUnauthorized, // unauthorized
			},
		}
	}

	userID := v.(*domain.ID)

	profile, err := pc.profileUsecase.GetProfileByID(ctx, userID)
	if err != nil {
		fmt.Println(err.Error())
		return httpx.Response[domain.Profile]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	return httpx.Response[domain.Profile]{
		Body: *profile,
	}
}
