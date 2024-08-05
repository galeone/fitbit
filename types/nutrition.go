// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package types

type FoodLog struct {
	IsFavorite  bool        `json:"isFavorite"`
	LogDate     FitbitDate  `json:"logDate"`
	LogId       int64       `json:"logId"`
	LoggedFood  LoggedFood  `json:"loggedFood"`
	Nutritional Nutritional `json:"nutritionalValues"`
}

type FoodLogList struct {
	Foods []FoodLog `json:"foods"`
	Goals struct {
		Calories int64 `json:"calories"`
	} `json:"goals"`
	Summary Nutritional `json:"summary"`
}

type LoggedFood struct {
	AccessLevel string         `json:"accessLevel"`
	Amount      float64        `json:"amount"`
	Brand       string         `json:"brand"`
	Calories    int64          `json:"calories"`
	FoodID      int64          `json:"foodId"`
	Locale      string         `json:"locale"`
	MealType    int64          `json:"mealTypeId"`
	Name        string         `json:"name"`
	Unit        LoggedFoodUnit `json:"unit"`
	Units       []int64        `json:"units"`
}

type LoggedFoodUnit struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Plural string `json:"plural"`
}

type Nutritional struct {
	Calories int64   `json:"calories"`
	Carbs    float64 `json:"carbs"`
	Fat      float64 `json:"fat"`
	Fiber    float64 `json:"fiber"`
	Protein  float64 `json:"protein"`
	Sodium   float64 `json:"sodium"`
	Water    float64 `json:"water,omitempty"`
}
