package service

import (
	"cmp"
	"lms/models"
	"slices"
	"time"
)

type ClientService struct {
	db *DB
}

func NewClientService(db *DB) *ClientService {
	return &ClientService{
		db: db,
	}
}

func (s *ClientService) GetClients() []*models.Client {
	return s.db.GetClients()
}

func (s *ClientService) GetClient(id uint64) *models.Client {
	return s.db.GetClient(id)
}

func (s *ClientService) CreateClient(name string, workingHours models.WorkingHours, priority int, leadCapacity int) *models.Client {
	return s.db.CreateClient(name, workingHours, priority, leadCapacity)
}

func (s *ClientService) IncrementLeadCount(id uint64) *models.Client {
	return s.db.IncrementLeadCount(id)
}

func (s *ClientService) FindLead(
	meetingStartTime time.Time,
	meetingEndTime time.Time,
) *models.Client {
	clients := s.db.GetClients()

	if len(clients) == 0 {
		return nil
	}

	var availableClients []*models.Client

	// only take into account clients with not exhausted lead capacity
	for _, client := range clients {
		if client.LeadCapacity > client.LeadCount {
			availableClients = append(availableClients, client)
		}
	}

	if len(availableClients) == 0 {
		return nil
	}

	// find intersection with working hours
	availableClientsWithMatchingTime, ok := s.db.ClientsWorkingHoursSearchTree.AllIntersections(
		meetingStartTime, meetingEndTime,
	)

	if !ok {
		return nil
	}

	// sort clients by priority desc
	slices.SortFunc(availableClientsWithMatchingTime, func(l, r *models.Client) int {
		return cmp.Compare(l.Priority, r.Priority) * -1
	})

	// sort client by remaining lead count asc
	slices.SortFunc(availableClientsWithMatchingTime, func(l, r *models.Client) int {
		return cmp.Compare(l.LeadCapacity-l.LeadCount, r.LeadCapacity-r.LeadCount) * -1
	})

	return availableClientsWithMatchingTime[0]
}
