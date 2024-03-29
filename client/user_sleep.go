// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/galeone/fitbit/v2/types"
)

// UserSleepLogList retrieves a list of a user's sleep log entries before or after a given date,
// and specifying offset, limit and sort order.
//
// GET: /1.2/user/[user-id]/sleep/list.json
func (c *Client) UserSleepLogList(pagination *types.Pagination) (ret *types.SleepLogs, err error) {
	var sb strings.Builder
	sb.WriteString("/sleep/list.json?sort=")
	sb.WriteString(pagination.Sort)
	sb.WriteString("&offset=")
	sb.WriteString(strconv.Itoa(int(pagination.Offset)))
	sb.WriteString("&limit=")
	sb.WriteString(strconv.Itoa(int(pagination.Limit)))

	if !pagination.BeforeDate.IsZero() {
		sb.WriteString("&beforeDate=")
		sb.WriteString(url.QueryEscape(pagination.BeforeDate.Format(types.DateTimeLayout)))
	}

	if !pagination.AfterDate.IsZero() {
		sb.WriteString("&afterDate=")
		sb.WriteString(url.QueryEscape(pagination.AfterDate.Format(types.DateTimeLayout)))
	}

	path := UserV12(sb.String())

	var res *http.Response
	if res, err = c.req.Get(path); err != nil {
		return
	}
	var body []byte
	if body, err = c.resRead(res); err != nil {
		return
	}
	ret = &types.SleepLogs{}
	err = json.Unmarshal(body, ret)
	return
}

// UserSleepLog retrieves a list of a user's sleep log entries for a date range.
//
// The endDate parameter is optional. When present it returns the summary, day-by-day, from startDate to endDate.
func (c *Client) UserSleepLog(startDate, endDate *time.Time) (ret *types.SleepLogs, err error) {
	var res *http.Response
	var sb strings.Builder

	// /1/user/[user-id]/sleep/date/[date].json
	sb.WriteString(fmt.Sprintf("/sleep/date/%s", startDate.Format(types.DateLayout)))
	if endDate != nil && !endDate.IsZero() {
		// /1/user/[user-id]/sleep/date/[start-date]/[end-date].json
		sb.WriteString(fmt.Sprintf("/%s", endDate.Format(types.DateLayout)))
	}
	sb.WriteString(".json")
	if res, err = c.req.Get(UserV12(sb.String())); err != nil {
		return
	}
	var body []byte
	if body, err = c.resRead(res); err != nil {
		return
	}
	ret = &types.SleepLogs{}
	err = json.Unmarshal(body, ret)
	return
}

// UserSleepGoalReport retrieves the user's current sleep goal.
func (c *Client) UserSleepGoalReport() (ret *types.SleepGoalReport, err error) {
	var res *http.Response
	if res, err = c.req.Get(UserV12("/sleep/goal.json")); err != nil {
		return
	}
	var body []byte
	if body, err = c.resRead(res); err != nil {
		return
	}
	ret = &types.SleepGoalReport{}
	err = json.Unmarshal(body, ret)
	return
}
