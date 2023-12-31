// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package client

import (
	"encoding/json"
	"net/http"

	"github.com/galeone/fitbit/v2/types"
)

// AllActivityTypes a list of all valid Fitbit public activities and the private, user-created activities from the Fitbit activities database.
// If available, activity level details will display.
//
// GET: /1/activities.json
func (c *Client) AllActivityTypes() (ret *types.ActivityCatalog, err error) {
	var res *http.Response
	if res, err = c.req.Get(V1("/activities.json")); err != nil {
		return
	}
	var body []byte
	if body, err = c.resRead(res); err != nil {
		return
	}
	ret = &types.ActivityCatalog{}
	err = json.Unmarshal(body, ret)
	return
}
