package service

import (
	"lms/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindLead(t *testing.T) {
	testCases := []struct {
		name               string
		clients            []*models.Client
		startTime          time.Time
		endTime            time.Time
		expectedClientName string
	}{
		{
			name:      "no clients, return nil",
			clients:   make([]*models.Client, 0),
			startTime: time.Now().UTC(),
			endTime:   time.Now().UTC().Add(time.Hour),
		},
		{
			name: "all clients LeadCapacityExhausted, return nil",
			clients: []*models.Client{
				{
					Name:         "client1",
					WorkingHours: models.WorkingHours{Start: time.Now().UTC(), End: time.Now().UTC().Add(time.Hour)},
					Priority:     1,
					LeadCapacity: 1,
					LeadCount:    1,
				},
				{
					Name:         "client2",
					WorkingHours: models.WorkingHours{Start: time.Now().UTC(), End: time.Now().UTC().Add(time.Hour)},
					Priority:     1,
					LeadCapacity: 3,
					LeadCount:    3,
				},
			},
			startTime: time.Now().UTC(),
			endTime:   time.Now().UTC().Add(time.Hour),
		},
		// {
		// 	name: "find client with LeadCapacity not exhausted, based on priority, return client2",
		// 	clients: []*models.Client{
		// 		{
		// 			Name:         "client1",
		// 			WorkingHours: models.WorkingHours{Start: time.Now().UTC(), End: time.Now().UTC().Add(time.Hour)},
		// 			Priority:     1,
		// 			LeadCapacity: 3,
		// 			LeadCount:    0,
		// 		},
		// 		{
		// 			Name:         "client2",
		// 			WorkingHours: models.WorkingHours{Start: time.Now().UTC(), End: time.Now().UTC().Add(time.Hour)},
		// 			Priority:     2,
		// 			LeadCapacity: 3,
		// 			LeadCount:    0,
		// 		},
		// 		{
		// 			Name:         "client3",
		// 			WorkingHours: models.WorkingHours{Start: time.Now().UTC(), End: time.Now().UTC().Add(time.Hour)},
		// 			Priority:     2,
		// 			LeadCapacity: 1,
		// 			LeadCount:    0,
		// 		},
		// 	},
		// 	startTime:          time.Now().UTC(),
		// 	endTime:            time.Now().UTC().Add(time.Hour),
		// 	expectedClientName: "client2",
		// },
		// {
		// 	name: "find client with LeadCapacity not exhausted, based on lead capacity, return client1",
		// 	clients: []*models.Client{
		// 		{
		// 			Name:         "client1",
		// 			WorkingHours: models.WorkingHours{Start: time.Now().UTC(), End: time.Now().UTC().Add(time.Hour)},
		// 			Priority:     2,
		// 			LeadCapacity: 3,
		// 			LeadCount:    0,
		// 		},
		// 		{
		// 			Name:         "client2",
		// 			WorkingHours: models.WorkingHours{Start: time.Now().UTC(), End: time.Now().UTC().Add(time.Hour)},
		// 			Priority:     2,
		// 			LeadCapacity: 2,
		// 			LeadCount:    1,
		// 		},
		// 		{
		// 			Name:         "client3",
		// 			WorkingHours: models.WorkingHours{Start: time.Now().AddDate(0, -2, 0).UTC(), End: time.Now().UTC().Add(time.Hour)},
		// 			Priority:     2,
		// 			LeadCapacity: 2,
		// 			LeadCount:    2,
		// 		},
		// 	},
		// 	startTime:          time.Now().UTC(),
		// 	endTime:            time.Now().UTC().Add(time.Hour),
		// 	expectedClientName: "client1",
		// },
	}

	for _, testCase := range testCases {
		s := NewClientService(NewDB())

		for _, client := range testCase.clients {
			newClient := s.CreateClient(client.Name, client.WorkingHours, client.Priority, client.LeadCapacity)
			for i := 0; i < client.LeadCount; i++ {
				s.IncrementLeadCount(newClient.ID)
			}
		}

		res := s.FindLead(testCase.startTime, testCase.endTime)

		if testCase.expectedClientName == "" {
			assert.Nil(t, res)
		} else {
			assert.Equal(t, testCase.expectedClientName, res.Name)
		}
	}
}
