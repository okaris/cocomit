// SPDX-License-Identifier: MIT

package processor

import (
	"math"
)

var CocomoProjectType = "organic"

var projectType = map[string][]float64{
	"organic":       {2.4, 1.05, 2.5, 0.38},
	"semi-detached": {3.0, 1.12, 2.5, 0.35},
	"embedded":      {3.6, 1.20, 2.5, 0.32},
}

func EstimateCost(effortApplied float64, averageWage int64, overhead float64) float64 {
	return effortApplied * float64(averageWage/12) * overhead
}

func EstimateEffort(sloc int64, eaf float64) float64 {
	return projectType[CocomoProjectType][0] *
		math.Pow(float64(sloc)/1000, projectType[CocomoProjectType][1]) * eaf
}

func EstimateScheduleMonths(effortApplied float64) float64 {
	return projectType[CocomoProjectType][2] *
		math.Pow(effortApplied, projectType[CocomoProjectType][3])
}
