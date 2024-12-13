package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"airbnb/database"
	"airbnb/models"
	"airbnb/utils"

	"github.com/gin-gonic/gin"
)

func GetRoomMetrics(c *gin.Context) {
	roomID := c.Param("room_id")
	var room models.Room
	if err := database.DB.First(&room, "id = ?", roomID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}
	monthlyOccupancy, err := utils.CalculateMonthlyOccupancy(room.AvailableDates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed"})
		return
	}

	var availableDates []string
	err = json.Unmarshal(room.AvailableDates, &availableDates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed"})
		return
	}

	now := time.Now()
	next30Days := now.AddDate(0, 0, 30)
	nightlyRates := []map[string]interface{}{}
	rates := []float64{}

	for _, dateStr := range availableDates {
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			continue
		}
		if date.After(now) && date.Before(next30Days) {
			rates = append(rates, room.RatePerNight)
			nightlyRates = append(nightlyRates, map[string]interface{}{
				"date": dateStr,
				"rate": room.RatePerNight,
			})
		}
	}

	avgRate, maxRate, minRate := utils.CalculateRates(rates)

	c.JSON(http.StatusOK, gin.H{
		"room_id":           roomID,
		"occupancy_monthly": monthlyOccupancy,
		"rate_metrics": gin.H{
			"average_rate": avgRate,
			"highest_rate": maxRate,
			"lowest_rate":  minRate,
		},
		"nightly_rates": nightlyRates,
	})
}
