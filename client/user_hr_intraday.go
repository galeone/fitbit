// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/galeone/fitbit/types"
)

// UserHeartRateIntraday retrieves the heart rate intraday time series data on a specific date range for a 24 hour period.
// The response will include the activity detail second by second.
// The endDate parameter is optional. When present it should be in a 24 hours range between the start date.
// Minutes are considered.
func (c *Client) UserHeartRateIntraday(startDate, endDate *time.Time) (ret *types.HeartRateSeries, err error) {
	var res *http.Response
	hasEndDate := endDate != nil && !endDate.IsZero()

	var path string
	// Same route, but with a period of 1d instead of and end date
	if hasEndDate {
		// /1/user/[user-id]/activities/heart/date/[date]/1d/[detail-level]/time/[start-time]/[end-time].json
		path = fmt.Sprintf("/activities/heart/date/%s/1d/1sec/time/%s/%s.json", startDate.Format(types.DateLayout), startDate.Format(types.TimeLayout), endDate.Format(types.TimeLayout))
	} else {
		// /1/user/[user-id]/activities/heart/date/[date]/1d/[detail-level].json
		path = fmt.Sprintf("/activities/heart/date/%s/%s/1sec.json", startDate.Format(types.DateLayout), types.Period1Day)
	}
	if res, err = c.req.Get(UserV1(path)); err != nil {
		return
	}
	var body []byte
	if body, err = c.resRead(res); err != nil {
		return
	}
	err = json.Unmarshal(body, ret)
	return

}
