package createGoogleSignInURLUseCase

import (
	createGoogleSignInURLUseCaseDto "github.com/shinYeongHyeon/eveningIsFood-api/src/oauth/application/CreateGoogleSignInURLUseCase/dto"
	oauthDomain "github.com/shinYeongHyeon/eveningIsFood-api/src/oauth/domain"
	coreUseCaseResponseCode "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/useCase"
)

// Execute : Create google sign in url
func Execute() createGoogleSignInURLUseCaseDto.CreateGoogleSignInURLUseCaseResponse {
	GoogleOAuthURLBuilder := oauthDomain.CreateNewGoogleOAuthURL()

	oauthURL := GoogleOAuthURLBuilder.
		SetScope("https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile").
		SetAccessType("online").
		SetIncludeGrantedScopes("true").
		SetResponseType("code").
		SetState("").
		SetRedirectUri("http://localhost/oauth/google/redirect").
		SetClientId("606338658560-f788riniq0g1fi8bbjp9ni5nojb83b35.apps.googleusercontent.com").
		Build()

	return createGoogleSignInURLUseCaseDto.CreateGoogleSignInURLUseCaseResponse{
		Code: coreUseCaseResponseCode.SUCCESS,
		Url:  oauthURL,
	}
}
