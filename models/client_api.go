package models

import "time"

type CreateClientRequest struct {
	Name         string       `json:"name" binding:"required"`
	WorkingHours WorkingHours `json:"working_hours" binding:"required"`
	Priority     int          `json:"priority" binding:"required"`
	LeadCapacity int          `json:"lead_capacity" binding:"required"`
}

type AssignLeadRequest struct {
	MeetingStartTime time.Time `json:"meeting_start_time" binding:"required"`
	MeetingEndTime   time.Time `json:"meeting_end_time" binding:"required"`
}

type Error struct {
	Error string `json:"error"`
}
