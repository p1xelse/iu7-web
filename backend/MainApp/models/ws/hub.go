package models

import "writesend/MainApp/models"

type Subscription struct {
	conn *connection
	room uint64
}

type Hub struct {
	rooms      map[uint64]map[*connection]bool
	broadcast  chan models.Message
	register   chan Subscription
	unregister chan Subscription
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan models.Message),
		register:   make(chan Subscription),
		unregister: make(chan Subscription),
		rooms:      make(map[uint64]map[*connection]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case s := <-h.register:
			connections := h.rooms[s.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.room] = connections
			}
			h.rooms[s.room][s.conn] = true
		case s := <-h.unregister:
			connections := h.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}
		case m := <-h.broadcast:
			connections := h.rooms[m.DialogID]
			for c := range connections {
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.DialogID)
					}
				}
			}
		}
	}
}
