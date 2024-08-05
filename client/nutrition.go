// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/galeone/fitbit/v2/types"
)

// /1/user/[user-id]/foods/log/date/[date].json
func (c *Client) FoodLogList(date time.Time) (ret *types.FoodLogList, err error) {
	path := fmt.Sprintf("/foods/log/date/%s.json", date.Format(types.DateLayout))

	var res *http.Response
	if res, err = c.req.Get(UserV1(path)); err != nil {
		return
	}

	err = json.NewDecoder(res.Body).Decode(&ret)
	return
}
