// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package types

// /spo2/date/%s.json

type OxygenSaturation struct {
	DateTime FitbitDate            `json:"dateTime"`
	Value    OxygenSaturationValue `json:"value"`
}

type OxygenSaturationValue struct {
	Avg float64 `json:"avg"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type OxygenSaturations []OxygenSaturation
