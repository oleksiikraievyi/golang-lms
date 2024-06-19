package service

import (
	"cmp"
	"lms/models"
	"sync"
	"sync/atomic"
	"time"

	"slices"

	"github.com/rdleal/intervalst/interval"
)

type DB struct {
	ClientCounter                 *atomic.Uint64
	ClientsMap                    *sync.Map
	ClientsWorkingHoursSearchTree *interval.SearchTree[*models.Client, time.Time]
}

func NewDB() *DB {
	var clientCounter atomic.Uint64

	searchTreeCmpFn := func(t1, t2 time.Time) int {
		switch {
		case t1.After(t2):
			return 1
		case t1.Before(t2):
			return -1
		default:
			return 0
		}
	}

	clientsWorkingHoursSearchTree := interval.NewSearchTree[*models.Client](searchTreeCmpFn)

	return &DB{
		ClientsMap:                    new(sync.Map),
		ClientCounter:                 &clientCounter,
		ClientsWorkingHoursSearchTree: clientsWorkingHoursSearchTree,
	}
}

func (db *DB) GetClients() []*models.Client {
	var clients []*models.Client

	// copying clients from sync.Map to a slice to not have race conditions
	db.ClientsMap.Range(func(key, value any) bool {
		clients = append(clients, value.(*models.Client))
		return true
	})

	// sort by priority asc for convenience
	slices.SortFunc(clients, func(l, r *models.Client) int {
		return cmp.Compare(l.Priority, r.Priority)
	})

	return clients
}

func (db *DB) GetClient(id uint64) *models.Client {
	if client, ok := db.ClientsMap.Load(id); ok {
		return client.(*models.Client)
	}

	return nil
}

func (db *DB) CreateClient(name string, workingHours models.WorkingHours, priority int, leadCapacity int) *models.Client {
	newClient := &models.Client{
		ID:           db.ClientCounter.Load(),
		Name:         name,
		WorkingHours: workingHours,
		Priority:     priority,
		LeadCapacity: leadCapacity,
		LeadCount:    0,
	}

	db.ClientsMap.Store(newClient.ID, newClient)
	db.ClientsWorkingHoursSearchTree.Insert(newClient.WorkingHours.Start, newClient.WorkingHours.End, newClient)
	db.ClientCounter.Add(1)

	return newClient
}

func (db *DB) IncrementLeadCount(id uint64) *models.Client {
	availableClient := db.GetClient(id)

	if availableClient == nil {
		return nil
	}

	availableClient.LeadCount++
	db.ClientsMap.Store(availableClient.ID, availableClient)

	return availableClient
}
