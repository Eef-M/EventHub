package controllers

import (
	"net/http"

	"github.com/Eef-M/EventHub/backend/initializers"
	"github.com/Eef-M/EventHub/backend/models"
	"github.com/gin-gonic/gin"
)

func OrganizerDashboard(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	organizer := user.(models.User)

	var totalEvents int64
	var totalRegistrations int64
	var pendingApprovals int64
	var totalFeedbacks int64
	var recentRegistrations []models.EventRegistration

	// Total Events
	initializers.DB.Model(&models.Event{}).
		Where("organizer_id = ?", organizer.ID).
		Count(&totalEvents)

		// Total Registrations
	initializers.DB.Table("event_registrations").
		Joins("JOIN events ON event_registrations.event_id = events.id").
		Where("events.organizer_id = ?", organizer.ID).
		Count(&totalRegistrations)

		// Pending Approvals
	initializers.DB.Table("event_registrations").
		Joins("JOIN events ON event_registrations.event_id = events.id").
		Where("events.organizer_id = ? AND event_registrations.status = ?", organizer.ID, "pending").
		Count(&pendingApprovals)

		// Feedback Received
	initializers.DB.Table("event_feedbacks").
		Joins("JOIN events ON event_feedbacks.event_id = events.id").
		Where("events.organizer_id = ?", organizer.ID).
		Count(&totalFeedbacks)

	// Recent Registrations (limit 5)
	initializers.DB.Preload("Event").
		Preload("User").
		Joins("JOIN events ON event_registrations.event_id = events.id").
		Where("events.organizer_id = ?", organizer.ID).
		Order("event_registrations.created_at DESC").
		Limit(5).
		Find(&recentRegistrations)

	c.JSON(http.StatusOK, gin.H{
		"total_events":         totalEvents,
		"total_registrations":  totalRegistrations,
		"pending_approvals":    pendingApprovals,
		"feedback_received":    totalFeedbacks,
		"recent_registrations": recentRegistrations,
	})
}
