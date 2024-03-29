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

// UserHeartRateTimeseries retrieves the heart rate time series data over a period of time by specifying a date range.
// The response will include only the daily summary values.
// The endDate parameter is optional. When present it returns the summary, day-by-day, from startDate to endDate.
func (c *Client) UserHeartRateTimeseries(startDate, endDate *time.Time) (ret *types.HeartRateSeries, err error) {
	var res *http.Response
	hasEndDate := endDate != nil && !endDate.IsZero()

	var path string
	// Same route, but with a period of 1d instead of and end date
	if hasEndDate {
		// GET: /1/user/[user-id]/activities/heart/date/[start-date]/[end-date].json
		path = fmt.Sprintf("/activities/heart/date/%s/%s.json", startDate.Format(types.DateLayout), endDate.Format(types.DateLayout))
	} else {
		// GET: /1/user/[user-id]/activities/heart/date/[date]/[period].json
		path = fmt.Sprintf("/activities/heart/date/%s/%s.json", startDate.Format(types.DateLayout), types.Period1Day)
	}
	if res, err = c.req.Get(UserV1(path)); err != nil {
		return
	}
	var body []byte
	if body, err = c.resRead(res); err != nil {
		return
	}

	ret = &types.HeartRateSeries{}
	err = json.Unmarshal(body, ret)
	return

}
