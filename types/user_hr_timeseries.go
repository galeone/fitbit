// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package types

// /activities/heart/date/%s/%s.json

type HeartRateSeries struct {
	ActivitiesHeart []HeartRateActivities `json:"activities-heart"`
}

type HeartRateTimePointValue struct {
	CustomHeartRateZones []HeartRateZone `json:"customHeartRateZones"`
	HeartRateZones       []HeartRateZone `json:"heartRateZones"`
	RestingHeartRate     int64           `json:"restingHeartRate"`
}

type HeartRateActivities struct {
	DateTime FitbitDate              `json:"dateTime"`
	Value    HeartRateTimePointValue `json:"value"`
}
