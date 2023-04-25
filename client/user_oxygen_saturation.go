// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/galeone/fitbit/types"
)

// UserOxygenSaturation retrieves the SpO2 summary data for a single date.
// SpO2 applies specifically to a user’s “main sleep”, which is the longest single period of time asleep on a given date.
//
// The endDate parameter is optional. When present it returns the summary, day-by-day, from startDate to endDate.
func (c *Client) UserOxygenSaturation(startDate, endDate *time.Time) (ret *types.OxygenSaturations, err error) {
	var res *http.Response
	var sb strings.Builder

	// /1/user/[user-id]/spo2/date/[date].json
	sb.WriteString(fmt.Sprintf("/spo2/date/%s", startDate.Format(types.DateLayout)))
	if endDate != nil && !endDate.IsZero() {
		// /1/user/[user-id]/spo2/date/[start-date]/[end-date].json
		sb.WriteString(fmt.Sprintf("/%s", endDate.Format(types.DateLayout)))
	}
	sb.WriteString(".json")
	if res, err = c.req.Get(UserV1(sb.String())); err != nil {
		return
	}
	var body []byte
	if body, err = c.resRead(res); err != nil {
		return
	}
	ret = &types.OxygenSaturations{}
	err = json.Unmarshal(body, ret)
	return
}
