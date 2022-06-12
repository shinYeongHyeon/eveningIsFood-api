package oAuthQueryController

import (
	"github.com/gofiber/fiber/v2"
	createGoogleSignInURLUseCase "github.com/shinYeongHyeon/eveningIsFood-api/src/oauth/application/CreateGoogleSignInURLUseCase"
	coreUseCaseResponseCode "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/useCase"
)

func CreateGoogleSignIn(context *fiber.Ctx) error {
	createGoogleSignInURLUseCaseResponse := createGoogleSignInURLUseCase.Execute()

	if createGoogleSignInURLUseCaseResponse.Code != coreUseCaseResponseCode.SUCCESS {
		return context.Status(fiber.StatusInternalServerError).JSON(createGoogleSignInURLUseCaseResponse)
	}

	return context.Redirect(createGoogleSignInURLUseCaseResponse.Url)
}
