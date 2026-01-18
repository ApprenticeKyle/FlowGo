package devops

import (
	"sync"
)

// LogEvent Log event structure
type LogEvent struct {
	Type       string `json:"type"`        // "log" or "status"
	PipelineID uint64 `json:"pipeline_id"` // Related pipeline ID
	Content    string `json:"content"`     // Log content or new status
	Timestamp  int64  `json:"timestamp"`
}

// LogBroadcaster manages SSE clients
type LogBroadcaster struct {
	clients map[chan LogEvent]bool
	mu      sync.RWMutex
}

func NewLogBroadcaster() *LogBroadcaster {
	return &LogBroadcaster{
		clients: make(map[chan LogEvent]bool),
	}
}

// Subscribe adds a new client
func (b *LogBroadcaster) Subscribe() chan LogEvent {
	ch := make(chan LogEvent, 100) // Buffer to prevent blocking
	b.mu.Lock()
	defer b.mu.Unlock()
	b.clients[ch] = true
	return ch
}

// Unsubscribe removes a client
func (b *LogBroadcaster) Unsubscribe(ch chan LogEvent) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.clients[ch]; ok {
		delete(b.clients, ch)
		close(ch)
	}
}

// Broadcast sends an event to all clients
func (b *LogBroadcaster) Broadcast(event LogEvent) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for ch := range b.clients {
		select {
		case ch <- event:
		default:
			// Client too slow, skip to avoid blocking broadcaster
		}
	}
}

// SendLog helper
func (b *LogBroadcaster) SendLog(pipelineID uint64, content string) {
	b.Broadcast(LogEvent{
		Type:       "log",
		PipelineID: pipelineID,
		Content:    content,
		Timestamp:  0, // Frontend receives it real-time
	})
}

// SendStatus helper
func (b *LogBroadcaster) SendStatus(pipelineID uint64, status string) {
	b.Broadcast(LogEvent{
		Type:       "status",
		PipelineID: pipelineID,
		Content:    status,
		Timestamp:  0,
	})
}
