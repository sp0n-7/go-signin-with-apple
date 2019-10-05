package apple

// WebValidationTokenRequest is based off of https://developer.apple.com/documentation/signinwithapplerestapi/generate_and_validate_tokens
type WebValidationTokenRequest struct {
	// ClientID is the "Services ID" value that you get when navigating to your "sign in with Apple"-enabled service ID
	ClientID string

	// ClientSecret is secret generated as a JSON Web Token that uses the secret key generated by the WWDR portal.
	// It can also be generated using the GenerateClientSecret function provided in this package
	ClientSecret string

	// Code is the authorization code received from your application’s user agent.
	// The code is single use only and valid for five minutes.
	Code string

	// RedirectURI is the destination URI the code was originally sent to.
	// Redirect URLs must be registered with Apple. You can register up to 10. Apple will throw an error with IP address
	// URLs on the authorization screen, and will not let you add localhost in the developer portal.
	RedirectURI string
}

// ValidationResponse is based off of https://developer.apple.com/documentation/signinwithapplerestapi/tokenresponse
type ValidationResponse struct {
	// (Reserved for future use) A token used to access allowed data. Currently, no data set has been defined for access.
	AccessToken string `json:"access_token"`

	// The type of access token. It will always be "bearer".
	TokenType string `json:"token_type"`

	// The amount of time, in seconds, before the access token expires. You can revalidate with the "RefreshToken"
	ExpiresIn int `json:"expires_in"`

	// The refresh token used to regenerate new access tokens. Store this token securely on your server.
	RefreshToken string `json:"refresh_token"`

	// A JSON Web Token that contains the user’s identity information.
	IDToken string `json:"id_token"`

	// Used to capture any error returned by the endpoint. Do not trust the response if this error is not nil
	Error string `json:"error"`
}
