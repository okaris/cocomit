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

func EstimateCost(effortApplied float64, hourlyWage float64, overhead float64) float64 {
	hours := effortApplied * 152 // 152 hours per person-month
	return hours * hourlyWage * overhead
}

func EstimateEffort(sloc int64, eaf float64) float64 {
	// Assume 1 developer writing 50 LOC/hour
	// 1 person-month = 152 hours
	hours := float64(sloc) / 50.0
	return (hours / 152.0) * eaf
}

func EstimateScheduleMonths(effortApplied float64) float64 {
	return projectType[CocomoProjectType][2] *
		math.Pow(effortApplied, projectType[CocomoProjectType][3])
}
