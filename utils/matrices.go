package utils

import (
	"encoding/json"
	"time"
)

func CalculateMonthlyOccupancy(datesJSON []byte) (map[string]float64, error) {
	var dates []string
	err := json.Unmarshal(datesJSON, &dates)
	if err != nil {
		return nil, err
	}

	monthlyOccupancy := make(map[string]int)
	now := time.Now()

	for i := 0; i < 5; i++ {
		monthStart := now.AddDate(0, i, 0)
		monthEnd := monthStart.AddDate(0, 1, -1)

		for _, dateStr := range dates {
			date, _ := time.Parse("2006-01-02", dateStr)
			if date.After(monthStart.AddDate(0, 0, -1)) && date.Before(monthEnd.AddDate(0, 0, 1)) {
				monthKey := monthStart.Format("2006-01")
				monthlyOccupancy[monthKey]++
			}
		}
	}

	monthlyPercentages := make(map[string]float64)
	for month, availableDays := range monthlyOccupancy {
		totalDays := time.Date(now.Year(), now.Month()+time.Month(len(monthlyOccupancy)), 0, 0, 0, 0, 0, now.Location()).Day()
		monthlyPercentages[month] = float64(availableDays) / float64(totalDays) * 100
	}

	return monthlyPercentages, nil
}

func CalculateRates(rates []float64) (average, highest, lowest float64) {
	if len(rates) == 0 {
		return 0, 0, 0
	}

	highest = rates[0]
	lowest = rates[0]
	total := 0.0

	for _, rate := range rates {
		if rate > highest {
			highest = rate
		}
		if rate < lowest {
			lowest = rate
		}
		total += rate
	}

	average = total / float64(len(rates))
	return average, highest, lowest
}
