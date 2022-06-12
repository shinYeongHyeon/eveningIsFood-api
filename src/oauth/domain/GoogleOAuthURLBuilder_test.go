package oauthDomain

import "testing"

func TestNewGoogleOAuthURLBuilder(t *testing.T) {
	builder := CreateNewGoogleOAuthURL()

	string := builder.
		SetScope("test_scope").
		SetAccessType("test_access_type").
		SetIncludeGrantedScopes("test_include_granted_scopes").
		SetResponseType("test_response_type").
		SetState("test_state").
		SetRedirectUri("test_redirect_uri").
		SetClientId("test_client_id").
		Build()

	if string != "https://accounts.google.com/o/oauth2/v2/auth?scope=test_scope&access_type=test_access_type&include_granted_scopes=test_include_granted_scopes&response_type=test_response_type&state=test_state&redirect_uri=test_redirect_uri&client_id=test_client_id" {
		t.Errorf("CreateNewGoogleOAuthURL() should return valid Google OAuth URL, but got: %s", string)
	}
}
