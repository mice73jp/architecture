package service2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/net/context"
)

var src = rand.NewSource(time.Now().UnixNano())

// MockGateway ...
type MockGateway struct {
	db *MockDB
}

// MockDB ...
type MockDB struct {
	mu   sync.RWMutex
	data map[int64]Person
}

// NewMockDB ...
func NewMockDB() *MockDB {
	return &MockDB{data: make(map[int64]Person)}
}

// NewMockGateway ...
func NewMockGateway(db *MockDB) *MockGateway {
	return &MockGateway{db}
}

// RegisterPerson ...
func (r *MockGateway) RegisterPerson(ctx context.Context, p Person) (Person, error) {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()

	var id int64
	for {
		id = src.Int63()
		_, ok := r.db.data[id]
		if !ok {
			break
		}
	}

	p.ID = id
	r.db.data[p.ID] = p

	return p, nil
}

// GetPersonByID ...
func (r *MockGateway) GetPersonByID(ctx context.Context, id int64) (Person, error) {
	r.db.mu.Lock()
	defer r.db.mu.Unlock()

	if p, ok := r.db.data[id]; ok {
		return p, nil
	}

	return Person{}, fmt.Errorf("Person not found : id: %d", id)
}
