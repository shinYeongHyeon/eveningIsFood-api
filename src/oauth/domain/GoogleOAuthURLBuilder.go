package oauthDomain

type GoogleOAuthURL struct {
	scope                string
	accessType           string
	includeGrantedScopes string
	responseType         string
	state                string
	redirectUri          string
	clientId             string
}

// CreateNewGoogleOAuthURL CreateNew returns empty GoogleOAuthURL.
func CreateNewGoogleOAuthURL() *GoogleOAuthURL {
	return &GoogleOAuthURL{}
}

// SetScope sets the scope of the GoogleOAuthURL.
func (builder *GoogleOAuthURL) SetScope(scope string) *GoogleOAuthURL {
	builder.scope = scope

	return builder
}

// SetAccessType sets the access type of the GoogleOAuthURL.
func (builder *GoogleOAuthURL) SetAccessType(accessType string) *GoogleOAuthURL {
	builder.accessType = accessType

	return builder
}

// SetIncludeGrantedScopes sets the include granted scopes of the GoogleOAuthURL.
func (builder *GoogleOAuthURL) SetIncludeGrantedScopes(includeGrantedScopes string) *GoogleOAuthURL {
	builder.includeGrantedScopes = includeGrantedScopes

	return builder
}

// SetResponseType sets the response type of the GoogleOAuthURL.
func (builder *GoogleOAuthURL) SetResponseType(responseType string) *GoogleOAuthURL {
	builder.responseType = responseType

	return builder
}

// SetState sets the state of the GoogleOAuthURL.
func (builder *GoogleOAuthURL) SetState(state string) *GoogleOAuthURL {
	builder.state = state

	return builder
}

// SetRedirectUri sets the redirect uri of the GoogleOAuthURL.
func (builder *GoogleOAuthURL) SetRedirectUri(redirectUri string) *GoogleOAuthURL {
	builder.redirectUri = redirectUri

	return builder
}

// SetClientId sets the client id of the GoogleOAuthURL.
func (builder *GoogleOAuthURL) SetClientId(clientId string) *GoogleOAuthURL {
	builder.clientId = clientId

	return builder
}

// Build returns string from GoogleOAuthURL.
func (builder *GoogleOAuthURL) Build() string {
	return "https://accounts.google.com/o/oauth2/v2/auth?" +
		"scope=" + builder.scope + "&" +
		"access_type=" + builder.accessType + "&" +
		"include_granted_scopes=" + builder.includeGrantedScopes + "&" +
		"response_type=" + builder.responseType + "&" +
		"state=" + builder.state + "&" +
		"redirect_uri=" + builder.redirectUri + "&" +
		"client_id=" + builder.clientId
}
