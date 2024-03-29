// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package api contains the implementation of the REST
// [Fitbit Web API][1].
//
// [1] https://dev.fitbit.com/build/reference/web-api/explore/
package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/galeone/fitbit/v2"
)

const (
	apiV1  string = "https://api.fitbit.com/1"
	apiV12 string = "https://api.fitbit.com/1.2"
)

// V1 returns the Fitbit API v1 pointing to the desired endpoint.
func V1(endpoint string) string {
	return fmt.Sprintf("%s/%s", apiV1, strings.TrimLeft(endpoint, "/"))
}

// UserV1 returns the Fitbit API v1 pointing to the desired endpoint.
// The UserV1 call differs from V1, because it passes the `/user/-/` arguments
// after the V1 base url.
// This endpoint shall be used when requiring user-specific info.
func UserV1(endpoint string) string {
	return fmt.Sprintf("%s/user/-/%s", apiV1, strings.TrimLeft(endpoint, "/"))
}

// UserV12 is the same of UserV1, with the only difference that it points to the
// Fitbit API v1.2 instead of v1.
// This endpoint is used for the sleep data requests.
func UserV12(endpoint string) string {
	return fmt.Sprintf("%s/user/-/%s", apiV12, strings.TrimLeft(endpoint, "/"))
}

// Client is the implementation of the [Fitbit Web API][1].
// [1] https://dev.fitbit.com/build/reference/web-api/explore/
type Client struct {
	authorizer *fitbit.Authorizer
	req        *http.Client
}

// NewClient creates a new *Client
func NewClient(authorizer *fitbit.Authorizer) (ret *Client, err error) {
	ret = &Client{
		authorizer,
		nil,
	}
	if err = ret.Req(); err != nil {
		return nil, err
	}
	return
}

// Req refreshes the HTTP client. It uses the Authorizer instance
// that automatically handles the refresh token exchange when needed.
//
// Call this method when the various Client methods are failing because of
// the expired access token.
func (c *Client) Req() (err error) {
	var req *http.Client
	if req, err = c.authorizer.HTTP(); err == nil {
		c.req = req
	}
	return
}

func (c *Client) resRead(res *http.Response) (body []byte, err error) {
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode: %d. Message: %s", res.StatusCode, string(body))
	}
	return

}
