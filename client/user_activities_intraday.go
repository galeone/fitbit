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

func (c *Client) userActivityIntradayByRange(resource string, startDate, endDate *time.Time) (ret interface{}, err error) {
	var res *http.Response
	hasEndDate := endDate != nil && !endDate.IsZero()

	var path string
	// Same route, but with a period of 1d instead of and end date
	if hasEndDate {
		// /1/user/[user-id]/activities/[resource]/date/[date]/1d/[detail-level]/time/[start-time]/[end-time].json
		path = fmt.Sprintf("/activities/%s/date/%s/1d/1min/time/%s/%s/.json", resource, startDate.Format(types.DateLayout), startDate.Format(types.TimeLayout), endDate.Format(types.TimeLayout))
	} else {
		// /1/user/[user-id]/activities/[resource]/date/[date]/1d/[detail-level].json
		path = fmt.Sprintf("/activities/%s/date/%s/1d/1min.json", resource, startDate.Format(types.DateLayout))
	}
	if res, err = c.req.Get(UserV1(path)); err != nil {
		return
	}
	var body []byte
	if body, err = c.resRead(res); err != nil {
		return
	}
	// https://dev.fitbit.com/build/reference/web-api/intraday/get-activity-intraday-by-date/ (resource)
	switch resource {
	case "calories":
		ret = &types.CaloriesSeriesIntraday{}
	case "distance":
		ret = &types.DistanceSeriesIntraday{}
	case "elevation":
		ret = &types.ElevationSeriesIntraday{}
	case "floors":
		ret = &types.FloorsSeriesIntraday{}
	case "steps":
		ret = &types.StepsSeriesIntraday{}
	default:
		panic(fmt.Sprintf("resouce %s not supported", resource))
	}
	err = json.Unmarshal(body, ret)
	return
}

// UserCaloriesIntraday retrieves the calories intraday time series data on a specific date or 24 hour period.
// The response will include the activity detail minute by minute
// The endDate parameter is optional. When present it should be in a 24 hours range between the start date.
// Minutes are considered.
func (c *Client) UserCaloriesIntraday(startDate, endDate *time.Time) (ret *types.CaloriesSeriesIntraday, err error) {
	var val interface{}
	if val, err = c.userActivityIntradayByRange("calories", startDate, endDate); err != nil {
		return nil, err
	}
	return val.(*types.CaloriesSeriesIntraday), err
}

// UserDistanceIntraday retrieves the calories intraday time series data on a specific date or 24 hour period.
// The response will include the activity detail minute by minute
// The endDate parameter is optional. When present it should be in a 24 hours range between the start date.
// Minutes are considered.
func (c *Client) UserDistanceIntraday(startDate, endDate *time.Time) (ret *types.DistanceSeriesIntraday, err error) {
	var val interface{}
	if val, err = c.userActivityIntradayByRange("distance", startDate, endDate); err != nil {
		return nil, err
	}
	return val.(*types.DistanceSeriesIntraday), err
}

// UserElevationIntraday retrieves the calories intraday time series data on a specific date or 24 hour period.
// The response will include the activity detail minute by minute
// The endDate parameter is optional. When present it should be in a 24 hours range between the start date.
// Minutes are considered.
func (c *Client) UserElevationIntraday(startDate, endDate *time.Time) (ret *types.ElevationSeriesIntraday, err error) {
	var val interface{}
	if val, err = c.userActivityIntradayByRange("elevation", startDate, endDate); err != nil {
		return nil, err
	}
	return val.(*types.ElevationSeriesIntraday), err
}

// UserFloorsIntraday retrieves the calories intraday time series data on a specific date or 24 hour period.
// The response will include the activity detail minute by minute
// The endDate parameter is optional. When present it should be in a 24 hours range between the start date.
// Minutes are considered.
func (c *Client) UserFloorsIntraday(startDate, endDate *time.Time) (ret *types.FloorsSeriesIntraday, err error) {
	var val interface{}
	if val, err = c.userActivityIntradayByRange("floors", startDate, endDate); err != nil {
		return nil, err
	}
	return val.(*types.FloorsSeriesIntraday), err
}

// UserStepsIntraday retrieves the calories intraday time series data on a specific date or 24 hour period.
// The response will include the activity detail minute by minute
// The endDate parameter is optional. When present it should be in a 24 hours range between the start date.
// Minutes are considered.
func (c *Client) UserStepsIntraday(startDate, endDate *time.Time) (ret *types.StepsSeriesIntraday, err error) {
	var val interface{}
	if val, err = c.userActivityIntradayByRange("steps", startDate, endDate); err != nil {
		return nil, err
	}
	return val.(*types.StepsSeriesIntraday), err
}
