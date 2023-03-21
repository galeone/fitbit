// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package types

// AuthorizedUser represents the payload received
// after succesfully exchaing the Authorization Code
// with the Fitbit OAuth2 server (Server Application Type)
// See the [documentation] - step 4.
//
// [documentation]: https://dev.fitbit.com/build/reference/web-api/developer-guide/authorization/#Authorization-Code-Grant-Flow-with-PKCE
type AuthorizedUser struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	UserID       string `json:"user_id"`
}

// OAuth2Error represents the payload received in case of error, during the
// OAuth2 authorization flow
type OAuth2Error struct {
	Errors  []APIError
	Success bool `json:"success"`
}

// AuthorizingUser is the type used during the exchange of the
// "Code" for the tokens in the OAuth2 flow.
// There's no JSON decoration because the code is taken from the URL.
// The same goes for the CSRFToken that's placed inside the "state" URL
// parameter.
type AuthorizingUser struct {
	Code      string
	CSRFToken string
}
