package models

import "time"

type WorkingHours struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type Client struct {
	ID           uint64       `json:"id"`
	Name         string       `json:"name" binding:"required"`
	WorkingHours WorkingHours `json:"working_hours" binding:"required"`
	Priority     int          `json:"priority" binding:"required"`
	LeadCapacity int          `json:"lead_capacity" binding:"required"`
	LeadCount    int          `json:"lead_count" binding:"required"`
}
