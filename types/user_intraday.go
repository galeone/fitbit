// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package types

// /activities/%s/date/%s/%s/%s/time/%s/%s.json

type CaloriesSeriesIntraday struct {
	CaloriesSeries
	CaloriesIntraday TimeSeriesIntraday `json:"activities-calories-intraday"`
}

type TimeSeriesIntraday struct {
	Dataset         []Dataset `json:"dataset"`
	DatasetInterval int64     `json:"datasetInterval"`
	DatasetType     string    `json:"datasetType"`
}

type Dataset struct {
	Level int64      `json:"level,omitempty"`
	Mets  int64      `json:"mets,omitempty"`
	Time  FitbitTime `json:"time"`
	Value float64    `json:"value"`
}

// /activities/%s/date/%s/%s/%s/time/%s/%s.json

type DistanceSeriesIntraday struct {
	DistanceSeries
	DistanceIntraday TimeSeriesIntraday `json:"activities-distance-intraday"`
}

// /activities/%s/date/%s/%s/%s/time/%s/%s.json

type ElevationSeriesIntraday struct {
	ElevationSeries
	ElevationIntraday TimeSeriesIntraday `json:"activities-elevation-intraday"`
}

// /activities/%s/date/%s/%s/%s/time/%s/%s.json

type FloorsSeriesIntraday struct {
	FloorsSeries
	FloorsIntraday TimeSeriesIntraday `json:"activities-floors-intraday"`
}

// /activities/%s/date/%s/%s/%s/time/%s/%s.json

type StepsSeriesIntraday struct {
	StepsSeries
	StepsIntraday TimeSeriesIntraday `json:"activities-steps-intraday"`
}

// /br/date/%s/%s/all.json

type BreathingRateIntraday struct {
	Br []BreathingRateTimePointIntraday `json:"br"`
}

type BreathingRateTimePointIntraday struct {
	DateTime FitbitDate                   `json:"dateTime"`
	Value    BreathingRateIntradaySummary `json:"value"`
}

type BreathingRateIntradaySummary struct {
	DeepSleepSummary  BreathingRateValue `json:"deepSleepSummary"`
	FullSleepSummary  BreathingRateValue `json:"fullSleepSummary"`
	LightSleepSummary BreathingRateValue `json:"lightSleepSummary"`
	RemSleepSummary   BreathingRateValue `json:"remSleepSummary"`
}

// /activities/heart/date/%s/%s/%s/time/%s/%s.json

type HeartRateIntraday struct {
	HeartRateSeries
	HeartRateIntraday TimeSeriesIntraday `json:"activities-heart-intraday"`
}

// /hrv/date/%s/%s/all.json

type HeartRateVariabilityIntraday struct {
	Hrv []HeartRateVariabilityTimeStepIntraday `json:"hrv"`
}

type HeartRateVariabilityValueIntraday struct {
	Coverage float64 `json:"coverage"`
	Hf       float64 `json:"hf"`
	Lf       float64 `json:"lf"`
	Rmssd    float64 `json:"rmssd"`
}

type HeartRateVariabilityTimeStepIntraday struct {
	DateTime FitbitDate                   `json:"dateTime"`
	Minutes  []HeartRateVariabilityMinute `json:"minutes"`
}

type HeartRateVariabilityMinute struct {
	Minute FitbitDate                        `json:"minute"`
	Value  HeartRateVariabilityValueIntraday `json:"value"`
}

// /spo2/date/%s/%s/all.json

type OxygenSaturationIntraday struct {
	DateTime FitbitDate               `json:"dateTime"`
	Minutes  []OxygenSaturationMinute `json:"minutes"`
}

type OxygenSaturationMinute struct {
	Minute FitbitDateTime `json:"minute"`
	Value  float64        `json:"value"`
}
