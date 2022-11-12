// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package types

// APIError represents the message sent by the Fitbit API
// when there's an error (in the request, or on the server).
type APIError struct {
	ErrorType string `json:"errorType"`
	FieldName string `json:"fieldName,omitempty"`
	Message   string `json:"message"`
}
