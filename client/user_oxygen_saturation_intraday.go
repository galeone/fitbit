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

// UserOxygenSaturationIntraday retrieves s the SpO2 intraday data for a specified date range.
// SpO2 applies specifically to a user’s “main sleep”, which is the longest single period of time asleep on a given date.
// Spo2 values are calculated on a 5-minute exponentially-moving average.
// The measurement is provided at the end of a period of sleep.
//
// The endDate parameter is optional. When present it returns the summary, day-by-day, from startDate to endDate.
func (c *Client) UserOxygenSaturationIntraday(startDate, endDate *time.Time) (ret *types.OxygenSaturationIntraday, err error) {
	var res *http.Response
	var sb strings.Builder

	// /1/user/[user-id]/spo2/date/[date]/all.json
	sb.WriteString(fmt.Sprintf("/spo2/date/%s", startDate.Format(types.DateLayout)))
	if endDate != nil && !endDate.IsZero() {
		// /1/user/[user-id]/spo2/date/[start-date]/[end-date]/all.json
		sb.WriteString(fmt.Sprintf("/%s", endDate.Format(types.DateLayout)))
	}
	sb.WriteString("/all.json")
	if res, err = c.req.Get(UserV1(sb.String())); err != nil {
		return
	}

	var body []byte
	if body, err = c.resRead(res); err != nil {
		return
	}

	ret = &types.OxygenSaturationIntraday{}
	err = json.Unmarshal(body, ret)
	return
}
