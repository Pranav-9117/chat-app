package room

import (
	"log"
	"sync"
)

type Manager struct{

	rooms map[string]*Room
	mu sync.RWMutex
}

func NewManager()*Manager{
	return &Manager{
		rooms: make(map[string]*Room),

	}
}

func(m *Manager)GetRoom(id string)*Room{

	m.mu.Lock()
	defer m.mu.Unlock()

	room,exists := m.rooms[id]

	if !exists{
		log.Println("Creating new room",id)

		room = NewRoom(id)
		m.rooms[id]=room
		go room.Run()
	}
	return room
}