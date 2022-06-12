package createGoogleSignInURLUseCase

import (
	coreUseCaseResponseCode "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/useCase"
	"testing"
)

func TestExecute(t *testing.T) {
	response := Execute()

	if response.Code != coreUseCaseResponseCode.SUCCESS {
		t.Errorf("Expected code : %s, but got code : %s", coreUseCaseResponseCode.SUCCESS, response.Code)
	}
}
