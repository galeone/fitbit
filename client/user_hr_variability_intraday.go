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

	"github.com/galeone/fitbit/v2/types"
)

// UserHeartRateVariabilityIntraday retrieves the Heart Rate Variability (HRV) intraday data for a date range.
// HRV data applies specifically to a user’s “main sleep,” which is the longest single period of time asleep on a given date.
// It measures your HRV rate at various times and returns:
// - Root Mean Square of Successive Differences (rmssd)
// - Low Frequency (LF)
// - High Frequency (HF)
// - Coverage data for a given measurement.
// Rmssd measures short-term variability in your heart rate while asleep.
// LF and HF capture the power in interbeat interval fluctuations within either high frequency or low frequency bands.
// Finally, coverage refers to data completeness in terms of the number of interbeat intervals.
//
// The endDate parameter is optional. When present it returns the summary, day-by-day, from startDate to endDate.
func (c *Client) UserHeartRateVariabilityIntraday(startDate, endDate *time.Time) (ret *types.HeartRateVariabilityIntraday, err error) {
	var res *http.Response
	var sb strings.Builder

	// /1/user/[user-id]/hrv/date/[date]/all.json
	sb.WriteString(fmt.Sprintf("/hrv/date/%s", startDate.Format(types.DateLayout)))
	if endDate != nil && !endDate.IsZero() {
		// /1/user/[user-id]/hrv/date/[startDate]/[endDate]/all.json
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
	ret = &types.HeartRateVariabilityIntraday{}
	err = json.Unmarshal(body, ret)
	return
}
